package select_

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/expr"
	"strings"
)

/**
 *
 *
 *
 *
 */
type ISQLSelectQuery interface {
	expr.ISQLExpr

	OrderByClause() *SQLOrderByClause
	SetOrderByClause(x *SQLOrderByClause)

	LimitClause() ISQLLimitClause
	SetLimitClause(x ISQLLimitClause)

	LockClause() ISQLLockClause
	SetLockClause(x ISQLLockClause)
}
type AbstractSQLSelectQuery struct {
	*expr.SQLExpr

	orderByClause *SQLOrderByClause
	limitClause   ISQLLimitClause
	lockClause    ISQLLockClause
}

func NewAbstractSelectQuery() *AbstractSQLSelectQuery {
	x := new(AbstractSQLSelectQuery)
	x.SQLExpr = expr.NewExpr()
	return x
}

func (x *AbstractSQLSelectQuery) OrderByClause() *SQLOrderByClause {
	return x.orderByClause
}

func (x *AbstractSQLSelectQuery) SetOrderByClause(orderByClause *SQLOrderByClause) {
	if orderByClause == nil {
		return
	}
	orderByClause.SetParent(orderByClause)
	x.orderByClause = orderByClause
}

func (x *AbstractSQLSelectQuery) LimitClause() ISQLLimitClause {
	return x.limitClause
}

func (x *AbstractSQLSelectQuery) SetLimitClause(limitClause ISQLLimitClause) {
	if limitClause == nil {
		return
	}
	limitClause.SetParent(x)
	x.limitClause = limitClause
}

func (x *AbstractSQLSelectQuery) LockClause() ISQLLockClause {
	return x.lockClause
}

func (x *AbstractSQLSelectQuery) SetLockClause(lockClause ISQLLockClause) {
	if ast.IsNil(lockClause) {
		return
	}
	lockClause.SetParent(x)
	x.lockClause = lockClause
}

type SQLSelectElement struct {
	*expr.SQLExpr

	expr  expr.ISQLExpr
	As    bool
	alias expr.ISQLExpr
}

func NewSelectElement() *SQLSelectElement {
	x := new(SQLSelectElement)
	x.SQLExpr = expr.NewExpr()
	return x
}

func NewSelectElementWithExpr(expr expr.ISQLExpr) *SQLSelectElement {
	x := NewSelectElement()
	x.SetExpr(expr)
	return x
}

func NewSelectElementWithAlias(expr expr.ISQLExpr, as bool, alias expr.ISQLExpr) *SQLSelectElement {
	x := NewSelectElement()
	x.SetExpr(expr)
	x.As = as
	x.SetAlias(alias)
	return x
}

func (x *SQLSelectElement) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLSelectElement) SetExpr(expr expr.ISQLExpr) {
	if ast.IsNil(expr) {
		return
	}
	expr.SetParent(x)
	x.expr = expr
}
func (x *SQLSelectElement) Alias() expr.ISQLExpr {
	return x.alias
}
func (x *SQLSelectElement) SetAlias(alias expr.ISQLExpr) {
	if ast.IsNil(alias) {
		return
	}
	alias.SetParent(x)
	x.alias = alias
}

type SQLSelectTargetElement struct {
	*expr.SQLExpr
	expr expr.ISQLExpr
}

func NewSelectTargetElement() *SQLSelectTargetElement {
	x := new(SQLSelectTargetElement)
	x.SQLExpr = expr.NewExpr()
	return x
}
func NewSelectTargetElementWithExpr(expr expr.ISQLExpr) *SQLSelectTargetElement {
	x := NewSelectTargetElement()
	x.SetExpr(expr)
	return x
}
func (x *SQLSelectTargetElement) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLSelectTargetElement) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

// select ... from ... where ..
type SQLSelectQuery struct {
	*AbstractSQLSelectQuery
	selectElements       []*SQLSelectElement
	BulkCollect          bool
	selectTargetElements []*SQLSelectTargetElement
	fromClause           *SQLFromClause
	whereClause          *SQLWhereClause
	groupByClause        ISQLGroupByClause
	windowClause         ISQLGroupByClause
}

