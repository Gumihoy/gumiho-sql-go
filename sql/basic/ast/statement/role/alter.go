package role

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 *
 */
type SQLAlterRoleStatement struct {
	*statement.AbstractSQLStatement
}
func NewAlterRoleStatement(dbType db.Type) *SQLAlterRoleStatement {
	x := new(SQLAlterRoleStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}