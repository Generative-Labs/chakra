package db

import (
	"context"
	"strconv"

	"github.com/generativelabs/btcserver/internal/db/ent"
	"github.com/generativelabs/btcserver/internal/db/ent/globalstate"
	"github.com/generativelabs/btcserver/internal/utils"
)

func (c *Backend) IsTimeWheelExist() (bool, error) {
	return c.dbClient.GlobalState.Query().
		Where(globalstate.Key("TimeWheelProgress")).Exist(context.Background())
}

func (c *Backend) GetTimeWheel() (*ent.GlobalState, error) {
	return c.dbClient.GlobalState.Query().
		Where(globalstate.Key("TimeWheelProgress")).Only(context.Background())
}

func (c *Backend) CreateTimeWheel(timeWheel int64) error {
	_, err := c.dbClient.GlobalState.Create().
		SetKey("TimeWheelProgress").
		SetValue(strconv.Itoa(int(timeWheel))).
		SetCreateAt(utils.MakeTimestamp()).
		Save(context.Background())

	return err
}

func (c *Backend) UpdateTimeWheel(timeWheel int64) error {
	_, err := c.dbClient.GlobalState.Update().
		Where(globalstate.Key("TimeWheelProgress")).
		SetValue(strconv.Itoa(int(timeWheel))).
		Save(context.Background())
	return err
}

func (c *Backend) UpsertTimeWheel(timeWheel int64) error {
	exist, err := c.dbClient.GlobalState.Query().
		Where(globalstate.Key("TimeWheelProgress")).Exist(context.Background())
	if err != nil {
		return err
	}

	if exist {
		_, err = c.dbClient.GlobalState.Update().
			Where(globalstate.Key("TimeWheelProgress")).
			SetValue(strconv.Itoa(int(timeWheel))).
			Save(context.Background())
	} else {
		_, err = c.dbClient.GlobalState.Create().
			SetKey("TimeWheelProgress").
			SetValue(strconv.Itoa(int(timeWheel))).
			SetCreateAt(utils.MakeTimestamp()).
			Save(context.Background())
	}

	return err
}