func NewSelectQuery() *SQLSelectQuery {
	x := new(SQLSelectQuery)
	x.AbstractSQLSelectQuery = NewAbstractSelectQuery()
	return x
}

func (x *SQLSelectQuery) SelectElements() []*SQLSelectElement {
	return x.selectElements
}

func (x *SQLSelectQuery) AddSelectElement(selectElement *SQLSelectElement) {
	selectElement.SetParent(x)
	x.selectElements = append(x.selectElements, selectElement)
}

func (x *SQLSelectQuery) SelectTargetElements() []*SQLSelectTargetElement {
	return x.selectTargetElements
}

func (x *SQLSelectQuery) AddSelectTargetElement(expr expr.ISQLExpr) {
	if ast.IsNil(expr) {
		return
	}
	var selectTargetElement *SQLSelectTargetElement
	switch expr.(type) {
	case SQLSelectTargetElement:
		selectTargetElement = expr.(*SQLSelectTargetElement)
		break
	default:
		selectTargetElement = NewSelectTargetElementWithExpr(expr)
	}
	selectTargetElement.SetParent(x)
	x.selectTargetElements = append(x.selectTargetElements, selectTargetElement)
}

func (x *SQLSelectQuery) FromClause() *SQLFromClause {
	return x.fromClause
}

func (x *SQLSelectQuery) SetFromClause(fromClause *SQLFromClause) {
	if ast.IsNil(fromClause) {
		return
	}
	fromClause.SetParent(x)
	x.fromClause = fromClause
}

func (x *SQLSelectQuery) WhereClause() *SQLWhereClause {
	return x.whereClause
}

func (x *SQLSelectQuery) SetWhereClause(whereClause *SQLWhereClause) {
	if ast.IsNil(whereClause) {
		return
	}
	whereClause.SetParent(x)
	x.whereClause = whereClause
}

func (x *SQLSelectQuery) GroupByClause() ISQLGroupByClause {
	return x.groupByClause
}

func (x *SQLSelectQuery) SetGroupByClause(groupByClause ISQLGroupByClause) {
	if ast.IsNil(groupByClause) {
		return
	}
	groupByClause.SetParent(x)
	x.groupByClause = groupByClause
}

func (x *SQLSelectQuery) WindowClause() ISQLGroupByClause {
	return x.windowClause
}

func (x *SQLSelectQuery) SetWindowClause(windowClause ISQLGroupByClause) {
	if ast.IsNil(windowClause) {
		return
	}
	windowClause.SetParent(x)
	x.windowClause = windowClause
}

type SQLParenSelectQuery struct {
	*AbstractSQLSelectQuery
	subQuery ISQLSelectQuery
}

func NewParenSelectQuery(subQuery ISQLSelectQuery) *SQLParenSelectQuery {
	x := new(SQLParenSelectQuery)
	x.AbstractSQLSelectQuery = NewAbstractSelectQuery()
	x.subQuery = subQuery
	return x
}

func (x *SQLParenSelectQuery) SubQuery() ISQLSelectQuery {
	return x.subQuery
}
func (x *SQLParenSelectQuery) SetSubQuery(subQuery ISQLSelectQuery) {
	if ast.IsNil(subQuery) {
		return
	}
	subQuery.SetParent(x)
	x.subQuery = subQuery
}

type SQLSelectUnionQuery struct {
	*AbstractSQLSelectQuery

	left     ISQLSelectQuery
	Operator SQLUnionOperator
	right    ISQLSelectQuery
}

func NewSelectUnionQuery(left ISQLSelectQuery, operator SQLUnionOperator, right ISQLSelectQuery) *SQLSelectUnionQuery {
	x := new(SQLSelectUnionQuery)
	x.AbstractSQLSelectQuery = NewAbstractSelectQuery()
	x.SetLeft(left)
	x.Operator = operator
	x.SetRight(right)
	return x
}

