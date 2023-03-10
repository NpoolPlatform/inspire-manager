// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/goodorderpercent"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodOrderPercentQuery is the builder for querying GoodOrderPercent entities.
type GoodOrderPercentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodOrderPercent
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodOrderPercentQuery builder.
func (gopq *GoodOrderPercentQuery) Where(ps ...predicate.GoodOrderPercent) *GoodOrderPercentQuery {
	gopq.predicates = append(gopq.predicates, ps...)
	return gopq
}

// Limit adds a limit step to the query.
func (gopq *GoodOrderPercentQuery) Limit(limit int) *GoodOrderPercentQuery {
	gopq.limit = &limit
	return gopq
}

// Offset adds an offset step to the query.
func (gopq *GoodOrderPercentQuery) Offset(offset int) *GoodOrderPercentQuery {
	gopq.offset = &offset
	return gopq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gopq *GoodOrderPercentQuery) Unique(unique bool) *GoodOrderPercentQuery {
	gopq.unique = &unique
	return gopq
}

// Order adds an order step to the query.
func (gopq *GoodOrderPercentQuery) Order(o ...OrderFunc) *GoodOrderPercentQuery {
	gopq.order = append(gopq.order, o...)
	return gopq
}

// First returns the first GoodOrderPercent entity from the query.
// Returns a *NotFoundError when no GoodOrderPercent was found.
func (gopq *GoodOrderPercentQuery) First(ctx context.Context) (*GoodOrderPercent, error) {
	nodes, err := gopq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodorderpercent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) FirstX(ctx context.Context) *GoodOrderPercent {
	node, err := gopq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodOrderPercent ID from the query.
// Returns a *NotFoundError when no GoodOrderPercent ID was found.
func (gopq *GoodOrderPercentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gopq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodorderpercent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gopq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodOrderPercent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodOrderPercent entity is found.
// Returns a *NotFoundError when no GoodOrderPercent entities are found.
func (gopq *GoodOrderPercentQuery) Only(ctx context.Context) (*GoodOrderPercent, error) {
	nodes, err := gopq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodorderpercent.Label}
	default:
		return nil, &NotSingularError{goodorderpercent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) OnlyX(ctx context.Context) *GoodOrderPercent {
	node, err := gopq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodOrderPercent ID in the query.
// Returns a *NotSingularError when more than one GoodOrderPercent ID is found.
// Returns a *NotFoundError when no entities are found.
func (gopq *GoodOrderPercentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gopq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodorderpercent.Label}
	default:
		err = &NotSingularError{goodorderpercent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gopq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodOrderPercents.
func (gopq *GoodOrderPercentQuery) All(ctx context.Context) ([]*GoodOrderPercent, error) {
	if err := gopq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gopq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) AllX(ctx context.Context) []*GoodOrderPercent {
	nodes, err := gopq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodOrderPercent IDs.
func (gopq *GoodOrderPercentQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gopq.Select(goodorderpercent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gopq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gopq *GoodOrderPercentQuery) Count(ctx context.Context) (int, error) {
	if err := gopq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gopq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) CountX(ctx context.Context) int {
	count, err := gopq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gopq *GoodOrderPercentQuery) Exist(ctx context.Context) (bool, error) {
	if err := gopq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gopq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gopq *GoodOrderPercentQuery) ExistX(ctx context.Context) bool {
	exist, err := gopq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodOrderPercentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gopq *GoodOrderPercentQuery) Clone() *GoodOrderPercentQuery {
	if gopq == nil {
		return nil
	}
	return &GoodOrderPercentQuery{
		config:     gopq.config,
		limit:      gopq.limit,
		offset:     gopq.offset,
		order:      append([]OrderFunc{}, gopq.order...),
		predicates: append([]predicate.GoodOrderPercent{}, gopq.predicates...),
		// clone intermediate query.
		sql:    gopq.sql.Clone(),
		path:   gopq.path,
		unique: gopq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodOrderPercent.Query().
//		GroupBy(goodorderpercent.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gopq *GoodOrderPercentQuery) GroupBy(field string, fields ...string) *GoodOrderPercentGroupBy {
	grbuild := &GoodOrderPercentGroupBy{config: gopq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gopq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gopq.sqlQuery(ctx), nil
	}
	grbuild.label = goodorderpercent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.GoodOrderPercent.Query().
//		Select(goodorderpercent.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (gopq *GoodOrderPercentQuery) Select(fields ...string) *GoodOrderPercentSelect {
	gopq.fields = append(gopq.fields, fields...)
	selbuild := &GoodOrderPercentSelect{GoodOrderPercentQuery: gopq}
	selbuild.label = goodorderpercent.Label
	selbuild.flds, selbuild.scan = &gopq.fields, selbuild.Scan
	return selbuild
}

func (gopq *GoodOrderPercentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gopq.fields {
		if !goodorderpercent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gopq.path != nil {
		prev, err := gopq.path(ctx)
		if err != nil {
			return err
		}
		gopq.sql = prev
	}
	if goodorderpercent.Policy == nil {
		return errors.New("ent: uninitialized goodorderpercent.Policy (forgotten import ent/runtime?)")
	}
	if err := goodorderpercent.Policy.EvalQuery(ctx, gopq); err != nil {
		return err
	}
	return nil
}

func (gopq *GoodOrderPercentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GoodOrderPercent, error) {
	var (
		nodes = []*GoodOrderPercent{}
		_spec = gopq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*GoodOrderPercent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &GoodOrderPercent{config: gopq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gopq.modifiers) > 0 {
		_spec.Modifiers = gopq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gopq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gopq *GoodOrderPercentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gopq.querySpec()
	if len(gopq.modifiers) > 0 {
		_spec.Modifiers = gopq.modifiers
	}
	_spec.Node.Columns = gopq.fields
	if len(gopq.fields) > 0 {
		_spec.Unique = gopq.unique != nil && *gopq.unique
	}
	return sqlgraph.CountNodes(ctx, gopq.driver, _spec)
}

func (gopq *GoodOrderPercentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gopq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gopq *GoodOrderPercentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodorderpercent.Table,
			Columns: goodorderpercent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodorderpercent.FieldID,
			},
		},
		From:   gopq.sql,
		Unique: true,
	}
	if unique := gopq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gopq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodorderpercent.FieldID)
		for i := range fields {
			if fields[i] != goodorderpercent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gopq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gopq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gopq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gopq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gopq *GoodOrderPercentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gopq.driver.Dialect())
	t1 := builder.Table(goodorderpercent.Table)
	columns := gopq.fields
	if len(columns) == 0 {
		columns = goodorderpercent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gopq.sql != nil {
		selector = gopq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gopq.unique != nil && *gopq.unique {
		selector.Distinct()
	}
	for _, m := range gopq.modifiers {
		m(selector)
	}
	for _, p := range gopq.predicates {
		p(selector)
	}
	for _, p := range gopq.order {
		p(selector)
	}
	if offset := gopq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gopq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gopq *GoodOrderPercentQuery) ForUpdate(opts ...sql.LockOption) *GoodOrderPercentQuery {
	if gopq.driver.Dialect() == dialect.Postgres {
		gopq.Unique(false)
	}
	gopq.modifiers = append(gopq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gopq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gopq *GoodOrderPercentQuery) ForShare(opts ...sql.LockOption) *GoodOrderPercentQuery {
	if gopq.driver.Dialect() == dialect.Postgres {
		gopq.Unique(false)
	}
	gopq.modifiers = append(gopq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gopq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gopq *GoodOrderPercentQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodOrderPercentSelect {
	gopq.modifiers = append(gopq.modifiers, modifiers...)
	return gopq.Select()
}

// GoodOrderPercentGroupBy is the group-by builder for GoodOrderPercent entities.
type GoodOrderPercentGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gopgb *GoodOrderPercentGroupBy) Aggregate(fns ...AggregateFunc) *GoodOrderPercentGroupBy {
	gopgb.fns = append(gopgb.fns, fns...)
	return gopgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gopgb *GoodOrderPercentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gopgb.path(ctx)
	if err != nil {
		return err
	}
	gopgb.sql = query
	return gopgb.sqlScan(ctx, v)
}

func (gopgb *GoodOrderPercentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gopgb.fields {
		if !goodorderpercent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gopgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gopgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gopgb *GoodOrderPercentGroupBy) sqlQuery() *sql.Selector {
	selector := gopgb.sql.Select()
	aggregation := make([]string, 0, len(gopgb.fns))
	for _, fn := range gopgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gopgb.fields)+len(gopgb.fns))
		for _, f := range gopgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gopgb.fields...)...)
}

// GoodOrderPercentSelect is the builder for selecting fields of GoodOrderPercent entities.
type GoodOrderPercentSelect struct {
	*GoodOrderPercentQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gops *GoodOrderPercentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gops.prepareQuery(ctx); err != nil {
		return err
	}
	gops.sql = gops.GoodOrderPercentQuery.sqlQuery(ctx)
	return gops.sqlScan(ctx, v)
}

func (gops *GoodOrderPercentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gops.sql.Query()
	if err := gops.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gops *GoodOrderPercentSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodOrderPercentSelect {
	gops.modifiers = append(gops.modifiers, modifiers...)
	return gops
}
