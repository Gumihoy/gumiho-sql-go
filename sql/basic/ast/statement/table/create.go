package table

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type TableScope string

const (
	TEMPORARY        TableScope = "TEMPORARY"
	GLOBAL                      = "GLOBAL"
	GLOBAL_TEMPORARY            = "GLOBAL TEMPORARY"
	LOCAL_TEMPORARY             = "LOCAL TEMPORARY"
)

/**
 *
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20definition
 *
 * CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
    (create_definition,...)
    [table_options]
    [partition_options]

CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
    [(create_definition,...)]
    [table_options]
    [partition_options]
    [IGNORE | REPLACE]
    [AS] query_expression

CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
    { LIKE old_tbl_name | (LIKE old_tbl_name) }

 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLCreateTableStatement struct {
	*statement.AbstractSQLStatement

	TableScope  TableScope
	IfNotExists bool
	name        expr.ISQLName

	Paren    bool
	elements []table.ISQLTableElement

	// MySQL: table_options
	options []expr.ISQLExpr

	partitionBy table.ISQLPartitionBy

	As       bool
	subQuery select_.ISQLSelectQuery
}

func NewCreateTableStatement(dbType db.Type) *SQLCreateTableStatement {
	x := new(SQLCreateTableStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	x.Paren = true
	x.elements = make([]table.ISQLTableElement, 0, 10)
	return x
}

func (x *SQLCreateTableStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLCreateTableStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLCreateTableStatement) ObjectType() db.SQLObjectType {
	return db.TABLE
}

func (x *SQLCreateTableStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateTableStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLCreateTableStatement) Elements() []table.ISQLTableElement {
	return x.elements
}
func (x *SQLCreateTableStatement) AddElement(element table.ISQLTableElement) {
	if element == nil {
		return
	}
	element.SetParent(x)
	x.elements = append(x.elements, element)
}
func (x *SQLCreateTableStatement) AddElements(elements []table.ISQLTableElement) {
	if elements == nil || len(elements) == 0 {
		return
	}
	for _, element := range elements {
		x.AddElement(element)
	}
}
func (x *SQLCreateTableStatement) Options() []expr.ISQLExpr {
	return x.options
}
func (x *SQLCreateTableStatement) Option(i int) expr.ISQLExpr {
	return x.options[i]
}
func (x *SQLCreateTableStatement) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
func (x *SQLCreateTableStatement) PartitionBy() table.ISQLPartitionBy {
	return x.partitionBy
}

func (x *SQLCreateTableStatement) SetPartitionBy(partitionBy table.ISQLPartitionBy) {
	if partitionBy == nil {
		return
	}
	partitionBy.SetParent(x)
	x.partitionBy = partitionBy
}

func (x *SQLCreateTableStatement) SubQuery() select_.ISQLSelectQuery {
	return x.subQuery
}

func (x *SQLCreateTableStatement) SetSubQuery(subQuery select_.ISQLSelectQuery) {
	subQuery.SetParent(x)
	x.subQuery = subQuery
}
