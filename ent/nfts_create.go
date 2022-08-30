// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"rei.io/rei/ent/nfts"
)

// NFTsCreate is the builder for creating a NFTs entity.
type NFTsCreate struct {
	config
	mutation *NFTsMutation
	hooks    []Hook
}

// SetObjectID sets the "ObjectID" field.
func (ntc *NFTsCreate) SetObjectID(s string) *NFTsCreate {
	ntc.mutation.SetObjectID(s)
	return ntc
}

// SetType sets the "Type" field.
func (ntc *NFTsCreate) SetType(s string) *NFTsCreate {
	ntc.mutation.SetType(s)
	return ntc
}

// SetMetadata sets the "Metadata" field.
func (ntc *NFTsCreate) SetMetadata(m map[string]interface{}) *NFTsCreate {
	ntc.mutation.SetMetadata(m)
	return ntc
}

// SetSequenceID sets the "SequenceID" field.
func (ntc *NFTsCreate) SetSequenceID(u uint64) *NFTsCreate {
	ntc.mutation.SetSequenceID(u)
	return ntc
}

// Mutation returns the NFTsMutation object of the builder.
func (ntc *NFTsCreate) Mutation() *NFTsMutation {
	return ntc.mutation
}

// Save creates the NFTs in the database.
func (ntc *NFTsCreate) Save(ctx context.Context) (*NFTs, error) {
	var (
		err  error
		node *NFTs
	)
	if len(ntc.hooks) == 0 {
		if err = ntc.check(); err != nil {
			return nil, err
		}
		node, err = ntc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NFTsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ntc.check(); err != nil {
				return nil, err
			}
			ntc.mutation = mutation
			if node, err = ntc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ntc.hooks) - 1; i >= 0; i-- {
			if ntc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ntc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ntc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*NFTs)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NFTsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ntc *NFTsCreate) SaveX(ctx context.Context) *NFTs {
	v, err := ntc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntc *NFTsCreate) Exec(ctx context.Context) error {
	_, err := ntc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntc *NFTsCreate) ExecX(ctx context.Context) {
	if err := ntc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntc *NFTsCreate) check() error {
	if _, ok := ntc.mutation.ObjectID(); !ok {
		return &ValidationError{Name: "ObjectID", err: errors.New(`ent: missing required field "NFTs.ObjectID"`)}
	}
	if _, ok := ntc.mutation.GetType(); !ok {
		return &ValidationError{Name: "Type", err: errors.New(`ent: missing required field "NFTs.Type"`)}
	}
	if _, ok := ntc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "Metadata", err: errors.New(`ent: missing required field "NFTs.Metadata"`)}
	}
	if _, ok := ntc.mutation.SequenceID(); !ok {
		return &ValidationError{Name: "SequenceID", err: errors.New(`ent: missing required field "NFTs.SequenceID"`)}
	}
	return nil
}

func (ntc *NFTsCreate) sqlSave(ctx context.Context) (*NFTs, error) {
	_node, _spec := ntc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ntc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ntc *NFTsCreate) createSpec() (*NFTs, *sqlgraph.CreateSpec) {
	var (
		_node = &NFTs{config: ntc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: nfts.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nfts.FieldID,
			},
		}
	)
	if value, ok := ntc.mutation.ObjectID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nfts.FieldObjectID,
		})
		_node.ObjectID = value
	}
	if value, ok := ntc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nfts.FieldType,
		})
		_node.Type = value
	}
	if value, ok := ntc.mutation.Metadata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: nfts.FieldMetadata,
		})
		_node.Metadata = value
	}
	if value, ok := ntc.mutation.SequenceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: nfts.FieldSequenceID,
		})
		_node.SequenceID = value
	}
	return _node, _spec
}

// NFTsCreateBulk is the builder for creating many NFTs entities in bulk.
type NFTsCreateBulk struct {
	config
	builders []*NFTsCreate
}

// Save creates the NFTs entities in the database.
func (ntcb *NFTsCreateBulk) Save(ctx context.Context) ([]*NFTs, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ntcb.builders))
	nodes := make([]*NFTs, len(ntcb.builders))
	mutators := make([]Mutator, len(ntcb.builders))
	for i := range ntcb.builders {
		func(i int, root context.Context) {
			builder := ntcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NFTsMutation)
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
					_, err = mutators[i+1].Mutate(root, ntcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ntcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ntcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ntcb *NFTsCreateBulk) SaveX(ctx context.Context) []*NFTs {
	v, err := ntcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntcb *NFTsCreateBulk) Exec(ctx context.Context) error {
	_, err := ntcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntcb *NFTsCreateBulk) ExecX(ctx context.Context) {
	if err := ntcb.Exec(ctx); err != nil {
		panic(err)
	}
}