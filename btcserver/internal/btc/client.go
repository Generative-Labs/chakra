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
)

// Transaction requires the number of Confirmations for finalization
const TxFinalizedConfirmations = 6

type Client struct {
	rpcClient     *rpcclient.Client
	networkParams *chaincfg.Params
}

type StakeRecord struct {
	TxID            string
	StakerPubKeyStr string
	Amount          uint64
	Duration        int64
	Status          StakeRecordStatus
}

type StakeRecordStatus int

const (
	// TxPending defines the StakeRecord status where the Tx has not been included in a block yet.
	TxPending StakeRecordStatus = iota
	// TxIncluded defines the StakeRecord status where the Tx has been included in a block.
	TxIncluded
	// TxFinalized defines the StakeRecord status where the Tx has been confirmed.
	TxFinalized
	// Mismatch defines the StakeRecord status where the Tx in the record does not match the content.
	Mismatch
)

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

func (c *Client) CheckRewardReceiverSignature(stakerPubKeyStr, rewardReceiver,
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

	message := assembleRewardSignatureMessage(rewardReceiver, timestamp)
	messageHash := chainhash.DoubleHashB([]byte(message))

	if !signature.Verify(messageHash, stakerPubKey) {
		return errors.New("signature verify failed")
	}

	return nil
}

func (c *Client) CheckStakeRecords(stakeRecords []StakeRecord) ([]StakeRecordStatus, error) {
	var recordStatuses []StakeRecordStatus
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
			recordStatuses = append(recordStatuses, stakeRecords[i].Status)
			continue
		}
		if stakeRecords[i].Status == Mismatch || stakeRecords[i].Status == TxFinalized {
			// TODO should not come here. Mismatch/Finalized tx record should't been checked again.
			recordStatuses = append(recordStatuses, stakeRecords[i].Status)
			continue
		}

		if stakeRecords[i].Status == TxPending {
			err = c.CheckStake(txRes, stakeRecords[i].StakerPubKeyStr, stakeRecords[i].Amount, stakeRecords[i].Duration)
			if err != nil {
				recordStatuses = append(recordStatuses, Mismatch)
				continue
			}
		}

		if txRes.Confirmations >= TxFinalizedConfirmations { //nolint
			recordStatuses = append(recordStatuses, TxFinalized)
		} else if txRes.Confirmations == 0 {
			recordStatuses = append(recordStatuses, TxPending)
		} else {
			recordStatuses = append(recordStatuses, TxIncluded)
		}

	}

	return recordStatuses, nil
}

func (c *Client) CheckStake(tx *btcjson.TxRawResult, stakerPubKeyStr string, amount uint64, duration int64) error {
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

func (c *Client) calculateAddressRedeemScriptHash(stakerPubKey *secp256k1.PublicKey, duration int64) (*btcutil.AddressScriptHash, error) {
	builder := txscript.NewScriptBuilder()
	builder.AddInt64(duration)
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

func assembleRewardSignatureMessage(rewardReceiveAddress string, timestamp int32) string {
	return fmt.Sprintf("%s/%d", rewardReceiveAddress, timestamp)
}
