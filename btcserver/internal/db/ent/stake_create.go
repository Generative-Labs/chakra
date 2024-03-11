// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/generativelabs/btcserver/internal/db/ent/stake"
)

// StakeCreate is the builder for creating a Stake entity.
type StakeCreate struct {
	config
	mutation *StakeMutation
	hooks    []Hook
}

// SetStaker sets the "Staker" field.
func (sc *StakeCreate) SetStaker(s string) *StakeCreate {
	sc.mutation.SetStaker(s)
	return sc
}

// SetStakerPublicKey sets the "StakerPublicKey" field.
func (sc *StakeCreate) SetStakerPublicKey(s string) *StakeCreate {
	sc.mutation.SetStakerPublicKey(s)
	return sc
}

// SetTx sets the "Tx" field.
func (sc *StakeCreate) SetTx(s string) *StakeCreate {
	sc.mutation.SetTx(s)
	return sc
}

// SetStart sets the "Start" field.
func (sc *StakeCreate) SetStart(u uint64) *StakeCreate {
	sc.mutation.SetStart(u)
	return sc
}

// SetDuration sets the "Duration" field.
func (sc *StakeCreate) SetDuration(u uint64) *StakeCreate {
	sc.mutation.SetDuration(u)
	return sc
}

// SetDeadline sets the "Deadline" field.
func (sc *StakeCreate) SetDeadline(u uint64) *StakeCreate {
	sc.mutation.SetDeadline(u)
	return sc
}

// SetAmount sets the "Amount" field.
func (sc *StakeCreate) SetAmount(u uint64) *StakeCreate {
	sc.mutation.SetAmount(u)
	return sc
}

// SetRewardReceiver sets the "RewardReceiver" field.
func (sc *StakeCreate) SetRewardReceiver(s string) *StakeCreate {
	sc.mutation.SetRewardReceiver(s)
	return sc
}

// SetFinalizedStatus sets the "FinalizedStatus" field.
func (sc *StakeCreate) SetFinalizedStatus(i int) *StakeCreate {
	sc.mutation.SetFinalizedStatus(i)
	return sc
}

// SetNillableFinalizedStatus sets the "FinalizedStatus" field if the given value is not nil.
func (sc *StakeCreate) SetNillableFinalizedStatus(i *int) *StakeCreate {
	if i != nil {
		sc.SetFinalizedStatus(*i)
	}
	return sc
}

// SetReleaseStatus sets the "ReleaseStatus" field.
func (sc *StakeCreate) SetReleaseStatus(i int) *StakeCreate {
	sc.mutation.SetReleaseStatus(i)
	return sc
}

// SetNillableReleaseStatus sets the "ReleaseStatus" field if the given value is not nil.
func (sc *StakeCreate) SetNillableReleaseStatus(i *int) *StakeCreate {
	if i != nil {
		sc.SetReleaseStatus(*i)
	}
	return sc
}

// SetBtcSig sets the "BtcSig" field.
func (sc *StakeCreate) SetBtcSig(s string) *StakeCreate {
	sc.mutation.SetBtcSig(s)
	return sc
}

// SetReceiverSig sets the "ReceiverSig" field.
func (sc *StakeCreate) SetReceiverSig(s string) *StakeCreate {
	sc.mutation.SetReceiverSig(s)
	return sc
}

// SetTimestamp sets the "Timestamp" field.
func (sc *StakeCreate) SetTimestamp(u uint64) *StakeCreate {
	sc.mutation.SetTimestamp(u)
	return sc
}

// SetCreateAt sets the "CreateAt" field.
func (sc *StakeCreate) SetCreateAt(u uint64) *StakeCreate {
	sc.mutation.SetCreateAt(u)
	return sc
}

// SetUpdateAt sets the "UpdateAt" field.
func (sc *StakeCreate) SetUpdateAt(u uint64) *StakeCreate {
	sc.mutation.SetUpdateAt(u)
	return sc
}

