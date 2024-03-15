package chakra

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/stretchr/testify/assert"
)

const BTCTxID = "e1280cbfb6fea0c4b234da582cff735b03b672e2719725f369d76b4f78f260d6"

// NewChakraAccount
func TestSubmitTXInfo(t *testing.T) {
	url := "https://madara.to3.io"                                                    //nolint
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"   //nolint
	addr := "0x3"                                                                     //nolint
	contractAddr := "0x4ca7cade3c35486817d8a0880cc8919cb82d92dd79208ffdecc31328f6174" //nolint
	startAt := time.Now().UnixNano()
	expireAt := time.Now().Add(time.Hour * 72).UnixNano()
	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	// ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b048
	txID := BTCTxID

	// SubmitTXInfo
	res, err := SubmitTXInfo(context.Background(), acc, contractAddr, txID, "200000000000000000", startAt, expireAt, "0x1")
	if err != nil {
		t.Fatalf("Submit TX Info err:%s", err)
	}
	t.Logf("res %v", res)

	receipt, err := acc.WaitForTransactionReceipt(context.Background(), res.TransactionHash, time.Second)
	if err != nil {
		t.Fatalf("Submit TX Info transactionReceipt err:%s", err)
	}
	t.Logf("receipt %v", *receipt)
}

func TestRewardTo(t *testing.T) {
	url := "https://madara.to3.io"
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"
	addr := "0x3"
	contractAddr := "0x4ca7cade3c35486817d8a0880cc8919cb82d92dd79208ffdecc31328f6174"

	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	// ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b047
	txID := []string{BTCTxID}

	// RewardTo
	res, err := RewardTo(context.Background(), acc, contractAddr, txID)
	if err != nil {
		t.Fatalf("RewardTo err:%s", err)
	}

	t.Logf("res %v", res)

	receipt, err := acc.WaitForTransactionReceipt(context.Background(), res.TransactionHash, time.Second)
	if err != nil {
		t.Fatalf("RewardTo transactionReceipt err:%s", err)
	}
	t.Logf("receipt %v", *receipt)
}

func TestTxRewardsof(t *testing.T) {
	url := "https://madara.to3.io"
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"
	addr := "0x3"
	contractAddr := "0x4ca7cade3c35486817d8a0880cc8919cb82d92dd79208ffdecc31328f6174"

	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	txIDs := BTCTxID
	params := utils.BigIntToFelt(utils.HexToBN(txIDs))
	//params, err := utils.HexToFelt(txIDs)
	//if err != nil {
	//	t.Fatalf("HexToFelt err:%s", err)
	//}

	//lenP := utils.BigIntToFelt(big.NewInt(int64(len(txIDs))))
	callData := make([]*felt.Felt, 0)
	//callData = append(callData, lenP)
	callData = append(callData, params)

	// RPCCall
	res, err := RPCCall(acc, contractAddr, "rewardsOf", callData)
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
	// 0x41f88729384193e8a51892d8e0bb815420f4a0b047cc852134e1a63fa261f0f
	// 0x51967d02a9d2dd355f1ab7a9f56f50728ce55927b6d22f32603d5dd05880d6e
	// 0x67c00d275c5fb25553b40fc5a0eaee006f86b74f1d0f0140764995619d8b71e

	// submit
	// 0x754821689f87e4ceb3fb54701fe08889edc5f842160ea2c8e6d5ec24a709ff0
	// 0x68e8676e2098828cc54401c30e57c34e564685f9f0161dff4cf8c5159dccc97
	// 0x642bf80a5d0e5c72ce6333e591e9ba6da4bfb1242fbad5b28d352b87799271c

	hash, err := utils.HexToFelt("0x51967d02a9d2dd355f1ab7a9f56f50728ce55927b6d22f32603d5dd05880d6e")
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
