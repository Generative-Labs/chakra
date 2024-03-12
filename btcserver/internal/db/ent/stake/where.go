// Code generated by ent, DO NOT EDIT.

package stake

import (
	"entgo.io/ent/dialect/sql"
	"github.com/generativelabs/btcserver/internal/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldID, id))
}

// Staker applies equality check predicate on the "Staker" field. It's identical to StakerEQ.
func Staker(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldStaker, v))
}

// StakerPublicKey applies equality check predicate on the "StakerPublicKey" field. It's identical to StakerPublicKeyEQ.
func StakerPublicKey(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldStakerPublicKey, v))
}

// Tx applies equality check predicate on the "Tx" field. It's identical to TxEQ.
func Tx(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldTx, v))
}

// Start applies equality check predicate on the "Start" field. It's identical to StartEQ.
func Start(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldStart, v))
}

// Duration applies equality check predicate on the "Duration" field. It's identical to DurationEQ.
func Duration(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldDuration, v))
}

// Deadline applies equality check predicate on the "Deadline" field. It's identical to DeadlineEQ.
func Deadline(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldDeadline, v))
}

// ReleasingTime applies equality check predicate on the "ReleasingTime" field. It's identical to ReleasingTimeEQ.
func ReleasingTime(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldReleasingTime, v))
}

// Amount applies equality check predicate on the "Amount" field. It's identical to AmountEQ.
func Amount(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldAmount, v))
}

// RewardReceiver applies equality check predicate on the "RewardReceiver" field. It's identical to RewardReceiverEQ.
func RewardReceiver(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldRewardReceiver, v))
}

// FinalizedStatus applies equality check predicate on the "FinalizedStatus" field. It's identical to FinalizedStatusEQ.
func FinalizedStatus(v int) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldFinalizedStatus, v))
}

// ReleaseStatus applies equality check predicate on the "ReleaseStatus" field. It's identical to ReleaseStatusEQ.
func ReleaseStatus(v int) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldReleaseStatus, v))
}

// BtcSig applies equality check predicate on the "BtcSig" field. It's identical to BtcSigEQ.
func BtcSig(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldBtcSig, v))
}

// ReceiverSig applies equality check predicate on the "ReceiverSig" field. It's identical to ReceiverSigEQ.
func ReceiverSig(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldReceiverSig, v))
}

// Timestamp applies equality check predicate on the "Timestamp" field. It's identical to TimestampEQ.
func Timestamp(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldTimestamp, v))
}

// CreateAt applies equality check predicate on the "CreateAt" field. It's identical to CreateAtEQ.
func CreateAt(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldCreateAt, v))
}

// UpdateAt applies equality check predicate on the "UpdateAt" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldUpdateAt, v))
}

// StakerEQ applies the EQ predicate on the "Staker" field.
func StakerEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldStaker, v))
}

// StakerNEQ applies the NEQ predicate on the "Staker" field.
func StakerNEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldStaker, v))
}

// StakerIn applies the In predicate on the "Staker" field.
func StakerIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldStaker, vs...))
}

// StakerNotIn applies the NotIn predicate on the "Staker" field.
func StakerNotIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldStaker, vs...))
}

// StakerGT applies the GT predicate on the "Staker" field.
func StakerGT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldStaker, v))
}

// StakerGTE applies the GTE predicate on the "Staker" field.
func StakerGTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldStaker, v))
}

// StakerLT applies the LT predicate on the "Staker" field.
func StakerLT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldStaker, v))
}

// StakerLTE applies the LTE predicate on the "Staker" field.
func StakerLTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldStaker, v))
}

// StakerContains applies the Contains predicate on the "Staker" field.
func StakerContains(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContains(FieldStaker, v))
}

// StakerHasPrefix applies the HasPrefix predicate on the "Staker" field.
func StakerHasPrefix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasPrefix(FieldStaker, v))
}

// StakerHasSuffix applies the HasSuffix predicate on the "Staker" field.
func StakerHasSuffix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasSuffix(FieldStaker, v))
}

