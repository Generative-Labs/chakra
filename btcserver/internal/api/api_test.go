package api_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/generativelabs/btcserver/internal/api"
	"github.com/generativelabs/btcserver/internal/btc"
	"github.com/generativelabs/btcserver/internal/db"
	"github.com/generativelabs/btcserver/internal/db/ent/enttest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createMemoryTestDB(t *testing.T) *db.Backend {
	t.Helper()

	dbClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	return db.CreateBackendWithDB(dbClient)
}

func TestPostStakeBtc(t *testing.T) {
	memoryDB := createMemoryTestDB(t)
	btcClient, err := btc.NewClient(btc.Config{
		NetworkName: chaincfg.RegressionNetParams.Name,
		RPCHost:     "localhost:18332",
		RPCUser:     "rpcuser",
		RPCPass:     "rpcpass",
		DisableTLS:  true,
	})
	assert.NoError(t, err)

	server := api.NewServer(context.Background(), memoryDB, nil, "null", btcClient)

	router := gin.Default()
	router.Use(api.CORSMiddleware())
	api.SetupRoutes(router, server)

	data := map[string]interface{}{
		"staker":             "bc123",
		"staker_public_key":  "0x123",
		"tx_id":              "0x123",
		"start":              1234567890, // start timestamp
		"duration":           3600,       // duration in seconds
		"amount":             100,        // stake amount
		"reward":             10,         // reward amount
		"reward_receiver":    "0x123",
		"btc_signature":      "btc_signature_here",
		"receiver_signature": "0x123",
		"timestamp":          1234567890, // timestamp
	}

	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error encoding JSON:%s", err)
		return
	}

	req, _ := http.NewRequest("POST", "/api/stake_btc", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	// public key is invalid
	assert.True(t, strings.Contains(w.Body.String(), "public key"))
}

func TestPostStakeBtcWithValidRewardSignature(t *testing.T) {
	memoryDB := createMemoryTestDB(t)
	btcClient, err := btc.NewClient(btc.Config{
		NetworkName: chaincfg.RegressionNetParams.Name,
		RPCHost:     "localhost:18332",
		RPCUser:     "rpcuser",
		RPCPass:     "rpcpass",
		DisableTLS:  true,
	})
	assert.NoError(t, err)

	server := api.NewServer(context.Background(), memoryDB, nil, "null", btcClient)

	router := gin.Default()
	router.Use(api.CORSMiddleware())
	api.SetupRoutes(router, server)

	cairoRewardAddr := "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7"
	timestamp := int32(1710387487)

	message := btc.AssembleRewardSignatureMessage(cairoRewardAddr, timestamp)
	msgH := chainhash.DoubleHashB([]byte(message))

	bobPrivKey, err := btcutil.DecodeWIF("5JoQtsKQuH8hC9MyvfJAqo6qmKLm8ePYNucs7tPu2YxG12trzBt")
	if err != nil {
		fmt.Println(err)
		return
	}
	signature := ecdsa.Sign(bobPrivKey.PrivKey, msgH)
	sigB := signature.Serialize()
	rewardSigHex := hex.EncodeToString(sigB)

	data := map[string]interface{}{
		"staker":             "bc1xxxxxxxxxx",
		"staker_public_key":  "0x024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10",
		"tx_id":              "transaction_id_here",
		"start":              1234567890, // start timestamp
		"duration":           3600,       // duration in seconds
		"amount":             100,        // stake amount
		"reward":             10,         // reward amount
		"reward_receiver":    "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7",
		"btc_signature":      "0x123",
		"receiver_signature": rewardSigHex,
		"timestamp":          1710387487, // timestamp
	}

	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error encoding JSON:%s", err)
		return
	}

	req, _ := http.NewRequest("POST", "/api/stake_btc", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// public key is invalid
	assert.True(t, strings.Contains(w.Body.String(), "Ok"))

	// check record in database
	stakes, err := memoryDB.QueryStakesByStaker("bc1xxxxxxxxxx", 1, 10)
	assert.NoError(t, err)
	assert.Equal(t, len(stakes), 1)
	assert.Equal(t, stakes[0].Staker, "bc1xxxxxxxxxx")
	assert.Equal(t, stakes[0].Tx, "transaction_id_here")
	assert.Equal(t, stakes[0].StakerPublicKey, "0x024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10")
}
