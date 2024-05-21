package db

import (
	"context"
	"github.com/generativelabs/btcserver/internal/db/ent"
	"github.com/generativelabs/btcserver/internal/db/ent/stakeindex"
	"time"
)

func (c *Backend) CreateStakeIndex(staker, txID string, start int64) (*ent.StakeIndex, error) {
	si, err := c.dbClient.StakeIndex.Create().
		SetStaker(staker).
		SetTx(txID).
		SetStart(start).
		SetCreateAt(time.Now().Unix()).
		Save(context.Background())
	return si, err
}

func (c *Backend) UpdateStakeIndex(staker, index string) error {
	_, err := c.dbClient.StakeIndex.Update().
		Where(stakeindex.Staker(staker)).
		SetStaker(staker).
		Save(context.Background())
	return err
}

func (c *Backend) IsStakeExist(staker string) (bool, error) {
	return c.dbClient.StakeIndex.Query().
		Where(stakeindex.Staker(staker)).
		Exist(context.Background())
}

func (c *Backend) GetStakeIndex(staker string) (int, error) {
	index, err := c.dbClient.StakeIndex.Query().
		Where(stakeindex.Staker(staker)).
		OnlyID(context.Background())

	return index, err
}
