package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestPostStakeBtc(t *testing.T) {
	url := "http://localhost:8080/api/stake_btc"
	data := map[string]interface{}{
		"staker":             "bc1xxxxxxxxxx",
		"staker_public_key":  "public_key_here",
		"tx_id":              "transaction_id_here",
		"start":              1234567890, // start timestamp
		"duration":           3600,       // duration in seconds
		"amount":             100,        // stake amount
		"reward":             10,         // reward amount
		"reward_receiver":    "receiver_address_here",
		"btc_signature":      "btc_signature_here",
		"receiver_signature": "receiver_signature_here",
		"timestamp":          1234567890, // timestamp
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error encoding JSON:%s", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Error sending request:%s", err)
		return
	}
	defer resp.Body.Close()

	t.Logf("Response Status:%v", resp.Status)
}
