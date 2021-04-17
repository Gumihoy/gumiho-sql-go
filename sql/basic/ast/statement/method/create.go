package function

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE <SQL-invoked function>
 * { <function specification> | <method specification designator> } <routine body>
 <function specification> 	::=
		FUNCTION <schema qualified routine name>
         <SQL parameter declaration list> <returns clause> <routine characteristics> [ <dispatch clause> ]
<method specification designator>    ::=
         SPECIFIC METHOD <specific method name>
     |     [ INSTANCE | STATIC | CONSTRUCTOR ] METHOD <method name> <SQL parameter declaration list>
         [ <returns clause> ] FOR <schema-resolved user-defined type name>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#schema%20function
 *
 *
 */
type SQLCreateMethodStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}
func NewCreateFunctionStatement(dbType db.Type) *SQLCreateMethodStatement {
	x := new(SQLCreateMethodStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
