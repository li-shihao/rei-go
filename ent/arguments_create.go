// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"rei.io/rei/ent/arguments"
)

// ArgumentsCreate is the builder for creating a Arguments entity.
type ArgumentsCreate struct {
	config
	mutation *ArgumentsMutation
	hooks    []Hook
}

// SetName sets the "Name" field.
func (ac *ArgumentsCreate) SetName(s string) *ArgumentsCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetType sets the "Type" field.
func (ac *ArgumentsCreate) SetType(s string) *ArgumentsCreate {
	ac.mutation.SetType(s)
	return ac
}

// SetTransactionID sets the "TransactionID" field.
func (ac *ArgumentsCreate) SetTransactionID(s string) *ArgumentsCreate {
	ac.mutation.SetTransactionID(s)
	return ac
}

// SetData sets the "Data" field.
func (ac *ArgumentsCreate) SetData(s string) *ArgumentsCreate {
	ac.mutation.SetData(s)
	return ac
}

// Mutation returns the ArgumentsMutation object of the builder.
func (ac *ArgumentsCreate) Mutation() *ArgumentsMutation {
	return ac.mutation
}

// Save creates the Arguments in the database.
func (ac *ArgumentsCreate) Save(ctx context.Context) (*Arguments, error) {
	var (
		err  error
		node *Arguments
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArgumentsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Arguments)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ArgumentsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArgumentsCreate) SaveX(ctx context.Context) *Arguments {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArgumentsCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArgumentsCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArgumentsCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "Name", err: errors.New(`ent: missing required field "Arguments.Name"`)}
	}
	if _, ok := ac.mutation.GetType(); !ok {
		return &ValidationError{Name: "Type", err: errors.New(`ent: missing required field "Arguments.Type"`)}
	}
	if _, ok := ac.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "TransactionID", err: errors.New(`ent: missing required field "Arguments.TransactionID"`)}
	}
	if _, ok := ac.mutation.Data(); !ok {
		return &ValidationError{Name: "Data", err: errors.New(`ent: missing required field "Arguments.Data"`)}
	}
	return nil
}

func (ac *ArgumentsCreate) sqlSave(ctx context.Context) (*Arguments, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *ArgumentsCreate) createSpec() (*Arguments, *sqlgraph.CreateSpec) {
	var (
		_node = &Arguments{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: arguments.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: arguments.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arguments.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arguments.FieldType,
		})
		_node.Type = value
	}
	if value, ok := ac.mutation.TransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arguments.FieldTransactionID,
		})
		_node.TransactionID = value
	}
	if value, ok := ac.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: arguments.FieldData,
		})
		_node.Data = value
	}
	return _node, _spec
}

// ArgumentsCreateBulk is the builder for creating many Arguments entities in bulk.
type ArgumentsCreateBulk struct {
	config
	builders []*ArgumentsCreate
}

// Save creates the Arguments entities in the database.
func (acb *ArgumentsCreateBulk) Save(ctx context.Context) ([]*Arguments, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Arguments, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArgumentsMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArgumentsCreateBulk) SaveX(ctx context.Context) []*Arguments {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArgumentsCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArgumentsCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
