package statement

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"reflect"
)

type SQLDropBehavior string

const (
	RESTRICT SQLDropBehavior = "RESTRICT"
	CASCADE                  = "CASCADE"

	FORCE    = "FORCE"
	VALIDATE = "VALIDATE"

	CASCADE_CONSTRAINTS = "CASCADE CONSTRAINTS"
)

type ISQLStatement interface {
	ast.ISQLObject
	ObjectType() db.SQLObjectType
}

type AbstractSQLStatement struct {
	*ast.SQLObject
}

func NewAbstractSQLStatementWithDBType(dbType db.Type) *AbstractSQLStatement {
	x := new(AbstractSQLStatement)
	x.SQLObject = ast.NewObjectWithDBType(dbType)
	return x
}
func (x *AbstractSQLStatement) ObjectType() db.SQLObjectType {
	xType := reflect.TypeOf(x)
	panic(xType)
}

/**
 * SELECT
    [ALL | DISTINCT | DISTINCTROW ]
    [HIGH_PRIORITY]
    [STRAIGHT_JOIN]
    [SQL_SMALL_RESULT] [SQL_BIG_RESULT] [SQL_BUFFER_RESULT]
    [SQL_NO_CACHE] [SQL_CALC_FOUND_ROWS]
    select_expr [, select_expr] ...
    [into_option]
    [FROM table_references
      [PARTITION partition_list]]
    [WHERE where_condition]
    [GROUP BY {col_name | expr | position}, ... [WITH ROLLUP]]
    [HAVING where_condition]
    [WINDOW window_name AS (window_spec)
        [, window_name AS (window_spec)] ...]
    [ORDER BY {col_name | expr | position}
      [ASC | DESC], ... [WITH ROLLUP]]
    [LIMIT {[offset,] row_count | row_count OFFSET offset}]
    [into_option]
    [FOR {UPDATE | SHARE}
        [OF tbl_name [, tbl_name] ...]
        [NOWAIT | SKIP LOCKED]
      | LOCK IN SHARE MODE]
    [into_option]

into_option: {
    INTO OUTFILE 'file_name'
        [CHARACTER SET charset_name]
        export_options
  | INTO DUMPFILE 'file_name'
  | INTO var_name [, var_name] ...
}
 * https://dev.mysql.com/doc/refman/8.0/en/select.html
 *
 *
 *
 */
type SQLSelectStatement struct {
	*AbstractSQLStatement
	query select_.ISQLSelectQuery
}

