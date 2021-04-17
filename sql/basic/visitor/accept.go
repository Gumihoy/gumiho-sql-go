package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/common"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/condition"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/function"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/index"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/literal"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/operator"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	exprSequence "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/sequence"
	exprServer "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/server"
	exprTable "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	exprUser "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/user"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/variable"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/comment"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/database"
	statementFunction "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/function"
	statementIndex "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/index"
	statementPackage "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/package"
	statementPackageBody "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/packagebody"
	statementProcedure "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/procedure"
	statementRole "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/role"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/schema"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/sequence"
	statementServer "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/server"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/set"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/show"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/synonym"
	statementTable "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/table"
	statementTrigger "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/trigger"
	statementType "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/type"
	statementTypeBody "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/typebody"
	statementUser "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/user"
	statementView "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/view"
	"reflect"
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

	xType := reflect.TypeOf(x)

	visitor.BeforeVisit(visitor, x)
	switch x.(type) {

	// ---------------------------- Comment Start ------------------------------------
	case *ast.SQLMinusComment:
		x, _ := x.(*ast.SQLMinusComment)
		AcceptMinusComment(visitor, x)
	case *ast.SQLMultiLineComment:
		x, _ := x.(*ast.SQLMultiLineComment)
		AcceptMultiLineComment(visitor, x)
	case *ast.SQLSharpComment:
		x, _ := x.(*ast.SQLSharpComment)
		AcceptSharpComment(visitor, x)
	// ---------------------------- Comment End ------------------------------------

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
	case *literal.SQLHexadecimalLiteral:
		x, _ := x.(*literal.SQLHexadecimalLiteral)
		AcceptHexadecimalLiteral(visitor, x)
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

	case *expr.SQLDBLinkExpr:
		x, _ := x.(*expr.SQLDBLinkExpr)
		AcceptDBLinkExpr(visitor, x)

	// ---------------------------- Identifier End ------------------------------------

	// ---------------------------- Variable Start ------------------------------------
	case *variable.SQLVariableExpr:
		x, _ := x.(*variable.SQLVariableExpr)
		AcceptVariableExpr(visitor, x)
	case *variable.SQLAtVariableExpr:
		x, _ := x.(*variable.SQLAtVariableExpr)
		AcceptAtVariableExpr(visitor, x)
	case *variable.SQLAtAtVariableExpr:
		x, _ := x.(*variable.SQLAtAtVariableExpr)
		AcceptAtAtVariableExpr(visitor, x)
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

	// ---------------------------- Condition Start ------------------------------------
	case *condition.SQLIsCondition:
		x, _ := x.(*condition.SQLIsCondition)
		AcceptIsCondition(visitor, x)

	case *condition.SQLLikeCondition:
		x, _ := x.(*condition.SQLLikeCondition)
		AcceptLikeCondition(visitor, x)

	case *condition.SQLBetweenCondition:
		x, _ := x.(*condition.SQLBetweenCondition)
		AcceptBetweenCondition(visitor, x)

	case *condition.SQLInCondition:
		x, _ := x.(*condition.SQLInCondition)
		AcceptInCondition(visitor, x)

	case *condition.SQLIsASetCondition:
		x, _ := x.(*condition.SQLIsASetCondition)
		AcceptIsASetCondition(visitor, x)

	// ---------------------------- Condition End ------------------------------------

	// ---------------------------- Expr Start ------------------------------------
	case *expr.SQLAllColumnExpr:
		x, _ := x.(*expr.SQLAllColumnExpr)
		AcceptAllColumnExpr(visitor, x)

	case *expr.SQLNullExpr:
		x, _ := x.(*expr.SQLNullExpr)
		AcceptNullExpr(visitor, x)

	case *expr.SQLListExpr:
		x, _ := x.(*expr.SQLListExpr)
		AcceptListExpr(visitor, x)

	case *common.SQLSubQueryExpr:
		x, _ := x.(*common.SQLSubQueryExpr)
		AcceptSubQuery(visitor, x)


	case *expr.SQLAssignExpr:
		x, _ := x.(*expr.SQLAssignExpr)
		AcceptAssignExpr(visitor, x)
		break

	// ---------------------------- Expr End ------------------------------------

	// ---------------------------- Function Start ------------------------------------

	case *function.SQLMethodInvocation:
		x, _ := x.(*function.SQLMethodInvocation)
		AcceptMethodInvocation(visitor, x)
		break

	case *function.SQLStaticMethodInvocation:
		x, _ := x.(*function.SQLStaticMethodInvocation)
		AcceptStaticMethodInvocation(visitor, x)
		break

	case *function.SQLCastFunctionArgument:
		x, _ := x.(*function.SQLCastFunctionArgument)
		AcceptCastFunctionArgument(visitor, x)
		break
	// ---------------------------- Function End ------------------------------------

	// ---------------------------- DataType Start ------------------------------------
	case *datatype.SQLDataType:
		x, _ := x.(*datatype.SQLDataType)
		AcceptDataType(visitor, x)

	case *datatype.SQLIntervalDataType:
		x, _ := x.(*datatype.SQLIntervalDataType)
		AcceptIntervalDataType(visitor, x)

	case *datatype.SQLIntervalDataTypeField:
		x, _ := x.(*datatype.SQLIntervalDataTypeField)
		AcceptIntervalDataTypeField(visitor, x)

	case *datatype.SQLDateDataType:
		x, _ := x.(*datatype.SQLDateDataType)
		AcceptDateDataType(visitor, x)
		break
	case *datatype.SQLDateTimeDataType:
		x, _ := x.(*datatype.SQLDateTimeDataType)
		AcceptDateTimeDataType(visitor, x)
		break
	case *datatype.SQLTimeDataType:
		x, _ := x.(*datatype.SQLTimeDataType)
		AcceptTimeDataType(visitor, x)
		break
	case *datatype.SQLTimestampDataType:
		x, _ := x.(*datatype.SQLTimestampDataType)
		AcceptTimestampDataType(visitor, x)
		break
	// ---------------------------- DataType End ------------------------------------


	// -------------------------- DDL Expr --------------------------
	case *statement.SQLEditionAbleExpr:
		x, _ := x.(*statement.SQLEditionAbleExpr)
		AcceptEditionAbleExpr(visitor, x)
	case *statement.SQLNonEditionAbleExpr:
		x, _ := x.(*statement.SQLNonEditionAbleExpr)
		AcceptNonEditionAbleExpr(visitor, x)
	case *statement.SQLCompileExpr:
		x, _ := x.(*statement.SQLCompileExpr)
		AcceptCompileExpr(visitor, x)

		// ------------- Index Expr
		case *index.SQLIndexColumn:
			x, _ := x.(*index.SQLIndexColumn)
			AcceptIndexColumn(visitor, x)

		// ------------- Sequence Expr
	case *exprSequence.SQLIncrementBySequenceOption:
		x, _ := x.(*exprSequence.SQLIncrementBySequenceOption)
		AcceptIncrementBySequenceOption(visitor, x)
	case *exprSequence.SQLStartWithSequenceOption:
		x, _ := x.(*exprSequence.SQLStartWithSequenceOption)
		AcceptStartWithSequenceOption(visitor, x)
	case *exprSequence.SQLMaxValueSequenceOption:
		x, _ := x.(*exprSequence.SQLMaxValueSequenceOption)
		AcceptMaxValueSequenceOption(visitor, x)
	case *exprSequence.SQLNoMaxValueSequenceOption:
		x, _ := x.(*exprSequence.SQLNoMaxValueSequenceOption)
		AcceptNoMaxValueSequenceOption(visitor, x)
	case *exprSequence.SQLMinValueSequenceOption:
		x, _ := x.(*exprSequence.SQLMinValueSequenceOption)
		AcceptMinValueSequenceOption(visitor, x)
	case *exprSequence.SQLNoMinValueSequenceOption:
		x, _ := x.(*exprSequence.SQLNoMinValueSequenceOption)
		AcceptNoMinValueSequenceOption(visitor, x)
	case *exprSequence.SQLCycleSequenceOption:
		x, _ := x.(*exprSequence.SQLCycleSequenceOption)
		AcceptCycleSequenceOption(visitor, x)
	case *exprSequence.SQLNoCycleSequenceOption:
		x, _ := x.(*exprSequence.SQLNoCycleSequenceOption)
		AcceptNoCycleSequenceOption(visitor, x)
	case *exprSequence.SQLCacheSequenceOption:
		x, _ := x.(*exprSequence.SQLCacheSequenceOption)
		AcceptCacheSequenceOption(visitor, x)
	case *exprSequence.SQLNoCacheSequenceOption:
		x, _ := x.(*exprSequence.SQLNoCacheSequenceOption)
		AcceptNoCacheSequenceOption(visitor, x)
	case *exprSequence.SQLOrderSequenceOption:
		x, _ := x.(*exprSequence.SQLOrderSequenceOption)
		AcceptOrderSequenceOption(visitor, x)
	case *exprSequence.SQLNoOrderSequenceOption:
		x, _ := x.(*exprSequence.SQLNoOrderSequenceOption)
		AcceptNoOrderSequenceOption(visitor, x)
	case *exprSequence.SQLKeepSequenceOption:
		x, _ := x.(*exprSequence.SQLKeepSequenceOption)
		AcceptKeepSequenceOption(visitor, x)
	case *exprSequence.SQLNoKeepSequenceOption:
		x, _ := x.(*exprSequence.SQLNoKeepSequenceOption)
		AcceptNoKeepSequenceOption(visitor, x)
	case *exprSequence.SQLScaleSequenceOption:
		x, _ := x.(*exprSequence.SQLScaleSequenceOption)
		AcceptScaleSequenceOption(visitor, x)
	case *exprSequence.SQLNoScaleSequenceOption:
		x, _ := x.(*exprSequence.SQLNoScaleSequenceOption)
		AcceptNoScaleSequenceOption(visitor, x)
	case *exprSequence.SQLSessionSequenceOption:
		x, _ := x.(*exprSequence.SQLSessionSequenceOption)
		AcceptSessionSequenceOption(visitor, x)
	case *exprSequence.SQLGlobalSequenceOption:
		x, _ := x.(*exprSequence.SQLGlobalSequenceOption)
		AcceptGlobalSequenceOption(visitor, x)


		// ------------- Server Expr
	case *exprServer.SQLHostOption:
		x, _ := x.(*exprServer.SQLHostOption)
		AcceptHostOption(visitor, x)
	case *exprServer.SQLDatabaseOption:
		x, _ := x.(*exprServer.SQLDatabaseOption)
		AcceptDatabaseOption(visitor, x)
	case *exprServer.SQLUserOption:
		x, _ := x.(*exprServer.SQLUserOption)
		AcceptUserOption(visitor, x)
	case *exprServer.SQLPasswordOption:
		x, _ := x.(*exprServer.SQLPasswordOption)
		AcceptPasswordOption(visitor, x)
	case *exprServer.SQLOwnerOption:
		x, _ := x.(*exprServer.SQLOwnerOption)
		AcceptOwnerOption(visitor, x)
	case *exprServer.SQLSocketOption:
		x, _ := x.(*exprServer.SQLSocketOption)
		AcceptSocketOption(visitor, x)
	case *exprServer.SQLPortOption:
		x, _ := x.(*exprServer.SQLPortOption)
		AcceptPortOption(visitor, x)

	// ------------- Table Expr
	case *exprTable.SQLTableColumn:
		x, _ := x.(*exprTable.SQLTableColumn)
		AcceptTableColumn(visitor, x)

	case *exprTable.SQLPrimaryKeyTableConstraint:
		x, _ := x.(*exprTable.SQLPrimaryKeyTableConstraint)
		AcceptPrimaryKeyTableConstraint(visitor, x)
	case *exprTable.SQLUniqueTableConstraint:
		x, _ := x.(*exprTable.SQLUniqueTableConstraint)
		AcceptUniqueTableConstraint(visitor, x)
	case *exprTable.SQLUniqueIndexTableConstraint:
		x, _ := x.(*exprTable.SQLUniqueIndexTableConstraint)
		AcceptUniqueIndexTableConstraint(visitor, x)
	case *exprTable.SQLUniqueKeyTableConstraint:
		x, _ := x.(*exprTable.SQLUniqueKeyTableConstraint)
		AcceptUniqueKeyTableConstraint(visitor, x)
	case *exprTable.SQLForeignKeyTableConstraint:
		x, _ := x.(*exprTable.SQLForeignKeyTableConstraint)
		AcceptForeignKeyTableConstraint(visitor, x)
	case *exprTable.SQLCheckTableConstraint:
		x, _ := x.(*exprTable.SQLCheckTableConstraint)
		AcceptCheckTableConstraint(visitor, x)

	case *exprTable.SQLNullColumnConstraint:
		x, _ := x.(*exprTable.SQLNullColumnConstraint)
		AcceptNullColumnConstraint(visitor, x)
	case *exprTable.SQLNotNullColumnConstraint:
		x, _ := x.(*exprTable.SQLNotNullColumnConstraint)
		AcceptNotNullColumnConstraint(visitor, x)
	case *exprTable.SQLPrimaryKeyColumnConstraint:
		x, _ := x.(*exprTable.SQLPrimaryKeyColumnConstraint)
		AcceptPrimaryKeyColumnConstraint(visitor, x)
	case *exprTable.SQLKeyColumnConstraint:
		x, _ := x.(*exprTable.SQLKeyColumnConstraint)
		AcceptKeyColumnConstraint(visitor, x)
	case *exprTable.SQLUniqueColumnConstraint:
		x, _ := x.(*exprTable.SQLUniqueColumnConstraint)
		AcceptUniqueColumnConstraint(visitor, x)
	case *exprTable.SQLCheckColumnConstraint:
		x, _ := x.(*exprTable.SQLCheckColumnConstraint)
		AcceptCheckColumnConstraint(visitor, x)

	case *exprTable.SQLDefaultClause:
		x, _ := x.(*exprTable.SQLDefaultClause)
		AcceptDefaultClause(visitor, x)
	case *exprTable.SQLAutoIncrementExpr:
		x, _ := x.(*exprTable.SQLAutoIncrementExpr)
		AcceptAutoIncrementExpr(visitor, x)
	case *exprTable.SQLVisibleExpr:
		x, _ := x.(*exprTable.SQLVisibleExpr)
		AcceptVisibleExpr(visitor, x)
	case *exprTable.SQLInvisibleExpr:
		x, _ := x.(*exprTable.SQLInvisibleExpr)
		AcceptInvisibleExpr(visitor, x)
	case *exprTable.SQLCommentExpr:
		x, _ := x.(*exprTable.SQLCommentExpr)
		AcceptCommentExpr(visitor, x)

	case *exprTable.SQLTableLikeClause:
		x, _ := x.(*exprTable.SQLTableLikeClause)
		AcceptTableLikeClause(visitor, x)

	case *exprTable.SQLCharsetAssignExpr:
		x, _ := x.(*exprTable.SQLCharsetAssignExpr)
		AcceptCharsetAssignExpr(visitor, x)
	case *exprTable.SQLCharacterSetAssignExpr:
		x, _ := x.(*exprTable.SQLCharacterSetAssignExpr)
		AcceptCharacterSetAssignExpr(visitor, x)

	case *exprTable.SQLPartitionByHash:
		x, _ := x.(*exprTable.SQLPartitionByHash)
		AcceptPartitionByHash(visitor, x)
	case *exprTable.SQLPartitionByKey:
		x, _ := x.(*exprTable.SQLPartitionByKey)
		AcceptPartitionByKey(visitor, x)
	case *exprTable.SQLPartitionByRange:
		x, _ := x.(*exprTable.SQLPartitionByRange)
		AcceptPartitionByRange(visitor, x)
	case *exprTable.SQLPartitionByList:
		x, _ := x.(*exprTable.SQLPartitionByList)
		AcceptPartitionByList(visitor, x)

	case *exprTable.SQLSubPartitionByHash:
		x, _ := x.(*exprTable.SQLSubPartitionByHash)
		AcceptSubPartitionByHash(visitor, x)
	case *exprTable.SQLSubPartitionByKey:
		x, _ := x.(*exprTable.SQLSubPartitionByKey)
		AcceptSubPartitionByKey(visitor, x)

	case *exprTable.SQLPartitionDefinition:
		x, _ := x.(*exprTable.SQLPartitionDefinition)
		AcceptPartitionDefinition(visitor, x)

	case *exprTable.SQLPartitionValuesLessThan:
		x, _ := x.(*exprTable.SQLPartitionValuesLessThan)
		AcceptPartitionValuesLessThan(visitor, x)
	case *exprTable.SQLPartitionValuesLessThanMaxValue:
		x, _ := x.(*exprTable.SQLPartitionValuesLessThanMaxValue)
		AcceptPartitionValuesLessThanMaxValue(visitor, x)
	case *exprTable.SQLPartitionValuesIn:
		x, _ := x.(*exprTable.SQLPartitionValuesIn)
		AcceptPartitionValuesIn(visitor, x)

	case *exprTable.SQLSubPartitionDefinition:
		x, _ := x.(*exprTable.SQLSubPartitionDefinition)
		AcceptSubPartitionDefinition(visitor, x)

	// ------------- Alter OnTable Expr
	case *exprTable.SQLAddColumnAlterTableAction:
		x, _ := x.(*exprTable.SQLAddColumnAlterTableAction)
		AcceptAddColumnAlterTableAction(visitor, x)
	case *exprTable.SQLAlterColumnAlterTableAction:
		x, _ := x.(*exprTable.SQLAlterColumnAlterTableAction)
		AcceptAlterColumnAlterTableAction(visitor, x)
	case *exprTable.SQLDropColumnAlterTableAction:
		x, _ := x.(*exprTable.SQLDropColumnAlterTableAction)
		AcceptDropColumnAlterTableAction(visitor, x)
	case *exprTable.SQLAddTableConstraintAlterTableAction:
		x, _ := x.(*exprTable.SQLAddTableConstraintAlterTableAction)
		AcceptAddTableConstraintAlterTableAction(visitor, x)

	case *exprTable.SQLDropIndexAlterTableAction:
		x, _ := x.(*exprTable.SQLDropIndexAlterTableAction)
		AcceptDropIndexAlterTableAction(visitor, x)
	case *exprTable.SQLDropKeyAlterTableAction:
		x, _ := x.(*exprTable.SQLDropKeyAlterTableAction)
		AcceptDropKeyAlterTableAction(visitor, x)

	case *exprTable.SQLDropTableConstraintAlterTableAction:
		x, _ := x.(*exprTable.SQLDropTableConstraintAlterTableAction)
		AcceptDropConstraintTableConstraintAlterTableAction(visitor, x)

	case *exprTable.SQLDropPrimaryKeyTableConstraintAlterTableAction:
		x, _ := x.(*exprTable.SQLDropPrimaryKeyTableConstraintAlterTableAction)
		AcceptDropPrimaryKeyTableConstraintAlterTableAction(visitor, x)
	case *exprTable.SQLDropUniqueTableConstraintAlterTableAction:
		x, _ := x.(*exprTable.SQLDropUniqueTableConstraintAlterTableAction)
		AcceptDropUniqueTableConstraintAlterTableAction(visitor, x)
	case *exprTable.SQLDropForeignKeyTableConstraintAlterTableAction:
		x, _ := x.(*exprTable.SQLDropForeignKeyTableConstraintAlterTableAction)
		AcceptDropForeignKeyTableConstraintAlterTableAction(visitor, x)
	case *exprTable.SQLDropCheckTableConstraintAlterTableAction:
		x, _ := x.(*exprTable.SQLDropCheckTableConstraintAlterTableAction)
		AcceptDropCheckTableConstraintAlterTableAction(visitor, x)


	case *exprTable.SQLAddPartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLAddPartitionAlterTableAction)
		AcceptAddPartitionAlterTableAction(visitor, x)
	case *exprTable.SQLDropPartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLDropPartitionAlterTableAction)
		AcceptDropPartitionAlterTableAction(visitor, x)
	case *exprTable.SQLDiscardPartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLDiscardPartitionAlterTableAction)
		AcceptDiscardPartitionAlterTableAction(visitor, x)
	case *exprTable.SQLImportPartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLImportPartitionAlterTableAction)
		AcceptImportPartitionAlterTableAction(visitor, x)
	case *exprTable.SQLTruncatePartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLTruncatePartitionAlterTableAction)
		AcceptTruncatePartitionAlterTableAction(visitor, x)
	case *exprTable.SQLCoalescePartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLCoalescePartitionAlterTableAction)
		AcceptCoalescePartitionAlterTableAction(visitor, x)
	case *exprTable.SQLReorganizePartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLReorganizePartitionAlterTableAction)
		AcceptReorganizePartitionAlterTableAction(visitor, x)
	case *exprTable.SQLExchangePartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLExchangePartitionAlterTableAction)
		AcceptExchangePartitionAlterTableAction(visitor, x)
	case *exprTable.SQLAnalyzePartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLAnalyzePartitionAlterTableAction)
		AcceptAnalyzePartitionAlterTableAction(visitor, x)
	case *exprTable.SQLCheckPartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLCheckPartitionAlterTableAction)
		AcceptCheckPartitionAlterTableAction(visitor, x)
	case *exprTable.SQLOptimizePartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLOptimizePartitionAlterTableAction)
		AcceptOptimizePartitionAlterTableAction(visitor, x)
	case *exprTable.SQLRebuildPartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLRebuildPartitionAlterTableAction)
		AcceptRebuildPartitionAlterTableAction(visitor, x)
	case *exprTable.SQLRepairPartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLRepairPartitionAlterTableAction)
		AcceptRepairPartitionAlterTableAction(visitor, x)
	case *exprTable.SQLRemovePartitionAlterTableAction:
		x, _ := x.(*exprTable.SQLRemovePartitionAlterTableAction)
		AcceptRemovePartitionAlterTableAction(visitor, x)



		// ------------- Drop Table Expr
	case *exprTable.SQLDropTableStatementRestrictOption:
		x, _ := x.(*exprTable.SQLDropTableStatementRestrictOption)
		AcceptDropTableStatementRestrictOption(visitor, x)
		break
	case *exprTable.SQLDropTableStatementCascadeOption:
		x, _ := x.(*exprTable.SQLDropTableStatementCascadeOption)
		AcceptDropTableStatementCascadeOption(visitor, x)
		break

		// ------------- User Expr
	case *exprUser.SQLUserName:
		x, _ := x.(*exprUser.SQLUserName)
		AcceptUserName(visitor, x)

	case *exprUser.SQLIdentifiedByAuthOption:
		x, _ := x.(*exprUser.SQLIdentifiedByAuthOption)
		AcceptIdentifiedByAuthOption(visitor, x)
	case *exprUser.SQLIdentifiedByPasswordAuthOption:
		x, _ := x.(*exprUser.SQLIdentifiedByPasswordAuthOption)
		AcceptIdentifiedByPasswordAuthOption(visitor, x)
	case *exprUser.SQLIdentifiedByRandomPasswordAuthOption:
		x, _ := x.(*exprUser.SQLIdentifiedByRandomPasswordAuthOption)
		AcceptIdentifiedByRandomPasswordAuthOption(visitor, x)
	case *exprUser.SQLIdentifiedWithAuthOption:
		x, _ := x.(*exprUser.SQLIdentifiedWithAuthOption)
		AcceptIdentifiedWithAuthOption(visitor, x)
	case *exprUser.SQLIdentifiedWithByAuthOption:
		x, _ := x.(*exprUser.SQLIdentifiedWithByAuthOption)
		AcceptIdentifiedWithByAuthOption(visitor, x)
	case *exprUser.SQLIdentifiedWithByRandomPasswordAuthOption:
		x, _ := x.(*exprUser.SQLIdentifiedWithByRandomPasswordAuthOption)
		AcceptIdentifiedWithByRandomPasswordAuthOption(visitor, x)
	case *exprUser.SQLIdentifiedWithAsAuthOption:
		x, _ := x.(*exprUser.SQLIdentifiedWithAsAuthOption)
		AcceptIdentifiedWithAsAuthOption(visitor, x)

		// ------------- View Expr

	case *view.SQLViewColumn:
		x, _ := x.(*view.SQLViewColumn)
		AcceptViewColumn(visitor, x)

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

	case *select_.SQLWithClause:
		x, _ := x.(*select_.SQLWithClause)
		AcceptWithClause(visitor, x)
		break
	case *select_.SQLSubQueryFactoringClause:
		x, _ := x.(*select_.SQLSubQueryFactoringClause)
		AcceptSubQueryFactoringClause(visitor, x)
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

	case *select_.SQLTableReference:
		x, _ := x.(*select_.SQLTableReference)
		AcceptTableReference(visitor, x)

	case *select_.SQLPartitionClause:
		x, _ := x.(*select_.SQLPartitionClause)
		AcceptPartitionClause(visitor, x)
	case *select_.SQLPartitionForClause:
		x, _ := x.(*select_.SQLPartitionForClause)
		AcceptPartitionForClause(visitor, x)
	case *select_.SQLSubPartitionClause:
		x, _ := x.(*select_.SQLSubPartitionClause)
		AcceptSubPartitionClause(visitor, x)
	case *select_.SQLSubPartitionForClause:
		x, _ := x.(*select_.SQLSubPartitionForClause)
		AcceptSubPartitionForClause(visitor, x)

	case *select_.SQLSampleClause:
		x, _ := x.(*select_.SQLSampleClause)
		AcceptSampleClause(visitor, x)

		case *select_.SQLOJTableReference:
		x, _ := x.(*select_.SQLOJTableReference)
		AcceptOJTableReference(visitor, x)

	case *select_.SQLSubQueryTableReference:
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

	case *select_.SQLHierarchicalQueryClauseConnectBy:
		x, _ := x.(*select_.SQLHierarchicalQueryClauseConnectBy)
		AcceptHierarchicalQueryClauseConnectBy(visitor, x)
	case *select_.SQLHierarchicalQueryClauseStartWith:
		x, _ := x.(*select_.SQLHierarchicalQueryClauseStartWith)
		AcceptHierarchicalQueryClauseStartWith(visitor, x)

	case *select_.SQLGroupByHavingClause:
		x, _ := x.(*select_.SQLGroupByHavingClause)
		AcceptGroupByHavingClause(visitor, x)
		break
	case *select_.SQLHavingGroupByClause:
		x, _ := x.(*select_.SQLHavingGroupByClause)
		AcceptHavingGroupByClause(visitor, x)
	case *select_.SQLGroupByElement:
		x, _ := x.(*select_.SQLGroupByElement)
		AcceptGroupByElement(visitor, x)

	case *select_.SQLOrderByClause:
		x, _ := x.(*select_.SQLOrderByClause)
		AcceptOrderByClause(visitor, x)
	case *select_.SQLOrderByElement:
		x, _ := x.(*select_.SQLOrderByElement)
		AcceptOrderByElement(visitor, x)

	case *select_.SQLLimitOffsetClause:
		x, _ := x.(*select_.SQLLimitOffsetClause)
		AcceptLimitOffsetClause(visitor, x)
		break
	case *select_.SQLOffsetFetchClause:
		x, _ := x.(*select_.SQLOffsetFetchClause)
		AcceptOffsetFetchClause(visitor, x)

	case *select_.SQLForUpdateClause:
		x, _ := x.(*select_.SQLForUpdateClause)
		AcceptForUpdateClause(visitor, x)
	case *select_.SQLForShareClause:
		x, _ := x.(*select_.SQLForShareClause)
		AcceptForShareClause(visitor, x)
	case *select_.SQLLockInShareModeClause:
		x, _ := x.(*select_.SQLLockInShareModeClause)
		AcceptLockInShareModeClause(visitor, x)


		/********************************************* Statement ********************************************/

		// -------------------------- Comment
	case *comment.SQLCommentOnAuditPolicyStatement:
		x, _ := x.(*comment.SQLCommentOnAuditPolicyStatement)
		AcceptCommentOnAuditPolicyStatement(visitor, x)
	case *comment.SQLCommentOnColumnStatement:
		x, _ := x.(*comment.SQLCommentOnColumnStatement)
		AcceptCommentOnColumnStatement(visitor, x)
	case *comment.SQLCommentOnEditionStatement:
		x, _ := x.(*comment.SQLCommentOnEditionStatement)
		AcceptCommentOnEditionStatement(visitor, x)
	case *comment.SQLCommentOnIndextypeStatement:
		x, _ := x.(*comment.SQLCommentOnIndextypeStatement)
		AcceptCommentOnIndextypeStatement(visitor, x)
	case *comment.SQLCommentOnMaterializedViewStatement:
		x, _ := x.(*comment.SQLCommentOnMaterializedViewStatement)
		AcceptCommentOnMaterializedViewStatement(visitor, x)
	case *comment.SQLCommentOnMiningModelStatement:
		x, _ := x.(*comment.SQLCommentOnMiningModelStatement)
		AcceptCommentOnMiningModelStatement(visitor, x)
	case *comment.SQLCommentOnOperatorStatement:
		x, _ := x.(*comment.SQLCommentOnOperatorStatement)
		AcceptCommentOnOperatorStatement(visitor, x)
	case *comment.SQLCommentOnTableStatement:
		x, _ := x.(*comment.SQLCommentOnTableStatement)
		AcceptCommentOnTableStatement(visitor, x)

		// -------------------------- DDL --------------------------
		// -------------------------- Database
	case *database.SQLAlterDatabaseStatement:
		x, _ := x.(*database.SQLAlterDatabaseStatement)
		AcceptAlterDatabaseStatement(visitor, x)
	case *database.SQLCreateDatabaseStatement:
		x, _ := x.(*database.SQLCreateDatabaseStatement)
		AcceptCreateDatabaseStatement(visitor, x)
	case *database.SQLDropDatabaseStatement:
		x, _ := x.(*database.SQLDropDatabaseStatement)
		AcceptDropDatabaseStatement(visitor, x)

		// -------------------------- Function
	case *statementFunction.SQLAlterFunctionStatement:
		x, _ := x.(*statementFunction.SQLAlterFunctionStatement)
		AcceptAlterFunctionStatement(visitor, x)
	case *statementFunction.SQLCreateFunctionStatement:
		x, _ := x.(*statementFunction.SQLCreateFunctionStatement)
		AcceptCreateFunctionStatement(visitor, x)
	case *statementFunction.SQLDropFunctionStatement:
		x, _ := x.(*statementFunction.SQLDropFunctionStatement)
		AcceptDropFunctionStatement(visitor, x)

		// -------------------------- Index
	case *statementIndex.SQLAlterIndexStatement:
		x, _ := x.(*statementIndex.SQLAlterIndexStatement)
		AcceptAlterIndexStatement(visitor, x)
	case *statementIndex.SQLCreateIndexStatement:
		x, _ := x.(*statementIndex.SQLCreateIndexStatement)
		AcceptCreateIndexStatement(visitor, x)
	case *statementIndex.SQLDropIndexStatement:
		x, _ := x.(*statementIndex.SQLDropIndexStatement)
		AcceptDropIndexStatement(visitor, x)

		// -------------------------- Package
	case *statementPackage.SQLAlterPackageStatement:
		x, _ := x.(*statementPackage.SQLAlterPackageStatement)
		AcceptAlterPackageStatement(visitor, x)
	case *statementPackage.SQLCreatePackageStatement:
		x, _ := x.(*statementPackage.SQLCreatePackageStatement)
		AcceptCreatePackageStatement(visitor, x)
	case *statementPackage.SQLDropPackageStatement:
		x, _ := x.(*statementPackage.SQLDropPackageStatement)
		AcceptDropPackageStatement(visitor, x)

		// -------------------------- Package Body
	case *statementPackageBody.SQLAlterPackageBodyStatement:
		x, _ := x.(*statementPackageBody.SQLAlterPackageBodyStatement)
		AcceptAlterPackageBodyStatement(visitor, x)
	case *statementPackageBody.SQLCreatePackageBoydStatement:
		x, _ := x.(*statementPackageBody.SQLCreatePackageBoydStatement)
		AcceptCreatePackageBoydStatement(visitor, x)
	case *statementPackageBody.SQLDropPackageBodyStatement:
		x, _ := x.(*statementPackageBody.SQLDropPackageBodyStatement)
		AcceptDropPackageBodyStatement(visitor, x)


		// -------------------------- Procedure
	case *statementProcedure.SQLAlterProcedureStatement:
		x, _ := x.(*statementProcedure.SQLAlterProcedureStatement)
		AcceptAlterProcedureStatement(visitor, x)
	case *statementProcedure.SQLCreateProcedureStatement:
		x, _ := x.(*statementProcedure.SQLCreateProcedureStatement)
		AcceptCreateProcedureStatement(visitor, x)
	case *statementProcedure.SQLDropProcedureStatement:
		x, _ := x.(*statementProcedure.SQLDropProcedureStatement)
		AcceptDropProcedureStatement(visitor, x)

		// -------------------------- Role
	case *statementRole.SQLAlterRoleStatement:
		x, _ := x.(*statementRole.SQLAlterRoleStatement)
		AcceptAlterRoleStatement(visitor, x)
	case *statementRole.SQLCreateRoleStatement:
		x, _ := x.(*statementRole.SQLCreateRoleStatement)
		AcceptCreateRoleStatement(visitor, x)
	case *statementRole.SQLDropRoleStatement:
		x, _ := x.(*statementRole.SQLDropRoleStatement)
		AcceptDropRoleStatement(visitor, x)

		// -------------------------- Schema
	case *schema.SQLAlterSchemaStatement:
		x, _ := x.(*schema.SQLAlterSchemaStatement)
		AcceptAlterSchemaStatement(visitor, x)
	case *schema.SQLCreateSchemaStatement:
		x, _ := x.(*schema.SQLCreateSchemaStatement)
		AcceptCreateSchemaStatement(visitor, x)
	case *schema.SQLDropSchemaStatement:
		x, _ := x.(*schema.SQLDropSchemaStatement)
		AcceptDropSchemaStatement(visitor, x)

		// -------------------------- Sequence
	case *sequence.SQLAlterSequenceStatement:
		x, _ := x.(*sequence.SQLAlterSequenceStatement)
		AcceptAlterSequenceStatement(visitor, x)
	case *sequence.SQLCreateSequenceStatement:
		x, _ := x.(*sequence.SQLCreateSequenceStatement)
		AcceptCreateSequenceStatement(visitor, x)
	case *sequence.SQLDropSequenceStatement:
		x, _ := x.(*sequence.SQLDropSequenceStatement)
		AcceptDropSequenceStatement(visitor, x)

		// -------------------------- Server
	case *statementServer.SQLAlterServerStatement:
		x, _ := x.(*statementServer.SQLAlterServerStatement)
		AcceptAlterServerStatement(visitor, x)
	case *statementServer.SQLCreateServerStatement:
		x, _ := x.(*statementServer.SQLCreateServerStatement)
		AcceptCreateServerStatement(visitor, x)
	case *statementServer.SQLDropServerStatement:
		x, _ := x.(*statementServer.SQLDropServerStatement)
		AcceptDropServerStatement(visitor, x)

		// -------------------------- Synonym
	case *synonym.SQLAlterSynonymStatement:
		x, _ := x.(*synonym.SQLAlterSynonymStatement)
		AcceptAlterSynonymStatement(visitor, x)
	case *synonym.SQLCreateSynonymStatement:
		x, _ := x.(*synonym.SQLCreateSynonymStatement)
		AcceptCreateSynonymStatement(visitor, x)
	case *synonym.SQLDropSynonymStatement:
		x, _ := x.(*synonym.SQLDropSynonymStatement)
		AcceptDropSynonymStatement(visitor, x)

		// -------------------------- OnTable
	case *statementTable.SQLAlterTableStatement:
		x, _ := x.(*statementTable.SQLAlterTableStatement)
		AcceptAlterTableStatement(visitor, x)
	case *statementTable.SQLCreateTableStatement:
		x, _ := x.(*statementTable.SQLCreateTableStatement)
		AcceptCreateTableStatement(visitor, x)
	case *statementTable.SQLDropTableStatement:
		x, _ := x.(*statementTable.SQLDropTableStatement)
		AcceptDropTableStatement(visitor, x)

		// -------------------------- Trigger
	case *statementTrigger.SQLAlterTriggerStatement:
		x, _ := x.(*statementTrigger.SQLAlterTriggerStatement)
		AcceptAlterTriggerStatement(visitor, x)
	case *statementTrigger.SQLCreateTriggerStatement:
		x, _ := x.(*statementTrigger.SQLCreateTriggerStatement)
		AcceptCreateTriggerStatement(visitor, x)
	case *statementTrigger.SQLDropTriggerStatement:
		x, _ := x.(*statementTrigger.SQLDropTriggerStatement)
		AcceptDropTriggerStatement(visitor, x)

		// -------------------------- Type
	case *statementType.SQLAlterTypeStatement:
		x, _ := x.(*statementType.SQLAlterTypeStatement)
		AcceptAlterTypeStatement(visitor, x)
	case *statementType.SQLCreateTypeStatement:
		x, _ := x.(*statementType.SQLCreateTypeStatement)
		AcceptCreateTypeStatement(visitor, x)
	case *statementType.SQLDropTypeStatement:
		x, _ := x.(*statementType.SQLDropTypeStatement)
		AcceptDropTypeStatement(visitor, x)

		// -------------------------- Type Body
	case *statementTypeBody.SQLAlterTypeBodyStatement:
		x, _ := x.(*statementTypeBody.SQLAlterTypeBodyStatement)
		AcceptAlterTypeBodyStatement(visitor, x)
	case *statementTypeBody.SQLCreateTypeBodyStatement:
		x, _ := x.(*statementTypeBody.SQLCreateTypeBodyStatement)
		AcceptCreateTypeBodyStatement(visitor, x)
	case *statementTypeBody.SQLDropTypeBodyStatement:
		x, _ := x.(*statementTypeBody.SQLDropTypeBodyStatement)
		AcceptDropTypeBodyStatement(visitor, x)

		// -------------------------- User
	case *statementUser.SQLAlterUserStatement:
			x, _ := x.(*statementUser.SQLAlterUserStatement)
			AcceptAlterUserStatement(visitor, x)
	case *statementUser.SQLCreateUserStatement:
		x, _ := x.(*statementUser.SQLCreateUserStatement)
		AcceptCreateUserStatement(visitor, x)
	case *statementUser.SQLDropUserStatement:
		x, _ := x.(*statementUser.SQLDropUserStatement)
		AcceptDropUserStatement(visitor, x)

		// -------------------------- View
	case *statementView.SQLAlterViewStatement:
		x, _ := x.(*statementView.SQLAlterViewStatement)
		AcceptAlterViewStatement(visitor, x)
	case *statementView.SQLCreateViewStatement:
		x, _ := x.(*statementView.SQLCreateViewStatement)
		AcceptCreateViewStatement(visitor, x)
	case *statementView.SQLDropViewStatement:
		x, _ := x.(*statementView.SQLDropViewStatement)
		AcceptDropViewStatement(visitor, x)

		// -------------------------- DML --------------------------

		// ------------- Delete Statement
	case *statement.SQLDeleteStatement:
		x, _ := x.(*statement.SQLDeleteStatement)
		AcceptDeleteStatement(visitor, x)

	case *statement.SQLInsertStatement:
		x, _ := x.(*statement.SQLInsertStatement)
		AcceptInsertStatement(visitor, x)

		// ------------- Select Statement
	case *statement.SQLSelectStatement:
		x, _ := x.(*statement.SQLSelectStatement)
		AcceptSelectStatement(visitor, x)

	case *statement.SQLUpdateStatement:
		x, _ := x.(*statement.SQLUpdateStatement)
		AcceptUpdateStatement(visitor, x)

		// -------------------------- SET --------------------------
	case *set.SQLSetVariableAssignmentStatement:
		x, _ := x.(*set.SQLSetVariableAssignmentStatement)
		AcceptSetVariableAssignmentStatement(visitor, x)

	case *set.SQLSetCharacterSetStatement:
		x, _ := x.(*set.SQLSetCharacterSetStatement)
		AcceptSetCharacterSetStatement(visitor, x)

	case *set.SQLSetCharsetStatement:
		x, _ := x.(*set.SQLSetCharsetStatement)
		AcceptSetCharsetStatement(visitor, x)

	case *set.SQLSetNamesStatement:
		x, _ := x.(*set.SQLSetNamesStatement)
		AcceptSetNamesStatement(visitor, x)

		// -------------------------- SHOW --------------------------
	case *show.SQLShowCreateDatabaseStatement:
		x, _ := x.(*show.SQLShowCreateDatabaseStatement)
		AcceptShowCreateDatabaseStatement(visitor, x)

	case *show.SQLShowCreateEventStatement:
		x, _ := x.(*show.SQLShowCreateEventStatement)
		AcceptShowCreateEventStatement(visitor, x)

	case *show.SQLShowCreateFunctionStatement:
		x, _ := x.(*show.SQLShowCreateFunctionStatement)
		AcceptShowCreateFunctionStatement(visitor, x)

	case *show.SQLShowCreateProcedureStatement:
		x, _ := x.(*show.SQLShowCreateProcedureStatement)
		AcceptShowCreateProcedureStatement(visitor, x)

	case *show.SQLShowCreateTableStatement:
		x, _ := x.(*show.SQLShowCreateTableStatement)
		AcceptShowCreateTableStatement(visitor, x)

	case *show.SQLShowCreateTriggerStatement:
		x, _ := x.(*show.SQLShowCreateTriggerStatement)
		AcceptShowCreateTriggerStatement(visitor, x)

	case *show.SQLShowCreateViewStatement:
		x, _ := x.(*show.SQLShowCreateViewStatement)
		AcceptShowCreateViewStatement(visitor, x)


		// -------------------------- Explain --------------------------
	case *statement.SQLDescStatement:
		x, _ := x.(*statement.SQLDescStatement)
		AcceptDescStatement(visitor, x)
	case *statement.SQLDescribeStatement:
		x, _ := x.(*statement.SQLDescribeStatement)
		AcceptDescribeStatement(visitor, x)
	case *statement.SQLExplainStatement:
		x, _ := x.(*statement.SQLExplainStatement)
		AcceptExplainStatement(visitor, x)

		// -------------------------- Help --------------------------
	case *statement.SQLHelpStatement:
		x, _ := x.(*statement.SQLHelpStatement)
		AcceptHelpStatement(visitor, x)
		// -------------------------- Use --------------------------
	case *statement.SQLUseStatement:
		x, _ := x.(*statement.SQLUseStatement)
		AcceptUseStatement(visitor, x)
	default:
		panic("TODO: " + xType.String())
	}


	visitor.AfterVisit(visitor, x)
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
