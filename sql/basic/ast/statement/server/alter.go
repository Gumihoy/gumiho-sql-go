package server

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER SERVER  server_name
    OPTIONS (option [, option] ...)
 * https://dev.mysql.com/doc/refman/8.0/en/alter-server.html
 */
type SQLAlterServerStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLName
	options []expr.ISQLExpr
}
func NewAlterServerStatement(dbType db.Type) *SQLAlterServerStatement {
	x := new(SQLAlterServerStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLAlterServerStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLAlterServerStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterServerStatement) Options() []expr.ISQLExpr {
	return x.options
}
func (x *SQLAlterServerStatement) Option(i int) expr.ISQLExpr {
	return x.options[i]
}
func (x *SQLAlterServerStatement) AddOption(option view.ISQLViewElement) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
func (x *SQLAlterServerStatement) AddOptions(options []view.ISQLViewElement) {
	if options == nil || len(options) == 0 {
		return
	}
	for _, option := range options {
		x.AddOption(option)
	}
}