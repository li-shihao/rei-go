// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"rei.io/rei/ent/objects"
)

// ObjectsCreate is the builder for creating a Objects entity.
type ObjectsCreate struct {
	config
	mutation *ObjectsMutation
	hooks    []Hook
}

// SetStatus sets the "Status" field.
func (oc *ObjectsCreate) SetStatus(s string) *ObjectsCreate {
	oc.mutation.SetStatus(s)
	return oc
}

// SetDataType sets the "DataType" field.
func (oc *ObjectsCreate) SetDataType(s string) *ObjectsCreate {
	oc.mutation.SetDataType(s)
	return oc
}

// SetType sets the "Type" field.
func (oc *ObjectsCreate) SetType(s string) *ObjectsCreate {
	oc.mutation.SetType(s)
	return oc
}

// SetHasPublicTransfer sets the "Has_public_transfer" field.
func (oc *ObjectsCreate) SetHasPublicTransfer(b bool) *ObjectsCreate {
	oc.mutation.SetHasPublicTransfer(b)
	return oc
}

// SetFields sets the "Fields" field.
func (oc *ObjectsCreate) SetFields(m map[string]interface{}) *ObjectsCreate {
	oc.mutation.SetFields(m)
	return oc
}

// SetOwner sets the "Owner" field.
func (oc *ObjectsCreate) SetOwner(s string) *ObjectsCreate {
	oc.mutation.SetOwner(s)
	return oc
}

// SetObjectID sets the "ObjectID" field.
func (oc *ObjectsCreate) SetObjectID(s string) *ObjectsCreate {
	oc.mutation.SetObjectID(s)
	return oc
}

// SetSequenceID sets the "SequenceID" field.
func (oc *ObjectsCreate) SetSequenceID(u uint64) *ObjectsCreate {
	oc.mutation.SetSequenceID(u)
	return oc
}

// Mutation returns the ObjectsMutation object of the builder.
func (oc *ObjectsCreate) Mutation() *ObjectsMutation {
	return oc.mutation
}

// Save creates the Objects in the database.
func (oc *ObjectsCreate) Save(ctx context.Context) (*Objects, error) {
	var (
		err  error
		node *Objects
	)
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObjectsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Objects)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ObjectsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *ObjectsCreate) SaveX(ctx context.Context) *Objects {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *ObjectsCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *ObjectsCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *ObjectsCreate) check() error {
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "Status", err: errors.New(`ent: missing required field "Objects.Status"`)}
	}
	if _, ok := oc.mutation.DataType(); !ok {
		return &ValidationError{Name: "DataType", err: errors.New(`ent: missing required field "Objects.DataType"`)}
	}
	if _, ok := oc.mutation.GetType(); !ok {
		return &ValidationError{Name: "Type", err: errors.New(`ent: missing required field "Objects.Type"`)}
	}
	if _, ok := oc.mutation.HasPublicTransfer(); !ok {
		return &ValidationError{Name: "Has_public_transfer", err: errors.New(`ent: missing required field "Objects.Has_public_transfer"`)}
	}
	if _, ok := oc.mutation.GetFields(); !ok {
		return &ValidationError{Name: "Fields", err: errors.New(`ent: missing required field "Objects.Fields"`)}
	}
	if _, ok := oc.mutation.Owner(); !ok {
		return &ValidationError{Name: "Owner", err: errors.New(`ent: missing required field "Objects.Owner"`)}
	}
	if _, ok := oc.mutation.ObjectID(); !ok {
		return &ValidationError{Name: "ObjectID", err: errors.New(`ent: missing required field "Objects.ObjectID"`)}
	}
	if _, ok := oc.mutation.SequenceID(); !ok {
		return &ValidationError{Name: "SequenceID", err: errors.New(`ent: missing required field "Objects.SequenceID"`)}
	}
	return nil
}

func (oc *ObjectsCreate) sqlSave(ctx context.Context) (*Objects, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *ObjectsCreate) createSpec() (*Objects, *sqlgraph.CreateSpec) {
	var (
		_node = &Objects{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: objects.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: objects.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := oc.mutation.DataType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldDataType,
		})
		_node.DataType = value
	}
	if value, ok := oc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldType,
		})
		_node.Type = value
	}
	if value, ok := oc.mutation.HasPublicTransfer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: objects.FieldHasPublicTransfer,
		})
		_node.HasPublicTransfer = value
	}
	if value, ok := oc.mutation.GetFields(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: objects.FieldFields,
		})
		_node.Fields = value
	}
	if value, ok := oc.mutation.Owner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldOwner,
		})
		_node.Owner = value
	}
	if value, ok := oc.mutation.ObjectID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldObjectID,
		})
		_node.ObjectID = value
	}
	if value, ok := oc.mutation.SequenceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: objects.FieldSequenceID,
		})
		_node.SequenceID = value
	}
	return _node, _spec
}

// ObjectsCreateBulk is the builder for creating many Objects entities in bulk.
type ObjectsCreateBulk struct {
	config
	builders []*ObjectsCreate
}

// Save creates the Objects entities in the database.
func (ocb *ObjectsCreateBulk) Save(ctx context.Context) ([]*Objects, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Objects, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ObjectsMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *ObjectsCreateBulk) SaveX(ctx context.Context) []*Objects {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *ObjectsCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *ObjectsCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}