// StakerEqualFold applies the EqualFold predicate on the "Staker" field.
func StakerEqualFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEqualFold(FieldStaker, v))
}

// StakerContainsFold applies the ContainsFold predicate on the "Staker" field.
func StakerContainsFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContainsFold(FieldStaker, v))
}

// StakerPublicKeyEQ applies the EQ predicate on the "StakerPublicKey" field.
func StakerPublicKeyEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldStakerPublicKey, v))
}

// StakerPublicKeyNEQ applies the NEQ predicate on the "StakerPublicKey" field.
func StakerPublicKeyNEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldStakerPublicKey, v))
}

// StakerPublicKeyIn applies the In predicate on the "StakerPublicKey" field.
func StakerPublicKeyIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldStakerPublicKey, vs...))
}

// StakerPublicKeyNotIn applies the NotIn predicate on the "StakerPublicKey" field.
func StakerPublicKeyNotIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldStakerPublicKey, vs...))
}

// StakerPublicKeyGT applies the GT predicate on the "StakerPublicKey" field.
func StakerPublicKeyGT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldStakerPublicKey, v))
}

// StakerPublicKeyGTE applies the GTE predicate on the "StakerPublicKey" field.
func StakerPublicKeyGTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldStakerPublicKey, v))
}

// StakerPublicKeyLT applies the LT predicate on the "StakerPublicKey" field.
func StakerPublicKeyLT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldStakerPublicKey, v))
}

// StakerPublicKeyLTE applies the LTE predicate on the "StakerPublicKey" field.
func StakerPublicKeyLTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldStakerPublicKey, v))
}

// StakerPublicKeyContains applies the Contains predicate on the "StakerPublicKey" field.
func StakerPublicKeyContains(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContains(FieldStakerPublicKey, v))
}

// StakerPublicKeyHasPrefix applies the HasPrefix predicate on the "StakerPublicKey" field.
func StakerPublicKeyHasPrefix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasPrefix(FieldStakerPublicKey, v))
}

// StakerPublicKeyHasSuffix applies the HasSuffix predicate on the "StakerPublicKey" field.
func StakerPublicKeyHasSuffix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasSuffix(FieldStakerPublicKey, v))
}

// StakerPublicKeyEqualFold applies the EqualFold predicate on the "StakerPublicKey" field.
func StakerPublicKeyEqualFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEqualFold(FieldStakerPublicKey, v))
}

// StakerPublicKeyContainsFold applies the ContainsFold predicate on the "StakerPublicKey" field.
func StakerPublicKeyContainsFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContainsFold(FieldStakerPublicKey, v))
}

// TxEQ applies the EQ predicate on the "Tx" field.
func TxEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldTx, v))
}

// TxNEQ applies the NEQ predicate on the "Tx" field.
func TxNEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldTx, v))
}

// TxIn applies the In predicate on the "Tx" field.
func TxIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldTx, vs...))
}

// TxNotIn applies the NotIn predicate on the "Tx" field.
func TxNotIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldTx, vs...))
}

// TxGT applies the GT predicate on the "Tx" field.
func TxGT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldTx, v))
}

// TxGTE applies the GTE predicate on the "Tx" field.
func TxGTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldTx, v))
}

// TxLT applies the LT predicate on the "Tx" field.
func TxLT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldTx, v))
}

// TxLTE applies the LTE predicate on the "Tx" field.
func TxLTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldTx, v))
}

// TxContains applies the Contains predicate on the "Tx" field.
func TxContains(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContains(FieldTx, v))
}

// TxHasPrefix applies the HasPrefix predicate on the "Tx" field.
func TxHasPrefix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasPrefix(FieldTx, v))
}

// TxHasSuffix applies the HasSuffix predicate on the "Tx" field.
func TxHasSuffix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasSuffix(FieldTx, v))
}

// TxEqualFold applies the EqualFold predicate on the "Tx" field.
func TxEqualFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEqualFold(FieldTx, v))
}

