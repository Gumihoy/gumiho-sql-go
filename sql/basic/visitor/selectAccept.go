package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
)

func AcceptSelectQuery(visitor ISQLVisitor, x *select_.SQLSelectQuery) {
	if visitor.VisitSelectQuery(visitor, x) {

		Accept(visitor, x.WithClause())

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
	if visitor.VisitParenSelectQuery(visitor, x) {
		Accept(visitor, x.SubQuery())
		Accept(visitor, x.OrderByClause())
		Accept(visitor, x.LimitClause())
		Accept(visitor, x.LockClause())
	}
}
func AcceptSelectUnionQuery(visitor ISQLVisitor, x *select_.SQLSelectUnionQuery) {
	if visitor.VisitSelectUnionQuery(visitor, x) {
		Accept(visitor, x.Left())
		Accept(visitor, x.Right())
	}
}

func AcceptWithClause(visitor ISQLVisitor, x *select_.SQLWithClause) {
	if visitor.VisitWithClause(visitor, x) {

		for _, child := range x.FactoringClause() {
			Accept(visitor, child)
		}
	}
}
func AcceptSubQueryFactoringClause(visitor ISQLVisitor, x *select_.SQLSubQueryFactoringClause) {
	if visitor.VisitSubQueryFactoringClause(visitor, x) {
		Accept(visitor, x.Name())

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}

		Accept(visitor, x.SubQuery())
	}
}
func AcceptSearchClause(visitor ISQLVisitor, x *select_.SQLSearchClause) {
	if visitor.VisitSearchClause(visitor, x) {

	}
}

func AcceptSubAvFactoringClause(visitor ISQLVisitor, x *select_.SQLSubAvFactoringClause) {
	if visitor.VisitSubAvFactoringClause(visitor, x) {
		Accept(visitor, x.Name())

		Accept(visitor, x.SubAvClause())
	}
}
func AcceptSubAvClause(visitor ISQLVisitor, x *select_.SQLSubAvClause) {
	if visitor.VisitSubAvClause(visitor, x) {
		Accept(visitor, x.Name())
	}
}

func AcceptSelectElement(visitor ISQLVisitor, x *select_.SQLSelectElement) {
	if visitor.VisitSelectElement(visitor, x) {
		Accept(visitor, x.Expr())
		Accept(visitor, x.Alias())
	}
}

func AcceptSelectTargetElement(visitor ISQLVisitor, x *select_.SQLSelectTargetElement) {

}
func AcceptFromClause(visitor ISQLVisitor, x *select_.SQLFromClause) {
	if visitor.VisitFromClause(visitor, x) {
		Accept(visitor, x.TableReference())
	}
}
func AcceptTableReference(visitor ISQLVisitor, x *select_.SQLTableReference) {
	if visitor.VisitTableReference(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.PartitionExtensionClause())
		Accept(visitor, x.Alias())

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}

