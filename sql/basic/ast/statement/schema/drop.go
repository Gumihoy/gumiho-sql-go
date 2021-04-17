package schema

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP SCHEMA <schema name> <drop behavior>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20schema%20statement
 */
type SQLDropSchemaStatement struct {
	*statement.AbstractSQLStatement

	IfExists bool
	name     expr.ISQLName
	Behavior statement.SQLDropBehavior
}

func NewDropSchemaStatement(dbType db.Type) *SQLDropSchemaStatement {
	x := new(SQLDropSchemaStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLDropSchemaStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropSchemaStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}
