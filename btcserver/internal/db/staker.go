package db

import (
	"context"

	"github.com/generativelabs/btcserver/internal/db/ent"
	"github.com/generativelabs/btcserver/internal/db/ent/stake"
)

//Staker          string `form:"staker" json:"staker,omitempty"`
//	TxID            string `json:"tx_id,omitempty"`
//	Start           int64  `json:"start,omitempty"`
//	Duration        int64  `json:"duration,omitempty"`
//	Amount          int64  `json:"amount,omitempty"`
//	ReceiverAddress string `json:"receiver_address,omitempty"`
//	BtcSignature    string `json:"btc_signature,omitempty"`
//	RewordSignature string `json:"reword_signature,omitempty"`

func (c *Backend) CreateStake(
	staker string,
	txID string,
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
		SetTx(txID).
		SetStart(start).
		SetDuration(duration).
		SetAmount(amount).
		SetRewardReceiver(rewardReceiver).
		SetBtcSig(btcSignature).
		SetReceiverSig(receiverSignature).
		SetTimestamp(timestamp).
		Save(context.Background())

	return err
}

func (c *Backend) UpdateStakeReleaseStatus(staker string, Status bool) error {
	_, err := c.dbClient.Stake.Update().
		Where(stake.StakerEQ(staker)).
		SetReleaseStatus(Status).
		Save(context.Background())
	return err
}

func (c *Backend) UpdateStakeFinalizedStatus(staker string, Status bool) error {
	_, err := c.dbClient.Stake.Update().
		Where(stake.StakerEQ(staker)).
		SetFinalizedStatus(Status).
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

func (c *Backend) QueryNotEndStatesTx(limit int) ([]string, error) {
	return c.dbClient.Stake.Query().Where(stake.ReleaseStatus(false)).
		Limit(limit).Select(stake.FieldTx).Strings(context.Background())
}