// TxContainsFold applies the ContainsFold predicate on the "Tx" field.
func TxContainsFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContainsFold(FieldTx, v))
}

// StartEQ applies the EQ predicate on the "Start" field.
func StartEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldStart, v))
}

// StartNEQ applies the NEQ predicate on the "Start" field.
func StartNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldStart, v))
}

// StartIn applies the In predicate on the "Start" field.
func StartIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldStart, vs...))
}

// StartNotIn applies the NotIn predicate on the "Start" field.
func StartNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldStart, vs...))
}

// StartGT applies the GT predicate on the "Start" field.
func StartGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldStart, v))
}

// StartGTE applies the GTE predicate on the "Start" field.
func StartGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldStart, v))
}

// StartLT applies the LT predicate on the "Start" field.
func StartLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldStart, v))
}

// StartLTE applies the LTE predicate on the "Start" field.
func StartLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldStart, v))
}

// DurationEQ applies the EQ predicate on the "Duration" field.
func DurationEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "Duration" field.
func DurationNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "Duration" field.
func DurationIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "Duration" field.
func DurationNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "Duration" field.
func DurationGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "Duration" field.
func DurationGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "Duration" field.
func DurationLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "Duration" field.
func DurationLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldDuration, v))
}

// DeadlineEQ applies the EQ predicate on the "Deadline" field.
func DeadlineEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldDeadline, v))
}

// DeadlineNEQ applies the NEQ predicate on the "Deadline" field.
func DeadlineNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldDeadline, v))
}

// DeadlineIn applies the In predicate on the "Deadline" field.
func DeadlineIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldDeadline, vs...))
}

// DeadlineNotIn applies the NotIn predicate on the "Deadline" field.
func DeadlineNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldDeadline, vs...))
}

// DeadlineGT applies the GT predicate on the "Deadline" field.
func DeadlineGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldDeadline, v))
}

// DeadlineGTE applies the GTE predicate on the "Deadline" field.
func DeadlineGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldDeadline, v))
}

// DeadlineLT applies the LT predicate on the "Deadline" field.
func DeadlineLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldDeadline, v))
}

// DeadlineLTE applies the LTE predicate on the "Deadline" field.
func DeadlineLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldDeadline, v))
}

// ReleasingTimeEQ applies the EQ predicate on the "ReleasingTime" field.
func ReleasingTimeEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldReleasingTime, v))
}

// ReleasingTimeNEQ applies the NEQ predicate on the "ReleasingTime" field.
func ReleasingTimeNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldReleasingTime, v))
}

// ReleasingTimeIn applies the In predicate on the "ReleasingTime" field.
func ReleasingTimeIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldReleasingTime, vs...))
}

// ReleasingTimeNotIn applies the NotIn predicate on the "ReleasingTime" field.
func ReleasingTimeNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldReleasingTime, vs...))
}

// ReleasingTimeGT applies the GT predicate on the "ReleasingTime" field.
func ReleasingTimeGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldReleasingTime, v))
}

// ReleasingTimeGTE applies the GTE predicate on the "ReleasingTime" field.
func ReleasingTimeGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldReleasingTime, v))
}

// ReleasingTimeLT applies the LT predicate on the "ReleasingTime" field.
func ReleasingTimeLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldReleasingTime, v))
}

// ReleasingTimeLTE applies the LTE predicate on the "ReleasingTime" field.
func ReleasingTimeLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldReleasingTime, v))
}

// AmountEQ applies the EQ predicate on the "Amount" field.
func AmountEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "Amount" field.
func AmountNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "Amount" field.
func AmountIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "Amount" field.
func AmountNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "Amount" field.
func AmountGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "Amount" field.
func AmountGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "Amount" field.
func AmountLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "Amount" field.
func AmountLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldAmount, v))
}

// RewardReceiverEQ applies the EQ predicate on the "RewardReceiver" field.
func RewardReceiverEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldRewardReceiver, v))
}

