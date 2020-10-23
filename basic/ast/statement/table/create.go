package table

import (
	"gumihoy.com/sql/basic/ast/expr"
	"gumihoy.com/sql/basic/ast/expr/datatype"
)

type TableScope string

const (
	GLOBAL           TableScope = "GLOBAL"
	GLOBAL_TEMPORARY TableScope = "GLOBAL TEMPORARY"
	LOCAL_TEMPORARY  TableScope = "LOCAL TEMPORARY"
)

type SQLCreateTableStatement struct {
	tableScope TableScope
	name       expr.ISQLName
	elements   []ISQLTableElement
}

type ISQLTableElement interface {
	expr.ISQLExpr
}

type SQLTableColumn struct {
	*expr.SQLExpr
	name     expr.ISQLIdentifier
	dataType datatype.ISQLDataType
}

func NewColumn() *SQLTableColumn {
	x := new(SQLTableColumn)
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
	return x.name
}
func (x *SQLTableColumn) SetDataType(datatype datatype.ISQLDataType) {
	datatype.SetParent(x)
	x.dataType = datatype
}



// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20constraint%20definition
type ISQLTableConstraint interface {
	ISQLTableElement
	Name() expr.ISQLName
	SetName(name expr.ISQLName)
}

type SQLTableConstraint struct {
	expr.SQLExpr
	name expr.ISQLName
}

func (x *SQLTableConstraint) Name() expr.ISQLName {
	return x.name
}

func (x *SQLTableConstraint) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#unique%20constraint%20definition
type SQLPrimaryKeyTableConstraint struct {
	SQLTableConstraint
}

type SQLUniqueTableConstraint struct {
	SQLTableConstraint
}

func NewUniqueTableConstraint() *SQLUniqueTableConstraint {
	x := new(SQLUniqueTableConstraint)
	return x
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#referential%20constraint%20definition
type SQLForeignKeyTableConstraint struct {
	SQLTableConstraint

	referencingColumn []expr.ISQLName

	referencedTable  expr.ISQLName
	referencedColumn []expr.ISQLName
}

type SQLMatchType string

const (
	FULL    SQLMatchType = "FULL"
	PARTIAL SQLMatchType = "PARTIAL"
	SIMPLE  SQLMatchType = "SIMPLE"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#check%20constraint%20definition
type SQLCheckTableConstraint struct {
	SQLTableConstraint
	condition expr.ISQLExpr
}

type SQLTableLikeClause struct {
	expr.SQLExpr
}

type SQLSelfReferencingColumn struct {
	expr.SQLExpr
}
type SQLColumnOption struct {
	expr.SQLExpr
}

// ---------------------------------------------- Column Constraint ----------------------------------------------

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20constraint%20definition
type ISQLColumnConstraint interface {
	expr.ISQLExpr
	Name() expr.ISQLName
	SetName(name expr.ISQLName)
}

type SQLColumnConstraint struct {
	expr.SQLExpr
	name expr.ISQLName
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#constraint%20characteristics
type SQLColumnConstraintCharacteristics struct {
	expr.SQLExpr
}

func (x *SQLColumnConstraint) Name() expr.ISQLName {
	return x.name
}

func (x *SQLColumnConstraint) SetName(name expr.ISQLName) {
	name.SetParent(x)
}

type SQLPrimaryKeyColumnConstraint struct {
	SQLColumnConstraint
}

type SQLUniqueColumnConstraint struct {
	SQLColumnConstraint
}
type SQLNotNullColumnConstraint struct {
	SQLColumnConstraint
}
type SQLNullColumnConstraint struct {
	SQLColumnConstraint
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#references%20specification
type SQLReferencesColumnConstraint struct {
	SQLColumnConstraint
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#check%20constraint%20definition
type SQLCheckColumnConstraint struct {
	SQLColumnConstraint
	condition expr.ISQLExpr
}
