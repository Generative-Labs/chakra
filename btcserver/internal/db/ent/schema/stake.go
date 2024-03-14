package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Stake struct {
	ent.Schema
}

func (Stake) Fields() []ent.Field {
	return []ent.Field{
		field.String("Staker").MaxLen(90), // btc address
		field.String("StakerPublicKey"),
		field.String("Tx").MaxLen(66).Unique(),    // btc transaction id len is 64byte, and len of prefix "0x" is 2byte.
		field.Int64("Start"),                      // btc stake start timestamp
		field.Int64("Duration"),                   // btc stake duration
		field.Int64("Deadline"),                   // btc stake end timestamp
		field.Int64("ReleasingTime"),              // Time to release rewards every day
		field.Int64("Amount"),                     // btc stake amount
		field.String("RewardReceiver").MaxLen(66), // starknet address to receive reward. length is 64byte, and length of prefix "0x" is 2byte.
		field.Int("FinalizedStatus").Default(0),   // btc transaction weather finalized(0 means not on the chain; 1 means it has been uploaded but not finalized; 2 is finalized, 3 is mismatch).
		field.Int("ReleaseStatus").Default(0),     // stake epoch is over(0 means the rewards have not been released yet; 1 means rewards have been released).
		field.String("ReceiverSig"),               // signature for receiver address.
		field.Int64("Timestamp"),
		field.Int64("CreateAt").Immutable(),
		field.Int64("UpdateAt").Default(0),
	}
}

func (Stake) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Staker", "Tx"),
		index.Fields("ReleaseStatus", "Tx"),
		index.Fields("FinalizedStatus", "Tx"),
	}
}
