// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/generativelabs/btcserver/internal/db/ent/stake"
)

// Stake is the model entity for the Stake schema.
type Stake struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Staker holds the value of the "Staker" field.
	Staker string `json:"Staker,omitempty"`
	// StakerPublicKey holds the value of the "StakerPublicKey" field.
	StakerPublicKey string `json:"StakerPublicKey,omitempty"`
	// Tx holds the value of the "Tx" field.
	Tx string `json:"Tx,omitempty"`
	// Start holds the value of the "Start" field.
	Start uint64 `json:"Start,omitempty"`
	// Duration holds the value of the "Duration" field.
	Duration uint64 `json:"Duration,omitempty"`
	// Deadline holds the value of the "Deadline" field.
	Deadline uint64 `json:"Deadline,omitempty"`
	// Amount holds the value of the "Amount" field.
	Amount uint64 `json:"Amount,omitempty"`
	// RewardReceiver holds the value of the "RewardReceiver" field.
	RewardReceiver string `json:"RewardReceiver,omitempty"`
	// FinalizedStatus holds the value of the "FinalizedStatus" field.
	FinalizedStatus int `json:"FinalizedStatus,omitempty"`
	// ReleaseStatus holds the value of the "ReleaseStatus" field.
	ReleaseStatus int `json:"ReleaseStatus,omitempty"`
	// BtcSig holds the value of the "BtcSig" field.
	BtcSig string `json:"BtcSig,omitempty"`
	// ReceiverSig holds the value of the "ReceiverSig" field.
	ReceiverSig string `json:"ReceiverSig,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp uint64 `json:"Timestamp,omitempty"`
	// CreateAt holds the value of the "CreateAt" field.
	CreateAt uint64 `json:"CreateAt,omitempty"`
	// UpdateAt holds the value of the "UpdateAt" field.
	UpdateAt     uint64 `json:"UpdateAt,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Stake) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case stake.FieldID, stake.FieldStart, stake.FieldDuration, stake.FieldDeadline, stake.FieldAmount, stake.FieldFinalizedStatus, stake.FieldReleaseStatus, stake.FieldTimestamp, stake.FieldCreateAt, stake.FieldUpdateAt:
			values[i] = new(sql.NullInt64)
		case stake.FieldStaker, stake.FieldStakerPublicKey, stake.FieldTx, stake.FieldRewardReceiver, stake.FieldBtcSig, stake.FieldReceiverSig:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Stake fields.
func (s *Stake) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case stake.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case stake.FieldStaker:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Staker", values[i])
			} else if value.Valid {
				s.Staker = value.String
			}
		case stake.FieldStakerPublicKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field StakerPublicKey", values[i])
			} else if value.Valid {
				s.StakerPublicKey = value.String
			}
		case stake.FieldTx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Tx", values[i])
			} else if value.Valid {
				s.Tx = value.String
			}
		case stake.FieldStart:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Start", values[i])
			} else if value.Valid {
				s.Start = uint64(value.Int64)
			}
		case stake.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Duration", values[i])
			} else if value.Valid {
				s.Duration = uint64(value.Int64)
			}
		case stake.FieldDeadline:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Deadline", values[i])
			} else if value.Valid {
				s.Deadline = uint64(value.Int64)
			}
		case stake.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Amount", values[i])
			} else if value.Valid {
				s.Amount = uint64(value.Int64)
			}
		case stake.FieldRewardReceiver:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field RewardReceiver", values[i])
			} else if value.Valid {
				s.RewardReceiver = value.String
			}
		case stake.FieldFinalizedStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field FinalizedStatus", values[i])
			} else if value.Valid {
				s.FinalizedStatus = int(value.Int64)
			}
		case stake.FieldReleaseStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ReleaseStatus", values[i])
			} else if value.Valid {
				s.ReleaseStatus = int(value.Int64)
			}
		case stake.FieldBtcSig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BtcSig", values[i])
			} else if value.Valid {
				s.BtcSig = value.String
			}
		case stake.FieldReceiverSig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ReceiverSig", values[i])
			} else if value.Valid {
				s.ReceiverSig = value.String
			}
		case stake.FieldTimestamp:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				s.Timestamp = uint64(value.Int64)
			}
		case stake.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CreateAt", values[i])
			} else if value.Valid {
				s.CreateAt = uint64(value.Int64)
			}
		case stake.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field UpdateAt", values[i])
			} else if value.Valid {
				s.UpdateAt = uint64(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Stake.
// This includes values selected through modifiers, order, etc.
func (s *Stake) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Stake.
// Note that you need to call Stake.Unwrap() before calling this method if this Stake
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Stake) Update() *StakeUpdateOne {
	return NewStakeClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Stake entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Stake) Unwrap() *Stake {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Stake is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Stake) String() string {
	var builder strings.Builder
	builder.WriteString("Stake(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("Staker=")
	builder.WriteString(s.Staker)
	builder.WriteString(", ")
	builder.WriteString("StakerPublicKey=")
	builder.WriteString(s.StakerPublicKey)
	builder.WriteString(", ")
	builder.WriteString("Tx=")
	builder.WriteString(s.Tx)
	builder.WriteString(", ")
	builder.WriteString("Start=")
	builder.WriteString(fmt.Sprintf("%v", s.Start))
	builder.WriteString(", ")
	builder.WriteString("Duration=")
	builder.WriteString(fmt.Sprintf("%v", s.Duration))
	builder.WriteString(", ")
	builder.WriteString("Deadline=")
	builder.WriteString(fmt.Sprintf("%v", s.Deadline))
	builder.WriteString(", ")
	builder.WriteString("Amount=")
	builder.WriteString(fmt.Sprintf("%v", s.Amount))
	builder.WriteString(", ")
	builder.WriteString("RewardReceiver=")
	builder.WriteString(s.RewardReceiver)
	builder.WriteString(", ")
	builder.WriteString("FinalizedStatus=")
	builder.WriteString(fmt.Sprintf("%v", s.FinalizedStatus))
	builder.WriteString(", ")
	builder.WriteString("ReleaseStatus=")
	builder.WriteString(fmt.Sprintf("%v", s.ReleaseStatus))
	builder.WriteString(", ")
	builder.WriteString("BtcSig=")
	builder.WriteString(s.BtcSig)
	builder.WriteString(", ")
	builder.WriteString("ReceiverSig=")
	builder.WriteString(s.ReceiverSig)
	builder.WriteString(", ")
	builder.WriteString("Timestamp=")
	builder.WriteString(fmt.Sprintf("%v", s.Timestamp))
	builder.WriteString(", ")
	builder.WriteString("CreateAt=")
	builder.WriteString(fmt.Sprintf("%v", s.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("UpdateAt=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdateAt))
	builder.WriteByte(')')
	return builder.String()
}

// Stakes is a parsable slice of Stake.
type Stakes []*Stake
