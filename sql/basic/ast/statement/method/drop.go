package function

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type SQLDropMethodStatement struct {
	*statement.AbstractSQLStatement
}
func NewDropFunctionStatement(dbType db.Type) *SQLDropMethodStatement {
	x := new(SQLDropMethodStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}