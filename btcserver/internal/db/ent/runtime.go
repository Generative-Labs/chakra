// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/generativelabs/btcserver/internal/db/ent/schema"
	"github.com/generativelabs/btcserver/internal/db/ent/stake"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	stakeFields := schema.Stake{}.Fields()
	_ = stakeFields
	// stakeDescStaker is the schema descriptor for staker field.
	stakeDescStaker := stakeFields[0].Descriptor()
	// stake.StakerValidator is a validator for the "staker" field. It is called by the builders before save.
	stake.StakerValidator = stakeDescStaker.Validators[0].(func(string) error)
	// stakeDescTx is the schema descriptor for tx field.
	stakeDescTx := stakeFields[1].Descriptor()
	// stake.TxValidator is a validator for the "tx" field. It is called by the builders before save.
	stake.TxValidator = stakeDescTx.Validators[0].(func(string) error)
	// stakeDescRewardReceiver is the schema descriptor for reward_receiver field.
	stakeDescRewardReceiver := stakeFields[5].Descriptor()
	// stake.RewardReceiverValidator is a validator for the "reward_receiver" field. It is called by the builders before save.
	stake.RewardReceiverValidator = stakeDescRewardReceiver.Validators[0].(func(string) error)
	// stakeDescEnd is the schema descriptor for end field.
	stakeDescEnd := stakeFields[6].Descriptor()
	// stake.DefaultEnd holds the default value on creation for the end field.
	stake.DefaultEnd = stakeDescEnd.Default.(bool)
}
