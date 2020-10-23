package visitor

import (
	"gumihoy.com/sql/basic/ast/expr/variable"
)

func AcceptVariableExpr(visitor ISQLVisitor, x *variable.SQLVariableExpr)  {
	visitor.VisitVariableExpr(x)
}

func AcceptBindVariableExpr(visitor ISQLVisitor, x *variable.SQLBindVariableExpr)  {
	if visitor.VisitBindVariableExpr(x) {
		Accept(visitor, x.Name())
	}

}