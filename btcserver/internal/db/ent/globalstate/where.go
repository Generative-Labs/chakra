// Code generated by ent, DO NOT EDIT.

package globalstate

import (
	"entgo.io/ent/dialect/sql"
	"github.com/generativelabs/btcserver/internal/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLTE(FieldID, id))
}

// Key applies equality check predicate on the "Key" field. It's identical to KeyEQ.
func Key(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "Value" field. It's identical to ValueEQ.
func Value(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldValue, v))
}

// CreateAt applies equality check predicate on the "CreateAt" field. It's identical to CreateAtEQ.
func CreateAt(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "UpdateAt" field. It's identical to UpdateAtEQ.
func UpdateAt(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldUpdateAt, v))
}

// KeyEQ applies the EQ predicate on the "Key" field.
func KeyEQ(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "Key" field.
func KeyNEQ(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "Key" field.
func KeyIn(vs ...string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "Key" field.
func KeyNotIn(vs ...string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "Key" field.
func KeyGT(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "Key" field.
func KeyGTE(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "Key" field.
func KeyLT(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "Key" field.
func KeyLTE(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "Key" field.
func KeyContains(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "Key" field.
func KeyHasPrefix(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "Key" field.
func KeyHasSuffix(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "Key" field.
func KeyEqualFold(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "Key" field.
func KeyContainsFold(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "Value" field.
func ValueEQ(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "Value" field.
func ValueNEQ(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "Value" field.
func ValueIn(vs ...string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "Value" field.
func ValueNotIn(vs ...string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "Value" field.
func ValueGT(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "Value" field.
func ValueGTE(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "Value" field.
func ValueLT(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "Value" field.
func ValueLTE(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "Value" field.
func ValueContains(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "Value" field.
func ValueHasPrefix(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "Value" field.
func ValueHasSuffix(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "Value" field.
func ValueEqualFold(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "Value" field.
func ValueContainsFold(v string) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldContainsFold(FieldValue, v))
}

// CreateAtEQ applies the EQ predicate on the "CreateAt" field.
func CreateAtEQ(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "CreateAt" field.
func CreateAtNEQ(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "CreateAt" field.
func CreateAtIn(vs ...int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "CreateAt" field.
func CreateAtNotIn(vs ...int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "CreateAt" field.
func CreateAtGT(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "CreateAt" field.
func CreateAtGTE(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "CreateAt" field.
func CreateAtLT(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "CreateAt" field.
func CreateAtLTE(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "UpdateAt" field.
func UpdateAtEQ(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "UpdateAt" field.
func UpdateAtNEQ(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "UpdateAt" field.
func UpdateAtIn(vs ...int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "UpdateAt" field.
func UpdateAtNotIn(vs ...int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "UpdateAt" field.
func UpdateAtGT(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "UpdateAt" field.
func UpdateAtGTE(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "UpdateAt" field.
func UpdateAtLT(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "UpdateAt" field.
func UpdateAtLTE(v int64) predicate.GlobalState {
	return predicate.GlobalState(sql.FieldLTE(FieldUpdateAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GlobalState) predicate.GlobalState {
	return predicate.GlobalState(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GlobalState) predicate.GlobalState {
	return predicate.GlobalState(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GlobalState) predicate.GlobalState {
	return predicate.GlobalState(sql.NotPredicates(p))
}
