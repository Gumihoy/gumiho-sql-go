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
	exprView "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/comment"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/database"
	statementFunction "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/function"
	statementIndex "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/index"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/package"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/packagebody"
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
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	"strings"
)

type ISQLVisitor interface {
	BeforeVisit(child ISQLVisitor, x ast.ISQLObject)
	AfterVisit(child ISQLVisitor, x ast.ISQLObject)

	// ---------------------------- Comment Start ------------------------------------
	VisitMultiLineComment(child ISQLVisitor, x *ast.SQLMultiLineComment) bool
	VisitMinusComment(child ISQLVisitor, x *ast.SQLMinusComment) bool
	VisitSharpComment(child ISQLVisitor, x *ast.SQLSharpComment) bool

	VisitMultiLineHint(child ISQLVisitor, x *ast.SQLMultiLineHint) bool
	VisitMinusHint(child ISQLVisitor, x *ast.SQLMinusHint) bool
	// ---------------------------- Comment End ------------------------------------

	// ---------------------------- Literal Start ------------------------------------
	// visitNQStringLiteral(x SQLNQStringLiteral) bool
	// visitNStringLiteral(x SQLNStringLiteral) bool
	// visitQStringLiteral(x SQLQStringLiteral) bool
	VisitStringLiteral(child ISQLVisitor, x *literal.SQLStringLiteral) bool
	VisitCharacterStringLiteral(child ISQLVisitor, x *literal.SQLCharacterStringLiteral) bool

	// visitBinaryDoubleLiteral(x SQLBinaryDoubleLiteral) bool
	// visitBinaryFloatLiteral(x SQLBinaryFloatLiteral) bool
	// visitDecimalLiteral(x SQLDecimalLiteral) bool
	// visitFloatingPointLiteral(x SQLFloatingPointLiteral) bool
	VisitIntegerLiteral(child ISQLVisitor, x *literal.SQLIntegerLiteral) bool
	VisitFloatingPointLiteral(child ISQLVisitor, x *literal.SQLFloatingPointLiteral) bool
	VisitHexadecimalLiteral(child ISQLVisitor, x *literal.SQLHexadecimalLiteral) bool
	// visitBitValueLiteral(x literal.SQLBitValueLiteral) bool
	// visitHexaDecimalLiteral(x SQLHexaDecimalLiteral) bool
	//
	VisitDateLiteral(child ISQLVisitor, x *literal.SQLDateLiteral) bool
	VisitTimeLiteral(child ISQLVisitor, x *literal.SQLTimeLiteral) bool
	VisitTimestampLiteral(child ISQLVisitor, x *literal.SQLTimestampLiteral) bool

	VisitBooleanLiteral(child ISQLVisitor, x *literal.SQLBooleanLiteral) bool
	// ---------------------------- Literal End ------------------------------------

	// ---------------------------- Identifier Start ------------------------------------
	VisitIdentifier(child ISQLVisitor, x *expr.SQLUnQuotedIdentifier) bool
	VisitDoubleQuotedIdentifier(child ISQLVisitor, x *expr.SQLDoubleQuotedIdentifier) bool
	VisitReverseQuotedIdentifier(child ISQLVisitor, x *expr.SQLReverseQuotedIdentifier) bool

	VisitName(child ISQLVisitor, x *expr.SQLName) bool

	// ---------------------------- Identifier End ------------------------------------

	// ---------------------------- Variable Start ------------------------------------
	VisitVariableExpr(child ISQLVisitor, x *variable.SQLVariableExpr) bool
	VisitAtVariableExpr(child ISQLVisitor, x *variable.SQLAtVariableExpr) bool
	VisitAtAtVariableExpr(child ISQLVisitor, x *variable.SQLAtAtVariableExpr) bool
	VisitBindVariableExpr(child ISQLVisitor, x *variable.SQLBindVariableExpr) bool

	// ---------------------------- Variable End ------------------------------------

	// ---------------------------- Operator Start ------------------------------------
	VisitUnaryOperatorExpr(child ISQLVisitor, x *operator.SQLUnaryOperatorExpr) bool
	VisitBinaryOperatorExpr(child ISQLVisitor, x *operator.SQLBinaryOperatorExpr) bool

	// ---------------------------- Operator End ------------------------------------

	// ---------------------------- Condition Start ------------------------------------
	VisitIsCondition(child ISQLVisitor, x *condition.SQLIsCondition) bool
	VisitLikeCondition(child ISQLVisitor, x *condition.SQLLikeCondition) bool
	VisitBetweenCondition(child ISQLVisitor, x *condition.SQLBetweenCondition) bool
	VisitInCondition(child ISQLVisitor, x *condition.SQLInCondition) bool
	VisitIsASetCondition(child ISQLVisitor, x *condition.SQLIsASetCondition) bool

	// ---------------------------- Condition End ------------------------------------

	// ---------------------------- Expr Start ------------------------------------
	VisitDBLinkExpr(child ISQLVisitor, x *expr.SQLDBLinkExpr) bool
	VisitAllColumnExpr(child ISQLVisitor, x *expr.SQLAllColumnExpr) bool
	VisitNullExpr(child ISQLVisitor, x *expr.SQLNullExpr) bool
	VisitListExpr(child ISQLVisitor, x *expr.SQLListExpr) bool
	VisitSubQuery(child ISQLVisitor, x *common.SQLSubQueryExpr) bool
	VisitAssignExpr(child ISQLVisitor, x *expr.SQLAssignExpr) bool
	// ---------------------------- Expr End ------------------------------------

	// ---------------------------- Function Start ------------------------------------
	VisitMethodInvocation(child ISQLVisitor, x *function.SQLMethodInvocation) bool
	VisitStaticMethodInvocation(child ISQLVisitor, x *function.SQLStaticMethodInvocation) bool
	VisitCastFunctionArgument(child ISQLVisitor, x *function.SQLCastFunctionArgument) bool

	// ---------------------------- Function End ------------------------------------

	// ---------------------------- DataType Start ------------------------------------
	VisitDataType(child ISQLVisitor, x *datatype.SQLDataType) bool
	VisitIntervalDataType(child ISQLVisitor, x *datatype.SQLIntervalDataType) bool
	VisitIntervalDataTypeField(child ISQLVisitor, x *datatype.SQLIntervalDataTypeField) bool
	VisitDateDataType(child ISQLVisitor, x *datatype.SQLDateDataType) bool
	VisitDateTimeDataType(child ISQLVisitor, x *datatype.SQLDateTimeDataType) bool
	VisitTimeDataType(child ISQLVisitor, x *datatype.SQLTimeDataType) bool
	VisitTimestampDataType(child ISQLVisitor, x *datatype.SQLTimestampDataType) bool
	// ---------------------------- DataType End ------------------------------------

	// ---------------------------- Index Start ------------------------------------
	VisitIndexColumn(child ISQLVisitor, x *index.SQLIndexColumn) bool
	// ---------------------------- Index End ------------------------------------

	// ---------------------------- Sequence Start ------------------------------------
	VisitIncrementBySequenceOption(child ISQLVisitor, x *exprSequence.SQLIncrementBySequenceOption) bool
	VisitStartWithSequenceOption(child ISQLVisitor, x *exprSequence.SQLStartWithSequenceOption) bool
	VisitMaxValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLMaxValueSequenceOption) bool
	VisitNoMaxValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoMaxValueSequenceOption) bool
	VisitMinValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLMinValueSequenceOption) bool
	VisitNoMinValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoMinValueSequenceOption) bool
	VisitCycleSequenceOption(child ISQLVisitor, x *exprSequence.SQLCycleSequenceOption) bool
	VisitNoCycleSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoCycleSequenceOption) bool
	VisitCacheSequenceOption(child ISQLVisitor, x *exprSequence.SQLCacheSequenceOption) bool
	VisitNoCacheSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoCacheSequenceOption) bool
	VisitOrderSequenceOption(child ISQLVisitor, x *exprSequence.SQLOrderSequenceOption) bool
	VisitNoOrderSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoOrderSequenceOption) bool
	VisitKeepSequenceOption(child ISQLVisitor, x *exprSequence.SQLKeepSequenceOption) bool
	VisitNoKeepSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoKeepSequenceOption) bool
	VisitScaleSequenceOption(child ISQLVisitor, x *exprSequence.SQLScaleSequenceOption) bool
	VisitNoScaleSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoScaleSequenceOption) bool
	VisitSessionSequenceOption(child ISQLVisitor, x *exprSequence.SQLSessionSequenceOption) bool
	VisitGlobalSequenceOption(child ISQLVisitor, x *exprSequence.SQLGlobalSequenceOption) bool
	// ---------------------------- Sequence End ------------------------------------

	// ---------------------------- Server Start ------------------------------------
	VisitHostOption(child ISQLVisitor, x *exprServer.SQLHostOption) bool
	VisitDatabaseOption(child ISQLVisitor, x *exprServer.SQLDatabaseOption) bool
	VisitUserOption(child ISQLVisitor, x *exprServer.SQLUserOption) bool
	VisitPasswordOption(child ISQLVisitor, x *exprServer.SQLPasswordOption) bool
	VisitSocketOption(child ISQLVisitor, x *exprServer.SQLSocketOption) bool
	VisitOwnerOption(child ISQLVisitor, x *exprServer.SQLOwnerOption) bool
	VisitPortOption(child ISQLVisitor, x *exprServer.SQLPortOption) bool
	// ---------------------------- Server End ------------------------------------

	// ---------------------------- OnTable Start ------------------------------------
	VisitTableColumn(child ISQLVisitor, x *exprTable.SQLTableColumn) bool
	VisitPrimaryKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLPrimaryKeyTableConstraint) bool
	VisitUniqueTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueTableConstraint) bool
	VisitUniqueIndexTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueIndexTableConstraint) bool
	VisitUniqueKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueKeyTableConstraint) bool
	VisitForeignKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLForeignKeyTableConstraint) bool
	VisitCheckTableConstraint(child ISQLVisitor, x *exprTable.SQLCheckTableConstraint) bool
	VisitTableLikeClause(child ISQLVisitor, x *exprTable.SQLTableLikeClause) bool

	VisitPrimaryKeyColumnConstraint(child ISQLVisitor, x *exprTable.SQLPrimaryKeyColumnConstraint) bool
	VisitKeyColumnConstraint(child ISQLVisitor, x *exprTable.SQLKeyColumnConstraint) bool
	VisitUniqueColumnConstraint(child ISQLVisitor, x *exprTable.SQLUniqueColumnConstraint) bool
	VisitNullColumnConstraint(child ISQLVisitor, x *exprTable.SQLNullColumnConstraint) bool
	VisitNotNullColumnConstraint(child ISQLVisitor, x *exprTable.SQLNotNullColumnConstraint) bool
	VisitCheckColumnConstraint(child ISQLVisitor, x *exprTable.SQLCheckColumnConstraint) bool

	VisitDefaultClause(child ISQLVisitor, x *exprTable.SQLDefaultClause) bool
	VisitAutoIncrementExpr(child ISQLVisitor, x *exprTable.SQLAutoIncrementExpr) bool
	VisitVisibleExpr(child ISQLVisitor, x *exprTable.SQLVisibleExpr) bool
	VisitInvisibleExpr(child ISQLVisitor, x *exprTable.SQLInvisibleExpr) bool
	VisitCommentExpr(child ISQLVisitor, x *exprTable.SQLCommentExpr) bool

	VisitCharsetAssignExpr(child ISQLVisitor, x *exprTable.SQLCharsetAssignExpr) bool
	VisitCharacterSetAssignExpr(child ISQLVisitor, x *exprTable.SQLCharacterSetAssignExpr) bool

	VisitPartitionByHash(child ISQLVisitor, x *exprTable.SQLPartitionByHash) bool
	VisitPartitionByKey(child ISQLVisitor, x *exprTable.SQLPartitionByKey) bool
	VisitPartitionByRange(child ISQLVisitor, x *exprTable.SQLPartitionByRange) bool
	VisitPartitionByList(child ISQLVisitor, x *exprTable.SQLPartitionByList) bool

	VisitSubPartitionByHash(child ISQLVisitor, x *exprTable.SQLSubPartitionByHash) bool
	VisitSubPartitionByKey(child ISQLVisitor, x *exprTable.SQLSubPartitionByKey) bool

	VisitPartitionDefinition(child ISQLVisitor, x *exprTable.SQLPartitionDefinition) bool

	VisitPartitionValuesLessThan(child ISQLVisitor, x *exprTable.SQLPartitionValuesLessThan) bool
	VisitPartitionValuesLessThanMaxValue(child ISQLVisitor, x *exprTable.SQLPartitionValuesLessThanMaxValue) bool
	VisitPartitionValuesIn(child ISQLVisitor, x *exprTable.SQLPartitionValuesIn) bool

	VisitSubPartitionDefinition(child ISQLVisitor, x *exprTable.SQLSubPartitionDefinition) bool

	// ----- Alter OnTable
	VisitAddColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddColumnAlterTableAction) bool
	VisitAlterColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLAlterColumnAlterTableAction) bool
	VisitDropColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropColumnAlterTableAction) bool
	VisitAddTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddTableConstraintAlterTableAction) bool

	VisitDropIndexAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropIndexAlterTableAction) bool
	VisitDropKeyAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropKeyAlterTableAction) bool

	VisitDropConstraintTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropTableConstraintAlterTableAction) bool
	VisitDropPrimaryKeyTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropPrimaryKeyTableConstraintAlterTableAction) bool
	VisitDropUniqueTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropUniqueTableConstraintAlterTableAction) bool
	VisitDropForeignKeyTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropForeignKeyTableConstraintAlterTableAction) bool
	VisitDropCheckTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropCheckTableConstraintAlterTableAction) bool

	VisitAddPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddPartitionAlterTableAction) bool
	VisitDropPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropPartitionAlterTableAction) bool
	VisitDiscardPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLDiscardPartitionAlterTableAction) bool
	VisitImportPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLImportPartitionAlterTableAction) bool
	VisitTruncatePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLTruncatePartitionAlterTableAction) bool
	VisitCoalescePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLCoalescePartitionAlterTableAction) bool
	VisitReorganizePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLReorganizePartitionAlterTableAction) bool
	VisitExchangePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLExchangePartitionAlterTableAction) bool
	VisitAnalyzePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLAnalyzePartitionAlterTableAction) bool
	VisitCheckPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLCheckPartitionAlterTableAction) bool
	VisitOptimizePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLOptimizePartitionAlterTableAction) bool
	VisitRebuildPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRebuildPartitionAlterTableAction) bool
	VisitRepairPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRepairPartitionAlterTableAction) bool
	VisitRemovePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRemovePartitionAlterTableAction) bool

	// ---------------------------- Table End ------------------------------------

	// ---------------------------- User Start ------------------------------------
	VisitUserName(child ISQLVisitor, x *exprUser.SQLUserName) bool
	VisitIdentifiedByAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByAuthOption) bool
	VisitIdentifiedByPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByPasswordAuthOption) bool
	VisitIdentifiedByRandomPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByRandomPasswordAuthOption) bool
	VisitIdentifiedWithAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithAuthOption) bool
	VisitIdentifiedWithByAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithByAuthOption) bool
	VisitIdentifiedWithByRandomPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithByRandomPasswordAuthOption) bool
	VisitIdentifiedWithAsAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithAsAuthOption) bool
	// ---------------------------- User End ------------------------------------

	// ---------------------------- View Start ------------------------------------
	VisitViewColumn(child ISQLVisitor, x *exprView.SQLViewColumn) bool
	// ---------------------------- View End ------------------------------------

	// ---------------------------- Select Start ------------------------------------
	VisitSelectQuery(child ISQLVisitor, x *select_.SQLSelectQuery) bool
	VisitParenSelectQuery(child ISQLVisitor, x *select_.SQLParenSelectQuery) bool
	VisitSelectUnionQuery(child ISQLVisitor, x *select_.SQLSelectUnionQuery) bool

	VisitWithClause(child ISQLVisitor, x *select_.SQLWithClause) bool
	VisitSubQueryFactoringClause(child ISQLVisitor, x *select_.SQLSubQueryFactoringClause) bool
	VisitSearchClause(child ISQLVisitor, x *select_.SQLSearchClause) bool
	VisitSubAvFactoringClause(child ISQLVisitor, x *select_.SQLSubAvFactoringClause) bool
	VisitSubAvClause(child ISQLVisitor, x *select_.SQLSubAvClause) bool

	VisitSelectElement(child ISQLVisitor, x *select_.SQLSelectElement) bool
	VisitSelectTargetElement(child ISQLVisitor, x *select_.SQLSelectTargetElement) bool

	VisitFromClause(child ISQLVisitor, x *select_.SQLFromClause) bool
	VisitTableReference(child ISQLVisitor, x *select_.SQLTableReference) bool

	VisitPartitionClause(child ISQLVisitor, x *select_.SQLPartitionClause) bool
	VisitPartitionForClause(child ISQLVisitor, x *select_.SQLPartitionForClause) bool
	VisitSubPartitionClause(child ISQLVisitor, x *select_.SQLSubPartitionClause) bool
	VisitSubPartitionForClause(child ISQLVisitor, x *select_.SQLSubPartitionForClause) bool

	VisitSampleClause(child ISQLVisitor, x *select_.SQLSampleClause) bool

	VisitOJTableReference(child ISQLVisitor, x *select_.SQLOJTableReference) bool
	VisitSubQueryTableReference(child ISQLVisitor, x *select_.SQLSubQueryTableReference) bool
	VisitJoinTableReference(child ISQLVisitor, x *select_.SQLJoinTableReference) bool
	VisitJoinOnCondition(child ISQLVisitor, x *select_.SQLJoinOnCondition) bool
	VisitJoinUsingCondition(child ISQLVisitor, x *select_.SQLJoinUsingCondition) bool

	VisitWhereClause(child ISQLVisitor, x *select_.SQLWhereClause) bool

	VisitHierarchicalQueryClauseConnectBy(child ISQLVisitor, x *select_.SQLHierarchicalQueryClauseConnectBy) bool
	VisitHierarchicalQueryClauseStartWith(child ISQLVisitor, x *select_.SQLHierarchicalQueryClauseStartWith) bool

	VisitGroupByHavingClause(child ISQLVisitor, x *select_.SQLGroupByHavingClause) bool
	VisitHavingGroupByClause(child ISQLVisitor, x *select_.SQLHavingGroupByClause) bool
	VisitGroupByElement(child ISQLVisitor, x *select_.SQLGroupByElement) bool

	VisitOrderByClause(child ISQLVisitor, x *select_.SQLOrderByClause) bool
	VisitOrderByElement(child ISQLVisitor, x *select_.SQLOrderByElement) bool
	VisitLimitOffsetClause(child ISQLVisitor, x *select_.SQLLimitOffsetClause) bool
	VisitOffsetFetchClause(child ISQLVisitor, x *select_.SQLOffsetFetchClause) bool

	VisitForUpdateClause(child ISQLVisitor, x *select_.SQLForUpdateClause) bool
	VisitForShareClause(child ISQLVisitor, x *select_.SQLForShareClause) bool
	VisitLockInShareModeClause(child ISQLVisitor, x *select_.SQLLockInShareModeClause) bool

	// ---------------------------- Select End ------------------------------------

	// ---------------------------- Statement Expr Start ------------------------------------
	VisitEditionAbleExpr(child ISQLVisitor, x *statement.SQLEditionAbleExpr) bool
	VisitNonEditionAbleExpr(child ISQLVisitor, x *statement.SQLNonEditionAbleExpr) bool
	VisitCompileExpr(child ISQLVisitor, x *statement.SQLCompileExpr) bool
	// ---------------------------- Statement Expr End ------------------------------------

	// ---------------------------- Statement Start ------------------------------------
	VisitCommentOnAuditPolicyStatement(child ISQLVisitor, x *comment.SQLCommentOnAuditPolicyStatement) bool
	VisitCommentOnColumnStatement(child ISQLVisitor, x *comment.SQLCommentOnColumnStatement) bool
	VisitCommentOnEditionStatement(child ISQLVisitor, x *comment.SQLCommentOnEditionStatement) bool
	VisitCommentOnIndextypeStatement(child ISQLVisitor, x *comment.SQLCommentOnIndextypeStatement) bool
	VisitCommentOnMaterializedViewStatement(child ISQLVisitor, x *comment.SQLCommentOnMaterializedViewStatement) bool
	VisitCommentOnMiningModelStatement(child ISQLVisitor, x *comment.SQLCommentOnMiningModelStatement) bool
	VisitCommentOnOperatorStatement(child ISQLVisitor, x *comment.SQLCommentOnOperatorStatement) bool
	VisitCommentOnTableStatement(child ISQLVisitor, x *comment.SQLCommentOnTableStatement) bool

	VisitAlterDatabaseStatement(child ISQLVisitor, x *database.SQLAlterDatabaseStatement) bool
	VisitCreateDatabaseStatement(child ISQLVisitor, x *database.SQLCreateDatabaseStatement) bool
	VisitDropDatabaseStatement(child ISQLVisitor, x *database.SQLDropDatabaseStatement) bool

	VisitAlterFunctionStatement(child ISQLVisitor, x *statementFunction.SQLAlterFunctionStatement) bool
	VisitCreateFunctionStatement(child ISQLVisitor, x *statementFunction.SQLCreateFunctionStatement) bool
	VisitDropFunctionStatement(child ISQLVisitor, x *statementFunction.SQLDropFunctionStatement) bool

	VisitAlterIndexStatement(child ISQLVisitor, x *statementIndex.SQLAlterIndexStatement) bool
	VisitCreateIndexStatement(child ISQLVisitor, x *statementIndex.SQLCreateIndexStatement) bool
	VisitDropIndexStatement(child ISQLVisitor, x *statementIndex.SQLDropIndexStatement) bool

	VisitAlterPackageStatement(child ISQLVisitor, x *package_.SQLAlterPackageStatement) bool
	VisitCreatePackageStatement(child ISQLVisitor, x *package_.SQLCreatePackageStatement) bool
	VisitDropPackageStatement(child ISQLVisitor, x *package_.SQLDropPackageStatement) bool

	VisitAlterPackageBodyStatement(child ISQLVisitor, x *packagebody.SQLAlterPackageBodyStatement) bool
	VisitCreatePackageBoydStatement(child ISQLVisitor, x *packagebody.SQLCreatePackageBoydStatement) bool
	VisitDropPackageBodyStatement(child ISQLVisitor, x *packagebody.SQLDropPackageBodyStatement) bool

	VisitAlterProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLAlterProcedureStatement) bool
	VisitCreateProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLCreateProcedureStatement) bool
	VisitDropProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLDropProcedureStatement) bool

	VisitAlterRoleStatement(child ISQLVisitor, x *statementRole.SQLAlterRoleStatement) bool
	VisitCreateRoleStatement(child ISQLVisitor, x *statementRole.SQLCreateRoleStatement) bool
	VisitDropRoleStatement(child ISQLVisitor, x *statementRole.SQLDropRoleStatement) bool

	VisitAlterSchemaStatement(child ISQLVisitor, x *schema.SQLAlterSchemaStatement) bool
	VisitCreateSchemaStatement(child ISQLVisitor, x *schema.SQLCreateSchemaStatement) bool
	VisitDropSchemaStatement(child ISQLVisitor, x *schema.SQLDropSchemaStatement) bool

	VisitAlterSequenceStatement(child ISQLVisitor, x *sequence.SQLAlterSequenceStatement) bool
	VisitCreateSequenceStatement(child ISQLVisitor, x *sequence.SQLCreateSequenceStatement) bool
	VisitDropSequenceStatement(child ISQLVisitor, x *sequence.SQLDropSequenceStatement) bool

	VisitAlterServerStatement(child ISQLVisitor, x *statementServer.SQLAlterServerStatement) bool
	VisitCreateServerStatement(child ISQLVisitor, x *statementServer.SQLCreateServerStatement) bool
	VisitDropServerStatement(child ISQLVisitor, x *statementServer.SQLDropServerStatement) bool

	VisitAlterSynonymStatement(child ISQLVisitor, x *synonym.SQLAlterSynonymStatement) bool
	VisitCreateSynonymStatement(child ISQLVisitor, x *synonym.SQLCreateSynonymStatement) bool
	VisitDropSynonymStatement(child ISQLVisitor, x *synonym.SQLDropSynonymStatement) bool

	VisitAlterTableStatement(child ISQLVisitor, x *statementTable.SQLAlterTableStatement) bool
	VisitCreateTableStatement(child ISQLVisitor, x *statementTable.SQLCreateTableStatement) bool
	VisitDropTableStatement(child ISQLVisitor, x *statementTable.SQLDropTableStatement) bool

	VisitAlterTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLAlterTriggerStatement) bool
	VisitCreateTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLCreateTriggerStatement) bool
	VisitDropTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLDropTriggerStatement) bool

	VisitAlterTypeStatement(child ISQLVisitor, x *statementType.SQLAlterTypeStatement) bool
	VisitCreateTypeStatement(child ISQLVisitor, x *statementType.SQLCreateTypeStatement) bool
	VisitDropTypeStatement(child ISQLVisitor, x *statementType.SQLDropTypeStatement) bool

	VisitAlterTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLAlterTypeBodyStatement) bool
	VisitCreateTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLCreateTypeBodyStatement) bool
	VisitDropTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLDropTypeBodyStatement) bool

	VisitAlterUserStatement(child ISQLVisitor, x *statementUser.SQLAlterUserStatement) bool
	VisitCreateUserStatement(child ISQLVisitor, x *statementUser.SQLCreateUserStatement) bool
	VisitDropUserStatement(child ISQLVisitor, x *statementUser.SQLDropUserStatement) bool

	VisitAlterViewStatement(child ISQLVisitor, x *statementView.SQLAlterViewStatement) bool
	VisitCreateViewStatement(child ISQLVisitor, x *statementView.SQLCreateViewStatement) bool
	VisitDropViewStatement(child ISQLVisitor, x *statementView.SQLDropViewStatement) bool

	VisitDeleteStatement(child ISQLVisitor, x *statement.SQLDeleteStatement) bool
	VisitInsertStatement(child ISQLVisitor, x *statement.SQLInsertStatement) bool
	VisitSelectStatement(child ISQLVisitor, x *statement.SQLSelectStatement) bool
	VisitUpdateStatement(child ISQLVisitor, x *statement.SQLUpdateStatement) bool

	VisitSetVariableAssignmentStatement(child ISQLVisitor, x *set.SQLSetVariableAssignmentStatement) bool
	VisitSetCharacterSetStatement(child ISQLVisitor, x *set.SQLSetCharacterSetStatement) bool
	VisitSetCharsetStatement(child ISQLVisitor, x *set.SQLSetCharsetStatement) bool
	VisitSetNamesStatement(child ISQLVisitor, x *set.SQLSetNamesStatement) bool

	VisitShowCreateDatabaseStatement(child ISQLVisitor, x *show.SQLShowCreateDatabaseStatement) bool
	VisitShowCreateEventStatement(child ISQLVisitor, x *show.SQLShowCreateEventStatement) bool
	VisitShowCreateFunctionStatement(child ISQLVisitor, x *show.SQLShowCreateFunctionStatement) bool
	VisitShowCreateProcedureStatement(child ISQLVisitor, x *show.SQLShowCreateProcedureStatement) bool
	VisitShowCreateTableStatement(child ISQLVisitor, x *show.SQLShowCreateTableStatement) bool
	VisitShowCreateTriggerStatement(child ISQLVisitor, x *show.SQLShowCreateTriggerStatement) bool
	VisitShowCreateViewStatement(child ISQLVisitor, x *show.SQLShowCreateViewStatement) bool

	VisitDescStatement(child ISQLVisitor, x *statement.SQLDescStatement) bool
	VisitDescribeStatement(child ISQLVisitor, x *statement.SQLDescribeStatement) bool
	VisitExplainStatement(child ISQLVisitor, x *statement.SQLExplainStatement) bool

	VisitHelpStatement(child ISQLVisitor, x *statement.SQLHelpStatement) bool
	VisitUseStatement(child ISQLVisitor, x *statement.SQLUseStatement) bool
	// ---------------------------- Statement End ------------------------------------
}

