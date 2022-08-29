// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"rei.io/rei/ent/transactions"
)

// TransactionsCreate is the builder for creating a Transactions entity.
type TransactionsCreate struct {
	config
	mutation *TransactionsMutation
	hooks    []Hook
}

// SetType sets the "Type" field.
func (tc *TransactionsCreate) SetType(s string) *TransactionsCreate {
	tc.mutation.SetType(s)
	return tc
}

// SetTime sets the "Time" field.
func (tc *TransactionsCreate) SetTime(t time.Time) *TransactionsCreate {
	tc.mutation.SetTime(t)
	return tc
}

// SetTransactionID sets the "TransactionID" field.
func (tc *TransactionsCreate) SetTransactionID(s string) *TransactionsCreate {
	tc.mutation.SetTransactionID(s)
	return tc
}

// SetStatus sets the "Status" field.
func (tc *TransactionsCreate) SetStatus(b bool) *TransactionsCreate {
	tc.mutation.SetStatus(b)
	return tc
}

// SetSender sets the "Sender" field.
func (tc *TransactionsCreate) SetSender(s string) *TransactionsCreate {
	tc.mutation.SetSender(s)
	return tc
}

// SetRecipient sets the "Recipient" field.
func (tc *TransactionsCreate) SetRecipient(s string) *TransactionsCreate {
	tc.mutation.SetRecipient(s)
	return tc
}

// SetNillableRecipient sets the "Recipient" field if the given value is not nil.
func (tc *TransactionsCreate) SetNillableRecipient(s *string) *TransactionsCreate {
	if s != nil {
		tc.SetRecipient(*s)
	}
	return tc
}

// SetAmount sets the "Amount" field.
func (tc *TransactionsCreate) SetAmount(f float64) *TransactionsCreate {
	tc.mutation.SetAmount(f)
	return tc
}

// SetNillableAmount sets the "Amount" field if the given value is not nil.
func (tc *TransactionsCreate) SetNillableAmount(f *float64) *TransactionsCreate {
	if f != nil {
		tc.SetAmount(*f)
	}
	return tc
}

// SetPackage sets the "Package" field.
func (tc *TransactionsCreate) SetPackage(s string) *TransactionsCreate {
	tc.mutation.SetPackage(s)
	return tc
}

// SetNillablePackage sets the "Package" field if the given value is not nil.
func (tc *TransactionsCreate) SetNillablePackage(s *string) *TransactionsCreate {
	if s != nil {
		tc.SetPackage(*s)
	}
	return tc
}

// SetModule sets the "Module" field.
func (tc *TransactionsCreate) SetModule(s string) *TransactionsCreate {
	tc.mutation.SetModule(s)
	return tc
}

// SetNillableModule sets the "Module" field if the given value is not nil.
func (tc *TransactionsCreate) SetNillableModule(s *string) *TransactionsCreate {
	if s != nil {
		tc.SetModule(*s)
	}
	return tc
}

// SetFunction sets the "Function" field.
func (tc *TransactionsCreate) SetFunction(s string) *TransactionsCreate {
	tc.mutation.SetFunction(s)
	return tc
}

// SetNillableFunction sets the "Function" field if the given value is not nil.
func (tc *TransactionsCreate) SetNillableFunction(s *string) *TransactionsCreate {
	if s != nil {
		tc.SetFunction(*s)
	}
	return tc
}

// SetGas sets the "Gas" field.
func (tc *TransactionsCreate) SetGas(u uint32) *TransactionsCreate {
	tc.mutation.SetGas(u)
	return tc
}

// Mutation returns the TransactionsMutation object of the builder.
func (tc *TransactionsCreate) Mutation() *TransactionsMutation {
	return tc.mutation
}

// Save creates the Transactions in the database.
func (tc *TransactionsCreate) Save(ctx context.Context) (*Transactions, error) {
	var (
		err  error
		node *Transactions
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Transactions)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TransactionsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionsCreate) SaveX(ctx context.Context) *Transactions {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionsCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionsCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionsCreate) check() error {
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "Type", err: errors.New(`ent: missing required field "Transactions.Type"`)}
	}
	if _, ok := tc.mutation.Time(); !ok {
		return &ValidationError{Name: "Time", err: errors.New(`ent: missing required field "Transactions.Time"`)}
	}
	if _, ok := tc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "TransactionID", err: errors.New(`ent: missing required field "Transactions.TransactionID"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "Status", err: errors.New(`ent: missing required field "Transactions.Status"`)}
	}
	if _, ok := tc.mutation.Sender(); !ok {
		return &ValidationError{Name: "Sender", err: errors.New(`ent: missing required field "Transactions.Sender"`)}
	}
	if _, ok := tc.mutation.Gas(); !ok {
		return &ValidationError{Name: "Gas", err: errors.New(`ent: missing required field "Transactions.Gas"`)}
	}
	return nil
}

func (tc *TransactionsCreate) sqlSave(ctx context.Context) (*Transactions, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TransactionsCreate) createSpec() (*Transactions, *sqlgraph.CreateSpec) {
	var (
		_node = &Transactions{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transactions.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactions.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactions.FieldType,
		})
		_node.Type = value
	}
	if value, ok := tc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactions.FieldTime,
		})
		_node.Time = value
	}
	if value, ok := tc.mutation.TransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactions.FieldTransactionID,
		})
		_node.TransactionID = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: transactions.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := tc.mutation.Sender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactions.FieldSender,
		})
		_node.Sender = value
	}
	if value, ok := tc.mutation.Recipient(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactions.FieldRecipient,
		})
		_node.Recipient = value
	}
	if value, ok := tc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transactions.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := tc.mutation.Package(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactions.FieldPackage,
		})
		_node.Package = value
	}
	if value, ok := tc.mutation.Module(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactions.FieldModule,
		})
		_node.Module = value
	}
	if value, ok := tc.mutation.Function(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactions.FieldFunction,
		})
		_node.Function = value
	}
	if value, ok := tc.mutation.Gas(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: transactions.FieldGas,
		})
		_node.Gas = value
	}
	return _node, _spec
}

// TransactionsCreateBulk is the builder for creating many Transactions entities in bulk.
type TransactionsCreateBulk struct {
	config
	builders []*TransactionsCreate
}

// Save creates the Transactions entities in the database.
func (tcb *TransactionsCreateBulk) Save(ctx context.Context) ([]*Transactions, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transactions, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionsCreateBulk) SaveX(ctx context.Context) []*Transactions {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionsCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionsCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