func (x *SQLSelectUnionQuery) Left() ISQLSelectQuery {
	return x.left
}

func (x *SQLSelectUnionQuery) SetLeft(left ISQLSelectQuery) {
	left.SetParent(x)
	x.left = left
}
func (x *SQLSelectUnionQuery) Right() ISQLSelectQuery {
	return x.right
}
func (x *SQLSelectUnionQuery) SetRight(right ISQLSelectQuery) {
	right.SetParent(x)
	x.right = right
}

type SQLUnionOperator string

var (
	UNION          SQLUnionOperator = "UNION"
	UNION_ALL      SQLUnionOperator = "UNION ALL"
	UNION_DISTINCT SQLUnionOperator = "UNION DISTINCT"

	MINUS SQLUnionOperator = "MINUS"

	EXCEPT          SQLUnionOperator = "EXCEPT"
	EXCEPT_ALL      SQLUnionOperator = "EXCEPT ALL"
	EXCEPT_DISTINCT SQLUnionOperator = "EXCEPT DISTINCT"

	INTERSECT          SQLUnionOperator = "INTERSECT"
	INTERSECT_ALL      SQLUnionOperator = "INTERSECT ALL"
	INTERSECT_DISTINCT SQLUnionOperator = "INTERSECT DISTINCT"
)

/**
 *
 */
type ISQLFromClause interface {
	expr.ISQLExpr
}

/**
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#from%20clause
 */
type SQLFromClause struct {
	*expr.SQLExpr
	tableReference ISQLTableReference
}

func NewFromClause(tableReference ISQLTableReference) *SQLFromClause {
	var x SQLFromClause
	x.SQLExpr = expr.NewExpr()
	x.SetTableReference(tableReference)
	return &x
}

