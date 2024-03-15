package btc

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

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
	case chaincfg.TestNet3Params.Name:
		networkParams = &chaincfg.TestNet3Params
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
	stakerPubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(stakerPubKeyStr, "0x"))
	if err != nil {
		return errors.New("public key should hex string")
	}

	stakerPubKey, err := btcec.ParsePubKey(stakerPubKeyBytes)
	if err != nil {
		return errors.New("invalid staker public key")
	}

	sigBytes, err := hex.DecodeString(strings.TrimPrefix(sigHexStr, "0x"))
	if err != nil {
		return errors.New("signature should be a hex string")
	}

	signature, err := ecdsa.ParseSignature(sigBytes)
	if err != nil {
		return errors.New("invalid reward receiver signature")
	}

	message := AssembleRewardSignatureMessage(rewardReceiver, timestamp)
	messageHash := chainhash.DoubleHashB([]byte(message))

	if !signature.Verify(messageHash, stakerPubKey) {
		return errors.New("reward address signature verify failed")
	}

	return nil
}

func (c *Client) CheckTxID(txID string) error {
	_, err := chainhash.NewHashFromStr(strings.TrimPrefix(txID, "0x"))
	if err != nil {
		return errors.New("invalid tx id")
	}
	return nil
}

func (c *Client) UpdateStakeRecordFinalizedStatus(stakeRecords []*types.StakeVerificationParam) ([]types.StakeRecordStatus, error) {
	recordStatuses := make([]types.StakeRecordStatus, len(stakeRecords))
	rawTxFutures := make([]*rpcclient.FutureGetRawTransactionVerboseResult, len(stakeRecords))
	for i, record := range stakeRecords {
		txHash, err := chainhash.NewHashFromStr(record.TxID)
		if err != nil {
			recordStatuses[i] = types.Mismatch
			rawTxFutures[i] = nil
			continue
		}
		rawTxFuture := c.rpcClient.GetRawTransactionVerboseAsync(txHash)
		rawTxFutures[i] = &rawTxFuture
	}

	for i := 0; i < len(rawTxFutures); i++ {
		future := rawTxFutures[i]
		if future == nil {
			continue
		}
		txRes, err := future.Receive()
		if err != nil {
			recordStatuses[i] = stakeRecords[i].FinalizedStatus
			continue
		}
		if stakeRecords[i].FinalizedStatus == types.Mismatch || stakeRecords[i].FinalizedStatus == types.TxFinalized {
			// TODO should not come here. Mismatch/Finalized tx record should't been checked again.
			recordStatuses[i] = stakeRecords[i].FinalizedStatus
			continue
		}
		if stakeRecords[i].FinalizedStatus == types.TxPending {
			err = c.CheckStake(txRes, stakeRecords[i].StakerPublicKey, stakeRecords[i].Amount, stakeRecords[i].Duration)
			if err != nil {
				recordStatuses[i] = types.Mismatch
				continue
			}
		}

		if txRes.Confirmations >= TxFinalizedConfirmations { //nolint
			recordStatuses[i] = types.TxFinalized
		} else if txRes.Confirmations == 0 {
			recordStatuses[i] = types.TxPending
		} else {
			recordStatuses[i] = types.TxIncluded
		}
	}

	return recordStatuses, nil
}

func (c *Client) CheckStake(tx *btcjson.TxRawResult, stakerPubKeyStr string, amount uint64, duration uint64) error {
	if len(tx.Vout) != 1 {
		return errors.New("stake tx should has 1 out")
	}

	stakerPubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(stakerPubKeyStr, "0x"))
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
