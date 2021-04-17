package server

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE SERVER server_name
    FOREIGN DATA WRAPPER wrapper_name
    OPTIONS (option [, option] ...)

option: {
    HOST character-literal
  | DATABASE character-literal
  | USER character-literal
  | PASSWORD character-literal
  | SOCKET character-literal
  | OWNER character-literal
  | PORT numeric-literal
}
 * https://dev.mysql.com/doc/refman/8.0/en/create-server.html
 */
type SQLCreateServerStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLName
	wrapperName expr.ISQLName
	options []expr.ISQLExpr
}
func NewCreateServerStatement(dbType db.Type) *SQLCreateServerStatement {
	x := new(SQLCreateServerStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCreateServerStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateServerStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLCreateServerStatement) WrapperName() expr.ISQLName {
	return x.wrapperName
}
func (x *SQLCreateServerStatement) SetWrapperName(wrapperName expr.ISQLName) {
	if wrapperName == nil {
		return
	}
	wrapperName.SetParent(x)
	x.wrapperName = wrapperName
}
func (x *SQLCreateServerStatement) Options() []expr.ISQLExpr {
	return x.options
}
func (x *SQLCreateServerStatement) Option(i int) expr.ISQLExpr {
	return x.options[i]
}
func (x *SQLCreateServerStatement) AddOption(option view.ISQLViewElement) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
func (x *SQLCreateServerStatement) AddOptions(options []view.ISQLViewElement) {
	if options == nil || len(options) == 0 {
		return
	}
	for _, option := range options {
		x.AddOption(option)
	}
}