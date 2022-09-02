// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"rei.io/rei/ent/nft"
	"rei.io/rei/ent/predicate"
)

// NFTUpdate is the builder for updating NFT entities.
type NFTUpdate struct {
	config
	hooks    []Hook
	mutation *NFTMutation
}

// Where appends a list predicates to the NFTUpdate builder.
func (nu *NFTUpdate) Where(ps ...predicate.NFT) *NFTUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetObjectID sets the "ObjectID" field.
func (nu *NFTUpdate) SetObjectID(s string) *NFTUpdate {
	nu.mutation.SetObjectID(s)
	return nu
}

// SetType sets the "Type" field.
func (nu *NFTUpdate) SetType(s string) *NFTUpdate {
	nu.mutation.SetType(s)
	return nu
}

// SetMetadata sets the "Metadata" field.
func (nu *NFTUpdate) SetMetadata(m map[string]interface{}) *NFTUpdate {
	nu.mutation.SetMetadata(m)
	return nu
}

// SetSequenceID sets the "SequenceID" field.
func (nu *NFTUpdate) SetSequenceID(u uint64) *NFTUpdate {
	nu.mutation.ResetSequenceID()
	nu.mutation.SetSequenceID(u)
	return nu
}

// AddSequenceID adds u to the "SequenceID" field.
func (nu *NFTUpdate) AddSequenceID(u int64) *NFTUpdate {
	nu.mutation.AddSequenceID(u)
	return nu
}

// Mutation returns the NFTMutation object of the builder.
func (nu *NFTUpdate) Mutation() *NFTMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NFTUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NFTMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NFTUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NFTUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NFTUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NFTUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nft.Table,
			Columns: nft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nft.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.ObjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nft.FieldObjectID,
		})
	}
	if value, ok := nu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nft.FieldType,
		})
	}
	if value, ok := nu.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: nft.FieldMetadata,
		})
	}
	if value, ok := nu.mutation.SequenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: nft.FieldSequenceID,
		})
	}
	if value, ok := nu.mutation.AddedSequenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: nft.FieldSequenceID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// NFTUpdateOne is the builder for updating a single NFT entity.
type NFTUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NFTMutation
}

// SetObjectID sets the "ObjectID" field.
func (nuo *NFTUpdateOne) SetObjectID(s string) *NFTUpdateOne {
	nuo.mutation.SetObjectID(s)
	return nuo
}

// SetType sets the "Type" field.
func (nuo *NFTUpdateOne) SetType(s string) *NFTUpdateOne {
	nuo.mutation.SetType(s)
	return nuo
}

// SetMetadata sets the "Metadata" field.
func (nuo *NFTUpdateOne) SetMetadata(m map[string]interface{}) *NFTUpdateOne {
	nuo.mutation.SetMetadata(m)
	return nuo
}

// SetSequenceID sets the "SequenceID" field.
func (nuo *NFTUpdateOne) SetSequenceID(u uint64) *NFTUpdateOne {
	nuo.mutation.ResetSequenceID()
	nuo.mutation.SetSequenceID(u)
	return nuo
}

// AddSequenceID adds u to the "SequenceID" field.
func (nuo *NFTUpdateOne) AddSequenceID(u int64) *NFTUpdateOne {
	nuo.mutation.AddSequenceID(u)
	return nuo
}

// Mutation returns the NFTMutation object of the builder.
func (nuo *NFTUpdateOne) Mutation() *NFTMutation {
	return nuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NFTUpdateOne) Select(field string, fields ...string) *NFTUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated NFT entity.
func (nuo *NFTUpdateOne) Save(ctx context.Context) (*NFT, error) {
	var (
		err  error
		node *NFT
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NFTMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*NFT)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NFTMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NFTUpdateOne) SaveX(ctx context.Context) *NFT {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NFTUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NFTUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NFTUpdateOne) sqlSave(ctx context.Context) (_node *NFT, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nft.Table,
			Columns: nft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nft.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NFT.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nft.FieldID)
		for _, f := range fields {
			if !nft.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nft.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.ObjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nft.FieldObjectID,
		})
	}
	if value, ok := nuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nft.FieldType,
		})
	}
	if value, ok := nuo.mutation.Metadata(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: nft.FieldMetadata,
		})
	}
	if value, ok := nuo.mutation.SequenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: nft.FieldSequenceID,
		})
	}
	if value, ok := nuo.mutation.AddedSequenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: nft.FieldSequenceID,
		})
	}
	_node = &NFT{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}