// RewardReceiverNEQ applies the NEQ predicate on the "RewardReceiver" field.
func RewardReceiverNEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldRewardReceiver, v))
}

// RewardReceiverIn applies the In predicate on the "RewardReceiver" field.
func RewardReceiverIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldRewardReceiver, vs...))
}

// RewardReceiverNotIn applies the NotIn predicate on the "RewardReceiver" field.
func RewardReceiverNotIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldRewardReceiver, vs...))
}

// RewardReceiverGT applies the GT predicate on the "RewardReceiver" field.
func RewardReceiverGT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldRewardReceiver, v))
}

// RewardReceiverGTE applies the GTE predicate on the "RewardReceiver" field.
func RewardReceiverGTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldRewardReceiver, v))
}

// RewardReceiverLT applies the LT predicate on the "RewardReceiver" field.
func RewardReceiverLT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldRewardReceiver, v))
}

// RewardReceiverLTE applies the LTE predicate on the "RewardReceiver" field.
func RewardReceiverLTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldRewardReceiver, v))
}

// RewardReceiverContains applies the Contains predicate on the "RewardReceiver" field.
func RewardReceiverContains(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContains(FieldRewardReceiver, v))
}

// RewardReceiverHasPrefix applies the HasPrefix predicate on the "RewardReceiver" field.
func RewardReceiverHasPrefix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasPrefix(FieldRewardReceiver, v))
}

// RewardReceiverHasSuffix applies the HasSuffix predicate on the "RewardReceiver" field.
func RewardReceiverHasSuffix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasSuffix(FieldRewardReceiver, v))
}

// RewardReceiverEqualFold applies the EqualFold predicate on the "RewardReceiver" field.
func RewardReceiverEqualFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEqualFold(FieldRewardReceiver, v))
}

// RewardReceiverContainsFold applies the ContainsFold predicate on the "RewardReceiver" field.
func RewardReceiverContainsFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContainsFold(FieldRewardReceiver, v))
}

// FinalizedStatusEQ applies the EQ predicate on the "FinalizedStatus" field.
func FinalizedStatusEQ(v int) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldFinalizedStatus, v))
}

// FinalizedStatusNEQ applies the NEQ predicate on the "FinalizedStatus" field.
func FinalizedStatusNEQ(v int) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldFinalizedStatus, v))
}

// FinalizedStatusIn applies the In predicate on the "FinalizedStatus" field.
func FinalizedStatusIn(vs ...int) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldFinalizedStatus, vs...))
}

// FinalizedStatusNotIn applies the NotIn predicate on the "FinalizedStatus" field.
func FinalizedStatusNotIn(vs ...int) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldFinalizedStatus, vs...))
}

// FinalizedStatusGT applies the GT predicate on the "FinalizedStatus" field.
func FinalizedStatusGT(v int) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldFinalizedStatus, v))
}

// FinalizedStatusGTE applies the GTE predicate on the "FinalizedStatus" field.
func FinalizedStatusGTE(v int) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldFinalizedStatus, v))
}

// FinalizedStatusLT applies the LT predicate on the "FinalizedStatus" field.
func FinalizedStatusLT(v int) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldFinalizedStatus, v))
}

// FinalizedStatusLTE applies the LTE predicate on the "FinalizedStatus" field.
func FinalizedStatusLTE(v int) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldFinalizedStatus, v))
}

// ReleaseStatusEQ applies the EQ predicate on the "ReleaseStatus" field.
func ReleaseStatusEQ(v int) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldReleaseStatus, v))
}

// ReleaseStatusNEQ applies the NEQ predicate on the "ReleaseStatus" field.
func ReleaseStatusNEQ(v int) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldReleaseStatus, v))
}

// ReleaseStatusIn applies the In predicate on the "ReleaseStatus" field.
func ReleaseStatusIn(vs ...int) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldReleaseStatus, vs...))
}

