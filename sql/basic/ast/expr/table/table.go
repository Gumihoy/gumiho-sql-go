package table

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
)

// ---------------------------------------------- Table Element Start ----------------------------------------------

/**
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20contents%20source
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20element%20list
 */
type ISQLTableElement interface {
	expr.ISQLExpr
}

/**
 * <column name> [ <data type> | <domain name> ] [ <reference scope check> ] [ <default clause> | <identity column specification> | <generation clause> ] [ <column constraint definition> ... ] [ <collate clause> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20definition
 *
 *  <column name> data_type [NOT NULL | NULL] [DEFAULT {literal | (expr)} ]
      [VISIBLE | INVISIBLE]
      [AUTO_INCREMENT] [UNIQUE [KEY]] [[PRIMARY] KEY]
      [COMMENT 'string']
      [COLLATE collation_name]
      [COLUMN_FORMAT {FIXED | DYNAMIC | DEFAULT}]
      [ENGINE_ATTRIBUTE [=] 'string']
      [SECONDARY_ENGINE_ATTRIBUTE [=] 'string']
      [STORAGE {DISK | MEMORY}]
      [reference_definition]
      [check_constraint_definition]
| <column name> data_type
      [COLLATE collation_name]
      [GENERATED ALWAYS] AS (expr)
      [VIRTUAL | STORED] [NOT NULL | NULL]
      [VISIBLE | INVISIBLE]
      [UNIQUE [KEY]] [[PRIMARY] KEY]
      [COMMENT 'string']
      [reference_definition]
      [check_constraint_definition]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
 */
type SQLTableColumn struct {
	*expr.AbstractSQLExpr
	name     expr.ISQLIdentifier
	dataType datatype.ISQLDataType

	// option or column constraint
	options []expr.ISQLExpr
}

func NewColumn() *SQLTableColumn {
	x := new(SQLTableColumn)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.options = make([]expr.ISQLExpr, 0, 10)
	return x
}

func NewColumnWitNameAndDataType(name expr.ISQLIdentifier, dataType datatype.ISQLDataType) *SQLTableColumn {
	x := new(SQLTableColumn)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetName(name)
	x.SetDataType(dataType)
	x.options = make([]expr.ISQLExpr, 0, 10)
	return x
}

func (x *SQLTableColumn) Name() expr.ISQLIdentifier {
	return x.name
}

func (x *SQLTableColumn) SetName(name expr.ISQLIdentifier) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLTableColumn) DataType() datatype.ISQLDataType {
	return x.dataType
}
func (x *SQLTableColumn) SetDataType(datatype datatype.ISQLDataType) {
	datatype.SetParent(x)
	x.dataType = datatype
}

func (x *SQLTableColumn) Options() []expr.ISQLExpr {
	return x.options
}
func (x *SQLTableColumn) Option(i int) expr.ISQLExpr {
	return x.options[i]
}
func (x *SQLTableColumn) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}

func (x *SQLTableColumn) DefaultClause() *SQLDefaultClause {
	for _, cc := range x.options {
		switch cc.(type) {
		case *SQLDefaultClause:
			return cc.(*SQLDefaultClause)
		}
	}
	return nil
}

// [ <constraint name definition> ] <table constraint> [ <constraint characteristics> ]
// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20constraint%20definition
type ISQLTableConstraint interface {
	ISQLTableElement
	Name() expr.ISQLName
	SetName(name expr.ISQLName)
}

/**
 * CONSTRAINT name
 */
type AbstractSQLTableConstraint struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewAbstractSQLTableConstraint() *AbstractSQLTableConstraint {
	x := new(AbstractSQLTableConstraint)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *AbstractSQLTableConstraint) Name() expr.ISQLName {
	return x.name
}

func (x *AbstractSQLTableConstraint) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 * [CONSTRAINT name] PRIMARY KEY ()
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#unique%20constraint%20definition
 */
type SQLPrimaryKeyTableConstraint struct {
	*AbstractSQLTableConstraint
	columns []expr.ISQLExpr
}

func NewPrimaryKeyTableConstraint() *SQLPrimaryKeyTableConstraint {
	x := new(SQLPrimaryKeyTableConstraint)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}
func (x *SQLPrimaryKeyTableConstraint) Columns() []expr.ISQLExpr {
	return x.columns
}
func (x *SQLPrimaryKeyTableConstraint) Column(i int) expr.ISQLExpr {
	return x.columns[i]
}
func (x *SQLPrimaryKeyTableConstraint) AddColumn(column expr.ISQLExpr) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}

