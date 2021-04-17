package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/operator"
)

func AcceptUnaryOperatorExpr(visitor ISQLVisitor, x *operator.SQLUnaryOperatorExpr)  {
	visitor.VisitUnaryOperatorExpr(visitor, x)
}

func AcceptBinaryOperatorExpr(visitor ISQLVisitor, x *operator.SQLBinaryOperatorExpr)  {
	if visitor.VisitBinaryOperatorExpr(visitor, x) {
		Accept(visitor, x.Left())
		Accept(visitor, x.Right())
	}

}