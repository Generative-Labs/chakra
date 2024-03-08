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
		field.String("Staker").MaxLen(90),          // btc address
		field.String("Tx").MaxLen(66).Unique(),     // btc transaction id len is 64byte, and len of prefix "0x" is 2byte.
		field.Uint64("Start"),                      // btc stake start timestamp
		field.Uint64("Duration"),                   // btc stake end timestamp
		field.Uint64("Amount"),                     // btc stake amount
		field.String("RewardReceiver").MaxLen(66),  // starknet address to receive reward. length is 64byte, and length of prefix "0x" is 2byte.
		field.Bool("FinalizedStatus"),              // btc transaction weather finalized.
		field.Bool("ReleaseStatus").Default(false), // stake epoch is over.
		field.String("BtcSig"),                     // signature for btc transaction.
		field.String("ReceiverSig"),                // signature for receiver address.
		field.Uint64("Timestamp"),
		field.Uint64("CreateAt").Immutable(),
		field.Uint64("UpdateAt").Immutable(),
	}
}

func (Stake) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Staker", "Tx"),
		index.Fields("ReleaseStatus", "Tx"),
	}
}
