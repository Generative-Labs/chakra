package chakra

import (
	"context"
	"testing"
)

// NewChakraAccount
func TestSubmitTXInfo(t *testing.T) {
	url := "https://madara.to3.io"
	priKeyt := "0x00c1cf1490de1352865301bb8705143f3ef938f97fdf892f1090dcb5ac7bcd1d"
	addr := "0x3"
	contractAddr := "0x2c38b9a62cabbcb1c8e0041ac96bff436c0805e1adfee2137a2bf41f8bf68"

	acc, err := NewChakraAccount(context.Background(), url, priKeyt, addr)
	if err != nil {
		t.Fatalf("New chakra account err:%s", err)
	}

	txID := "ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b047"

	//SubmitTXInfo
	res, err := SubmitTXInfo(context.Background(), acc, contractAddr, txID, "111111", 11111111, 111111111, "0x1")
	//res, err := RPCCall(acc, addr, "getBalance", nil)
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

	txID := []string{"ab44db940bdd7ededefd8b452a219c146b742161fe48512eb370d8b681e1b047"}

	//SubmitTXInfo
	res, err := RewardTo(context.Background(), acc, contractAddr, txID)
	//res, err := RPCCall(acc, addr, "getBalance", nil)
	if err != nil {
		t.Fatalf("RewardTo err:%s", err)
	}

	t.Logf("res %v", res)
}
