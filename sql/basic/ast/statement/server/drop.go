package server

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP SERVER [ IF EXISTS ] server_name
 * https://dev.mysql.com/doc/refman/8.0/en/drop-server.html
 */
type SQLDropServerStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	name     expr.ISQLName
}
func NewDropServerStatement(dbType db.Type) *SQLDropServerStatement {
	x := new(SQLDropServerStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropServerStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropServerStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}