// Mutation returns the StakeMutation object of the builder.
func (sc *StakeCreate) Mutation() *StakeMutation {
	return sc.mutation
}

// Save creates the Stake in the database.
func (sc *StakeCreate) Save(ctx context.Context) (*Stake, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StakeCreate) SaveX(ctx context.Context) *Stake {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StakeCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StakeCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StakeCreate) defaults() {
	if _, ok := sc.mutation.FinalizedStatus(); !ok {
		v := stake.DefaultFinalizedStatus
		sc.mutation.SetFinalizedStatus(v)
	}
	if _, ok := sc.mutation.ReleaseStatus(); !ok {
		v := stake.DefaultReleaseStatus
		sc.mutation.SetReleaseStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StakeCreate) check() error {
	if _, ok := sc.mutation.Staker(); !ok {
		return &ValidationError{Name: "Staker", err: errors.New(`ent: missing required field "Stake.Staker"`)}
	}
	if v, ok := sc.mutation.Staker(); ok {
		if err := stake.StakerValidator(v); err != nil {
			return &ValidationError{Name: "Staker", err: fmt.Errorf(`ent: validator failed for field "Stake.Staker": %w`, err)}
		}
	}
	if _, ok := sc.mutation.StakerPublicKey(); !ok {
		return &ValidationError{Name: "StakerPublicKey", err: errors.New(`ent: missing required field "Stake.StakerPublicKey"`)}
	}
	if _, ok := sc.mutation.GetTx(); !ok {
		return &ValidationError{Name: "Tx", err: errors.New(`ent: missing required field "Stake.Tx"`)}
	}
	if v, ok := sc.mutation.GetTx(); ok {
		if err := stake.TxValidator(v); err != nil {
			return &ValidationError{Name: "Tx", err: fmt.Errorf(`ent: validator failed for field "Stake.Tx": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Start(); !ok {
		return &ValidationError{Name: "Start", err: errors.New(`ent: missing required field "Stake.Start"`)}
	}
	if _, ok := sc.mutation.Duration(); !ok {
		return &ValidationError{Name: "Duration", err: errors.New(`ent: missing required field "Stake.Duration"`)}
	}
	if _, ok := sc.mutation.Deadline(); !ok {
		return &ValidationError{Name: "Deadline", err: errors.New(`ent: missing required field "Stake.Deadline"`)}
	}
	if _, ok := sc.mutation.Amount(); !ok {
		return &ValidationError{Name: "Amount", err: errors.New(`ent: missing required field "Stake.Amount"`)}
	}
	if _, ok := sc.mutation.RewardReceiver(); !ok {
		return &ValidationError{Name: "RewardReceiver", err: errors.New(`ent: missing required field "Stake.RewardReceiver"`)}
	}
	if v, ok := sc.mutation.RewardReceiver(); ok {
		if err := stake.RewardReceiverValidator(v); err != nil {
			return &ValidationError{Name: "RewardReceiver", err: fmt.Errorf(`ent: validator failed for field "Stake.RewardReceiver": %w`, err)}
		}
	}
	if _, ok := sc.mutation.FinalizedStatus(); !ok {
		return &ValidationError{Name: "FinalizedStatus", err: errors.New(`ent: missing required field "Stake.FinalizedStatus"`)}
	}
	if _, ok := sc.mutation.ReleaseStatus(); !ok {
		return &ValidationError{Name: "ReleaseStatus", err: errors.New(`ent: missing required field "Stake.ReleaseStatus"`)}
	}
	if _, ok := sc.mutation.BtcSig(); !ok {
		return &ValidationError{Name: "BtcSig", err: errors.New(`ent: missing required field "Stake.BtcSig"`)}
	}
	if _, ok := sc.mutation.ReceiverSig(); !ok {
		return &ValidationError{Name: "ReceiverSig", err: errors.New(`ent: missing required field "Stake.ReceiverSig"`)}
	}
	if _, ok := sc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "Timestamp", err: errors.New(`ent: missing required field "Stake.Timestamp"`)}
	}
	if _, ok := sc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "CreateAt", err: errors.New(`ent: missing required field "Stake.CreateAt"`)}
	}
	if _, ok := sc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "UpdateAt", err: errors.New(`ent: missing required field "Stake.UpdateAt"`)}
	}
	return nil
}

func (sc *StakeCreate) sqlSave(ctx context.Context) (*Stake, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StakeCreate) createSpec() (*Stake, *sqlgraph.CreateSpec) {
	var (
		_node = &Stake{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(stake.Table, sqlgraph.NewFieldSpec(stake.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Staker(); ok {
		_spec.SetField(stake.FieldStaker, field.TypeString, value)
		_node.Staker = value
	}
	if value, ok := sc.mutation.StakerPublicKey(); ok {
		_spec.SetField(stake.FieldStakerPublicKey, field.TypeString, value)
		_node.StakerPublicKey = value
	}
	if value, ok := sc.mutation.GetTx(); ok {
		_spec.SetField(stake.FieldTx, field.TypeString, value)
		_node.Tx = value
	}
	if value, ok := sc.mutation.Start(); ok {
		_spec.SetField(stake.FieldStart, field.TypeUint64, value)
		_node.Start = value
	}
	if value, ok := sc.mutation.Duration(); ok {
		_spec.SetField(stake.FieldDuration, field.TypeUint64, value)
		_node.Duration = value
	}
	if value, ok := sc.mutation.Deadline(); ok {
		_spec.SetField(stake.FieldDeadline, field.TypeUint64, value)
		_node.Deadline = value
	}
	if value, ok := sc.mutation.Amount(); ok {
		_spec.SetField(stake.FieldAmount, field.TypeUint64, value)
		_node.Amount = value
	}
	if value, ok := sc.mutation.RewardReceiver(); ok {
		_spec.SetField(stake.FieldRewardReceiver, field.TypeString, value)
		_node.RewardReceiver = value
	}
	if value, ok := sc.mutation.FinalizedStatus(); ok {
		_spec.SetField(stake.FieldFinalizedStatus, field.TypeInt, value)
		_node.FinalizedStatus = value
	}
	if value, ok := sc.mutation.ReleaseStatus(); ok {
		_spec.SetField(stake.FieldReleaseStatus, field.TypeInt, value)
		_node.ReleaseStatus = value
	}
	if value, ok := sc.mutation.BtcSig(); ok {
		_spec.SetField(stake.FieldBtcSig, field.TypeString, value)
		_node.BtcSig = value
	}
	if value, ok := sc.mutation.ReceiverSig(); ok {
		_spec.SetField(stake.FieldReceiverSig, field.TypeString, value)
		_node.ReceiverSig = value
	}
	if value, ok := sc.mutation.Timestamp(); ok {
		_spec.SetField(stake.FieldTimestamp, field.TypeUint64, value)
		_node.Timestamp = value
	}
	if value, ok := sc.mutation.CreateAt(); ok {
		_spec.SetField(stake.FieldCreateAt, field.TypeUint64, value)
		_node.CreateAt = value
	}
	if value, ok := sc.mutation.UpdateAt(); ok {
		_spec.SetField(stake.FieldUpdateAt, field.TypeUint64, value)
		_node.UpdateAt = value
	}
	return _node, _spec
}

// StakeCreateBulk is the builder for creating many Stake entities in bulk.
type StakeCreateBulk struct {
	config
	err      error
	builders []*StakeCreate
}

// Save creates the Stake entities in the database.
func (scb *StakeCreateBulk) Save(ctx context.Context) ([]*Stake, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stake, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StakeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StakeCreateBulk) SaveX(ctx context.Context) []*Stake {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StakeCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StakeCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
