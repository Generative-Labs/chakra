package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// StakeIndex holds the schema definition for the StakeIndex entity.
type StakeIndex struct {
	ent.Schema
}

// Fields of the StakeIndex.
func (StakeIndex) Fields() []ent.Field {
	return []ent.Field{
		field.String("Staker").MaxLen(90).Unique(), // btc address
		field.String("Tx").MaxLen(66).Unique(),     // btc transaction id len is 64byte, and len of prefix "0x" is 2byte.
		field.Uint64("Index").Default(0),
		field.Int64("Start").Default(0),
		field.Int64("CreateAt").Immutable(),
		field.Int64("UpdateAt").Default(0),
	}
}

// Edges of the StakeIndex.
func (StakeIndex) Edges() []ent.Edge {
	return nil
}