func (x *SQLFromClause) TableReference() ISQLTableReference {
	return x.tableReference
}
func (x *SQLFromClause) SetTableReference(tableReference ISQLTableReference) {
	if ast.IsNil(tableReference) {
		return
	}
	tableReference.SetParent(x)
	x.tableReference = tableReference
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20reference%20list
//
type ISQLTableReference interface {
	expr.ISQLExpr

	SetParen(paren bool)
	SetAs(as bool)
	SetAlias(alias expr.ISQLIdentifier)
}

type AbstractSQLTableReference struct {
	*expr.SQLExpr
	paren   bool
	as      bool
	alias   expr.ISQLIdentifier
	columns []expr.ISQLExpr
}

func NewAbstractTableReference() *AbstractSQLTableReference {
	x := new(AbstractSQLTableReference)
	x.SQLExpr = expr.NewExpr()
	return x
}

func (x *AbstractSQLTableReference) Paren() bool {
	return x.paren
}

func (x *AbstractSQLTableReference) SetParen(paren bool) {
	x.paren = paren
}

func (x *AbstractSQLTableReference) As() bool {
	return x.as
}

func (x *AbstractSQLTableReference) SetAs(as bool) {
	x.as = as
}
func (x *AbstractSQLTableReference) Alias() expr.ISQLIdentifier {
	return x.alias
}

func (x *AbstractSQLTableReference) SetAlias(alias expr.ISQLIdentifier) {
	if ast.IsNil(alias) {
		return
	}
	alias.SetParent(x)
	x.alias = alias
}
func (x *AbstractSQLTableReference) Columns() []expr.ISQLExpr {
	return x.columns
}

func (x *AbstractSQLTableReference) AddColumn(column expr.ISQLExpr) {
	if ast.IsNil(column) {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}

type SQLTableReference struct {
	*AbstractSQLTableReference
	name expr.ISQLName
}

func NewTableReference() *SQLTableReference {
	x := new(SQLTableReference)
	x.AbstractSQLTableReference = NewAbstractTableReference()
	return x
}

func NewTableReferenceWithAlias(name expr.ISQLName, as bool, alias expr.ISQLIdentifier) *SQLTableReference {
	x := new(SQLTableReference)
	x.AbstractSQLTableReference = NewAbstractTableReference()
	x.SetName(name)
	x.as = as
	x.SetAlias(alias)
	return x
}

func (x *SQLTableReference) Name() expr.ISQLName {
	return x.name
}

func (x *SQLTableReference) SetName(name expr.ISQLName) {
	if ast.IsNil(name) {
		return
	}
	name.SetParent(x)
	x.name = name
}

//
type SQLSubQueryTableReference struct {
	*AbstractSQLTableReference
	subQuery ISQLSelectQuery
}

func NewSubQueryTableReference(subQuery ISQLSelectQuery) *SQLSubQueryTableReference {
	x := new(SQLSubQueryTableReference)
	x.SetSubQuery(subQuery)
	return x
}

func (x *SQLSubQueryTableReference) SubQuery() ISQLSelectQuery {
	return x.subQuery
}

func (x *SQLSubQueryTableReference) SetSubQuery(subQuery ISQLSelectQuery) {
	if ast.IsNil(subQuery) {
		return
	}
	subQuery.SetParent(x)
	x.subQuery = subQuery
}

type SQLJoinTableReference struct {
	*AbstractSQLTableReference

	left ISQLTableReference

	JoinType SQLJoinType

	right ISQLTableReference

	condition ISQLJoinCondition
}

func NewJoinTableReference(left ISQLTableReference, joinType SQLJoinType, right ISQLTableReference) *SQLJoinTableReference {
	x := new(SQLJoinTableReference)
	x.SetLeft(left)
	x.JoinType = joinType
	x.SetRight(right)
	return x
}

func (x *SQLJoinTableReference) Left() ISQLTableReference {
	return x.left
}

func (x *SQLJoinTableReference) SetLeft(left ISQLTableReference) {
	left.SetParent(x)
	x.left = left
}

func (x *SQLJoinTableReference) Right() ISQLTableReference {
	return x.right
}

func (x *SQLJoinTableReference) SetRight(right ISQLTableReference) {
	right.SetParent(x)
	x.right = right
}
func (x *SQLJoinTableReference) Condition() ISQLJoinCondition {
	return x.condition
}

func (x *SQLJoinTableReference) SetCondition(condition ISQLJoinCondition) {
	condition.SetParent(x)
	x.condition = condition
}

// ON condition
// USING (...,...)
type ISQLJoinCondition interface {
	expr.ISQLExpr
}

type SQLJoinOnCondition struct {
	*expr.SQLExpr
	condition expr.ISQLExpr
}

func NewJoinOnCondition() *SQLJoinOnCondition {
	x := new(SQLJoinOnCondition)
	return x
}

func (x *SQLJoinOnCondition) Condition() expr.ISQLExpr {
	return x.condition
}
func (x *SQLJoinOnCondition) SetCondition(condition expr.ISQLExpr) {
	if condition == nil {
		return
	}
	condition.SetParent(x)
	x.condition = condition
}

type SQLJoinUsingCondition struct {
	*expr.SQLExpr
	columns []expr.ISQLExpr
}

func NewJoinUsingCondition() *SQLJoinUsingCondition {
	x := new(SQLJoinUsingCondition)
	return x
}
func (x *SQLJoinUsingCondition) Columns() []expr.ISQLExpr {
	return x.columns
}
func (x *SQLJoinUsingCondition) AddColumn(column expr.ISQLExpr) {
	column.SetParent(x)
	x.columns = append(x.columns, column)
}

type SQLJoinType string

const (
	COMMA      = ","
	JOIN       = "JOIN"
	INNER_JOIN = "INNER JOIN"

	CROSS_JOIN  = "CROSS_JOIN"
	CROSS_APPLY = "CROSS_APPLY"

	LEFT_JOIN       = "LEFT_JOIN"
	LEFT_OUTER_JOIN = "LEFT_OUTER_JOIN"

	RIGHT_JOIN       = "RIGHT_JOIN"
	RIGHT_OUTER_JOIN = "RIGHT_OUTER_JOIN"

	FULL_JOIN       = "FULL_JOIN"
	FULL_OUTER_JOIN = "FULL_OUTER_JOIN"

	NATURAL_JOIN             = "NATURAL_JOIN"
	NATURAL_INNER_JOIN       = "NATURAL_INNER_JOIN"
	NATURAL_LEFT_JOIN        = "NATURAL_LEFT_JOIN"
	NATURAL_LEFT_OUTER_JOIN  = "NATURAL_LEFT_OUTER_JOIN"
	NATURAL_RIGHT_JOIN       = "NATURAL_RIGHT_JOIN"
	NATURAL_RIGHT_OUTER_JOIN = "NATURAL_RIGHT_OUTER_JOIN"
	NATURAL_FULL_JOIN        = "NATURAL_FULL_JOIN"
	NATURAL_FULL_OUTER_JOIN  = "NATURAL_FULL_OUTER_JOIN"

	OUTER_APPLY = "OUTER_APPLY"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#where%20clause
// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
type SQLWhereClause struct {
	*expr.SQLExpr
	condition expr.ISQLExpr
}

func NewWhereClause(condition expr.ISQLExpr) *SQLWhereClause {
	if ast.IsNil(condition) {
		panic("condition is nil.")
	}
	var x SQLWhereClause
	x.SQLExpr = expr.NewExpr()
	x.SetCondition(condition)
	return &x
}

func (x *SQLWhereClause) Condition() expr.ISQLExpr {
	return x.condition
}
func (x *SQLWhereClause) SetCondition(condition expr.ISQLExpr) {
	if ast.IsNil(condition) {
		return
	}
	condition.SetParent(x)
	x.condition = condition
}

/**
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#group%20by%20clause
 * https://dev.mysql.com/doc/refman/8.0/en/select.html
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/img_text/group_by_clause.html
 */

type ISQLGroupByClause interface {
	expr.ISQLExpr

	Quantifier() expr.SQLSetQuantifier
	SetQuantifier(quantifier expr.SQLSetQuantifier)

	Elements() []*SQLGroupByElement
	AddElement(element *SQLGroupByElement)

	WithRollup() bool
	SetWithRollup(withRollup bool)

	Having() expr.ISQLExpr
	SetHaving(having expr.ISQLExpr)
}

// GROUP BY ...,.. [HAVING having]
// HAVING having [GROUP BY ...,..]
type AbstractSQLGroupByClause struct {
	*expr.SQLExpr

	quantifier expr.SQLSetQuantifier
	elements   []*SQLGroupByElement
	withRollup bool
	having     expr.ISQLExpr
}

func (x *AbstractSQLGroupByClause) Quantifier() expr.SQLSetQuantifier {
	return x.quantifier
}

func (x *AbstractSQLGroupByClause) SetQuantifier(quantifier expr.SQLSetQuantifier) {
	x.quantifier = quantifier
}

func (x *AbstractSQLGroupByClause) Elements() []*SQLGroupByElement {
	return x.elements
}

func (x *AbstractSQLGroupByClause) AddElement(element *SQLGroupByElement) {
	element.SetParent(x)
	x.elements = append(x.elements, element)
}

func (x *AbstractSQLGroupByClause) WithRollup() bool {
	return x.withRollup
}

func (x *AbstractSQLGroupByClause) SetWithRollup(withRollup bool) {
	x.withRollup = withRollup
}

func (x *AbstractSQLGroupByClause) Having() expr.ISQLExpr {
	return x.having
}

func (x *AbstractSQLGroupByClause) SetHaving(having expr.ISQLExpr) {
	having.SetParent(x)
	x.having = having
}

// GROUP BY ...,.. [HAVING having]
type SQLGroupByHavingClause struct {
	*AbstractSQLGroupByClause
}

func NewGroupByHavingClause() *SQLGroupByHavingClause {
	x := new(SQLGroupByHavingClause)
	return x
}

// HAVING having [GROUP BY ...,..]
type SQLHavingGroupByClause struct {
	*AbstractSQLGroupByClause
}

func NewHavingGroupByClause() *SQLHavingGroupByClause {
	x := new(SQLHavingGroupByClause)
	return x
}

type SQLGroupByElement struct {
	*expr.SQLExpr
	expr expr.ISQLExpr
}

func NewGroupByElement(expr expr.ISQLExpr) *SQLGroupByElement {
	x := new(SQLGroupByElement)
	x.SetExpr(expr)
	return x
}

func (x *SQLGroupByElement) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLGroupByElement) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#order%20by%20clause
// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
type SQLOrderByClause struct {
	expr.SQLExpr

	siblings bool
	elements []*SQLOrderByElement
}

func NewOrderByClause() *SQLOrderByClause {
	x := new(SQLOrderByClause)
	x.elements = make([]*SQLOrderByElement, 10)
	return x
}

func (s *SQLOrderByClause) AddElement(x *SQLOrderByElement) {
	s.elements = append(s.elements, x)
}

type SQLOrderByElement struct {
	expr.SQLExpr
	key           expr.ISQLExpr
	specification SQLOrderingSpecification
	nullOrdering  SQLNullOrdering
}

func (x *SQLOrderByElement) SetKey(key expr.ISQLExpr) {
	key.SetParent(x)
	x.key = key
}

func NewOrderByElement(key expr.ISQLExpr) *SQLOrderByElement {
	x := new(SQLOrderByElement)
	x.SetKey(key)
	return x
}

func NewOrderByElementWithSpecification(key expr.ISQLExpr, specification SQLOrderingSpecification) *SQLOrderByElement {
	x := new(SQLOrderByElement)
	x.SetKey(key)
	x.specification = specification
	return x
}

type SQLOrderingSpecification string

const (
	ASC  SQLOrderingSpecification = "ASC"
	DESC                          = "DESC"
)

type SQLNullOrdering string

func (x SQLNullOrdering) Lower() string {
	return strings.ToLower(string(x))
}
func (x SQLNullOrdering) Upper() string {
	return string(x)
}

const (
	NULLS_FIRST SQLNullOrdering = "NULLS FIRST"
	NULLS_LAST                  = "NULLS LAST"
)

func newOrderByClause(condition expr.ISQLExpr) *SQLOrderByClause {
	return &SQLOrderByClause{}
}

/**
 * [LIMIT {[offset,] row_count | row_count OFFSET offset}]
 * [ LIMIT { count | ALL } ] [ OFFSET start [ ROW | ROWS ] ]
 * [ OFFSET start [ ROW | ROWS ] ] [ FETCH { FIRST | NEXT } [ count ] { ROW | ROWS } ONLY ]
 * [ OFFSET offset { ROW | ROWS } ] [ FETCH { FIRST | NEXT } [ { rowcount | percent PERCENT } ] { ROW | ROWS } { ONLY | WITH TIES } ]
 */
type ISQLLimitClause interface {
	expr.ISQLExpr
}

type SQLRowType string

const (
	ROW  = "ROW"
	ROWS = "ROWS"
)

/**
 * [LIMIT {[offset,] row_count | row_count OFFSET offset}]
 * https://dev.mysql.com/doc/refman/8.0/en/select.html
 * <p>
 * [ LIMIT { count | ALL } ] [ OFFSET start [ ROW | ROWS ] ]
 * https://www.postgresql.org/docs/devel/static/sql-select.html
 *
 */
type SQLLimitOffsetClause struct {
	*expr.SQLExpr
	Offset        bool
	offsetExpr    expr.ISQLExpr
	countExpr     expr.ISQLExpr
	OffSetRowType SQLRowType
}

func NewLimitOffsetClause() *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	return x
}

func NewLimitOffsetClauseWithCount(countExpr expr.ISQLExpr) *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	x.SetCountExpr(countExpr)
	return x
}