/**
 * [CONSTRAINT name] UNIQUE ()
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#unique%20constraint%20definition
 *
 * [CONSTRAINT [symbol]] UNIQUE [INDEX | KEY]
      [index_name] [index_type] (key_part,...)
      [index_option] ...
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLUniqueTableConstraint struct {
	*AbstractSQLTableConstraint
	columns []expr.ISQLExpr
}

func NewUniqueTableConstraint() *SQLUniqueTableConstraint {
	x := new(SQLUniqueTableConstraint)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}
func (x *SQLUniqueTableConstraint) Columns() []expr.ISQLExpr {
	return x.columns
}
func (x *SQLUniqueTableConstraint) Column(i int) expr.ISQLExpr {
	return x.columns[i]
}
func (x *SQLUniqueTableConstraint) AddColumn(column expr.ISQLExpr) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}

type SQLUniqueIndexTableConstraint struct {
	*SQLUniqueTableConstraint
}

func NewUniqueIndexTableConstraint() *SQLUniqueIndexTableConstraint {
	x := new(SQLUniqueIndexTableConstraint)
	x.SQLUniqueTableConstraint = NewUniqueTableConstraint()
	return x
}

type SQLUniqueKeyTableConstraint struct {
	*SQLUniqueTableConstraint
}

func NewUniqueKeyTableConstraint() *SQLUniqueKeyTableConstraint {
	x := new(SQLUniqueKeyTableConstraint)
	x.SQLUniqueTableConstraint = NewUniqueTableConstraint()
	return x
}

/**
 * [CONSTRAINT name] FOREIGN KEY [index_name] <left paren> <referencing columns> <right paren> REFERENCES <referenced table and columns> [ MATCH <match type> ] [ <referential triggered action> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#referential%20constraint%20definition
 *
 *
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLForeignKeyTableConstraint struct {
	*AbstractSQLTableConstraint

	referencingIndex   expr.ISQLName
	referencingColumns []expr.ISQLExpr

	referencedTable   expr.ISQLName
	referencedColumns []expr.ISQLExpr
}

func NewForeignKeyTableConstraint() *SQLForeignKeyTableConstraint {
	x := new(SQLForeignKeyTableConstraint)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}
func (x *SQLForeignKeyTableConstraint) ReferencingIndex(i int) expr.ISQLName {
	return x.referencingIndex
}
func (x *SQLForeignKeyTableConstraint) SetReferencingIndex(referencingIndex expr.ISQLName) {
	if referencingIndex == nil {
		return
	}
	referencingIndex.SetParent(x)
	x.referencingIndex = referencingIndex
}
func (x *SQLForeignKeyTableConstraint) ReferencingColumns() []expr.ISQLExpr {
	return x.referencingColumns
}
func (x *SQLForeignKeyTableConstraint) ReferencingColumn(i int) expr.ISQLExpr {
	return x.referencingColumns[i]
}
func (x *SQLForeignKeyTableConstraint) AddReferencingColumn(referencingColumn expr.ISQLExpr) {
	if referencingColumn == nil {
		return
	}
	referencingColumn.SetParent(x)
	x.referencingColumns = append(x.referencingColumns, referencingColumn)
}
func (x *SQLForeignKeyTableConstraint) ReferencedTable(i int) expr.ISQLName {
	return x.referencedTable
}
func (x *SQLForeignKeyTableConstraint) SetReferencedTable(referencedTable expr.ISQLName) {
	if referencedTable == nil {
		return
	}
	referencedTable.SetParent(x)
	x.referencedTable = referencedTable
}
func (x *SQLForeignKeyTableConstraint) ReferencedColumns() []expr.ISQLExpr {
	return x.referencedColumns
}
func (x *SQLForeignKeyTableConstraint) ReferencedColumn(i int) expr.ISQLExpr {
	return x.referencedColumns[i]
}
func (x *SQLForeignKeyTableConstraint) AddReferencedColumn(referencedColumn expr.ISQLExpr) {
	if referencedColumn == nil {
		return
	}
	referencedColumn.SetParent(x)
	x.referencedColumns = append(x.referencedColumns, referencedColumn)
}

type SQLMatchType string

const (
	FULL    SQLMatchType = "FULL"
	PARTIAL SQLMatchType = "PARTIAL"
	SIMPLE  SQLMatchType = "SIMPLE"
)

/**
 * [CONSTRAINT name] CHECK ( condition )
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#check%20constraint%20definition
 *
 * [CONSTRAINT [symbol]] CHECK (expr) [[NOT] ENFORCED]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLCheckTableConstraint struct {
	*AbstractSQLTableConstraint
	condition expr.ISQLExpr
}

func NewCheckTableConstraint() *SQLCheckTableConstraint {
	x := new(SQLCheckTableConstraint)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}
func NewCheckTableConstraintWithCondition(condition expr.ISQLExpr) *SQLCheckTableConstraint {
	x := new(SQLCheckTableConstraint)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	x.SetCondition(condition)
	return x
}

func (x *SQLCheckTableConstraint) Condition() expr.ISQLExpr {
	return x.condition
}

func (x *SQLCheckTableConstraint) SetCondition(condition expr.ISQLExpr) {
	condition.SetParent(x)
	x.condition = condition
}

/**
 * INDEX [index_name] [index_type] (key_part,...)
      [index_option] ...
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLIndexTableElement struct {
	*AbstractSQLTableConstraint
}

func NewIndexTableElement() *SQLIndexTableElement {
	x := new(SQLIndexTableElement)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}

/**
*  KEY [index_name] [index_type] (key_part,...)
	 [index_option] ...
* https://dev.mysql.com/doc/refman/8.0/en/create-table.html
*/
type SQLKeyTableElement struct {
	*AbstractSQLTableConstraint
}

func NewKeyTableElement() *SQLKeyTableElement {
	x := new(SQLKeyTableElement)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}

/**
*  FULLTEXT [INDEX | KEY] [index_name] (key_part,...)
      [index_option] ...
* https://dev.mysql.com/doc/refman/8.0/en/create-table.html
*/
type SQLFulltextIndexTableElement struct {
	*AbstractSQLTableConstraint
}

func NewFulltextIndexTableElement() *SQLFulltextIndexTableElement {
	x := new(SQLFulltextIndexTableElement)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}

type SQLFulltextKeyTableElement struct {
	*AbstractSQLTableConstraint
}

func NewFulltextKeyTableElement() *SQLFulltextKeyTableElement {
	x := new(SQLFulltextKeyTableElement)
	x.AbstractSQLTableConstraint = NewAbstractSQLTableConstraint()
	return x
}

/**
*  SPATIAL [INDEX | KEY] [index_name] (key_part,...)
      [index_option] ...
* https://dev.mysql.com/doc/refman/8.0/en/create-table.html
*/
type SQLSpatialIndexTableElement struct {
	*AbstractSQLTableConstraint
}
type SQLSpatialKeyTableElement struct {
	*AbstractSQLTableConstraint
}

/**
 *  LIKE <table name> [ <like options> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#like%20clause
 *
 * LIKE old_tbl_name
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLTableLikeClause struct {
	*expr.AbstractSQLExpr

	name expr.ISQLName
}

func NewTableLikeClause() *SQLTableLikeClause {
	x := new(SQLTableLikeClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *SQLTableLikeClause) Name() expr.ISQLName {
	return x.name
}

func (x *SQLTableLikeClause) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 *
 */
type SQLSelfReferencingColumn struct {
	expr.AbstractSQLExpr
}
type SQLColumnOption struct {
	expr.AbstractSQLExpr
}

// ----------- Default Clause
/**
 * DEFAULT {literal | (value)}
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLDefaultClause struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}

func NewDefaultClause() *SQLDefaultClause {
	x := new(SQLDefaultClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *SQLDefaultClause) Value() expr.ISQLExpr {
	return x.value
}

func (x *SQLDefaultClause) SetValue(value expr.ISQLExpr) {
	value.SetParent(x)
	x.value = value
}

/**
 * GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ <left paren> <common sequence generator options> <right paren> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20definition
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#identity%20column%20specification
 *
 * identity_clause: GENERATED [ ALWAYS | BY DEFAULT [ ON NULL ] ] AS IDENTITY [ ( identity_options ) ]
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
 */
type SQLIdentityGeneratedKind string

const (
	ALWAYS     SQLIdentityGeneratedKind = "ALWAYS"
	BY_DEFAULT                          = "BY DEFAULT"
)

