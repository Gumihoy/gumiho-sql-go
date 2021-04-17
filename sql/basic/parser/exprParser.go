package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/common"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/condition"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/function"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/literal"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/operator"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/sequence"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/variable"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
)

var complexFunctionNameMap = make(map[string]bool)
var nonParametricFunctionNameMap = make(map[string]bool)

func init() {

}

type ISQLExprParser interface {
	ISQLParser

	CreateSQLCommentOnAuditPolicyStatementParser() ISQLCommentStatementParser
	CreateSQLCommentOnColumnStatementParser() ISQLCommentStatementParser
	CreateSQLCommentOnEditionStatementParser() ISQLCommentStatementParser
	CreateSQLCommentOnIndextypeStatementParser() ISQLCommentStatementParser
	CreateSQLCommentOnMaterializedViewStatementParser() ISQLCommentStatementParser
	CreateSQLCommentOnMiningModelStatementParser() ISQLCommentStatementParser
	CreateSQLCommentOnOperatorStatementParser() ISQLCommentStatementParser
	CreateSQLCommentOnTableStatementParser() ISQLCommentStatementParser

	CreateSQLDatabaseStatementParser() ISQLDatabaseStatementParser
	CreateSQLFunctionStatementParser() ISQLFunctionStatementParser
	CreateSQLIndexStatementParser() ISQLIndexStatementParser
	CreateSQLPackageStatementParser() ISQLPackageStatementParser
	CreateSQLPackageBodyStatementParser() ISQLPackageBodyStatementParser
	CreateSQLProcedureStatementParser() ISQLProcedureStatementParser
	CreateSQLRoleStatementParser() ISQLRoleStatementParser
	CreateSQLSchemaStatementParser() ISQLSchemaStatementParser
	CreateSQLSequenceStatementParser() ISQLSequenceStatementParser
	CreateSQLServerStatementParser() ISQLServerStatementParser
	CreateSQLSynonymStatementParser() ISQLSynonymStatementParser
	CreateSQLTableStatementParser() ISQLTableStatementParser
	CreateSQLTriggerStatementParser() ISQLTriggerStatementParser
	CreateSQLTypeStatementParser() ISQLTypeStatementParser
	CreateSQLTypeBodyStatementParser() ISQLTypeBodyStatementParser
	CreateSQLUserStatementParser() ISQLUserStatementParser
	CreateSQLViewStatementParser() ISQLViewStatementParser

	CreateSQLDeleteStatementParser() ISQLDeleteStatementParser
	CreateSQLInsertStatementParser() ISQLInsertStatementParser
	CreateSQLSelectStatementParser() ISQLSelectStatementParser
	CreateSQLUpdateStatementParser() ISQLUpdateStatementParser

	CreateSQLSetVariableAssignmentStatementParser() ISQLSetStatementParser
	CreateSQLSetCharacterSetStatementParser() ISQLSetStatementParser
	CreateSQLSetCharsetStatementParser() ISQLSetStatementParser
	CreateSQLSetNamesStatementParser() ISQLSetStatementParser

	CreateSQLShowCreateDatabaseStatementParser() ISQLShowStatementParser
	CreateSQLShowCreateEventStatementParser() ISQLShowStatementParser
	CreateSQLShowCreateFunctionStatementParser() ISQLShowStatementParser
	CreateSQLShowCreateProcedureStatementParser() ISQLShowStatementParser
	CreateSQLShowCreateTableStatementParser() ISQLShowStatementParser
	CreateSQLShowCreateTriggerStatementParser() ISQLShowStatementParser
	CreateSQLShowCreateViewStatementParser() ISQLShowStatementParser

	CreateSQLShowDatabasesStatementParser() ISQLShowStatementParser

	CreateSQLDescStatementParser() ISQLExplainStatementParser
	CreateSQLDescribeStatementParser() ISQLExplainStatementParser
	CreateSQLExplainStatementParser() ISQLExplainStatementParser

	CreateSQLHelpStatementParser() ISQLHelpStatementParser
	CreateSQLUseStatementParser() ISQLUseStatementParser

	ParseComments(child ISQLExprParser) []ast.ISQLComment

	IsComplexFunction(name string) bool
	ParseComplexFunction(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr

	IsNonParametricFunction(name string) bool
	ParseNonParametricFunction(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr

	ParseIdentifier(child ISQLExprParser) expr.ISQLIdentifier
	ParseNameRest(child ISQLExprParser, owner expr.ISQLName) expr.ISQLName

	ParseExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr
	ParsePrimaryExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr

	ParseAtExpr(child ISQLExprParser) expr.ISQLExpr
	ParseOuterJoinExprRest(x expr.ISQLExpr) (expr.ISQLExpr, bool)
	ParseCallExpr(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr

	IsParseLikeCondition() bool
	ParseLikeConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr
	ParseRegexpConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr
	ParseBetweenConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr
	ParseInConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr

	ParseDateDataType(child ISQLExprParser) datatype.ISQLDataType
	ParseDateTimeDataType(child ISQLExprParser) datatype.ISQLDataType
	ParseTimeDataType(child ISQLExprParser) datatype.ISQLDataType
	ParseTimestampDataType(child ISQLExprParser) datatype.ISQLDataType

	// --------------------------------------------- DDL --------------------------------------------------------
	// ------------------- Statement ---------------
	// ------------ Delete
	ParseDeleteStatement(child ISQLExprParser) statement.ISQLStatement
	// ------------ Update
	ParseInsertStatement(child ISQLExprParser) statement.ISQLStatement
	// ------------ Update
	ParseUpdateStatement(child ISQLExprParser) statement.ISQLStatement

	// --------------------------------------------- DDL Expr --------------------------------------------------------
	// ------------ Element
	ParseParameterDeclaration(child ISQLExprParser) *statement.SQLParameterDeclaration

	// ------------ Sequence
	ParseSequenceOption(child ISQLExprParser) expr.ISQLExpr

	// ------------ Synonym
	ParseAlterSynonymAction(child ISQLExprParser) expr.ISQLExpr

	// ------------ OnTable
	ParseTableElement(child ISQLExprParser) table.ISQLTableElement
	ParseTableColumn(child ISQLExprParser) *table.SQLTableColumn
	IsParseTableColumnOption(child ISQLExprParser) bool
	ParseTableColumnOption(child ISQLExprParser) (expr.ISQLExpr, bool)
	ParseGeneratedClause(child ISQLExprParser) expr.ISQLExpr
	ParseIdentityColumnClauseOption(child ISQLExprParser) expr.ISQLExpr
	IsParseColumnConstraint() bool
	ParseColumnConstraint(child ISQLExprParser) (table.ISQLColumnConstraint, bool)
	ParseColumnConstraintRest(child ISQLExprParser, name expr.ISQLName) (table.ISQLColumnConstraint, bool)

	IsParseTableConstraint() bool
	ParseTableConstraint(child ISQLExprParser) table.ISQLTableConstraint
	ParseTableConstraintRest(child ISQLExprParser, name expr.ISQLName) (table.ISQLTableConstraint, bool)

	ParsePartitionBy(child ISQLExprParser) table.ISQLPartitionBy
	ParseSubPartitionBy(child ISQLExprParser) table.ISQLSubPartitionBy
	ParsePartitionDefinition(child ISQLExprParser) *table.SQLPartitionDefinition
	ParsePartitionDefinitionOption(child ISQLExprParser) (expr.ISQLExpr, bool)
	ParseSubPartitionDefinition(child ISQLExprParser) *table.SQLSubPartitionDefinition

	ParseAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction
	ParseAddAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction
	ParseAddColumnAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction
	ParseAlterAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction
	ParseDropAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction
	ParseChangeAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction
	ParseModifyAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction
	ParseRenameAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction

	// ------------ View
	ParseViewElement(child ISQLExprParser) view.ISQLViewElement
	ParseViewColumn(child ISQLExprParser) *view.SQLViewColumn
	ParseViewColumnOption(child ISQLExprParser) (expr.ISQLExpr, bool)

	// ------------ Select
	ParseWithClause(child ISQLExprParser) select_.ISQLWithClause
	ParseSubQueryFactoringClauseRest(child ISQLExprParser, name expr.ISQLExpr) select_.ISQLFactoringClause
	ParseSubAvFactoringClauseRest(child ISQLExprParser, name expr.ISQLExpr) select_.ISQLFactoringClause

	ParseTableReference(child ISQLExprParser) select_.ISQLTableReference
	ParsePartitionExtensionClause(child ISQLExprParser) expr.ISQLExpr

	ParseLimitOffsetClause(child ISQLExprParser) select_.ISQLLimitClause
	ParseOffsetFetchClause(child ISQLExprParser) select_.ISQLLimitClause

	ParseLockClause(child ISQLExprParser) select_.ISQLLockClause
	ParseForUpdate(child ISQLExprParser) select_.ISQLLockClause
	ParseLockInShareModeClause(child ISQLExprParser) select_.ISQLLockClause
}

type SQLExprParser struct {
	*SQLParser
}

func NewExprParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLExprParser {
	return NewExprParserByLexer(NewLexer(sql), dbType, config)
}

func NewExprParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLExprParser {
	x := new(SQLExprParser)
	x.SQLParser = NewParserByLexer(lexer, dbType, config)
	return x
}

func (x *SQLExprParser) ParseStatements() []ast.ISQLObject {
	return x.ParseStatementsWithParent(nil)
}

func (x *SQLExprParser) ParseStatementsWithParent(parent ast.ISQLObject) []ast.ISQLObject {
	var comments []ast.ISQLComment
	var stmts []ast.ISQLObject
	for {
		if x.Accept(EOF) {
			break
		}

		if x.AcceptAndNextToken(SYMB_SEMI) {
			if len(stmts) > 0 {
				lastStmt := stmts[len(stmts)-1]
				lastStmt.SetAfterSemi(true)
			}
			continue
		}

		if x.AcceptAndNextToken(SYMB_SLASH) {
			continue
		}

		comments = x.ParseComments(x)
		stmt := ParseStatement(x)
		if stmt == nil {
			break
		}

		if len(comments) != 0 {
			stmt.AddBeforeComments(comments)
			comments = ClearComments(comments)
		}

		stmt.SetParent(parent)
		stmts = append(stmts, stmt)
	}

	if len(comments) != 0 {
		if len(stmts) > 0 {
			lastStmt := stmts[len(stmts)-1]
			lastStmt.AddAfterComments(comments)
			comments = ClearComments(comments)
		}
	}

	return stmts
}

func ParseStatement(x ISQLExprParser) statement.ISQLStatement {
	// COMMENT
	if x.Accept(COMMENT) {
		return ParseCommentStatement(x)
	}

	// DDL
	if x.Accept(ALTER) {
		return ParseAlterStatement(x)
	}
	// if x.Accept(ANALYZE) {
	// 	return parseAnalyzeStatement()
	// }
	// if x.Accept(ASSOCIATE) {
	// 	panic("UnSupport .")
	// 	return nil
	// }
	// if x.Accept(AUDIT) {
	// 	panic("UnSupport .")
	// 	return nil
	// }
	// if x.Accept(COMMENT) {
	// 	return parseCommentStatement()
	// }
	if x.Accept(CREATE) {
		return ParseCreateStatement(x)
	}
	if x.Accept(DROP) {
		return ParseDropStatement(x)
	}
	// if x.Accept(RENAME)) {
	// 	return null
	// }

	// DML
	// if x.Accept(CALL)) {
	// 	return parseCallStatement()
	// }
	if x.Accept(DELETE) {
		return x.ParseDeleteStatement(x)
	}
	// if x.Accept(EXPLAIN)) {
	// 	return parseExplainStatement()
	// }
	if x.Accept(INSERT) {
		return x.ParseInsertStatement(x)
	}
	// if x.Accept(LOCK)) {
	// 	return parseLockStatement()
	// }
	// if x.Accept(WITH)) {
	// 	return parseWithStatement()
	// }
	if IsSelect(x) {
		return ParseSelectStatement(x)
	}
	// if x.Accept(LPAREN)) {
	// 	return parseLParenStatement()
	// }
	if x.Accept(UPDATE) {
		return x.ParseUpdateStatement(x)
	}

	// FCL
	// if x.Accept(CASE)) {
	// 	return parseCaseStatement()
	// }
	// if x.Accept(CLOSE)) {
	// 	return parseCloseStatement()
	// }
	// if x.Accept(CONTINUE)) {
	// 	return parseContinueStatement()
	// }
	// if x.Accept(EXIT)) {
	// 	return parseExitStatement()
	// }
	// if x.Accept(FETCH)) {
	// 	return parseFetchStatement()
	// }
	// if x.Accept(SQLToken.TokenKind.FOR)) {
	// 	return parseForLoopStatement()
	// }
	// if x.Accept(GOTO)) {
	// 	return parseGotoStatement()
	// }
	// if x.Accept(IF)) {
	// 	return parseIfStatement()
	// }
	// if x.Accept(LOOP)) {
	// 	return parseLoopStatement()
	// }
	// if x.Accept(NULL)) {
	// 	return parseNullStatement()
	// }
	// if x.Accept(OPEN)) {
	// 	return parseOpenStatement()
	// }
	// if x.Accept(PIPE)) {
	// 	return parsePipeRowStatement()
	// }
	// if x.Accept(RAISE)) {
	// 	return parseRaiseStatement()
	// }
	// if x.Accept(REPEAT)) {
	// 	return parseRepeatStatement()
	// }
	// if x.Accept(RETURN)) {
	// 	return parseReturnStatement()
	// }
	// if x.Accept(WHILE)) {
	// 	return parseWhileLoopStatement()
	// }

	// TC
	// if x.Accept(COMMIT) {
	// 	return parseCommitStatement()
	// }
	// if x.Accept(ROLLBACK) {
	// 	return parseRollbackStatement()
	// }
	// if x.Accept(SAVEPOINT) {
	// 	return null
	// }

	// Database Administration Statements
	if x.Accept(SET) {
		return ParseSetStatement(x)
	}

	if x.Accept(SHOW) {
		return ParseShowStatement(x)
	}

	// Utility Statements
	if x.Accept(DESC) {
		return ParseDescStatement(x)
	}
	if x.Accept(DESCRIBE) {
		return ParseDescribeStatement(x)
	}
	if x.Accept(EXPLAIN) {
		return ParseExplainStatement(x)
	}

	if x.Accept(HELP) {
		return ParseHelpStatement(x)
	}
	if x.Accept(USE) {
		return ParseUseStatement(x)
	}

	// SC

	// make := x.make()
	// expr := x.ParseExpr()
	// if x.Accept(SQLToken.TokenKind.ASSIGN) {
	// 	// this.reset(make)
	// 	return parseAssignmentStatement()
	// }
	//
	//
	// if x.Accept(LOOP) {
	// 	// x.lexer.Reset()
	// 	return parseLoopStatement()
	//
	// }
	// if x.Accept(WHILE) {
	// 	// this.reset(make)
	// 	return parseWhileStatement()
	// }
	// this.reset(make)

	return nil
}

func ParseMinusComment(x ISQLExprParser) *ast.SQLMinusComment {
	x.Lexer().ScanMinusCommentRest()
	x.ClearComments()
	comment := ast.NewMinusCommentWithComment(x.StringValue())
	NextTokenByParser(x)
	return comment
}
func ParseMultiLineComment(x ISQLExprParser) *ast.SQLMultiLineComment {
	x.Lexer().ScanMultiLineCommentRest()
	x.ClearComments()
	comment := ast.NewMultiLineCommentWithComment(x.StringValue())
	NextTokenByParser(x)
	return comment
}
func ParseSharpComment(x ISQLExprParser) *ast.SQLSharpComment {
	x.Lexer().ScanSharpCommentRest()
	x.ClearComments()
	comment := ast.NewSharpCommentWithComment(x.StringValue())
	NextTokenByParser(x)
	return comment
}

func ParseCommentStatement(x ISQLExprParser) statement.ISQLStatement {
	mark := x.Mark()
	x.AcceptAndNextTokenWithError(COMMENT, true)

	x.AcceptAndNextTokenWithError(ON, true)

	// COLUMN
	if x.Accept(COLUMN) {
		x.ResetWithMark(mark)
		return ParseCommentOnColumnStatement(x)
	}

	// MATERIALIZED VIEW
	if x.AcceptAndNextToken(MATERIALIZED) {
		x.AcceptAndNextTokenWithError(VIEW, true)
		x.ResetWithMark(mark)
		return ParseCommentOnMaterializedViewStatement(x)
	}

	// TABLE
	if x.Accept(TABLE) {
		x.ResetWithMark(mark)
		return ParseCommentOnTableStatement(x)
	}

	panic(x.UnSupport())
}

func ParseCommentOnColumnStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLCommentOnColumnStatementParser()
	return parser.Parse()
}
func ParseCommentOnMaterializedViewStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLCommentOnMaterializedViewStatementParser()
	return parser.Parse()
}
func ParseCommentOnTableStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLCommentOnTableStatementParser()
	return parser.Parse()
}

