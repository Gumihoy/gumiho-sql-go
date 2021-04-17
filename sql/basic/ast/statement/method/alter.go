package function

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type SQLAlterMethodStatement struct {
	*statement.AbstractSQLStatement
}
func NewAlterFunctionStatement(dbType db.Type) *SQLAlterMethodStatement {
	x := new(SQLAlterMethodStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}