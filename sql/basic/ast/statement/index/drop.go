package index

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 *
 * DROP INDEX index_name ON tbl_name
    [algorithm_option | lock_option] ...
 * https://dev.mysql.com/doc/refman/8.0/en/drop-index.html
 *
 *
 */
type SQLDropIndexStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLName
	onTable expr.ISQLName
	options []expr.ISQLExpr
}

func NewDropIndexStatement(dbType db.Type) *SQLDropIndexStatement {
	x := new(SQLDropIndexStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLDropIndexStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropIndexStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLDropIndexStatement) OnTable() expr.ISQLName {
	return x.onTable
}
func (x *SQLDropIndexStatement) SetOnTable(onTable expr.ISQLName) {
	if onTable == nil {
		return
	}
	onTable.SetParent(x)
	x.onTable = onTable
}
func (x *SQLDropIndexStatement) Options() []expr.ISQLExpr{
	return x.options
}
func (x *SQLDropIndexStatement) Option(i int) expr.ISQLExpr{
	return x.options[i]
}
func (x *SQLDropIndexStatement) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