func NewSelectStatement(dbType db.Type, query select_.ISQLSelectQuery) *SQLSelectStatement {
	x := new(SQLSelectStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	x.SetQuery(query)
	return x
}

func (x *SQLSelectStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLSelectStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLSelectStatement) ObjectType() db.SQLObjectType {
	return db.SELECT
}

func (x *SQLSelectStatement) Query() select_.ISQLSelectQuery {
	return x.query
}
func (x *SQLSelectStatement) SetQuery(query select_.ISQLSelectQuery) {
	if query == nil {
		return
	}
	query.SetParent(x)
	x.query = query
}

/**
 * INSERT INTO <insertion target> <insert columns and source>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#insert%20statement
 *
 * INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    { {VALUES | VALUE} (value_list) [, (value_list)] ...
      |
      VALUES row_constructor_list
    }
    [AS row_alias[(col_alias [, col_alias] ...)]]
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [AS row_alias[(col_alias [, col_alias] ...)]]
    SET assignment_list
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    [AS row_alias[(col_alias [, col_alias] ...)]]
    {SELECT ... | TABLE table_name}
    [ON DUPLICATE KEY UPDATE assignment_list]
 * https://dev.mysql.com/doc/refman/8.0/en/insert.html
 *
 * INSERT [ hint ] INTO dml_table_expression_clause [ t_alias ] [ (column [, column ]...) ] { values_clause [ returning_clause ] | subquery} [ error_logging_clause ]
 * INSERT [ hint ]  { ALL { INTO dml_table_expression_clause [ t_alias ] [ (column [, column ]...) ] [ values_clause ] [error_logging_clause] }... | conditional_insert_clause } subquery
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/INSERT.html#GUID-903F8043-0254-4EE9-ACC1-CB8AC0AF3423
 */
type PriorityKind string

const (
	LOW_PRIORITY  PriorityKind = "LOW_PRIORITY"
	DELAYED                    = "DELAYED"
	HIGH_PRIORITY              = "HIGH_PRIORITY"
)

type ValueKind string

const (
	VALUES ValueKind = "VALUES"
	VALUE            = "VALUE"
)

/**
 * INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    { {VALUES | VALUE} (value_list) [, (value_list)] ...
      |
      VALUES row_constructor_list
    }
    [AS row_alias[(col_alias [, col_alias] ...)]]
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [AS row_alias[(col_alias [, col_alias] ...)]]
    SET assignment_list
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    [AS row_alias[(col_alias [, col_alias] ...)]]
    {SELECT ... | TABLE table_name}
    [ON DUPLICATE KEY UPDATE assignment_list]
 * https://dev.mysql.com/doc/refman/8.0/en/insert.html
 *
 *
 */
type SQLInsertStatement struct {
	*AbstractSQLStatement

	PriorityKind PriorityKind
	Ignore       bool

	Into           bool
	tableReference select_.ISQLTableReference

	columns []expr.ISQLIdentifier

	ValueKind ValueKind
	values    []expr.ISQLExpr

	// AS row_alias[(col_alias [, col_alias] ...)]

	subQuery select_.ISQLSelectQuery

	// MySQL: TABLE table_name
	table expr.ISQLName

	// MySQL: SET assignment_list
	setAssignments []expr.ISQLExpr

	// MySQL: ON DUPLICATE KEY UPDATE assignment_list
	updateAssignments []expr.ISQLExpr
}

func NewInsertStatement(dbType db.Type) *SQLInsertStatement {
	x := new(SQLInsertStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	x.Into = true
	x.ValueKind = VALUES
	return x
}
func (x *SQLInsertStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLInsertStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLInsertStatement) ObjectType() db.SQLObjectType {
	return db.INSERT
}

func (x *SQLInsertStatement) TableReference() select_.ISQLTableReference {
	return x.tableReference
}
func (x *SQLInsertStatement) SetTableReference(tableReference select_.ISQLTableReference) {
	tableReference.SetParent(x)
	x.tableReference = tableReference
}
func (x *SQLInsertStatement) Columns() []expr.ISQLIdentifier {
	return x.columns
}
func (x *SQLInsertStatement) Column(i int) expr.ISQLIdentifier {
	return x.columns[i]
}
func (x *SQLInsertStatement) AddColumn(column expr.ISQLIdentifier) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}
func (x *SQLInsertStatement) Values() []expr.ISQLExpr {
	return x.values
}
func (x *SQLInsertStatement) Value(i int) expr.ISQLExpr {
	return x.values[i]
}
func (x *SQLInsertStatement) AddValue(value expr.ISQLExpr) {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.values = append(x.values, value)
}
func (x *SQLInsertStatement) SubQuery() select_.ISQLSelectQuery {
	return x.subQuery
}
func (x *SQLInsertStatement) SetSubQuery(subQuery select_.ISQLSelectQuery) {
	subQuery.SetParent(x)
	x.subQuery = subQuery
}

func (x *SQLInsertStatement) UpdateAssignments() []expr.ISQLExpr {
	return x.updateAssignments
}
func (x *SQLInsertStatement) UpdateAssignment(i int) expr.ISQLExpr {
	return x.updateAssignments[i]
}
func (x *SQLInsertStatement) AddUpdateAssignment(assignment expr.ISQLExpr) {
	if assignment == nil {
		return
	}
	assignment.SetParent(x)
	x.updateAssignments = append(x.updateAssignments, assignment)
}
func (x *SQLInsertStatement) SetAssignments() []expr.ISQLExpr {
	return x.setAssignments
}
func (x *SQLInsertStatement) SetAssignment(i int) expr.ISQLExpr {
	return x.setAssignments[i]
}
func (x *SQLInsertStatement) AddSetAssignment(assignment expr.ISQLExpr) {
	if assignment == nil {
		return
	}
	assignment.SetParent(x)
	x.setAssignments = append(x.setAssignments, assignment)
}

/**
 * ROW(value_list)
 * https://dev.mysql.com/doc/refman/8.0/en/insert.html
 */
type SQLRowConstructorsExpr struct {
	*expr.AbstractSQLExpr
	values []expr.ISQLExpr
}

func NewRowConstructorsExpr() *SQLRowConstructorsExpr {
	x := new(SQLRowConstructorsExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *SQLRowConstructorsExpr) Values() []expr.ISQLExpr {
	return x.values
}
func (x *SQLRowConstructorsExpr) Value(i int) expr.ISQLExpr {
	return x.values[i]
}
func (x *SQLRowConstructorsExpr) AddTable(value expr.ISQLExpr) {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.values = append(x.values, value)
}

/**
 * UPDATE <target table> SET <set clause list> WHERE CURRENT OF <cursor name>
 * UPDATE <target table> SET <set clause list> [ WHERE <search condition> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#SQL%20data%20change%20statement
 *
 * UPDATE [LOW_PRIORITY] [IGNORE] table_reference
    SET assignment_list
    [WHERE where_condition]
    [ORDER BY ...]
    [LIMIT row_count]
 * https://dev.mysql.com/doc/refman/8.0/en/update.html
 *
 *
 */
type SQLUpdateStatement struct {
	*AbstractSQLStatement
	LowPriority bool
	Ignore      bool

	tableReference select_.ISQLTableReference

	assignments []expr.ISQLExpr

	whereClause   *select_.SQLWhereClause
	orderByClause *select_.SQLOrderByClause
	limitClause   select_.ISQLLimitClause
}

func NewUpdateStatement(dbType db.Type) *SQLUpdateStatement {
	x := new(SQLUpdateStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLUpdateStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLUpdateStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLUpdateStatement) ObjectType() db.SQLObjectType {
	return db.UPDATE
}

func (x *SQLUpdateStatement) TableReference() select_.ISQLTableReference {
	return x.tableReference
}
func (x *SQLUpdateStatement) SetTableReference(tableReference select_.ISQLTableReference) {
	tableReference.SetParent(x)
	x.tableReference = tableReference
}
func (x *SQLUpdateStatement) Assignments() []expr.ISQLExpr {
	return x.assignments
}
func (x *SQLUpdateStatement) Assignment(i int) expr.ISQLExpr {
	return x.assignments[i]
}
func (x *SQLUpdateStatement) AddAssignment(assignment expr.ISQLExpr) {
	if assignment == nil {
		return
	}
	assignment.SetParent(x)
	x.assignments = append(x.assignments, assignment)
}
func (x *SQLUpdateStatement) WhereClause() *select_.SQLWhereClause {
	return x.whereClause
}

func (x *SQLUpdateStatement) SetWhereClause(whereClause *select_.SQLWhereClause) {
	if ast.IsNil(whereClause) {
		return
	}
	whereClause.SetParent(x)
	x.whereClause = whereClause
}
func (x *SQLUpdateStatement) OrderByClause() *select_.SQLOrderByClause {
	return x.orderByClause
}

func (x *SQLUpdateStatement) SetOrderByClause(orderByClause *select_.SQLOrderByClause) {
	if orderByClause == nil {
		return
	}
	orderByClause.SetParent(orderByClause)
	x.orderByClause = orderByClause
}

func (x *SQLUpdateStatement) LimitClause() select_.ISQLLimitClause {
	return x.limitClause
}

func (x *SQLUpdateStatement) SetLimitClause(limitClause select_.ISQLLimitClause) {
	if limitClause == nil {
		return
	}
	limitClause.SetParent(x)
	x.limitClause = limitClause
}

/**
 * DELETE FROM <target table> WHERE CURRENT OF <cursor name>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#delete%20statement:%20positioned
 *
 * DELETE FROM <target table> [ WHERE <search condition> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#delete%20statement:%20searched
 *
 * Single-OnTable Syntax：
 * DELETE [LOW_PRIORITY] [QUICK] [IGNORE] FROM tbl_name [[AS] tbl_alias]
    [PARTITION (partition_name [, partition_name] ...)]
    [WHERE where_condition]
    [ORDER BY ...]
    [LIMIT row_count]
 *
 * Multiple-OnTable Syntax
 * DELETE [LOW_PRIORITY] [QUICK] [IGNORE]
    tbl_name[.*] [, tbl_name[.*]] ...
    FROM table_references
    [WHERE where_condition]

   DELETE [LOW_PRIORITY] [QUICK] [IGNORE]
    FROM tbl_name[.*] [, tbl_name[.*]] ...
    USING table_references
    [WHERE where_condition]
 * https://dev.mysql.com/doc/refman/8.0/en/delete.html
 *
 *
 * DELETE [ hint ]
   [ FROM ] { dml_table_expression_clause | ONLY (dml_table_expression_clause)} [ t_alias ]
     [ where_clause ]
     [ returning_clause ]
     [error_logging_clause];
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DELETE.html#GUID-156845A5-B626-412B-9F95-8869B988ABD7
 *
 */
type SQLDeleteStatement struct {
	*AbstractSQLStatement

	LowPriority bool
	Quick       bool
	Ignore      bool

	// MySQL: Multiple-OnTable
	tables []expr.ISQLName

	From           bool
	tableReference select_.ISQLTableReference

	// MySQL: USING table_references
	usingTableReference select_.ISQLTableReference

	whereClause   *select_.SQLWhereClause
	orderByClause *select_.SQLOrderByClause
	limitClause   select_.ISQLLimitClause

	// Oracle: Returning Clause
	returningClause select_.ISQLReturningClause
}

func NewDeleteStatement(dbType db.Type) *SQLDeleteStatement {
	x := new(SQLDeleteStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	x.From = true
	return x
}

func (x *SQLDeleteStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLDeleteStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLDeleteStatement) ObjectType() db.SQLObjectType {
	return db.DELETE
}

func (x *SQLDeleteStatement) Tables() []expr.ISQLName {
	return x.tables
}
func (x *SQLDeleteStatement) Table(i int) expr.ISQLName {
	return x.tables[i]
}
func (x *SQLDeleteStatement) AddTable(table expr.ISQLName) {
	if table == nil {
		return
	}
	table.SetParent(x)
	x.tables = append(x.tables, table)
}

func (x *SQLDeleteStatement) TableReference() select_.ISQLTableReference {
	return x.tableReference
}
func (x *SQLDeleteStatement) SetTableReference(tableReference select_.ISQLTableReference) {
	tableReference.SetParent(x)
	x.tableReference = tableReference
}

func (x *SQLDeleteStatement) UsingTableReference() select_.ISQLTableReference {
	return x.usingTableReference
}
func (x *SQLDeleteStatement) SetUsingTableReference(usingTableReference select_.ISQLTableReference) {
	if usingTableReference == nil {
		return
	}
	usingTableReference.SetParent(x)
	x.usingTableReference = usingTableReference
}

func (x *SQLDeleteStatement) WhereClause() *select_.SQLWhereClause {
	return x.whereClause
}

func (x *SQLDeleteStatement) SetWhereClause(whereClause *select_.SQLWhereClause) {
	if ast.IsNil(whereClause) {
		return
	}
	whereClause.SetParent(x)
	x.whereClause = whereClause
}
func (x *SQLDeleteStatement) OrderByClause() *select_.SQLOrderByClause {
	return x.orderByClause
}

func (x *SQLDeleteStatement) SetOrderByClause(orderByClause *select_.SQLOrderByClause) {
	if orderByClause == nil {
		return
	}
	orderByClause.SetParent(orderByClause)
	x.orderByClause = orderByClause
}

func (x *SQLDeleteStatement) LimitClause() select_.ISQLLimitClause {
	return x.limitClause
}

func (x *SQLDeleteStatement) SetLimitClause(limitClause select_.ISQLLimitClause) {
	if limitClause == nil {
		return
	}
	limitClause.SetParent(x)
	x.limitClause = limitClause
}

func (x *SQLDeleteStatement) ReturningClause() select_.ISQLReturningClause {
	return x.returningClause
}

func (x *SQLDeleteStatement) SetReturningClause(returningClause select_.ISQLReturningClause) {
	if returningClause == nil {
		return
	}
	returningClause.SetParent(x)
	x.returningClause = returningClause
}

/**
 *
 * EXPLAIN tbl_name [col_name | wild]
 *   EXPLAIN [explain_type] explainable_stmt
 explain_type: {
    EXTENDED
  | PARTITIONS
  | FORMAT = format_name
}
 * {EXPLAIN | DESCRIBE | DESC} ANALYZE [FORMAT = TREE] select_statement
 * https://dev.mysql.com/doc/refman/8.0/en/explain.html
 *
 * EXPLAIN PLAN [ SET STATEMENT_ID = string ]
   [ INTO [ schema. ] table [ @ dblink ] ] FOR statement ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/EXPLAIN-PLAN.html#GUID-FD540872-4ED3-4936-96A2-362539931BA0
 */
type abstractSQLExplainStatement struct {
	*AbstractSQLStatement
	// MySQL
	Analyze     bool
	explainType expr.ISQLExpr
	table       expr.ISQLName
	column      expr.ISQLExpr

	connectionId expr.ISQLExpr

	stmt ISQLStatement
}

func newAbstractSQLExplainStatement(dbType db.Type) *abstractSQLExplainStatement {
	x := new(abstractSQLExplainStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *abstractSQLExplainStatement) ExplainType() expr.ISQLExpr {
	return x.explainType
}
func (x *abstractSQLExplainStatement) SetExplainType(explainType expr.ISQLExpr) {
	if explainType == nil {
		return
	}
	explainType.SetParent(x)
	x.explainType = explainType
}
func (x *abstractSQLExplainStatement) Table() expr.ISQLName {
	return x.table
}
func (x *abstractSQLExplainStatement) SetTable(table expr.ISQLName) {
	if table == nil {
		return
	}
	table.SetParent(x)
	x.explainType = table
}
func (x *abstractSQLExplainStatement) Column() expr.ISQLExpr {
	return x.column
}
func (x *abstractSQLExplainStatement) SetColumn(column expr.ISQLExpr) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.column = column
}
func (x *abstractSQLExplainStatement) ConnectionId() expr.ISQLExpr {
	return x.connectionId
}
func (x *abstractSQLExplainStatement) SetConnectionId(connectionId expr.ISQLExpr) {
	if connectionId == nil {
		return
	}
	connectionId.SetParent(x)
	x.connectionId = connectionId
}

func (x *abstractSQLExplainStatement) Stmt() ISQLStatement {
	return x.stmt
}
func (x *abstractSQLExplainStatement) SetStmt(stmt ISQLStatement) {
	if stmt == nil {
		return
	}
	stmt.SetParent(x)
	x.stmt = stmt
}

/**
 * EXPLAIN PLAN
   [ SET STATEMENT_ID = string ]
   [ INTO [ schema. ] table [ @ dblink ] ]
FOR statement ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/EXPLAIN-PLAN.html#GUID-FD540872-4ED3-4936-96A2-362539931BA0
 */
type SQLExplainStatement struct {
	*abstractSQLExplainStatement

	// Oracle
	setStatementIdValueExpr expr.ISQLExpr
	intoTable               expr.ISQLName
}

func NewExplainStatement(dbType db.Type) *SQLExplainStatement {
	x := new(SQLExplainStatement)
	x.abstractSQLExplainStatement = newAbstractSQLExplainStatement(dbType)
	return x
}
func (x *SQLExplainStatement) SetStatementIdValueExpr() expr.ISQLExpr {
	return x.setStatementIdValueExpr
}
func (x *SQLExplainStatement) SetSetStatementIdValueExpr(setStatementIdValueExpr expr.ISQLExpr) {
	if setStatementIdValueExpr == nil {
		return
	}
	setStatementIdValueExpr.SetParent(x)
	x.setStatementIdValueExpr = setStatementIdValueExpr
}
func (x *SQLExplainStatement) IntoTable() expr.ISQLName {
	return x.intoTable
}
func (x *SQLExplainStatement) SetIntoTable(intoTable expr.ISQLName) {
	if intoTable == nil {
		return
	}
	intoTable.SetParent(x)
	x.intoTable = intoTable
}

/**
 * DESCRIBE
 * https://dev.mysql.com/doc/refman/8.0/en/explain.html
 */
type SQLDescribeStatement struct {
	*abstractSQLExplainStatement
}

func NewDescribeStatement(dbType db.Type) *SQLDescribeStatement {
	x := new(SQLDescribeStatement)
	x.abstractSQLExplainStatement = newAbstractSQLExplainStatement(dbType)
	return x
}

/**
 * DESC
 * https://dev.mysql.com/doc/refman/8.0/en/explain.html
 */
type SQLDescStatement struct {
	*abstractSQLExplainStatement
}

func NewDescStatement(dbType db.Type) *SQLDescStatement {
	x := new(SQLDescStatement)
	x.abstractSQLExplainStatement = newAbstractSQLExplainStatement(dbType)
	return x
}

/**
 * HELP ''
 * https://dev.mysql.com/doc/refman/8.0/en/use.html
 */
type SQLHelpStatement struct {
	*AbstractSQLStatement
	name expr.ISQLExpr
}

func NewHelpStatement(dbType db.Type) *SQLHelpStatement {
	x := new(SQLHelpStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLHelpStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLHelpStatement) SetName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 * USE db_name
 * https://dev.mysql.com/doc/refman/8.0/en/use.html
 */
type SQLUseStatement struct {
	*AbstractSQLStatement
	name expr.ISQLExpr
}

func NewUseStatement(dbType db.Type) *SQLUseStatement {
	x := new(SQLUseStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLUseStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLUseStatement) SetName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 * assignment_statement_target := expression ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/assignment-statement.html#GUID-4C3BEFDF-3FFA-4E9D-96D0-4C5E13E08643
 */
type SQLAssignmentStatement struct {
	*AbstractSQLStatement
	target expr.ISQLExpr
	expr   expr.ISQLExpr
}

func NewAssignmentStatement(dbType db.Type) *SQLAssignmentStatement {
	x := new(SQLAssignmentStatement)
	x.AbstractSQLStatement = NewAbstractSQLStatementWithDBType(dbType)
	return x
}

// EDITIONABLE | NONEDITIONABLE | COMPILE
/**
 * EDITIONABLE
 */
type SQLEditionAbleExpr struct {
	*expr.AbstractSQLExpr
}

func NewEditionAbleExpr() *SQLEditionAbleExpr {
	x := new(SQLEditionAbleExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * NONEDITIONABLE
 */
type SQLNonEditionAbleExpr struct {
	*expr.AbstractSQLExpr
}

func NewNonEditionAbleExpr() *SQLNonEditionAbleExpr {
	x := new(SQLNonEditionAbleExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * COMPILE
 */
type SQLCompileExpr struct {
	*expr.AbstractSQLExpr
}

func NewCompileExpr() *SQLCompileExpr {
	x := new(SQLCompileExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * READ ONLY
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLReadOnlyExpr struct {
	*expr.AbstractSQLExpr
}

func NewReadOnlyExpr() *SQLReadOnlyExpr {
	x := new(SQLReadOnlyExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * READ WRITE
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLReadWriteExpr struct {
	*expr.AbstractSQLExpr
}

func NewReadWriteExpr() *SQLReadWriteExpr {
	x := new(SQLReadWriteExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * [ <parameter mode> ] [ <SQL parameter name> ] <parameter type> [ RESULT ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#SQL%20parameter%20declaration
 *
 *
 */
type SQLParameterMode string

const (
	IN    SQLParameterMode = "IN"
	OUT                    = "OUT"
	INOUT                  = "INOUT"
)

type SQLParameterDeclaration struct {
	*expr.AbstractSQLExpr
	ParameterMode SQLParameterMode
	name          expr.ISQLName
	dataType      datatype.ISQLDataType
	Result        bool
}

func NewParameterDeclaration() *SQLParameterDeclaration {
	x := new(SQLParameterDeclaration)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLParameterDeclaration) Name() expr.ISQLName {
	return x.name
}
func (x *SQLParameterDeclaration) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLParameterDeclaration) DataType() datatype.ISQLDataType {
	return x.dataType
}
func (x *SQLParameterDeclaration) SetDataType(dataType datatype.ISQLDataType) {
	if dataType == nil {
		return
	}
	dataType.SetParent(x)
	x.dataType = dataType
}