type SQLVisitorAdapter struct {
}

func NewVisitorAdapter() *SQLVisitorAdapter {
	return &SQLVisitorAdapter{}
}

func (v *SQLVisitorAdapter) BeforeVisit(child ISQLVisitor, x ast.ISQLObject) {

}

func (v *SQLVisitorAdapter) VisitMultiLineComment(child ISQLVisitor, x *ast.SQLMultiLineComment) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitMinusComment(child ISQLVisitor, x *ast.SQLMinusComment) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitSharpComment(child ISQLVisitor, x *ast.SQLSharpComment) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitMultiLineHint(child ISQLVisitor, x *ast.SQLMultiLineHint) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitMinusHint(child ISQLVisitor, x *ast.SQLMinusHint) bool {
	panic("implement me")
}

// ---------------------------- Literal Start ------------------------------------

func (v *SQLVisitorAdapter) VisitStringLiteral(child ISQLVisitor, x *literal.SQLStringLiteral) bool                   { return true }
func (v *SQLVisitorAdapter) VisitCharacterStringLiteral(child ISQLVisitor, x *literal.SQLCharacterStringLiteral) bool { return true }

func (v *SQLVisitorAdapter) VisitIntegerLiteral(child ISQLVisitor, x *literal.SQLIntegerLiteral) bool             { return true }
func (v *SQLVisitorAdapter) VisitFloatingPointLiteral(child ISQLVisitor, x *literal.SQLFloatingPointLiteral) bool { return true }
func (v *SQLVisitorAdapter) VisitHexadecimalLiteral(child ISQLVisitor, x *literal.SQLHexadecimalLiteral) bool     { return true }

func (v *SQLVisitorAdapter) VisitDateLiteral(child ISQLVisitor, x *literal.SQLDateLiteral) bool           { return true }
func (v *SQLVisitorAdapter) VisitTimeLiteral(child ISQLVisitor, x *literal.SQLTimeLiteral) bool           { return true }
func (v *SQLVisitorAdapter) VisitTimestampLiteral(child ISQLVisitor, x *literal.SQLTimestampLiteral) bool { return true }
func (v *SQLVisitorAdapter) VisitBooleanLiteral(child ISQLVisitor, x *literal.SQLBooleanLiteral) bool     { return true }

// ---------------------------- Literal End ------------------------------------

// ---------------------------- Identifier Start ------------------------------------
func (v *SQLVisitorAdapter) VisitIdentifier(child ISQLVisitor, x *expr.SQLUnQuotedIdentifier) bool                   { return true }
func (v *SQLVisitorAdapter) VisitDoubleQuotedIdentifier(child ISQLVisitor, x *expr.SQLDoubleQuotedIdentifier) bool   { return true }
func (v *SQLVisitorAdapter) VisitReverseQuotedIdentifier(child ISQLVisitor, x *expr.SQLReverseQuotedIdentifier) bool { return true }

func (v *SQLVisitorAdapter) VisitName(child ISQLVisitor, x *expr.SQLName) bool {
	return true
}

// ---------------------------- Identifier End ------------------------------------

// ---------------------------- Variable Start ------------------------------------
func (v *SQLVisitorAdapter) VisitVariableExpr(child ISQLVisitor, x *variable.SQLVariableExpr) bool         { return true }
func (v *SQLVisitorAdapter) VisitAtVariableExpr(child ISQLVisitor, x *variable.SQLAtVariableExpr) bool     { return true }
func (v *SQLVisitorAdapter) VisitAtAtVariableExpr(child ISQLVisitor, x *variable.SQLAtAtVariableExpr) bool { return true }
func (v *SQLVisitorAdapter) VisitBindVariableExpr(child ISQLVisitor, x *variable.SQLBindVariableExpr) bool { return true }

// ---------------------------- Variable End ------------------------------------

// ---------------------------- Operator Start ------------------------------------
func (v *SQLVisitorAdapter) VisitUnaryOperatorExpr(child ISQLVisitor, x *operator.SQLUnaryOperatorExpr) bool   { return true }
func (v *SQLVisitorAdapter) VisitBinaryOperatorExpr(child ISQLVisitor, x *operator.SQLBinaryOperatorExpr) bool { return true }

// ---------------------------- Operator End ------------------------------------

// ---------------------------- Condition Start ------------------------------------
func (v *SQLVisitorAdapter) VisitIsCondition(child ISQLVisitor, x *condition.SQLIsCondition) bool           { return true }
func (v *SQLVisitorAdapter) VisitLikeCondition(child ISQLVisitor, x *condition.SQLLikeCondition) bool       { return true }
func (v *SQLVisitorAdapter) VisitBetweenCondition(child ISQLVisitor, x *condition.SQLBetweenCondition) bool { return true }
func (v *SQLVisitorAdapter) VisitInCondition(child ISQLVisitor, x *condition.SQLInCondition) bool           { return true }
func (v *SQLVisitorAdapter) VisitIsASetCondition(child ISQLVisitor, x *condition.SQLIsASetCondition) bool   { return true }

// ---------------------------- Condition End ------------------------------------

// ---------------------------- Expr Start ------------------------------------
func (v *SQLVisitorAdapter) VisitDBLinkExpr(child ISQLVisitor, x *expr.SQLDBLinkExpr) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitAllColumnExpr(child ISQLVisitor, x *expr.SQLAllColumnExpr) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitNullExpr(child ISQLVisitor, x *expr.SQLNullExpr) bool       { return true }
func (v *SQLVisitorAdapter) VisitListExpr(child ISQLVisitor, x *expr.SQLListExpr) bool       { return true }
func (v *SQLVisitorAdapter) VisitSubQuery(child ISQLVisitor, x *common.SQLSubQueryExpr) bool { return true }
func (v *SQLVisitorAdapter) VisitAssignExpr(child ISQLVisitor, x *expr.SQLAssignExpr) bool   { return true }

// ---------------------------- Expr End ------------------------------------

// ---------------------------- Function Start ------------------------------------
func (v *SQLVisitorAdapter) VisitMethodInvocation(child ISQLVisitor, x *function.SQLMethodInvocation) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitStaticMethodInvocation(child ISQLVisitor, x *function.SQLStaticMethodInvocation) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitCastFunctionArgument(child ISQLVisitor, x *function.SQLCastFunctionArgument) bool {
	return true
}

// ---------------------------- Function End ------------------------------------

// ---------------------------- DataType Start ------------------------------------
func (v *SQLVisitorAdapter) VisitDataType(child ISQLVisitor, x *datatype.SQLDataType) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitIntervalDataType(child ISQLVisitor, x *datatype.SQLIntervalDataType) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitIntervalDataTypeField(child ISQLVisitor, x *datatype.SQLIntervalDataTypeField) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitDateDataType(child ISQLVisitor, x *datatype.SQLDateDataType) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitDateTimeDataType(child ISQLVisitor, x *datatype.SQLDateTimeDataType) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitTimeDataType(child ISQLVisitor, x *datatype.SQLTimeDataType) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitTimestampDataType(child ISQLVisitor, x *datatype.SQLTimestampDataType) bool {
	return true
}

// ---------------------------- DataType End ------------------------------------

// ---------------------------- Index Start ------------------------------------
func (v *SQLVisitorAdapter) VisitIndexColumn(child ISQLVisitor, x *index.SQLIndexColumn) bool { return true }

// ---------------------------- Index End ------------------------------------

