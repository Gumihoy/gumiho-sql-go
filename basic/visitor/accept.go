package visitor

import (
	"fmt"
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/expr"
	"gumihoy.com/sql/basic/ast/expr/literal"
	"gumihoy.com/sql/basic/ast/expr/operator"
	"gumihoy.com/sql/basic/ast/expr/select"
	"gumihoy.com/sql/basic/ast/expr/variable"
	"gumihoy.com/sql/basic/ast/statement"
)

type ISQLAccept interface {
}

type SQLAccept struct {
}

func Accept(visitor ISQLVisitor, x ast.ISQLObject) {
	if ast.IsNil(x) {
		return
	}
	if visitor == nil {
		panic("visitor is nil.")
	}

	visitor.BeforeVisit(x)
	switch t := x.(type) {

	// ---------------------------- Literal Start ------------------------------------
	case *literal.SQLStringLiteral:
		x, _ := x.(*literal.SQLStringLiteral)
		AcceptStringLiteral(visitor, x)
		break
	case *literal.SQLCharacterStringLiteral:
		x, _ := x.(*literal.SQLCharacterStringLiteral)
		AcceptCharacterStringLiteral(visitor, x)
		break

	case *literal.SQLIntegerLiteral:
		x, _ := x.(*literal.SQLIntegerLiteral)
		AcceptIntegerLiteral(visitor, x)
		break
	case *literal.SQLFloatingPointLiteral:
		x, _ := x.(*literal.SQLFloatingPointLiteral)
		AcceptFloatingPointLiteral(visitor, x)
		break

	case *literal.SQLDateLiteral:
		x, _ := x.(*literal.SQLDateLiteral)
		AcceptDateLiteral(visitor, x)
		break
	case *literal.SQLTimeLiteral:
		x, _ := x.(*literal.SQLTimeLiteral)
		AcceptTimeLiteral(visitor, x)
		break
	case *literal.SQLTimestampLiteral:
		x, _ := x.(*literal.SQLTimestampLiteral)
		AcceptTimestampLiteral(visitor, x)
		break

	case *literal.SQLBooleanLiteral:
		x, _ := x.(*literal.SQLBooleanLiteral)
		AcceptBooleanLiteral(visitor, x)
		break
	// ---------------------------- Literal End ------------------------------------



	// ---------------------------- Identifier Start ------------------------------------
	case *expr.SQLUnQuotedIdentifier:
		x, _ := x.(*expr.SQLUnQuotedIdentifier)
		AcceptIdentifier(visitor, x)
		break
	case *expr.SQLDoubleQuotedIdentifier:
		x, _ := x.(*expr.SQLDoubleQuotedIdentifier)
		AcceptDoubleQuotedIdentifier(visitor, x)
		break
	case *expr.SQLReverseQuotedIdentifier:
		x, _ := x.(*expr.SQLReverseQuotedIdentifier)
		AcceptReverseQuotedIdentifier(visitor, x)
		break
	case *expr.SQLName:
		x, _ := x.(*expr.SQLName)
		AcceptName(visitor, x)
		break

	// ---------------------------- Identifier End ------------------------------------


	// ---------------------------- Variable Start ------------------------------------
	case *variable.SQLVariableExpr:
		x, _ := x.(*variable.SQLVariableExpr)
		AcceptVariableExpr(visitor, x)
		break
	case *variable.SQLBindVariableExpr:
		x, _ := x.(*variable.SQLBindVariableExpr)
		AcceptBindVariableExpr(visitor, x)
		break

	// ---------------------------- Variable End ------------------------------------


	// ---------------------------- Operator Start ------------------------------------
	case *operator.SQLUnaryOperatorExpr:
		x, _ := x.(*operator.SQLUnaryOperatorExpr)
		AcceptUnaryOperatorExpr(visitor, x)
		break
	case *operator.SQLBinaryOperatorExpr:
		x, _ := x.(*operator.SQLBinaryOperatorExpr)
		AcceptBinaryOperatorExpr(visitor, x)
		break

	// ---------------------------- Operator End ------------------------------------

	// ---------------------------- Expr Start ------------------------------------

	case *expr.SQLAllColumnExpr:
		x, _ := x.(*expr.SQLAllColumnExpr)
		AcceptAllColumnExpr(visitor, x)
		break

	// ---------------------------- Expr End ------------------------------------

	// ------------- Select Expr
	case *select_.SQLSelectQuery:
		x, _ := x.(*select_.SQLSelectQuery)
		AcceptSelectQuery(visitor, x)
		break
	case *select_.SQLParenSelectQuery:
		x, _ := x.(*select_.SQLParenSelectQuery)
		AcceptParenSelectQuery(visitor, x)
		break
	case *select_.SQLSelectUnionQuery:
		x, _ := x.(*select_.SQLSelectUnionQuery)
		AcceptSelectUnionQuery(visitor, x)
		break

	case *select_.SQLSelectElement:
		x, _ := x.(*select_.SQLSelectElement)
		AcceptSelectElement(visitor, x)
		break
	case *select_.SQLSelectTargetElement:
		x, _ := x.(*select_.SQLSelectTargetElement)
		AcceptSelectTargetElement(visitor, x)
		break

	case *select_.SQLFromClause:
		x, _ := x.(*select_.SQLFromClause)
		AcceptFromClause(visitor, x)
		break
	case *select_.SQLTableReference:
		x, _ := x.(*select_.SQLTableReference)
		AcceptTableReference(visitor, x)
		break
	case select_.SQLSubQueryTableReference:
		x, _ := x.(*select_.SQLSubQueryTableReference)
		AcceptSubQueryTableReference(visitor, x)
		break
	case *select_.SQLJoinTableReference:
		x, _ := x.(*select_.SQLJoinTableReference)
		AcceptJoinTableReference(visitor, x)
		break
	case *select_.SQLJoinOnCondition:
		x, _ := x.(*select_.SQLJoinOnCondition)
		AcceptJoinOnCondition(visitor, x)
		break
	case *select_.SQLJoinUsingCondition:
		x, _ := x.(*select_.SQLJoinUsingCondition)
		AcceptJoinUsingCondition(visitor, x)
		break

	case *select_.SQLWhereClause:
		x, _ := x.(*select_.SQLWhereClause)
		AcceptWhereClause(visitor, x)
		break

	case *select_.SQLGroupByHavingClause:
		x, _ := x.(*select_.SQLGroupByHavingClause)
		AcceptGroupByHavingClause(visitor, x)
		break
	case *select_.SQLHavingGroupByClause:
		x, _ := x.(*select_.SQLHavingGroupByClause)
		AcceptHavingGroupByClause(visitor, x)
		break

	case *select_.SQLOrderByClause:
		x, _ := x.(*select_.SQLOrderByClause)
		AcceptOrderByClause(visitor, x)
		break

	case *select_.SQLLimitOffsetClause:
		x, _ := x.(*select_.SQLLimitOffsetClause)
		AcceptLimitOffsetClause(visitor, x)
		break
	case *select_.SQLOffsetFetchClause:
		x, _ := x.(*select_.SQLOffsetFetchClause)
		AcceptOffsetFetchClause(visitor, x)
		break
	case *select_.SQLForUpdateClause:
		x, _ := x.(*select_.SQLForUpdateClause)
		AcceptForUpdateClause(visitor, x)
		break

		/********************************************* Statement ********************************************/

		// ------------- Select Statement
	case *statement.SQLSelectStatement:
		x, _ := x.(*statement.SQLSelectStatement)
		AcceptSelectStatement(visitor, x)
		break

	default:
		fmt.Println("TODO.", t)
		panic("TODO.")
	}
	visitor.AfterVisit(x)
}



func AcceptInOperator(visitor ISQLVisitor, x *operator.SQLInOperator) {
	// VisitStringLiteral(x)
}

func AcceptChildren(visitor ISQLVisitor, children []ast.ISQLObject) {
	if children == nil {
		return
	}

	for _, child := range children {
		Accept(visitor, child)
	}
}

// func AcceptChild(visitor ISQLVisitor, child ast.ISQLObject) {
// 	if ast.IsNil(child) {
// 		return
// 	}
// 	Accept(visitor, child)
// }