// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent/ledger"
	"github.com/leeliwei930/walletassignment/ent/predicate"
	"github.com/leeliwei930/walletassignment/ent/user"
	"github.com/leeliwei930/walletassignment/ent/wallet"
)

// WalletQuery is the builder for querying Wallet entities.
type WalletQuery struct {
	config
	ctx         *QueryContext
	order       []wallet.OrderOption
	inters      []Interceptor
	predicates  []predicate.Wallet
	withUser    *UserQuery
	withLedgers *LedgerQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WalletQuery builder.
func (wq *WalletQuery) Where(ps ...predicate.Wallet) *WalletQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WalletQuery) Limit(limit int) *WalletQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WalletQuery) Offset(offset int) *WalletQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WalletQuery) Unique(unique bool) *WalletQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WalletQuery) Order(o ...wallet.OrderOption) *WalletQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryUser chains the current query on the "user" edge.
func (wq *WalletQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(wallet.Table, wallet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, wallet.UserTable, wallet.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLedgers chains the current query on the "ledgers" edge.
func (wq *WalletQuery) QueryLedgers() *LedgerQuery {
	query := (&LedgerClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(wallet.Table, wallet.FieldID, selector),
			sqlgraph.To(ledger.Table, ledger.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wallet.LedgersTable, wallet.LedgersColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Wallet entity from the query.
// Returns a *NotFoundError when no Wallet was found.
func (wq *WalletQuery) First(ctx context.Context) (*Wallet, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wallet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WalletQuery) FirstX(ctx context.Context) *Wallet {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Wallet ID from the query.
// Returns a *NotFoundError when no Wallet ID was found.
func (wq *WalletQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wallet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WalletQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Wallet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Wallet entity is found.
// Returns a *NotFoundError when no Wallet entities are found.
func (wq *WalletQuery) Only(ctx context.Context) (*Wallet, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wallet.Label}
	default:
		return nil, &NotSingularError{wallet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WalletQuery) OnlyX(ctx context.Context) *Wallet {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Wallet ID in the query.
// Returns a *NotSingularError when more than one Wallet ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WalletQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wallet.Label}
	default:
		err = &NotSingularError{wallet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WalletQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Wallets.
func (wq *WalletQuery) All(ctx context.Context) ([]*Wallet, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryAll)
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Wallet, *WalletQuery]()
	return withInterceptors[[]*Wallet](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WalletQuery) AllX(ctx context.Context) []*Wallet {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Wallet IDs.
func (wq *WalletQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryIDs)
	if err = wq.Select(wallet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WalletQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WalletQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryCount)
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WalletQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WalletQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WalletQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, ent.OpQueryExist)
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WalletQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WalletQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WalletQuery) Clone() *WalletQuery {
	if wq == nil {
		return nil
	}
	return &WalletQuery{
		config:      wq.config,
		ctx:         wq.ctx.Clone(),
		order:       append([]wallet.OrderOption{}, wq.order...),
		inters:      append([]Interceptor{}, wq.inters...),
		predicates:  append([]predicate.Wallet{}, wq.predicates...),
		withUser:    wq.withUser.Clone(),
		withLedgers: wq.withLedgers.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WalletQuery) WithUser(opts ...func(*UserQuery)) *WalletQuery {
	query := (&UserClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withUser = query
	return wq
}

// WithLedgers tells the query-builder to eager-load the nodes that are connected to
// the "ledgers" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WalletQuery) WithLedgers(opts ...func(*LedgerQuery)) *WalletQuery {
	query := (&LedgerClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withLedgers = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Wallet.Query().
//		GroupBy(wallet.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WalletQuery) GroupBy(field string, fields ...string) *WalletGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WalletGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = wallet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.Wallet.Query().
//		Select(wallet.FieldUserID).
//		Scan(ctx, &v)
func (wq *WalletQuery) Select(fields ...string) *WalletSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WalletSelect{WalletQuery: wq}
	sbuild.label = wallet.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WalletSelect configured with the given aggregations.
func (wq *WalletQuery) Aggregate(fns ...AggregateFunc) *WalletSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WalletQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !wallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WalletQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Wallet, error) {
	var (
		nodes       = []*Wallet{}
		_spec       = wq.querySpec()
		loadedTypes = [2]bool{
			wq.withUser != nil,
			wq.withLedgers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Wallet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Wallet{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withUser; query != nil {
		if err := wq.loadUser(ctx, query, nodes, nil,
			func(n *Wallet, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withLedgers; query != nil {
		if err := wq.loadLedgers(ctx, query, nodes,
			func(n *Wallet) { n.Edges.Ledgers = []*Ledger{} },
			func(n *Wallet, e *Ledger) { n.Edges.Ledgers = append(n.Edges.Ledgers, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WalletQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Wallet, init func(*Wallet), assign func(*Wallet, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Wallet)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WalletQuery) loadLedgers(ctx context.Context, query *LedgerQuery, nodes []*Wallet, init func(*Wallet), assign func(*Wallet, *Ledger)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Wallet)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(ledger.FieldWalletID)
	}
	query.Where(predicate.Ledger(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(wallet.LedgersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.WalletID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "wallet_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wq *WalletQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WalletQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wallet.Table, wallet.Columns, sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeUUID))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wallet.FieldID)
		for i := range fields {
			if fields[i] != wallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if wq.withUser != nil {
			_spec.Node.AddColumnOnce(wallet.FieldUserID)
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WalletQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(wallet.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = wallet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range wq.modifiers {
		m(selector)
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (wq *WalletQuery) ForUpdate(opts ...sql.LockOption) *WalletQuery {
	if wq.driver.Dialect() == dialect.Postgres {
		wq.Unique(false)
	}
	wq.modifiers = append(wq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return wq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (wq *WalletQuery) ForShare(opts ...sql.LockOption) *WalletQuery {
	if wq.driver.Dialect() == dialect.Postgres {
		wq.Unique(false)
	}
	wq.modifiers = append(wq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return wq
}

// WalletGroupBy is the group-by builder for Wallet entities.
type WalletGroupBy struct {
	selector
	build *WalletQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WalletGroupBy) Aggregate(fns ...AggregateFunc) *WalletGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WalletGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, ent.OpQueryGroupBy)
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WalletQuery, *WalletGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WalletGroupBy) sqlScan(ctx context.Context, root *WalletQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WalletSelect is the builder for selecting fields of Wallet entities.
type WalletSelect struct {
	*WalletQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WalletSelect) Aggregate(fns ...AggregateFunc) *WalletSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WalletSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, ent.OpQuerySelect)
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WalletQuery, *WalletSelect](ctx, ws.WalletQuery, ws, ws.inters, v)
}

func (ws *WalletSelect) sqlScan(ctx context.Context, root *WalletQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
