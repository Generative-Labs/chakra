// Code generated by ent, DO NOT EDIT.

package globalstate

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the globalstate type in the database.
	Label = "global_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldCreateAt holds the string denoting the createat field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the updateat field in the database.
	FieldUpdateAt = "update_at"
	// Table holds the table name of the globalstate in the database.
	Table = "global_states"
)

// Columns holds all SQL columns for globalstate fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldValue,
	FieldCreateAt,
	FieldUpdateAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdateAt holds the default value on creation for the "UpdateAt" field.
	DefaultUpdateAt int64
)

// OrderOption defines the ordering options for the GlobalState queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKey orders the results by the Key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByValue orders the results by the Value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByCreateAt orders the results by the CreateAt field.
func ByCreateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateAt, opts...).ToFunc()
}

// ByUpdateAt orders the results by the UpdateAt field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}