func ParseAlterStatement(x ISQLExprParser) statement.ISQLStatement {
	mark := x.Mark()
	x.AcceptAndNextTokenWithError(ALTER, true)

	// Synonym
	x.AcceptAndNextToken(PUBLIC)

	// DATABASE
	if x.Accept(DATABASE) {
		x.ResetWithMark(mark)
		return ParseAlterDatabaseStatement(x)
	}

	// FUNCTION
	if x.Accept(FUNCTION) {
		x.ResetWithMark(mark)
		return ParseAlterFunctionStatement(x)
	}

	// INDEX
	if x.Accept(INDEX) {
		x.ResetWithMark(mark)
		return ParseAlterIndexStatement(x)
	}

	// PACKAGE
	if x.AcceptAndNextToken(PACKAGE) {
		if x.Accept(BODY) {
			x.ResetWithMark(mark)
			return ParseAlterPackageBodyStatement(x)
		}
		x.ResetWithMark(mark)
		return ParseAlterPackageStatement(x)
	}

	// PROCEDURE
	if x.Accept(PROCEDURE) {
		x.ResetWithMark(mark)
		return ParseAlterProcedureStatement(x)
	}

	// ROLE
	if x.Accept(ROLE) {
		x.ResetWithMark(mark)
		return ParseAlterRoleStatement(x)
	}

	// SCHEMA
	if x.Accept(SCHEMA) {
		x.ResetWithMark(mark)
		return ParseAlterSchemaStatement(x)
	}

	// SEQUENCE
	if x.Accept(SEQUENCE) {
		x.ResetWithMark(mark)
		return ParseAlterSequenceStatement(x)
	}

	// SERVER
	if x.Accept(SERVER) {
		x.ResetWithMark(mark)
		return ParseAlterServerStatement(x)
	}

	// SYNONYM
	if x.Accept(SYNONYM) {
		x.ResetWithMark(mark)
		return ParseAlterSynonymStatement(x)
	}

	// TABLE
	if x.Accept(TABLE) {
		x.ResetWithMark(mark)
		return ParseAlterTableStatement(x)
	}

	// TRIGGER
	if x.Accept(TRIGGER) {
		x.ResetWithMark(mark)
		return ParseAlterTriggerStatement(x)
	}

	// TYPE / TYPE BODY
	if x.AcceptAndNextToken(TYPE) {
		if x.Accept(BODY) {
			x.ResetWithMark(mark)
			return ParseAlterTypeBodyStatement(x)
		}

		x.ResetWithMark(mark)
		return ParseAlterTypeStatement(x)
	}

	// USER
	if x.Accept(USER) {
		x.ResetWithMark(mark)
		return ParseAlterUserStatement(x)
	}

	// USER
	if x.Accept(USER) {
		x.ResetWithMark(mark)
		return ParseAlterUserStatement(x)
	}

	// VIEW
	if x.Accept(VIEW) {
		x.ResetWithMark(mark)
		return ParseAlterViewStatement(x)
	}

	panic(x.UnSupport())
}

func ParseCreateStatement(x ISQLExprParser) statement.ISQLStatement {
	mark := x.Mark()

	x.AcceptAndNextTokenWithError(CREATE, true)

	// Database Link
	// x.AcceptAndNextToken(SHARED)

	// Database Link
	x.AcceptAndNextToken(PUBLIC)

	// Function
	// Type
	// Type Body
	ParseOrReplace(x)

	// OnTable
	// x.AcceptAndNextToken(GLOBAL)
	// x.AcceptAndNextToken(PRIVATE)
	x.AcceptAndNextToken(TEMPORARY)
	// x.AcceptAndNextToken(DUPLICATED)

	// DATABASE
	if x.Accept(DATABASE) {
		x.ResetWithMark(mark)
		return ParseCreateDatabaseStatement(x)
	}

	// FUNCTION
	if x.Accept(FUNCTION) {
		x.ResetWithMark(mark)
		return ParseCreateFunctionStatement(x)
	}

	// INDEX
	if x.Accept(INDEX) {
		x.ResetWithMark(mark)
		return ParseCreateIndexStatement(x)
	}

	// PACKAGE
	// PACKAGE BODY
	if x.AcceptAndNextToken(PACKAGE) {
		if x.Accept(BODY) {
			x.ResetWithMark(mark)
			return ParseCreatePackageBodyStatement(x)
		}
		x.ResetWithMark(mark)
		return ParseCreatePackageStatement(x)
	}

	// PROCEDURE
	if x.Accept(PROCEDURE) {
		x.ResetWithMark(mark)
		return ParseCreateProcedureStatement(x)
	}

	// ROLE
	if x.Accept(ROLE) {
		x.ResetWithMark(mark)
		return ParseCreateRoleStatement(x)
	}

	// SCHEMA
	if x.Accept(SCHEMA) {
		x.ResetWithMark(mark)
		return ParseCreateSchemaStatement(x)
	}

	// SEQUENCE
	if x.Accept(SEQUENCE) {
		x.ResetWithMark(mark)
		return ParseCreateSequenceStatement(x)
	}

	// SERVER
	if x.Accept(SERVER) {
		x.ResetWithMark(mark)
		return ParseCreateServerStatement(x)
	}

	// SYNONYM
	if x.Accept(SYNONYM) {
		x.ResetWithMark(mark)
		return ParseCreateSynonymStatement(x)
	}

	// TABLE
	if x.Accept(TABLE) {
		x.ResetWithMark(mark)
		return ParseCreateTableStatement(x)
	}

	// TRIGGER
	if x.Accept(TRIGGER) {
		x.ResetWithMark(mark)
		return ParseCreateTriggerStatement(x)
	}

	// TYPE
	// TYPE BODY
	if x.AcceptAndNextToken(TYPE) {
		if x.Accept(BODY) {
			x.ResetWithMark(mark)
			return ParseCreateTypeBodyStatement(x)
		}
		x.ResetWithMark(mark)
		return ParseCreateTypeStatement(x)
	}

	// USER
	if x.Accept(USER) {
		x.ResetWithMark(mark)
		return ParseCreateUserStatement(x)
	}

	// VIEW
	if x.Accept(VIEW) {
		x.ResetWithMark(mark)
		return ParseCreateViewStatement(x)
	}
	panic(x.UnSupport())
}

func ParseDropStatement(x ISQLExprParser) statement.ISQLStatement {
	mark := x.Mark()

	x.AcceptAndNextTokenWithError(DROP, true)

	// Database Link
	// x.AcceptAndNextToken(SHARED)

	// Database Link
	// Synonym
	x.AcceptAndNextToken(PUBLIC)

	// OnTable
	// x.AcceptAndNextToken(GLOBAL)
	// x.AcceptAndNextToken(PRIVATE)
	x.AcceptAndNextToken(TEMPORARY)
	// x.AcceptAndNextToken(DUPLICATED)

	// DATABASE
	if x.Accept(DATABASE) {
		x.ResetWithMark(mark)
		return ParseDropDatabaseStatement(x)
	}

	// FUNCTION
	if x.Accept(FUNCTION) {
		x.ResetWithMark(mark)
		return ParseDropFunctionStatement(x)
	}

	// INDEX
	if x.Accept(INDEX) {
		x.ResetWithMark(mark)
		return ParseDropIndexStatement(x)
	}

	// PACKAGE
	// PACKAGE BODY
	if x.AcceptAndNextToken(PACKAGE) {
		if x.Accept(BODY) {
			x.ResetWithMark(mark)
			return ParseDropPackageBodyStatement(x)
		}
		x.ResetWithMark(mark)
		return ParseDropPackageStatement(x)
	}

	// PROCEDURE
	if x.Accept(PROCEDURE) {
		x.ResetWithMark(mark)
		return ParseDropProcedureStatement(x)
	}

	// ROLE
	if x.Accept(ROLE) {
		x.ResetWithMark(mark)
		return ParseDropRoleStatement(x)
	}

	// SCHEMA
	if x.Accept(SCHEMA) {
		x.ResetWithMark(mark)
		return ParseDropSchemaStatement(x)
	}

	// SEQUENCE
	if x.Accept(SEQUENCE) {
		x.ResetWithMark(mark)
		return ParseDropSequenceStatement(x)
	}

	// SERVER
	if x.Accept(SERVER) {
		x.ResetWithMark(mark)
		return ParseDropServerStatement(x)
	}

	// SYNONYM
	if x.Accept(SYNONYM) {
		x.ResetWithMark(mark)
		return ParseDropSynonymStatement(x)
	}

	// TABLE
	if x.Accept(TABLE) {
		x.ResetWithMark(mark)
		return ParseDropTableStatement(x)
	}

	// TRIGGER
	if x.Accept(TRIGGER) {
		x.ResetWithMark(mark)
		return ParseDropTriggerStatement(x)
	}

	// TYPE
	// TYPE BODY
	if x.AcceptAndNextToken(TYPE) {
		if x.Accept(BODY) {
			x.ResetWithMark(mark)
			return ParseDropTypeBodyStatement(x)
		}
		x.ResetWithMark(mark)
		return ParseDropTypeStatement(x)
	}

	// USER
	if x.Accept(USER) {
		x.ResetWithMark(mark)
		return ParseDropUserStatement(x)
	}

	// VIEW
	if x.Accept(VIEW) {
		x.ResetWithMark(mark)
		return ParseDropViewStatement(x)
	}

	panic(x.UnSupport())
}

func ParseAlterDatabaseStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLDatabaseStatementParser()
	return parser.ParseAlter()
}
func ParseCreateDatabaseStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLDatabaseStatementParser()
	return parser.ParseCreate()
}
func ParseDropDatabaseStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLDatabaseStatementParser()
	return parser.ParseDrop()
}

func ParseAlterFunctionStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLFunctionStatementParser()
	return parser.ParseAlter()
}
func ParseCreateFunctionStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLFunctionStatementParser()
	return parser.ParseCreate()
}
func ParseDropFunctionStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLFunctionStatementParser()
	return parser.ParseDrop()
}

func ParseAlterIndexStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLIndexStatementParser()
	return parser.ParseAlter()
}
func ParseCreateIndexStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLIndexStatementParser()
	return parser.ParseCreate()
}
func ParseDropIndexStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLIndexStatementParser()
	return parser.ParseDrop()
}

func ParseAlterPackageStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLPackageStatementParser()
	return parser.ParseAlter()
}
func ParseCreatePackageStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLPackageStatementParser()
	return parser.ParseCreate()
}
func ParseDropPackageStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLPackageStatementParser()
	return parser.ParseDrop()
}

func ParseAlterPackageBodyStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLPackageBodyStatementParser()
	return parser.ParseAlter()
}
func ParseCreatePackageBodyStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLPackageBodyStatementParser()
	return parser.ParseCreate()
}
func ParseDropPackageBodyStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLPackageBodyStatementParser()
	return parser.ParseDrop()
}

func ParseAlterProcedureStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLProcedureStatementParser()
	return parser.ParseAlter()
}
func ParseCreateProcedureStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLProcedureStatementParser()
	return parser.ParseCreate()
}
func ParseDropProcedureStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLProcedureStatementParser()
	return parser.ParseDrop()
}

func ParseAlterRoleStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLRoleStatementParser()
	return parser.ParseAlter()
}
func ParseCreateRoleStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLRoleStatementParser()
	return parser.ParseCreate()
}
func ParseDropRoleStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLRoleStatementParser()
	return parser.ParseDrop()
}

func ParseAlterSchemaStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSchemaStatementParser()
	return parser.ParseAlter()
}
func ParseCreateSchemaStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSchemaStatementParser()
	return parser.ParseCreate()
}
func ParseDropSchemaStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSchemaStatementParser()
	return parser.ParseDrop()
}

func ParseAlterSequenceStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSequenceStatementParser()
	return parser.ParseAlter()
}
func ParseCreateSequenceStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSequenceStatementParser()
	return parser.ParseCreate()
}
func ParseDropSequenceStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSequenceStatementParser()
	return parser.ParseDrop()
}

func ParseAlterServerStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLServerStatementParser()
	return parser.ParseAlter()
}
func ParseCreateServerStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLServerStatementParser()
	return parser.ParseCreate()
}
func ParseDropServerStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLServerStatementParser()
	return parser.ParseDrop()
}

func ParseAlterSynonymStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSynonymStatementParser()
	return parser.ParseAlter()
}
func ParseCreateSynonymStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSynonymStatementParser()
	return parser.ParseCreate()
}
func ParseDropSynonymStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSynonymStatementParser()
	return parser.ParseDrop()
}

func ParseAlterTableStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTableStatementParser()
	return parser.ParseAlter()
}
func ParseCreateTableStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTableStatementParser()
	return parser.ParseCreate()
}
func ParseDropTableStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTableStatementParser()
	return parser.ParseDrop()
}

func ParseAlterTriggerStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTriggerStatementParser()
	return parser.ParseAlter()
}
func ParseCreateTriggerStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTriggerStatementParser()
	return parser.ParseCreate()
}
func ParseDropTriggerStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTriggerStatementParser()
	return parser.ParseDrop()
}

func ParseAlterTypeStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTypeStatementParser()
	return parser.ParseAlter()
}
func ParseCreateTypeStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTypeStatementParser()
	return parser.ParseCreate()
}
func ParseDropTypeStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTypeStatementParser()
	return parser.ParseDrop()
}
func ParseAlterTypeBodyStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTypeBodyStatementParser()
	return parser.ParseAlter()
}
func ParseCreateTypeBodyStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTypeBodyStatementParser()
	return parser.ParseCreate()
}
func ParseDropTypeBodyStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLTypeBodyStatementParser()
	return parser.ParseDrop()
}

func ParseAlterUserStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLUserStatementParser()
	return parser.ParseAlter()
}
func ParseCreateUserStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLUserStatementParser()
	return parser.ParseCreate()
}
func ParseDropUserStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLUserStatementParser()
	return parser.ParseDrop()
}

func ParseAlterViewStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLViewStatementParser()
	return parser.ParseAlter()
}
func ParseCreateViewStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLViewStatementParser()
	return parser.ParseCreate()
}
func ParseDropViewStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLViewStatementParser()
	return parser.ParseDrop()
}

/**
 * DELETE FROM <target table> WHERE CURRENT OF <cursor name>
 * DELETE FROM <target table> [ WHERE <search condition> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#delete%20statement:%20positioned
 */
func (sp *SQLExprParser) ParseDeleteStatement(child ISQLExprParser) statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DELETE) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(FROM, true)

	x := statement.NewDeleteStatement(sp.DBType())

	tableReference := ParseTableReference(child)
	x.SetTableReference(tableReference)

	whereClause := ParseWhereClause(child)
	x.SetWhereClause(whereClause)

	return x
}

/**
 * INSERT INTO ( <table name> | ONLY <left paren> <table name> <right paren> ) ( <from subquery> | <from constructor> |  <from default> )
 *
 * <from subquery>: [ <left paren> <insert column list> <right paren> ] [ <override clause> ] <query expression>
 * <from constructor>:  [ <left paren> <insert column list> <right paren> ] [ <override clause> ] <contextually typed table value constructor>
 * <from default>: DEFAULT VALUES
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#insert%20statement
 */
func (sp *SQLExprParser) ParseInsertStatement(child ISQLExprParser) statement.ISQLStatement {
	if !sp.AcceptAndNextToken(INSERT) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(INTO, true)

	x := statement.NewInsertStatement(sp.DBType())

	tableReference := sp.ParseTableReference(child)
	x.SetTableReference(tableReference)

	if sp.AcceptAndNextToken(SYMB_LEFT_PAREN) {

		for {
			column := child.ParseIdentifier(child)
			x.AddColumn(column)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	} else if sp.AcceptAndNextToken(DEFAULT) {
		sp.AcceptAndNextTokenWithError(VALUES, true)
	}

	if IsSelect(sp) {
		subQuery := ParseSelectQuery(child)
		x.SetSubQuery(subQuery)
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	} else {
		sp.AcceptAndNextTokenWithError(VALUES, true)

		for {
			value := ParseExpr(child)
			x.AddValue(value)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

	}

	return x
}

func IsSelect(sp ISQLExprParser) bool {
	return sp.Accept(WITH) || sp.Accept(SELECT)
}

func ParseSelectStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSelectStatementParser()
	return parser.Parse()
}

/**
 * UPDATE <target table> SET <set clause list> WHERE CURRENT OF <cursor name>
 * UPDATE <target table> SET <set clause list> [ WHERE <search condition> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#SQL%20data%20change%20statement
 */
func (sp *SQLExprParser) ParseUpdateStatement(child ISQLExprParser) statement.ISQLStatement {
	if !sp.AcceptAndNextToken(UPDATE) {
		return nil
	}
	x := statement.NewUpdateStatement(sp.DBType())

	tableReference := ParseTableReference(child)
	x.SetTableReference(tableReference)

	sp.AcceptAndNextTokenWithError(SET, true)
	for {
		assignment := ParseExpr(child)
		x.AddAssignment(assignment)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}

	whereClause := ParseWhereClause(sp)
	x.SetWhereClause(whereClause)

	return x
}
func ParseSetStatement(sp ISQLExprParser) statement.ISQLStatement {
	mark := sp.Mark()
	sp.AcceptAndNextTokenWithError(SET, true)

	if sp.Accept(CHARACTER) {

		sp.ResetWithMark(mark)
		return ParseSetCharacterSetStatement(sp)

	} else if sp.Accept(CHARSET) {
		sp.ResetWithMark(mark)

		return ParseSetCharsetStatement(sp)

	} else if sp.Accept(NAMES) {

		sp.ResetWithMark(mark)
		return ParseSetNamesStatement(sp)

	} else {
		sp.ResetWithMark(mark)
		return ParseSetVariableAssignmentStatement(sp)
	}
}

func ParseSetVariableAssignmentStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLSetVariableAssignmentStatementParser()
	return parser.Parse()
}
func ParseSetCharacterSetStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLSetCharacterSetStatementParser()
	return parser.Parse()
}
func ParseSetCharsetStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLSetCharsetStatementParser()
	return parser.Parse()
}
func ParseSetNamesStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLSetNamesStatementParser()
	return parser.Parse()
}

func ParseShowStatement(sp ISQLExprParser) statement.ISQLStatement {
	mark := sp.Mark()

	sp.AcceptAndNextTokenWithError(SHOW, true)

	// CREATE
	if sp.AcceptAndNextToken(CREATE) {

		if sp.Accept(DATABASE) {

			sp.ResetWithMark(mark)
			return ParseShowCreateDatabaseStatement(sp)

		} else if sp.Accept(EVENT) {

			sp.ResetWithMark(mark)
			return ParseShowCreateEventStatement(sp)

		} else if sp.Accept(FUNCTION) {

			sp.ResetWithMark(mark)
			return ParseShowCreateFunctionStatement(sp)

		} else if sp.Accept(PROCEDURE) {

			sp.ResetWithMark(mark)
			return ParseShowCreateProcedureStatement(sp)

		} else if sp.Accept(TABLE) {

			sp.ResetWithMark(mark)
			return ParseShowCreateTableStatement(sp)

		} else if sp.Accept(TRIGGER) {

			sp.ResetWithMark(mark)
			return ParseShowCreateTriggerStatement(sp)

		} else if sp.Accept(VIEW) {
			sp.ResetWithMark(mark)
			return ParseShowCreateViewStatement(sp)

		}

	} else if sp.Accept(DATABASES) {

		sp.ResetWithMark(mark)
		return ParseShowDatabasesStatement(sp)

	}

	panic(sp.UnSupport())
}

func ParseShowCreateDatabaseStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLShowCreateDatabaseStatementParser()
	return parser.Parse()
}
func ParseShowCreateEventStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLShowCreateEventStatementParser()
	return parser.Parse()
}
func ParseShowCreateFunctionStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLShowCreateFunctionStatementParser()
	return parser.Parse()
}
func ParseShowCreateProcedureStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLShowCreateProcedureStatementParser()
	return parser.Parse()
}
func ParseShowCreateTableStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLShowCreateTableStatementParser()
	return parser.Parse()
}
func ParseShowCreateTriggerStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLShowCreateTriggerStatementParser()
	return parser.Parse()
}
func ParseShowCreateViewStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLTableStatementParser()
	return parser.ParseCreate()
}
func ParseShowDatabasesStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLShowDatabasesStatementParser()
	return parser.Parse()
}

func ParseDescStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLDescStatementParser()
	return parser.Parse()
}
func ParseDescribeStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLDescribeStatementParser()
	return parser.Parse()
}
func ParseExplainStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLExplainStatementParser()
	return parser.Parse()
}

func ParseHelpStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLHelpStatementParser()
	return parser.Parse()
}
func ParseUseStatement(sp ISQLExprParser) statement.ISQLStatement {
	parser := sp.CreateSQLUseStatementParser()
	return parser.Parse()
}

