package user

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE USER [IF NOT EXISTS]
    user [auth_option] [, user [auth_option]] ...
    DEFAULT ROLE role [, role ] ...
    [REQUIRE {NONE | tls_option [[AND] tls_option] ...}]
    [WITH resource_option [resource_option] ...]
    [password_option | lock_option] ...
    [COMMENT 'comment_string' | ATTRIBUTE 'json_object']
 * https://dev.mysql.com/doc/refman/8.0/en/create-user.html
 */
type SQLCreateUserStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	names []expr.ISQLExpr

	// MySQL: DEFAULT ROLE role [, role ] ...
	roles []expr.ISQLExpr
	withOptions []expr.ISQLExpr
	options []expr.ISQLExpr
}

func NewCreateUserStatement(dbType db.Type) *SQLCreateUserStatement {
	x:=new(SQLCreateUserStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLCreateUserStatement) Names() []expr.ISQLExpr  {
	return x.names
}
func (x *SQLCreateUserStatement) Name(i int) expr.ISQLExpr {
	return x.names[i]
}
func (x *SQLCreateUserStatement) AddName(name expr.ISQLExpr)   {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}


func (x *SQLCreateUserStatement) Roles() []expr.ISQLExpr  {
	return x.roles
}
func (x *SQLCreateUserStatement) Role(i int) expr.ISQLExpr {
	return x.roles[i]
}
func (x *SQLCreateUserStatement) AddRole(role expr.ISQLExpr)   {
	if role == nil {
		return
	}
	role.SetParent(x)
	x.roles = append(x.roles, role)
}



