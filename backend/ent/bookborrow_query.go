// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/predicate"
	"github.com/team11/app/ent/servicepoint"
	"github.com/team11/app/ent/user"
)

// BookborrowQuery is the builder for querying Bookborrow entities.
type BookborrowQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Bookborrow
	// eager-loading edges.
	withUSER         *UserQuery
	withBOOK         *BookQuery
	withSERVICEPOINT *ServicePointQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (bq *BookborrowQuery) Where(ps ...predicate.Bookborrow) *BookborrowQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BookborrowQuery) Limit(limit int) *BookborrowQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BookborrowQuery) Offset(offset int) *BookborrowQuery {
	bq.offset = &offset
	return bq
}

// Order adds an order step to the query.
func (bq *BookborrowQuery) Order(o ...OrderFunc) *BookborrowQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryUSER chains the current query on the USER edge.
func (bq *BookborrowQuery) QueryUSER() *UserQuery {
	query := &UserQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookborrow.Table, bookborrow.FieldID, bq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookborrow.USERTable, bookborrow.USERColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBOOK chains the current query on the BOOK edge.
func (bq *BookborrowQuery) QueryBOOK() *BookQuery {
	query := &BookQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookborrow.Table, bookborrow.FieldID, bq.sqlQuery()),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookborrow.BOOKTable, bookborrow.BOOKColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySERVICEPOINT chains the current query on the SERVICEPOINT edge.
func (bq *BookborrowQuery) QuerySERVICEPOINT() *ServicePointQuery {
	query := &ServicePointQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookborrow.Table, bookborrow.FieldID, bq.sqlQuery()),
			sqlgraph.To(servicepoint.Table, servicepoint.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookborrow.SERVICEPOINTTable, bookborrow.SERVICEPOINTColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bookborrow entity in the query. Returns *NotFoundError when no bookborrow was found.
func (bq *BookborrowQuery) First(ctx context.Context) (*Bookborrow, error) {
	bs, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(bs) == 0 {
		return nil, &NotFoundError{bookborrow.Label}
	}
	return bs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BookborrowQuery) FirstX(ctx context.Context) *Bookborrow {
	b, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return b
}

// FirstID returns the first Bookborrow id in the query. Returns *NotFoundError when no id was found.
func (bq *BookborrowQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bookborrow.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (bq *BookborrowQuery) FirstXID(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Bookborrow entity in the query, returns an error if not exactly one entity was returned.
func (bq *BookborrowQuery) Only(ctx context.Context) (*Bookborrow, error) {
	bs, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(bs) {
	case 1:
		return bs[0], nil
	case 0:
		return nil, &NotFoundError{bookborrow.Label}
	default:
		return nil, &NotSingularError{bookborrow.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BookborrowQuery) OnlyX(ctx context.Context) *Bookborrow {
	b, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// OnlyID returns the only Bookborrow id in the query, returns an error if not exactly one id was returned.
func (bq *BookborrowQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = &NotSingularError{bookborrow.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BookborrowQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bookborrows.
func (bq *BookborrowQuery) All(ctx context.Context) ([]*Bookborrow, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BookborrowQuery) AllX(ctx context.Context) []*Bookborrow {
	bs, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return bs
}

// IDs executes the query and returns a list of Bookborrow ids.
func (bq *BookborrowQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(bookborrow.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BookborrowQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BookborrowQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BookborrowQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BookborrowQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BookborrowQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BookborrowQuery) Clone() *BookborrowQuery {
	return &BookborrowQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]OrderFunc{}, bq.order...),
		unique:     append([]string{}, bq.unique...),
		predicates: append([]predicate.Bookborrow{}, bq.predicates...),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

//  WithUSER tells the query-builder to eager-loads the nodes that are connected to
// the "USER" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BookborrowQuery) WithUSER(opts ...func(*UserQuery)) *BookborrowQuery {
	query := &UserQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withUSER = query
	return bq
}

//  WithBOOK tells the query-builder to eager-loads the nodes that are connected to
// the "BOOK" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BookborrowQuery) WithBOOK(opts ...func(*BookQuery)) *BookborrowQuery {
	query := &BookQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withBOOK = query
	return bq
}

//  WithSERVICEPOINT tells the query-builder to eager-loads the nodes that are connected to
// the "SERVICEPOINT" edge. The optional arguments used to configure the query builder of the edge.
func (bq *BookborrowQuery) WithSERVICEPOINT(opts ...func(*ServicePointQuery)) *BookborrowQuery {
	query := &ServicePointQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withSERVICEPOINT = query
	return bq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BORROWDATE time.Time `json:"BORROW_DATE,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bookborrow.Query().
//		GroupBy(bookborrow.FieldBORROWDATE).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BookborrowQuery) GroupBy(field string, fields ...string) *BookborrowGroupBy {
	group := &BookborrowGroupBy{config: bq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		BORROWDATE time.Time `json:"BORROW_DATE,omitempty"`
//	}
//
//	client.Bookborrow.Query().
//		Select(bookborrow.FieldBORROWDATE).
//		Scan(ctx, &v)
//
func (bq *BookborrowQuery) Select(field string, fields ...string) *BookborrowSelect {
	selector := &BookborrowSelect{config: bq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(), nil
	}
	return selector
}

func (bq *BookborrowQuery) prepareQuery(ctx context.Context) error {
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BookborrowQuery) sqlAll(ctx context.Context) ([]*Bookborrow, error) {
	var (
		nodes       = []*Bookborrow{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [3]bool{
			bq.withUSER != nil,
			bq.withBOOK != nil,
			bq.withSERVICEPOINT != nil,
		}
	)
	if bq.withUSER != nil || bq.withBOOK != nil || bq.withSERVICEPOINT != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, bookborrow.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Bookborrow{config: bq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bq.withUSER; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Bookborrow)
		for i := range nodes {
			if fk := nodes[i].USER_ID; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "USER_ID" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.USER = n
			}
		}
	}

	if query := bq.withBOOK; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Bookborrow)
		for i := range nodes {
			if fk := nodes[i].BOOK_ID; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(book.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "BOOK_ID" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BOOK = n
			}
		}
	}

	if query := bq.withSERVICEPOINT; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Bookborrow)
		for i := range nodes {
			if fk := nodes[i].SERVICEPOINT_ID; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(servicepoint.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "SERVICEPOINT_ID" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SERVICEPOINT = n
			}
		}
	}

	return nodes, nil
}

func (bq *BookborrowQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BookborrowQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bq *BookborrowQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bookborrow.Table,
			Columns: bookborrow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookborrow.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BookborrowQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bookborrow.Table)
	selector := builder.Select(t1.Columns(bookborrow.Columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(bookborrow.Columns...)...)
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookborrowGroupBy is the builder for group-by Bookborrow entities.
type BookborrowGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BookborrowGroupBy) Aggregate(fns ...AggregateFunc) *BookborrowGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scan the result into the given value.
func (bgb *BookborrowGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bgb *BookborrowGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookborrowGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bgb *BookborrowGroupBy) StringsX(ctx context.Context) []string {
	v, err := bgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bgb *BookborrowGroupBy) StringX(ctx context.Context) string {
	v, err := bgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookborrowGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bgb *BookborrowGroupBy) IntsX(ctx context.Context) []int {
	v, err := bgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bgb *BookborrowGroupBy) IntX(ctx context.Context) int {
	v, err := bgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookborrowGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bgb *BookborrowGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bgb *BookborrowGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BookborrowGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bgb *BookborrowGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (bgb *BookborrowGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bgb *BookborrowGroupBy) BoolX(ctx context.Context) bool {
	v, err := bgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bgb *BookborrowGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bgb.sqlQuery().Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BookborrowGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql
	columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
	columns = append(columns, bgb.fields...)
	for _, fn := range bgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(bgb.fields...)
}

// BookborrowSelect is the builder for select fields of Bookborrow entities.
type BookborrowSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (bs *BookborrowSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := bs.path(ctx)
	if err != nil {
		return err
	}
	bs.sql = query
	return bs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bs *BookborrowSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookborrowSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bs *BookborrowSelect) StringsX(ctx context.Context) []string {
	v, err := bs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bs *BookborrowSelect) StringX(ctx context.Context) string {
	v, err := bs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookborrowSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bs *BookborrowSelect) IntsX(ctx context.Context) []int {
	v, err := bs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bs *BookborrowSelect) IntX(ctx context.Context) int {
	v, err := bs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookborrowSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bs *BookborrowSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bs *BookborrowSelect) Float64X(ctx context.Context) float64 {
	v, err := bs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BookborrowSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bs *BookborrowSelect) BoolsX(ctx context.Context) []bool {
	v, err := bs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (bs *BookborrowSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookborrow.Label}
	default:
		err = fmt.Errorf("ent: BookborrowSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bs *BookborrowSelect) BoolX(ctx context.Context) bool {
	v, err := bs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bs *BookborrowSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bs.sqlQuery().Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bs *BookborrowSelect) sqlQuery() sql.Querier {
	selector := bs.sql
	selector.Select(selector.Columns(bs.fields...)...)
	return selector
}