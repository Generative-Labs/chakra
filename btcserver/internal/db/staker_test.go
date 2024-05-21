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

func CreateTestDB(t *testing.T) *Backend {
	t.Helper()

	path := "/Users/jiaxingsun/go/chakra/btcserver/cmd/temp/chakrastake.db"
	client := enttest.Open(t, "sqlite3", path+"?_fk=1")
	dbClient := &Backend{
		dbClient: client,
	}
	return dbClient
}

func MockBatchStakeInfo(size int) []*types.StakeInfoReq {
	stakeInfoReqList := make([]*types.StakeInfoReq, 0)
	start := TestTime - 24*time.Hour.Nanoseconds() //+ 5*time.Minute.Nanoseconds()

	for i := 1; i <= size; i++ {
		st := start + int64(i)*time.Minute.Nanoseconds()
		fmt.Printf("== start %s\n", utils.TimestampToTime(st))
		si := &types.StakeInfoReq{
			Staker:                  "bc1xxxxxxxxxx",
			StakerPublicKey:         "0x0000",
			TxID:                    "txidxxxxxxxxxxxxxxxxxxxxx" + strconv.Itoa(i),
			Duration:                7 * 24 * time.Hour.Nanoseconds(),
			Amount:                  uint64(5),
			RewardReceiver:          "0x1111111111",
			RewardReceiverSignature: "receiverSignature",
			Timestamp:               start + 10*time.Minute.Nanoseconds(),
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
	amount := uint64(5)
	rewardReceiver := "0x1111111111"
	reward := uint64(500)
	receiverSignature := "receiverSignature"
	timestamp := start + 10*time.Minute.Nanoseconds()

	err = cli.CreateStake(staker, stakerPublicKey, txID, duration, amount, rewardReceiver, reward, receiverSignature, timestamp)
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
	amount := uint64(5)
	rewardReceiver := "0x1111111111"
	reward := uint64(500)
	receiverSignature := "receiverSignature"
	timestamp := start + 10*time.Minute.Nanoseconds()

	err = cli.CreateStake(staker, stakerPublicKey, txID, duration, amount, rewardReceiver, reward, receiverSignature, timestamp)
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
		err = cli.CreateStake(si.Staker, si.StakerPublicKey, si.TxID, si.Duration, si.Amount, si.RewardReceiver, si.Reward, si.RewardReceiverSignature, si.Timestamp)
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
		err = cli.CreateStake(si.Staker, si.StakerPublicKey, si.TxID, si.Duration, si.Amount, si.RewardReceiver, si.Reward, si.RewardReceiverSignature, si.Timestamp)
		if err != nil {
			assert.NoError(t, err)
		}

		if i == 3 {
			err = cli.UpdateStakeFinalizedStatus(si.Staker, si.TxID, int(types.TxFinalized), 0, 0, 0)
			if err != nil {
				assert.NoError(t, err)
			}
		}

		if i == 5 {
			err = cli.UpdateStakeFinalizedStatus(si.Staker, si.TxID, int(types.TxIncluded), 0, 0, 0)
			if err != nil {
				assert.NoError(t, err)
			}
		}

		if i == 7 {
			err = cli.UpdateStakeFinalizedStatus(si.Staker, si.TxID, int(types.Mismatch), 0, 0, 0)
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

func TestGetStakeListDuringActivity(t *testing.T) {
	cli := CreateTestDB(t)

	//for i := 0; i < 3; i++ {
	//	staker := fmt.Sprintf("bc1xxxxxxxxxx%d", i)
	//	stakerPublicKey := "0x0000"
	//	txID := fmt.Sprintf("txid00000000000000000000%d", i)
	//	start := time.Now().UnixNano() + (4+int64(i))*time.Hour.Nanoseconds()
	//	t.Log("tttt:", start)
	//	duration := 7 * 24 * time.Hour.Nanoseconds()
	//	amount := uint64(5)
	//	rewardReceiver := "0x1111111111"
	//	reward := uint64(500)
	//	receiverSignature := "receiverSignature"
	//	timestamp := start + 10*time.Minute.Nanoseconds()
	//
	//	err := cli.CreateStake(staker, stakerPublicKey, txID, duration, amount, rewardReceiver, reward, receiverSignature, timestamp)
	//	if err != nil {
	//		t.Fatalf("CreateStake err:%s", err)
	//	}
	//
	//	err = cli.UpdateStakeFinalizedStatus(staker, txID, int(types.TxFinalized), start, start+duration, 0)
	//	if err != nil {
	//		t.Fatalf("UpdateStakeFinalizedStatus err:%s", err)
	//	}
	//}
	//
	//for i := 0; i < 3; i++ {
	//	staker := fmt.Sprintf("bc1xxxxxxxxxx%d", i)
	//	stakerPublicKey := "0x0000"
	//	txID := fmt.Sprintf("txid00000000000000000001%d", i)
	//	start := time.Now().UnixNano() + (4+int64(i))*time.Hour.Nanoseconds()
	//	t.Log("tttt:", start)
	//	duration := 7 * 24 * time.Hour.Nanoseconds()
	//	amount := uint64(5)
	//	rewardReceiver := "0x1111111111"
	//	reward := uint64(500)
	//	receiverSignature := "receiverSignature"
	//	timestamp := start + 10*time.Minute.Nanoseconds()
	//
	//	err := cli.CreateStake(staker, stakerPublicKey, txID, duration, amount, rewardReceiver, reward, receiverSignature, timestamp)
	//	if err != nil {
	//		t.Fatalf("CreateStake err:%s", err)
	//	}
	//
	//	err = cli.UpdateStakeFinalizedStatus(staker, txID, int(types.TxFinalized), start, start+duration, 0)
	//	if err != nil {
	//		t.Fatalf("UpdateStakeFinalizedStatus err:%s", err)
	//	}
	//}

	sList, err := cli.GetStakeListDuringActivity(int64(1716311269005570000), int64(1716318469009887000))
	if err != nil {
		t.Fatalf("GetStakeListDuringActivity err:%s", err)
	}

	for _, stakeInfo := range sList {
		t.Logf("stake info :%+v", stakeInfo)

		exist, err := cli.IsStakeExist(stakeInfo.Staker)
		if err != nil {
			t.Fatalf("IsStakeExist err:%s", err)
		}
		if exist {
			t.Errorf("%s is exist", stakeInfo.Staker)
			continue
		}
		si, err := cli.CreateStakeIndex(stakeInfo.Staker, stakeInfo.Tx, stakeInfo.Start)
		if err != nil {
			t.Fatalf("CreateStakeIndex err:%s", err)
			continue
		}
		t.Logf("CreateStakeIndex :%+v", si)
	}

}
