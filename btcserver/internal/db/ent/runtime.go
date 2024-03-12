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
	// stakeDescStaker is the schema descriptor for Staker field.
	stakeDescStaker := stakeFields[0].Descriptor()
	// stake.StakerValidator is a validator for the "Staker" field. It is called by the builders before save.
	stake.StakerValidator = stakeDescStaker.Validators[0].(func(string) error)
	// stakeDescTx is the schema descriptor for Tx field.
	stakeDescTx := stakeFields[2].Descriptor()
	// stake.TxValidator is a validator for the "Tx" field. It is called by the builders before save.
	stake.TxValidator = stakeDescTx.Validators[0].(func(string) error)
	// stakeDescRewardReceiver is the schema descriptor for RewardReceiver field.
	stakeDescRewardReceiver := stakeFields[8].Descriptor()
	// stake.RewardReceiverValidator is a validator for the "RewardReceiver" field. It is called by the builders before save.
	stake.RewardReceiverValidator = stakeDescRewardReceiver.Validators[0].(func(string) error)
	// stakeDescFinalizedStatus is the schema descriptor for FinalizedStatus field.
	stakeDescFinalizedStatus := stakeFields[9].Descriptor()
	// stake.DefaultFinalizedStatus holds the default value on creation for the FinalizedStatus field.
	stake.DefaultFinalizedStatus = stakeDescFinalizedStatus.Default.(int)
	// stakeDescReleaseStatus is the schema descriptor for ReleaseStatus field.
	stakeDescReleaseStatus := stakeFields[10].Descriptor()
	// stake.DefaultReleaseStatus holds the default value on creation for the ReleaseStatus field.
	stake.DefaultReleaseStatus = stakeDescReleaseStatus.Default.(int)
}
