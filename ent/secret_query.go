// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/TinyMonitor/ent/predicate"
	"github.com/lbrictson/TinyMonitor/ent/secret"
)

// SecretQuery is the builder for querying Secret entities.
type SecretQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Secret
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SecretQuery builder.
func (sq *SecretQuery) Where(ps ...predicate.Secret) *SecretQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SecretQuery) Limit(limit int) *SecretQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SecretQuery) Offset(offset int) *SecretQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SecretQuery) Unique(unique bool) *SecretQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SecretQuery) Order(o ...OrderFunc) *SecretQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// First returns the first Secret entity from the query.
// Returns a *NotFoundError when no Secret was found.
func (sq *SecretQuery) First(ctx context.Context) (*Secret, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{secret.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SecretQuery) FirstX(ctx context.Context) *Secret {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Secret ID from the query.
// Returns a *NotFoundError when no Secret ID was found.
func (sq *SecretQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{secret.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SecretQuery) FirstIDX(ctx context.Context) string {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Secret entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Secret entity is found.
// Returns a *NotFoundError when no Secret entities are found.
func (sq *SecretQuery) Only(ctx context.Context) (*Secret, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{secret.Label}
	default:
		return nil, &NotSingularError{secret.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SecretQuery) OnlyX(ctx context.Context) *Secret {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Secret ID in the query.
// Returns a *NotSingularError when more than one Secret ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SecretQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{secret.Label}
	default:
		err = &NotSingularError{secret.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SecretQuery) OnlyIDX(ctx context.Context) string {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Secrets.
func (sq *SecretQuery) All(ctx context.Context) ([]*Secret, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SecretQuery) AllX(ctx context.Context) []*Secret {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Secret IDs.
func (sq *SecretQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := sq.Select(secret.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SecretQuery) IDsX(ctx context.Context) []string {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SecretQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SecretQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SecretQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SecretQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SecretQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SecretQuery) Clone() *SecretQuery {
	if sq == nil {
		return nil
	}
	return &SecretQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		predicates: append([]predicate.Secret{}, sq.predicates...),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Secret.Query().
//		GroupBy(secret.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SecretQuery) GroupBy(field string, fields ...string) *SecretGroupBy {
	grbuild := &SecretGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = secret.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Secret.Query().
//		Select(secret.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *SecretQuery) Select(fields ...string) *SecretSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SecretSelect{SecretQuery: sq}
	selbuild.label = secret.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SecretSelect configured with the given aggregations.
func (sq *SecretQuery) Aggregate(fns ...AggregateFunc) *SecretSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SecretQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !secret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SecretQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Secret, error) {
	var (
		nodes = []*Secret{}
		_spec = sq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Secret).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Secret{config: sq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sq *SecretQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SecretQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *SecretQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   secret.Table,
			Columns: secret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: secret.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, secret.FieldID)
		for i := range fields {
			if fields[i] != secret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SecretQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(secret.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = secret.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SecretGroupBy is the group-by builder for Secret entities.
type SecretGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SecretGroupBy) Aggregate(fns ...AggregateFunc) *SecretGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SecretGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SecretGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !secret.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SecretGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SecretSelect is the builder for selecting fields of Secret entities.
type SecretSelect struct {
	*SecretQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SecretSelect) Aggregate(fns ...AggregateFunc) *SecretSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SecretSelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SecretQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SecretSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(ss.sql))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ss.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ss.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}