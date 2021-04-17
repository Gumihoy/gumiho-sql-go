package statement

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
)

type IMySQLStatement interface {
	statement.ISQLStatement
}

type MySQLStatement struct {
	*statement.AbstractSQLStatement
}

type MySQLSelectStatement struct {
	*statement.SQLSelectStatement
}

func NewSelectStatement(query select_.ISQLSelectQuery) *MySQLSelectStatement {
	x := new(MySQLSelectStatement)
	x.SQLSelectStatement = statement.NewSelectStatement(query)
	return x
}


