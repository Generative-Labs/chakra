package api_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
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
	// memoryDB := createMemoryTestDB(t)
	// btcClient, err := btc.NewClient(btc.Config{
	//	NetworkName: chaincfg.RegressionNetParams.Name,
	//	RPCHost:     "localhost:18332",
	//	RPCUser:     "rpcuser",
	//	RPCPass:     "rpcpass",
	//	DisableTLS:  true,
	// })
	// assert.NoError(t, err)

	// server := api.NewServer(context.Background(), memoryDB, nil, "null", btcClient)
	//
	// router := gin.Default()
	// router.Use(api.CORSMiddleware())
	// api.SetupRoutes(router, server)

	// {"staker":"","staker_public_key":"","tx_id":"","start":,"durnation":7,"amount":90,"reward":0.063,"receiver_receiver":"","receiver_address_signature":"","timestamp":1710481332057}

	data := map[string]interface{}{
		"staker":             "tb1qdp844q47twdjffxc3ptxu7atm6uw5ss3h98td2",
		"staker_public_key":  "02435f6406512081c715b3dd1c80166e1443c553470028f1991b6e0270b1a607c0",
		"tx_id":              "39d93ae35e841ec14e83205b1a4f5660894983a96f94c5bedb3273e58afde756",
		"start":              1710481332057, // start timestamp
		"duration":           7,             // duration in seconds
		"amount":             90,            // stake amount
		"reward":             0.063,         // reward amount
		"receiver_receiver":  "0x65fbbc6ed72f28f38e9b7b440b4115b143a35cfe7ceb390f448fa0a1bcbd8dc",
		"receiver_signature": "H3Ma5HAYXV6HlBmWpcfAbuT08wbVmzyRUdkEFJBjzuBUejZuk+VKgifIdF50cAIGSWkWlYhqt2ck4RWSb1YZBwM=",
		"timestamp":          1710481332057, // timestamp
	}

	body, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error encoding JSON:%s", err)
		return
	}

	req, err := http.NewRequest("POST", "http://3.0.18.212:8080/api/stake_btc", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Error NewRequest :%s", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	t.Logf("req:%+v", resp.Body)
	t.Logf("StatusCode:%+v", resp.StatusCode)

	// req.Header.Set("Content-Type", "application/json")
	// w := httptest.NewRecorder()
	// router.ServeHTTP(w, req)

	// assert.Equal(t, 400, w.Code)
	// // public key is invalid
	// assert.True(t, strings.Contains(w.Body.String(), "public key"))
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
	timestamp := int64(1710387487)

	msgH := btc.AssembleRewardSignatureMessage(cairoRewardAddr, timestamp)

	bobPrivKey, err := btcutil.DecodeWIF("5JoQtsKQuH8hC9MyvfJAqo6qmKLm8ePYNucs7tPu2YxG12trzBt")
	assert.NoError(t, err)

	signature, err := ecdsa.SignCompact(bobPrivKey.PrivKey, msgH, true)
	assert.NoError(t, err)
	rewardSigBase64 := base64.StdEncoding.EncodeToString(signature)

	data := map[string]interface{}{
		"staker":                     "bc1xxxxxxxxxx",
		"staker_public_key":          "0x024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10",
		"tx_id":                      "39d93ae35e841ec14e83205b1a4f5660894983a96f94c5bedb3273e58afde756",
		"duration":                   7,   // duration in days
		"amount":                     100, // stake amount
		"reward":                     10,  // reward amount
		"receiver_receiver":          "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7",
		"receiver_address_signature": rewardSigBase64,
		"timestamp":                  1710387487, // timestamp
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
	assert.Equal(t, stakes[0].Tx, "39d93ae35e841ec14e83205b1a4f5660894983a96f94c5bedb3273e58afde756")
	assert.Equal(t, stakes[0].StakerPublicKey, "0x024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10")
	assert.Equal(t, stakes[0].Duration, int64(7*(time.Hour*24)))
}

func TestGetStakesList(_ *testing.T) {
	// 构建 API 地址
	apiURL := "http://localhost:8080/api/stakes_list?staker=example&page=1&size=10"

	// 发送 GET 请求
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// 处理响应数据
	fmt.Println("Total Count:", data["total_count"])
	fmt.Println("Data List:")
	for _, item := range data["data_list"].([]interface{}) {
		fmt.Println(item)
	}
}
