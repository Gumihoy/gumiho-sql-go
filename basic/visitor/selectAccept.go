package visitor

import (
	"gumihoy.com/sql/basic/ast/expr/select"
	"gumihoy.com/sql/basic/ast/statement"
)

func AcceptSelectStatement(visitor ISQLVisitor, x *statement.SQLSelectStatement) {
	if visitor.VisitSelectStatement(x) {
		Accept(visitor, x.Query())
	}
}

func AcceptSelectQuery(visitor ISQLVisitor, x *select_.SQLSelectQuery) {
	if visitor.VisitSelectQuery(x) {

		for _, child := range x.SelectElements() {
			Accept(visitor, child)
		}

		for _, child := range x.SelectTargetElements() {
			Accept(visitor, child)
		}

		Accept(visitor, x.FromClause())
		Accept(visitor, x.WhereClause())
		Accept(visitor, x.GroupByClause())
		Accept(visitor, x.WindowClause())
		Accept(visitor, x.OrderByClause())
		Accept(visitor, x.LimitClause())
		Accept(visitor, x.LockClause())
	}
}
func AcceptParenSelectQuery(visitor ISQLVisitor, x *select_.SQLParenSelectQuery) {
	if visitor.VisitParenSelectQuery(x) {
		Accept(visitor, x.SubQuery())
		Accept(visitor, x.OrderByClause())
		Accept(visitor, x.LimitClause())
		Accept(visitor, x.LockClause())
	}
}
func AcceptSelectUnionQuery(visitor ISQLVisitor, x *select_.SQLSelectUnionQuery) {
	if visitor.VisitSelectUnionQuery(x) {
		Accept(visitor, x.Left())
		Accept(visitor, x.Right())
	}
}

func AcceptSelectElement(visitor ISQLVisitor, x *select_.SQLSelectElement) {
	if visitor.VisitSelectElement(x) {
		Accept(visitor, x.Expr())
		Accept(visitor, x.Alias())
	}
}
func AcceptSelectTargetElement(visitor ISQLVisitor, x *select_.SQLSelectTargetElement) {

}
func AcceptFromClause(visitor ISQLVisitor, x *select_.SQLFromClause)                         {
	if visitor.VisitFromClause(x) {
		Accept(visitor, x.TableReference())
	}
}
func AcceptTableReference(visitor ISQLVisitor, x *select_.SQLTableReference)                 {
	if visitor.VisitTableReference(x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Alias())

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptSubQueryTableReference(visitor ISQLVisitor, x *select_.SQLSubQueryTableReference) {
	if visitor.VisitSubQueryTableReference(x) {
		Accept(visitor, x.SubQuery())
	}
}
func AcceptJoinTableReference(visitor ISQLVisitor, x *select_.SQLJoinTableReference)         {
	if visitor.VisitJoinTableReference(x) {
		Accept(visitor, x.Left())
		Accept(visitor, x.Right())
		Accept(visitor, x.Condition())
	}
}
func AcceptJoinOnCondition(visitor ISQLVisitor, x *select_.SQLJoinOnCondition)               {
	if visitor.VisitJoinOnCondition(x) {
		Accept(visitor, x.Condition())
	}
}
func AcceptJoinUsingCondition(visitor ISQLVisitor, x *select_.SQLJoinUsingCondition)         {
	if visitor.VisitJoinUsingCondition(x) {
		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}

func AcceptWhereClause(visitor ISQLVisitor, x *select_.SQLWhereClause)                 {
	if visitor.VisitWhereClause(x) {
		Accept(visitor, x.Condition())
	}
}
func AcceptGroupByHavingClause(visitor ISQLVisitor, x *select_.SQLGroupByHavingClause) {

}
func AcceptHavingGroupByClause(visitor ISQLVisitor, x *select_.SQLHavingGroupByClause) {

}
func AcceptOrderByClause(visitor ISQLVisitor, x *select_.SQLOrderByClause)             {

}

func AcceptLimitOffsetClause(visitor ISQLVisitor, x *select_.SQLLimitOffsetClause) {

}
func AcceptOffsetFetchClause(visitor ISQLVisitor, x *select_.SQLOffsetFetchClause) {

}

func AcceptForUpdateClause(visitor ISQLVisitor, x *select_.SQLForUpdateClause) {

}
