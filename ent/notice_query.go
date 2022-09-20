// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cateiru/cateir.com/ent/notice"
	"github.com/cateiru/cateir.com/ent/predicate"
)

// NoticeQuery is the builder for querying Notice entities.
type NoticeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Notice
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NoticeQuery builder.
func (nq *NoticeQuery) Where(ps ...predicate.Notice) *NoticeQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NoticeQuery) Limit(limit int) *NoticeQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NoticeQuery) Offset(offset int) *NoticeQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NoticeQuery) Unique(unique bool) *NoticeQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NoticeQuery) Order(o ...OrderFunc) *NoticeQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// First returns the first Notice entity from the query.
// Returns a *NotFoundError when no Notice was found.
func (nq *NoticeQuery) First(ctx context.Context) (*Notice, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NoticeQuery) FirstX(ctx context.Context) *Notice {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Notice ID from the query.
// Returns a *NotFoundError when no Notice ID was found.
func (nq *NoticeQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NoticeQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Notice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Notice entity is found.
// Returns a *NotFoundError when no Notice entities are found.
func (nq *NoticeQuery) Only(ctx context.Context) (*Notice, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notice.Label}
	default:
		return nil, &NotSingularError{notice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NoticeQuery) OnlyX(ctx context.Context) *Notice {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Notice ID in the query.
// Returns a *NotSingularError when more than one Notice ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NoticeQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notice.Label}
	default:
		err = &NotSingularError{notice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NoticeQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Notices.
func (nq *NoticeQuery) All(ctx context.Context) ([]*Notice, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NoticeQuery) AllX(ctx context.Context) []*Notice {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Notice IDs.
func (nq *NoticeQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := nq.Select(notice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NoticeQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NoticeQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NoticeQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NoticeQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NoticeQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NoticeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NoticeQuery) Clone() *NoticeQuery {
	if nq == nil {
		return nil
	}
	return &NoticeQuery{
		config:     nq.config,
		limit:      nq.limit,
		offset:     nq.offset,
		order:      append([]OrderFunc{}, nq.order...),
		predicates: append([]predicate.Notice{}, nq.predicates...),
		// clone intermediate query.
		sql:    nq.sql.Clone(),
		path:   nq.path,
		unique: nq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DiscordWebhook string `json:"discord_webhook,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Notice.Query().
//		GroupBy(notice.FieldDiscordWebhook).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NoticeQuery) GroupBy(field string, fields ...string) *NoticeGroupBy {
	grbuild := &NoticeGroupBy{config: nq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(ctx), nil
	}
	grbuild.label = notice.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DiscordWebhook string `json:"discord_webhook,omitempty"`
//	}
//
//	client.Notice.Query().
//		Select(notice.FieldDiscordWebhook).
//		Scan(ctx, &v)
func (nq *NoticeQuery) Select(fields ...string) *NoticeSelect {
	nq.fields = append(nq.fields, fields...)
	selbuild := &NoticeSelect{NoticeQuery: nq}
	selbuild.label = notice.Label
	selbuild.flds, selbuild.scan = &nq.fields, selbuild.Scan
	return selbuild
}

func (nq *NoticeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nq.fields {
		if !notice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NoticeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Notice, error) {
	var (
		nodes = []*Notice{}
		_spec = nq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Notice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Notice{config: nq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (nq *NoticeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.fields
	if len(nq.fields) > 0 {
		_spec.Unique = nq.unique != nil && *nq.unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NoticeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nq *NoticeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   notice.Table,
			Columns: notice.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: notice.FieldID,
			},
		},
		From:   nq.sql,
		Unique: true,
	}
	if unique := nq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notice.FieldID)
		for i := range fields {
			if fields[i] != notice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NoticeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(notice.Table)
	columns := nq.fields
	if len(columns) == 0 {
		columns = notice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.unique != nil && *nq.unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NoticeGroupBy is the group-by builder for Notice entities.
type NoticeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NoticeGroupBy) Aggregate(fns ...AggregateFunc) *NoticeGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scans the result into the given value.
func (ngb *NoticeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.sql = query
	return ngb.sqlScan(ctx, v)
}

func (ngb *NoticeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ngb.fields {
		if !notice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NoticeGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql.Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
		for _, f := range ngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ngb.fields...)...)
}

// NoticeSelect is the builder for selecting fields of Notice entities.
type NoticeSelect struct {
	*NoticeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NoticeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.sql = ns.NoticeQuery.sqlQuery(ctx)
	return ns.sqlScan(ctx, v)
}

func (ns *NoticeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ns.sql.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
