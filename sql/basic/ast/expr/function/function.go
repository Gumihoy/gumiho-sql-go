package function

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
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

type abstractSQLFunction struct {
	*expr.AbstractSQLExpr
	name      expr.ISQLExpr
	Paren     bool
	arguments []expr.ISQLExpr
}

func NewAbstractSQLFunction() *abstractSQLFunction {
	x := new(abstractSQLFunction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.Paren = true
	x.arguments = make([]expr.ISQLExpr, 0, 10)
	return x
}
func (x *abstractSQLFunction) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.name {
		x.SetName(target)
		return true
	}
	for i, child := range x.arguments {
		if source == child {
			return x.SetArgument(i, target)
		}
	}
	return false
}
func (x *abstractSQLFunction) Clone() ast.ISQLObject {
	panic("implement me")
}
func (x *abstractSQLFunction) Name() expr.ISQLExpr {
	return x.name
}

func (x *abstractSQLFunction) SetName(name expr.ISQLExpr) {
	name.SetParent(x)
	x.name = name
}

func (x *abstractSQLFunction) Arguments() []expr.ISQLExpr {
	return x.arguments
}
func (x *abstractSQLFunction) SetArgument(i int, argument expr.ISQLExpr) bool {
	if argument == nil {
		return false
	}
	x.arguments[i] = argument
	return true
}
func (x *abstractSQLFunction) AddArgument(argument expr.ISQLExpr) {
	if argument == nil {
		return
	}
	argument.SetParent(x)
	x.arguments = append(x.arguments, argument)
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#method%20invocation
type SQLMethodInvocation struct {
	*abstractSQLFunction
}

func NewMethodInvocation(name expr.ISQLExpr, arguments ...expr.ISQLExpr) *SQLMethodInvocation {
	x := new(SQLMethodInvocation)
	x.abstractSQLFunction = NewAbstractSQLFunction()

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
	*abstractSQLFunction
	typeName expr.ISQLName
}

func NewStaticMethodInvocation() *SQLStaticMethodInvocation {
	x := new(SQLStaticMethodInvocation)
	x.abstractSQLFunction = NewAbstractSQLFunction()
	return x
}

func (self *SQLStaticMethodInvocation) TypeName() expr.ISQLName {
	return self.typeName
}

func (self *SQLStaticMethodInvocation) SetTypeName(typeName expr.ISQLName) {
	if typeName == nil {
		return
	}
	typeName.SetParent(self)
	self.typeName = typeName
}

/**
 * cast( expr as datatype[p])
 * https://dev.mysql.com/doc/refman/8.0/en/cast-functions.html#function_cast
 */
// type SQLCastFunction struct {
// 	*abstractSQLFunction
// }

/**
 * expr as datatype[p]
 * https://dev.mysql.com/doc/refman/8.0/en/cast-functions.html#function_cast
 */
type SQLCastFunctionArgument struct {
	*expr.AbstractSQLExpr

	expr     expr.ISQLExpr
	dataType datatype.ISQLDataType
}

func NewCastFunctionArgument(left expr.ISQLExpr, dataType datatype.ISQLDataType) *SQLCastFunctionArgument {
	x := new(SQLCastFunctionArgument)
	x.AbstractSQLExpr = expr.NewAbstractExpr()

	x.SetExpr(left)
	x.SetDataType(dataType)
	return x
}
func (x *SQLCastFunctionArgument) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.expr {
		x.SetExpr(target)
		return true
	}
	if source == x.dataType {
		x.SetDataType(target)
		return true
	}
	return false
}

func (x *SQLCastFunctionArgument) Expr() expr.ISQLExpr {
	return x.expr
}

func (x *SQLCastFunctionArgument) SetExpr(expr expr.ISQLExpr) {
	if expr == nil {
		return
	}
	expr.SetParent(x)
	x.expr = expr
}

func (self *SQLCastFunctionArgument) DataType() datatype.ISQLDataType {
	return self.dataType
}

func (self *SQLCastFunctionArgument) SetDataType(dataType datatype.ISQLDataType) {
	if dataType == nil {
		return
	}
	dataType.SetParent(self)
	self.dataType = dataType
}


//
type SQLSystemDate struct {
	*abstractSQLFunction
}

func NewSystemDate() {

}
