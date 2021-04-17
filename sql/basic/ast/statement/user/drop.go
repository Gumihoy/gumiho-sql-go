package user

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP USER [IF EXISTS] user [, user ] ...
 * https://dev.mysql.com/doc/refman/8.0/en/drop-role.html
 */
type SQLDropUserStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	names []expr.ISQLExpr
}
func NewDropRoleStatement(dbType db.Type) *SQLDropUserStatement {
	x := new(SQLDropUserStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropUserStatement) Names() []expr.ISQLExpr  {
	return x.names
}
func (x *SQLDropUserStatement) Name(i int) expr.ISQLExpr {
	return x.names[i]
}
func (x *SQLDropUserStatement) AddName(name expr.ISQLExpr)   {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}