// ---------------------------- Sequence Start ------------------------------------
func (v *SQLVisitorAdapter) VisitIncrementBySequenceOption(child ISQLVisitor, x *exprSequence.SQLIncrementBySequenceOption) bool { return true }
func (v *SQLVisitorAdapter) VisitStartWithSequenceOption(child ISQLVisitor, x *exprSequence.SQLStartWithSequenceOption) bool     { return true }
func (v *SQLVisitorAdapter) VisitMaxValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLMaxValueSequenceOption) bool       { return true }
func (v *SQLVisitorAdapter) VisitNoMaxValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoMaxValueSequenceOption) bool   { return true }
func (v *SQLVisitorAdapter) VisitMinValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLMinValueSequenceOption) bool       { return true }
func (v *SQLVisitorAdapter) VisitNoMinValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoMinValueSequenceOption) bool   { return true }
func (v *SQLVisitorAdapter) VisitCycleSequenceOption(child ISQLVisitor, x *exprSequence.SQLCycleSequenceOption) bool             { return true }
func (v *SQLVisitorAdapter) VisitNoCycleSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoCycleSequenceOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitCacheSequenceOption(child ISQLVisitor, x *exprSequence.SQLCacheSequenceOption) bool             { return true }
func (v *SQLVisitorAdapter) VisitNoCacheSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoCacheSequenceOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitOrderSequenceOption(child ISQLVisitor, x *exprSequence.SQLOrderSequenceOption) bool             { return true }
func (v *SQLVisitorAdapter) VisitNoOrderSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoOrderSequenceOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitKeepSequenceOption(child ISQLVisitor, x *exprSequence.SQLKeepSequenceOption) bool               { return true }
func (v *SQLVisitorAdapter) VisitNoKeepSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoKeepSequenceOption) bool           { return true }
func (v *SQLVisitorAdapter) VisitScaleSequenceOption(child ISQLVisitor, x *exprSequence.SQLScaleSequenceOption) bool             { return true }
func (v *SQLVisitorAdapter) VisitNoScaleSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoScaleSequenceOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitSessionSequenceOption(child ISQLVisitor, x *exprSequence.SQLSessionSequenceOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitGlobalSequenceOption(child ISQLVisitor, x *exprSequence.SQLGlobalSequenceOption) bool           { return true }

// ---------------------------- Sequence End ------------------------------------

// ---------------------------- Server Start ------------------------------------
func (v *SQLVisitorAdapter) VisitHostOption(child ISQLVisitor, x *exprServer.SQLHostOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitDatabaseOption(child ISQLVisitor, x *exprServer.SQLDatabaseOption) bool { return true }
func (v *SQLVisitorAdapter) VisitUserOption(child ISQLVisitor, x *exprServer.SQLUserOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitPasswordOption(child ISQLVisitor, x *exprServer.SQLPasswordOption) bool { return true }
func (v *SQLVisitorAdapter) VisitSocketOption(child ISQLVisitor, x *exprServer.SQLSocketOption) bool     { return true }
func (v *SQLVisitorAdapter) VisitOwnerOption(child ISQLVisitor, x *exprServer.SQLOwnerOption) bool       { return true }
func (v *SQLVisitorAdapter) VisitPortOption(child ISQLVisitor, x *exprServer.SQLPortOption) bool         { return true }

// ---------------------------- Server End ------------------------------------

// ---------------------------- OnTable Start ------------------------------------
func (v *SQLVisitorAdapter) VisitTableColumn(child ISQLVisitor, x *exprTable.SQLTableColumn) bool                               { return true }
func (v *SQLVisitorAdapter) VisitPrimaryKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLPrimaryKeyTableConstraint) bool   { return true }
func (v *SQLVisitorAdapter) VisitUniqueTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueTableConstraint) bool           { return true }
func (v *SQLVisitorAdapter) VisitUniqueIndexTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueIndexTableConstraint) bool { return true }
func (v *SQLVisitorAdapter) VisitUniqueKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueKeyTableConstraint) bool     { return true }
func (v *SQLVisitorAdapter) VisitForeignKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLForeignKeyTableConstraint) bool   { return true }
func (v *SQLVisitorAdapter) VisitCheckTableConstraint(child ISQLVisitor, x *exprTable.SQLCheckTableConstraint) bool             { return true }

func (v *SQLVisitorAdapter) VisitTableLikeClause(child ISQLVisitor, x *exprTable.SQLTableLikeClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitPrimaryKeyColumnConstraint(child ISQLVisitor, x *exprTable.SQLPrimaryKeyColumnConstraint) bool { return true }
func (v *SQLVisitorAdapter) VisitKeyColumnConstraint(child ISQLVisitor, x *exprTable.SQLKeyColumnConstraint) bool               { return true }
func (v *SQLVisitorAdapter) VisitUniqueColumnConstraint(child ISQLVisitor, x *exprTable.SQLUniqueColumnConstraint) bool         { return true }
func (v *SQLVisitorAdapter) VisitNullColumnConstraint(child ISQLVisitor, x *exprTable.SQLNullColumnConstraint) bool             { return true }
func (v *SQLVisitorAdapter) VisitNotNullColumnConstraint(child ISQLVisitor, x *exprTable.SQLNotNullColumnConstraint) bool       { return true }
func (v *SQLVisitorAdapter) VisitCheckColumnConstraint(child ISQLVisitor, x *exprTable.SQLCheckColumnConstraint) bool           { return true }

func (v *SQLVisitorAdapter) VisitDefaultClause(child ISQLVisitor, x *exprTable.SQLDefaultClause) bool         { return true }
func (v *SQLVisitorAdapter) VisitAutoIncrementExpr(child ISQLVisitor, x *exprTable.SQLAutoIncrementExpr) bool { return true }
func (v *SQLVisitorAdapter) VisitVisibleExpr(child ISQLVisitor, x *exprTable.SQLVisibleExpr) bool             { return true }
func (v *SQLVisitorAdapter) VisitInvisibleExpr(child ISQLVisitor, x *exprTable.SQLInvisibleExpr) bool         { return true }
func (v *SQLVisitorAdapter) VisitCommentExpr(child ISQLVisitor, x *exprTable.SQLCommentExpr) bool             { return true }

func (v *SQLVisitorAdapter) VisitCharsetAssignExpr(child ISQLVisitor, x *exprTable.SQLCharsetAssignExpr) bool           { return true }
func (v *SQLVisitorAdapter) VisitCharacterSetAssignExpr(child ISQLVisitor, x *exprTable.SQLCharacterSetAssignExpr) bool { return true }

func (v *SQLVisitorAdapter) VisitPartitionByHash(child ISQLVisitor, x *exprTable.SQLPartitionByHash) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitPartitionByKey(child ISQLVisitor, x *exprTable.SQLPartitionByKey) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitPartitionByRange(child ISQLVisitor, x *exprTable.SQLPartitionByRange) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitPartitionByList(child ISQLVisitor, x *exprTable.SQLPartitionByList) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSubPartitionByHash(child ISQLVisitor, x *exprTable.SQLSubPartitionByHash) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSubPartitionByKey(child ISQLVisitor, x *exprTable.SQLSubPartitionByKey) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitPartitionDefinition(child ISQLVisitor, x *exprTable.SQLPartitionDefinition) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitPartitionValuesLessThan(child ISQLVisitor, x *exprTable.SQLPartitionValuesLessThan) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitPartitionValuesLessThanMaxValue(child ISQLVisitor, x *exprTable.SQLPartitionValuesLessThanMaxValue) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitPartitionValuesIn(child ISQLVisitor, x *exprTable.SQLPartitionValuesIn) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSubPartitionDefinition(child ISQLVisitor, x *exprTable.SQLSubPartitionDefinition) bool {
	return true
}

// ----- Alter OnTable
func (v *SQLVisitorAdapter) VisitAddColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddColumnAlterTableAction) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitAlterColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLAlterColumnAlterTableAction) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitDropColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropColumnAlterTableAction) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitAddTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddTableConstraintAlterTableAction) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitDropIndexAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropIndexAlterTableAction) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitDropKeyAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropKeyAlterTableAction) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitDropConstraintTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropTableConstraintAlterTableAction) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitDropPrimaryKeyTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropPrimaryKeyTableConstraintAlterTableAction) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitDropUniqueTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropUniqueTableConstraintAlterTableAction) bool         { return true }
func (v *SQLVisitorAdapter) VisitDropForeignKeyTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropForeignKeyTableConstraintAlterTableAction) bool { return true }
func (v *SQLVisitorAdapter) VisitDropCheckTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropCheckTableConstraintAlterTableAction) bool           { return true }

func (v *SQLVisitorAdapter) VisitAddPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddPartitionAlterTableAction) bool               { return true }
func (v *SQLVisitorAdapter) VisitDropPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropPartitionAlterTableAction) bool             { return true }
func (v *SQLVisitorAdapter) VisitDiscardPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLDiscardPartitionAlterTableAction) bool       { return true }
func (v *SQLVisitorAdapter) VisitImportPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLImportPartitionAlterTableAction) bool         { return true }
func (v *SQLVisitorAdapter) VisitTruncatePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLTruncatePartitionAlterTableAction) bool     { return true }
func (v *SQLVisitorAdapter) VisitCoalescePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLCoalescePartitionAlterTableAction) bool     { return true }
func (v *SQLVisitorAdapter) VisitReorganizePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLReorganizePartitionAlterTableAction) bool { return true }
func (v *SQLVisitorAdapter) VisitExchangePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLExchangePartitionAlterTableAction) bool     { return true }
func (v *SQLVisitorAdapter) VisitAnalyzePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLAnalyzePartitionAlterTableAction) bool       { return true }
func (v *SQLVisitorAdapter) VisitCheckPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLCheckPartitionAlterTableAction) bool           { return true }
func (v *SQLVisitorAdapter) VisitOptimizePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLOptimizePartitionAlterTableAction) bool     { return true }
func (v *SQLVisitorAdapter) VisitRebuildPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRebuildPartitionAlterTableAction) bool       { return true }
func (v *SQLVisitorAdapter) VisitRepairPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRepairPartitionAlterTableAction) bool         { return true }
func (v *SQLVisitorAdapter) VisitRemovePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRemovePartitionAlterTableAction) bool         { return true }

// ---------------------------- OnTable End ------------------------------------

// ---------------------------- User Start ------------------------------------
func (v *SQLVisitorAdapter) VisitUserName(child ISQLVisitor, x *exprUser.SQLUserName) bool                                                                 { return true }
func (v *SQLVisitorAdapter) VisitIdentifiedByAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByAuthOption) bool                                     { return true }
func (v *SQLVisitorAdapter) VisitIdentifiedByPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByPasswordAuthOption) bool                     { return true }
func (v *SQLVisitorAdapter) VisitIdentifiedByRandomPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByRandomPasswordAuthOption) bool         { return true }
func (v *SQLVisitorAdapter) VisitIdentifiedWithAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithAuthOption) bool                                 { return true }
func (v *SQLVisitorAdapter) VisitIdentifiedWithByAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithByAuthOption) bool                             { return true }
func (v *SQLVisitorAdapter) VisitIdentifiedWithByRandomPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithByRandomPasswordAuthOption) bool { return true }
func (v *SQLVisitorAdapter) VisitIdentifiedWithAsAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithAsAuthOption) bool                             { return true }

// ---------------------------- User End ------------------------------------

// ---------------------------- View Start ------------------------------------
func (v *SQLVisitorAdapter) VisitViewColumn(child ISQLVisitor, x *exprView.SQLViewColumn) bool { return true }

// ---------------------------- View End ------------------------------------

