package statement

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/expr/select"
)

type ISQLStatement interface {
	ast.ISQLObject
}

type SQLStatement struct {
	*ast.SQLObject
}

func NewStatement() *SQLStatement {
	x :=new(SQLStatement)
	x.SQLObject = ast.NewObject()
	return x
}


/**
 *
 */
type SQLSelectStatement struct {
	*SQLStatement
	query select_.ISQLSelectQuery
}

func NewSelectStatement(query select_.ISQLSelectQuery) *SQLSelectStatement {
	x := new(SQLSelectStatement)
	x.SQLStatement = NewStatement()
	x.SetQuery(query)
	return x
}

func (x *SQLSelectStatement)Query() select_.ISQLSelectQuery {
	return x.query
}
func (x *SQLSelectStatement)SetQuery(query select_.ISQLSelectQuery)  {
	if query == nil {
		return
	}
	query.SetParent(x)
	x.query = query
}



/**
 *
 */
type SQLInsertStatement struct {
}
func NewInsertStatement() *SQLInsertStatement {
	return nil
}

/**
 *
 */
type SQLUpdateStatement struct {
}

func NewUpdateStatement() *SQLUpdateStatement {
	return nil
}

/**
 *
 */
type SQLDeleteStatement struct {
}

func NewDeleteStatement() *SQLDeleteStatement {
	return nil
}
