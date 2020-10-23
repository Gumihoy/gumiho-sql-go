package statement

import (
	"gumihoy.com/sql/basic/ast/expr/select"
	"gumihoy.com/sql/basic/ast/statement"
)

type IMySQLStatement interface {
	statement.ISQLStatement
}

type MySQLStatement struct {
	statement.SQLStatement
}

type MySQLSelectStatement struct {
	*statement.SQLSelectStatement
}

func NewSelectStatement(query select_.ISQLSelectQuery) *MySQLSelectStatement {
	x := new(MySQLSelectStatement)
	x.SQLSelectStatement = statement.NewSelectStatement(query)
	return x
}