type SQLIdentityColumnClause struct {
	*expr.AbstractSQLExpr
	IdentityGeneratedKind SQLIdentityGeneratedKind
	options               []expr.ISQLExpr
}

func NewIdentityColumnClause() *SQLIdentityColumnClause {
	x := new(SQLIdentityColumnClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIdentityColumnClause) Options() []expr.ISQLExpr {
	return x.options
}
func (x *SQLIdentityColumnClause) Option(i int) expr.ISQLExpr {
	return x.options[i]
}
func (x *SQLIdentityColumnClause) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}

/**
 * GENERATED ALWAYS AS <left paren> <value expression> <right paren>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20definition
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#generation%20clause
 *
 * [ GENERATED ALWAYS ] AS (column_expression)
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
 */
type SQLGenerationClause struct {
	*expr.AbstractSQLExpr
	GeneratedAlways bool
	expr            expr.ISQLExpr
}

func NewGenerationClause() *SQLGenerationClause {
	x := new(SQLGenerationClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

// ---------------------------------------------- Column Constraint ----------------------------------------------

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20constraint%20definition
type ISQLColumnConstraint interface {
	expr.ISQLExpr
	// Name() expr.ISQLName
	// SetName(name expr.ISQLName)
}

type AbstractSQLColumnConstraint struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewAbstractSQLColumnConstraint() *AbstractSQLColumnConstraint {
	x := new(AbstractSQLColumnConstraint)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#constraint%20characteristics
type SQLColumnConstraintCharacteristics struct {
	expr.AbstractSQLExpr
}

func (x *AbstractSQLColumnConstraint) Name() expr.ISQLName {
	return x.name
}

func (x *AbstractSQLColumnConstraint) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 * [ CONSTRAINT <constraint name>] PRIMARY KEY <constraint check time> [ [ NOT ] DEFERRABLE ] | [ NOT ] DEFERRABLE [ <constraint check time> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20constraint%20definition
 *
 * PRIMARY KEY
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLPrimaryKeyColumnConstraint struct {
	*AbstractSQLColumnConstraint
}

func NewPrimaryKeyColumnConstraint() *SQLPrimaryKeyColumnConstraint {
	x := new(SQLPrimaryKeyColumnConstraint)
	x.AbstractSQLColumnConstraint = NewAbstractSQLColumnConstraint()
	return x
}

/**
 * [CONSTRAINT <constraint name>] UNIQUE <constraint check time> [ [ NOT ] DEFERRABLE ] | [ NOT ] DEFERRABLE [ <constraint check time> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#unique%20constraint%20definition
 *
 * UNIQUE [KEY]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLUniqueColumnConstraint struct {
	*AbstractSQLColumnConstraint
	Key bool
}

func NewUniqueColumnConstraint() *SQLUniqueColumnConstraint {
	x := new(SQLUniqueColumnConstraint)
	x.AbstractSQLColumnConstraint = NewAbstractSQLColumnConstraint()
	return x
}

/**
 * NOT NULL
 */
type SQLNotNullColumnConstraint struct {
	*AbstractSQLColumnConstraint
}

func NewNotNullColumnConstraint() *SQLNotNullColumnConstraint {
	x := new(SQLNotNullColumnConstraint)
	x.AbstractSQLColumnConstraint = NewAbstractSQLColumnConstraint()
	return x
}

/**
 * NULL
 */
type SQLNullColumnConstraint struct {
	*AbstractSQLColumnConstraint
}

func NewNullColumnConstraint() *SQLNullColumnConstraint {
	x := new(SQLNullColumnConstraint)
	x.AbstractSQLColumnConstraint = NewAbstractSQLColumnConstraint()
	return x
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#references%20specification
type SQLReferencesColumnConstraint struct {
	AbstractSQLColumnConstraint
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#check%20constraint%20definition
type SQLCheckColumnConstraint struct {
	*AbstractSQLColumnConstraint
	condition expr.ISQLExpr
}

func NewCheckColumnConstraint() *SQLCheckColumnConstraint {
	x := new(SQLCheckColumnConstraint)
	x.AbstractSQLColumnConstraint = NewAbstractSQLColumnConstraint()
	return x
}

/**
 * KEY
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLKeyColumnConstraint struct {
	*expr.AbstractSQLExpr
}

func NewKeyColumnConstraint() *SQLKeyColumnConstraint {
	x := new(SQLKeyColumnConstraint)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * AUTO_INCREMENT
 *
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLAutoIncrementExpr struct {
	*expr.AbstractSQLExpr
}

func NewAutoIncrementExpr() *SQLAutoIncrementExpr {
	x := new(SQLAutoIncrementExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * VISIBLE
 *
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLVisibleExpr struct {
	*expr.AbstractSQLExpr
	comment expr.ISQLExpr
}

func NewVisibleExpr() *SQLVisibleExpr {
	x := new(SQLVisibleExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * INVISIBLE
 *
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLInvisibleExpr struct {
	*expr.AbstractSQLExpr
	comment expr.ISQLExpr
}

func NewInvisibleExpr() *SQLInvisibleExpr {
	x := new(SQLInvisibleExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * COMMENT 'string'
 *
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLCommentExpr struct {
	*expr.AbstractSQLExpr
	comment expr.ISQLExpr
}

func NewCommentExpr() *SQLCommentExpr {
	x := new(SQLCommentExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLCommentExpr) Comment() expr.ISQLExpr {
	return x.comment
}

func (x *SQLCommentExpr) SetComment(comment expr.ISQLExpr) {
	comment.SetParent(x)
	x.comment = comment
}

// ---------------------------------------------- Table Element End ----------------------------------------------

// ---------------------------------------------- MySQL Table Option Start ----------------------------------------------
/**
 * [DEFAULT] CHARACTER SET [=] charset_name
 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLCharacterSetAssignExpr struct {
	*expr.SQLAssignExpr
	Default bool
}

func NewCharacterSetAssignExpr() *SQLCharacterSetAssignExpr {
	x := new(SQLCharacterSetAssignExpr)
	x.SQLAssignExpr = expr.NewAssignExpr()
	return x
}

/**
 * [DEFAULT] CHARSET [=] charset_name
 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLCharsetAssignExpr struct {
	*expr.SQLAssignExpr
	Default bool
}

func NewCharsetAssignExpr() *SQLCharsetAssignExpr {
	x := new(SQLCharsetAssignExpr)
	x.SQLAssignExpr = expr.NewAssignExpr()
	return x
}

/**
 * [DEFAULT] COLLATE [=] collation_name
 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLCollateAssignExpr struct {
	*expr.SQLAssignExpr
	Default bool
}

func NewCollateAssignExpr() *SQLCollateAssignExpr {
	x := new(SQLCollateAssignExpr)
	x.SQLAssignExpr = expr.NewAssignExpr()
	return x
}

/**
 * DATA DIRECTORY [=] 'absolute path to directory'
 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLDataDirectoryAssignExpr struct {
	*expr.SQLAssignExpr
	Default bool
}

func NewDataDirectoryAssignExpr() *SQLDataDirectoryAssignExpr {
	x := new(SQLDataDirectoryAssignExpr)
	x.SQLAssignExpr = expr.NewAssignExpr()
	return x
}

/**
 * INDEX DIRECTORY [=] 'absolute path to directory'
 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLIndexDirectoryAssignExpr struct {
	*expr.SQLAssignExpr
	Default bool
}

func NewIndexDirectoryAssignExpr() *SQLIndexDirectoryAssignExpr {
	x := new(SQLIndexDirectoryAssignExpr)
	x.SQLAssignExpr = expr.NewAssignExpr()
	return x
}

/**
 * TABLESPACE tablespace_name [STORAGE {DISK | MEMORY}]
 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLTablespaceAssignExpr struct {
	*expr.SQLAssignExpr
}

func NewTablespaceAssignExpr() *SQLTablespaceAssignExpr {
	x := new(SQLTablespaceAssignExpr)
	x.SQLAssignExpr = expr.NewAssignExpr()
	return x
}

// ---------------------------------------------- MySQL OnTable Option End ----------------------------------------------

// ---------------------------------------------- OnTable Partitioning Start ----------------------------------------------

type ISQLTablePartitioning interface {
	expr.ISQLExpr
}

/**
 * PARTITION BY
        { [LINEAR] HASH(value)
        	| [LINEAR] KEY [ALGORITHM={1 | 2}] (column_list)
        	| RANGE{(value) | COLUMNS(column_list)}
        	| LIST{(value) | COLUMNS(column_list)}
		}
    [PARTITIONS num]
    [SUBPARTITION BY
        { [LINEAR] HASH(value)
        | [LINEAR] KEY [ALGORITHM={1 | 2}] (column_list) }
      [SUBPARTITIONS num]
    ]
    [(partition_definition [, partition_definition] ...)]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 *
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
 */
type ISQLPartitionBy interface {
	ISQLTablePartitioning

	Columns() []expr.ISQLExpr
	Column(i int) expr.ISQLExpr
	AddColumn(column expr.ISQLExpr)

	PartitionsNum() expr.ISQLExpr
	SetPartitionsNum(partitionsNum expr.ISQLExpr)

	SubPartitionBy() ISQLSubPartitionBy
	SetSubPartitionBy(subPartitionBy ISQLSubPartitionBy)

	PartitionDefinitions() []*SQLPartitionDefinition
	PartitionDefinition(i int) *SQLPartitionDefinition
	AddPartitionDefinition(partitionDefinition *SQLPartitionDefinition)
}

type AbstractSQLPartitionBy struct {
	*expr.AbstractSQLExpr

	columns       []expr.ISQLExpr
	partitionsNum expr.ISQLExpr

	subPartitionBy       ISQLSubPartitionBy
	partitionDefinitions []*SQLPartitionDefinition
}

func NewAbstractSQLPartitionBy() *AbstractSQLPartitionBy {
	x := new(AbstractSQLPartitionBy)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.columns = make([]expr.ISQLExpr, 0, 10)
	return x
}

func (x *AbstractSQLPartitionBy) Columns() []expr.ISQLExpr {
	return x.columns
}

func (x *AbstractSQLPartitionBy) Column(i int) expr.ISQLExpr {
	return x.columns[i]
}

func (x *AbstractSQLPartitionBy) AddColumn(column expr.ISQLExpr) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}

func (x *AbstractSQLPartitionBy) PartitionsNum() expr.ISQLExpr {
	return x.partitionsNum
}
func (x *AbstractSQLPartitionBy) SetPartitionsNum(partitionsNum expr.ISQLExpr) {
	if partitionsNum == nil {
		return
	}
	partitionsNum.SetParent(x)
	x.partitionsNum = partitionsNum
}

func (x *AbstractSQLPartitionBy) SubPartitionBy() ISQLSubPartitionBy {
	return x.subPartitionBy
}
func (x *AbstractSQLPartitionBy) SetSubPartitionBy(subPartitionBy ISQLSubPartitionBy) {
	x.subPartitionBy = subPartitionBy
}

func (x *AbstractSQLPartitionBy) PartitionDefinitions() []*SQLPartitionDefinition {
	return x.partitionDefinitions
}
func (x *AbstractSQLPartitionBy) PartitionDefinition(i int) *SQLPartitionDefinition {
	return x.partitionDefinitions[i]
}
func (x *AbstractSQLPartitionBy) AddPartitionDefinition(partitionDefinition *SQLPartitionDefinition) {
	if partitionDefinition == nil {
		return
	}
	partitionDefinition.SetParent(x)
	x.partitionDefinitions = append(x.partitionDefinitions, partitionDefinition)
}

/**
 * PARTITION BY [LINEAR] HASH(value)
	[PARTITIONS num]
    [SUBPARTITION BY
        { [LINEAR] HASH(expr)
        | [LINEAR] KEY [ALGORITHM={1 | 2}] (column_list) }
      [SUBPARTITIONS num]
    ]
    [(partition_definition [, partition_definition] ...)]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLPartitionByHash struct {
	*AbstractSQLPartitionBy
	Linear bool
}

func NewPartitionByHash() *SQLPartitionByHash {
	x := new(SQLPartitionByHash)
	x.AbstractSQLPartitionBy = NewAbstractSQLPartitionBy()
	return x
}

/**
 * PARTITION BY [LINEAR] KEY [ALGORITHM = {1 | 2}](value)
	[PARTITIONS num]
    [SUBPARTITION BY
        { [LINEAR] HASH(expr)
        | [LINEAR] KEY [ALGORITHM={1 | 2}] (column_list) }
      [SUBPARTITIONS num]
    ]
    [(partition_definition [, partition_definition] ...)]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLPartitionByKey struct {
	*AbstractSQLPartitionBy
	Linear         bool
	algorithmValue expr.ISQLExpr
}

func NewPartitionByKey() *SQLPartitionByKey {
	x := new(SQLPartitionByKey)
	x.AbstractSQLPartitionBy = NewAbstractSQLPartitionBy()
	return x
}
func (x *SQLPartitionByKey) AlgorithmValue() expr.ISQLExpr {
	return x.algorithmValue
}
func (x *SQLPartitionByKey) SetAlgorithmValue(algorithmValue expr.ISQLExpr) {
	if algorithmValue == nil {
		return
	}
	algorithmValue.SetParent(x)
	x.algorithmValue = algorithmValue
}

/**
 *  PARTITION BY RANGE [COLUMNS] (expr)   [PARTITIONS num]
    [SUBPARTITION BY
        { [LINEAR] HASH(expr)
        | [LINEAR] KEY [ALGORITHM={1 | 2}] (column_list) }
      [SUBPARTITIONS num]
    ]
    [(partition_definition [, partition_definition] ...)]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 *
 * range_partitions:
 * PARTITION BY RANGE(column [,column]+)
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
 */
type SQLPartitionByRange struct {
	*AbstractSQLPartitionBy
	HasColumns bool
}

func NewPartitionByRange() *SQLPartitionByRange {
	x := new(SQLPartitionByRange)
	x.AbstractSQLPartitionBy = NewAbstractSQLPartitionBy()
	return x
}

/**
 *  PARTITION BY LIST [COLUMNS] (expr)   [PARTITIONS num]
    [SUBPARTITION BY
        { [LINEAR] HASH(expr)
        | [LINEAR] KEY [ALGORITHM={1 | 2}] (column_list) }
      [SUBPARTITIONS num]
    ]
    [(partition_definition [, partition_definition] ...)]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 *
 * list_partitions:
 * PARTITION BY LIST(column [,column]+)
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
 */
type SQLPartitionByList struct {
	*AbstractSQLPartitionBy
	HasColumns bool
}

func NewPartitionByList() *SQLPartitionByList {
	x := new(SQLPartitionByList)
	x.AbstractSQLPartitionBy = NewAbstractSQLPartitionBy()
	return x
}

/*
 *  [SUBPARTITION BY
        { [LINEAR] HASH(value)
        | [LINEAR] KEY [ALGORITHM={1 | 2}] (column_list) }
      [SUBPARTITIONS num]
    ]
 */
type ISQLSubPartitionBy interface {
	expr.ISQLExpr

	Columns() []expr.ISQLExpr
	Column(i int) expr.ISQLExpr
	AddColumn(column expr.ISQLExpr)

	SubPartitionsNum() expr.ISQLExpr
	SetSubPartitionsNum(subPartitionsNum expr.ISQLExpr)
}

type AbstractSQLSubPartitionBy struct {
	*expr.AbstractSQLExpr
	columns          []expr.ISQLExpr
	subPartitionsNum expr.ISQLExpr
}

func NewAbstractSQLSubPartitionBy() *AbstractSQLSubPartitionBy {
	x := new(AbstractSQLSubPartitionBy)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.columns = make([]expr.ISQLExpr, 0, 10)
	return x
}

func (x *AbstractSQLSubPartitionBy) Columns() []expr.ISQLExpr {
	return x.columns
}

func (x *AbstractSQLSubPartitionBy) Column(index int) expr.ISQLExpr {
	return x.columns[index]
}

func (x *AbstractSQLSubPartitionBy) AddColumn(column expr.ISQLExpr) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}
func (x *AbstractSQLSubPartitionBy) SubPartitionsNum() expr.ISQLExpr {
	return x.subPartitionsNum
}
func (x *AbstractSQLSubPartitionBy) SetSubPartitionsNum(subPartitionsNum expr.ISQLExpr) {
	if subPartitionsNum == nil {
		return
	}
	subPartitionsNum.SetParent(x)
	x.subPartitionsNum = subPartitionsNum
}

type SQLSubPartitionByHash struct {
	*AbstractSQLSubPartitionBy
	Linear bool
}

func NewSubPartitionByHash() *SQLSubPartitionByHash {
	x := new(SQLSubPartitionByHash)
	x.AbstractSQLSubPartitionBy = NewAbstractSQLSubPartitionBy()
	return x
}

type SQLSubPartitionByKey struct {
	*AbstractSQLSubPartitionBy
	Linear         bool
	algorithmValue expr.ISQLExpr
}

func NewSubPartitionByKey() *SQLSubPartitionByKey {
	x := new(SQLSubPartitionByKey)
	x.AbstractSQLSubPartitionBy = NewAbstractSQLSubPartitionBy()
	return x
}
func (x *SQLSubPartitionByKey) AlgorithmValue() expr.ISQLExpr {
	return x.algorithmValue
}
func (x *SQLSubPartitionByKey) SetAlgorithmValue(algorithmValue expr.ISQLExpr) {
	if algorithmValue == nil {
		return
	}
	algorithmValue.SetParent(x)
	x.algorithmValue = algorithmValue
}

/**
 * PARTITION partition_name
        [VALUES
			{LESS THAN {(value | value_list) | MAXVALUE}
            |
            IN (value_list)}
		]
        [[STORAGE] ENGINE [=] engine_name]
        [COMMENT [=] 'string' ]
        [DATA DIRECTORY [=] 'data_dir']
        [INDEX DIRECTORY [=] 'index_dir']
        [MAX_ROWS [=] max_number_of_rows]
        [MIN_ROWS [=] min_number_of_rows]
        [TABLESPACE [=] tablespace_name]
        [(subpartition_definition [, subpartition_definition] ...)]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 *
 *
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
 */
type SQLPartitionDefinition struct {
	*expr.AbstractSQLExpr

	name expr.ISQLName

	values ISQLPartitionValues

	options []expr.ISQLExpr

	subpartitionDefinitions []*SQLSubPartitionDefinition
}

func NewPartitionDefinition() *SQLPartitionDefinition {
	x := new(SQLPartitionDefinition)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.subpartitionDefinitions = make([]*SQLSubPartitionDefinition, 0, 10)
	return x
}
func (x *SQLPartitionDefinition) Name() expr.ISQLName {
	return x.name
}

func (x *SQLPartitionDefinition) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLPartitionDefinition) Values() ISQLPartitionValues {
	return x.values
}

func (x *SQLPartitionDefinition) SetValues(values ISQLPartitionValues) {
	if values == nil {
		return
	}
	values.SetParent(x)
	x.values = values
}
func (x *SQLPartitionDefinition) Options() []expr.ISQLExpr {
	return x.options
}

func (x *SQLPartitionDefinition) Option(i int) expr.ISQLExpr {
	return x.subpartitionDefinitions[i]
}

func (x *SQLPartitionDefinition) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
func (x *SQLPartitionDefinition) SubpartitionDefinitions() []*SQLSubPartitionDefinition {
	return x.subpartitionDefinitions
}

func (x *SQLPartitionDefinition) SubpartitionDefinition(i int) *SQLSubPartitionDefinition {
	return x.subpartitionDefinitions[i]
}

func (x *SQLPartitionDefinition) AddSubpartitionDefinition(subPartitionDefinition *SQLSubPartitionDefinition) {
	if subPartitionDefinition == nil {
		return
	}
	subPartitionDefinition.SetParent(x)
	x.subpartitionDefinitions = append(x.subpartitionDefinitions, subPartitionDefinition)
}

/**
 * [VALUES
			{LESS THAN {(value | value_list) | MAXVALUE}
            |
            IN (value_list)}
		]
 */
type ISQLPartitionValues interface {
	expr.ISQLExpr
}
type AbstractSQLPartitionValues struct {
	*expr.AbstractSQLExpr
	values []expr.ISQLExpr
}

func NewAbstractSQLPartitionValues() *AbstractSQLPartitionValues {
	x := new(AbstractSQLPartitionValues)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *AbstractSQLPartitionValues) Values() []expr.ISQLExpr {
	return x.values
}

func (x *AbstractSQLPartitionValues) Value(i int) expr.ISQLExpr {
	return x.values[i]
}

func (x *AbstractSQLPartitionValues) AddValue(value expr.ISQLExpr) {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.values = append(x.values, value)
}

/**
 * VALUES LESS THAN (expr, expr)
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLPartitionValuesLessThan struct {
	*AbstractSQLPartitionValues
}

func NewPartitionValuesLessThan() *SQLPartitionValuesLessThan {
	x := new(SQLPartitionValuesLessThan)
	x.AbstractSQLPartitionValues = NewAbstractSQLPartitionValues()
	return x
}

/**
 * VALUES LESS THAN (expr, expr)
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLPartitionValuesLessThanMaxValue struct {
	*expr.AbstractSQLExpr
}

func NewPartitionValuesLessThanMaxValue() *SQLPartitionValuesLessThanMaxValue {
	x := new(SQLPartitionValuesLessThanMaxValue)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * VALUES IN (expr, expr)
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 */
type SQLPartitionValuesIn struct {
	*AbstractSQLPartitionValues
}

func NewPartitionValuesIn() *SQLPartitionValuesIn {
	x := new(SQLPartitionValuesIn)
	x.AbstractSQLPartitionValues = NewAbstractSQLPartitionValues()
	return x
}

/**
 * subpartition_definition:
    SUBPARTITION logical_name
        [[STORAGE] ENGINE [=] engine_name]
        [COMMENT [=] 'string' ]
        [DATA DIRECTORY [=] 'data_dir']
        [INDEX DIRECTORY [=] 'index_dir']
        [MAX_ROWS [=] max_number_of_rows]
        [MIN_ROWS [=] min_number_of_rows]
        [TABLESPACE [=] tablespace_name]
 * https://dev.mysql.com/doc/refman/8.0/en/create-table.html
 *
 *
 */
type SQLSubPartitionDefinition struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewSubPartitionDefinition() *SQLSubPartitionDefinition {
	x := new(SQLSubPartitionDefinition)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLSubPartitionDefinition) Name() expr.ISQLName {
	return x.name
}

func (x *SQLSubPartitionDefinition) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

// ---------------------------------------------- Table Partitioning End ----------------------------------------------

// ---------------------------------------------- Alter Table Start ----------------------------------------------
/**
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20table%20action
 */
type ISQLAlterTableAction interface {
	expr.ISQLExpr
}

/**
 * ADD [ COLUMN ] <column definition>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#add%20column%20definition

 * ADD [COLUMN] col_name column_definition [FIRST | AFTER col_name]
        [FIRST | AFTER col_name]
 * ADD [COLUMN] (col_name column_definition,...)
 *https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLAddColumnAlterTableAction struct {
	*expr.AbstractSQLExpr
	HasColumn bool
	Paren     bool
	columns   []*SQLTableColumn
}

func NewAddColumnAlterTableAction() *SQLAddColumnAlterTableAction {
	x := new(SQLAddColumnAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLAddColumnAlterTableAction) Columns() []*SQLTableColumn {
	return x.columns
}
func (x *SQLAddColumnAlterTableAction) Column(i int) *SQLTableColumn {
	return x.columns[i]
}
func (x *SQLAddColumnAlterTableAction) AddColumn(column *SQLTableColumn) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}

/**
 *  ADD INDEX [index_name]
        [index_type] (key_part,...) [index_option] ...
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLAddIndexAlterTableAction struct {
	*expr.AbstractSQLExpr
	name    expr.ISQLName
	Paren   bool
	columns []*SQLTableColumn
}

func NewAddIndexAlterTableAction() *SQLAddIndexAlterTableAction {
	x := new(SQLAddIndexAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * ADD KEY [index_name]
        [index_type] (key_part,...) [index_option] ...
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLAddKeyAlterTableAction struct {
	*expr.AbstractSQLExpr
	HasColumn bool
	Paren     bool
	columns   []*SQLTableColumn
}

func NewAddKeyAlterTableAction() *SQLAddKeyAlterTableAction {
	x := new(SQLAddKeyAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * ALTER [ COLUMN ] <column name> <alter column action>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20definition
 */
type SQLAlterColumnAlterTableAction struct {
	*expr.AbstractSQLExpr
	HasColumn bool
	column    expr.ISQLName
	action    ISQLAlterColumnAction
}

func NewAlterColumnAlterTableAction() *SQLAlterColumnAlterTableAction {
	x := new(SQLAlterColumnAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLAlterColumnAlterTableAction) Column() expr.ISQLName {
	return x.column
}
func (x *SQLAlterColumnAlterTableAction) SetColumn(column expr.ISQLName) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.column = column
}
func (x *SQLAlterColumnAlterTableAction) Action() ISQLAlterColumnAction {
	return x.action
}
func (x *SQLAlterColumnAlterTableAction) SetAction(action ISQLAlterColumnAction) {
	if action == nil {
		return
	}
	action.SetParent(x)
	x.action = action
}

/**
 * <set column default clause>
     |     <drop column default clause>
     |     <add column scope clause>
     |     <drop column scope clause>
     |     <alter identity column specification>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20action
 */
type ISQLAlterColumnAction interface {
	expr.ISQLExpr
}

/**
 * SET <default clause>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20action
 */
type SQLSetDefaultAlterColumnAction struct {
	*expr.AbstractSQLExpr
	defaultClause *SQLDefaultClause
}

func NewSetDefaultAlterColumnAction() *SQLSetDefaultAlterColumnAction {
	x := new(SQLSetDefaultAlterColumnAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * <drop column default clause>: DROP DEFAULT
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20action
 */
type SQLDropDefaultAlterColumnAction struct {
	*expr.AbstractSQLExpr
}

func NewDropDefaultAlterColumnAction() *SQLDropDefaultAlterColumnAction {
	x := new(SQLDropDefaultAlterColumnAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 *  <add column scope clause>: ADD SCOPE table_name
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20action
 */
type SQLAddScopeAlterColumnAction struct {
	*expr.AbstractSQLExpr
}

func NewAddScopeAlterColumnAction() *SQLAddScopeAlterColumnAction {
	x := new(SQLAddScopeAlterColumnAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * <drop column scope clause>: DROP SCOPE table_name
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20action
 */
type SQLDropScopeAlterColumnAction struct {
	*expr.AbstractSQLExpr
}

func NewDropScopeAlterColumnAction() *SQLDropScopeAlterColumnAction {
	x := new(SQLDropScopeAlterColumnAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * <alter identity column specification>:
 	RESTART WITH <sequence generator restart value>
 *  SET <basic sequence generator option>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20action
 */
type SQLRestartWithAlterColumnAction struct {
	*expr.AbstractSQLExpr
}

func NewRestartWithAlterColumnAction() *SQLRestartWithAlterColumnAction {
	x := new(SQLRestartWithAlterColumnAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * DROP [ COLUMN ] <column name> <drop behavior>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20column%20definition

 * DROP [COLUMN] col_name
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 *
 * DROP [ COLUMN ] [ IF EXISTS ] column_name [ RESTRICT | CASCADE ]
 * https://www.postgresql.org/docs/devel/sql-altertable.html
 */
type SQLDropColumnAlterTableAction struct {
	*expr.AbstractSQLExpr
	HasColumn bool
	IfExists bool
	column    expr.ISQLName
}

func NewDropColumnAlterTableAction() *SQLDropColumnAlterTableAction {
	x := new(SQLDropColumnAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropColumnAlterTableAction) Column() expr.ISQLName {
	return x.column
}
func (x *SQLDropColumnAlterTableAction) SetColumn(column expr.ISQLName) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.column = column
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#add%20table%20constraint%20definition
type SQLAddTableConstraintAlterTableAction struct {
	*expr.AbstractSQLExpr
	tableConstraint ISQLTableConstraint
}

func NewAddTableConstraintAlterTableAction() *SQLAddTableConstraintAlterTableAction {
	x := new(SQLAddTableConstraintAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLAddTableConstraintAlterTableAction) TableConstraint() ISQLTableConstraint {
	return x.tableConstraint
}

func (x *SQLAddTableConstraintAlterTableAction) SetTableConstraint(tableConstraint ISQLTableConstraint) {
	tableConstraint.SetParent(x)
	x.tableConstraint = tableConstraint
}

/**
* DROP INDEX name
* https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
*/
type SQLDropIndexAlterTableAction struct {
	*expr.AbstractSQLExpr
	IfExists bool
	name     expr.ISQLName
}

func NewDropIndexAlterTableAction() *SQLDropIndexAlterTableAction {
	x := new(SQLDropIndexAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropIndexAlterTableAction) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropIndexAlterTableAction) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
* DROP KEY name
* https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
*/
type SQLDropKeyAlterTableAction struct {
	*expr.AbstractSQLExpr
	IfExists bool
	name     expr.ISQLName
}

func NewDropKeyAlterTableAction() *SQLDropKeyAlterTableAction {
	x := new(SQLDropKeyAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropKeyAlterTableAction) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropKeyAlterTableAction) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * DROP CONSTRAINT name
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 *
 * DROP CONSTRAINT [ IF EXISTS ]  constraint_name [ RESTRICT | CASCADE ]
 * https://www.postgresql.org/docs/devel/sql-altertable.html
 */
type SQLDropTableConstraintAlterTableAction struct {
	*expr.AbstractSQLExpr
	IfExists bool
	name     expr.ISQLName

}

func NewDropTableConstraintAlterTableAction() *SQLDropTableConstraintAlterTableAction {
	x := new(SQLDropTableConstraintAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropTableConstraintAlterTableAction) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropTableConstraintAlterTableAction) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * DROP PRIMARY KEY
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLDropPrimaryKeyTableConstraintAlterTableAction struct {
	*expr.AbstractSQLExpr
}

func NewDropPrimaryKeyTableConstraintAlterTableAction() *SQLDropPrimaryKeyTableConstraintAlterTableAction {
	x := new(SQLDropPrimaryKeyTableConstraintAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * drop_constraint_clause: DROP UNIQUE (column [, column ]...)
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-TABLE.html#GUID-552E7373-BF93-477D-9DA3-B2C9386F2877
 */
type SQLDropUniqueTableConstraintAlterTableAction struct {
	*expr.AbstractSQLExpr
	name []expr.ISQLName
}

func NewDropUniqueAlterTableAction() *SQLDropUniqueTableConstraintAlterTableAction {
	x := new(SQLDropUniqueTableConstraintAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * DROP FOREIGN KEY fk_symbol
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLDropForeignKeyTableConstraintAlterTableAction struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewDropForeignKeyTableConstraintAlterTableAction() *SQLDropForeignKeyTableConstraintAlterTableAction {
	x := new(SQLDropForeignKeyTableConstraintAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropForeignKeyTableConstraintAlterTableAction) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropForeignKeyTableConstraintAlterTableAction) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * DROP FOREIGN KEY fk_symbol
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLDropCheckTableConstraintAlterTableAction struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewDropCheckTableConstraintAlterTableAction() *SQLDropCheckTableConstraintAlterTableAction {
	x := new(SQLDropCheckTableConstraintAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropCheckTableConstraintAlterTableAction) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropCheckTableConstraintAlterTableAction) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * ADD PARTITION (partition_definition)
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 *
 * ADD PARTITION [ partition ] add_range_partition_clause [, PARTITION [ partition ] add_range_partition_clause ]...
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-TABLE.html#GUID-552E7373-BF93-477D-9DA3-B2C9386F2877
 */
type SQLAddPartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	partition *SQLPartitionDefinition
}

func NewAddPartitionAlterTableAction() *SQLAddPartitionAlterTableAction {
	x := new(SQLAddPartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLAddPartitionAlterTableAction) Partition() *SQLPartitionDefinition {
	return x.partition
}
func (x *SQLAddPartitionAlterTableAction) SetPartition(partition *SQLPartitionDefinition) {
	if partition == nil {
		return
	}
	partition.SetParent(x)
	x.partition = partition
}

/**
 * ADD PARTITION [ partition ] add_range_partition_clause [, PARTITION [ partition ] add_range_partition_clause ]...
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-TABLE.html#GUID-552E7373-BF93-477D-9DA3-B2C9386F2877
 */
type SQLAddPartitionsAlterTableAction struct {
	*expr.AbstractSQLExpr
	partitions []*SQLPartitionDefinition
}

func NewAddPartitionsAlterTableAction() *SQLAddPartitionsAlterTableAction {
	x := new(SQLAddPartitionsAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLAddPartitionsAlterTableAction) Partitions() []*SQLPartitionDefinition {
	return x.partitions
}
func (x *SQLAddPartitionsAlterTableAction) AddPartition(partition *SQLPartitionDefinition) {
	if partition == nil {
		return
	}
	partition.SetParent(x)
	x.partitions = append(x.partitions, partition)
}

/**
 * DROP PARTITION partition_names
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLDropPartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewDropPartitionAlterTableAction() *SQLDropPartitionAlterTableAction {
	x := new(SQLDropPartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * DISCARD PARTITION {partition_names | ALL} TABLESPACE
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLDiscardPartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewDiscardPartitionAlterTableAction() *SQLDiscardPartitionAlterTableAction {
	x := new(SQLDiscardPartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * IMPORT PARTITION {partition_names | ALL} TABLESPACE
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLImportPartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewImportPartitionAlterTableAction() *SQLImportPartitionAlterTableAction {
	x := new(SQLImportPartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * TRUNCATE PARTITION {partition_names | ALL}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLTruncatePartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewTruncatePartitionAlterTableAction() *SQLTruncatePartitionAlterTableAction {
	x := new(SQLTruncatePartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLTruncatePartitionAlterTableAction) Names() []expr.ISQLName {
	return x.names
}
func (x *SQLTruncatePartitionAlterTableAction) Name(i int) expr.ISQLName {
	return x.names[i]
}
func (x *SQLTruncatePartitionAlterTableAction) AddName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}

/**
 * COALESCE PARTITION number
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLCoalescePartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}

func NewCoalescePartitionAlterTableAction() *SQLCoalescePartitionAlterTableAction {
	x := new(SQLCoalescePartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * REORGANIZE PARTITION partition_names INTO (partition_definitions)
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLReorganizePartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names      []expr.ISQLName
	partitions []*SQLPartitionDefinition
}

func (x *SQLReorganizePartitionAlterTableAction) Names() []expr.ISQLName {
	return x.names
}
func (x *SQLReorganizePartitionAlterTableAction) Name(i int) expr.ISQLName {
	return x.names[i]
}
func (x *SQLReorganizePartitionAlterTableAction) AddName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}
func (x *SQLReorganizePartitionAlterTableAction) Partitions() []*SQLPartitionDefinition {
	return x.partitions
}
func (x *SQLReorganizePartitionAlterTableAction) Partition(i int) *SQLPartitionDefinition {
	return x.partitions[i]
}
func (x *SQLReorganizePartitionAlterTableAction) AddPartition(partition *SQLPartitionDefinition) {
	if partition == nil {
		return
	}
	partition.SetParent(x)
	x.partitions = append(x.partitions, partition)
}

func NewReorganizePartitionAlterTableAction() *SQLReorganizePartitionAlterTableAction {
	x := new(SQLReorganizePartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * EXCHANGE PARTITION partition_name WITH TABLE tbl_name [{WITH | WITHOUT} VALIDATION]
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLExchangePartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	name  expr.ISQLName
	table expr.ISQLName
}

func NewExchangePartitionAlterTableAction() *SQLExchangePartitionAlterTableAction {
	x := new(SQLExchangePartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * ANALYZE PARTITION {partition_names | ALL}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLAnalyzePartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func (x *SQLAnalyzePartitionAlterTableAction) Names() []expr.ISQLName {
	return x.names
}
func (x *SQLAnalyzePartitionAlterTableAction) Name(i int) expr.ISQLName {
	return x.names[i]
}
func (x *SQLAnalyzePartitionAlterTableAction) AddName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}

func NewAnalyzePartitionAlterTableAction() *SQLAnalyzePartitionAlterTableAction {
	x := new(SQLAnalyzePartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * CHECK PARTITION {partition_names | ALL}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLCheckPartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewCheckPartitionAlterTableAction() *SQLCheckPartitionAlterTableAction {
	x := new(SQLCheckPartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLCheckPartitionAlterTableAction) Names() []expr.ISQLName {
	return x.names
}
func (x *SQLCheckPartitionAlterTableAction) Name(i int) expr.ISQLName {
	return x.names[i]
}
func (x *SQLCheckPartitionAlterTableAction) AddName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}

/**
 * OPTIMIZE PARTITION {partition_names | ALL}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLOptimizePartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewOptimizePartitionAlterTableAction() *SQLOptimizePartitionAlterTableAction {
	x := new(SQLOptimizePartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * REBUILD PARTITION {partition_names | ALL}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLRebuildPartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewRebuildPartitionAlterTableAction() *SQLRebuildPartitionAlterTableAction {
	x := new(SQLRebuildPartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * REPAIR PARTITION {partition_names | ALL}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLRepairPartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewRepairPartitionAlterTableAction() *SQLRepairPartitionAlterTableAction {
	x := new(SQLRepairPartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * REMOVE PARTITIONING
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
type SQLRemovePartitionAlterTableAction struct {
	*expr.AbstractSQLExpr
}

func NewRemovePartitionAlterTableAction() *SQLRemovePartitionAlterTableAction {
	x := new(SQLRemovePartitionAlterTableAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

// ---------------------------------------------- Alter Table End ----------------------------------------------

// RESTRICT | CASCADE
type SQLDropTableStatementRestrictOption struct {
	*expr.AbstractSQLExpr
}

func NewDropTableStatementRestrictOption() *SQLDropTableStatementRestrictOption {
	x := new(SQLDropTableStatementRestrictOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

type SQLDropTableStatementCascadeOption struct {
	*expr.AbstractSQLExpr
}

func NewDropTableStatementCascadeOption() *SQLDropTableStatementCascadeOption {
	x := new(SQLDropTableStatementCascadeOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
