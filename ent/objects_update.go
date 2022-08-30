// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"rei.io/rei/ent/objects"
	"rei.io/rei/ent/predicate"
)

// ObjectsUpdate is the builder for updating Objects entities.
type ObjectsUpdate struct {
	config
	hooks    []Hook
	mutation *ObjectsMutation
}

// Where appends a list predicates to the ObjectsUpdate builder.
func (ou *ObjectsUpdate) Where(ps ...predicate.Objects) *ObjectsUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetStatus sets the "Status" field.
func (ou *ObjectsUpdate) SetStatus(s string) *ObjectsUpdate {
	ou.mutation.SetStatus(s)
	return ou
}

// SetDataType sets the "DataType" field.
func (ou *ObjectsUpdate) SetDataType(s string) *ObjectsUpdate {
	ou.mutation.SetDataType(s)
	return ou
}

// SetType sets the "Type" field.
func (ou *ObjectsUpdate) SetType(s string) *ObjectsUpdate {
	ou.mutation.SetType(s)
	return ou
}

// SetHasPublicTransfer sets the "Has_public_transfer" field.
func (ou *ObjectsUpdate) SetHasPublicTransfer(b bool) *ObjectsUpdate {
	ou.mutation.SetHasPublicTransfer(b)
	return ou
}

// SetFields sets the "Fields" field.
func (ou *ObjectsUpdate) SetFields(m map[string]interface{}) *ObjectsUpdate {
	ou.mutation.SetFields(m)
	return ou
}

// SetOwner sets the "Owner" field.
func (ou *ObjectsUpdate) SetOwner(s string) *ObjectsUpdate {
	ou.mutation.SetOwner(s)
	return ou
}

// SetObjectID sets the "ObjectID" field.
func (ou *ObjectsUpdate) SetObjectID(s string) *ObjectsUpdate {
	ou.mutation.SetObjectID(s)
	return ou
}

// SetSequenceID sets the "SequenceID" field.
func (ou *ObjectsUpdate) SetSequenceID(u uint64) *ObjectsUpdate {
	ou.mutation.ResetSequenceID()
	ou.mutation.SetSequenceID(u)
	return ou
}

// AddSequenceID adds u to the "SequenceID" field.
func (ou *ObjectsUpdate) AddSequenceID(u int64) *ObjectsUpdate {
	ou.mutation.AddSequenceID(u)
	return ou
}

// Mutation returns the ObjectsMutation object of the builder.
func (ou *ObjectsUpdate) Mutation() *ObjectsMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *ObjectsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObjectsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *ObjectsUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *ObjectsUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *ObjectsUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *ObjectsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   objects.Table,
			Columns: objects.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: objects.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldStatus,
		})
	}
	if value, ok := ou.mutation.DataType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldDataType,
		})
	}
	if value, ok := ou.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldType,
		})
	}
	if value, ok := ou.mutation.HasPublicTransfer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: objects.FieldHasPublicTransfer,
		})
	}
	if value, ok := ou.mutation.GetFields(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: objects.FieldFields,
		})
	}
	if value, ok := ou.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldOwner,
		})
	}
	if value, ok := ou.mutation.ObjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldObjectID,
		})
	}
	if value, ok := ou.mutation.SequenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: objects.FieldSequenceID,
		})
	}
	if value, ok := ou.mutation.AddedSequenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: objects.FieldSequenceID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{objects.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ObjectsUpdateOne is the builder for updating a single Objects entity.
type ObjectsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ObjectsMutation
}

// SetStatus sets the "Status" field.
func (ouo *ObjectsUpdateOne) SetStatus(s string) *ObjectsUpdateOne {
	ouo.mutation.SetStatus(s)
	return ouo
}

// SetDataType sets the "DataType" field.
func (ouo *ObjectsUpdateOne) SetDataType(s string) *ObjectsUpdateOne {
	ouo.mutation.SetDataType(s)
	return ouo
}

// SetType sets the "Type" field.
func (ouo *ObjectsUpdateOne) SetType(s string) *ObjectsUpdateOne {
	ouo.mutation.SetType(s)
	return ouo
}

// SetHasPublicTransfer sets the "Has_public_transfer" field.
func (ouo *ObjectsUpdateOne) SetHasPublicTransfer(b bool) *ObjectsUpdateOne {
	ouo.mutation.SetHasPublicTransfer(b)
	return ouo
}

// SetFields sets the "Fields" field.
func (ouo *ObjectsUpdateOne) SetFields(m map[string]interface{}) *ObjectsUpdateOne {
	ouo.mutation.SetFields(m)
	return ouo
}

// SetOwner sets the "Owner" field.
func (ouo *ObjectsUpdateOne) SetOwner(s string) *ObjectsUpdateOne {
	ouo.mutation.SetOwner(s)
	return ouo
}

// SetObjectID sets the "ObjectID" field.
func (ouo *ObjectsUpdateOne) SetObjectID(s string) *ObjectsUpdateOne {
	ouo.mutation.SetObjectID(s)
	return ouo
}

// SetSequenceID sets the "SequenceID" field.
func (ouo *ObjectsUpdateOne) SetSequenceID(u uint64) *ObjectsUpdateOne {
	ouo.mutation.ResetSequenceID()
	ouo.mutation.SetSequenceID(u)
	return ouo
}

// AddSequenceID adds u to the "SequenceID" field.
func (ouo *ObjectsUpdateOne) AddSequenceID(u int64) *ObjectsUpdateOne {
	ouo.mutation.AddSequenceID(u)
	return ouo
}

// Mutation returns the ObjectsMutation object of the builder.
func (ouo *ObjectsUpdateOne) Mutation() *ObjectsMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *ObjectsUpdateOne) Select(field string, fields ...string) *ObjectsUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Objects entity.
func (ouo *ObjectsUpdateOne) Save(ctx context.Context) (*Objects, error) {
	var (
		err  error
		node *Objects
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObjectsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (ouo *ObjectsUpdateOne) SaveX(ctx context.Context) *Objects {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *ObjectsUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *ObjectsUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *ObjectsUpdateOne) sqlSave(ctx context.Context) (_node *Objects, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   objects.Table,
			Columns: objects.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: objects.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Objects.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, objects.FieldID)
		for _, f := range fields {
			if !objects.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != objects.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldStatus,
		})
	}
	if value, ok := ouo.mutation.DataType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldDataType,
		})
	}
	if value, ok := ouo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldType,
		})
	}
	if value, ok := ouo.mutation.HasPublicTransfer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: objects.FieldHasPublicTransfer,
		})
	}
	if value, ok := ouo.mutation.GetFields(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: objects.FieldFields,
		})
	}
	if value, ok := ouo.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldOwner,
		})
	}
	if value, ok := ouo.mutation.ObjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: objects.FieldObjectID,
		})
	}
	if value, ok := ouo.mutation.SequenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: objects.FieldSequenceID,
		})
	}
	if value, ok := ouo.mutation.AddedSequenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: objects.FieldSequenceID,
		})
	}
	_node = &Objects{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{objects.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}