// ---------------------------- Select Start ------------------------------------
func (v *SQLVisitorAdapter) VisitSelectQuery(child ISQLVisitor, x *select_.SQLSelectQuery) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitParenSelectQuery(child ISQLVisitor, x *select_.SQLParenSelectQuery) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSelectUnionQuery(child ISQLVisitor, x *select_.SQLSelectUnionQuery) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitWithClause(child ISQLVisitor, x *select_.SQLWithClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSubQueryFactoringClause(child ISQLVisitor, x *select_.SQLSubQueryFactoringClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSearchClause(child ISQLVisitor, x *select_.SQLSearchClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSubAvFactoringClause(child ISQLVisitor, x *select_.SQLSubAvFactoringClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSubAvClause(child ISQLVisitor, x *select_.SQLSubAvClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSelectElement(child ISQLVisitor, x *select_.SQLSelectElement) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSelectTargetElement(child ISQLVisitor, x *select_.SQLSelectTargetElement) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitFromClause(child ISQLVisitor, x *select_.SQLFromClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitTableReference(child ISQLVisitor, x *select_.SQLTableReference) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitPartitionClause(child ISQLVisitor, x *select_.SQLPartitionClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitPartitionForClause(child ISQLVisitor, x *select_.SQLPartitionForClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSubPartitionClause(child ISQLVisitor, x *select_.SQLSubPartitionClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSubPartitionForClause(child ISQLVisitor, x *select_.SQLSubPartitionForClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSampleClause(child ISQLVisitor, x *select_.SQLSampleClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitOJTableReference(child ISQLVisitor, x *select_.SQLOJTableReference) bool { return true }
func (v *SQLVisitorAdapter) VisitSubQueryTableReference(child ISQLVisitor, x *select_.SQLSubQueryTableReference) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitJoinTableReference(child ISQLVisitor, x *select_.SQLJoinTableReference) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitJoinOnCondition(child ISQLVisitor, x *select_.SQLJoinOnCondition) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitJoinUsingCondition(child ISQLVisitor, x *select_.SQLJoinUsingCondition) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitWhereClause(child ISQLVisitor, x *select_.SQLWhereClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitHierarchicalQueryClauseConnectBy(child ISQLVisitor, x *select_.SQLHierarchicalQueryClauseConnectBy) bool { return true }
func (v *SQLVisitorAdapter) VisitHierarchicalQueryClauseStartWith(child ISQLVisitor, x *select_.SQLHierarchicalQueryClauseStartWith) bool { return true }

func (v *SQLVisitorAdapter) VisitGroupByHavingClause(child ISQLVisitor, x *select_.SQLGroupByHavingClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitHavingGroupByClause(child ISQLVisitor, x *select_.SQLHavingGroupByClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitGroupByElement(child ISQLVisitor, x *select_.SQLGroupByElement) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitOrderByClause(child ISQLVisitor, x *select_.SQLOrderByClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitOrderByElement(child ISQLVisitor, x *select_.SQLOrderByElement) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitLimitOffsetClause(child ISQLVisitor, x *select_.SQLLimitOffsetClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitOffsetFetchClause(child ISQLVisitor, x *select_.SQLOffsetFetchClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitForUpdateClause(child ISQLVisitor, x *select_.SQLForUpdateClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitForShareClause(child ISQLVisitor, x *select_.SQLForShareClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitLockInShareModeClause(child ISQLVisitor, x *select_.SQLLockInShareModeClause) bool {
	return true
}

// ---------------------------- Select End ------------------------------------

// ---------------------------- Statement Expr Start ------------------------------------
func (v *SQLVisitorAdapter) VisitEditionAbleExpr(child ISQLVisitor, x *statement.SQLEditionAbleExpr) bool       { return true }
func (v *SQLVisitorAdapter) VisitNonEditionAbleExpr(child ISQLVisitor, x *statement.SQLNonEditionAbleExpr) bool { return true }
func (v *SQLVisitorAdapter) VisitCompileExpr(child ISQLVisitor, x *statement.SQLCompileExpr) bool               { return true }

// ---------------------------- Statement Expr End ------------------------------------

// ---------------------------- Statement Start ------------------------------------
func (v *SQLVisitorAdapter) VisitCommentOnAuditPolicyStatement(child ISQLVisitor, x *comment.SQLCommentOnAuditPolicyStatement) bool           { return true }
func (v *SQLVisitorAdapter) VisitCommentOnColumnStatement(child ISQLVisitor, x *comment.SQLCommentOnColumnStatement) bool                     { return true }
func (v *SQLVisitorAdapter) VisitCommentOnEditionStatement(child ISQLVisitor, x *comment.SQLCommentOnEditionStatement) bool                   { return true }
func (v *SQLVisitorAdapter) VisitCommentOnIndextypeStatement(child ISQLVisitor, x *comment.SQLCommentOnIndextypeStatement) bool               { return true }
func (v *SQLVisitorAdapter) VisitCommentOnMaterializedViewStatement(child ISQLVisitor, x *comment.SQLCommentOnMaterializedViewStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitCommentOnMiningModelStatement(child ISQLVisitor, x *comment.SQLCommentOnMiningModelStatement) bool           { return true }
func (v *SQLVisitorAdapter) VisitCommentOnOperatorStatement(child ISQLVisitor, x *comment.SQLCommentOnOperatorStatement) bool                 { return true }
func (v *SQLVisitorAdapter) VisitCommentOnTableStatement(child ISQLVisitor, x *comment.SQLCommentOnTableStatement) bool                       { return true }

func (v *SQLVisitorAdapter) VisitAlterDatabaseStatement(child ISQLVisitor, x *database.SQLAlterDatabaseStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateDatabaseStatement(child ISQLVisitor, x *database.SQLCreateDatabaseStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropDatabaseStatement(child ISQLVisitor, x *database.SQLDropDatabaseStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterFunctionStatement(child ISQLVisitor, x *statementFunction.SQLAlterFunctionStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateFunctionStatement(child ISQLVisitor, x *statementFunction.SQLCreateFunctionStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropFunctionStatement(child ISQLVisitor, x *statementFunction.SQLDropFunctionStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterIndexStatement(child ISQLVisitor, x *statementIndex.SQLAlterIndexStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateIndexStatement(child ISQLVisitor, x *statementIndex.SQLCreateIndexStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropIndexStatement(child ISQLVisitor, x *statementIndex.SQLDropIndexStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterPackageStatement(child ISQLVisitor, x *package_.SQLAlterPackageStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreatePackageStatement(child ISQLVisitor, x *package_.SQLCreatePackageStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropPackageStatement(child ISQLVisitor, x *package_.SQLDropPackageStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterPackageBodyStatement(child ISQLVisitor, x *packagebody.SQLAlterPackageBodyStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreatePackageBoydStatement(child ISQLVisitor, x *packagebody.SQLCreatePackageBoydStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropPackageBodyStatement(child ISQLVisitor, x *packagebody.SQLDropPackageBodyStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLAlterProcedureStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLCreateProcedureStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLDropProcedureStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterRoleStatement(child ISQLVisitor, x *statementRole.SQLAlterRoleStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateRoleStatement(child ISQLVisitor, x *statementRole.SQLCreateRoleStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropRoleStatement(child ISQLVisitor, x *statementRole.SQLDropRoleStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterSchemaStatement(child ISQLVisitor, x *schema.SQLAlterSchemaStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateSchemaStatement(child ISQLVisitor, x *schema.SQLCreateSchemaStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropSchemaStatement(child ISQLVisitor, x *schema.SQLDropSchemaStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterSequenceStatement(child ISQLVisitor, x *sequence.SQLAlterSequenceStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateSequenceStatement(child ISQLVisitor, x *sequence.SQLCreateSequenceStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropSequenceStatement(child ISQLVisitor, x *sequence.SQLDropSequenceStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterServerStatement(child ISQLVisitor, x *statementServer.SQLAlterServerStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateServerStatement(child ISQLVisitor, x *statementServer.SQLCreateServerStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropServerStatement(child ISQLVisitor, x *statementServer.SQLDropServerStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterSynonymStatement(child ISQLVisitor, x *synonym.SQLAlterSynonymStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateSynonymStatement(child ISQLVisitor, x *synonym.SQLCreateSynonymStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropSynonymStatement(child ISQLVisitor, x *synonym.SQLDropSynonymStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterTableStatement(child ISQLVisitor, x *statementTable.SQLAlterTableStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateTableStatement(child ISQLVisitor, x *statementTable.SQLCreateTableStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropTableStatement(child ISQLVisitor, x *statementTable.SQLDropTableStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLAlterTriggerStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLCreateTriggerStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLDropTriggerStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterTypeStatement(child ISQLVisitor, x *statementType.SQLAlterTypeStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateTypeStatement(child ISQLVisitor, x *statementType.SQLCreateTypeStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropTypeStatement(child ISQLVisitor, x *statementType.SQLDropTypeStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLAlterTypeBodyStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLCreateTypeBodyStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLDropTypeBodyStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterUserStatement(child ISQLVisitor, x *statementUser.SQLAlterUserStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateUserStatement(child ISQLVisitor, x *statementUser.SQLCreateUserStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropUserStatement(child ISQLVisitor, x *statementUser.SQLDropUserStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitAlterViewStatement(child ISQLVisitor, x *statementView.SQLAlterViewStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitCreateViewStatement(child ISQLVisitor, x *statementView.SQLCreateViewStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitDropViewStatement(child ISQLVisitor, x *statementView.SQLDropViewStatement) bool     { return true }

func (v *SQLVisitorAdapter) VisitDeleteStatement(child ISQLVisitor, x *statement.SQLDeleteStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitInsertStatement(child ISQLVisitor, x *statement.SQLInsertStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitSelectStatement(child ISQLVisitor, x *statement.SQLSelectStatement) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitUpdateStatement(child ISQLVisitor, x *statement.SQLUpdateStatement) bool { return true }

func (v *SQLVisitorAdapter) VisitSetVariableAssignmentStatement(child ISQLVisitor, x *set.SQLSetVariableAssignmentStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitSetCharacterSetStatement(child ISQLVisitor, x *set.SQLSetCharacterSetStatement) bool             { return true }
func (v *SQLVisitorAdapter) VisitSetCharsetStatement(child ISQLVisitor, x *set.SQLSetCharsetStatement) bool                       { return true }
func (v *SQLVisitorAdapter) VisitSetNamesStatement(child ISQLVisitor, x *set.SQLSetNamesStatement) bool                           { return true }

func (v *SQLVisitorAdapter) VisitShowCreateDatabaseStatement(child ISQLVisitor, x *show.SQLShowCreateDatabaseStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitShowCreateEventStatement(child ISQLVisitor, x *show.SQLShowCreateEventStatement) bool         { return true }
func (v *SQLVisitorAdapter) VisitShowCreateFunctionStatement(child ISQLVisitor, x *show.SQLShowCreateFunctionStatement) bool   { return true }
func (v *SQLVisitorAdapter) VisitShowCreateProcedureStatement(child ISQLVisitor, x *show.SQLShowCreateProcedureStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitShowCreateTableStatement(child ISQLVisitor, x *show.SQLShowCreateTableStatement) bool         { return true }
func (v *SQLVisitorAdapter) VisitShowCreateTriggerStatement(child ISQLVisitor, x *show.SQLShowCreateTriggerStatement) bool     { return true }
func (v *SQLVisitorAdapter) VisitShowCreateViewStatement(child ISQLVisitor, x *show.SQLShowCreateViewStatement) bool           { return true }

func (v *SQLVisitorAdapter) VisitDescStatement(child ISQLVisitor, x *statement.SQLDescStatement) bool         { return true }
func (v *SQLVisitorAdapter) VisitDescribeStatement(child ISQLVisitor, x *statement.SQLDescribeStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitExplainStatement(child ISQLVisitor, x *statement.SQLExplainStatement) bool   { return true }

func (v *SQLVisitorAdapter) VisitHelpStatement(child ISQLVisitor, x *statement.SQLHelpStatement) bool { return true }
func (v *SQLVisitorAdapter) VisitUseStatement(child ISQLVisitor, x *statement.SQLUseStatement) bool   { return true }

// ---------------------------- Statement End ------------------------------------

func (v *SQLVisitorAdapter) AfterVisit(child ISQLVisitor, x ast.ISQLObject) {

}

type ISQLOutputVisitor interface {
	ISQLVisitor

	IsLowerCase() bool

	Write(s string)
	WriteKeyword(keyword *Keyword)

	WriteSpace()
	WriteSpaceAfterValue(s string)
	WriteSpaceAfterKeyword(keyword *Keyword)

	WriteLn()
	WriteLnAfterValue(value string)
	WriteLnAfterKeyword(keyword *Keyword)

	WriteIndent()

	IncrementIndentAndWriteLn()
	DecrementIndentAndWriteLn()
}

type SQLOutputVisitor struct {
	*SQLVisitorAdapter

	Builder   *strings.Builder
	Indent    int
	Line, Pos int

	SelectElementLimit int
	Config             *config.SQLOutputConfig
}

func NewOutputVisitor(builder *strings.Builder, config *config.SQLOutputConfig) *SQLOutputVisitor {
	var x SQLOutputVisitor
	x.SQLVisitorAdapter = NewVisitorAdapter()
	x.Builder = builder
	x.Indent = 0
	x.Line, x.Pos = 1, 0

	x.SelectElementLimit = 80
	x.Config = config
	return &x
}

func (v *SQLOutputVisitor) IsLowerCase() bool {
	return v.Config.LowerCase
}

func (v *SQLOutputVisitor) IncrementIndent() {
	v.Indent++
}

func (v *SQLOutputVisitor) IncrementIndentAndWriteLn() {
	v.IncrementIndent()
	v.WriteLn()
}

func (v *SQLOutputVisitor) DecrementIndent() {
	v.Indent--
}

func (v *SQLOutputVisitor) DecrementIndentAndWriteLn() {
	v.DecrementIndent()
	v.WriteLn()
}

func (v *SQLOutputVisitor) IncrementLine() {
	v.Line++
	v.Pos = 0
}

func (v *SQLOutputVisitor) Write(s string) {
	v.Builder.WriteString(s)
	v.Pos += len(s)
}

func (v *SQLOutputVisitor) WriteKeyword(keyword *Keyword) {
	if v.IsLowerCase() {
		v.Builder.WriteString(keyword.Lower)
	} else {
		v.Builder.WriteString(keyword.Upper)
	}
}

// func (v *SQLOutputVisitor) WriteAccept(x ast.ISQLObject) {
// 	if ast.IsNil(x) {
// 		return
// 	}
// 	Accept(v, x)
// }
func WriteAccept(v ISQLVisitor, x ast.ISQLObject) {
	if ast.IsNil(x) {
		return
	}
	Accept(v, x)
}

func (v *SQLOutputVisitor) WriteSpace() {
	v.Builder.WriteString(" ")
	v.Pos++
}
func (v *SQLOutputVisitor) WriteSpaceAfterValue(s string) {
	if s == "" {
		return
	}
	v.WriteSpace()
	v.Write(s)
}

func (v *SQLOutputVisitor) WriteSpaceAfterKeyword(keyword *Keyword) {
	v.WriteSpace()
	if v.IsLowerCase() {
		v.Write(keyword.Lower)
	} else {
		v.Write(keyword.Upper)
	}
}

// func (v *SQLOutputVisitor) WriteSpaceAfterAccept(x ast.ISQLObject) {
// 	if ast.IsNil(x) {
// 		return
// 	}
// 	v.Write(" ")
// 	Accept(v, x)
// }

func WriteSpaceAfterAccept(v ISQLVisitor, x ast.ISQLObject) {
	if ast.IsNil(x) {
		return
	}
	v.(ISQLOutputVisitor).Write(" ")
	Accept(v, x)
}
func WriteSpaceAndIndentLnAfterAccept(v ISQLVisitor, paren bool, x ast.ISQLObject) {
	if ast.IsNil(x) {
		return
	}
	v.(ISQLOutputVisitor).WriteSpace()
	if paren {
		v.(ISQLOutputVisitor).WriteKeyword(SYMB_LEFT_PAREN)
	}

	v.(ISQLOutputVisitor).IncrementIndentAndWriteLn()

	Accept(v, x)

	v.(ISQLOutputVisitor).DecrementIndentAndWriteLn()

	if paren {
		v.(ISQLOutputVisitor).WriteKeyword(SYMB_RIGHT_PAREN)
	}
}

func (v *SQLOutputVisitor) WriteLn() {
	v.Builder.WriteByte('\n')
	v.IncrementLine()
	v.WriteIndent()
}
func (v *SQLOutputVisitor) WriteLnAfterValue(value string) {
	v.WriteLn()
	v.Write(value)
}
func (v *SQLOutputVisitor) WriteLnAfterKeyword(keyword *Keyword) {
	if keyword == nil {
		return
	}
	v.WriteLn()
	v.WriteKeyword(keyword)
}

// func (v *SQLOutputVisitor) WriteLnAfterAccept(x ast.ISQLObject) {
// 	if ast.IsNil(x) {
// 		return
// 	}
// 	v.WriteLn()
// 	Accept(v, x)
// }
func WriteLnAfterAccept(v ISQLVisitor, x ast.ISQLObject) {
	if ast.IsNil(x) {
		return
	}
	v.(ISQLOutputVisitor).WriteLn()
	Accept(v, x)
}

func (v *SQLOutputVisitor) WriteIndent() {
	for i := 0; i < v.Indent; i++ {
		v.Builder.WriteByte('\t')
	}
}

func WriteIndentLnAfterAccept(v ISQLVisitor, paren bool, object ast.ISQLObject) {
	if paren {
		v.(ISQLOutputVisitor).WriteKeyword(SYMB_LEFT_PAREN)
	}

	v.(ISQLOutputVisitor).IncrementIndentAndWriteLn()

	Accept(v, object)

	v.(ISQLOutputVisitor).DecrementIndentAndWriteLn()

	if paren {
		v.(ISQLOutputVisitor).WriteKeyword(SYMB_RIGHT_PAREN)
	}
}

// ---------------------------- Comment Start ------------------------------------
func (v *SQLOutputVisitor) VisitMultiLineComment(child ISQLVisitor, x *ast.SQLMultiLineComment) bool {
	v.WriteKeyword(SYMB_SLASH)
	v.WriteKeyword(SYMB_STAR)
	v.Write(x.Comment())
	v.WriteKeyword(SYMB_STAR)
	v.WriteKeyword(SYMB_SLASH)
	return false
}

func (v *SQLOutputVisitor) VisitMinusComment(child ISQLVisitor, x *ast.SQLMinusComment) bool {
	v.WriteKeyword(SYMB_MINUS)
	v.WriteKeyword(SYMB_MINUS)
	v.WriteSpaceAfterValue(x.Comment())
	return false
}

func (v *SQLOutputVisitor) VisitSharpComment(child ISQLVisitor, x *ast.SQLSharpComment) bool {
	v.WriteKeyword(SYMB_SHARP)
	v.WriteSpaceAfterValue(x.Comment())
	return false
}

func (v *SQLOutputVisitor) VisitMultiLineHint(child ISQLVisitor, x *ast.SQLMultiLineHint) bool {
	return false
}
func (v *SQLOutputVisitor) VisitMinusHint(child ISQLVisitor, x *ast.SQLMinusHint) bool {
	return false
}

// ---------------------------- Comment End ------------------------------------

// ---------------------------- Literal Start ------------------------------------

func (v *SQLOutputVisitor) VisitStringLiteral(child ISQLVisitor, x *literal.SQLStringLiteral) bool {
	v.Write("'")
	v.Write(x.Value())
	v.Write("'")
	return false
}
func (v *SQLOutputVisitor) VisitCharacterStringLiteral(child ISQLVisitor, x *literal.SQLCharacterStringLiteral) bool {

	v.Write(x.Charset())
	v.Write("'")
	v.Write(x.Value())
	v.Write("'")

	return false
}
func (v *SQLOutputVisitor) VisitIntegerLiteral(child ISQLVisitor, x *literal.SQLIntegerLiteral) bool {
	v.Write(x.StringValue())
	return false
}
func (v *SQLOutputVisitor) VisitFloatingPointLiteral(child ISQLVisitor, x *literal.SQLFloatingPointLiteral) bool {
	v.Write(x.StringValue())
	return false
}

func (v *SQLOutputVisitor) VisitHexadecimalLiteral(child ISQLVisitor, x *literal.SQLHexadecimalLiteral) bool {
	v.Write(string(x.Hexadecimal()))

	quote := x.Hexadecimal() == literal.X_LOWER || x.Hexadecimal() == literal.X_UPPER
	if quote {
		v.WriteKeyword(SYMB_SINGLE_QUOTE)
	}
	v.Write(x.Value())
	if quote {
		v.WriteKeyword(SYMB_SINGLE_QUOTE)
	}

	return false
}

func (v *SQLOutputVisitor) VisitDateLiteral(child ISQLVisitor, x *literal.SQLDateLiteral) bool {
	v.Write("DATE")
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitTimeLiteral(child ISQLVisitor, x *literal.SQLTimeLiteral) bool {
	v.Write("TIME")
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitTimestampLiteral(child ISQLVisitor, x *literal.SQLTimestampLiteral) bool {
	v.Write("TIMESTAMP")
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitBooleanLiteral(child ISQLVisitor, x *literal.SQLBooleanLiteral) bool {
	if x.Value() {
		v.Write("true")
	} else {
		v.Write("false")
	}
	return false
}

// ---------------------------- Literal End ------------------------------------

// ---------------------------- Identifier Start ------------------------------------
func (v *SQLOutputVisitor) VisitIdentifier(child ISQLVisitor, x *expr.SQLUnQuotedIdentifier) bool {
	v.Write(x.StringName())
	return false
}
func (v *SQLOutputVisitor) VisitDoubleQuotedIdentifier(child ISQLVisitor, x *expr.SQLDoubleQuotedIdentifier) bool {
	v.Write("\"")
	v.Write(x.StringName())
	v.Write("\"")
	return false
}
func (v *SQLOutputVisitor) VisitReverseQuotedIdentifier(child ISQLVisitor, x *expr.SQLReverseQuotedIdentifier) bool {
	v.Write("`")
	v.Write(x.StringName())
	v.Write("`")
	return false
}

func (v *SQLOutputVisitor) VisitName(child ISQLVisitor, x *expr.SQLName) bool {
	Accept(v, x.Owner())
	v.WriteKeyword(SYMB_DOT)
	Accept(v, x.Name())
	return false
}

// ---------------------------- Identifier End ------------------------------------

// ---------------------------- Variable Start ------------------------------------
func (v *SQLOutputVisitor) VisitVariableExpr(child ISQLVisitor, x *variable.SQLVariableExpr) bool {
	v.Write("?")
	return false
}
func (v *SQLOutputVisitor) VisitAtVariableExpr(child ISQLVisitor, x *variable.SQLAtVariableExpr) bool {
	v.WriteKeyword(SYMB_AT)
	Accept(v, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitAtAtVariableExpr(child ISQLVisitor, x *variable.SQLAtAtVariableExpr) bool {
	if x.HasAtAT {
		v.WriteKeyword(SYMB_AT)
		v.WriteKeyword(SYMB_AT)

		if x.Kind != "" {
			v.Write(string(x.Kind))
			v.WriteKeyword(SYMB_DOT)
		}
		Accept(v, x.Name())
	} else {
		v.Write(string(x.Kind))
		WriteSpaceAfterAccept(child, x.Name())
	}

	return false
}
func (v *SQLOutputVisitor) VisitBindVariableExpr(child ISQLVisitor, x *variable.SQLBindVariableExpr) bool {
	v.Write(":")
	Accept(v, x.Name())
	return false
}

// ---------------------------- Variable End ------------------------------------

// ---------------------------- Operator Start ------------------------------------
func (v *SQLOutputVisitor) VisitUnaryOperatorExpr(child ISQLVisitor, x *operator.SQLUnaryOperatorExpr) bool {
	switch x.Operator {
	default:
		v.Write(string(x.Operator))
		WriteSpaceAfterAccept(child, x.Operand())
	}
	return false
}
func (v *SQLOutputVisitor) VisitBinaryOperatorExpr(child ISQLVisitor, x *operator.SQLBinaryOperatorExpr) bool {
	if x.Paren() {
		v.Write("(")
	}
	Accept(v, x.Left())
	v.WriteSpaceAfterValue(string(x.Operator))
	WriteSpaceAfterAccept(child, x.Right())
	if x.Paren() {
		v.Write(")")
	}
	return false
}

// ---------------------------- Operator End ------------------------------------

// ---------------------------- Condition Start ------------------------------------
func (v *SQLOutputVisitor) VisitIsCondition(child ISQLVisitor, x *condition.SQLIsCondition) bool {
	Accept(v, x.Expr())
	if x.Not {
		v.WriteSpaceAfterKeyword(NOT)
	}
	v.WriteSpaceAfterKeyword(IS)
	v.WriteSpaceAfterValue(string(x.Value))
	return false
}
func (v *SQLOutputVisitor) VisitLikeCondition(child ISQLVisitor, x *condition.SQLLikeCondition) bool {
	Accept(v, x.Expr())
	if x.Not {
		v.WriteSpaceAfterKeyword(NOT)
	}
	v.WriteSpaceAfterValue(string(x.Like))
	WriteSpaceAfterAccept(child, x.Pattern())

	if x.Escape() != nil {
		v.WriteSpaceAfterKeyword(ESCAPE)
		WriteSpaceAfterAccept(child, x.Escape())
	}

	return false
}
func (v *SQLOutputVisitor) VisitBetweenCondition(child ISQLVisitor, x *condition.SQLBetweenCondition) bool {
	Accept(v, x.Expr())
	if x.Not {
		v.WriteSpaceAfterKeyword(NOT)
	}
	v.WriteSpaceAfterKeyword(BETWEEN)
	WriteSpaceAfterAccept(child, x.Between())

	v.WriteSpaceAfterKeyword(AND)
	WriteSpaceAfterAccept(child, x.And())

	return false
}
func (v *SQLOutputVisitor) VisitInCondition(child ISQLVisitor, x *condition.SQLInCondition) bool {
	Accept(v, x.Expr())
	v.WriteSpaceAfterKeyword(IN)
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Values()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteSpace()
		}
		value := x.Value(i)
		switch value.(type) {
		case select_.ISQLSelectQuery:
			WriteIndentLnAfterAccept(child, false, value)
		default:
			Accept(v, value)
		}

	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitIsASetCondition(child ISQLVisitor, x *condition.SQLIsASetCondition) bool {
	return false
}

// ---------------------------- Condition End ------------------------------------

// ---------------------------- Expr Start ------------------------------------
func (v *SQLOutputVisitor) VisitDBLinkExpr(child ISQLVisitor, x *expr.SQLDBLinkExpr) bool {
	Accept(v, x.Name())
	v.WriteKeyword(SYMB_AT)
	Accept(v, x.DBLink())
	return false
}
func (v *SQLOutputVisitor) VisitAllColumnExpr(child ISQLVisitor, x *expr.SQLAllColumnExpr) bool {
	v.Write("*")
	return false
}
func (v *SQLOutputVisitor) VisitNullExpr(child ISQLVisitor, x *expr.SQLNullExpr) bool {
	v.WriteKeyword(NULL)
	return false
}
func (v *SQLOutputVisitor) VisitListExpr(child ISQLVisitor, x *expr.SQLListExpr) bool {
	v.WriteKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Elements()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteSpace()
		}
		Accept(v, x.Element(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitSubQuery(child ISQLVisitor, x *common.SQLSubQueryExpr) bool {
	WriteIndentLnAfterAccept(child, true, x.Query())
	return false
}
func (v *SQLOutputVisitor) VisitAssignExpr(child ISQLVisitor, x *expr.SQLAssignExpr) bool {
	Accept(v, x.Name())
	if x.Equal {
		v.WriteSpaceAfterKeyword(SYMB_EQUAL)
	}
	WriteSpaceAfterAccept(child, x.Value())
	return false
}

// ---------------------------- Expr End ------------------------------------

// ---------------------------- Function Start ------------------------------------
func (v *SQLOutputVisitor) VisitMethodInvocation(child ISQLVisitor, x *function.SQLMethodInvocation) bool {
	Accept(v, x.Name())

	if x.Paren {
		v.WriteKeyword(SYMB_LEFT_PAREN)
	}
	for i := 0; i < len(x.Arguments()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteSpace()
		}

		Accept(v, x.Arguments()[i])
	}
	if x.Paren {
		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}

	return false
}
func (v *SQLOutputVisitor) VisitStaticMethodInvocation(child ISQLVisitor, x *function.SQLStaticMethodInvocation) bool {

	return true
}

func (v *SQLOutputVisitor) VisitCastFunctionArgument(child ISQLVisitor, x *function.SQLCastFunctionArgument) bool {
	Accept(v, x.Expr())
	v.WriteSpaceAfterKeyword(AS)
	WriteSpaceAfterAccept(child, x.DataType())
	return false
}

// ---------------------------- Function End ------------------------------------

// ---------------------------- DataType Start ------------------------------------
func (v *SQLOutputVisitor) VisitDataType(child ISQLVisitor, x *datatype.SQLDataType) bool {
	Accept(v, x.Name())

	if x.Paren {
		v.WriteKeyword(SYMB_LEFT_PAREN)
	}
	for i := 0; i < len(x.Precisions()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteSpace()
		}
		Accept(v, x.Precisions()[i])
	}

	if x.Paren {
		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}

	return false
}
func (v *SQLOutputVisitor) VisitIntervalDataType(child ISQLVisitor, x *datatype.SQLIntervalDataType) bool {
	Accept(v, x.Start())
	WriteSpaceAfterAccept(child, x.End())
	return false
}
func (v *SQLOutputVisitor) VisitIntervalDataTypeField(child ISQLVisitor, x *datatype.SQLIntervalDataTypeField) bool {

	for i := 0; i < len(x.Precisions()); i++ {

	}
	return false
}
func (v *SQLOutputVisitor) VisitDateDataType(child ISQLVisitor, x *datatype.SQLDateDataType) bool {
	v.WriteKeyword(DATE)
	return false
}
func (v *SQLOutputVisitor) VisitDateTimeDataType(child ISQLVisitor, x *datatype.SQLDateTimeDataType) bool {
	v.WriteKeyword(DATETIME)
	return false
}
func (v *SQLOutputVisitor) VisitTimeDataType(child ISQLVisitor, x *datatype.SQLTimeDataType) bool {
	v.WriteKeyword(TIME)
	return false
}
func (v *SQLOutputVisitor) VisitTimestampDataType(child ISQLVisitor, x *datatype.SQLTimestampDataType) bool {
	v.WriteKeyword(TIMESTAMP)
	return false
}

// ---------------------------- DataType End ------------------------------------

// ---------------------------- Index Start ------------------------------------
func (v *SQLOutputVisitor) VisitIndexColumn(child ISQLVisitor, x *index.SQLIndexColumn) bool {
	Accept(v, x.Expr())

	return false
}

// ---------------------------- Index End ------------------------------------

// ---------------------------- Sequence Start ------------------------------------
func (v *SQLOutputVisitor) VisitIncrementBySequenceOption(child ISQLVisitor, x *exprSequence.SQLIncrementBySequenceOption) bool {
	v.WriteKeyword(INCREMENT)
	v.WriteSpaceAfterKeyword(BY)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitStartWithSequenceOption(child ISQLVisitor, x *exprSequence.SQLStartWithSequenceOption) bool {
	v.WriteKeyword(START)
	v.WriteSpaceAfterKeyword(WITH)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitMaxValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLMaxValueSequenceOption) bool {
	v.WriteKeyword(MAXVALUE)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitNoMaxValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoMaxValueSequenceOption) bool {
	v.WriteKeyword(NO)
	v.WriteSpaceAfterKeyword(MAXVALUE)
	return false
}
func (v *SQLOutputVisitor) VisitMinValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLMinValueSequenceOption) bool {
	v.WriteKeyword(MINVALUE)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitNoMinValueSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoMinValueSequenceOption) bool {
	v.WriteKeyword(NO)
	v.WriteKeyword(MINVALUE)
	return false
}
func (v *SQLOutputVisitor) VisitCycleSequenceOption(child ISQLVisitor, x *exprSequence.SQLCycleSequenceOption) bool {
	v.WriteKeyword(CYCLE)
	return false
}
func (v *SQLOutputVisitor) VisitNoCycleSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoCycleSequenceOption) bool {
	v.WriteKeyword(NOCYCLE)
	return false
}
func (v *SQLOutputVisitor) VisitCacheSequenceOption(child ISQLVisitor, x *exprSequence.SQLCacheSequenceOption) bool {
	v.WriteKeyword(CACHE)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitNoCacheSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoCacheSequenceOption) bool {
	v.WriteKeyword(NOCACHE)
	return true
}
func (v *SQLOutputVisitor) VisitOrderSequenceOption(child ISQLVisitor, x *exprSequence.SQLOrderSequenceOption) bool {
	v.WriteKeyword(ORDER)
	return false
}
func (v *SQLOutputVisitor) VisitNoOrderSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoOrderSequenceOption) bool {
	v.WriteKeyword(NOORDER)
	return false
}
func (v *SQLOutputVisitor) VisitKeepSequenceOption(child ISQLVisitor, x *exprSequence.SQLKeepSequenceOption) bool {
	v.WriteKeyword(KEEP)
	return false
}
func (v *SQLOutputVisitor) VisitNoKeepSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoKeepSequenceOption) bool {
	v.WriteKeyword(NOKEEP)
	return false
}
func (v *SQLOutputVisitor) VisitScaleSequenceOption(child ISQLVisitor, x *exprSequence.SQLScaleSequenceOption) bool {
	v.WriteKeyword(SCALE)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitNoScaleSequenceOption(child ISQLVisitor, x *exprSequence.SQLNoScaleSequenceOption) bool {
	v.WriteKeyword(NOSCALE)
	return false
}
func (v *SQLOutputVisitor) VisitSessionSequenceOption(child ISQLVisitor, x *exprSequence.SQLSessionSequenceOption) bool {
	v.WriteKeyword(SESSION)
	return false
}
func (v *SQLOutputVisitor) VisitGlobalSequenceOption(child ISQLVisitor, x *exprSequence.SQLGlobalSequenceOption) bool {
	v.WriteKeyword(GLOBAL)
	return false
}

// ---------------------------- Sequence End ------------------------------------

// ---------------------------- Server Start ------------------------------------
func (v *SQLOutputVisitor) VisitHostOption(child ISQLVisitor, x *exprServer.SQLHostOption) bool {
	v.WriteKeyword(HOST)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitDatabaseOption(child ISQLVisitor, x *exprServer.SQLDatabaseOption) bool {
	v.WriteKeyword(DATABASE)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitUserOption(child ISQLVisitor, x *exprServer.SQLUserOption) bool {
	v.WriteKeyword(USER)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitPasswordOption(child ISQLVisitor, x *exprServer.SQLPasswordOption) bool {
	v.WriteKeyword(PASSWORD)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitSocketOption(child ISQLVisitor, x *exprServer.SQLSocketOption) bool {
	v.WriteKeyword(SOCKET)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitOwnerOption(child ISQLVisitor, x *exprServer.SQLOwnerOption) bool {
	v.WriteKeyword(OWNER)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitPortOption(child ISQLVisitor, x *exprServer.SQLPortOption) bool {
	v.WriteKeyword(PORT)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}

// ---------------------------- Server End ------------------------------------

// ---------------------------- OnTable Start ------------------------------------
func (v *SQLOutputVisitor) VisitTableColumn(child ISQLVisitor, x *exprTable.SQLTableColumn) bool {
	Accept(v, x.Name())
	WriteSpaceAfterAccept(child, x.DataType())

	for i := 0; i < len(x.Options()); i++ {
		WriteSpaceAfterAccept(child, x.Option(i))
	}
	return false
}
func (v *SQLOutputVisitor) VisitPrimaryKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLPrimaryKeyTableConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}

	v.WriteKeyword(PRIMARY)
	v.WriteSpaceAfterKeyword(KEY)

	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Columns()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Column(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)

	return false
}
func (v *SQLOutputVisitor) VisitUniqueTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueTableConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}

	v.WriteKeyword(UNIQUE)

	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Columns()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Column(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitUniqueIndexTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueIndexTableConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}

	v.WriteKeyword(UNIQUE)
	v.WriteSpaceAfterKeyword(INDEX)

	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Columns()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Column(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitUniqueKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLUniqueKeyTableConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}

	v.WriteKeyword(UNIQUE)
	v.WriteSpaceAfterKeyword(KEY)

	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Columns()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Column(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitForeignKeyTableConstraint(child ISQLVisitor, x *exprTable.SQLForeignKeyTableConstraint) bool {

	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}

	v.WriteKeyword(FOREIGN)
	v.WriteSpaceAfterKeyword(KEY)

	return false
}
func (v *SQLOutputVisitor) VisitCheckTableConstraint(child ISQLVisitor, x *exprTable.SQLCheckTableConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}
	v.WriteKeyword(CHECK)
	v.WriteKeyword(SYMB_LEFT_PAREN)
	Accept(v, x.Condition())
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}

func (v *SQLOutputVisitor) VisitTableLikeClause(child ISQLVisitor, x *exprTable.SQLTableLikeClause) bool {
	v.WriteKeyword(LIKE)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitPrimaryKeyColumnConstraint(child ISQLVisitor, x *exprTable.SQLPrimaryKeyColumnConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}
	v.WriteKeyword(PRIMARY)
	v.WriteSpaceAfterKeyword(KEY)
	return false
}
func (v *SQLOutputVisitor) VisitKeyColumnConstraint(child ISQLVisitor, x *exprTable.SQLKeyColumnConstraint) bool {
	v.WriteKeyword(KEY)
	return false
}
func (v *SQLOutputVisitor) VisitUniqueColumnConstraint(child ISQLVisitor, x *exprTable.SQLUniqueColumnConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}
	v.WriteKeyword(UNIQUE)
	if x.Key {
		v.WriteSpaceAfterKeyword(KEY)
	}

	return false
}
func (v *SQLOutputVisitor) VisitNullColumnConstraint(child ISQLVisitor, x *exprTable.SQLNullColumnConstraint) bool {
	v.WriteKeyword(NULL)
	return false
}
func (v *SQLOutputVisitor) VisitNotNullColumnConstraint(child ISQLVisitor, x *exprTable.SQLNotNullColumnConstraint) bool {
	v.WriteKeyword(NOT)
	v.WriteSpaceAfterKeyword(NULL)
	return false
}
func (v *SQLOutputVisitor) VisitCheckColumnConstraint(child ISQLVisitor, x *exprTable.SQLCheckColumnConstraint) bool {
	if x.Name() != nil {
		v.WriteKeyword(CONSTRAINT)
		WriteSpaceAfterAccept(child, x.Name())
		v.WriteSpace()
	}
	v.WriteKeyword(CHECK)
	return false
}

func (v *SQLOutputVisitor) VisitDefaultClause(child ISQLVisitor, x *exprTable.SQLDefaultClause) bool {
	v.WriteKeyword(DEFAULT)
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitAutoIncrementExpr(child ISQLVisitor, x *exprTable.SQLAutoIncrementExpr) bool {
	v.WriteKeyword(AUTO_INCREMENT)
	return false
}

func (v *SQLOutputVisitor) VisitVisibleExpr(child ISQLVisitor, x *exprTable.SQLVisibleExpr) bool {
	v.WriteKeyword(VISIBLE)
	return false
}
func (v *SQLOutputVisitor) VisitInvisibleExpr(child ISQLVisitor, x *exprTable.SQLInvisibleExpr) bool {
	v.WriteKeyword(INVISIBLE)
	return false
}
func (v *SQLOutputVisitor) VisitCommentExpr(child ISQLVisitor, x *exprTable.SQLCommentExpr) bool {
	v.WriteKeyword(COMMENT)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}

func (v *SQLOutputVisitor) VisitCharsetAssignExpr(child ISQLVisitor, x *exprTable.SQLCharsetAssignExpr) bool {
	if x.Default {
		v.WriteKeyword(DEFAULT)
		v.WriteSpace()
	}
	v.WriteKeyword(CHARSET)
	if x.Equal {
		v.WriteSpaceAfterKeyword(SYMB_EQUAL)
	}
	WriteSpaceAfterAccept(child, x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitCharacterSetAssignExpr(child ISQLVisitor, x *exprTable.SQLCharacterSetAssignExpr) bool {
	if x.Default {
		v.WriteKeyword(DEFAULT)
		v.WriteSpace()
	}
	v.WriteKeyword(CHARACTER)
	v.WriteSpaceAfterKeyword(SET)
	if x.Equal {
		v.WriteSpaceAfterKeyword(SYMB_EQUAL)
	}
	WriteSpaceAfterAccept(child, x.Value())
	return false
}

func (v *SQLOutputVisitor) VisitPartitionByHash(child ISQLVisitor, x *exprTable.SQLPartitionByHash) bool {
	v.WriteKeyword(PARTITION)
	v.WriteSpaceAfterKeyword(BY)
	v.WriteSpaceAfterKeyword(HASH)

	v.WritePartitionByColumns(x.Columns())

	v.WritePartitionByPartitionsNum(child, x.PartitionsNum())

	if x.SubPartitionBy() != nil {
		WriteLnAfterAccept(child, x.SubPartitionBy())
	}

	v.WritePartitionDefinitions(x.PartitionDefinitions())
	return false
}
func (v *SQLOutputVisitor) VisitPartitionByKey(child ISQLVisitor, x *exprTable.SQLPartitionByKey) bool {
	v.WriteKeyword(PARTITION)
	v.WriteSpaceAfterKeyword(BY)
	v.WriteSpaceAfterKeyword(KEY)

	v.WritePartitionByColumns(x.Columns())

	v.WritePartitionByPartitionsNum(child, x.PartitionsNum())

	if x.SubPartitionBy() != nil {
		WriteLnAfterAccept(child, x.SubPartitionBy())
	}

	v.WritePartitionDefinitions(x.PartitionDefinitions())
	return false
}
func (v *SQLOutputVisitor) VisitPartitionByRange(child ISQLVisitor, x *exprTable.SQLPartitionByRange) bool {
	v.WriteKeyword(PARTITION)
	v.WriteSpaceAfterKeyword(BY)
	v.WriteSpaceAfterKeyword(RANGE)

	v.WritePartitionByColumns(x.Columns())

	v.WritePartitionByPartitionsNum(child, x.PartitionsNum())

	if x.SubPartitionBy() != nil {
		WriteLnAfterAccept(child, x.SubPartitionBy())
	}

	v.WritePartitionDefinitions(x.PartitionDefinitions())

	return false
}
func (v *SQLOutputVisitor) VisitPartitionByList(child ISQLVisitor, x *exprTable.SQLPartitionByList) bool {
	v.WriteKeyword(PARTITION)
	v.WriteSpaceAfterKeyword(BY)
	v.WriteSpaceAfterKeyword(LIST)

	v.WritePartitionByColumns(x.Columns())

	v.WritePartitionByPartitionsNum(child, x.PartitionsNum())

	if x.SubPartitionBy() != nil {
		WriteLnAfterAccept(child, x.SubPartitionBy())
	}

	v.WritePartitionDefinitions(x.PartitionDefinitions())
	return false
}
func (v *SQLOutputVisitor) WritePartitionByColumns(columns []expr.ISQLExpr) {
	if columns == nil || len(columns) == 0 {
		return
	}
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(columns); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteSpace()
		}
		Accept(v, columns[i])
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
}
func (v *SQLOutputVisitor) WritePartitionByPartitionsNum(child ISQLVisitor, partitionsNum expr.ISQLExpr) {
	if partitionsNum == nil {
		return
	}
	v.WriteSpaceAfterKeyword(PARTITIONS)
	WriteSpaceAfterAccept(child, partitionsNum)
}
func (v *SQLOutputVisitor) WritePartitionDefinitions(partitionDefinitions []*exprTable.SQLPartitionDefinition) {
	if partitionDefinitions == nil ||
		len(partitionDefinitions) == 0 {
		return
	}
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	v.IncrementIndentAndWriteLn()
	for i := 0; i < len(partitionDefinitions); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}
		Accept(v, partitionDefinitions[i])
	}
	v.DecrementIndentAndWriteLn()
	v.WriteKeyword(SYMB_RIGHT_PAREN)
}

func (v *SQLOutputVisitor) VisitSubPartitionByHash(child ISQLVisitor, x *exprTable.SQLSubPartitionByHash) bool {
	v.WriteKeyword(SUBPARTITION)
	v.WriteSpaceAfterKeyword(BY)

	if x.Linear {
		v.WriteSpaceAfterKeyword(LINEAR)
	}

	v.WriteSpaceAfterKeyword(HASH)

	v.WriteSubPartitionByColumns(x.Columns())

	v.WriteSubPartitionBySubPartitionsNum(child, x.SubPartitionsNum())

	return false
}
func (v *SQLOutputVisitor) VisitSubPartitionByKey(child ISQLVisitor, x *exprTable.SQLSubPartitionByKey) bool {
	v.WriteKeyword(SUBPARTITION)
	v.WriteSpaceAfterKeyword(BY)

	if x.Linear {
		v.WriteSpaceAfterKeyword(LINEAR)
	}

	v.WriteSpaceAfterKeyword(KEY)

	v.WriteSubPartitionByColumns(x.Columns())

	v.WriteSubPartitionBySubPartitionsNum(child, x.SubPartitionsNum())

	return false
}
func (v *SQLOutputVisitor) WriteSubPartitionByColumns(columns []expr.ISQLExpr) {
	if columns == nil || len(columns) == 0 {
		return
	}
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	v.IncrementIndentAndWriteLn()
	for i := 0; i < len(columns); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}
		Accept(v, columns[i])
	}
	v.DecrementIndentAndWriteLn()
	v.WriteKeyword(SYMB_RIGHT_PAREN)
}
func (v *SQLOutputVisitor) WriteSubPartitionBySubPartitionsNum(child ISQLVisitor, subPartitionsNum expr.ISQLExpr) {
	if subPartitionsNum == nil {
		return
	}
	v.WriteSpaceAfterKeyword(SUBPARTITIONS)
	WriteSpaceAfterAccept(child, subPartitionsNum)
}

func (v *SQLOutputVisitor) VisitPartitionDefinition(child ISQLVisitor, x *exprTable.SQLPartitionDefinition) bool {
	v.WriteKeyword(PARTITION)
	WriteSpaceAfterAccept(child, x.Name())

	WriteSpaceAfterAccept(child, x.Values())

	for i := 0; i < len(x.SubpartitionDefinitions()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}
		Accept(v, x.SubpartitionDefinition(i))
	}

	return false
}

func (v *SQLOutputVisitor) WritePartitionValues(values []expr.ISQLExpr) {
	if values == nil ||
		len(values) == 0 {
		return
	}
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	v.IncrementIndentAndWriteLn()
	for i := 0; i < len(values); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteSpace()
		}
		Accept(v, values[i])
	}
	v.DecrementIndentAndWriteLn()
	v.WriteKeyword(SYMB_RIGHT_PAREN)
}
func (v *SQLOutputVisitor) VisitPartitionValuesLessThan(child ISQLVisitor, x *exprTable.SQLPartitionValuesLessThan) bool {
	v.WriteKeyword(VALUES)
	v.WriteSpaceAfterKeyword(LESS)
	v.WriteSpaceAfterKeyword(THAN)

	v.WritePartitionValues(x.Values())
	return false
}
func (v *SQLOutputVisitor) VisitPartitionValuesLessThanMaxValue(child ISQLVisitor, x *exprTable.SQLPartitionValuesLessThanMaxValue) bool {
	v.WriteKeyword(VALUES)
	v.WriteSpaceAfterKeyword(LESS)
	v.WriteSpaceAfterKeyword(THAN)

	v.WriteSpaceAfterKeyword(MAXVALUE)
	return false
}
func (v *SQLOutputVisitor) VisitPartitionValuesIn(child ISQLVisitor, x *exprTable.SQLPartitionValuesIn) bool {
	v.WriteKeyword(VALUES)
	v.WriteSpaceAfterKeyword(IN)

	v.WritePartitionValues(x.Values())
	return false
}

func (v *SQLOutputVisitor) VisitSubPartitionDefinition(child ISQLVisitor, x *exprTable.SQLSubPartitionDefinition) bool {
	v.WriteKeyword(SUBPARTITION)
	WriteSpaceAfterAccept(child, x.Name())

	return false
}

// ----- Alter OnTable
func (v *SQLOutputVisitor) VisitAddColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddColumnAlterTableAction) bool {
	v.WriteKeyword(ADD)
	if x.HasColumn {
		v.WriteSpaceAfterKeyword(COLUMN)
	}

	paren := x.Paren || len(x.Columns()) > 1
	if paren {
		v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
		v.IncrementIndent()
	}
	for i := 0; i < len(x.Columns()); i++ {
		if paren {
			WriteLnAfterAccept(child, x.Column(i))
		} else {
			WriteSpaceAfterAccept(child, x.Column(i))
		}

	}
	if paren {
		v.DecrementIndentAndWriteLn()
		v.WriteKeyword(SYMB_LEFT_PAREN)
	}
	return false
}

func (v *SQLOutputVisitor) VisitAlterColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLAlterColumnAlterTableAction) bool {
	v.WriteKeyword(ALTER)
	if x.HasColumn {
		v.WriteSpaceAfterKeyword(COLUMN)
	}
	WriteSpaceAfterAccept(child, x.Column())
	WriteSpaceAfterAccept(child, x.Action())

	return false
}
func (v *SQLOutputVisitor) VisitDropColumnAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropColumnAlterTableAction) bool {
	v.WriteKeyword(DROP)
	if x.HasColumn {
		v.WriteSpaceAfterKeyword(COLUMN)
	}
	WriteSpaceAfterAccept(child, x.Column())
	return false
}
func (v *SQLOutputVisitor) VisitAddTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddTableConstraintAlterTableAction) bool {
	return false
}
func (v *SQLOutputVisitor) VisitDropIndexAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropIndexAlterTableAction) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(INDEX)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitDropKeyAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropKeyAlterTableAction) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(KEY)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitDropConstraintTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropTableConstraintAlterTableAction) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(CONSTRAINT)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitDropPrimaryKeyTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropPrimaryKeyTableConstraintAlterTableAction) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(PRIMARY)
	v.WriteSpaceAfterKeyword(KEY)
	return false
}
func (v *SQLOutputVisitor) VisitDropUniqueTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropUniqueTableConstraintAlterTableAction) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(UNIQUE)
	return false
}
func (v *SQLOutputVisitor) VisitDropForeignKeyTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropForeignKeyTableConstraintAlterTableAction) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(FOREIGN)
	v.WriteSpaceAfterKeyword(KEY)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitDropCheckTableConstraintAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropCheckTableConstraintAlterTableAction) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(CHECK)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitAddPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLAddPartitionAlterTableAction) bool               { return false }
func (v *SQLOutputVisitor) VisitDropPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLDropPartitionAlterTableAction) bool             { return false }
func (v *SQLOutputVisitor) VisitDiscardPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLDiscardPartitionAlterTableAction) bool       { return false }
func (v *SQLOutputVisitor) VisitImportPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLImportPartitionAlterTableAction) bool         { return false }
func (v *SQLOutputVisitor) VisitTruncatePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLTruncatePartitionAlterTableAction) bool     { return false }
func (v *SQLOutputVisitor) VisitCoalescePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLCoalescePartitionAlterTableAction) bool     { return false }
func (v *SQLOutputVisitor) VisitReorganizePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLReorganizePartitionAlterTableAction) bool { return false }
func (v *SQLOutputVisitor) VisitExchangePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLExchangePartitionAlterTableAction) bool     { return false }
func (v *SQLOutputVisitor) VisitAnalyzePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLAnalyzePartitionAlterTableAction) bool       { return false }
func (v *SQLOutputVisitor) VisitCheckPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLCheckPartitionAlterTableAction) bool           { return false }
func (v *SQLOutputVisitor) VisitOptimizePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLOptimizePartitionAlterTableAction) bool     { return false }
func (v *SQLOutputVisitor) VisitRebuildPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRebuildPartitionAlterTableAction) bool       { return false }
func (v *SQLOutputVisitor) VisitRepairPartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRepairPartitionAlterTableAction) bool         { return false }
func (v *SQLOutputVisitor) VisitRemovePartitionAlterTableAction(child ISQLVisitor, x *exprTable.SQLRemovePartitionAlterTableAction) bool         { return false }

// ---------------------------- Table End ------------------------------------

// ---------------------------- User Start ------------------------------------
func (v *SQLOutputVisitor) VisitUserName(child ISQLVisitor, x *exprUser.SQLUserName) bool {
	Accept(v, x.Name())
	v.WriteKeyword(SYMB_AT)
	Accept(v, x.Host())

	WriteSpaceAfterAccept(child, x.Option())
	return false
}
func (v *SQLOutputVisitor) VisitIdentifiedByAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByAuthOption) bool {
	v.WriteKeyword(IDENTIFIED)
	v.WriteSpaceAfterKeyword(BY)
	WriteSpaceAfterAccept(child, x.Auth())
	return false
}
func (v *SQLOutputVisitor) VisitIdentifiedByPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByPasswordAuthOption) bool {
	v.WriteKeyword(IDENTIFIED)
	v.WriteSpaceAfterKeyword(BY)
	v.WriteSpaceAfterKeyword(PASSWORD)
	WriteSpaceAfterAccept(child, x.Auth())
	return false
}
func (v *SQLOutputVisitor) VisitIdentifiedByRandomPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedByRandomPasswordAuthOption) bool {
	v.WriteKeyword(IDENTIFIED)
	v.WriteSpaceAfterKeyword(BY)
	v.WriteSpaceAfterKeyword(RANDOM)
	v.WriteSpaceAfterKeyword(PASSWORD)
	return false
}
func (v *SQLOutputVisitor) VisitIdentifiedWithAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithAuthOption) bool {
	v.WriteKeyword(IDENTIFIED)
	v.WriteSpaceAfterKeyword(WITH)
	WriteSpaceAfterAccept(child, x.Plugin())
	return false
}
func (v *SQLOutputVisitor) VisitIdentifiedWithByAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithByAuthOption) bool {
	v.WriteKeyword(IDENTIFIED)
	v.WriteSpaceAfterKeyword(WITH)
	WriteSpaceAfterAccept(child, x.Auth())
	return false
}
func (v *SQLOutputVisitor) VisitIdentifiedWithByRandomPasswordAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithByRandomPasswordAuthOption) bool {
	v.WriteKeyword(IDENTIFIED)
	v.WriteSpaceAfterKeyword(WITH)
	WriteSpaceAfterAccept(child, x.Plugin())
	return false
}
func (v *SQLOutputVisitor) VisitIdentifiedWithAsAuthOption(child ISQLVisitor, x *exprUser.SQLIdentifiedWithAsAuthOption) bool {
	v.WriteKeyword(IDENTIFIED)
	v.WriteSpaceAfterKeyword(WITH)
	WriteSpaceAfterAccept(child, x.Auth())
	return false
}

// ---------------------------- User End ------------------------------------

// ---------------------------- View Start ------------------------------------
func (v *SQLOutputVisitor) VisitViewColumn(child ISQLVisitor, x *exprView.SQLViewColumn) bool {
	Accept(v, x.Name())

	return false
}

// ---------------------------- View End ------------------------------------

// ---------------------------- Select Start ------------------------------------
func (v *SQLOutputVisitor) VisitSelectQuery(child ISQLVisitor, x *select_.SQLSelectQuery) bool {

	Accept(v, x.WithClause())

	v.WriteKeyword(SELECT)

	v.WriteSelectElements(x.SelectElements())

	WriteLnAfterAccept(child, x.FromClause())

	WriteLnAfterAccept(child, x.WhereClause())

	WriteLnAfterAccept(child, x.HierarchicalQueryClause())

	WriteLnAfterAccept(child, x.GroupByClause())

	WriteLnAfterAccept(child, x.OrderByClause())

	WriteLnAfterAccept(child, x.LimitClause())

	WriteLnAfterAccept(child, x.LockClause())

	return false
}

func (v *SQLOutputVisitor) WriteSelectElements(x []*select_.SQLSelectElement) {
	start := v.Builder.Len()
	end := v.Builder.Len()
	v.IncrementIndent()
	for i := 0; i < len(x); i++ {
		if i == 0 {
			v.WriteSpace()
		}
		if i != 0 {
			v.Write(",")
			if end-start > v.SelectElementLimit {
				v.WriteLn()
				start = v.Builder.Len()
			} else {
				v.WriteSpace()
			}
		}

		child := x[i]
		Accept(v, child)

		end = v.Builder.Len()
	}
	v.DecrementIndent()
}

func (v *SQLOutputVisitor) VisitParenSelectQuery(child ISQLVisitor, x *select_.SQLParenSelectQuery) bool {
	v.Write("(")

	v.IncrementIndent()
	v.WriteLn()

	Accept(v, x.SubQuery())

	v.DecrementIndent()
	v.WriteLn()
	v.Write(")")

	WriteLnAfterAccept(child, x.OrderByClause())
	WriteLnAfterAccept(child, x.LimitClause())

	return false
}

func (v *SQLOutputVisitor) VisitSelectUnionQuery(child ISQLVisitor, x *select_.SQLSelectUnionQuery) bool {

	Accept(v, x.Left())

	v.WriteLnAfterValue(string(x.Operator))

	WriteLnAfterAccept(child, x.Right())

	WriteLnAfterAccept(child, x.OrderByClause())
	WriteLnAfterAccept(child, x.LimitClause())

	return false
}

func (v *SQLOutputVisitor) VisitWithClause(child ISQLVisitor, x *select_.SQLWithClause) bool {
	v.WriteKeyword(WITH)

	if x.Recursive {
		v.WriteSpaceAfterKeyword(RECURSIVE)
	}

	v.IncrementIndent()
	v.WriteLn()

	for i := 0; i < len(x.FactoringClause()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}
		Accept(v, x.FactoringClause()[i])
	}

	v.DecrementIndent()
	v.WriteLn()
	return false
}

// cte_name [(col_name [, col_name] ...)] AS (subquery)
func (v *SQLOutputVisitor) VisitSubQueryFactoringClause(child ISQLVisitor, x *select_.SQLSubQueryFactoringClause) bool {
	Accept(v, x.Name())

	if len(x.Columns()) > 0 {
		v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)

		for i := 0; i < len(x.Columns()); i++ {
			if i != 0 {
				v.WriteKeyword(SYMB_COMMA)
				v.WriteSpace()
			}
			Accept(v, x.Columns()[i])
		}

		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}

	v.WriteSpaceAfterKeyword(AS)
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)

	v.IncrementIndent()

	WriteLnAfterAccept(child, x.SubQuery())

	v.DecrementIndent()
	v.WriteLnAfterKeyword(SYMB_RIGHT_PAREN)

	return false
}

func (v *SQLOutputVisitor) VisitSearchClause(child ISQLVisitor, x *select_.SQLSearchClause) bool {
	return false
}

func (v *SQLOutputVisitor) VisitSubAvFactoringClause(child ISQLVisitor, x *select_.SQLSubAvFactoringClause) bool {
	return false
}

func (v *SQLOutputVisitor) VisitSubAvClause(child ISQLVisitor, x *select_.SQLSubAvClause) bool {
	return false
}

func (v *SQLOutputVisitor) VisitSelectElement(child ISQLVisitor, x *select_.SQLSelectElement) bool {
	Accept(v, x.Expr())

	if x.As {
		v.WriteSpaceAfterValue("AS")
	}

	WriteSpaceAfterAccept(child, x.Alias())

	return false
}
func (v *SQLOutputVisitor) VisitSelectTargetElement(child ISQLVisitor, x *select_.SQLSelectTargetElement) bool {

	return false
}

func (v *SQLOutputVisitor) VisitFromClause(child ISQLVisitor, x *select_.SQLFromClause) bool {
	v.WriteKeyword(FROM)
	WriteSpaceAfterAccept(child, x.TableReference())
	return false
}
func (v *SQLOutputVisitor) VisitTableReference(child ISQLVisitor, x *select_.SQLTableReference) bool {
	if x.Paren() {
		v.WriteKeyword(SYMB_LEFT_PAREN)
	}

	Accept(v, x.Name())

	WriteSpaceAfterAccept(child, x.PartitionExtensionClause())

	WriteSpaceAfterAccept(child, x.SampleClause())

	if x.As() {
		v.WriteSpaceAfterKeyword(AS)
	}

	WriteSpaceAfterAccept(child, x.Alias())

	if x.Paren() {
		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}
	return false
}
func (v *SQLOutputVisitor) VisitPartitionClause(child ISQLVisitor, x *select_.SQLPartitionClause) bool {
	v.WriteKeyword(PARTITION)
	v.WriteKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Name(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitPartitionForClause(child ISQLVisitor, x *select_.SQLPartitionForClause) bool {
	v.WriteKeyword(PARTITION)
	v.WriteSpaceAfterKeyword(FOR)
	v.WriteKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Name(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitSubPartitionClause(child ISQLVisitor, x *select_.SQLSubPartitionClause) bool {
	v.WriteKeyword(SUBPARTITION)
	v.WriteKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Name(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitSubPartitionForClause(child ISQLVisitor, x *select_.SQLSubPartitionForClause) bool {
	v.WriteKeyword(SUBPARTITION)
	v.WriteSpaceAfterKeyword(FOR)
	v.WriteKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Name(i))
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}

func (v *SQLOutputVisitor) VisitSampleClause(child ISQLVisitor, x *select_.SQLSampleClause) bool {
	v.WriteKeyword(SAMPLE)

	if x.Block {
		v.WriteSpaceAfterKeyword(BLOCK)
	}

	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	Accept(v, x.Percent())
	v.WriteKeyword(SYMB_RIGHT_PAREN)

	if x.SeedValue() != nil {
		v.WriteSpaceAfterKeyword(SEED)
		v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
		Accept(v, x.SeedValue())
		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}
	return false
}
func (v *SQLOutputVisitor) VisitOJTableReference(child ISQLVisitor, x *select_.SQLOJTableReference) bool {
	v.WriteKeyword(SYMB_LERT_BRACE)
	v.IncrementIndentAndWriteLn()

	v.WriteKeyword(OJ)
	WriteSpaceAfterAccept(child, x.TableReference())

	v.DecrementIndentAndWriteLn()
	v.WriteKeyword(SYMB_RIGHT_BRACE)

	return false
}
func (v *SQLOutputVisitor) VisitSubQueryTableReference(child ISQLVisitor, x *select_.SQLSubQueryTableReference) bool {
	if x.Paren() {
		v.WriteKeyword(SYMB_LEFT_PAREN)
		v.IncrementIndent()
		v.WriteLn()
	}

	Accept(v, x.SubQuery())

	if x.Paren() {
		v.DecrementIndent()
		v.WriteLn()
		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}

	if x.As() {
		v.WriteSpaceAfterKeyword(AS)
	}

	WriteSpaceAfterAccept(child, x.Alias())

	return false
}
func (v *SQLOutputVisitor) VisitJoinTableReference(child ISQLVisitor, x *select_.SQLJoinTableReference) bool {
	if x.Paren() {
		v.WriteKeyword(SYMB_LEFT_PAREN)
	}

	Accept(v, x.Left())

	if x.JoinType != select_.COMMA {
		v.WriteSpace()
	}
	v.Write(string(x.JoinType))

	WriteSpaceAfterAccept(child, x.Right())

	WriteSpaceAfterAccept(child, x.Condition())

	if x.Paren() {
		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}
	return false
}

func (v *SQLOutputVisitor) VisitJoinOnCondition(child ISQLVisitor, x *select_.SQLJoinOnCondition) bool {
	v.WriteKeyword(ON)
	WriteSpaceAfterAccept(child, x.Condition())
	return false
}
func (v *SQLOutputVisitor) VisitJoinUsingCondition(child ISQLVisitor, x *select_.SQLJoinUsingCondition) bool {
	v.WriteKeyword(USING)
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i := 0; i < len(x.Columns()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, x.Column(i))
	}
	v.WriteSpaceAfterKeyword(SYMB_RIGHT_PAREN)
	return false
}

func (v *SQLOutputVisitor) VisitWhereClause(child ISQLVisitor, x *select_.SQLWhereClause) bool {
	v.WriteKeyword(WHERE)
	WriteSpaceAfterAccept(child, x.Condition())
	return false
}

func (v *SQLOutputVisitor) VisitHierarchicalQueryClauseConnectBy(child ISQLVisitor, x *select_.SQLHierarchicalQueryClauseConnectBy) bool {
	// CONNECT BY [ NOCYCLE ] condition [ START WITH condition ]
	v.WriteKeyword(CONNECT)
	v.WriteSpaceAfterKeyword(BY)
	if x.NoCycle {
		v.WriteSpaceAfterKeyword(NOCYCLE)
	}
	WriteSpaceAfterAccept(child, x.ConnectByCondition())

	if x.StartWithCondition() != nil {
		v.WriteSpaceAfterKeyword(START)
		v.WriteSpaceAfterKeyword(WITH)
		WriteSpaceAfterAccept(child, x.StartWithCondition())
	}
	return false
}
func (v *SQLOutputVisitor) VisitHierarchicalQueryClauseStartWith(child ISQLVisitor, x *select_.SQLHierarchicalQueryClauseStartWith) bool {
	v.WriteKeyword(START)
	v.WriteSpaceAfterKeyword(WITH)
	WriteSpaceAfterAccept(child, x.StartWithCondition())

	v.WriteSpaceAfterKeyword(CONNECT)
	v.WriteSpaceAfterKeyword(BY)
	if x.NoCycle {
		v.WriteSpaceAfterKeyword(NOCYCLE)
	}
	WriteSpaceAfterAccept(child, x.ConnectByCondition())

	return false
}

func (v *SQLOutputVisitor) VisitGroupByHavingClause(child ISQLVisitor, x *select_.SQLGroupByHavingClause) bool {
	v.WriteKeyword(GROUP)
	v.WriteSpaceAfterKeyword(BY)
	for i := 0; i < len(x.Elements()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Element(i))
	}

	if x.Having() != nil {
		v.WriteSpaceAfterKeyword(HAVING)
		WriteSpaceAfterAccept(child, x.Having())
	}

	return false
}
func (v *SQLOutputVisitor) VisitHavingGroupByClause(child ISQLVisitor, x *select_.SQLHavingGroupByClause) bool {
	v.WriteKeyword(HAVING)
	WriteSpaceAfterAccept(child, x.Having())

	if len(x.Elements()) > 0 {

	}
	return false
}
func (v *SQLOutputVisitor) VisitGroupByElement(child ISQLVisitor, x *select_.SQLGroupByElement) bool {
	Accept(v, x.Expr())
	return false
}
func (v *SQLOutputVisitor) VisitOrderByClause(child ISQLVisitor, x *select_.SQLOrderByClause) bool {
	v.WriteKeyword(ORDER)
	v.WriteSpaceAfterKeyword(BY)

	for i := 0; i < len(x.Elements()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Element(i))
	}
	return false
}
func (v *SQLOutputVisitor) VisitOrderByElement(child ISQLVisitor, x *select_.SQLOrderByElement) bool {
	Accept(v, x.Key())
	v.WriteSpaceAfterValue(string(x.Specification))
	v.WriteSpaceAfterValue(string(x.NullOrdering))
	return false
}
func (v *SQLOutputVisitor) VisitLimitOffsetClause(child ISQLVisitor, x *select_.SQLLimitOffsetClause) bool {
	v.WriteKeyword(LIMIT)

	if x.Offset {
		WriteSpaceAfterAccept(child, x.CountExpr())
		v.WriteSpaceAfterKeyword(OFFSET)
		WriteSpaceAfterAccept(child, x.OffsetExpr())

	} else {

		WriteSpaceAfterAccept(child, x.OffsetExpr())

		if x.OffsetExpr() != nil {
			v.WriteKeyword(SYMB_COMMA)
		}

		WriteSpaceAfterAccept(child, x.CountExpr())
	}
	return false
}
func (v *SQLOutputVisitor) VisitOffsetFetchClause(child ISQLVisitor, x *select_.SQLOffsetFetchClause) bool {

	return false
}

func (v *SQLOutputVisitor) VisitForUpdateClause(child ISQLVisitor, x *select_.SQLForUpdateClause) bool {
	v.WriteKeyword(FOR)
	v.WriteSpaceAfterKeyword(UPDATE)

	if len(x.Tables()) > 0 {
		v.WriteSpaceAfterKeyword(OF)
		for i := 0; i < len(x.Tables()); i++ {
			if i != 0 {
				v.WriteKeyword(SYMB_COMMA)
			}
			WriteSpaceAfterAccept(child, x.Table(i))
		}
	}

	WriteSpaceAfterAccept(child, x.WaitExpr())
	return false
}
func (v *SQLOutputVisitor) VisitForShareClause(child ISQLVisitor, x *select_.SQLForShareClause) bool {
	v.WriteKeyword(FOR)
	v.WriteSpaceAfterKeyword(SHARE)

	if len(x.Tables()) > 0 {
		v.WriteSpaceAfterKeyword(OF)
		for i := 0; i < len(x.Tables()); i++ {
			if i != 0 {
				v.WriteKeyword(SYMB_COMMA)
			}
			WriteSpaceAfterAccept(child, x.Table(i))
		}
	}

	WriteSpaceAfterAccept(child, x.WaitExpr())
	return false
}
func (v *SQLOutputVisitor) VisitLockInShareModeClause(child ISQLVisitor, x *select_.SQLLockInShareModeClause) bool {
	v.WriteKeyword(LOCK)
	v.WriteSpaceAfterKeyword(IN)
	v.WriteSpaceAfterKeyword(SHARE)
	v.WriteSpaceAfterKeyword(MODE)
	return false
}

// ---------------------------- Select End ------------------------------------

// ---------------------------- Statement Expr Start ------------------------------------
func (v *SQLOutputVisitor) VisitEditionAbleExpr(child ISQLVisitor, x *statement.SQLEditionAbleExpr) bool {
	v.WriteKeyword(EDITIONABLE)
	return false
}
func (v *SQLOutputVisitor) VisitNonEditionAbleExpr(child ISQLVisitor, x *statement.SQLNonEditionAbleExpr) bool {
	v.WriteKeyword(NONEDITIONABLE)
	return false
}
func (v *SQLOutputVisitor) VisitCompileExpr(child ISQLVisitor, x *statement.SQLCompileExpr) bool {
	v.WriteKeyword(COMPILE)
	return false
}

// ---------------------------- Statement Expr End ------------------------------------

// ---------------------------- Statement Start ------------------------------------
func (v *SQLOutputVisitor) VisitCommentOnAuditPolicyStatement(child ISQLVisitor, x *comment.SQLCommentOnAuditPolicyStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(AUDIT)
	v.WriteSpaceAfterKeyword(POLICY)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}
func (v *SQLOutputVisitor) VisitCommentOnColumnStatement(child ISQLVisitor, x *comment.SQLCommentOnColumnStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(COLUMN)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}
func (v *SQLOutputVisitor) VisitCommentOnEditionStatement(child ISQLVisitor, x *comment.SQLCommentOnEditionStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(EDITION)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}
func (v *SQLOutputVisitor) VisitCommentOnIndextypeStatement(child ISQLVisitor, x *comment.SQLCommentOnIndextypeStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(INDEXTYPE)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}
func (v *SQLOutputVisitor) VisitCommentOnMaterializedViewStatement(child ISQLVisitor, x *comment.SQLCommentOnMaterializedViewStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(MATERIALIZED)
	v.WriteSpaceAfterKeyword(VIEW)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}
func (v *SQLOutputVisitor) VisitCommentOnMiningModelStatement(child ISQLVisitor, x *comment.SQLCommentOnMiningModelStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(OPERATOR)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}
func (v *SQLOutputVisitor) VisitCommentOnOperatorStatement(child ISQLVisitor, x *comment.SQLCommentOnOperatorStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(OPERATOR)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}
func (v *SQLOutputVisitor) VisitCommentOnTableStatement(child ISQLVisitor, x *comment.SQLCommentOnTableStatement) bool {
	v.WriteKeyword(COMMENT)
	v.WriteSpaceAfterKeyword(ON)
	v.WriteSpaceAfterKeyword(TABLE)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(IS)
	WriteSpaceAfterAccept(child, x.Comment())
	return false
}

func (v *SQLOutputVisitor) VisitAlterDatabaseStatement(child ISQLVisitor, x *database.SQLAlterDatabaseStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(DATABASE)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteAlterDatabaseStatementAction(child, x.Actions())
	return false
}
func (v *SQLOutputVisitor) WriteAlterDatabaseStatementAction(child ISQLVisitor, actions []expr.ISQLExpr) {
	if actions == nil || len(actions) == 0 {
		return
	}
	for _, action := range actions {
		WriteSpaceAfterAccept(child, action)
	}
}
func (v *SQLOutputVisitor) VisitCreateDatabaseStatement(child ISQLVisitor, x *database.SQLCreateDatabaseStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(DATABASE)

	if x.IfNotExists {
		v.WriteSpaceAfterKeyword(IF_NOT_EXISTS)
	}

	WriteSpaceAfterAccept(child, x.Name())

	return false
}

func (v *SQLOutputVisitor) VisitDropDatabaseStatement(child ISQLVisitor, x *database.SQLDropDatabaseStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(DATABASE)

	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}

	WriteSpaceAfterAccept(child, x.Name())

	return false
}

func (v *SQLOutputVisitor) VisitAlterFunctionStatement(child ISQLVisitor, x *statementFunction.SQLAlterFunctionStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitCreateFunctionStatement(child ISQLVisitor, x *statementFunction.SQLCreateFunctionStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitDropFunctionStatement(child ISQLVisitor, x *statementFunction.SQLDropFunctionStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(FUNCTION)
	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitAlterIndexStatement(child ISQLVisitor, x *statementIndex.SQLAlterIndexStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitCreateIndexStatement(child ISQLVisitor, x *statementIndex.SQLCreateIndexStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(INDEX)

	WriteSpaceAfterAccept(child, x.Name())

	v.WriteSpaceAfterKeyword(ON)

	if x.Cluster {
		v.WriteSpaceAfterKeyword(CLUSTER)
	}
	WriteSpaceAfterAccept(child, x.OnName())

	v.VisitCreateIndexStatementColumns(x.Columns())

	for i := 0; i < len(x.Options()); i++ {
		WriteSpaceAfterAccept(child, x.Option(i))
	}
	return false
}
func (v *SQLOutputVisitor) VisitCreateIndexStatementColumns(columns []*index.SQLIndexColumn) {
	if columns == nil || len(columns) == 0 {
		return
	}
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	v.IncrementIndentAndWriteLn()
	for i := 0; i < len(columns); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}
		Accept(v, columns[i])
	}
	v.DecrementIndentAndWriteLn()
	v.WriteKeyword(SYMB_RIGHT_PAREN)
}
func (v *SQLOutputVisitor) VisitDropIndexStatement(child ISQLVisitor, x *statementIndex.SQLDropIndexStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(INDEX)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterKeyword(ON)
	WriteSpaceAfterAccept(child, x.OnTable())

	for i := 0; i < len(x.Options()); i++ {
		WriteSpaceAfterAccept(child, x.Option(i))
	}
	return false
}

func (v *SQLOutputVisitor) VisitAlterPackageStatement(child ISQLVisitor, x *package_.SQLAlterPackageStatement) bool {
	return true
}
func (v *SQLOutputVisitor) VisitCreatePackageStatement(child ISQLVisitor, x *package_.SQLCreatePackageStatement) bool {
	return true
}
func (v *SQLOutputVisitor) VisitDropPackageStatement(child ISQLVisitor, x *package_.SQLDropPackageStatement) bool {
	return true
}

func (v *SQLOutputVisitor) VisitAlterPackageBodyStatement(child ISQLVisitor, x *packagebody.SQLAlterPackageBodyStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitCreatePackageBoydStatement(child ISQLVisitor, x *packagebody.SQLCreatePackageBoydStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(PACKAGE)
	v.WriteSpaceAfterKeyword(BODY)
	return false
}
func (v *SQLOutputVisitor) VisitDropPackageBodyStatement(child ISQLVisitor, x *packagebody.SQLDropPackageBodyStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(PACKAGE)
	v.WriteSpaceAfterKeyword(BODY)
	// if x.IfExists {
	// 	v.WriteSpaceAfterKeyword(IF_EXISTS)
	// }
	// WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitAlterProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLAlterProcedureStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(PROCEDURE)
	return false
}
func (v *SQLOutputVisitor) VisitCreateProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLCreateProcedureStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(PROCEDURE)
	return false
}
func (v *SQLOutputVisitor) VisitDropProcedureStatement(child ISQLVisitor, x *statementProcedure.SQLDropProcedureStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(PROCEDURE)
	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitAlterRoleStatement(child ISQLVisitor, x *statementRole.SQLAlterRoleStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(ROLE)
	return false
}
func (v *SQLOutputVisitor) VisitCreateRoleStatement(child ISQLVisitor, x *statementRole.SQLCreateRoleStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(ROLE)
	if x.IfNotExists {
		v.WriteSpaceAfterKeyword(IF_NOT_EXISTS)
	}

	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Name(i))
	}

	return false
}
func (v *SQLOutputVisitor) VisitDropRoleStatement(child ISQLVisitor, x *statementRole.SQLDropRoleStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(ROLE)
	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}
	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Name(i))
	}
	return false
}

func (v *SQLOutputVisitor) VisitAlterSchemaStatement(child ISQLVisitor, x *schema.SQLAlterSchemaStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitCreateSchemaStatement(child ISQLVisitor, x *schema.SQLCreateSchemaStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(SCHEMA)

	if x.IfNotExists {
		v.WriteSpaceAfterKeyword(IF_NOT_EXISTS)
	}

	WriteSpaceAfterAccept(child, x.Name())

	return false
}
func (v *SQLOutputVisitor) VisitDropSchemaStatement(child ISQLVisitor, x *schema.SQLDropSchemaStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(SCHEMA)

	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}

	WriteSpaceAfterAccept(child, x.Name())

	return false
}

func (v *SQLOutputVisitor) VisitAlterSequenceStatement(child ISQLVisitor, x *sequence.SQLAlterSequenceStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(SEQUENCE)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSequenceStatementOptions(child, x.Options())
	return false
}
func (v *SQLOutputVisitor) VisitCreateSequenceStatement(child ISQLVisitor, x *sequence.SQLCreateSequenceStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(SEQUENCE)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSequenceStatementOptions(child, x.Options())
	return false
}
func (v *SQLOutputVisitor) WriteSequenceStatementOptions(child ISQLVisitor, options []expr.ISQLExpr) {
	if options == nil || len(options) == 0 {
		return
	}
	v.IncrementIndent()
	for i := 0; i < len(options); i++ {
		WriteLnAfterAccept(child, options[i])
	}
	v.DecrementIndent()
}

func (v *SQLOutputVisitor) VisitDropSequenceStatement(child ISQLVisitor, x *sequence.SQLDropSequenceStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(SEQUENCE)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterValue(string(x.DropBehavior))
	return false
}

func (v *SQLOutputVisitor) VisitAlterServerStatement(child ISQLVisitor, x *statementServer.SQLAlterServerStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(SERVER)
	WriteSpaceAfterAccept(child, x.Name())

	v.WriteSpaceAfterKeyword(OPTIONS)
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i, option := range x.Options() {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, option)
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitCreateServerStatement(child ISQLVisitor, x *statementServer.SQLCreateServerStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(SERVER)
	WriteSpaceAfterAccept(child, x.Name())

	// FOREIGN DATA WRAPPER
	v.WriteLnAfterKeyword(FOREIGN)
	v.WriteSpaceAfterKeyword(DATA)
	v.WriteSpaceAfterKeyword(WRAPPER)
	WriteSpaceAfterAccept(child, x.WrapperName())

	v.WriteLnAfterKeyword(OPTIONS)
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	for i, option := range x.Options() {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteSpace()
		}
		Accept(v, option)
	}
	v.WriteKeyword(SYMB_RIGHT_PAREN)
	return false
}
func (v *SQLOutputVisitor) VisitDropServerStatement(child ISQLVisitor, x *statementServer.SQLDropServerStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(SERVER)
	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitAlterSynonymStatement(child ISQLVisitor, x *synonym.SQLAlterSynonymStatement) bool {
	v.WriteKeyword(ALTER)
	if x.Public {
		v.WriteSpaceAfterKeyword(PUBLIC)
	}
	v.WriteSpaceAfterKeyword(SYNONYM)
	WriteSpaceAfterAccept(child, x.Name())
	WriteSpaceAfterAccept(child, x.Action())
	return false
}
func (v *SQLOutputVisitor) VisitCreateSynonymStatement(child ISQLVisitor, x *synonym.SQLCreateSynonymStatement) bool {
	v.WriteKeyword(CREATE)
	if x.OrReplace {
		v.WriteSpaceAfterKeyword(OR)
		v.WriteSpaceAfterKeyword(REPLACE)
	}
	if x.Public {
		v.WriteSpaceAfterKeyword(PUBLIC)
	}
	v.WriteSpaceAfterKeyword(SYNONYM)
	WriteSpaceAfterAccept(child, x.Name())

	v.WriteSpaceAfterKeyword(FOR)
	WriteSpaceAfterAccept(child, x.ForName())
	return false
}
func (v *SQLOutputVisitor) VisitDropSynonymStatement(child ISQLVisitor, x *synonym.SQLDropSynonymStatement) bool {
	v.WriteKeyword(DROP)
	if x.Public {
		v.WriteSpaceAfterKeyword(PUBLIC)
	}
	v.WriteSpaceAfterKeyword(SYNONYM)
	WriteSpaceAfterAccept(child, x.Name())
	if x.Force {
		v.WriteSpaceAfterKeyword(FORCE)
	}
	return false
}

func (v *SQLOutputVisitor) VisitAlterTableStatement(child ISQLVisitor, x *statementTable.SQLAlterTableStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(TABLE)

	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}

	if x.Only {
		v.WriteSpaceAfterKeyword(ONLY)
	}

	WriteSpaceAfterAccept(child, x.Name())

	v.IncrementIndentAndWriteLn()
	for i := 0; i < len(x.Actions()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}
		Accept(child, x.Action(i))
	}
	v.DecrementIndent()
	return false
}
func (v *SQLOutputVisitor) VisitCreateTableStatement(child ISQLVisitor, x *statementTable.SQLCreateTableStatement) bool {
	v.WriteKeyword(CREATE)

	v.WriteSpaceAfterValue(string(x.TableScope))

	v.WriteSpaceAfterKeyword(TABLE)

	if x.IfNotExists {
		v.WriteSpaceAfterKeyword(IF_NOT_EXISTS)
	}

	WriteSpaceAfterAccept(child, x.Name())

	v.WriteTableElement(x.Paren, x.Elements())

	for i := 0; i < len(x.Options()); i++ {
		WriteSpaceAfterAccept(child, x.Option(i))
	}

	WriteLnAfterAccept(child, x.PartitionBy())

	if x.As {
		v.WriteSpaceAfterKeyword(AS)
		WriteLnAfterAccept(child, x.SubQuery())
	}

	return false
}

func (v *SQLOutputVisitor) WriteTableElement(paren bool, x []exprTable.ISQLTableElement) {
	if x == nil || len(x) == 0 {
		return
	}
	if paren {
		v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
		v.IncrementIndentAndWriteLn()
	} else {
		v.WriteSpace()
	}

	for i := 0; i < len(x); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}

		Accept(v, x[i])
	}

	if paren {
		v.DecrementIndentAndWriteLn()
		v.WriteKeyword(SYMB_RIGHT_PAREN)
	}

}

func (v *SQLOutputVisitor) VisitDropTableStatement(child ISQLVisitor, x *statementTable.SQLDropTableStatement) bool {
	v.WriteKeyword(DROP)

	if x.Temporary {
		v.WriteSpaceAfterKeyword(TEMPORARY)
	}

	v.WriteSpaceAfterKeyword(TABLE)

	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}

	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Names()[i])
	}

	return false
}

func (v *SQLOutputVisitor) VisitAlterTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLAlterTriggerStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitCreateTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLCreateTriggerStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(TRIGGER)

	return false
}
func (v *SQLOutputVisitor) VisitDropTriggerStatement(child ISQLVisitor, x *statementTrigger.SQLDropTriggerStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(TRIGGER)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitAlterTypeStatement(child ISQLVisitor, x *statementType.SQLAlterTypeStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitCreateTypeStatement(child ISQLVisitor, x *statementType.SQLCreateTypeStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitDropTypeStatement(child ISQLVisitor, x *statementType.SQLDropTypeStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(TYPE)
	WriteSpaceAfterAccept(child, x.Name())
	v.WriteSpaceAfterValue(string(x.Behavior))
	return false
}

func (v *SQLOutputVisitor) VisitAlterTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLAlterTypeBodyStatement) bool {
	return false
}
func (v *SQLOutputVisitor) VisitCreateTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLCreateTypeBodyStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(TYPE)
	v.WriteSpaceAfterKeyword(BODY)
	return false
}
func (v *SQLOutputVisitor) VisitDropTypeBodyStatement(child ISQLVisitor, x *statementTypeBody.SQLDropTypeBodyStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(TYPE)
	v.WriteSpaceAfterKeyword(BODY)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitAlterUserStatement(child ISQLVisitor, x *statementUser.SQLAlterUserStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(USER)
	return false
}

func (v *SQLOutputVisitor) VisitCreateUserStatement(child ISQLVisitor, x *statementUser.SQLCreateUserStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(USER)

	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Name(i))
	}
	return false
}
func (v *SQLOutputVisitor) VisitDropUserStatement(child ISQLVisitor, x *statementUser.SQLDropUserStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(USER)
	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}
	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Name(i))
	}
	return false
}

func (v *SQLOutputVisitor) VisitAlterViewStatement(child ISQLVisitor, x *statementView.SQLAlterViewStatement) bool {
	v.WriteKeyword(ALTER)
	v.WriteSpaceAfterKeyword(VIEW)

	WriteSpaceAfterAccept(child, x.Name())
	v.WriteAlterViewStatementElements(x.Elements())

	// MySQL
	if x.SubQuery() != nil {
		v.WriteLnAfterKeyword(AS)
		WriteLnAfterAccept(child, x.SubQuery())
	}

	// Oracle
	v.WriteAlterViewStatementActions(child, x.Actions())

	return false
}
func (v *SQLOutputVisitor) WriteAlterViewStatementElements(elements []exprView.ISQLViewElement) {
	if elements == nil || len(elements) == 0 {
		return
	}
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	v.IncrementIndentAndWriteLn()
	for i, element := range elements {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
			v.WriteLn()
		}
		Accept(v, element)
	}
	v.DecrementIndentAndWriteLn()
	v.WriteKeyword(SYMB_RIGHT_PAREN)
}
func (v *SQLOutputVisitor) WriteAlterViewStatementActions(child ISQLVisitor, actions []expr.ISQLExpr) {
	if actions == nil || len(actions) == 0 {
		return
	}
	for _, action := range actions {
		WriteSpaceAfterAccept(child, action)
	}
}
func (v *SQLOutputVisitor) VisitCreateViewStatement(child ISQLVisitor, x *statementView.SQLCreateViewStatement) bool {
	v.WriteKeyword(CREATE)
	v.WriteSpaceAfterKeyword(VIEW)

	WriteSpaceAfterAccept(child, x.Name())

	v.WriteViewElements(x.Elements())

	v.WriteLnAfterKeyword(AS)
	WriteLnAfterAccept(child, x.SubQuery())

	return false
}

func (v *SQLOutputVisitor) WriteViewElements(elements []exprView.ISQLViewElement) {
	if len(elements) == 0 {
		return
	}
	v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
	v.IncrementIndentAndWriteLn()
	for i := 0; i < len(elements); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		Accept(v, elements[i])
	}
	v.DecrementIndentAndWriteLn()
	v.WriteKeyword(SYMB_RIGHT_PAREN)
}

func (v *SQLOutputVisitor) VisitDropViewStatement(child ISQLVisitor, x *statementView.SQLDropViewStatement) bool {
	v.WriteKeyword(DROP)
	v.WriteSpaceAfterKeyword(VIEW)

	if x.IfExists {
		v.WriteSpaceAfterKeyword(IF_EXISTS)
	}

	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Name(i))
	}

	v.WriteSpaceAfterValue(string(x.Behavior))

	return false
}

func (v *SQLOutputVisitor) VisitDeleteStatement(child ISQLVisitor, x *statement.SQLDeleteStatement) bool {
	v.WriteKeyword(DELETE)

	if x.LowPriority {
		v.WriteSpaceAfterKeyword(LOW_PRIORITY)
	}
	if x.Quick {
		v.WriteSpaceAfterKeyword(QUICK)
	}
	if x.Ignore {
		v.WriteSpaceAfterKeyword(IGNORE)
	}

	for i := 0; i < len(x.Tables()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Table(i))
	}

	if x.From {
		v.WriteSpaceAfterKeyword(FROM)
	}
	WriteSpaceAfterAccept(child, x.TableReference())

	if x.UsingTableReference() != nil {
		v.WriteSpaceAfterKeyword(USING)
		WriteSpaceAfterAccept(child, x.UsingTableReference())
	}

	WriteLnAfterAccept(child, x.WhereClause())
	WriteLnAfterAccept(child, x.OrderByClause())
	WriteLnAfterAccept(child, x.LimitClause())

	WriteLnAfterAccept(child, x.ReturningClause())

	return false
}
func (v *SQLOutputVisitor) VisitInsertStatement(child ISQLVisitor, x *statement.SQLInsertStatement) bool {
	v.WriteKeyword(INSERT)

	if x.Into {
		v.WriteSpaceAfterKeyword(INTO)
	}

	WriteSpaceAfterAccept(child, x.TableReference())
	if len(x.Columns()) > 0 {
		v.WriteSpaceAfterKeyword(SYMB_LEFT_PAREN)
		for i := 0; i < len(x.Columns()); i++ {
			if i != 0 {
				v.WriteKeyword(SYMB_COMMA)
				v.WriteSpace()
			}
			Accept(v, x.Column(i))
		}
		v.WriteSpaceAfterKeyword(SYMB_RIGHT_PAREN)
	}

	if len(x.Values()) > 0 {
		v.WriteLnAfterKeyword(VALUES)

		for i := 0; i < len(x.Values()); i++ {
			if i != 0 {
				v.WriteKeyword(SYMB_COMMA)
			}
			WriteSpaceAfterAccept(child, x.Value(i))
		}
	}

	if x.SubQuery() != nil {
		WriteLnAfterAccept(child, x.SubQuery())
	}

	if len(x.UpdateAssignments()) > 0 {
		v.WriteSpaceAfterKeyword(ON)
		v.WriteSpaceAfterKeyword(DUPLICATE)
		v.WriteSpaceAfterKeyword(KEY)
		v.WriteSpaceAfterKeyword(UPDATE)

		for i := 0; i < len(x.UpdateAssignments()); i++ {
			if i != 0 {
				v.WriteKeyword(SYMB_COMMA)
			}
			WriteSpaceAfterAccept(child, x.UpdateAssignment(i))
		}
	}

	return false
}
func (v *SQLOutputVisitor) VisitSelectStatement(child ISQLVisitor, x *statement.SQLSelectStatement) bool {
	Accept(v, x.Query())
	return false
}
func (v *SQLOutputVisitor) VisitUpdateStatement(child ISQLVisitor, x *statement.SQLUpdateStatement) bool {
	v.WriteKeyword(UPDATE)

	if x.LowPriority {
		v.WriteSpaceAfterKeyword(LOW_PRIORITY)
	}
	if x.Ignore {
		v.WriteSpaceAfterKeyword(IGNORE)
	}

	WriteSpaceAfterAccept(child, x.TableReference())

	v.WriteLnAfterKeyword(SET)
	for i := 0; i < len(x.Assignments()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Assignment(i))
	}

	WriteLnAfterAccept(child, x.WhereClause())

	WriteLnAfterAccept(child, x.OrderByClause())

	WriteLnAfterAccept(child, x.LimitClause())

	return false
}

func (v *SQLOutputVisitor) VisitSetVariableAssignmentStatement(child ISQLVisitor, x *set.SQLSetVariableAssignmentStatement) bool {
	v.WriteKeyword(SET)
	for i := 0; i < len(x.Elements()); i++ {
		if i != 0 {
			v.WriteKeyword(SYMB_COMMA)
		}
		WriteSpaceAfterAccept(child, x.Element(i))
	}
	return false
}
func (v *SQLOutputVisitor) VisitSetCharacterSetStatement(child ISQLVisitor, x *set.SQLSetCharacterSetStatement) bool {
	v.WriteKeyword(SET)
	v.WriteSpaceAfterKeyword(CHARACTER)
	v.WriteSpaceAfterKeyword(SET)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitSetCharsetStatement(child ISQLVisitor, x *set.SQLSetCharsetStatement) bool {
	v.WriteKeyword(SET)
	v.WriteSpaceAfterKeyword(CHARSET)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitSetNamesStatement(child ISQLVisitor, x *set.SQLSetNamesStatement) bool {
	v.WriteKeyword(SET)
	v.WriteSpaceAfterKeyword(NAMES)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitShowCreateDatabaseStatement(child ISQLVisitor, x *show.SQLShowCreateDatabaseStatement) bool {
	v.WriteKeyword(SHOW)
	v.WriteSpaceAfterKeyword(CREATE)
	v.WriteSpaceAfterKeyword(DATABASE)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitShowCreateEventStatement(child ISQLVisitor, x *show.SQLShowCreateEventStatement) bool {
	v.WriteKeyword(SHOW)
	v.WriteSpaceAfterKeyword(CREATE)
	v.WriteSpaceAfterKeyword(EVENT)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitShowCreateFunctionStatement(child ISQLVisitor, x *show.SQLShowCreateFunctionStatement) bool {
	v.WriteKeyword(SHOW)
	v.WriteSpaceAfterKeyword(CREATE)
	v.WriteSpaceAfterKeyword(FUNCTION)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitShowCreateProcedureStatement(child ISQLVisitor, x *show.SQLShowCreateProcedureStatement) bool {
	v.WriteKeyword(SHOW)
	v.WriteSpaceAfterKeyword(CREATE)
	v.WriteSpaceAfterKeyword(PROCEDURE)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitShowCreateTableStatement(child ISQLVisitor, x *show.SQLShowCreateTableStatement) bool {
	v.WriteKeyword(SHOW)
	v.WriteSpaceAfterKeyword(CREATE)
	v.WriteSpaceAfterKeyword(TABLE)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitShowCreateTriggerStatement(child ISQLVisitor, x *show.SQLShowCreateTriggerStatement) bool {
	v.WriteKeyword(SHOW)
	v.WriteSpaceAfterKeyword(CREATE)
	v.WriteSpaceAfterKeyword(TRIGGER)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitShowCreateViewStatement(child ISQLVisitor, x *show.SQLShowCreateViewStatement) bool {
	v.WriteKeyword(SHOW)
	v.WriteSpaceAfterKeyword(CREATE)
	v.WriteSpaceAfterKeyword(VIEW)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

func (v *SQLOutputVisitor) VisitDescStatement(child ISQLVisitor, x *statement.SQLDescStatement) bool {
	v.WriteKeyword(DESC)
	if x.Analyze {
		v.WriteSpaceAfterKeyword(ANALYZE)
	}

	WriteSpaceAfterAccept(child, x.Table())
	WriteSpaceAfterAccept(child, x.Column())

	WriteSpaceAfterAccept(child, x.ExplainType())
	if x.ConnectionId() != nil {
		v.WriteSpaceAfterKeyword(FOR)
		v.WriteSpaceAfterKeyword(CONNECTION)
		WriteSpaceAfterAccept(child, x.ConnectionId())
	}
	WriteIndentLnAfterAccept(child, false, x.Stmt())

	return false
}
func (v *SQLOutputVisitor) VisitDescribeStatement(child ISQLVisitor, x *statement.SQLDescribeStatement) bool {
	v.WriteKeyword(DESCRIBE)
	if x.Analyze {
		v.WriteSpaceAfterKeyword(ANALYZE)
	}

	WriteSpaceAfterAccept(child, x.Table())
	WriteSpaceAfterAccept(child, x.Column())

	WriteSpaceAfterAccept(child, x.ExplainType())
	if x.ConnectionId() != nil {
		v.WriteSpaceAfterKeyword(FOR)
		v.WriteSpaceAfterKeyword(CONNECTION)
		WriteSpaceAfterAccept(child, x.ConnectionId())
	}
	WriteIndentLnAfterAccept(child, false, x.Stmt())
	return false
}
func (v *SQLOutputVisitor) VisitExplainStatement(child ISQLVisitor, x *statement.SQLExplainStatement) bool {
	v.WriteKeyword(EXPLAIN)
	if x.Analyze {
		v.WriteSpaceAfterKeyword(ANALYZE)
	}

	WriteSpaceAfterAccept(child, x.Table())
	WriteSpaceAfterAccept(child, x.Column())

	WriteSpaceAfterAccept(child, x.ExplainType())
	if x.ConnectionId() != nil {
		v.WriteSpaceAfterKeyword(FOR)
		v.WriteSpaceAfterKeyword(CONNECTION)
		WriteSpaceAfterAccept(child, x.ConnectionId())
	}
	WriteIndentLnAfterAccept(child, false, x.Stmt())
	return false
}

func (v *SQLOutputVisitor) VisitHelpStatement(child ISQLVisitor, x *statement.SQLHelpStatement) bool {
	v.WriteKeyword(HELP)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitUseStatement(child ISQLVisitor, x *statement.SQLUseStatement) bool {
	v.WriteKeyword(USE)
	WriteSpaceAfterAccept(child, x.Name())
	return false
}

// ---------------------------- Statement End ------------------------------------

func (v *SQLOutputVisitor) BeforeVisit(child ISQLVisitor, x ast.ISQLObject) {
	v.WriteBeforeComments(x.BeforeComments())
}

func (v *SQLOutputVisitor) WriteBeforeComments(comments []ast.ISQLComment) {
	if comments == nil || len(comments) == 0 {
		return
	}
	v.WriteComments(comments)
	comment := comments[len(comments)-1]
	switch comment.(type) {
	case *ast.SQLMinusComment, *ast.SQLSharpComment:
		v.WriteLn()
	case *ast.SQLMultiLineComment:
		v.WriteSpace()
	}
}

func (v *SQLOutputVisitor) WriteAfterComments(comments []ast.ISQLComment) {
	if comments == nil || len(comments) == 0 {
		return
	}
	v.WriteSpace()
	v.WriteComments(comments)
}

func (v *SQLOutputVisitor) WriteComments(comments []ast.ISQLComment) {
	if comments == nil || len(comments) == 0 {
		return
	}
	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		if i != 0 {
			v.WriteLn()
		}

		Accept(v, comment)
	}
}

func (v *SQLOutputVisitor) AfterVisit(child ISQLVisitor, x ast.ISQLObject) {
	if x.IsAfterSemi() {
		v.WriteKeyword(SYMB_SEMI)
	}
	v.WriteAfterComments(x.AfterComments())
}
