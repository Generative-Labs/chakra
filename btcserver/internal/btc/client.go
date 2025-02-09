package btc

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/generativelabs/btcserver/internal/types"
	"github.com/rs/zerolog/log"
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
	sigBase64Str string, timestamp int64,
) error {
	sigB, err := base64.StdEncoding.DecodeString(sigBase64Str)
	if err != nil {
		return errors.New("signature should be compress as base64")
	}

	msgH := AssembleRewardSignatureMessage(rewardReceiver, timestamp)

	recoverPK, ok, err := ecdsa.RecoverCompact(sigB, msgH)
	if err != nil || !ok {
		return errors.New("invalid signature")
	}

	addressPK, err := btcutil.NewAddressPubKey(recoverPK.SerializeCompressed(), c.networkParams)
	if err != nil {
		return err
	}

	if addressPK.String() != strings.TrimPrefix(stakerPubKeyStr, "0x") {
		return errors.New("mismatch signature")
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

func (c *Client) UpdateStakeRecords(stakeRecords []*types.StakeVerificationParam) ([]types.StakeRecordUpdates, error) {
	recordStatuses := make([]types.StakeRecordUpdates, len(stakeRecords))
	rawTxFutures := make([]*rpcclient.FutureGetRawTransactionVerboseResult, len(stakeRecords))
	for i, record := range stakeRecords {
		txHash, err := chainhash.NewHashFromStr(record.TxID)
		if err != nil {
			recordStatuses[i].Status = types.Mismatch
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
			log.Error().Msgf("💥 failed receive waits for the Response by txID(%s) on chain: %s", stakeRecords[i].TxID, err)
			recordStatuses[i].Status = stakeRecords[i].FinalizedStatus
			continue
		}
		if stakeRecords[i].FinalizedStatus == types.Mismatch || stakeRecords[i].FinalizedStatus == types.TxFinalized {
			// TODO should not come here. Mismatch/Finalized tx record should't been checked again.
			recordStatuses[i].Status = stakeRecords[i].FinalizedStatus
			continue
		}
		if stakeRecords[i].FinalizedStatus == types.TxPending {
			// durationBlock is store as days in day
			durationBlock := stakeRecords[i].Duration / (24 * time.Hour.Nanoseconds()) * 144

			err = c.CheckStake(txRes, stakeRecords[i].StakerPublicKey, stakeRecords[i].Amount, durationBlock)
			if err != nil {
				recordStatuses[i].Status = types.Mismatch
				continue
			}
			log.Info().Msgf("🔨check stake for txID(%s) on chain: durationBlock[%d] ", stakeRecords[i].TxID, durationBlock)
		}

		if txRes.Confirmations >= TxFinalizedConfirmations { //nolint
			recordStatuses[i].Status = types.TxFinalized
			recordStatuses[i].Start = txRes.Blocktime + int64(10*time.Minute.Seconds())*int64(txRes.Confirmations)
			recordStatuses[i].Start *= 1000000000
		} else if txRes.Confirmations == 0 {
			recordStatuses[i].Status = types.TxPending
		} else {
			recordStatuses[i].Status = types.TxIncluded
			recordStatuses[i].Start = txRes.Blocktime
			recordStatuses[i].Start *= 1000000000
		}

		log.Info().Msgf("🔨check stake on chain: stakeRecords[%+v] ", recordStatuses[i])
	}

	return recordStatuses, nil
}

func (c *Client) CheckStake(tx *btcjson.TxRawResult, stakerPubKeyStr string, amount uint64, duration int64) error {
	if len(tx.Vout) < 1 {
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

const (
	messageSignatureHeader = "Bitcoin Signed Message:\n"
)

func AssembleRewardSignatureMessage(rewardReceiveAddress string, timestamp int64) []byte {
	var buf bytes.Buffer

	msg := fmt.Sprintf("%s/%d", rewardReceiveAddress, timestamp)

	msgH := chainhash.DoubleHashB([]byte(msg))
	_ = wire.WriteVarString(&buf, 0, messageSignatureHeader)
	_ = wire.WriteVarString(&buf, 0, hex.EncodeToString(msgH))

	return chainhash.DoubleHashB(buf.Bytes())
}