func NewLimitOffsetClauseWithOffset(offset bool, offsetExpr expr.ISQLExpr, countExpr expr.ISQLExpr) *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	x.Offset = offset
	x.SetOffsetExpr(offsetExpr)
	x.SetCountExpr(countExpr)
	return x
}

func NewLimitOffsetClauseWithOffsetRowType(offset bool, offsetExpr expr.ISQLExpr, countExpr expr.ISQLExpr, offSetRowType SQLRowType) *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	x.Offset = offset
	x.SetOffsetExpr(offsetExpr)
	x.SetCountExpr(countExpr)
	return x
}

func (x *SQLLimitOffsetClause) OffsetExpr() expr.ISQLExpr {
	return x.offsetExpr
}
func (x *SQLLimitOffsetClause) SetOffsetExpr(offsetExpr expr.ISQLExpr) {
	if offsetExpr == nil {
		panic("offsetExpr is nil.")
	}
	offsetExpr.SetParent(x)
	x.offsetExpr = offsetExpr
}

func (x *SQLLimitOffsetClause) CountExpr() expr.ISQLExpr {
	return x.countExpr
}
func (x *SQLLimitOffsetClause) SetCountExpr(countExpr expr.ISQLExpr) {
	if countExpr == nil {
		panic("countExpr is nil.")
	}
	countExpr.SetParent(x)
	x.countExpr = countExpr
}

