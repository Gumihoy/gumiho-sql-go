package select_

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
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
	*expr.AbstractSQLExpr

	orderByClause *SQLOrderByClause
	limitClause   ISQLLimitClause
	lockClause    ISQLLockClause
}

func NewAbstractSelectQuery() *AbstractSQLSelectQuery {
	x := new(AbstractSQLSelectQuery)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *AbstractSQLSelectQuery) OrderByClause() *SQLOrderByClause {
	return x.orderByClause
}

func (x *AbstractSQLSelectQuery) SetOrderByClause(orderByClause *SQLOrderByClause) {
	if orderByClause == nil {
		return
	}
	orderByClause.SetParent(x)
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
	*expr.AbstractSQLExpr

	expr  expr.ISQLExpr
	As    bool
	alias expr.ISQLExpr
}

func NewSelectElement() *SQLSelectElement {
	x := new(SQLSelectElement)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
	*expr.AbstractSQLExpr
	expr expr.ISQLExpr
}

func NewSelectTargetElement() *SQLSelectTargetElement {
	x := new(SQLSelectTargetElement)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
	withClause              ISQLWithClause
	selectElements          []*SQLSelectElement
	BulkCollect             bool
	selectTargetElements    []*SQLSelectTargetElement
	fromClause              *SQLFromClause
	whereClause             *SQLWhereClause
	hierarchicalQueryClause ISQLHierarchicalQueryClause
	groupByClause           ISQLGroupByClause
	windowClause            ISQLGroupByClause
}

func NewSelectQuery() *SQLSelectQuery {
	x := new(SQLSelectQuery)
	x.AbstractSQLSelectQuery = NewAbstractSelectQuery()
	return x
}

func (x *SQLSelectQuery) WithClause() ISQLWithClause {
	return x.withClause
}

