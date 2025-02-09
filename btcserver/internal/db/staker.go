package db

import (
	"context"
	"time"

	"github.com/generativelabs/btcserver/internal/db/ent"
	"github.com/generativelabs/btcserver/internal/db/ent/stake"
	"github.com/generativelabs/btcserver/internal/types"
	"github.com/generativelabs/btcserver/internal/utils"
)

func (c *Backend) CreateStake(
	staker, stakerPublicKey, txID string,
	duration int64,
	amount uint64,
	rewardReceiver string,
	reward uint64,
	receiverSignature string,
	timestamp int64,
) error {
	ts := utils.MakeTimestamp()
	_, err := c.dbClient.Stake.Create().
		SetStaker(staker).
		SetStakerPublicKey(stakerPublicKey).
		SetTx(txID).
		SetDuration(duration).
		SetAmount(amount).
		SetRewardReceiver(rewardReceiver).
		SetReward(reward).
		SetReceiverSig(receiverSignature).
		SetTimestamp(timestamp).
		SetCreateAt(ts).
		Save(context.Background())

	return err
}

func (c *Backend) UpdateStakeReleasingTime(staker, txID string) error {
	stakeInfo, err := c.dbClient.Stake.Query().
		Where(stake.And(stake.Staker(staker), stake.Tx(txID))).
		Only(context.Background())
	if err != nil {
		return err
	}

	fixedTime := stakeInfo.ReleasingTime + 24*time.Hour.Nanoseconds()
	_, err = c.dbClient.Stake.Update().
		Where(stake.And(stake.Staker(staker), stake.Tx(txID))).
		SetReleasingTime(fixedTime).
		Save(context.Background())
	return err
}

func (c *Backend) UpdateStakeReleaseStatus(staker, txID string, status int) error {
	_, err := c.dbClient.Stake.Update().
		Where(stake.And(stake.Staker(staker), stake.Tx(txID))).
		SetReleaseStatus(status).
		Save(context.Background())
	return err
}

func (c *Backend) UpdateStakeFinalizedStatus(staker, txID string, status int, start int64,
	deadline int64, releasingTime int64,
) error {
	_, err := c.dbClient.Stake.Update().
		Where(stake.And(stake.StakerEQ(staker), stake.TxEQ(txID))).
		SetFinalizedStatus(status).
		SetStart(start).
		SetDeadline(deadline).
		SetReleasingTime(releasingTime).
		Save(context.Background())
	return err
}

func (c *Backend) QueryStakeInfoByStakerAndTxID(staker, txID string) (*ent.Stake, error) {
	return c.dbClient.Stake.Query().
		Where(stake.And(stake.Staker(staker), stake.Tx(txID))).
		Only(context.Background())
}

func (c *Backend) QueryStakesByStaker(staker string, page int, size int) ([]*ent.Stake, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	offset := (page - 1) * size

	return c.dbClient.Stake.Query().
		Where(stake.StakerEQ(staker)).
		Offset(offset).
		Limit(size).
		All(context.Background())
}

func (c *Backend) QueryStakesCountByStaker(staker string) (int, error) {
	return c.dbClient.Stake.Query().
		Where(stake.StakerEQ(staker)).
		Count(context.Background())
}

func (c *Backend) QueryNotReleaseStatesTx(page int, size int) ([]string, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	offset := (page - 1) * size

	return c.dbClient.Stake.Query().
		Where(stake.ReleaseStatus(0)).
		Offset(offset).
		Limit(size).
		Select(stake.FieldTx).
		Strings(context.Background())
}

func (c *Backend) QueryAllNotReleaseStatesTx() ([]string, error) {
	return c.dbClient.Stake.Query().
		Where(stake.ReleaseStatus(0)).
		Select(stake.FieldTx).
		Strings(context.Background())
}

// Query all txids that have not yet been locked up
func (c *Backend) QueryAllNotYetLockedUpTx(timeStamp int64) ([]string, error) {
	return c.dbClient.Stake.Query().
		Where(stake.DeadlineLTE(timeStamp)).
		Select(stake.FieldTx).
		Strings(context.Background())
}

func (c *Backend) QueryAllAlreadyLockedUpTx(timeStamp int64) ([]string, error) {
	return c.dbClient.Stake.Query().
		Where(stake.DeadlineGTE(timeStamp)).
		Select(stake.FieldTx).
		Strings(context.Background())
}

// QueryAllNotYetLockedUpTxNextPeriod Addresses that need to be released in the next 5 minute
func (c *Backend) QueryAllNotYetLockedUpTxNextPeriod(timeStamp int64, timeWheelSize time.Duration) ([]*types.ReleaseTxsInfo, error) {
	feture := timeStamp + timeWheelSize.Nanoseconds()
	stakeInfoList, err := c.dbClient.Stake.Query().
		Where(stake.FinalizedStatus(2)).
		Where(stake.ReleaseStatus(0)).
		Where(stake.DeadlineGT(feture)).
		Where(stake.And(stake.ReleasingTimeGTE(timeStamp), stake.ReleasingTimeLT(feture))).
		Order(ent.Asc(stake.FieldReleasingTime)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	releaseTxsInfos := make([]*types.ReleaseTxsInfo, 0)
	for _, stakeInfo := range stakeInfoList {
		releaseTxsInfos = append(releaseTxsInfos, &types.ReleaseTxsInfo{
			Staker:        stakeInfo.Staker,
			Tx:            stakeInfo.Tx,
			ReleasingTime: stakeInfo.ReleasingTime,
		})
	}

	return releaseTxsInfos, nil
}

func (c *Backend) QueryNoFinalizedStakeTx() ([]*types.StakeVerificationParam, error) {
	verifyParams := make([]*types.StakeVerificationParam, 0)

	err := c.dbClient.Stake.Query().Where(stake.FinalizedStatusLTE(int(types.TxIncluded))).
		Select(stake.FieldStaker, stake.FieldTx, stake.FieldStakerPublicKey, stake.FieldAmount, stake.FieldStart, stake.FieldDuration, stake.FieldRewardReceiver).
		Scan(context.Background(), &verifyParams)
	if err != nil {
		return nil, err
	}

	return verifyParams, nil
}

func (c *Backend) UpdateCanBeSubmitStatus(staker, txID string, status int) error {
	_, err := c.dbClient.Stake.Update().
		Where(stake.And(stake.Staker(staker), stake.Tx(txID))).
		SetSubmitStatus(status).
		Save(context.Background())
	return err
}

func (c *Backend) QueryCanBeSubmitStakeTx() ([]*ent.Stake, error) {
	list, err := c.dbClient.Stake.Query().Where(stake.FinalizedStatusLTE(int(types.TxFinalized))).
		Where(stake.SubmitStatus(0)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return list, nil
}
