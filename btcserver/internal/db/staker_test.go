package db

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/generativelabs/btcserver/internal/db/ent/enttest"
	"github.com/generativelabs/btcserver/internal/types"
	"github.com/generativelabs/btcserver/internal/utils"
	"github.com/stretchr/testify/assert"
)

var TestTime = time.Now().UnixNano()

func CreateMemoryTestDB(t *testing.T) (*Backend, error) {
	t.Helper()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	dbClient := &Backend{
		dbClient: client,
	}

	return dbClient, nil
}

func MockBatchStakeInfo(size int) []*types.StakeInfoReq {
	stakeInfoReqList := make([]*types.StakeInfoReq, 0)
	start := TestTime - 24*time.Hour.Nanoseconds() //+ 5*time.Minute.Nanoseconds()

	for i := 1; i <= size; i++ {
		st := start + int64(i)*time.Minute.Nanoseconds()
		fmt.Printf("== start %s\n", utils.TimestampToTime(st))
		si := &types.StakeInfoReq{
			Staker:            "bc1xxxxxxxxxx",
			StakerPublicKey:   "0x0000",
			TxID:              "txidxxxxxxxxxxxxxxxxxxxxx" + strconv.Itoa(i),
			Start:             st,
			Duration:          7 * 24 * time.Hour.Nanoseconds(),
			Amount:            int64(5),
			RewardReceiver:    "0x1111111111",
			ReceiverSignature: "receiverSignature",
			Timestamp:         start + 10*time.Minute.Nanoseconds(),
		}

		stakeInfoReqList = append(stakeInfoReqList, si)
	}

	return stakeInfoReqList
}

func TestCreateStake(t *testing.T) {
	cli, err := CreateMemoryTestDB(t)
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
	cli, err := CreateMemoryTestDB(t)
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
	cli, err := CreateMemoryTestDB(t)
	if err != nil {
		t.Fatalf("Init db err:%s", err)
	}

	siList := MockBatchStakeInfo(10)
	for _, si := range siList {
		err = cli.CreateStake(si.Staker, si.StakerPublicKey, si.TxID, si.Start, si.Duration, si.Amount, si.RewardReceiver, si.ReceiverSignature, si.Timestamp)
		if err != nil {
			t.Fatalf("CreateStake err:%s", err)
		}
	}

	tt := TestTime //+ 24*time.Hour.Nanoseconds()
	fmt.Printf("== query %s\n", utils.TimestampToTime(tt))
	txs, err := cli.QueryAllNotYetLockedUpTxNextPeriod(tt, types.TimeWheelSize)
	if err != nil {
		t.Fatalf("QueryAllNotYetLockedUpTxNextPeriod err:%s", err)
	}

	t.Logf("Release reward for :tt %d len %d %+v", tt, len(txs), txs)
}

func TestQueryNoFinalizedStakeTx(t *testing.T) {
	cli, err := CreateMemoryTestDB(t)
	if err != nil {
		t.Fatalf("Init db err:%s", err)
	}

	stakeRecordSize := 10
	siList := MockBatchStakeInfo(stakeRecordSize)
	for i, si := range siList {
		err = cli.CreateStake(si.Staker, si.StakerPublicKey, si.TxID, si.Start, si.Duration, si.Amount, si.RewardReceiver, si.ReceiverSignature, si.Timestamp)
		if err != nil {
			assert.NoError(t, err)
		}

		if i == 3 {
			err = cli.UpdateStakeFinalizedStatus(si.Staker, si.TxID, int(types.TxFinalized))
			if err != nil {
				assert.NoError(t, err)
			}
		}

		if i == 5 {
			err = cli.UpdateStakeFinalizedStatus(si.Staker, si.TxID, int(types.TxIncluded))
			if err != nil {
				assert.NoError(t, err)
			}
		}

		if i == 7 {
			err = cli.UpdateStakeFinalizedStatus(si.Staker, si.TxID, int(types.Mismatch))
			if err != nil {
				assert.NoError(t, err)
			}
		}
	}

	noFinalizeStakes, err := cli.QueryNoFinalizedStakeTx()
	if err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, len(noFinalizeStakes), stakeRecordSize-2)
}