func (x *SQLSelectQuery) SetWithClause(withClause ISQLWithClause) {
	if withClause == nil {
		return
	}
	withClause.SetParent(x)
	x.withClause = withClause
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

func (x *SQLSelectQuery) HierarchicalQueryClause() ISQLHierarchicalQueryClause {
	return x.hierarchicalQueryClause
}

func (x *SQLSelectQuery) SetHierarchicalQueryClause(hierarchicalQueryClause ISQLHierarchicalQueryClause) {
	if ast.IsNil(hierarchicalQueryClause) {
		return
	}
	hierarchicalQueryClause.SetParent(x)
	x.hierarchicalQueryClause = hierarchicalQueryClause
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
 * WITH [RECURSIVE]
        cte_name [(col_name [, col_name] ...)] AS (subquery)
        [, cte_name [(col_name [, col_name] ...)] AS (subquery)] ...
 * https://dev.mysql.com/doc/refman/8.0/en/with.html
 */
type ISQLWithClause interface {
	expr.ISQLExpr
}

/**
 * WITH [RECURSIVE]
        cte_name [(col_name [, col_name] ...)] AS (subquery)
        [, cte_name [(col_name [, col_name] ...)] AS (subquery)] ...
 *
 * query_name ([c_alias [, c_alias]...]) AS (subquery) [search_clause] [cycle_clause]
 */
type SQLWithClause struct {
	*expr.AbstractSQLExpr
	Recursive bool

	factoringClauses []ISQLFactoringClause
}

func NewWithClause() *SQLWithClause {
	x := new(SQLWithClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()

	x.factoringClauses = make([]ISQLFactoringClause, 0, 10)

	return x
}

func (x *SQLWithClause) FactoringClause() []ISQLFactoringClause {
	return x.factoringClauses
}

func (x *SQLWithClause) AddFactoringClause(factoringClause ISQLFactoringClause) {
	if factoringClause == nil {
		return
	}
	factoringClause.SetParent(x)
	x.factoringClauses = append(x.factoringClauses, factoringClause)
}

/**
 * subquery_factoring_clause, or subav_factoring_clause
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type ISQLFactoringClause interface {
	expr.ISQLExpr
	Name() expr.ISQLExpr
	SetName(name expr.ISQLExpr)
}
type abstractSQLFactoringClause struct {
	*expr.AbstractSQLExpr
	name expr.ISQLExpr
}

func NewAbstractSQLFactoringClause() *abstractSQLFactoringClause {
	x := new(abstractSQLFactoringClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *abstractSQLFactoringClause) Name() expr.ISQLExpr {
	return x.name
}

func (x *abstractSQLFactoringClause) SetName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 * cte_name [(col_name [, col_name] ...)] AS (subquery)
 * https://dev.mysql.com/doc/refman/8.0/en/with.html
 *
 * subquery_factoring_clause: cte_name [(col_name [, col_name] ...)] AS (subquery)
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type SQLSubQueryFactoringClause struct {
	*abstractSQLFactoringClause
	columns  []expr.ISQLIdentifier
	subQuery ISQLSelectQuery
}

func NewSubQueryFactoringClause() *SQLSubQueryFactoringClause {
	x := new(SQLSubQueryFactoringClause)
	x.abstractSQLFactoringClause = NewAbstractSQLFactoringClause()

	x.columns = make([]expr.ISQLIdentifier, 0, 10)

	return x
}

func (x *SQLSubQueryFactoringClause) Columns() []expr.ISQLIdentifier {
	return x.columns
}

func (x *SQLSubQueryFactoringClause) AddColumn(column expr.ISQLIdentifier) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}

func (x *SQLSubQueryFactoringClause) SubQuery() ISQLSelectQuery {
	return x.subQuery
}

func (x *SQLSubQueryFactoringClause) SetSubQuery(subQuery ISQLSelectQuery) {
	if subQuery == nil {
		return
	}
	subQuery.SetParent(x)
	x.subQuery = subQuery
}

/**
 * SEARCH
        { DEPTH FIRST BY c_alias [, c_alias]...
            [ ASC | DESC ]
            [ NULLS FIRST | NULLS LAST ]
         | BREADTH FIRST BY c_alias [, c_alias]...
            [ ASC | DESC ]
            [ NULLS FIRST | NULLS LAST ]
        }
        SET ordering_column
 */
type SQLSearchClause struct {
	*expr.AbstractSQLExpr
	columns       []expr.ISQLIdentifier
	specification SQLOrderingSpecification
	nullOrdering  SQLNullOrdering
}

/**
 * subav_factoring_clause: subav_name ANALYTIC VIEW AS (subav_clause)
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type SQLSubAvFactoringClause struct {
	*abstractSQLFactoringClause
	subAvClause *SQLSubAvClause
}

func NewSubAvFactoringClause() *SQLSubAvFactoringClause {
	x := new(SQLSubAvFactoringClause)
	x.abstractSQLFactoringClause = NewAbstractSQLFactoringClause()
	return x
}

func (x *SQLSubAvFactoringClause) SubAvClause() *SQLSubAvClause {
	return x.subAvClause
}

func (x *SQLSubAvFactoringClause) SetSubAvClause(subAvClause *SQLSubAvClause) {
	if subAvClause == nil {
		return
	}
	subAvClause.SetParent(x)
	x.subAvClause = subAvClause
}

/**
 * USING  { [schema.]base_av_name [hierarchies_clause] [filter_clauses] [add_calcs_clause] }
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type SQLSubAvClause struct {
	*expr.AbstractSQLExpr

	name expr.ISQLName
}

func (x *SQLSubAvClause) Name() expr.ISQLName {
	return x.name
}

func (x *SQLSubAvClause) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

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
	*expr.AbstractSQLExpr
	tableReference ISQLTableReference
}

func NewFromClause(tableReference ISQLTableReference) *SQLFromClause {
	var x SQLFromClause
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetTableReference(tableReference)
	return &x
}

func (x *SQLFromClause) TableReference() ISQLTableReference {
	return x.tableReference
}
func (x *SQLFromClause) SetTableReference(tableReference ISQLTableReference) {
	if tableReference == nil {
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
	*expr.AbstractSQLExpr
	paren   bool
	as      bool
	alias   expr.ISQLIdentifier
	columns []expr.ISQLExpr
}

func NewAbstractTableReference() *AbstractSQLTableReference {
	x := new(AbstractSQLTableReference)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
	name                     expr.ISQLExpr
	partitionExtensionClause expr.ISQLExpr
	sameClause               *SQLSampleClause
}

func NewTableReference() *SQLTableReference {
	x := new(SQLTableReference)
	x.AbstractSQLTableReference = NewAbstractTableReference()
	return x
}

func NewTableReferenceWithAlias(name expr.ISQLExpr, as bool, alias expr.ISQLIdentifier) *SQLTableReference {
	x := new(SQLTableReference)
	x.AbstractSQLTableReference = NewAbstractTableReference()
	x.SetName(name)
	x.as = as
	x.SetAlias(alias)
	return x
}

func (x *SQLTableReference) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.name {
		x.SetName(target)
		return true
	}
	if source == x.partitionExtensionClause {
		x.SetPartitionExtensionClause(target)
		return true
	}
	if source == x.sameClause {
		switch target.(type) {
		case *SQLSampleClause:
			x.SetSampleClause(target.(*SQLSampleClause))
			return true
		}
	}
	return false
}

func (x *SQLTableReference) Name() expr.ISQLExpr {
	return x.name
}

func (x *SQLTableReference) SetName(name expr.ISQLExpr) {
	if ast.IsNil(name) {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLTableReference) PartitionExtensionClause() expr.ISQLExpr {
	return x.partitionExtensionClause
}

func (x *SQLTableReference) SetPartitionExtensionClause(partitionExtensionClause expr.ISQLExpr) {
	if partitionExtensionClause == nil {
		return
	}
	partitionExtensionClause.SetParent(x)
	x.partitionExtensionClause = partitionExtensionClause
}
func (x *SQLTableReference) SampleClause() *SQLSampleClause {
	return x.sameClause
}

func (x *SQLTableReference) SetSampleClause(sameClause *SQLSampleClause) {
	if sameClause == nil {
		return
	}
	sameClause.SetParent(x)
	x.sameClause = sameClause
}

/**
 * PARTITION (partition_names)
 * https://dev.mysql.com/doc/refman/8.0/en/join.html
 *
 * { PARTITION (partition)
| PARTITION FOR (partition_key_value [, partition_key_value]...)
| SUBPARTITION (subpartition)
| SUBPARTITION FOR (subpartition_key_value [, subpartition_key_value]...)
}
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/img_text/partition_extension_clause.html
 */
type AbstractSQLPartitionExtensionClause struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLExpr
}

func NewAbstractSQLPartitionExtensionClause() *AbstractSQLPartitionExtensionClause {
	x := new(AbstractSQLPartitionExtensionClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.names = make([]expr.ISQLExpr, 0, 10)
	return x
}

func (x *AbstractSQLPartitionExtensionClause) Names() []expr.ISQLExpr {
	return x.names
}
func (x *AbstractSQLPartitionExtensionClause) Name(i int) expr.ISQLExpr {
	return x.names[i]
}
func (x *AbstractSQLPartitionExtensionClause) AddName(name expr.ISQLExpr) {
	name.SetParent(x)
	x.names = append(x.names, name)
}

/**
 * PARTITION (name [, name]...)
 */
type SQLPartitionClause struct {
	*AbstractSQLPartitionExtensionClause
}

func NewPartitionClause() *SQLPartitionClause {
	x := new(SQLPartitionClause)
	x.AbstractSQLPartitionExtensionClause = NewAbstractSQLPartitionExtensionClause()
	return x
}

/**
 * PARTITION FOR (name [, name]...)
 */
type SQLPartitionForClause struct {
	*AbstractSQLPartitionExtensionClause
}

func NewPartitionForClause() *SQLPartitionForClause {
	x := new(SQLPartitionForClause)
	x.AbstractSQLPartitionExtensionClause = NewAbstractSQLPartitionExtensionClause()
	return x
}

/**
 * SUBPARTITION (name [, name]...)
 */
type SQLSubPartitionClause struct {
	*AbstractSQLPartitionExtensionClause
}

func NewSubPartitionClause() *SQLSubPartitionClause {
	x := new(SQLSubPartitionClause)
	x.AbstractSQLPartitionExtensionClause = NewAbstractSQLPartitionExtensionClause()
	return x
}

/**
 * SUBPARTITION FOR (name [, name]...)
 */
type SQLSubPartitionForClause struct {
	*AbstractSQLPartitionExtensionClause
}

func NewSubPartitionForClause() *SQLSubPartitionForClause {
	x := new(SQLSubPartitionForClause)
	x.AbstractSQLPartitionExtensionClause = NewAbstractSQLPartitionExtensionClause()
	return x
}

/**
 * sample_clause:
	SAMPLE [ BLOCK ]
       (sample_percent)
       [ SEED (seed_value) ]
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type SQLSampleClause struct {
	*expr.AbstractSQLExpr
	Block     bool
	percent   expr.ISQLExpr
	seedValue expr.ISQLExpr
}

func NewSampleClause() *SQLSampleClause {
	x := new(SQLSampleClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *SQLSampleClause) Percent() expr.ISQLExpr {
	return x.percent
}

func (x *SQLSampleClause) SetPercent(percent expr.ISQLExpr) {
	if percent == nil {
		return
	}
	percent.SetParent(x)
	x.percent = percent
}
func (x *SQLSampleClause) SeedValue() expr.ISQLExpr {
	return x.seedValue
}

func (x *SQLSampleClause) SetSeedValue(seedValue expr.ISQLExpr) {
	if seedValue == nil {
		return
	}
	seedValue.SetParent(x)
	x.seedValue = seedValue
}

/**
 * flashback_query_clause
 * { VERSIONS BETWEEN { SCN | TIMESTAMP }
    { expr | MINVALUE } AND { expr | MAXVALUE }
| VERSIONS PERIOD FOR valid_time_column BETWEEN
    { expr | MINVALUE } AND { expr | MAXVALUE }
| AS OF { SCN | TIMESTAMP } expr
| AS OF PERIOD FOR valid_time_column expr
}
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type ISQLFlashbackQueryClause interface {
	expr.ISQLExpr
}

/**
 * AS OF PERIOD FOR valid_time_column expr
 */
type SQLFlashbackQueryAsOfPeriodForClause struct {
	*expr.AbstractSQLExpr
}

/**
 * { OJ table_reference }
 * https://dev.mysql.com/doc/refman/8.0/en/join.html
 */
type SQLOJTableReference struct {
	*AbstractSQLTableReference
	tableReference ISQLTableReference
}

func NewOJTableReference() *SQLOJTableReference {
	x := new(SQLOJTableReference)
	x.AbstractSQLTableReference = NewAbstractTableReference()
	return x
}
func (x *SQLOJTableReference) TableReference() ISQLTableReference {
	return x.tableReference
}

func (x *SQLOJTableReference) SetTableReference(tableReference ISQLTableReference) {
	if tableReference == nil {
		return
	}
	tableReference.SetParent(x)
	x.tableReference = tableReference
}

/**
 * [LATERAL] table_subquery [AS] alias [(col_list)]
 * https://dev.mysql.com/doc/refman/8.0/en/join.html
 *
 * ONLY ([ LATERAL ] (subquery [ subquery_restriction_clause ]))
 * [ LATERAL ] (subquery [ subquery_restriction_clause ])
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type SQLSubQueryTableReference struct {
	*AbstractSQLTableReference
	subQuery ISQLSelectQuery
}

func NewSubQueryTableReference(subQuery ISQLSelectQuery) *SQLSubQueryTableReference {
	x := new(SQLSubQueryTableReference)
	x.AbstractSQLTableReference = NewAbstractTableReference()
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
	x.AbstractSQLTableReference = NewAbstractTableReference()
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
	*expr.AbstractSQLExpr
	condition expr.ISQLExpr
}

func NewJoinOnConditionWithCondition(condition expr.ISQLExpr) *SQLJoinOnCondition {
	x := new(SQLJoinOnCondition)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetCondition(condition)
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
	*expr.AbstractSQLExpr
	columns []expr.ISQLExpr
}

func NewJoinUsingCondition() *SQLJoinUsingCondition {
	x := new(SQLJoinUsingCondition)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLJoinUsingCondition) Columns() []expr.ISQLExpr {
	return x.columns
}
func (x *SQLJoinUsingCondition) Column(i int) expr.ISQLExpr {
	return x.columns[i]
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

	CROSS_JOIN  = "CROSS JOIN"
	CROSS_APPLY = "CROSS APPLY"

	LEFT_JOIN       = "LEFT JOIN"
	LEFT_OUTER_JOIN = "LEFT OUTER JOIN"

	RIGHT_JOIN       = "RIGHT JOIN"
	RIGHT_OUTER_JOIN = "RIGHT OUTER JOIN"

	FULL_JOIN       = "FULL JOIN"
	FULL_OUTER_JOIN = "FULL OUTER JOIN"

	NATURAL_JOIN             = "NATURAL JOIN"
	NATURAL_INNER_JOIN       = "NATURAL INNER JOIN"
	NATURAL_LEFT_JOIN        = "NATURAL LEFT JOIN"
	NATURAL_LEFT_OUTER_JOIN  = "NATURAL LEFT OUTER JOIN"
	NATURAL_RIGHT_JOIN       = "NATURAL RIGHT JOIN"
	NATURAL_RIGHT_OUTER_JOIN = "NATURAL RIGHT OUTER JOIN"
	NATURAL_FULL_JOIN        = "NATURAL FULL JOIN"
	NATURAL_FULL_OUTER_JOIN  = "NATURAL FULL OUTER JOIN"

	OUTER_APPLY = "OUTER_APPLY"

	STRAIGHT_JOIN = "STRAIGHT_JOIN"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#where%20clause
// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
type SQLWhereClause struct {
	*expr.AbstractSQLExpr
	condition expr.ISQLExpr
}

func NewWhereClause(condition expr.ISQLExpr) *SQLWhereClause {
	if ast.IsNil(condition) {
		panic("condition is nil.")
	}
	var x SQLWhereClause
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
 * CONNECT BY [ NOCYCLE ] condition [ START WITH condition ]
 * | START WITH condition CONNECT BY [ NOCYCLE ] condition
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type ISQLHierarchicalQueryClause interface {
	expr.ISQLExpr
	ConnectByCondition() expr.ISQLExpr
	SetConnectByCondition(condition expr.ISQLExpr)
	StartWithCondition() expr.ISQLExpr
	SetStartWithCondition(condition expr.ISQLExpr)
}

type abstractSQLHierarchicalQueryClause struct {
	*expr.AbstractSQLExpr
	NoCycle            bool
	connectByCondition expr.ISQLExpr
	startWithCondition expr.ISQLExpr
}

func NewAbstractSQLHierarchicalQueryClause() *abstractSQLHierarchicalQueryClause {
	x := new(abstractSQLHierarchicalQueryClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *abstractSQLHierarchicalQueryClause) ConnectByCondition() expr.ISQLExpr {
	return x.connectByCondition
}
func (x *abstractSQLHierarchicalQueryClause) SetConnectByCondition(condition expr.ISQLExpr) {
	if condition == nil {
		return
	}
	condition.SetParent(x)
	x.connectByCondition = condition
}

func (x *abstractSQLHierarchicalQueryClause) StartWithCondition() expr.ISQLExpr {
	return x.startWithCondition
}
func (x *abstractSQLHierarchicalQueryClause) SetStartWithCondition(condition expr.ISQLExpr) {
	if condition == nil {
		return
	}
	condition.SetParent(x)
	x.startWithCondition = condition
}

type SQLHierarchicalQueryClauseConnectBy struct {
	*abstractSQLHierarchicalQueryClause
}

func NewHierarchicalQueryClauseConnectBy() *SQLHierarchicalQueryClauseConnectBy {
	x := new(SQLHierarchicalQueryClauseConnectBy)
	x.abstractSQLHierarchicalQueryClause = NewAbstractSQLHierarchicalQueryClause()
	return x
}

type SQLHierarchicalQueryClauseStartWith struct {
	*abstractSQLHierarchicalQueryClause
}

func NewHierarchicalQueryClauseStartWith() *SQLHierarchicalQueryClauseStartWith {
	x := new(SQLHierarchicalQueryClauseStartWith)
	x.abstractSQLHierarchicalQueryClause = NewAbstractSQLHierarchicalQueryClause()
	return x
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
	*expr.AbstractSQLExpr

	quantifier expr.SQLSetQuantifier
	elements   []*SQLGroupByElement
	withRollup bool
	having     expr.ISQLExpr
}

func NewAbstractSQLGroupByClause() *AbstractSQLGroupByClause {
	x := new(AbstractSQLGroupByClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x

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
func (x *AbstractSQLGroupByClause) Element(i int) *SQLGroupByElement {
	return x.elements[i]
}
func (x *AbstractSQLGroupByClause) AddElement(element *SQLGroupByElement) {
	if element == nil {
		return
	}
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
	if having == nil {
		return
	}
	having.SetParent(x)
	x.having = having
}

// GROUP BY ...,.. [HAVING having]
type SQLGroupByHavingClause struct {
	*AbstractSQLGroupByClause
}

func NewGroupByHavingClause() *SQLGroupByHavingClause {
	x := new(SQLGroupByHavingClause)
	x.AbstractSQLGroupByClause = NewAbstractSQLGroupByClause()
	return x
}

// HAVING having [GROUP BY ...,..]
type SQLHavingGroupByClause struct {
	*AbstractSQLGroupByClause
}

func NewHavingGroupByClause() *SQLHavingGroupByClause {
	x := new(SQLHavingGroupByClause)
	x.AbstractSQLGroupByClause = NewAbstractSQLGroupByClause()
	return x
}

type SQLGroupByElement struct {
	*expr.AbstractSQLExpr
	expr expr.ISQLExpr
}

func NewGroupByElement(ele expr.ISQLExpr) *SQLGroupByElement {
	x := new(SQLGroupByElement)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetExpr(ele)
	return x
}

func (x *SQLGroupByElement) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLGroupByElement) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

/**
 * MODEL
   [ cell_reference_options ]
   [ return_rows_clause ]
   [ reference_model ]...
main_model
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type SQLModelClause struct {
	*expr.AbstractSQLExpr

	name expr.ISQLExpr
}

func NewModelClause() *SQLModelClause {
	x := new(SQLModelClause)

	return x
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#order%20by%20clause
// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
type SQLOrderByClause struct {
	*expr.AbstractSQLExpr

	siblings bool
	elements []*SQLOrderByElement
}

func NewOrderByClause() *SQLOrderByClause {
	x := new(SQLOrderByClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.elements = make([]*SQLOrderByElement, 0, 10)
	return x
}
func (x *SQLOrderByClause) Elements() []*SQLOrderByElement {
	return x.elements
}
func (x *SQLOrderByClause) Element(i int) *SQLOrderByElement {
	return x.elements[i]
}
func (x *SQLOrderByClause) AddElement(element *SQLOrderByElement) {
	x.elements = append(x.elements, element)
}

type SQLOrderByElement struct {
	*expr.AbstractSQLExpr
	key           expr.ISQLExpr
	Specification SQLOrderingSpecification
	NullOrdering  SQLNullOrdering
}

func NewOrderByElement(key expr.ISQLExpr) *SQLOrderByElement {
	x := new(SQLOrderByElement)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetKey(key)
	return x
}

func NewOrderByElementWithSpecification(key expr.ISQLExpr, specification SQLOrderingSpecification) *SQLOrderByElement {
	x := new(SQLOrderByElement)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetKey(key)
	x.Specification = specification
	return x
}

func (x *SQLOrderByElement) Key() expr.ISQLExpr {
	return x.key
}
func (x *SQLOrderByElement) SetKey(key expr.ISQLExpr) {
	key.SetParent(x)
	x.key = key
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
	*expr.AbstractSQLExpr
	Offset        bool
	offsetExpr    expr.ISQLExpr
	countExpr     expr.ISQLExpr
	OffSetRowType SQLRowType
}

func NewLimitOffsetClause() *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func NewLimitOffsetClauseWithCount(countExpr expr.ISQLExpr) *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetCountExpr(countExpr)
	return x
}

func NewLimitOffsetClauseWithOffset(offset bool, offsetExpr expr.ISQLExpr, countExpr expr.ISQLExpr) *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.Offset = offset
	x.SetOffsetExpr(offsetExpr)
	x.SetCountExpr(countExpr)
	return x
}

func NewLimitOffsetClauseWithOffsetRowType(offset bool, offsetExpr expr.ISQLExpr, countExpr expr.ISQLExpr, offSetRowType SQLRowType) *SQLLimitOffsetClause {
	x := new(SQLLimitOffsetClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
	*expr.AbstractSQLExpr
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
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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

/**
 * FOR {UPDATE | SHARE} [OF tbl_name [, tbl_name] ...] [NOWAIT | SKIP LOCKED]
 * LOCK IN SHARE MODE
 * https://dev.mysql.com/doc/refman/8.0/en/select.html
 *
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type ISQLLockClause interface {
	expr.ISQLExpr
}

/**
 * FOR UPDATE [OF tbl_name [, tbl_name] ...] [NOWAIT | SKIP LOCKED]
 * FOR SHARE [OF tbl_name [, tbl_name] ...] [NOWAIT | SKIP LOCKED]
 */
type AbstractSQLLockForClause struct {
	*expr.AbstractSQLExpr
	tables   []expr.ISQLName
	waitExpr expr.ISQLExpr
}

func NewAbstractLockForClause() *AbstractSQLLockForClause {
	x := new(AbstractSQLLockForClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *AbstractSQLLockForClause) Tables() []expr.ISQLName {
	return x.tables
}
func (x *AbstractSQLLockForClause) Table(i int) expr.ISQLName {
	return x.tables[i]
}
func (x *AbstractSQLLockForClause) AddTable(table expr.ISQLName) {
	if table == nil {
		return
	}
	table.SetParent(x)
	x.tables = append(x.tables, table)
}
func (x *AbstractSQLLockForClause) WaitExpr() expr.ISQLExpr {
	return x.waitExpr
}
func (x *AbstractSQLLockForClause) SetWaitExpr(waitExpr expr.ISQLExpr) {
	if waitExpr == nil {
		return
	}
	waitExpr.SetParent(x)
	x.waitExpr = waitExpr
}

/**
 * NOWAIT
 */
type SQLLockForNoWaitExpr struct {
	*expr.AbstractSQLExpr
}

func NewLockForNoWaitExpr() *SQLLockForNoWaitExpr {
	x := new(SQLLockForNoWaitExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * WAIT expr
 */
type SQLLockForWaitExpr struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}

func NewLockForWaitExpr() *SQLLockForWaitExpr {
	x := new(SQLLockForWaitExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func NewLockForWaitExprWithValue(value expr.ISQLExpr) *SQLLockForWaitExpr {
	x := new(SQLLockForWaitExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetValue(value)
	return x
}
func (x *SQLLockForWaitExpr) Value() expr.ISQLExpr {
	return x.value
}
func (x *SQLLockForWaitExpr) SetValue(value expr.ISQLExpr) {
	if value == nil {
		panic("value is nil.")
	}
	value.SetParent(x)
	x.value = value
}

/**
 * SKIP LOCKED
 */
type SQLLockForSkipLockedExpr struct {
	*expr.AbstractSQLExpr
}

func NewLockForSkipLockedExpr() *SQLLockForSkipLockedExpr {
	x := new(SQLLockForSkipLockedExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * FOR UPDATE [OF tbl_name [, tbl_name] ...] [NOWAIT | SKIP LOCKED]
 * https://dev.mysql.com/doc/refman/8.0/en/select.html
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
type SQLForUpdateClause struct {
	*AbstractSQLLockForClause
}

func NewForUpdateClause() *SQLForUpdateClause {
	x := new(SQLForUpdateClause)
	x.AbstractSQLLockForClause = NewAbstractLockForClause()
	return x
}

/**
 * FOR SHARE [OF tbl_name [, tbl_name] ...] [NOWAIT | SKIP LOCKED]
 * https://dev.mysql.com/doc/refman/8.0/en/select.html
 *
 *
 */
type SQLForShareClause struct {
	*AbstractSQLLockForClause
}

func NewForShareClause() *SQLForShareClause {
	x := new(SQLForShareClause)
	x.AbstractSQLLockForClause = NewAbstractLockForClause()
	return x
}

/**
 * LOCK IN SHARE MODE
 * https://dev.mysql.com/doc/refman/8.0/en/select.html
 */
type SQLLockInShareModeClause struct {
	*expr.AbstractSQLExpr
}

func NewLockInShareModeClause() *SQLLockInShareModeClause {
	x := new(SQLLockInShareModeClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * WITH { READ ONLY | CHECK OPTION} [ CONSTRAINT constraint ]
 */
type SQLSubQueryRestrictionClause struct {
	*expr.AbstractSQLExpr

	name expr.ISQLName
}

/*
 *
 */
type ISQLReturningClause interface {
	expr.ISQLExpr
}

type SQLReturnIntoClause struct {
	*expr.AbstractSQLExpr
}

func NewReturnIntoClause() *SQLReturnIntoClause {
	x := new(SQLReturnIntoClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

type SQLReturningIntoClause struct {
	*expr.AbstractSQLExpr
}

func NewReturningIntoClause() *SQLReturningIntoClause {
	x := new(SQLReturningIntoClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
