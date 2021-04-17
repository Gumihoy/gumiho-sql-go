package role

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE ROLE <role name> [ WITH ADMIN <grantor> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#role%20definition
 *
 * CREATE ROLE [IF NOT EXISTS] role [, role ] ...
 * https://dev.mysql.com/doc/refman/8.0/en/create-role.html
 */
type WithAdminGrantor string

const (
	CURRENT_USER WithAdminGrantor = "CURRENT_USER"
	CURRENT_ROLE                  = "CURRENT_ROLE"
)

type SQLCreateRoleStatement struct {
	*statement.AbstractSQLStatement
	IfNotExists bool
	names       []expr.ISQLExpr

	WithAdminGrantor WithAdminGrantor
}

func NewCreateRoleStatement(dbType db.Type) *SQLCreateRoleStatement {
	x := new(SQLCreateRoleStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCreateRoleStatement) Names() []expr.ISQLExpr {
	return x.names
}
func (x *SQLCreateRoleStatement) Name(i int) expr.ISQLExpr {
	return x.names[i]
}
func (x *SQLCreateRoleStatement) AddName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}