/**
 * [ OFFSET start [ ROW | ROWS ] ] [ FETCH { FIRST | NEXT } [ count ] { ROW | ROWS } ONLY ]
 * [ OFFSET offset { ROW | ROWS } ] [ FETCH { FIRST | NEXT } [ { rowcount | percent PERCENT } ] { ROW | ROWS } { ONLY | WITH TIES } ]
 * <p>
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 * <p>
 * https://www.postgresql.org/docs/devel/static/sql-values.html
 *
 */
type SQLFetchType string

const (
	FIRST SQLFetchType = "FIRST"
	NEXT  SQLFetchType = "NEXT"
)

type SQLOnly string

const (
	ONLY      SQLOnly = "ONLY"
	WITH_TIES SQLOnly = "WITH TIES"
)

type SQLOffsetFetchClause struct {
	*expr.SQLExpr
	offsetExpr    expr.ISQLExpr
	OffSetRowType SQLRowType

	FetchType    SQLFetchType
	countExpr    expr.ISQLExpr
	Percent      bool
	FetchRowType SQLRowType
	OnlyType     SQLOnly
}

func NewOffsetFetchClause() *SQLOffsetFetchClause {
	x := new(SQLOffsetFetchClause)
	return x
}

func (x *SQLOffsetFetchClause) OffsetExpr() expr.ISQLExpr {
	return x.offsetExpr
}
func (x *SQLOffsetFetchClause) SetOffsetExpr(offsetExpr expr.ISQLExpr) {
	if offsetExpr == nil {
		panic("offsetExpr is nil.")
	}
	offsetExpr.SetParent(x)
	x.offsetExpr = offsetExpr
}

func (x *SQLOffsetFetchClause) CountExpr() expr.ISQLExpr {
	return x.countExpr
}
func (x *SQLOffsetFetchClause) SetCountExpr(countExpr expr.ISQLExpr) {
	if countExpr == nil {
		panic("countExpr is nil.")
	}
	countExpr.SetParent(x)
	x.countExpr = countExpr
}

type ISQLLockClause interface {
	expr.ISQLExpr
}

type SQLForUpdateClause struct {
	*expr.SQLExpr
}

func NewForUpdateClause() *SQLForUpdateClause {
	x := new(SQLForUpdateClause)
	return x
}

type ISQLReturningClause interface {
	expr.ISQLExpr
}

type SQLReturnIntoClause struct {
	expr.SQLExpr
}
type SQLReturningIntoClause struct {
	expr.SQLExpr
}
