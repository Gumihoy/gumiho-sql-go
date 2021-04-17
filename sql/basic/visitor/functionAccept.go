package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/function"
)

func AcceptMethodInvocation(visitor ISQLVisitor, x *function.SQLMethodInvocation) {
	if visitor.VisitMethodInvocation(visitor, x) {
		Accept(visitor, x.Name())

		for _, child := range x.Arguments() {
			Accept(visitor, child)
		}

	}
}


func AcceptStaticMethodInvocation(visitor ISQLVisitor, x *function.SQLStaticMethodInvocation) {
	if visitor.VisitStaticMethodInvocation(visitor, x) {
		Accept(visitor, x.TypeName())

		Accept(visitor, x.Name())
	}
}

func AcceptCastFunctionArgument(visitor ISQLVisitor, x *function.SQLCastFunctionArgument) {
	if visitor.VisitCastFunctionArgument(visitor, x) {
		Accept(visitor, x.Expr())

		Accept(visitor, x.DataType())
	}
}