func AcceptPartitionClause(visitor ISQLVisitor, x *select_.SQLPartitionClause) {
	if visitor.VisitPartitionClause(visitor, x) {

		for _, child := range x.Names() {
			Accept(visitor, child)
		}
	}
}
func AcceptPartitionForClause(visitor ISQLVisitor, x *select_.SQLPartitionForClause) {
	if visitor.VisitPartitionForClause(visitor, x) {

		for _, child := range x.Names() {
			Accept(visitor, child)
		}
	}
}
func AcceptSubPartitionClause(visitor ISQLVisitor, x *select_.SQLSubPartitionClause) {
	if visitor.VisitSubPartitionClause(visitor, x) {

		for _, child := range x.Names() {
			Accept(visitor, child)
		}
	}
}
func AcceptSubPartitionForClause(visitor ISQLVisitor, x *select_.SQLSubPartitionForClause) {
	if visitor.VisitSubPartitionForClause(visitor, x) {

		for _, child := range x.Names() {
			Accept(visitor, child)
		}
	}
}
func AcceptSampleClause(visitor ISQLVisitor, x *select_.SQLSampleClause) {
	if visitor.VisitSampleClause(visitor, x) {

		Accept(visitor, x.Percent())
		Accept(visitor, x.SeedValue())
	}
}
func AcceptOJTableReference(visitor ISQLVisitor, x *select_.SQLOJTableReference) {
	if visitor.VisitOJTableReference(visitor, x) {
		Accept(visitor, x.TableReference())
	}
}
func AcceptSubQueryTableReference(visitor ISQLVisitor, x *select_.SQLSubQueryTableReference) {
	if visitor.VisitSubQueryTableReference(visitor, x) {
		Accept(visitor, x.SubQuery())
	}
}
func AcceptJoinTableReference(visitor ISQLVisitor, x *select_.SQLJoinTableReference) {
	if visitor.VisitJoinTableReference(visitor, x) {
		Accept(visitor, x.Left())
		Accept(visitor, x.Right())
		Accept(visitor, x.Condition())
	}
}
func AcceptJoinOnCondition(visitor ISQLVisitor, x *select_.SQLJoinOnCondition) {
	if visitor.VisitJoinOnCondition(visitor, x) {
		Accept(visitor, x.Condition())
	}
}
func AcceptJoinUsingCondition(visitor ISQLVisitor, x *select_.SQLJoinUsingCondition) {
	if visitor.VisitJoinUsingCondition(visitor, x) {
		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}

func AcceptWhereClause(visitor ISQLVisitor, x *select_.SQLWhereClause) {
	if visitor.VisitWhereClause(visitor, x) {
		Accept(visitor, x.Condition())
	}
}

func AcceptHierarchicalQueryClauseConnectBy(visitor ISQLVisitor, x *select_.SQLHierarchicalQueryClauseConnectBy) {
	if visitor.VisitHierarchicalQueryClauseConnectBy(visitor, x) {
		Accept(visitor, x.ConnectByCondition())
		Accept(visitor, x.StartWithCondition())
	}
}
func AcceptHierarchicalQueryClauseStartWith(visitor ISQLVisitor, x *select_.SQLHierarchicalQueryClauseStartWith) {
	if visitor.VisitHierarchicalQueryClauseStartWith(visitor, x) {
		Accept(visitor, x.StartWithCondition())
		Accept(visitor, x.ConnectByCondition())
	}
}

func AcceptGroupByHavingClause(visitor ISQLVisitor, x *select_.SQLGroupByHavingClause) {
	if visitor.VisitGroupByHavingClause(visitor, x) {

		for _, child := range x.Elements() {
			Accept(visitor, child)
		}

		Accept(visitor, x.Having())
	}
}
func AcceptHavingGroupByClause(visitor ISQLVisitor, x *select_.SQLHavingGroupByClause) {
	if visitor.VisitHavingGroupByClause(visitor, x) {

		Accept(visitor, x.Having())

		for _, child := range x.Elements() {
			Accept(visitor, child)
		}

	}
}
func AcceptGroupByElement(visitor ISQLVisitor, x *select_.SQLGroupByElement) {
	if visitor.VisitGroupByElement(visitor, x) {

		Accept(visitor, x.Expr())

	}
}

func AcceptOrderByClause(visitor ISQLVisitor, x *select_.SQLOrderByClause) {

	if visitor.VisitOrderByClause(visitor, x) {

		for _, child := range x.Elements() {
			Accept(visitor, child)
		}

	}
}
func AcceptOrderByElement(visitor ISQLVisitor, x *select_.SQLOrderByElement) {

	if visitor.VisitOrderByElement(visitor, x) {

		Accept(visitor, x.Key())

	}
}

func AcceptLimitOffsetClause(visitor ISQLVisitor, x *select_.SQLLimitOffsetClause) {
	if visitor.VisitLimitOffsetClause(visitor, x) {

		if x.Offset {
			Accept(visitor, x.OffsetExpr())
			Accept(visitor, x.CountExpr())

		} else {

			Accept(visitor, x.CountExpr())
			Accept(visitor, x.OffsetExpr())
		}

	}
}
func AcceptOffsetFetchClause(visitor ISQLVisitor, x *select_.SQLOffsetFetchClause) {
	if visitor.VisitOffsetFetchClause(visitor, x) {

		Accept(visitor, x.OffsetExpr())
		Accept(visitor, x.CountExpr())

	}
}

func AcceptForUpdateClause(visitor ISQLVisitor, x *select_.SQLForUpdateClause) {
	if visitor.VisitForUpdateClause(visitor, x) {
		for _, child := range x.Tables() {
			Accept(visitor, child)
		}
	}
}
func AcceptForShareClause(visitor ISQLVisitor, x *select_.SQLForShareClause)  {
	if visitor.VisitForShareClause(visitor, x) {
		for _, child := range x.Tables() {
			Accept(visitor, child)
		}
	}
}
func AcceptLockInShareModeClause(visitor ISQLVisitor, x *select_.SQLLockInShareModeClause)  {
	if visitor.VisitLockInShareModeClause(visitor, x) {
	}
}