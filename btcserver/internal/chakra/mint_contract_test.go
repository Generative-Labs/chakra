package chakra

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/stretchr/testify/assert"
)

// NewChakraAccount
func TestSubmitTXInfo(t *testing.T) {
	url := "https://madara.to3.io"                                                    //nolint
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"   //nolint
	addr := "0x3"                                                                     //nolint
	contractAddr := "0x2c38b9a62cabbcb1c8e0041ac96bff436c0805e1adfee2137a2bf41f8bf68" //nolint

	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	// ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b048
	txID := "0x1"

	// SubmitTXInfo
	res, err := SubmitTXInfo(context.Background(), acc, contractAddr, txID, "111111", 11111111, 111111111, "0x1")
	if err != nil {
		t.Fatalf("Submit TX Info err:%s", err)
	}

	t.Logf("res %v", res)
}

func TestRewardTo(t *testing.T) {
	url := "https://madara.to3.io"
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"
	addr := "0x3"
	contractAddr := "0x2c38b9a62cabbcb1c8e0041ac96bff436c0805e1adfee2137a2bf41f8bf68"

	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	// ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b047
	txID := []string{"0x1"}

	// RewardTo
	res, err := RewardTo(context.Background(), acc, contractAddr, txID)
	if err != nil {
		t.Fatalf("RewardTo err:%s", err)
	}

	t.Logf("res %v", res)
}

func TestTxRewardsof(t *testing.T) {
	url := "https://madara.to3.io"
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"
	addr := "0x3"
	contractAddr := "0x2c38b9a62cabbcb1c8e0041ac96bff436c0805e1adfee2137a2bf41f8bf68"

	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	txIDs := []string{"0x1"}
	params := ArrBtcTxIDToFelt(txIDs)
	if err != nil {
		t.Fatalf("ArrBtcTxIDToFelt err:%s", err)
	}

	lenP := utils.BigIntToFelt(big.NewInt(int64(len(txIDs))))
	callData := make([]*felt.Felt, 0)
	callData = append(callData, lenP)
	callData = append(callData, params...)

	// RPCCall
	res, err := RPCCall(acc, contractAddr, "RPCCall", callData)
	if err != nil {
		t.Fatalf("RPCCall err:%s", err)
	}
	t.Logf("receipt %v", res)
}

func TestTxReceipt(t *testing.T) {
	url := "https://madara.to3.io"
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"
	addr := "0x3"

	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	// RewardTo
	// 0x5d4f285810fcb00facdc67e3b10ef35c631d2b886191871c4e2f144487bbdf4
	// 0x79ac97faa71ac1c9b9a86d3f814735d0a3ff4cc9469395a580af72e7a9aa969
	// 0x180e1b78954a76c8e450a6f3b2e58c4537a1c333efdcfa0575966bdf07a1145
	// 0x48bc822784e5b25347d18fdccf83db0bbbb5ae29c97d718663cf276f8bab8d9

	// submit
	// 0x13496bf70f595ed5e80c83241e73c4eeaaa49b78e1df0e0f2e83f9db6952f11
	// 0x2b142e0547bbd6272a865abf7628f7870633e7b3274c0c870718111b98118a
	hash, err := utils.HexToFelt("0x48bc822784e5b25347d18fdccf83db0bbbb5ae29c97d718663cf276f8bab8d9")
	if err != nil {
		t.Fatalf("HexToFelt err:%s", err)
	}
	receipt, err := acc.TransactionReceipt(context.Background(), hash)
	if err != nil {
		t.Fatalf("TransactionReceipt err:%s", err)
	}
	t.Logf("receipt %v", receipt)
}

func TestArrayTxID(t *testing.T) {
	err := ArrayTxIDToFelt(t)
	assert.NoError(t, err)

	err = TxIDToFelt(t)
	assert.NoError(t, err)
}

func ArrayTxIDToFelt(t *testing.T) error {
	t.Helper()

	txID := []string{"ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b047"}

	txByte, err := json.Marshal(txID)
	if err != nil {
		t.Fatalf("Marshal err:%s", err)
	}
	bigf := utils.BytesToBig(txByte)
	f := utils.BigIntToFelt(bigf)
	t.Logf("BigIntToFelt %v", f)
	return nil
}

func TxIDToFelt(t *testing.T) error {
	t.Helper()

	txID := "ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b047"

	txByte, err := json.Marshal(txID)
	if err != nil {
		return err
	}
	bigf := utils.BytesToBig(txByte)
	f := utils.BigIntToFelt(bigf)
	t.Logf("BigIntToFelt %v", f)
	return nil
}