// ReleaseStatusNotIn applies the NotIn predicate on the "ReleaseStatus" field.
func ReleaseStatusNotIn(vs ...int) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldReleaseStatus, vs...))
}

// ReleaseStatusGT applies the GT predicate on the "ReleaseStatus" field.
func ReleaseStatusGT(v int) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldReleaseStatus, v))
}

// ReleaseStatusGTE applies the GTE predicate on the "ReleaseStatus" field.
func ReleaseStatusGTE(v int) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldReleaseStatus, v))
}

// ReleaseStatusLT applies the LT predicate on the "ReleaseStatus" field.
func ReleaseStatusLT(v int) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldReleaseStatus, v))
}

// ReleaseStatusLTE applies the LTE predicate on the "ReleaseStatus" field.
func ReleaseStatusLTE(v int) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldReleaseStatus, v))
}

// BtcSigEQ applies the EQ predicate on the "BtcSig" field.
func BtcSigEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldBtcSig, v))
}

// BtcSigNEQ applies the NEQ predicate on the "BtcSig" field.
func BtcSigNEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldBtcSig, v))
}

// BtcSigIn applies the In predicate on the "BtcSig" field.
func BtcSigIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldBtcSig, vs...))
}

// BtcSigNotIn applies the NotIn predicate on the "BtcSig" field.
func BtcSigNotIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldBtcSig, vs...))
}

// BtcSigGT applies the GT predicate on the "BtcSig" field.
func BtcSigGT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldBtcSig, v))
}

// BtcSigGTE applies the GTE predicate on the "BtcSig" field.
func BtcSigGTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldBtcSig, v))
}

// BtcSigLT applies the LT predicate on the "BtcSig" field.
func BtcSigLT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldBtcSig, v))
}

// BtcSigLTE applies the LTE predicate on the "BtcSig" field.
func BtcSigLTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldBtcSig, v))
}

// BtcSigContains applies the Contains predicate on the "BtcSig" field.
func BtcSigContains(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContains(FieldBtcSig, v))
}

// BtcSigHasPrefix applies the HasPrefix predicate on the "BtcSig" field.
func BtcSigHasPrefix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasPrefix(FieldBtcSig, v))
}

// BtcSigHasSuffix applies the HasSuffix predicate on the "BtcSig" field.
func BtcSigHasSuffix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasSuffix(FieldBtcSig, v))
}

// BtcSigEqualFold applies the EqualFold predicate on the "BtcSig" field.
func BtcSigEqualFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEqualFold(FieldBtcSig, v))
}

// BtcSigContainsFold applies the ContainsFold predicate on the "BtcSig" field.
func BtcSigContainsFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContainsFold(FieldBtcSig, v))
}

// ReceiverSigEQ applies the EQ predicate on the "ReceiverSig" field.
func ReceiverSigEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldReceiverSig, v))
}

// ReceiverSigNEQ applies the NEQ predicate on the "ReceiverSig" field.
func ReceiverSigNEQ(v string) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldReceiverSig, v))
}

// ReceiverSigIn applies the In predicate on the "ReceiverSig" field.
func ReceiverSigIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldReceiverSig, vs...))
}

// ReceiverSigNotIn applies the NotIn predicate on the "ReceiverSig" field.
func ReceiverSigNotIn(vs ...string) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldReceiverSig, vs...))
}

// ReceiverSigGT applies the GT predicate on the "ReceiverSig" field.
func ReceiverSigGT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldReceiverSig, v))
}

// ReceiverSigGTE applies the GTE predicate on the "ReceiverSig" field.
func ReceiverSigGTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldReceiverSig, v))
}

// ReceiverSigLT applies the LT predicate on the "ReceiverSig" field.
func ReceiverSigLT(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldReceiverSig, v))
}

// ReceiverSigLTE applies the LTE predicate on the "ReceiverSig" field.
func ReceiverSigLTE(v string) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldReceiverSig, v))
}

// ReceiverSigContains applies the Contains predicate on the "ReceiverSig" field.
func ReceiverSigContains(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContains(FieldReceiverSig, v))
}

