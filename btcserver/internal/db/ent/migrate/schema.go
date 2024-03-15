// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GlobalStatesColumns holds the columns for the "global_states" table.
	GlobalStatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64, Default: 0},
	}
	// GlobalStatesTable holds the schema information for the "global_states" table.
	GlobalStatesTable = &schema.Table{
		Name:       "global_states",
		Columns:    GlobalStatesColumns,
		PrimaryKey: []*schema.Column{GlobalStatesColumns[0]},
	}
	// StakesColumns holds the columns for the "stakes" table.
	StakesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "staker", Type: field.TypeString, Size: 90},
		{Name: "staker_public_key", Type: field.TypeString},
		{Name: "tx", Type: field.TypeString, Unique: true, Size: 66},
		{Name: "start", Type: field.TypeInt64},
		{Name: "duration", Type: field.TypeInt64},
		{Name: "deadline", Type: field.TypeInt64},
		{Name: "releasing_time", Type: field.TypeInt64},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "reward_receiver", Type: field.TypeString, Size: 66},
		{Name: "reward", Type: field.TypeUint64},
		{Name: "finalized_status", Type: field.TypeInt, Default: 0},
		{Name: "release_status", Type: field.TypeInt, Default: 0},
		{Name: "submit_status", Type: field.TypeInt, Default: 0},
		{Name: "receiver_sig", Type: field.TypeString},
		{Name: "timestamp", Type: field.TypeInt64},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64, Default: 0},
	}
	// StakesTable holds the schema information for the "stakes" table.
	StakesTable = &schema.Table{
		Name:       "stakes",
		Columns:    StakesColumns,
		PrimaryKey: []*schema.Column{StakesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "stake_staker_tx",
				Unique:  false,
				Columns: []*schema.Column{StakesColumns[1], StakesColumns[3]},
			},
			{
				Name:    "stake_release_status_tx",
				Unique:  false,
				Columns: []*schema.Column{StakesColumns[12], StakesColumns[3]},
			},
			{
				Name:    "stake_finalized_status_tx",
				Unique:  false,
				Columns: []*schema.Column{StakesColumns[11], StakesColumns[3]},
			},
			{
				Name:    "stake_submit_status",
				Unique:  false,
				Columns: []*schema.Column{StakesColumns[13]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GlobalStatesTable,
		StakesTable,
	}
)

func init() {
}
