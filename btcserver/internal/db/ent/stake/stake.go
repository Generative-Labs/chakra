// Code generated by ent, DO NOT EDIT.

package stake

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the stake type in the database.
	Label = "stake"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStaker holds the string denoting the staker field in the database.
	FieldStaker = "staker"
	// FieldStakerPublicKey holds the string denoting the stakerpublickey field in the database.
	FieldStakerPublicKey = "staker_public_key"
	// FieldTx holds the string denoting the tx field in the database.
	FieldTx = "tx"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// FieldReleasingTime holds the string denoting the releasingtime field in the database.
	FieldReleasingTime = "releasing_time"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldRewardReceiver holds the string denoting the rewardreceiver field in the database.
	FieldRewardReceiver = "reward_receiver"
	// FieldReward holds the string denoting the reward field in the database.
	FieldReward = "reward"
	// FieldFinalizedStatus holds the string denoting the finalizedstatus field in the database.
	FieldFinalizedStatus = "finalized_status"
	// FieldReleaseStatus holds the string denoting the releasestatus field in the database.
	FieldReleaseStatus = "release_status"
	// FieldSubmitStatus holds the string denoting the submitstatus field in the database.
	FieldSubmitStatus = "submit_status"
	// FieldReceiverSig holds the string denoting the receiversig field in the database.
	FieldReceiverSig = "receiver_sig"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldCreateAt holds the string denoting the createat field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the updateat field in the database.
	FieldUpdateAt = "update_at"
	// Table holds the table name of the stake in the database.
	Table = "stakes"
)

// Columns holds all SQL columns for stake fields.
var Columns = []string{
	FieldID,
	FieldStaker,
	FieldStakerPublicKey,
	FieldTx,
	FieldStart,
	FieldDuration,
	FieldDeadline,
	FieldReleasingTime,
	FieldAmount,
	FieldRewardReceiver,
	FieldReward,
	FieldFinalizedStatus,
	FieldReleaseStatus,
	FieldSubmitStatus,
	FieldReceiverSig,
	FieldTimestamp,
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
	// StakerValidator is a validator for the "Staker" field. It is called by the builders before save.
	StakerValidator func(string) error
	// TxValidator is a validator for the "Tx" field. It is called by the builders before save.
	TxValidator func(string) error
	// DefaultStart holds the default value on creation for the "Start" field.
	DefaultStart int64
	// DefaultDeadline holds the default value on creation for the "Deadline" field.
	DefaultDeadline int64
	// DefaultReleasingTime holds the default value on creation for the "ReleasingTime" field.
	DefaultReleasingTime int64
	// RewardReceiverValidator is a validator for the "RewardReceiver" field. It is called by the builders before save.
	RewardReceiverValidator func(string) error
	// DefaultReward holds the default value on creation for the "Reward" field.
	DefaultReward uint64
	// DefaultFinalizedStatus holds the default value on creation for the "FinalizedStatus" field.
	DefaultFinalizedStatus int
	// DefaultReleaseStatus holds the default value on creation for the "ReleaseStatus" field.
	DefaultReleaseStatus int
	// DefaultSubmitStatus holds the default value on creation for the "SubmitStatus" field.
	DefaultSubmitStatus int
	// DefaultUpdateAt holds the default value on creation for the "UpdateAt" field.
	DefaultUpdateAt int64
)

// OrderOption defines the ordering options for the Stake queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStaker orders the results by the Staker field.
func ByStaker(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStaker, opts...).ToFunc()
}

// ByStakerPublicKey orders the results by the StakerPublicKey field.
func ByStakerPublicKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStakerPublicKey, opts...).ToFunc()
}

// ByTx orders the results by the Tx field.
func ByTx(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTx, opts...).ToFunc()
}

// ByStart orders the results by the Start field.
func ByStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStart, opts...).ToFunc()
}

// ByDuration orders the results by the Duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByDeadline orders the results by the Deadline field.
func ByDeadline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeadline, opts...).ToFunc()
}

// ByReleasingTime orders the results by the ReleasingTime field.
func ByReleasingTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleasingTime, opts...).ToFunc()
}

// ByAmount orders the results by the Amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByRewardReceiver orders the results by the RewardReceiver field.
func ByRewardReceiver(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRewardReceiver, opts...).ToFunc()
}

// ByReward orders the results by the Reward field.
func ByReward(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReward, opts...).ToFunc()
}

// ByFinalizedStatus orders the results by the FinalizedStatus field.
func ByFinalizedStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFinalizedStatus, opts...).ToFunc()
}

// ByReleaseStatus orders the results by the ReleaseStatus field.
func ByReleaseStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleaseStatus, opts...).ToFunc()
}

// BySubmitStatus orders the results by the SubmitStatus field.
func BySubmitStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubmitStatus, opts...).ToFunc()
}

// ByReceiverSig orders the results by the ReceiverSig field.
func ByReceiverSig(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceiverSig, opts...).ToFunc()
}

// ByTimestamp orders the results by the Timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// ByCreateAt orders the results by the CreateAt field.
func ByCreateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateAt, opts...).ToFunc()
}

// ByUpdateAt orders the results by the UpdateAt field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}