// ReceiverSigHasPrefix applies the HasPrefix predicate on the "ReceiverSig" field.
func ReceiverSigHasPrefix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasPrefix(FieldReceiverSig, v))
}

// ReceiverSigHasSuffix applies the HasSuffix predicate on the "ReceiverSig" field.
func ReceiverSigHasSuffix(v string) predicate.Stake {
	return predicate.Stake(sql.FieldHasSuffix(FieldReceiverSig, v))
}

// ReceiverSigEqualFold applies the EqualFold predicate on the "ReceiverSig" field.
func ReceiverSigEqualFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldEqualFold(FieldReceiverSig, v))
}

// ReceiverSigContainsFold applies the ContainsFold predicate on the "ReceiverSig" field.
func ReceiverSigContainsFold(v string) predicate.Stake {
	return predicate.Stake(sql.FieldContainsFold(FieldReceiverSig, v))
}

// TimestampEQ applies the EQ predicate on the "Timestamp" field.
func TimestampEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "Timestamp" field.
func TimestampNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "Timestamp" field.
func TimestampIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "Timestamp" field.
func TimestampNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "Timestamp" field.
func TimestampGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "Timestamp" field.
func TimestampGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "Timestamp" field.
func TimestampLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "Timestamp" field.
func TimestampLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldTimestamp, v))
}

// CreateAtEQ applies the EQ predicate on the "CreateAt" field.
func CreateAtEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldCreateAt, v))
}

// CreateAtNEQ applies the NEQ predicate on the "CreateAt" field.
func CreateAtNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldCreateAt, v))
}

// CreateAtIn applies the In predicate on the "CreateAt" field.
func CreateAtIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldCreateAt, vs...))
}

// CreateAtNotIn applies the NotIn predicate on the "CreateAt" field.
func CreateAtNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldCreateAt, vs...))
}

// CreateAtGT applies the GT predicate on the "CreateAt" field.
func CreateAtGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldCreateAt, v))
}

// CreateAtGTE applies the GTE predicate on the "CreateAt" field.
func CreateAtGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldCreateAt, v))
}

// CreateAtLT applies the LT predicate on the "CreateAt" field.
func CreateAtLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldCreateAt, v))
}

// CreateAtLTE applies the LTE predicate on the "CreateAt" field.
func CreateAtLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldCreateAt, v))
}

// UpdateAtEQ applies the EQ predicate on the "UpdateAt" field.
func UpdateAtEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldEQ(FieldUpdateAt, v))
}

// UpdateAtNEQ applies the NEQ predicate on the "UpdateAt" field.
func UpdateAtNEQ(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNEQ(FieldUpdateAt, v))
}

// UpdateAtIn applies the In predicate on the "UpdateAt" field.
func UpdateAtIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldIn(FieldUpdateAt, vs...))
}

// UpdateAtNotIn applies the NotIn predicate on the "UpdateAt" field.
func UpdateAtNotIn(vs ...uint64) predicate.Stake {
	return predicate.Stake(sql.FieldNotIn(FieldUpdateAt, vs...))
}

// UpdateAtGT applies the GT predicate on the "UpdateAt" field.
func UpdateAtGT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGT(FieldUpdateAt, v))
}

// UpdateAtGTE applies the GTE predicate on the "UpdateAt" field.
func UpdateAtGTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldGTE(FieldUpdateAt, v))
}

// UpdateAtLT applies the LT predicate on the "UpdateAt" field.
func UpdateAtLT(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLT(FieldUpdateAt, v))
}

// UpdateAtLTE applies the LTE predicate on the "UpdateAt" field.
func UpdateAtLTE(v uint64) predicate.Stake {
	return predicate.Stake(sql.FieldLTE(FieldUpdateAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Stake) predicate.Stake {
	return predicate.Stake(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Stake) predicate.Stake {
	return predicate.Stake(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Stake) predicate.Stake {
	return predicate.Stake(sql.NotPredicates(p))
}
