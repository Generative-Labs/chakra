package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type GlobalState struct {
	ent.Schema
}

func (GlobalState) Fields() []ent.Field {
	return []ent.Field{
		field.String("Key"),
		field.Uint64("Value"),
		field.Uint64("CreateAt").Immutable(),
		field.Uint64("UpdateAt").Immutable(),
	}
}
