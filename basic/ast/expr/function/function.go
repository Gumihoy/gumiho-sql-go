package function

import (
	"gumihoy.com/sql/basic/ast/expr"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#method%20invocation
// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#static%20method%20invocation
// https://dev.mysql.com/doc/refman/8.0/en/sql-function-reference.html
// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/Functions.html#GUID-D079EFD3-C683-441F-977E-2C9503089982
type ISQLFunction interface {
	expr.ISQLExpr

	Name() expr.ISQLIdentifier
	SetName(name expr.ISQLIdentifier)

	Arguments() []expr.ISQLExpr
	AddArgument(argument expr.ISQLExpr)
}

type AbstractSQLFunction struct {
	expr.SQLExpr
	name      expr.ISQLIdentifier
	arguments []expr.ISQLExpr
}

func (x *AbstractSQLFunction) Name() expr.ISQLIdentifier {
	return x.name
}

func (x *AbstractSQLFunction) SetName(name expr.ISQLIdentifier) {
	name.SetParent(x)
	x.name = name
}

func (x *AbstractSQLFunction) Arguments() []expr.ISQLExpr {
	return x.arguments
}

func (x *AbstractSQLFunction) AddArgument(argument expr.ISQLExpr) {
	argument.SetParent(x)
	x.arguments = append(x.arguments, argument)
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#method%20invocation
type SQLMethodInvocation struct {
	AbstractSQLFunction
}

func NewMethodInvocation(name expr.ISQLIdentifier, arguments ...expr.ISQLExpr) *SQLMethodInvocation {
	x := new(SQLMethodInvocation)
	x.SetName(name)
	for _, argument := range arguments {
		argument.SetParent(x)
	}
	x.arguments = arguments
	return x
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#static%20method%20invocation
// <type name>::<method name> [ <SQL argument list> ]
type SQLStaticMethodInvocation struct {
	AbstractSQLFunction

	typeName expr.ISQLName
}

func (SQLStaticMethodInvocation) Name() expr.ISQLIdentifier {
	panic("implement me")
}

func (SQLStaticMethodInvocation) SetName(name expr.ISQLIdentifier) {
	panic("implement me")
}

func (SQLStaticMethodInvocation) Arguments() []expr.ISQLExpr {
	panic("implement me")
}

func (SQLStaticMethodInvocation) AddArgument(argument expr.ISQLExpr) {
	panic("implement me")
}
