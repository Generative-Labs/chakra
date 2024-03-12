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
		field.String("Value"),
		field.Int64("CreateAt").Immutable(),
		field.Int64("UpdateAt").Default(0),
	}
}
