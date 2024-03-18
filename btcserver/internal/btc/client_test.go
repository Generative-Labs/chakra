package btc_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
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
		NetworkName: chaincfg.TestNet3Params.Name,
		RPCHost:     "localhost:18332",
		RPCUser:     "rpcuser",
		RPCPass:     "rpcpass",
		DisableTLS:  true,
	})
	assert.NoError(t, err)

	// pubkeyStr := "02435f6406512081c715b3dd1c80166e1443c553470028f1991b6e0270b1a607c0"
	// cairoRewardAddr := "0x65fbbc6ed72f28f38e9b7b440b4115b143a35cfe7ceb390f448fa0a1bcbd8dc"
	// timestamp := int64(1710587088607000)
	// sigBase64 := "IPAl0/CTHAnbqR70r3POQyPzp1Y1hcmru80DF8l72HcOApucgbrMbS5xIx0appoiN8I7VnacRkNlzkxJv389u0c="

	pubkeyStr := "03d87175c1ca3222d1500def9e79692fbd924b85c83e784907b1d1babded7cc72e"
	cairoRewardAddr := "0x1f511713a342ebc4320982295a769dd3d5491ca6416315695b3f62d3e782b71"
	timestamp := int64(1710643417776000)
	sigBase64 := "H3Puyb97JIkLZp5MoiiJIjh65bNJ+25ophK61Bkkb22SDVvk5g3etNL5u0+vLwWXPe7gSSE/f22vOxOAb7LWyK0="

	// pubkeyStr := "03d87175c1ca3222d1500def9e79692fbd924b85c83e784907b1d1babded7cc72e"
	// cairoRewardAddr := "0x1f511713a342ebc4320982295a769dd3d5491ca6416315695b3f62d3e782b71"
	// timestamp := int64(1710641568435000)
	// sigBase64 := "HxZgPvzmYY5EB2d+4VDhEqNzX6zjPiabAKwZps5tunSHKg4l5JWOtKxEQkCcxYQecOm2R/S+iQy0ycPwJQCtUWg="

	err = client.CheckRewardAddressSignature(pubkeyStr, cairoRewardAddr, sigBase64, timestamp)
	assert.NoError(t, err)
}
