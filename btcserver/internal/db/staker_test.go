package db

import (
	"strconv"
	"testing"
	"time"

	"github.com/generativelabs/btcserver/internal/types"
)

var (
	TestTime = time.Now().UnixNano()
)

func InitDB() (*Backend, error) {
	client, err := CreateSqliteDB("/Users/jiaxingsun/go/chakra/btcserver/cmd/temp/btc_server_db")
	if err != nil {
		return nil, err
	}

	dbClient := &Backend{
		dbClient: client,
	}

	return dbClient, nil
}

func InitBatchStakeInfo() []*types.StakeInfoReq {
	stakeInfoReqList := make([]*types.StakeInfoReq, 0)
	start := TestTime

	for i := 1; i <= 10; i++ {
		si := &types.StakeInfoReq{
			Staker:            "bc1xxxxxxxxxx",
			StakerPublicKey:   "0x0000",
			TxID:              "txidxxxxxxxxxxxxxxxxxxxxx" + strconv.Itoa(i),
			Start:             start + int64(i)*time.Minute.Nanoseconds(),
			Duration:          7 * 24 * time.Hour.Nanoseconds(),
			Amount:            int64(5),
			RewardReceiver:    "0x1111111111",
			ReceiverSignature: "receiverSignature",
			Timestamp:         start + 10*time.Minute.Nanoseconds(),
		}

		stakeInfoReqList = append(stakeInfoReqList, si)
	}

	//for i := 0; i < 10; i++ {
	//	si := &types.StakeInfoReq{
	//		Staker:            "bc1yyyyyyyyyy",
	//		StakerPublicKey:   "0x0000",
	//		TxID:              "txidyyyyyyyyyyyyyyyyy" + strconv.Itoa(i),
	//		Start:             start + int64(i),
	//		Duration:          7 * 24 * time.Hour.Nanoseconds(),
	//		Amount:            int64(5),
	//		RewardReceiver:    "0x1111111111",
	//		ReceiverSignature: "receiverSignature",
	//		Timestamp:         start + 10*time.Minute.Nanoseconds(),
	//	}
	//
	//	stakeInfoReqList = append(stakeInfoReqList, si)
	//}

	return stakeInfoReqList
}

func TestCreateStake(t *testing.T) {
	cli, err := InitDB()
	if err != nil {
		t.Fatalf("Init db err:%s", err)
	}

	staker := "bc1xxxxxxxxxx"
	stakerPublicKey := "0x0000"
	txID := "txid00000000000000000000"
	start := time.Now().UnixNano() + 4*time.Minute.Nanoseconds()
	duration := 7 * 24 * time.Hour.Nanoseconds()
	amount := int64(5)
	rewardReceiver := "0x1111111111"
	receiverSignature := "receiverSignature"
	timestamp := start + 10*time.Minute.Nanoseconds()

	err = cli.CreateStake(staker, stakerPublicKey, txID, start, duration, amount, rewardReceiver, receiverSignature, timestamp)
	if err != nil {
		t.Fatalf("CreateStake err:%s", err)
	}
}

func TestUpdateStakeReleasingTime(t *testing.T) {
	cli, err := InitDB()
	if err != nil {
		t.Fatalf("Init db err:%s", err)
	}

	staker := "bc1xxxxxxxxxx"
	txID := "txid00000000000000000000"

	si, err := cli.QueryStakeInfoByStakerAndTxID(staker, txID)
	if err != nil {
		t.Fatalf("CreateStake err:%s", err)
	}

	t.Logf("StakeInfo :%d", si.ReleasingTime)

	err = cli.UpdateStakeReleasingTime(staker, txID)
	if err != nil {
		t.Fatalf("CreateStake err:%s", err)
	}

	si, err = cli.QueryStakeInfoByStakerAndTxID(staker, txID)
	if err != nil {
		t.Fatalf("CreateStake err:%s", err)
	}

	t.Logf("After update:StakeInfo :%d", si.ReleasingTime)
}

// QueryAllNotYetLockedUpTxNextPeriod
func TestQueryAllNotYetLockedUpTxNextPeriod(t *testing.T) {
	cli, err := InitDB()
	if err != nil {
		t.Fatalf("Init db err:%s", err)
	}

	siList := InitBatchStakeInfo()
	for _, si := range siList {
		err = cli.CreateStake(si.Staker, si.StakerPublicKey, si.TxID, si.Start, si.Duration, si.Amount, si.RewardReceiver, si.ReceiverSignature, si.Timestamp)
		if err != nil {
			t.Fatalf("CreateStake err:%s", err)
		}
	}

	tt := TestTime + 24*time.Hour.Nanoseconds()

	txs, err := cli.QueryAllNotYetLockedUpTxNextPeriod(tt, types.TimeWheelSize)
	if err != nil {
		t.Fatalf("QueryAllNotYetLockedUpTxNextPeriod err:%s", err)
	}

	t.Logf("Release reward for :tt %d len %d %+v", tt, len(txs), txs)
}
