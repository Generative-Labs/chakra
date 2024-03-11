package db

import (
	"context"
	"time"

	"github.com/generativelabs/btcserver/internal/db/ent"
	"github.com/generativelabs/btcserver/internal/db/ent/stake"
)

func (c *Backend) CreateStake(
	staker, stakerPublicKey, txID string,
	start uint64,
	duration uint64,
	amount uint64,
	rewardReceiver string,
	btcSignature string,
	receiverSignature string,
	timestamp uint64,
) error {
	_, err := c.dbClient.Stake.Create().
		SetStaker(staker).
		SetStakerPublicKey(stakerPublicKey).
		SetTx(txID).
		SetStart(start).
		SetDuration(duration).
		SetDeadline(start + duration).
		SetAmount(amount).
		SetRewardReceiver(rewardReceiver).
		SetBtcSig(btcSignature).
		SetReceiverSig(receiverSignature).
		SetTimestamp(timestamp).
		Save(context.Background())

	return err
}

func (c *Backend) UpdateStakeReleaseStatus(staker string, status int) error {
	_, err := c.dbClient.Stake.Update().
		Where(stake.StakerEQ(staker)).
		SetReleaseStatus(status).
		Save(context.Background())
	return err
}

func (c *Backend) UpdateStakeFinalizedStatus(staker string, status int) error {
	_, err := c.dbClient.Stake.Update().
		Where(stake.StakerEQ(staker)).
		SetFinalizedStatus(status).
		Save(context.Background())
	return err
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
func (c *Backend) QueryAllNotYetLockedUpTx() ([]string, error) {
	return c.dbClient.Stake.Query().
		Where(stake.DeadlineLTE(uint64(time.Now().UnixMilli()))).
		Select(stake.FieldTx).
		Strings(context.Background())
}

func (c *Backend) QueryAllAlreadyLockedUpTx() ([]string, error) {
	return c.dbClient.Stake.Query().
		Where(stake.DeadlineGTE(uint64(time.Now().UnixMilli()))).
		Select(stake.FieldTx).
		Strings(context.Background())
}