func (x *SQLExprParser) CreateSQLCommentOnAuditPolicyStatementParser() ISQLCommentStatementParser {
	return nil
}
func (x *SQLExprParser) CreateSQLCommentOnColumnStatementParser() ISQLCommentStatementParser {
	return NewCommentOnColumnStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLCommentOnEditionStatementParser() ISQLCommentStatementParser {
	return nil
}
func (x *SQLExprParser) CreateSQLCommentOnIndextypeStatementParser() ISQLCommentStatementParser {
	return nil
}
func (x *SQLExprParser) CreateSQLCommentOnMaterializedViewStatementParser() ISQLCommentStatementParser {
	return NewCommentOnMaterializedViewStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLCommentOnMiningModelStatementParser() ISQLCommentStatementParser {
	return nil
}
func (x *SQLExprParser) CreateSQLCommentOnOperatorStatementParser() ISQLCommentStatementParser {
	return nil
}
func (x *SQLExprParser) CreateSQLCommentOnTableStatementParser() ISQLCommentStatementParser {
	return NewCommentOnTableStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLDatabaseStatementParser() ISQLDatabaseStatementParser {
	return NewDatabaseStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLFunctionStatementParser() ISQLFunctionStatementParser {
	return NewFunctionStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLIndexStatementParser() ISQLIndexStatementParser {
	return NewIndexStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLPackageStatementParser() ISQLPackageStatementParser {
	return NewPackageStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLPackageBodyStatementParser() ISQLPackageBodyStatementParser {
	return NewPackageBodyStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLProcedureStatementParser() ISQLProcedureStatementParser {
	return NewProcedureStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLRoleStatementParser() ISQLRoleStatementParser {
	return NewRoleStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLSchemaStatementParser() ISQLSchemaStatementParser {
	return NewSchemaStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLSequenceStatementParser() ISQLSequenceStatementParser {
	return NewSequenceStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLServerStatementParser() ISQLServerStatementParser {
	return NewServerStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLSynonymStatementParser() ISQLSynonymStatementParser {
	return NewSynonymStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLTableStatementParser() ISQLTableStatementParser {
	return NewTableStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLTriggerStatementParser() ISQLTriggerStatementParser {
	return NewTriggerStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLTypeStatementParser() ISQLTypeStatementParser {
	return NewTypeStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLTypeBodyStatementParser() ISQLTypeBodyStatementParser {
	return NewTypeBodyStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLUserStatementParser() ISQLUserStatementParser {
	return NewUserStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLViewStatementParser() ISQLViewStatementParser {
	return NewViewStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLDeleteStatementParser() ISQLDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLInsertStatementParser() ISQLInsertStatementParser {
	return NewInsertStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLSelectStatementParser() ISQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLUpdateStatementParser() ISQLUpdateStatementParser {
	return NewUpdateStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLSetVariableAssignmentStatementParser() ISQLSetStatementParser {
	return NewSetVariableAssignmentStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLSetCharacterSetStatementParser() ISQLSetStatementParser {
	return NewSetCharacterSetStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLSetCharsetStatementParser() ISQLSetStatementParser {
	return NewSetCharsetStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLSetNamesStatementParser() ISQLSetStatementParser {
	return NewSetNamesStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLShowCreateDatabaseStatementParser() ISQLShowStatementParser {
	return NewShowCreateDatabaseByExprParser(x)
}
func (x *SQLExprParser) CreateSQLShowCreateEventStatementParser() ISQLShowStatementParser {
	return NewShowCreateEventByExprParser(x)
}
func (x *SQLExprParser) CreateSQLShowCreateFunctionStatementParser() ISQLShowStatementParser {
	return NewShowCreateFunctionByExprParser(x)
}
func (x *SQLExprParser) CreateSQLShowCreateProcedureStatementParser() ISQLShowStatementParser {
	return NewShowCreateProcedureByExprParser(x)
}
func (x *SQLExprParser) CreateSQLShowCreateTableStatementParser() ISQLShowStatementParser {
	return NewShowCreateTableByExprParser(x)
}
func (x *SQLExprParser) CreateSQLShowCreateTriggerStatementParser() ISQLShowStatementParser {
	return NewShowCreateTriggerByExprParser(x)
}
func (x *SQLExprParser) CreateSQLShowCreateViewStatementParser() ISQLShowStatementParser {
	return NewShowCreateViewByExprParser(x)
}
func (x *SQLExprParser) CreateSQLShowDatabasesStatementParser() ISQLShowStatementParser {
	return NewShowDatabasesByExprParser(x)
}

func (x *SQLExprParser) CreateSQLDescStatementParser() ISQLExplainStatementParser {
	return NewDescStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLDescribeStatementParser() ISQLExplainStatementParser {
	return NewDescribeStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLExplainStatementParser() ISQLExplainStatementParser {
	return NewExplainStatementParserByExprParser(x)
}

func (x *SQLExprParser) CreateSQLHelpStatementParser() ISQLHelpStatementParser {
	return NewHelpStatementParserByExprParser(x)
}
func (x *SQLExprParser) CreateSQLUseStatementParser() ISQLUseStatementParser {
	return NewUseStatementParserByExprParser(x)
}

func (sp *SQLExprParser) ParseComments(child ISQLExprParser) []ast.ISQLComment {
	comments := make([]ast.ISQLComment, 0, 10)
loop:
	for {
		switch sp.Kind() {
		case COMMENT_MINUS:
			comment := ParseMinusComment(child)
			comments = append(comments, comment)

		case COMMENT_MULTI_LINE:
			comment := ParseMultiLineComment(child)
			comments = append(comments, comment)

		case COMMENT_SHARP:
			comment := ParseSharpComment(child)
			comments = append(comments, comment)
		default:
			break loop
		}
	}

	if child.Config().SkipComment {
		return nil
	}
	return comments
}

/**
 * E: T.T.T
 * T: ID
 */
func ParseName(x ISQLExprParser) expr.ISQLName {
	owner := x.ParseIdentifier(x)
	return x.ParseNameRest(x, owner)
}

func (sp SQLExprParser) ParseIdentifier(child ISQLExprParser) expr.ISQLIdentifier {
	var name expr.ISQLIdentifier
	switch sp.Kind() {
	case IDENTIFIER:
		name = expr.NewUnQuotedIdentifier(sp.StringValue())
		NextTokenByParser(child)
		break
	case IDENTIFIER_DOUBLE_QUOTE:
		name = expr.NewDoubleQuotedIdentifier(sp.StringValue())
		NextTokenByParser(child)
		break
	case IDENTIFIER_REVERSE_QUOTE:
		name = expr.NewReverseQuotedIdentifier(sp.StringValue())
		NextTokenByParser(sp)

	case SYMB_STAR:
		name = expr.NewAllColumnExpr()
		NextTokenByParser(sp)
	}

	return name
}

func (sp *SQLExprParser) ParseNameRest(child ISQLExprParser, owner expr.ISQLName) expr.ISQLName {
	switch sp.Kind() {
	case SYMB_DOT:
		NextTokenByParser(sp)
		right := child.ParseIdentifier(child)
		name := expr.NewNameWithOwnerAndName(owner, right)
		newName := sp.ParseNameRest(child, name)
		return newName
	}
	return owner
}

/**
 * expr
 */
func ParseExpr(sp ISQLExprParser) expr.ISQLExpr {
	if sp.Kind() == SYMB_SEMI ||
		sp.Kind() == SYMB_RIGHT_PAREN ||
		sp.Kind() == EOF {
		return nil
	}
	comments := sp.ParseComments(sp)
	left := ParsePrimaryExpr(sp)
	if left == nil {
		return left
	}

	leftAfterComments := sp.ParseComments(sp)
	left.AddAfterComments(leftAfterComments)

	x := sp.ParseExprRest(sp, left)
	if x != nil {
		x.AddBeforeComments(comments)
	}
	return x
}

func ParsePrimaryExpr(sp ISQLExprParser) expr.ISQLExpr {
	if sp.Kind() == SYMB_SEMI ||
		sp.Kind() == SYMB_RIGHT_PAREN ||
		sp.Kind() == EOF {
		return nil
	}

	var x expr.ISQLExpr

	switch sp.Kind() {
	case IDENTIFIER:
		x = expr.NewUnQuotedIdentifier(sp.StringValue())
		NextTokenByParser(sp)

	case IDENTIFIER_DOUBLE_QUOTE:
		x = expr.NewDoubleQuotedIdentifier(sp.StringValue())
		NextTokenByParser(sp)

	case IDENTIFIER_REVERSE_QUOTE:
		x = expr.NewReverseQuotedIdentifier(sp.StringValue())
		NextTokenByParser(sp)
		return x

	case LITERAL_STRING:
		x = literal.NewStringLiteral(sp.StringValue())
		NextTokenByParser(sp)
		return x

	case LITERAL_INTEGER:
		x = literal.NewIntegerLiteralWithString(sp.StringValue())
		NextTokenByParser(sp)
		return x

	case LITERAL_FLOATING_POINT:
		x = literal.NewFloatingPointLiteralWith(sp.StringValue())
		NextTokenByParser(sp)
		return x

	case LITERAL_HEXADECIMAL_X:
		x = literal.NewHexadecimalLiteral(literal.X_UPPER, sp.StringValue())
		NextTokenByParser(sp)
		return x

	case LITERAL_HEXADECIMAL_0X:
		x = literal.NewHexadecimalLiteral(literal.ZERO_X, sp.StringValue())
		NextTokenByParser(sp)
		return x

	case DATE:
		NextTokenByParser(sp)
		x = literal.NewDateLiteral(ParseExpr(sp))
		return x

	case TIME:
		NextTokenByParser(sp)
		x = literal.NewTimeLiteral(ParseExpr(sp))
		return x

	case TIMESTAMP:
		NextTokenByParser(sp)
		x = literal.NewTimestampLiteral(ParseExpr(sp))
		return x

	case TRUE:
		x = literal.NewBooleanLiteral(true)
		NextTokenByParser(sp)
		return x

	case FALSE:
		x = literal.NewBooleanLiteral(false)
		NextTokenByParser(sp)
		return x

	case INTERVAL:
		x = ParseIntervalLiteral(sp)
		return x

	case SYMB_STAR:
		x = expr.NewAllColumnExpr()
		NextTokenByParser(sp)
		return x

	case SYMB_QUESTION:
		x = variable.NewVariableExpr()
		NextTokenByParser(sp)
		return x

	case PRIOR:
		NextTokenByParser(sp)
		operand := ParsePrimaryExpr(sp)
		x = operator.NewUnaryOperatorExpr(operator.PRIOR, operand)
		return x
	case CONNECT_BY_ROOT:
		NextTokenByParser(sp)
		operand := ParsePrimaryExpr(sp)
		x = operator.NewUnaryOperatorExpr(operator.CONNECT_BY_ROOT, operand)
		return x
	case COLLATE:
		NextTokenByParser(sp)
		operand := ParsePrimaryExpr(sp)
		x = operator.NewUnaryOperatorExpr(operator.COLLATE, operand)
		return x

	case SYMB_AT:
		return sp.ParseAtExpr(sp)

	case SYMB_COLON:
		NextTokenByParser(sp)
		name := ParsePrimaryExpr(sp)
		x = variable.NewBindVariableExprWithName(name)
		return x

	case WITH, SELECT:
		return ParseSubQueryExpr(sp)

	case SYMB_LEFT_PAREN:
		NextTokenByParser(sp)
		x = ParseExpr(sp)

		if sp.Accept(SYMB_COMMA) {
			listExpr := expr.NewListExpr()
			listExpr.AddElement(x)
			for {
				NextTokenByParser(sp)
				listExpr.AddElement(ParseExpr(sp))
				if !sp.Accept(SYMB_COMMA) {
					break
				}
			}
			x = listExpr
		}

		switch x.(type) {
		case *operator.SQLUnaryOperatorExpr:
			x.(*operator.SQLUnaryOperatorExpr).SetParen(true)

		case *operator.SQLBinaryOperatorExpr:
			x.(*operator.SQLBinaryOperatorExpr).SetParen(true)

		case *common.SQLSubQueryExpr:
			x.(*common.SQLSubQueryExpr).Paren = true

		case *expr.SQLListExpr:
		default:
			listExpr := expr.NewListExpr()
			listExpr.AddElement(x)
			x = listExpr
		}

		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		return x
	}

	if x == nil {
		return x
	}

	return sp.ParsePrimaryExprRest(sp, x)
}

func (sp *SQLExprParser) ParsePrimaryExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if left == nil {
		panic("expr is nil.")
	}

	if sp.Accept(SYMB_SEMI) ||
		sp.Accept(EOF) {
		return left
	}

	// .
	if sp.Accept(SYMB_DOT) {
		return sp.parseDotExpr(child, left)
	}

	// @
	if sp.Accept(SYMB_AT) {
		return sp.ParseDBLinkExpr(child, left)
	}

	// =>
	if sp.Accept(SYMB_EQUAL_GREATER_THAN) {
		return child.ParseCallExpr(child, left)
	}

	// (
	if sp.Accept(SYMB_LEFT_PAREN) {

		outerJoinExpr, ok := child.ParseOuterJoinExprRest(left)
		if ok {
			return outerJoinExpr
		}
	}

	left = sp.ParseFunctionWithName(child, left)

	return left
}

func ParseSubQueryExpr(child ISQLExprParser) *common.SQLSubQueryExpr {
	if !child.Accept(WITH) && !child.Accept(SELECT) {
		return nil
	}
	x := common.NewSubQueryExpr()
	query := ParseSelectQuery(child)
	x.SetQuery(query)
	return x
}

func (sp *SQLExprParser) ParseAtExpr(child ISQLExprParser) expr.ISQLExpr {
	panic(sp.UnSupport())
}

func ParseIntervalLiteral(sp ISQLExprParser) *literal.SQLIntervalLiteral {
	if !sp.AcceptAndNextToken(INTERVAL) {
		return nil
	}

	x := literal.NewIntervalLiteral()
	value := ParseExpr(sp)
	x.SetValue(value)
	start := ParseIntervalLiteralField(sp)
	x.SetStart(start)

	if sp.AcceptAndNextToken(TO) {
		end := ParseIntervalLiteralField(sp)
		x.SetEnd(end)
	}

	return x
}
func ParseIntervalLiteralField(sp ISQLExprParser) *literal.SQLIntervalLiteralField {

	var unit literal.SQLIntervalUnit

	if sp.AcceptAndNextToken(YEAR) {
		unit = literal.YEAR
	} else if sp.AcceptAndNextToken(MONTH) {
		unit = literal.MONTH
	} else if sp.AcceptAndNextToken(DAY) {
		unit = literal.DAY
	} else if sp.AcceptAndNextToken(HOUR) {
		unit = literal.HOUR
	} else if sp.AcceptAndNextToken(MINUTE) {
		unit = literal.MINUTE
	} else if sp.AcceptAndNextToken(SECOND) {
		unit = literal.SECOND
	} else {
		panic(sp.UnSupport())
	}

	x := literal.NewIntervalLiteralFieldWitUnit(unit)
	paren := sp.AcceptAndNextToken(SYMB_LEFT_PAREN)
	if paren {
		for {

			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
	}

	return x
}

/**
 * x .xx
 *
 * @param expr
 * @return xx.xx
 */
func (sp *SQLExprParser) parseDotExpr(child ISQLExprParser, owner expr.ISQLExpr) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(SYMB_DOT) {
		return owner
	}

	name := child.ParseIdentifier(child)
	owner = expr.NewNameWithOwnerAndName(owner, name)
	owner = sp.parseDotExpr(child, owner)

	return child.ParsePrimaryExprRest(child, owner)
}

/**
 * xx@xx
 */
func (sp *SQLExprParser) ParseDBLinkExpr(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(SYMB_AT) {
		return name
	}
	dbLink := ParseName(child)
	return expr.NewDBLinkExpr(name, dbLink)
}

// (+)
func (self *SQLExprParser) ParseOuterJoinExprRest(x expr.ISQLExpr) (expr.ISQLExpr, bool) {
	mark := self.Mark()

	if !self.AcceptAndNextToken(SYMB_LEFT_PAREN) {
		return x, false
	}

	if self.Accept(SYMB_PLUS) {
		self.ResetWithMark(mark)
		panic(self.Token().UnSupport())
	}

	self.ResetWithMark(mark)

	return x, false
}

/**
 * =>
 */
func (self *SQLExprParser) ParseCallExpr(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {
	panic(self.Token().UnSupport())
}

/**
 * name [=] value
 */
func ParseAssignExpr(child ISQLExprParser) *expr.SQLAssignExpr {
	x := expr.NewAssignExpr()
	name := ParsePrimaryExpr(child)
	x.SetName(name)

	equal := child.AcceptAndNextToken(SYMB_EQUAL)
	x.Equal = equal

	value := ParseExpr(child)
	x.SetValue(value)

	return x
}

// --------------------------- function --------------------------------------------------

/**
 * xx
 * x.xx()
 *
 * @param expr
 * @return xx.xx
 */
func (sp *SQLExprParser) ParseFunctionWithName(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {

	if sp.AcceptAndNextToken(SYMB_LEFT_PAREN) {

		// 特殊处理
		switch name.(type) {
		case *expr.SQLUnQuotedIdentifier:
			identifier := name.(*expr.SQLUnQuotedIdentifier)
			if child.IsComplexFunction(identifier.Name()) {
				method := child.ParseComplexFunction(child, name)
				sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
				return sp.ParsePrimaryExprRest(child, method)
			}
		}

		methodInvocation := function.NewMethodInvocation(name)

		if !sp.Accept(SYMB_RIGHT_PAREN) {
			for {
				argument := ParseExpr(child)
				methodInvocation.AddArgument(argument)
				if !sp.AcceptAndNextToken(SYMB_COMMA) {
					break
				}
			}
		}

		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		return sp.ParsePrimaryExprRest(child, methodInvocation)

	} else {

	}

	return name
}

func (sp *SQLExprParser) IsComplexFunction(name string) bool {
	return complexFunctionNameMap[strings.ToUpper(name)]
}

func (sp *SQLExprParser) ParseComplexFunction(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {
	return name
}

func (sp *SQLExprParser) IsNonParametricFunction(name string) bool {
	return nonParametricFunctionNameMap[strings.ToUpper(name)]
}

func (sp *SQLExprParser) ParseNonParametricFunction(child ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {
	return name
}

// --------------------------- function --------------------------------------------------

func (sp *SQLExprParser) ParseExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if sp.Accept(SYMB_SEMI) ||
		sp.Accept(EOF) {
		return left
	}

	left = sp.ParseBitXorOperatorExprRest(child, left)
	left = sp.ParseMultiplicativeOperatorExprRest(child, left)
	left = sp.ParseAdditiveOperatorExprRest(child, left)
	left = sp.ParseShiftOperatorExprRest(child, left)
	left = sp.ParseBitAndOperatorExprRest(child, left)
	left = sp.ParseBitOrOperatorExprRest(child, left)
	left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	left = sp.ParseLowPriorityComparisonOperatorExprRest(child, left)
	left = sp.ParseAndOperatorExprRest(child, left)
	left = sp.ParseXOROperatorExprRest(child, left)
	left = sp.ParseOrOperatorExprRest(child, left)
	return left
}

/**
 * E: T ^ T ^ T
 * T: primaryExpr
 * https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *SQLExprParser) ParseBitXorOperatorExpr(child ISQLExprParser) expr.ISQLExpr {
	left := ParsePrimaryExpr(child)
	return x.ParseBitXorOperatorExprRest(child, left)
}
func (x *SQLExprParser) ParseBitXorOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	switch x.Kind() {
	case SYMB_BIT_XOR:
		NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseBitXorOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X ^ X
* OP: *, /, DIV, %, MOD
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *SQLExprParser) ParseMultiplicativeOperatorExpr(child ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseBitXorOperatorExpr(child)
	return sp.ParseMultiplicativeOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseMultiplicativeOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	switch sp.Kind() {
	case SYMB_STAR:
		NextTokenByParser(sp)
		right := sp.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MULTIPLY, right)
		return sp.ParseMultiplicativeOperatorExprRest(child, left)

	case SYMB_SLASH:
		NextTokenByParser(sp)
		right := sp.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.DIVIDE, right)
		return sp.ParseMultiplicativeOperatorExprRest(child, left)

	case SYMB_PERCENT:
		NextTokenByParser(sp)
		right := sp.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MODULO, right)
		return sp.ParseMultiplicativeOperatorExprRest(child, left)

	case DIV:
		NextTokenByParser(sp)
		right := sp.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.DIV, right)
		return sp.ParseMultiplicativeOperatorExprRest(child, left)

	case MOD:
		NextTokenByParser(sp)
		right := sp.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MOD, right)
		return sp.ParseMultiplicativeOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T (+/-) T (+/-) T ...
* T: X (*, /, DIV, %, MOD) X
* +, -
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *SQLExprParser) ParseAdditiveOperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseMultiplicativeOperatorExpr(child)
	return sp.ParseAdditiveOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseAdditiveOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	switch sp.Kind() {
	case SYMB_PLUS:
		NextTokenByParser(sp)
		right := sp.ParseMultiplicativeOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.PLUS, right)
		return sp.ParseAdditiveOperatorExprRest(child, left)

	case SYMB_MINUS:
		NextTokenByParser(sp)
		right := sp.ParseMultiplicativeOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MINUS, right)
		return sp.ParseAdditiveOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T
* T: X +/- X
* op: <<, >>
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *SQLExprParser) ParseShiftOperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseAdditiveOperatorExpr(child)
	return sp.ParseShiftOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseShiftOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch sp.Kind() {
	case SYMB_LESS_THAN_LESS_THAN:
		NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return sp.ParseShiftOperatorExprRest(child, left)
	case SYMB_GREATER_THAN_GREATER_THAN:
		NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_RIGHT, right)
		return sp.ParseShiftOperatorExprRest(child, left)
	}

	return left
}

/**
* E: T & T & T
* T: X (<<,>> ) X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *SQLExprParser) ParseBitAndOperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseShiftOperatorExpr(child)
	return sp.ParseBitAndOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseBitAndOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch sp.Kind() {
	case SYMB_BIT_AND:
		NextTokenByParser(sp)
		right := sp.ParseShiftOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.BIT_AND, right)
		return sp.ParseBitAndOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T | T | T
* T: X & X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *SQLExprParser) ParseBitOrOperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseBitAndOperatorExpr(child)
	return sp.ParseBitOrOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseBitOrOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch sp.Kind() {
	case SYMB_BIT_OR:
		NextTokenByParser(sp)
		right := sp.ParseBitAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.BIT_OR, right)
		return sp.ParseBitOrOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X | X
* op: = (comparison), <=>, >=, >, <=, <, <>, !=
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *SQLExprParser) ParseHighPriorityComparisonOperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseBitOrOperatorExpr(child)
	return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseHighPriorityComparisonOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	if sp.AcceptAndNextToken(SYMB_EQUAL) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.EQ, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(SYMB_LESS_THAN) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.LESS_THAN, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(SYMB_LESS_THAN_EQUAL) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.LESS_THAN_EQ, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(SYMB_GREATER_THAN) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(SYMB_GREATER_THAN_EQUAL) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN_EQ, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	}

	return left
}

/**
* E: T op T op T ...
* T: X | X
* op: IS [NOT] NULL, LIKE, [NOT] BETWEEN, [NOT] IN
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *SQLExprParser) ParseLowPriorityComparisonOperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseHighPriorityComparisonOperatorExpr(child)
	return sp.ParseLowPriorityComparisonOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseLowPriorityComparisonOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	if sp.Accept(IS) {

		return sp.ParseIsConditionRest(child, left)

	} else if sp.AcceptAndNextToken(NOT) {

		if child.IsParseLikeCondition() {

			return child.ParseLikeConditionRest(child, true, left)

		} else if sp.Accept(REGEXP) {

			return sp.ParseRegexpConditionRest(child, true, left)

		} else if sp.Accept(BETWEEN) {

			return sp.ParseBetweenConditionRest(child, true, left)

		} else if sp.Accept(IN) {

			return sp.ParseInConditionRest(child, true, left)
		}

	} else if sp.IsParseLikeCondition() {

		return sp.ParseLikeConditionRest(child, false, left)

	} else if sp.Accept(REGEXP) {

		return sp.ParseRegexpConditionRest(child, false, left)

	} else if sp.Accept(BETWEEN) {

		return sp.ParseBetweenConditionRest(child, false, left)

	} else if sp.Accept(IN) {

		return sp.ParseInConditionRest(child, false, left)
	}

	return left
}

/**
 * boolean_primary IS [NOT] {TRUE | FALSE | UNKNOWN | NULL}
 *
 *
 * https://dev.mysql.com/doc/refman/8.0/en/expressions.html
 */
func (sp *SQLExprParser) ParseIsConditionRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(IS) {
		return left
	}
	x := condition.NewIsCondition()
	x.SetExpr(left)

	x.Not = sp.AcceptAndNextToken(NOT)

	if sp.AcceptAndNextToken(NULL) {

		x.Value = condition.NULL

	} else if sp.AcceptAndNextToken(TRUE) {

		x.Value = condition.TRUE

	} else if sp.AcceptAndNextToken(FALSE) {

		x.Value = condition.FALSE

	} else if sp.AcceptAndNextToken(UNKNOWN) {

		x.Value = condition.UNKNOWN

	} else {
		panic(sp.UnSupport())
	}

	return x
}

func (sp *SQLExprParser) IsParseLikeCondition() bool {
	return sp.Accept(LIKE)
}

func (sp *SQLExprParser) ParseLikeConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr {
	if !sp.IsParseLikeCondition() {
		return left
	}
	sp.AcceptAndNextTokenWithError(LIKE, true)

	x := condition.NewLikeCondition()
	x.SetExpr(left)
	x.Not = not
	x.Like = condition.LIKE

	pattern := ParsePrimaryExpr(child)
	x.SetPattern(pattern)

	if sp.AcceptAndNextToken(ESCAPE) {
		escape := ParsePrimaryExpr(child)
		x.SetEscape(escape)
	}

	return x
}
func (sp *SQLExprParser) ParseRegexpConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(REGEXP) {
		return left
	}

	x := condition.NewBetweenCondition()
	x.SetExpr(left)
	x.Not = not

	between := ParseExpr(child)
	x.SetBetween(between)

	and := ParseExpr(child)
	x.SetAnd(and)

	return x
}
func (sp *SQLExprParser) ParseBetweenConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(BETWEEN) {
		return left
	}

	x := condition.NewBetweenCondition()
	x.SetExpr(left)
	x.Not = not

	between := ParsePrimaryExpr(sp)
	x.SetBetween(between)

	sp.AcceptAndNextTokenWithError(AND, true)
	and := ParsePrimaryExpr(child)
	x.SetAnd(and)

	return x
}
func (sp *SQLExprParser) ParseInConditionRest(child ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(IN) {
		return left
	}

	x := condition.NewInCondition()
	x.SetExpr(left)
	x.Not = not

	sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
	for {
		var value expr.ISQLExpr
		if IsSelect(child) {
			value = ParseSelectQuery(child)
		} else {
			value = ParseExpr(child)
		}
		x.AddValue(value)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	return x
}

func (sp *SQLExprParser) ParseNotComparisonOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if sp.AcceptAndNextToken(LIKE) {

	} else if sp.AcceptAndNextToken(IN) {

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)

		inCondition := condition.NewInConditionWithExpr(left)
		inCondition.Not = true
		for {

			value := ParseExpr(child)
			inCondition.AddValue(value)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, inCondition)

	}
	return left
}

/**
* E: NOT T
* T:
* op: NOT
* https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#boolean%20factor
*/
func (x *SQLExprParser) ParseNotExpr(child ISQLExprParser) expr.ISQLExpr {
	left := x.ParseHighPriorityComparisonOperatorExpr(child)
	return x.ParseNotExprRest(child, left)
}

func (sp *SQLExprParser) ParseNotExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	if sp.AcceptAndNextToken(AND) {
		right := sp.ParseHighPriorityComparisonOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.AND, right)
		return sp.ParseNotExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T
* T: X comparisonOperator X
* op: AND
* https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#boolean%20term
*/
func (x *SQLExprParser) ParseAndOperatorExpr(child ISQLExprParser) expr.ISQLExpr {
	left := x.ParseHighPriorityComparisonOperatorExpr(child)
	return x.ParseAndOperatorExprRest(child, left)
}

func (sp *SQLExprParser) ParseAndOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	if sp.AcceptAndNextToken(AND) {
		right := sp.ParseHighPriorityComparisonOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.AND, right)
		return sp.ParseAndOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T XOR T XOR T
* T: X (AND, &&) X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseXOROperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := x.ParseAndOperatorExpr(child)
	return x.ParseXOROperatorExprRest(child, left)
}

func (x *SQLExprParser) ParseXOROperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_LESS_THAN_LESS_THAN:
		NextTokenByParser(x)
		right := x.ParseAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return x.ParseXOROperatorExprRest(child, left)
	case SYMB_GREATER_THAN_GREATER_THAN:
		NextTokenByParser(x)
		right := x.ParseAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_RIGHT, right)
		return x.ParseXOROperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T
* T: X AND X
* op: OR
* https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#boolean%20value%20expression
*/
func (x *SQLExprParser) ParseOrOperatorExpr(child ISQLExprParser, ) expr.ISQLExpr {
	left := x.ParseAndOperatorExpr(child)
	return x.ParseOrOperatorExprRest(child, left)
}
func (sp *SQLExprParser) ParseOrOperatorExprRest(child ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	if sp.AcceptAndNextToken(OR) {
		right := sp.ParseAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.OR, right)
		return sp.ParseOrOperatorExprRest(child, left)
	}

	return left
}

// --------------------------- DataType --------------------------------------------------

func ParseDataType(sp ISQLExprParser) (datatype.ISQLDataType) {
	if sp.Accept(EOF) ||
		sp.Accept(SYMB_LEFT_PAREN) ||
		sp.Accept(SYMB_RIGHT_PAREN) ||
		sp.Accept(SYMB_COMMA) ||
		sp.Accept(SYMB_SEMI) {
		return nil
	}

	// Interval DataType
	if sp.Accept(INTERVAL) {
		return ParseIntervalDataType(sp)
	}

	name := ParseName(sp)
	if name == nil {
		panic(sp.SyntaxError())
	}

	// 特殊处理
	// switch name.(type) {
	// case *expr.SQLUnQuotedIdentifier:
	// 	identifier := name.(*expr.SQLUnQuotedIdentifier)
	//
	// }

	dateType := datatype.NewDataTypeWithSQLName(name, sp.DBType())

	paren := sp.AcceptAndNextToken(SYMB_LEFT_PAREN)
	dateType.Paren = paren

	if paren {
		for {
			precision := ParseExpr(sp)
			if precision == nil {
				panic(sp.SyntaxError())
			}
			dateType.AddPrecision(precision)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
	}

	return dateType
}

func ParseIntervalDataType(sp ISQLExprParser) *datatype.SQLIntervalDataType {
	if !sp.AcceptAndNextToken(INTERVAL) {
		return nil
	}

	x := datatype.NewIntervalDataType(sp.DBType())
	start := ParseIntervalDataTypeField(sp)
	x.SetStart(start)

	if sp.AcceptAndNextToken(TO) {
		end := ParseIntervalDataTypeField(sp)
		x.SetEnd(end)
	}

	return x
}
func ParseIntervalDataTypeField(sp ISQLExprParser) *datatype.SQLIntervalDataTypeField {

	var unit datatype.SQLIntervalUnit

	if sp.AcceptAndNextToken(YEAR) {
		unit = datatype.YEAR
	} else if sp.AcceptAndNextToken(MONTH) {
		unit = datatype.MONTH
	} else if sp.AcceptAndNextToken(DAY) {
		unit = datatype.DAY
	} else if sp.AcceptAndNextToken(HOUR) {
		unit = datatype.HOUR
	} else if sp.AcceptAndNextToken(MINUTE) {
		unit = datatype.MINUTE
	} else if sp.AcceptAndNextToken(SECOND) {
		unit = datatype.SECOND
	} else {
		panic(sp.UnSupport())
	}

	x := datatype.NewIntervalDataTypeFieldWitUnit(unit)
	paren := sp.AcceptAndNextToken(SYMB_LEFT_PAREN)
	if paren {
		for {

			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
	}

	return x
}

func (sep *SQLExprParser) ParseDateDataType(child ISQLExprParser) datatype.ISQLDataType {
	if !sep.AcceptAndNextToken(DATE) {
		return nil
	}

	x := datatype.NewDateDataType(sep.dbType)
	return x
}
func (spe *SQLExprParser) ParseDateTimeDataType(child ISQLExprParser) datatype.ISQLDataType {
	return nil
}
func (x *SQLExprParser) ParseTimeDataType(child ISQLExprParser) datatype.ISQLDataType {
	return nil
}

func (x *SQLExprParser) ParseTimestampDataType(child ISQLExprParser) datatype.ISQLDataType {
	return nil
}

// --------------------------- DataType --------------------------------------------------

// ------------------------------------------------------ DDL --------------------------------------------------
// --------------------------- Element --------------------------------------------------
func ParseIfNotExists(x ISQLExprParser) bool {
	if !x.AcceptAndNextToken(IF) {
		return false
	}

	x.AcceptAndNextTokenWithError(NOT, true)
	x.AcceptAndNextTokenWithError(EXISTS, true)

	return true
}

func ParseIfExists(x ISQLExprParser) bool {
	if !x.AcceptAndNextToken(IF) {
		return false
	}

	x.AcceptAndNextTokenWithError(EXISTS, true)

	return true
}

func ParseOrReplace(x ISQLExprParser) bool {
	if !x.AcceptAndNextToken(OR) {
		return false
	}
	x.AcceptAndNextTokenWithError(REPLACE, true)
	return true
}

func (sp *SQLExprParser) ParseParameterDeclaration(child ISQLExprParser) *statement.SQLParameterDeclaration {
	x := statement.NewParameterDeclaration()
	parameterMode := ParseParameterMode(child)
	x.ParameterMode = parameterMode

	name := ParseName(child)
	x.SetName(name)

	dataType := ParseDataType(child)
	x.SetDataType(dataType)

	result := sp.AcceptAndNextToken(RESULT)
	x.Result = result
	return x
}
func ParseParameterMode(sp ISQLExprParser) statement.SQLParameterMode {
	if sp.AcceptAndNextToken(IN) {
		return statement.IN

	} else if sp.AcceptAndNextToken(OUT) {
		return statement.OUT

	} else if sp.AcceptAndNextToken(INOUT) {
		return statement.INOUT

	}
	return ""
}

// --------------------------- Sequence --------------------------------------------------
func (sp *SQLExprParser) ParseSequenceOption(child ISQLExprParser) expr.ISQLExpr {
	if sp.AcceptAndNextToken(START) {
		sp.AcceptAndNextTokenWithError(WITH, true)

		x := sequence.NewStartWithSequenceOption()

		value := ParseExpr(child)
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(INCREMENT) {
		sp.AcceptAndNextTokenWithError(BY, true)

		x := sequence.NewIncrementBySequenceOption()

		value := ParseExpr(child)
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(MAXVALUE) {
		x := sequence.NewMaxValueSequenceOption()

		value := ParseExpr(child)
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(MINVALUE) {
		x := sequence.NewMaxValueSequenceOption()

		value := ParseExpr(child)
		x.SetValue(value)
		return x
	} else if sp.AcceptAndNextToken(CYCLE) {
		x := sequence.NewCycleSequenceOption()
		return x
	} else if sp.AcceptAndNextToken(NO) {
		if sp.AcceptAndNextToken(MAXVALUE) {
			return sequence.NewNoMaxValueSequenceOption()

		} else if sp.AcceptAndNextToken(MINVALUE) {
			return sequence.NewNoMinValueSequenceOption()

		} else if sp.AcceptAndNextToken(CYCLE) {
			return sequence.NewNoCycleSequenceOption()

		} else {
			panic(sp.UnSupport())
		}

	}
	return nil
}

// --------------------------- Synonym --------------------------------------------------
func (sp *SQLExprParser) ParseAlterSynonymAction(child ISQLExprParser) expr.ISQLExpr {
	if sp.AcceptAndNextToken(EDITIONABLE) {
		return statement.NewEditionAbleExpr()

	} else if sp.AcceptAndNextToken(NONEDITIONABLE) {
		return statement.NewNonEditionAbleExpr()

	} else if sp.AcceptAndNextToken(COMPILE) {

		return statement.NewCompileExpr()
	}
	return nil
}

// --------------------------- OnTable --------------------------------------------------
/**
* (
* 	TableElements
* )
*/
func ParseTableElements(sp ISQLExprParser) []table.ISQLTableElement {
	accept := sp.AcceptAndNextToken(SYMB_LEFT_PAREN)
	if !accept {
		return nil
	}

	tableElements := make([]table.ISQLTableElement, 0, 10)

	for {
		comments := sp.ParseComments(sp)
		tableElement := sp.ParseTableElement(sp)
		if tableElement != nil {
			tableElement.AddBeforeComments(comments)
		}

		tableElements = append(tableElements, tableElement)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	return tableElements
}

func (sp SQLExprParser) ParseTableElement(child ISQLExprParser) table.ISQLTableElement {

	var x table.ISQLTableElement = nil

	if child.IsParseTableConstraint() {
		return child.ParseTableConstraint(child)

	} else if sp.Accept(LIKE) {
		x = ParseLikeClause(child)

	} else if IsIdentifier(sp.Kind()) {
		x = child.ParseTableColumn(child)
	}

	return x
}

/**
 *
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20definition
 */
func (sp *SQLExprParser) ParseTableColumn(child ISQLExprParser) *table.SQLTableColumn {

	if !IsIdentifier(sp.Kind()) {
		return nil
	}

	x := table.NewColumn()

	name := sp.ParseIdentifier(sp)
	x.SetName(name)

	dataType := ParseDataType(sp)
	x.SetDataType(dataType)

	for {
		option, ok := child.ParseTableColumnOption(child)
		if !ok {
			break
		}
		x.AddOption(option)
	}
	return x
}
func (sp *SQLExprParser) IsParseTableColumnOption(child ISQLExprParser) bool {
	return child.IsParseColumnConstraint()
}
func (sp *SQLExprParser) ParseTableColumnOption(child ISQLExprParser) (expr.ISQLExpr, bool) {
	if sp.Accept(REFERENCES) {

		panic(sp.UnSupport())

	} else if sp.Accept(DEFAULT) {

		return sp.ParseDefaultClause(child), true

	} else if sp.Accept(GENERATED) {

		return child.ParseGeneratedClause(child), true

	} else if sp.Accept(COLLATE) {

		return sp.ParseCollateClause(child), true

	} else if child.IsParseColumnConstraint() {

		return child.ParseColumnConstraint(child)
	}

	return nil, false
}
func (sp *SQLExprParser) IsParseColumnConstraint() bool {
	if sp.Accept(CONSTRAINT) ||
		sp.Accept(NOT) ||
		sp.Accept(NULL) ||
		sp.Accept(UNIQUE) ||
		sp.Accept(PRIMARY) ||
		sp.Accept(CHECK) {
		return true
	}
	return false
}
func (sp *SQLExprParser) ParseColumnConstraint(child ISQLExprParser) (table.ISQLColumnConstraint, bool) {

	ok := false
	var name expr.ISQLName
	if sp.AcceptAndNextToken(CONSTRAINT) {
		name = ParseName(child)
	}

	if sp.AcceptAndNextToken(NOT) {
		sp.AcceptAndNextTokenWithError(NULL, true)
		x := table.NewNotNullColumnConstraint()
		x.SetName(name)
		ok = true
		return x, ok

	} else if sp.AcceptAndNextToken(NULL) {
		x := table.NewNullColumnConstraint()
		x.SetName(name)
		ok = true
		return x, ok

	} else if sp.AcceptAndNextToken(UNIQUE) {
		x := table.NewUniqueColumnConstraint()
		key := sp.AcceptAndNextToken(KEY)
		x.Key = key
		x.SetName(name)
		ok = true
		return x, ok

	} else if sp.AcceptAndNextToken(PRIMARY) {
		sp.AcceptAndNextTokenWithError(KEY, true)
		x := table.NewPrimaryKeyColumnConstraint()
		ok = true
		return x, ok

	} else if sp.AcceptAndNextToken(CHECK) {
		x := table.NewCheckColumnConstraint()
		x.SetName(name)
		ok = true
		return x, ok

	} else {
		x, ok := child.ParseColumnConstraintRest(child, name)
		if ok {
			return x, ok
		}
	}

	if name != nil {
		panic(sp.UnSupport())
	}

	return nil, ok
}
func (sp *SQLExprParser) ParseColumnConstraintRest(child ISQLExprParser, name expr.ISQLName) (table.ISQLColumnConstraint, bool) {
	return nil, false
}

/**
  * DEFAULT expr
 */
func (sp *SQLExprParser) ParseDefaultClause(child ISQLExprParser) *table.SQLDefaultClause {
	if !sp.AcceptAndNextToken(DEFAULT) {
		return nil
	}
	x := table.NewDefaultClause()
	value := ParseExpr(child)
	x.SetValue(value)
	return x
}

/**
 * GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ <left paren> <common sequence generator options> <right paren> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#identity%20column%20specification
 *
 * GENERATED ALWAYS AS <left paren> <value expression> <right paren>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#generation%20clause
 */
func (sp *SQLExprParser) ParseGeneratedClause(child ISQLExprParser) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(GENERATED) {
		return nil
	}
	var identityGeneratedKind table.SQLIdentityGeneratedKind
	if sp.AcceptAndNextToken(ALWAYS) {
		identityGeneratedKind = table.ALWAYS

	} else if sp.AcceptAndNextToken(BY) {
		sp.AcceptAndNextTokenWithError(DEFAULT, true)
		identityGeneratedKind = table.BY_DEFAULT
	}

	sp.AcceptAndNextTokenWithError(AS, true)

	identity := sp.AcceptAndNextToken(IDENTITY)
	if identity {
		x := table.NewIdentityColumnClause()
		x.IdentityGeneratedKind = identityGeneratedKind
		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			option := child.ParseIdentityColumnClauseOption(child)
			if option == nil {
				break
			}
			x.AddOption(option)
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		return x
	} else {
		x := table.NewGenerationClause()
		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)

		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		return x
	}
}

func (sp *SQLExprParser) ParseIdentityColumnClauseOption(child ISQLExprParser) expr.ISQLExpr {
	return nil
}

func (sp *SQLExprParser) ParseCommentExpr(child ISQLExprParser) *table.SQLCommentExpr {
	if !sp.AcceptAndNextToken(COMMENT) {
		return nil
	}
	x := table.NewCommentExpr()
	comment := ParseExpr(child)
	x.SetComment(comment)
	return x
}
func (sp *SQLExprParser) ParseCollateClause(child ISQLExprParser) *common.SQLCollateClause {
	if !sp.AcceptAndNextToken(COLLATE) {
		return nil
	}
	x := common.NewCollateClause()
	name := ParseName(child)
	x.SetName(name)
	return x
}
func (sp *SQLExprParser) IsParseTableConstraint() bool {
	if sp.Accept(CONSTRAINT) ||
		sp.Accept(UNIQUE) ||
		sp.Accept(PRIMARY) ||
		sp.Accept(FOREIGN) ||
		sp.Accept(CHECK) {
		return true
	}
	return false
}
func (sp *SQLExprParser) ParseTableConstraint(child ISQLExprParser) table.ISQLTableConstraint {
	if sp.IsParseTableConstraint() {
		return nil
	}

	var name expr.ISQLName
	if sp.AcceptAndNextToken(CONSTRAINT) {
		name = ParseName(child)
	}

	if sp.AcceptAndNextToken(UNIQUE) {

		x := table.NewUniqueTableConstraint()
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			column := ParseExpr(child)
			x.AddColumn(column)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		return x

	} else if sp.AcceptAndNextToken(PRIMARY) {
		sp.AcceptAndNextTokenWithError(KEY, true)

		x := table.NewPrimaryKeyTableConstraint()
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			column := ParseExpr(child)
			x.AddColumn(column)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		return x

	} else if sp.AcceptAndNextToken(FOREIGN) {
		x := table.NewForeignKeyTableConstraint()
		x.SetName(name)

		referencingIndex := ParseName(child)
		x.SetReferencingIndex(referencingIndex)

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			referencingColumn := ParseExpr(child)
			x.AddReferencingColumn(referencingColumn)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		sp.AcceptAndNextTokenWithError(REFERENCES, true)
		referencedTable := ParseName(child)
		x.SetReferencedTable(referencedTable)

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			referencedColumn := ParseExpr(child)
			x.AddReferencedColumn(referencedColumn)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		return x

	} else if sp.AcceptAndNextToken(CHECK) {
		x := table.NewCheckTableConstraint()
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		condition := ParseExpr(child)
		x.SetCondition(condition)
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		return x

	} else {
		x, ok := child.ParseTableConstraintRest(child, name)
		if ok {
			return x
		}
	}

	if name != nil {
		panic(sp.UnSupport())
	}

	return nil
}

func (sp *SQLExprParser) ParseTableConstraintRest(child ISQLExprParser, name expr.ISQLName) (table.ISQLTableConstraint, bool) {
	return nil, false
}

func ParseLikeClause(p ISQLExprParser) *table.SQLTableLikeClause {
	if !p.AcceptAndNextToken(LIKE) {
		return nil
	}

	x := table.NewTableLikeClause()

	name := ParseName(p)
	if name == nil {
		panic(p.SyntaxError())
	}
	x.SetName(name)

	return x
}

func (sp *SQLExprParser) ParsePartitionBy(child ISQLExprParser) table.ISQLPartitionBy {
	if !sp.AcceptAndNextToken(PARTITION) {
		return nil
	}
	child.AcceptAndNextTokenWithError(BY, true)

	if child.AcceptAndNextToken(LINEAR) {

	}

	if child.AcceptAndNextToken(HASH) {
		x := table.NewPartitionByHash()
		child.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			column := ParseExpr(child)
			x.AddColumn(column)
			if !child.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		child.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		return x
	}

	panic(child.SyntaxError())
}
func (sp *SQLExprParser) ParsePartitionByRest(child ISQLExprParser, x table.ISQLPartitionBy) {
	sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
	for {
		column := ParseExpr(child)
		x.AddColumn(column)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	// PARTITIONS num
	if sp.AcceptAndNextToken(PARTITIONS) {
		partitionsNum := ParseExpr(child)
		x.SetPartitionsNum(partitionsNum)
	}

	// SUBPARTITION BY
	if sp.Accept(SUBPARTITION) {
		subPartitionBy := child.ParseSubPartitionBy(child)
		x.SetSubPartitionBy(subPartitionBy)
	}

	// (partition_definitions)
	if sp.AcceptAndNextToken(SYMB_LEFT_PAREN) {
		for {
			partitionDefinition := child.ParsePartitionDefinition(child)
			x.AddPartitionDefinition(partitionDefinition)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
	}
}

func (sp *SQLExprParser) ParseSubPartitionBy(child ISQLExprParser) table.ISQLSubPartitionBy {
	if !sp.AcceptAndNextToken(SUBPARTITION) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(BY, true)

	linear := sp.AcceptAndNextToken(LINEAR)

	if sp.AcceptAndNextToken(HASH) {
		x := table.NewSubPartitionByHash()
		x.Linear = linear

		sp.ParseSubPartitionByRest(child, x)
		return x

	} else if sp.AcceptAndNextToken(KEY) {
		x := table.NewSubPartitionByKey()
		x.Linear = linear

		sp.ParseSubPartitionByRest(child, x)
		return x

	}

	panic(sp.UnSupport())
}
func (sp *SQLExprParser) ParseSubPartitionByRest(child ISQLExprParser, x table.ISQLSubPartitionBy) {
	sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
	for {
		column := ParseExpr(child)
		x.AddColumn(column)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	// SUBPARTITIONS num
	if sp.AcceptAndNextToken(SUBPARTITIONS) {
		subPartitionsNum := ParseExpr(child)
		x.SetSubPartitionsNum(subPartitionsNum)
	}
}

/**
 * PARTITION partition_name
        [VALUES
            {LESS THAN {(expr | value_list) | MAXVALUE}
            |
            IN (value_list)}]
        [[STORAGE] ENGINE [=] engine_name]
        [COMMENT [=] 'string' ]
        [DATA DIRECTORY [=] 'data_dir']
        [INDEX DIRECTORY [=] 'index_dir']
        [MAX_ROWS [=] max_number_of_rows]
        [MIN_ROWS [=] min_number_of_rows]
        [TABLESPACE [=] tablespace_name]
        [(subpartition_definition [, subpartition_definition] ...)]
 */
func (sp *SQLExprParser) ParsePartitionDefinition(child ISQLExprParser) *table.SQLPartitionDefinition {
	if !sp.AcceptAndNextToken(PARTITION) {
		return nil
	}

	x := table.NewPartitionDefinition()
	name := ParseName(child)
	x.SetName(name)

	return x
}

func (sp *SQLExprParser) ParsePartitionDefinitionOption(child ISQLExprParser) (expr.ISQLExpr, bool) {
	return nil, false
}

/**
 * SUBPARTITION logical_name
        [[STORAGE] ENGINE [=] engine_name]
        [COMMENT [=] 'string' ]
        [DATA DIRECTORY [=] 'data_dir']
        [INDEX DIRECTORY [=] 'index_dir']
        [MAX_ROWS [=] max_number_of_rows]
        [MIN_ROWS [=] min_number_of_rows]
        [TABLESPACE [=] tablespace_name]
 */
func (sp *SQLExprParser) ParseSubPartitionDefinition(child ISQLExprParser) *table.SQLSubPartitionDefinition {
	if !sp.AcceptAndNextToken(SUBPARTITION) {
		return nil
	}

	x := table.NewSubPartitionDefinition()
	name := ParseName(child)
	x.SetName(name)

	return x
}

/**
 * [VALUES
			{LESS THAN {(expr | value_list) | MAXVALUE}
            |
            IN (value_list)}
		]
 */
func (sp *SQLExprParser) ParsePartitionValues(child ISQLExprParser) table.ISQLPartitionValues {
	if !sp.AcceptAndNextToken(VALUES) {
		return nil
	}

	if sp.AcceptAndNextToken(LESS) {
		sp.AcceptAndNextTokenWithError(THAN, true)

		if sp.AcceptAndNextToken(MAXVALUE) {

			return table.NewPartitionValuesLessThanMaxValue()

		} else {
			x := table.NewPartitionValuesLessThan()
			sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
			for {
				value := ParseExpr(child)
				x.AddValue(value)
				if !sp.AcceptAndNextToken(SYMB_COMMA) {
					break
				}
			}
			sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

			return x
		}

	} else if sp.AcceptAndNextToken(IN) {
		x := table.NewPartitionValuesIn()

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			value := ParseExpr(child)
			x.AddValue(value)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		return x
	}

	panic(child.SyntaxError())
}

/**
 * <add column definition>
     |     <alter column definition>
     |     <drop column definition>
     |     <add table constraint definition>
     |     <drop table constraint definition>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20table%20action
 *
 *
 *
 */
func (sp *SQLExprParser) ParseAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if sp.Accept(ADD) {
		return child.ParseAddAlterTableAction(child)

	} else if sp.Accept(ALTER) {
		return child.ParseAlterAlterTableAction(child)

	} else if sp.Accept(DROP) {
		return child.ParseDropAlterTableAction(child)

	}

	return nil
}
func (sp *SQLExprParser) ParseAddAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(ADD) {
		return nil
	}
	if sp.AcceptAndNextToken(CONSTRAINT) {

		ParseName(child)

		if sp.AcceptAndNextToken(PRIMARY) {
			sp.AcceptAndNextTokenWithError(KEY, true)

		} else if sp.AcceptAndNextToken(UNIQUE) {

		} else if sp.AcceptAndNextToken(FOREIGN) {

		} else if sp.AcceptAndNextToken(CHECK) {

		}

	} else if sp.AcceptAndNextToken(PRIMARY) {
		sp.AcceptAndNextTokenWithError(KEY, true)

	} else if sp.AcceptAndNextToken(UNIQUE) {

	} else if sp.AcceptAndNextToken(FOREIGN) {

	} else if sp.Accept(COLUMN) || IsIdentifier(sp.Kind()) {
		return child.ParseAddColumnAlterTableAction(child)
	}

	panic(sp.UnSupport())
}
func (sp *SQLExprParser) ParseAddColumnAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.Accept(COLUMN) && !IsIdentifier(sp.Kind()) {
		return nil
	}

	x := table.NewAddColumnAlterTableAction()

	hasColumn := sp.AcceptAndNextToken(COLUMN)
	x.HasColumn = hasColumn

	paren := sp.AcceptAndNextToken(SYMB_LEFT_PAREN)
	column := child.ParseTableColumn(child)
	x.AddColumn(column)
	if paren {
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
	}
	return x
}

/**
 *  ALTER [ COLUMN ] <column name> <alter column action>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20definition
 */
func (sp *SQLExprParser) ParseAlterAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(ALTER) {
		return nil
	}
	hasColumn := sp.AcceptAndNextToken(COLUMN)
	x := table.NewAddColumnAlterTableAction()
	x.HasColumn = hasColumn

	return x
}
func (sp *SQLExprParser) ParseDropAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}

	if sp.AcceptAndNextToken(PRIMARY) {
		sp.AcceptAndNextTokenWithError(KEY, true)
		return table.NewDropPrimaryKeyTableConstraintAlterTableAction()

	} else if sp.AcceptAndNextToken(FOREIGN) {
		sp.AcceptAndNextTokenWithError(KEY, true)
		x := table.NewDropForeignKeyTableConstraintAlterTableAction()

		name := ParseName(child)
		x.SetName(name)

		return x

	} else if sp.AcceptAndNextToken(CHECK) {

		x := table.NewDropCheckTableConstraintAlterTableAction()

		name := ParseName(child)
		x.SetName(name)

		return x

	} else if sp.Accept(COLUMN) || IsIdentifier(sp.Kind()) {
		x := table.NewDropColumnAlterTableAction()

		hasColumn := sp.AcceptAndNextToken(COLUMN)
		x.HasColumn = hasColumn

		column := ParseName(child)
		x.SetColumn(column)
		return x
	}

	panic(sp.UnSupport())
}
func (sp *SQLExprParser) ParseChangeAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(CHANGE) {
		return nil
	}

	panic(sp.UnSupport())
}
func (sp *SQLExprParser) ParseModifyAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(MODIFY) {
		return nil
	}

	panic(sp.UnSupport())
}
func (sp *SQLExprParser) ParseRenameAlterTableAction(child ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(RENAME) {
		return nil
	}

	panic(sp.UnSupport())
}

// --------------------------- View --------------------------------------------------
func ParseViewElements(sp ISQLExprParser) []view.ISQLViewElement {
	if !sp.AcceptAndNextToken(SYMB_LEFT_PAREN) {
		return nil
	}

	viewElements := make([]view.ISQLViewElement, 0, 10)

	for {
		comments := sp.ParseComments(sp)
		viewElement := sp.ParseViewElement(sp)
		if viewElement != nil {
			viewElement.AddBeforeComments(comments)
		}

		viewElements = append(viewElements, viewElement)

		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	return viewElements
}
func (sp *SQLExprParser) ParseViewElement(child ISQLExprParser) view.ISQLViewElement {
	if sp.IsParseTableConstraint() {
		return child.ParseTableConstraint(child)

	} else if IsIdentifier(sp.Kind()) {

		return child.ParseViewColumn(child)
	}
	return nil
}
func (sp *SQLExprParser) ParseViewColumn(child ISQLExprParser) *view.SQLViewColumn {
	x := view.NewViewColumn()
	name := ParseName(child)
	x.SetName(name)

	for {
		option, ok := child.ParseViewColumnOption(child)
		if !ok {
			break
		}
		x.AddOption(option)
	}

	return x
}
func (sp *SQLExprParser) ParseViewColumnOption(child ISQLExprParser) (expr.ISQLExpr, bool) {
	return nil, false
}

// ------------------------------------------------------ DML --------------------------------------------------

// ------------------------ Delete

/**
 *  Select
 */
func ParseSelectQuery(sp ISQLExprParser) select_.ISQLSelectQuery {
	x := ParseQueryBlock(sp)
	if x == nil {
		return nil
	}

	orderByClause := ParseOrderByClause(sp)
	x.SetOrderByClause(orderByClause)

	limitClause := ParseLimitClause(sp)
	x.SetLimitClause(limitClause)

	lockClause := sp.ParseLockClause(sp)
	x.SetLockClause(lockClause)

	return x
}

func ParseQueryBlock(sp ISQLExprParser) select_.ISQLSelectQuery {
	x := ParsePrimaryQueryBlock(sp)
	x = parseUnionQueryBlock(sp, x)
	return x
}

func ParsePrimaryQueryBlock(sp ISQLExprParser) select_.ISQLSelectQuery {
	if sp.AcceptAndNextToken(SYMB_LEFT_PAREN) {

		selectQuery := ParseSelectQuery(sp)
		query := select_.NewParenSelectQuery(selectQuery)
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		return query
	}

	withClause := sp.ParseWithClause(sp)
	if !sp.AcceptAndNextToken(SELECT) {
		return nil
	}

	query := select_.NewSelectQuery()
	query.SetWithClause(withClause)

	// SQLSetQuantifier setQuantifier := parseSetQuantifier()
	// query.setSetQuantifier(setQuantifier)

	ParseSelectElements(sp, query)

	ParseSelectTargetElements(sp, query)

	fromClause := ParseFrom(sp)
	query.SetFromClause(fromClause)

	whereClause := ParseWhereClause(sp)
	query.SetWhereClause(whereClause)

	hierarchicalQueryClause := ParseHierarchicalQueryClause(sp)
	query.SetHierarchicalQueryClause(hierarchicalQueryClause)

	groupByClause := parseGroupBy(sp)
	query.SetGroupByClause(groupByClause)

	// windowClause := parseWindow()
	// query.setWindowClause(windowClause)

	orderByClause := ParseOrderByClause(sp)
	query.SetOrderByClause(orderByClause)

	limitClause := ParseLimitClause(sp)
	query.SetLimitClause(limitClause)

	lockClause := sp.ParseLockClause(sp)
	query.SetLockClause(lockClause)

	return query
}

func (sp *SQLExprParser) ParseWithClause(child ISQLExprParser) select_.ISQLWithClause {
	if !sp.AcceptAndNextToken(WITH) {
		return nil
	}

	x := select_.NewWithClause()
	x.Recursive = sp.AcceptAndNextToken(RECURSIVE)

	name := ParseExpr(sp)

	if sp.IsParserSubQueryFactoringClauseRest() {

		sp.ParseSubQueryFactoringClauseRest(child, name)

	} else if sp.IsParserSubAvFactoringClauseRest() {

		sp.ParseSubQueryFactoringClauseRest(child, name)

	} else {
		panic(sp.UnSupport())
	}

	return x
}

func (sp *SQLExprParser) IsParserSubQueryFactoringClauseRest() bool {
	return sp.Accept(SYMB_LEFT_PAREN) || sp.Accept(AS)
}

func (sp *SQLExprParser) ParseSubQueryFactoringClauseRest(child ISQLExprParser, name expr.ISQLExpr) select_.ISQLFactoringClause {
	if !sp.IsParserSubQueryFactoringClauseRest() {
		return nil
	}

	x := select_.NewSubQueryFactoringClause()

	x.SetName(name)

	if sp.AcceptAndNextToken(SYMB_LEFT_PAREN) {

		for {

			column := child.ParseIdentifier(child)
			x.AddColumn(column)

			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
	}

	sp.AcceptAndNextTokenWithError(AS, true)

	sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)

	subQuery := ParseSelectQuery(child)
	x.SetSubQuery(subQuery)

	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	return x

}
func (sp *SQLExprParser) IsParserSubAvFactoringClauseRest() bool {
	return false
}

func (sp *SQLExprParser) ParseSubAvFactoringClauseRest(child ISQLExprParser, name expr.ISQLExpr) select_.ISQLFactoringClause {
	if !sp.IsParserSubAvFactoringClauseRest() {
		return nil
	}
	sp.AcceptAndNextTokenWithError(ANALYTIC, true)
	sp.AcceptAndNextTokenWithError(VIEW, true)
	sp.AcceptAndNextTokenWithError(AS, true)

	x := select_.NewSubQueryFactoringClause()

	sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)

	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	return x
}

func ParseSelectElements(sp ISQLExprParser, parent *select_.SQLSelectQuery) {
	for {
		comments := sp.ParseComments(sp)
		element := ParseSelectElement(sp)
		if element != nil {
			element.AddBeforeComments(comments)
		}

		parent.AddSelectElement(element)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
}

func ParseSelectTargetElements(parser ISQLExprParser, parent *select_.SQLSelectQuery) {
	if !parser.AcceptAndNextToken(INTO) {
		return
	}

	for {
		expr := ParseExpr(parser)
		parent.AddSelectTargetElement(expr)
		if !parser.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
}

func ParseSelectElement(sp ISQLExprParser) *select_.SQLSelectElement {
	expr := ParseExpr(sp)

	if sp.Accept(SYMB_COMMA) || sp.Accept(FROM) || sp.Accept(BULK) {
		return select_.NewSelectElementWithExpr(expr)
	}

	as := sp.AcceptAndNextToken(AS)

	alias := ParseExpr(sp)
	if as && alias == nil {
		panic("alias is null. " + sp.Token().Error())
	}

	return select_.NewSelectElementWithAlias(expr, as, alias)
}

func parseUnionQueryBlock(parser ISQLExprParser, left select_.ISQLSelectQuery) select_.ISQLSelectQuery {

	var operator select_.SQLUnionOperator
	if parser.AcceptAndNextToken(UNION) {

		operator = select_.UNION

		if parser.AcceptAndNextToken(ALL) {

			operator = select_.UNION_ALL

		} else if parser.AcceptAndNextToken(DISTINCT) {

			operator = select_.UNION_DISTINCT
		}

	} else if parser.AcceptAndNextToken(MINUS) {

		operator = select_.MINUS

	} else if parser.AcceptAndNextToken(EXCEPT) {
		operator = select_.EXCEPT

		if parser.AcceptAndNextToken(ALL) {
			operator = select_.EXCEPT_ALL

		} else if parser.AcceptAndNextToken(DISTINCT) {
			operator = select_.EXCEPT_DISTINCT
		}

	} else if parser.AcceptAndNextToken(INTERSECT) {
		operator = select_.INTERSECT

		if parser.AcceptAndNextToken(ALL) {
			operator = select_.INTERSECT_ALL

		} else if parser.AcceptAndNextToken(DISTINCT) {
			operator = select_.INTERSECT_DISTINCT
		}
	}

	if operator != "" {
		right := ParsePrimaryQueryBlock(parser)
		unionQuery := select_.NewSelectUnionQuery(left, operator, right)
		return parseUnionQueryBlock(parser, unionQuery)
	}

	return left
}

func ParseFrom(sp ISQLExprParser) *select_.SQLFromClause {
	if !sp.AcceptAndNextToken(FROM) {
		return nil
	}
	tableReference := ParseTableReference(sp)

	return select_.NewFromClause(tableReference)
}

/**
 * E: T (JOIN, COMM) T JOIN T
 * T: primary
 */
func ParseTableReference(sp ISQLExprParser) select_.ISQLTableReference {
	tableReference := ParsePrimaryTableReference(sp)
	return parseJoinTableReference(sp, tableReference)
}

/**
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20primary
 */
func ParsePrimaryTableReference(sp ISQLExprParser) select_.ISQLTableReference {

	if sp.AcceptAndNextToken(SYMB_LEFT_PAREN) {
		x := ParseTableReference(sp)
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		as := false
		if sp.AcceptAndNextToken(AS) {
			as = true
		}

		alias := sp.ParseIdentifier(sp)
		if as && alias == nil {
			panic("TODO....")
		}

		x.SetParen(true)
		x.SetAs(as)
		x.SetAlias(alias)

		return x

	} else if sp.AcceptAndNextToken(SYMB_LERT_BRACE) {
		sp.AcceptAndNextTokenWithError(OJ, true)

		x := select_.NewOJTableReference()
		tableReference := ParseTableReference(sp)
		x.SetTableReference(tableReference)

		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_BRACE, true)

		return x
	} else if sp.Accept(WITH) || sp.Accept(SELECT) {
		subQuery := ParseSelectQuery(sp)

		x := select_.NewSubQueryTableReference(subQuery)

		return x

	} else if sp.Accept(ONLY) {

	} else if sp.Accept(TABLE) {
		// tableReference = parserTableFunctionTableReference()

	} else if sp.AcceptAndNextToken(UNNEST) {

	} else if IsIdentifier(sp.Kind()) {

		return sp.ParseTableReference(sp)
	}

	return nil
}

func (sp *SQLExprParser) ParseTableReference(child ISQLExprParser) select_.ISQLTableReference {
	if !IsIdentifier(sp.Kind()) {
		return nil
	}

	x := select_.NewTableReference()

	name := ParseName(child)
	x.SetName(name)

	partitionExtensionClause := sp.ParsePartitionExtensionClause(sp)
	x.SetPartitionExtensionClause(partitionExtensionClause)

	as := false
	if sp.AcceptAndNextToken(AS) {
		as = true
	}
	x.SetAs(as)

	alias := child.ParseIdentifier(child)
	x.SetAlias(alias)

	if as && alias == nil {
		panic("")
	}

	return x
}

/**
 * PARTITION (name, name...)
 */
func (sp *SQLExprParser) ParsePartitionExtensionClause(child ISQLExprParser) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(PARTITION) {
		return nil
	}

	x := select_.NewPartitionClause()

	sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
	for {
		name := ParseExpr(child)
		x.AddName(name)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	return nil
}

/**
* Join , comma
*/
func parseJoinTableReference(sp ISQLExprParser, left select_.ISQLTableReference) select_.ISQLTableReference {

	if sp.Accept(SYMB_RIGHT_PAREN) || sp.Accept(WHERE) {
		return left
	}

	if left == nil {
		panic("TableReference is nil.")
	}

	var joinType select_.SQLJoinType
	if sp.AcceptAndNextToken(SYMB_COMMA) {

		joinType = select_.COMMA

	} else if sp.AcceptAndNextToken(JOIN) {

		joinType = select_.JOIN

	} else if sp.AcceptAndNextToken(INNER) {

		sp.AcceptAndNextTokenWithError(JOIN, true)

		joinType = select_.INNER_JOIN

	} else if sp.AcceptAndNextToken(CROSS) {

		if sp.AcceptAndNextToken(JOIN) {
			joinType = select_.CROSS_JOIN
		} else if sp.AcceptAndNextToken(APPLY) {

			joinType = select_.CROSS_APPLY

		} else {
			panic("")
		}

	} else if sp.AcceptAndNextToken(LEFT) {

		joinType = select_.LEFT_JOIN
		if sp.AcceptAndNextToken(OUTER) {
			joinType = select_.LEFT_OUTER_JOIN
		}

		sp.AcceptAndNextTokenWithError(JOIN, true)

	} else if sp.AcceptAndNextToken(RIGHT) {

		joinType = select_.RIGHT_JOIN
		if sp.AcceptAndNextToken(OUTER) {
			joinType = select_.RIGHT_OUTER_JOIN
		}

		sp.AcceptAndNextTokenWithError(JOIN, true)

	} else if sp.AcceptAndNextToken(FULL) {

		joinType = select_.FULL_JOIN
		if sp.AcceptAndNextToken(OUTER) {
			joinType = select_.FULL_OUTER_JOIN
		}

		sp.AcceptAndNextTokenWithError(JOIN, true)

	} else if sp.AcceptAndNextToken(NATURAL) {

		joinType = select_.NATURAL_JOIN

		if sp.AcceptAndNextToken(INNER) {

			joinType = select_.NATURAL_INNER_JOIN
		} else if sp.AcceptAndNextToken(LEFT) {

			joinType = select_.NATURAL_LEFT_JOIN

			if sp.AcceptAndNextToken(OUTER) {
				joinType = select_.NATURAL_LEFT_OUTER_JOIN
			}

		} else if sp.AcceptAndNextToken(RIGHT) {

			joinType = select_.NATURAL_RIGHT_JOIN
			if sp.AcceptAndNextToken(OUTER) {
				joinType = select_.NATURAL_RIGHT_OUTER_JOIN
			}

		} else if sp.AcceptAndNextToken(FULL) {

			joinType = select_.NATURAL_FULL_JOIN
			if sp.AcceptAndNextToken(OUTER) {
				joinType = select_.NATURAL_FULL_OUTER_JOIN
			}
		}

		sp.AcceptAndNextTokenWithError(JOIN, true)

	} else if sp.AcceptAndNextToken(OUTER) {

		if sp.AcceptAndNextToken(APPLY) {
			joinType = select_.OUTER_APPLY

		} else {
			panic("")
		}

	} else if sp.AcceptAndNextToken(STRAIGHT_JOIN) {
		joinType = select_.STRAIGHT_JOIN
	}

	if joinType != "" {
		right := ParsePrimaryTableReference(sp)
		joinTableReference := select_.NewJoinTableReference(left, joinType, right)
		ParseJoinConditions(sp, joinTableReference)
		return parseJoinTableReference(sp, joinTableReference)
	}

	return left
}

func ParseJoinConditions(sp ISQLExprParser, x *select_.SQLJoinTableReference) {

	mark := sp.Mark()

	if sp.AcceptAndNextToken(ON) {

		condition := ParseExpr(sp)
		onCondition := select_.NewJoinOnConditionWithCondition(condition)
		x.SetCondition(onCondition)

	} else if sp.AcceptAndNextToken(USING) {

		// fix MySQL Delete Statement: USING table_references
		if !sp.Accept(SYMB_LEFT_PAREN) {
			sp.ResetWithMark(mark)
			return
		}

		usingCondition := select_.NewJoinUsingCondition()

		sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
		for {
			column := ParseExpr(sp)
			usingCondition.AddColumn(column)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

		x.SetCondition(usingCondition)
	}
}

func ParseWhereClause(sp ISQLExprParser) *select_.SQLWhereClause {
	if !sp.AcceptAndNextToken(WHERE) {
		return nil
	}

	var condition expr.ISQLExpr
	if sp.AcceptAndNextToken(CURRENT) {
		sp.AcceptAndNextTokenWithError(OF, true)

		ParseExpr(sp)

	} else {
		condition = ParseExpr(sp)
	}

	return select_.NewWhereClause(condition)
}

/**
 * { CONNECT BY [ NOCYCLE ] condition [ START WITH condition ]
| START WITH condition CONNECT BY [ NOCYCLE ] condition
}
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/SELECT.html#GUID-CFA006CA-6FF1-4972-821E-6996142A51C6
 */
func ParseHierarchicalQueryClause(sp ISQLExprParser) select_.ISQLHierarchicalQueryClause {
	if sp.AcceptAndNextToken(CONNECT) {
		sp.AcceptAndNextTokenWithError(BY, true)
		x := select_.NewHierarchicalQueryClauseConnectBy()

		nocycle := sp.AcceptAndNextToken(NOCYCLE)
		x.NoCycle = nocycle

		connectByCondition := ParseExpr(sp)
		x.SetConnectByCondition(connectByCondition)

		if sp.AcceptAndNextToken(START) {
			sp.AcceptAndNextTokenWithError(WITH, true)
			startWithCondition := ParseExpr(sp)
			x.SetStartWithCondition(startWithCondition)
		}

		return x
	}

	if sp.AcceptAndNextToken(START) {
		sp.AcceptAndNextTokenWithError(WITH, true)

		x := select_.NewHierarchicalQueryClauseStartWith()

		startWithCondition := ParseExpr(sp)
		x.SetStartWithCondition(startWithCondition)

		sp.AcceptAndNextTokenWithError(CONNECT, true)
		sp.AcceptAndNextTokenWithError(BY, true)

		nocycle := sp.AcceptAndNextToken(NOCYCLE)
		x.NoCycle = nocycle

		connectByCondition := ParseExpr(sp)
		x.SetConnectByCondition(connectByCondition)
		return x
	}
	return nil
}

func parseGroupBy(sp ISQLExprParser) select_.ISQLGroupByClause {
	var x select_.ISQLGroupByClause
	if sp.AcceptAndNextToken(GROUP) {
		sp.AcceptAndNextTokenWithError(BY, true)

		x = select_.NewGroupByHavingClause()

		// SQLSetQuantifier setQuantifier = parseSetQuantifier()
		// x.setQuantifier(setQuantifier)

		for {
			element := parseGroupByElement(sp)
			x.AddElement(element)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

		havingClause := parseHavingClause(sp)
		x.SetHaving(havingClause)

	} else if sp.Accept(HAVING) {
		x = select_.NewHavingGroupByClause()

		// x.setOrder(false)

		havingClause := parseHavingClause(sp)
		x.SetHaving(havingClause)

		sp.AcceptAndNextTokenWithError(GROUP, true)
		sp.AcceptAndNextTokenWithError(BY, true)

		// setQuantifier := parseSetQuantifier()
		// x.setQuantifier(setQuantifier)

		for {
			element := parseGroupByElement(sp)
			x.AddElement(element)

			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
	}

	return x
}

func parseGroupByElement(sp ISQLExprParser) *select_.SQLGroupByElement {
	expr := ParseExpr(sp)
	return select_.NewGroupByElement(expr)
}
func parseHavingClause(sp ISQLExprParser) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(HAVING) {
		return nil
	}
	return ParseExpr(sp)
}

func ParseOrderByClause(parser ISQLExprParser) *select_.SQLOrderByClause {
	if !parser.AcceptAndNextToken(ORDER) {
		return nil
	}

	parser.AcceptAndNextToken(BY)

	x := select_.NewOrderByClause()
	for {
		element := ParseOrderByElement(parser)
		x.AddElement(element)

		if parser.Kind() != SYMB_COMMA {
			break
		}
		NextTokenByParser(parser)
	}

	return x
}

func ParseOrderByElement(sp ISQLExprParser) *select_.SQLOrderByElement {
	key := ParseExpr(sp)

	var specification select_.SQLOrderingSpecification
	if sp.AcceptAndNextToken(ASC) {

		specification = select_.ASC

	} else if sp.AcceptAndNextToken(DESC) {

		specification = select_.DESC

	} else if sp.AcceptAndNextToken(USING) {

	}

	x := select_.NewOrderByElementWithSpecification(key, specification)

	return x
}

func ParseLimitClause(sp ISQLExprParser) select_.ISQLLimitClause {
	if sp.Accept(LIMIT) {

		return sp.ParseLimitOffsetClause(sp)

	} else if sp.Accept(OFFSET) || sp.Accept(FETCH) {

		return sp.ParseOffsetFetchClause(sp)

	}
	return nil
}

/**
* LIMIT row_count
* LIMIT offset, row_count
* LIMIT row_count OFFSET offset
*/
func (sp *SQLExprParser) ParseLimitOffsetClause(child ISQLExprParser) select_.ISQLLimitClause {
	if !child.AcceptAndNextToken(LIMIT) {
		return nil
	}

	x := select_.NewLimitOffsetClause()

	expr1 := ParseExpr(child)
	offset := false
	var offsetExpr expr.ISQLExpr
	var countExpr expr.ISQLExpr

	if child.AcceptAndNextToken(SYMB_COMMA) {
		offsetExpr = expr1
		countExpr = ParseExpr(child)

		x.SetOffsetExpr(offsetExpr)
		x.SetCountExpr(countExpr)

	} else if child.AcceptAndNextToken(OFFSET) {

		offset = true
		countExpr = expr1
		offsetExpr = ParseExpr(child)

		x.SetCountExpr(countExpr)
		x.SetOffsetExpr(offsetExpr)
	} else {

		countExpr = expr1
		x.SetCountExpr(countExpr)
	}
	x.Offset = offset

	return x
}

func (sp *SQLExprParser) ParseOffsetFetchClause(child ISQLExprParser) select_.ISQLLimitClause {
	if !sp.Accept(OFFSET) && !sp.Accept(FETCH) {
		return nil
	}
	panic(sp.UnSupport())
}

func (sp SQLExprParser) ParseLockClause(child ISQLExprParser) select_.ISQLLockClause {
	return nil
}

/**
* FOR { UPDATE | NO KEY UPDATE | SHARE | KEY SHARE } [ OF table_name [, ...] ] [ NOWAIT | SKIP LOCKED| WAIT integer  ] [...]
*/
func (sp *SQLExprParser) ParseForUpdate(child ISQLExprParser) select_.ISQLLockClause {
	return nil
}

func (sp *SQLExprParser) ParseForUpdateRest(child ISQLExprParser, x *select_.AbstractSQLLockForClause) {
	if sp.AcceptAndNextToken(OF) {
		for {
			x.AddTable(ParseName(child))
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
	}

	if sp.AcceptAndNextToken(NOWAIT) {

		x.SetWaitExpr(select_.NewLockForNoWaitExpr())

	} else if sp.AcceptAndNextToken(SKIP) {

		sp.AcceptAndNextTokenWithError(LOCKED, true)
		x.SetWaitExpr(select_.NewLockForSkipLockedExpr())
	}
}

/**
 * LOCK IN SHARE MODE
 */
func (sp *SQLExprParser) ParseLockInShareModeClause(child ISQLExprParser) select_.ISQLLockClause {
	if !sp.AcceptAndNextToken(LOCK) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(IN, true)
	sp.AcceptAndNextTokenWithError(SHARE, true)
	sp.AcceptAndNextTokenWithError(MODE, true)
	return select_.NewLockInShareModeClause()
}

/**
 * { RETURN | RETURNING } expr [, expr ]...INTO data_item [, data_item ]...
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DELETE.html#GUID-156845A5-B626-412B-9F95-8869B988ABD7
 */
func ParseIReturningClause(sp ISQLExprParser) select_.ISQLReturningClause {
	if !sp.Accept(RETURN) && !sp.Accept(RETURNING) {
		return nil
	}

	var x select_.ISQLReturningClause = nil
	if sp.AcceptAndNextToken(RETURN) {

		for {
			ParseExpr(sp)
			// x.addIntoItem(intoItem)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

		sp.AcceptAndNextTokenWithError(INTO, true)
		for {
			ParseExpr(sp)
			// x.addIntoItem(intoItem)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

	} else if sp.AcceptAndNextToken(RETURNING) {

		for {
			ParseExpr(sp)
			// x.addIntoItem(intoItem)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

		sp.AcceptAndNextTokenWithError(INTO, true)
		for {
			ParseExpr(sp)
			// x.addIntoItem(intoItem)
			if !sp.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

	}

	return x
}
