package role

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP ROLE <role name>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20role%20statement
 *
 * DROP ROLE [IF EXISTS] role [, role ] ...
 * https://dev.mysql.com/doc/refman/8.0/en/drop-role.html
 */
type SQLDropRoleStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	names []expr.ISQLExpr
}
func NewDropRoleStatement(dbType db.Type) *SQLDropRoleStatement {
	x := new(SQLDropRoleStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropRoleStatement) Names() []expr.ISQLExpr  {
	return x.names
}
func (x *SQLDropRoleStatement) Name(i int) expr.ISQLExpr {
	return x.names[i]
}
func (x *SQLDropRoleStatement) AddName(name expr.ISQLExpr)   {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}