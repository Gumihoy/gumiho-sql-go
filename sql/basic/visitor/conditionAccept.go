package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/condition"
)

func AcceptIsCondition(visitor ISQLVisitor, x *condition.SQLIsCondition)  {
	if visitor.VisitIsCondition(visitor, x) {
		Accept(visitor, x.Expr())
	}
}

func AcceptLikeCondition(visitor ISQLVisitor, x *condition.SQLLikeCondition)  {
	if visitor.VisitLikeCondition(visitor, x) {
		Accept(visitor, x.Expr())
		Accept(visitor, x.Pattern())
		Accept(visitor, x.Escape())
	}
}

func AcceptBetweenCondition(visitor ISQLVisitor, x *condition.SQLBetweenCondition)  {
	if visitor.VisitBetweenCondition(visitor, x) {
		Accept(visitor, x.Expr())
		Accept(visitor, x.Between())
		Accept(visitor, x.And())
	}
}

func AcceptInCondition(visitor ISQLVisitor, x *condition.SQLInCondition)  {
	if visitor.VisitInCondition(visitor, x) {
		Accept(visitor, x.Expr())

		for _, child := range x.Values() {
			Accept(visitor, child)
		}
	}
}

func AcceptIsASetCondition(visitor ISQLVisitor, x *condition.SQLIsASetCondition)  {
	if visitor.VisitIsASetCondition(visitor, x) {
		Accept(visitor, x.Expr())

	}
}