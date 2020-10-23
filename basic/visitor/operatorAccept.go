package visitor

import (
	"gumihoy.com/sql/basic/ast/expr/operator"
)

func AcceptUnaryOperatorExpr(visitor ISQLVisitor, x *operator.SQLUnaryOperatorExpr)  {
	visitor.VisitUnaryOperatorExpr(x)
}

func AcceptBinaryOperatorExpr(visitor ISQLVisitor, x *operator.SQLBinaryOperatorExpr)  {
	if visitor.VisitBinaryOperatorExpr(x) {
		Accept(visitor, x.Left())
		Accept(visitor, x.Right())
	}

}