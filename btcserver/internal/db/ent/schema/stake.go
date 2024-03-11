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
		field.Uint64("Start"),                     // btc stake start timestamp
		field.Uint64("Duration"),                  // btc stake duration
		field.Uint64("Deadline"),                  //  btc stake end timestamp
		field.Uint64("ReleasingTime"),             // Time to release rewards every day
		field.Uint64("Amount"),                    // btc stake amount
		field.String("RewardReceiver").MaxLen(66), // starknet address to receive reward. length is 64byte, and length of prefix "0x" is 2byte.
		field.Int("FinalizedStatus").Default(0),   // btc transaction weather finalized(0 means not on the chain; 1 means it has been uploaded but not finalized; 2 is finalized).
		field.Int("ReleaseStatus").Default(0),     // stake epoch is over(0 means the rewards have not been released yet; 1 means rewards have been released).
		field.String("BtcSig"),                    // signature for btc transaction.
		field.String("ReceiverSig"),               // signature for receiver address.
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
