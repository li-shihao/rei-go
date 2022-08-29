// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"rei.io/rei/ent/nfts"
	"rei.io/rei/ent/predicate"
)

// NFTsQuery is the builder for querying NFTs entities.
type NFTsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NFTs
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NFTsQuery builder.
func (ntq *NFTsQuery) Where(ps ...predicate.NFTs) *NFTsQuery {
	ntq.predicates = append(ntq.predicates, ps...)
	return ntq
}

// Limit adds a limit step to the query.
func (ntq *NFTsQuery) Limit(limit int) *NFTsQuery {
	ntq.limit = &limit
	return ntq
}

// Offset adds an offset step to the query.
func (ntq *NFTsQuery) Offset(offset int) *NFTsQuery {
	ntq.offset = &offset
	return ntq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ntq *NFTsQuery) Unique(unique bool) *NFTsQuery {
	ntq.unique = &unique
	return ntq
}

// Order adds an order step to the query.
func (ntq *NFTsQuery) Order(o ...OrderFunc) *NFTsQuery {
	ntq.order = append(ntq.order, o...)
	return ntq
}

// First returns the first NFTs entity from the query.
// Returns a *NotFoundError when no NFTs was found.
func (ntq *NFTsQuery) First(ctx context.Context) (*NFTs, error) {
	nodes, err := ntq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nfts.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ntq *NFTsQuery) FirstX(ctx context.Context) *NFTs {
	node, err := ntq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NFTs ID from the query.
// Returns a *NotFoundError when no NFTs ID was found.
func (ntq *NFTsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ntq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nfts.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ntq *NFTsQuery) FirstIDX(ctx context.Context) int {
	id, err := ntq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NFTs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NFTs entity is found.
// Returns a *NotFoundError when no NFTs entities are found.
func (ntq *NFTsQuery) Only(ctx context.Context) (*NFTs, error) {
	nodes, err := ntq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nfts.Label}
	default:
		return nil, &NotSingularError{nfts.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ntq *NFTsQuery) OnlyX(ctx context.Context) *NFTs {
	node, err := ntq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NFTs ID in the query.
// Returns a *NotSingularError when more than one NFTs ID is found.
// Returns a *NotFoundError when no entities are found.
func (ntq *NFTsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ntq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nfts.Label}
	default:
		err = &NotSingularError{nfts.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ntq *NFTsQuery) OnlyIDX(ctx context.Context) int {
	id, err := ntq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NFTsSlice.
func (ntq *NFTsQuery) All(ctx context.Context) ([]*NFTs, error) {
	if err := ntq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ntq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ntq *NFTsQuery) AllX(ctx context.Context) []*NFTs {
	nodes, err := ntq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NFTs IDs.
func (ntq *NFTsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ntq.Select(nfts.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ntq *NFTsQuery) IDsX(ctx context.Context) []int {
	ids, err := ntq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ntq *NFTsQuery) Count(ctx context.Context) (int, error) {
	if err := ntq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ntq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ntq *NFTsQuery) CountX(ctx context.Context) int {
	count, err := ntq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ntq *NFTsQuery) Exist(ctx context.Context) (bool, error) {
	if err := ntq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ntq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ntq *NFTsQuery) ExistX(ctx context.Context) bool {
	exist, err := ntq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NFTsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ntq *NFTsQuery) Clone() *NFTsQuery {
	if ntq == nil {
		return nil
	}
	return &NFTsQuery{
		config:     ntq.config,
		limit:      ntq.limit,
		offset:     ntq.offset,
		order:      append([]OrderFunc{}, ntq.order...),
		predicates: append([]predicate.NFTs{}, ntq.predicates...),
		// clone intermediate query.
		sql:    ntq.sql.Clone(),
		path:   ntq.path,
		unique: ntq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ObjectID string `json:"ObjectID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NFTs.Query().
//		GroupBy(nfts.FieldObjectID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ntq *NFTsQuery) GroupBy(field string, fields ...string) *NFTsGroupBy {
	grbuild := &NFTsGroupBy{config: ntq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ntq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ntq.sqlQuery(ctx), nil
	}
	grbuild.label = nfts.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ObjectID string `json:"ObjectID,omitempty"`
//	}
//
//	client.NFTs.Query().
//		Select(nfts.FieldObjectID).
//		Scan(ctx, &v)
func (ntq *NFTsQuery) Select(fields ...string) *NFTsSelect {
	ntq.fields = append(ntq.fields, fields...)
	selbuild := &NFTsSelect{NFTsQuery: ntq}
	selbuild.label = nfts.Label
	selbuild.flds, selbuild.scan = &ntq.fields, selbuild.Scan
	return selbuild
}

func (ntq *NFTsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ntq.fields {
		if !nfts.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ntq.path != nil {
		prev, err := ntq.path(ctx)
		if err != nil {
			return err
		}
		ntq.sql = prev
	}
	return nil
}

func (ntq *NFTsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NFTs, error) {
	var (
		nodes = []*NFTs{}
		_spec = ntq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*NFTs).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &NFTs{config: ntq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ntq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ntq *NFTsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ntq.querySpec()
	_spec.Node.Columns = ntq.fields
	if len(ntq.fields) > 0 {
		_spec.Unique = ntq.unique != nil && *ntq.unique
	}
	return sqlgraph.CountNodes(ctx, ntq.driver, _spec)
}

func (ntq *NFTsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ntq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ntq *NFTsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   nfts.Table,
			Columns: nfts.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: nfts.FieldID,
			},
		},
		From:   ntq.sql,
		Unique: true,
	}
	if unique := ntq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ntq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nfts.FieldID)
		for i := range fields {
			if fields[i] != nfts.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ntq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ntq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ntq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ntq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ntq *NFTsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ntq.driver.Dialect())
	t1 := builder.Table(nfts.Table)
	columns := ntq.fields
	if len(columns) == 0 {
		columns = nfts.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ntq.sql != nil {
		selector = ntq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ntq.unique != nil && *ntq.unique {
		selector.Distinct()
	}
	for _, p := range ntq.predicates {
		p(selector)
	}
	for _, p := range ntq.order {
		p(selector)
	}
	if offset := ntq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ntq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NFTsGroupBy is the group-by builder for NFTs entities.
type NFTsGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ntgb *NFTsGroupBy) Aggregate(fns ...AggregateFunc) *NFTsGroupBy {
	ntgb.fns = append(ntgb.fns, fns...)
	return ntgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ntgb *NFTsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ntgb.path(ctx)
	if err != nil {
		return err
	}
	ntgb.sql = query
	return ntgb.sqlScan(ctx, v)
}

func (ntgb *NFTsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ntgb.fields {
		if !nfts.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ntgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ntgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ntgb *NFTsGroupBy) sqlQuery() *sql.Selector {
	selector := ntgb.sql.Select()
	aggregation := make([]string, 0, len(ntgb.fns))
	for _, fn := range ntgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ntgb.fields)+len(ntgb.fns))
		for _, f := range ntgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ntgb.fields...)...)
}

// NFTsSelect is the builder for selecting fields of NFTs entities.
type NFTsSelect struct {
	*NFTsQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nts *NFTsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nts.prepareQuery(ctx); err != nil {
		return err
	}
	nts.sql = nts.NFTsQuery.sqlQuery(ctx)
	return nts.sqlScan(ctx, v)
}

func (nts *NFTsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nts.sql.Query()
	if err := nts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
