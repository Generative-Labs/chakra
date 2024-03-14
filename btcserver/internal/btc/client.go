package btc

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/generativelabs/btcserver/internal/types"
)

// Transaction requires the number of Confirmations for finalization
const TxFinalizedConfirmations = 6

type Client struct {
	rpcClient     *rpcclient.Client
	networkParams *chaincfg.Params
}

func NewClient(config Config) (*Client, error) {
	rpcClient, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         config.RPCHost,
		User:         config.RPCUser,
		Pass:         config.RPCPass,
		DisableTLS:   config.DisableTLS,
		HTTPPostMode: true,
	}, nil)
	if err != nil {
		return nil, err
	}

	var networkParams *chaincfg.Params

	switch config.NetworkName {
	case chaincfg.RegressionNetParams.Name:
		networkParams = &chaincfg.RegressionNetParams
	case chaincfg.SigNetParams.Name:
		networkParams = &chaincfg.SigNetParams
	case chaincfg.MainNetParams.Name:
		networkParams = &chaincfg.MainNetParams
	default:
		return nil, errors.New("unknown network")
	}

	client := Client{
		rpcClient:     rpcClient,
		networkParams: networkParams,
	}

	return &client, nil
}

func (c *Client) CheckRewardAddressSignature(stakerPubKeyStr, rewardReceiver,
	sigHexStr string, timestamp int32,
) error {
	stakerPubKeyBytes, err := hex.DecodeString(stakerPubKeyStr)
	if err != nil {
		return err
	}

	stakerPubKey, err := btcec.ParsePubKey(stakerPubKeyBytes)
	if err != nil {
		return err
	}

	sigBytes, err := hex.DecodeString(sigHexStr)
	if err != nil {
		return err
	}

	signature, err := ecdsa.ParseSignature(sigBytes)
	if err != nil {
		return err
	}

	message := AssembleRewardSignatureMessage(rewardReceiver, timestamp)
	messageHash := chainhash.DoubleHashB([]byte(message))

	if !signature.Verify(messageHash, stakerPubKey) {
		return errors.New("reward receiver signature verify failed")
	}

	return nil
}

func (c *Client) CheckStakeRecords(stakeRecords []*types.StakeVerificationParam) ([]types.StakeRecordStatus, error) {
	var recordStatuses []types.StakeRecordStatus
	var rawTxFutures []rpcclient.FutureGetRawTransactionVerboseResult
	for _, record := range stakeRecords {
		txHash, err := chainhash.NewHashFromStr(record.TxID)
		if err != nil {
			return nil, err
		}
		rawTxFuture := c.rpcClient.GetRawTransactionVerboseAsync(txHash)
		rawTxFutures = append(rawTxFutures, rawTxFuture)
	}

	for i, future := range rawTxFutures {
		txRes, err := future.Receive()
		if err != nil {
			recordStatuses = append(recordStatuses, stakeRecords[i].FinalizedStatus)
			continue
		}
		if stakeRecords[i].FinalizedStatus == types.Mismatch || stakeRecords[i].FinalizedStatus == types.TxFinalized {
			// TODO should not come here. Mismatch/Finalized tx record should't been checked again.
			recordStatuses = append(recordStatuses, stakeRecords[i].FinalizedStatus)
			continue
		}

		if stakeRecords[i].FinalizedStatus == types.TxPending {
			err = c.CheckStake(txRes, stakeRecords[i].StakerPubKey, stakeRecords[i].Amount, stakeRecords[i].Duration)
			if err != nil {
				recordStatuses = append(recordStatuses, types.Mismatch)
				continue
			}
		}

		if txRes.Confirmations >= TxFinalizedConfirmations { //nolint
			recordStatuses = append(recordStatuses, types.TxFinalized)
		} else if txRes.Confirmations == 0 {
			recordStatuses = append(recordStatuses, types.TxPending)
		} else {
			recordStatuses = append(recordStatuses, types.TxIncluded)
		}

	}

	return recordStatuses, nil
}

func (c *Client) CheckStake(tx *btcjson.TxRawResult, stakerPubKeyStr string, amount uint64, duration uint64) error {
	if len(tx.Vout) != 1 {
		return errors.New("stake tx should has 1 out")
	}

	stakerPubKeyBytes, err := hex.DecodeString(stakerPubKeyStr)
	if err != nil {
		return err
	}

	stakerPubKey, err := btcec.ParsePubKey(stakerPubKeyBytes)
	if err != nil {
		return err
	}

	redeemAddressScriptHash, err := c.calculateAddressRedeemScriptHash(stakerPubKey, duration)
	if err != nil {
		return err
	}

	redeemP2SH, err := txscript.PayToAddrScript(redeemAddressScriptHash)
	if err != nil {
		return err
	}

	redeemP2SHHex := hex.EncodeToString(redeemP2SH)

	btcAmount := btcutil.Amount(amount).ToBTC()

	for _, out := range tx.Vout {
		if out.ScriptPubKey.Hex != redeemP2SHHex {
			continue
		}
		if out.ScriptPubKey.Type != "scripthash" {
			continue
		}

		if out.Value == btcAmount {
			return nil
		}
	}

	return errors.New("verify stake tx failed")
}

func (c *Client) calculateAddressRedeemScriptHash(stakerPubKey *secp256k1.PublicKey, duration uint64) (*btcutil.AddressScriptHash, error) {
	builder := txscript.NewScriptBuilder()
	builder.AddInt64(int64(duration))
	builder.AddOp(txscript.OP_CHECKSEQUENCEVERIFY)
	builder.AddOp(txscript.OP_DROP)
	builder.AddData(stakerPubKey.SerializeCompressed())
	builder.AddOp(txscript.OP_CHECKSIG)

	csvScript, err := builder.Script()
	if err != nil {
		return nil, err
	}

	addressHash, err := btcutil.NewAddressScriptHash(csvScript, c.networkParams)
	if err != nil {
		return nil, err
	}

	return addressHash, nil
}

func AssembleRewardSignatureMessage(rewardReceiveAddress string, timestamp int32) string {
	return fmt.Sprintf("%s/%d", rewardReceiveAddress, timestamp)
}
