package btc_test

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/generativelabs/btcserver/internal/btc"
	"github.com/stretchr/testify/assert"
)

func TestCheckStakeTx(t *testing.T) {
	normalStakeTxResult := `
	{
		"hex": "02000000019853741549cc1b40e566f7e327a95272f43d66efaa8b60ba26ea54a9b3632459010000006b483045022100fd15ab8ea0a27aa3adf45d3085f88ab9c26ff2330708fd8d647266ae6fde95eb02204504ee6d9808082a719f21921d6540d1baf8b563aaac534c1b4002da850213c50121024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10ffffffff0150c300000000000017a914b2c88fae1cdbdeb20a7d574366d6d441e69e57b88700000000",
		"txid": "b986aa0ecb01c1604f5e800088977f51ac5c1506ae9d2d8baa36b5a6d85197ea",
		"hash": "b986aa0ecb01c1604f5e800088977f51ac5c1506ae9d2d8baa36b5a6d85197ea",
		"size": 190,
		"vsize": 190,
		"weight": 760,
		"version": 2,
		"locktime": 0,
		"vin": [{
			"txid": "592463b3a954ea26ba608baaef663df47252a927e3f766e5401bcc4915745398",
			"vout": 1,
			"scriptSig": {
				"asm": "3045022100fd15ab8ea0a27aa3adf45d3085f88ab9c26ff2330708fd8d647266ae6fde95eb02204504ee6d9808082a719f21921d6540d1baf8b563aaac534c1b4002da850213c5[ALL] 024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10",
				"hex": "483045022100fd15ab8ea0a27aa3adf45d3085f88ab9c26ff2330708fd8d647266ae6fde95eb02204504ee6d9808082a719f21921d6540d1baf8b563aaac534c1b4002da850213c50121024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10"
			},
			"sequence": 4294967295
		}],
		"vout": [{
			"value": 0.0005,
			"n": 0,
			"scriptPubKey": {
				"asm": "OP_HASH160 b2c88fae1cdbdeb20a7d574366d6d441e69e57b8 OP_EQUAL",
				"hex": "a914b2c88fae1cdbdeb20a7d574366d6d441e69e57b887",
				"type": "scripthash",
				"address": "2N9YYXfoJUt4iffCh4teBuGDbjk9dg84Jdi"
			}
		}]
	}
	`

	client, err := btc.NewClient(btc.Config{
		NetworkName: chaincfg.RegressionNetParams.Name,
		RPCHost:     "localhost:18332",
		RPCUser:     "rpcuser",
		RPCPass:     "rpcpass",
		DisableTLS:  true,
	})
	assert.NoError(t, err)

	var rawTxRes btcjson.TxRawResult
	err = json.Unmarshal([]byte(normalStakeTxResult), &rawTxRes)
	assert.NoError(t, err)

	tests := []struct {
		description   string
		pubKeyStr     string
		stakeAmount   uint64
		lockDuration  int64
		expectedError error
	}{
		{
			description:   "match",
			pubKeyStr:     "024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10", // BOB public key
			stakeAmount:   50000,
			lockDuration:  20,
			expectedError: nil,
		},
		{
			description:   "mismatch public key",
			pubKeyStr:     "025b81f0017e2091e2edcd5eecf10d5bdd120a5514cb3ee65b8447ec18bfc4575c",
			stakeAmount:   50000,
			lockDuration:  20,
			expectedError: errors.New("verify stake tx failed"),
		},
		{
			description:   "mismatch stake amount",
			pubKeyStr:     "024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10", // BOB public key
			stakeAmount:   10000,
			lockDuration:  20,
			expectedError: errors.New("verify stake tx failed"),
		},
		{
			description:   "mismatch stake duration",
			pubKeyStr:     "024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10", // BOB public key
			stakeAmount:   10000,
			lockDuration:  20,
			expectedError: errors.New("verify stake tx failed"),
		},
	}

	for _, test := range tests {
		err := client.CheckStake(&rawTxRes, test.pubKeyStr, test.stakeAmount, test.lockDuration)
		if test.expectedError != nil {
			assert.Equal(t, err, test.expectedError, test.description)
		} else {
			assert.NoError(t, err, test.description)
		}
	}
}

func TestCheckRewardAddressSignature(t *testing.T) {
	client, err := btc.NewClient(btc.Config{
		NetworkName: chaincfg.RegressionNetParams.Name,
		RPCHost:     "localhost:18332",
		RPCUser:     "rpcuser",
		RPCPass:     "rpcpass",
		DisableTLS:  true,
	})
	assert.NoError(t, err)

	bobPrivKey, err := btcutil.DecodeWIF(bobPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	cairoRewardAddr := "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7"
	timestamp := int32(1710387487)

	message := btc.AssembleRewardSignatureMessage(cairoRewardAddr, timestamp)
	msgH := chainhash.DoubleHashB([]byte(message))

	signature := ecdsa.Sign(bobPrivKey.PrivKey, msgH)
	sigB := signature.Serialize()
	sigH := hex.EncodeToString(sigB)

	pubkeyStr := "024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10"

	err = client.CheckRewardAddressSignature(pubkeyStr, cairoRewardAddr, sigH, timestamp)
	assert.NoError(t, err)

	mismatchCairoRewardAddr := "0x3d19214f89175a68b1874341ac8afa0e4f30dc114820d0e4039ee8c2be0a30f"
	err = client.CheckRewardAddressSignature(pubkeyStr, mismatchCairoRewardAddr, sigH, timestamp)
	assert.EqualError(t, err, "reward address signature verify failed")
}
