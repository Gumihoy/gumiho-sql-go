package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/index"
	view2 "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/function"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/procedure"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/role"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/schema"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/sequence"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/trigger"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type ISQLStatementParser interface {
	ISQLParser
	ExprParser() ISQLExprParser
	SetExprParser(exprParser ISQLExprParser)
}

type SQLStatementParser struct {
	*SQLParser
	exprParser ISQLExprParser
}

func (x *SQLStatementParser) ExprParser() ISQLExprParser {
	return x.exprParser
}

func (x *SQLStatementParser) SetExprParser(exprParser ISQLExprParser) {
	x.exprParser = exprParser
}

func NewStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLStatementParser {
	return NewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewStatementParserByExprParser(exprParser ISQLExprParser) *SQLStatementParser {
	var x SQLStatementParser
	x.SQLParser = NewParserByLexer(exprParser.Lexer(), exprParser.DBType(), exprParser.Config())
	x.exprParser = exprParser
	return &x
}

func ParseStatements(parser ISQLStatementParser) []statement.ISQLStatement {
	return ParseStatementsWithParent(parser, nil)
}

func ParseStatementsWithParent(x ISQLStatementParser, parent ast.ISQLObject) []statement.ISQLStatement {
	NextTokenByParser(x)

	var comments []ast.ISQLComment
	var stmts []statement.ISQLStatement

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

		comments = x.ExprParser().ParseComments(x.ExprParser())
		stmt := ParseStatement(x.ExprParser())
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

	if !x.Accept(EOF) {
		panic(x.UnSupport())
	}

	return stmts
}

// ------------------------------------------------------------- DML -------------------------------------------------------------

// ------------------------ Select ------------------------

type ISQLSelectStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLSelectStatementParser struct {
	*SQLStatementParser
}

func NewSelectStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSelectStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSelectStatementParserByExprParser(exprParser ISQLExprParser) *SQLSelectStatementParser {
	x := new(SQLSelectStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

func (x *SQLSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != SELECT {
		return nil
	}
	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(x.dbType, query)
}

// ------------------------ Delete ------------------------

type ISQLDeleteStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLDeleteStatementParser struct {
	*SQLStatementParser
}

func NewDeleteStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLDeleteStatementParser {
	return NewDeleteStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDeleteStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDeleteStatementParserByExprParser(exprParser ISQLExprParser) *SQLDeleteStatementParser {
	return &SQLDeleteStatementParser{NewStatementParserByExprParser(exprParser)}
}

func (x *SQLDeleteStatementParser) Parse() statement.ISQLStatement {
	if x.Accept(DELETE) {
		return nil
	}

	return nil
}

// ------------------------ Insert ------------------------
type ISQLInsertStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLInsertStatementParser struct {
	*SQLStatementParser
}

func (x *SQLInsertStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != SELECT {
		return nil
	}
	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(x.dbType, query)
}

func NewInsertStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLInsertStatementParser {
	return NewInsertStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewInsertStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLInsertStatementParser {
	return NewInsertStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewInsertStatementParserByExprParser(exprParser ISQLExprParser) *SQLInsertStatementParser {
	return &SQLInsertStatementParser{NewStatementParserByExprParser(exprParser)}
}

// ------------------------ Update ------------------------

type ISQLUpdateStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLUpdateStatementParser struct {
	*SQLStatementParser
}

func (x *SQLUpdateStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != SELECT {
		return nil
	}
	return nil
}

func NewUpdateStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLUpdateStatementParser {
	return NewUpdateStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewUpdateStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLUpdateStatementParser {
	return NewUpdateStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewUpdateStatementParserByExprParser(exprParser ISQLExprParser) *SQLUpdateStatementParser {
	return &SQLUpdateStatementParser{NewStatementParserByExprParser(exprParser)}
}

// ------------------------------------------------------------- DDL -------------------------------------------------------------

type ISQLDDLStatementParser interface {
	ISQLStatementParser
	ParseAlter() statement.ISQLStatement
	ParseCreate() statement.ISQLStatement
	ParseDrop() statement.ISQLStatement
}

type abstractSQLDDLStatementParser struct {
	*SQLStatementParser
}

func NewAbstractSQLDDLStatementParserByExprParser(exprParser ISQLExprParser) *abstractSQLDDLStatementParser {
	x := new(abstractSQLDDLStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

func (x *abstractSQLDDLStatementParser) ParseAlter() statement.ISQLStatement {
	panic("implement me")
}

func (x *abstractSQLDDLStatementParser) ParseCreate() statement.ISQLStatement {
	panic("implement me")
}

func (x *abstractSQLDDLStatementParser) ParseDrop() statement.ISQLStatement {
	panic("implement me")
}

// ------------------------ COMMENT ------------------------
type ISQLCommentStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}
type SQLCommentOnColumnStatementParser struct {
	*SQLStatementParser
}
func NewCommentOnColumnStatementParserByExprParser(exprParser ISQLExprParser) *SQLCommentOnColumnStatementParser {
	x := new(SQLCommentOnColumnStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

func (x *SQLCommentOnColumnStatementParser) Parse() statement.ISQLStatement {
	panic("implement me")
}


type SQLCommentOnMaterializedViewStatementParser struct {
	*SQLStatementParser
}
func NewCommentOnMaterializedViewStatementParserByExprParser(exprParser ISQLExprParser) *SQLCommentOnMaterializedViewStatementParser {
	x := new(SQLCommentOnMaterializedViewStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

func (x *SQLCommentOnMaterializedViewStatementParser) Parse() statement.ISQLStatement {
	panic("implement me")
}


type SQLCommentOnTableStatementParser struct {
	*SQLStatementParser
}
func NewCommentOnTableStatementParserByExprParser(exprParser ISQLExprParser) *SQLCommentOnTableStatementParser {
	x := new(SQLCommentOnTableStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (x *SQLCommentOnTableStatementParser) Parse() statement.ISQLStatement {
	panic("implement me")
}

// ------------------------ Database ------------------------

type ISQLDatabaseStatementParser interface {
	ISQLDDLStatementParser
}

type SQLDatabaseStatementParser struct {
	*SQLStatementParser
}

func NewDatabaseStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLDatabaseStatementParser {
	return NewDatabaseStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDatabaseStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLDatabaseStatementParser {
	return NewDatabaseStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDatabaseStatementParserByExprParser(exprParser ISQLExprParser) *SQLDatabaseStatementParser {
	x := new(SQLDatabaseStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

func (sp *SQLDatabaseStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLDatabaseStatementParser) ParseAlterDatabaseAction() expr.ISQLExpr {
	return nil
}
func (sp *SQLDatabaseStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLDatabaseStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ Function ------------------------

type ISQLFunctionStatementParser interface {
	ISQLDDLStatementParser
}

type SQLFunctionStatementParser struct {
	*SQLStatementParser
}

func NewFunctionStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLFunctionStatementParser {
	return NewFunctionStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewFunctionStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLFunctionStatementParser {
	return NewFunctionStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewFunctionStatementParserByExprParser(exprParser ISQLExprParser) *SQLFunctionStatementParser {
	x := new(SQLFunctionStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLFunctionStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLFunctionStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(CREATE) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(FUNCTION, true)

	x := function.NewCreateFunctionStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(SYMB_LEFT_PAREN, true)
	for {
		parameter := sp.ExprParser().ParseParameterDeclaration(sp.ExprParser())
		x.AddParameter(parameter)
		if !sp.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	sp.AcceptAndNextTokenWithError(RETURNS, true)
	ParseDataType(sp.ExprParser())


	return x
}

func (sp *SQLFunctionStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(FUNCTION, true)
	x := function.NewDropFunctionStatement(sp.DBType())
	ifExists := ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	name := ParseName(sp.ExprParser())
	x.SetName(name)
	return x
}

// ------------------------ Index ------------------------

type ISQLIndexStatementParser interface {
	ISQLDDLStatementParser
}

type SQLIndexStatementParser struct {
	*SQLStatementParser
}

func NewIndexStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLIndexStatementParser {
	return NewIndexStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewIndexStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLIndexStatementParser {
	return NewIndexStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewIndexStatementParserByExprParser(exprParser ISQLExprParser) *SQLIndexStatementParser {
	x := new(SQLIndexStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLIndexStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLIndexStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLIndexStatementParser) ParseIndexColumn() *index.SQLIndexColumn {
	x := index.NewViewColumn()
	expr := ParseExpr(sp.ExprParser())
	x.SetExpr(expr)

	return x
}
func (sp *SQLIndexStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ Package ------------------------

type ISQLPackageStatementParser interface {
	ISQLDDLStatementParser
}

type SQLPackageStatementParser struct {
	*SQLStatementParser
}

func NewPackageStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLPackageStatementParser {
	return NewPackageStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewPackageStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLPackageStatementParser {
	return NewPackageStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewPackageStatementParserByExprParser(exprParser ISQLExprParser) *SQLPackageStatementParser {
	x := new(SQLPackageStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (x *SQLPackageStatementParser) ParseAlter() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *SQLPackageStatementParser) ParseCreate() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *SQLPackageStatementParser) ParseDrop() statement.ISQLStatement {
	panic(x.UnSupport())
}

// ------------------------ Package Body ------------------------

type ISQLPackageBodyStatementParser interface {
	ISQLDDLStatementParser
}

type SQLPackageBodyStatementParser struct {
	*SQLStatementParser
}

func NewPackageBodyStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLPackageBodyStatementParser {
	return NewPackageBodyStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewPackageBodyStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLPackageBodyStatementParser {
	return NewPackageBodyStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewPackageBodyStatementParserByExprParser(exprParser ISQLExprParser) *SQLPackageBodyStatementParser {
	x := new(SQLPackageBodyStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (x *SQLPackageBodyStatementParser) ParseAlter() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *SQLPackageBodyStatementParser) ParseCreate() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *SQLPackageBodyStatementParser) ParseDrop() statement.ISQLStatement {
	panic(x.UnSupport())
}

// ------------------------ Procedure ------------------------

type ISQLProcedureStatementParser interface {
	ISQLDDLStatementParser
}

type SQLProcedureStatementParser struct {
	*SQLStatementParser
}

func NewProcedureStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLProcedureStatementParser {
	return NewProcedureStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewProcedureStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLProcedureStatementParser {
	return NewProcedureStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewProcedureStatementParserByExprParser(exprParser ISQLExprParser) *SQLProcedureStatementParser {
	x := new(SQLProcedureStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (x *SQLProcedureStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(ALTER) {
		return nil
	}

	return nil
}
func (x *SQLProcedureStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(CREATE) {
		return nil
	}

	return nil
}
func (sp *SQLProcedureStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(PROCEDURE, true)

	x := procedure.NewDropProcedureStatement(sp.DBType())

	ifExists := ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	name := ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Role ------------------------

type ISQLRoleStatementParser interface {
	ISQLDDLStatementParser
}

type SQLRoleStatementParser struct {
	*SQLStatementParser
}

func NewRoleStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLRoleStatementParser {
	return NewRoleStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewRoleStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLRoleStatementParser {
	return NewRoleStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewRoleStatementParserByExprParser(exprParser ISQLExprParser) *SQLRoleStatementParser {
	x := new(SQLRoleStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLRoleStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}

/**
 * CREATE ROLE <role name> [ WITH ADMIN <grantor> ]
 */
func (sp *SQLRoleStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(CREATE) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(ROLE, true)

	x := role.NewCreateRoleStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.AddName(name)

	if sp.AcceptAndNextToken(WITH) {
		sp.AcceptAndNextTokenWithError(ADMIN, true)
		if sp.AcceptAndNextToken(CURRENT_USER) {
			x.WithAdminGrantor = role.CURRENT_USER

		} else if sp.AcceptAndNextToken(CURRENT_ROLE) {
			x.WithAdminGrantor = role.CURRENT_ROLE

		} else {
			panic(sp.UnSupport())
		}
	}

	return x
}

/**
 * DROP ROLE <role name>
 */
func (sp *SQLRoleStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(ROLE, true)

	x := role.NewDropRoleStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.AddName(name)

	return x
}

// ------------------------ Schema ------------------------

type ISQLSchemaStatementParser interface {
	ISQLDDLStatementParser
}

type SQLSchemaStatementParser struct {
	*SQLStatementParser
}

func NewSchemaStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLSchemaStatementParser {
	return NewSchemaStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSchemaStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLSchemaStatementParser {
	return NewSchemaStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSchemaStatementParserByExprParser(exprParser ISQLExprParser) *SQLSchemaStatementParser {
	x := new(SQLSchemaStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

func (x *SQLSchemaStatementParser) ParseAlter() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (sp *SQLSchemaStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(CREATE) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(SCHEMA, true)

	x := schema.NewCreateSchemaStatement(sp.DBType())

	name := ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}
func (sp *SQLSchemaStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(SCHEMA, true)

	x := schema.NewDropSchemaStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.SetName(name)

	behavior := ParseDropBehavior(sp)
	x.Behavior = behavior
	return x
}

// ------------------------ Sequence ------------------------

type ISQLSequenceStatementParser interface {
	ISQLDDLStatementParser
}

type SQLSequenceStatementParser struct {
	*SQLStatementParser
}

func NewSequenceStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLSequenceStatementParser {
	return NewSequenceStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSequenceStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLSequenceStatementParser {
	return NewSequenceStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSequenceStatementParserByExprParser(exprParser ISQLExprParser) *SQLSequenceStatementParser {
	x := new(SQLSequenceStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (x *SQLSequenceStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(ALTER) {
		return nil
	}

	return nil
}
func (sp *SQLSequenceStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(CREATE) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(SEQUENCE, true)

	x := sequence.NewCreateSequenceStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}
func (sp *SQLSequenceStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(SEQUENCE, true)

	x := sequence.NewDropSequenceStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.SetName(name)

	dropBehavior := ParseDropBehavior(sp)
	x.DropBehavior = dropBehavior
	return x
}

// ------------------------ Server ------------------------
type ISQLServerStatementParser interface {
	ISQLDDLStatementParser
}

type SQLServerStatementParser struct {
	*SQLStatementParser
}

func NewServerStatementParserByExprParser(exprParser ISQLExprParser) *SQLServerStatementParser {
	x := new(SQLServerStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLServerStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLServerStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLServerStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ Synonym ------------------------

type ISQLSynonymStatementParser interface {
	ISQLDDLStatementParser
}

type SQLSynonymStatementParser struct {
	*SQLStatementParser
}

func NewSynonymStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLSynonymStatementParser {
	return NewSynonymStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSynonymStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLSynonymStatementParser {
	return NewSynonymStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSynonymStatementParserByExprParser(exprParser ISQLExprParser) *SQLSynonymStatementParser {
	x := new(SQLSynonymStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLSynonymStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLSynonymStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLSynonymStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ Table ------------------------

type ISQLTableStatementParser interface {
	ISQLDDLStatementParser
}

type SQLTableStatementParser struct {
	*SQLStatementParser
}

func NewTableStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLTableStatementParser {
	return NewTableStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTableStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLTableStatementParser {
	return NewTableStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTableStatementParserByExprParser(exprParser ISQLExprParser) *SQLTableStatementParser {
	x := new(SQLTableStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

/**
 * ALTER TABLE <table name> <alter table action>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20table%20statement
 */
func (sp *SQLTableStatementParser) ParseAlter() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(ALTER, true)
	sp.AcceptAndNextTokenWithError(TABLE, true)

	x := table.NewAlterTableStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.SetName(name)

	action := sp.ExprParser().ParseAlterTableAction(sp.ExprParser())
	x.AddAction(action)

	return x
}

/**
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20definition
 */
func (sp *SQLTableStatementParser) ParseCreate() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(CREATE, true)

	x := table.NewCreateTableStatement(sp.DBType())

	if sp.AcceptAndNextToken(GLOBAL) {

		sp.AcceptAndNextTokenWithError(TEMPORARY, true)
		x.TableScope = table.GLOBAL_TEMPORARY

	} else if sp.AcceptAndNextToken(LOCAL) {

		sp.AcceptAndNextTokenWithError(TEMPORARY, true)
		x.TableScope = table.LOCAL_TEMPORARY
	}

	sp.AcceptAndNextTokenWithError(TABLE, true)

	name := ParseName(sp.ExprParser())
	x.SetName(name)

	// (TableElements)
	elements := ParseTableElements(sp.ExprParser())
	x.AddElements(elements)

	return x
}
func (sp *SQLTableStatementParser) ParseDrop() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(DROP, true)
	sp.AcceptAndNextTokenWithError(TABLE, true)

	x := table.NewDropTableStatement(sp.DBType())
	return x
}

// ------------------------ Trigger ------------------------

type ISQLTriggerStatementParser interface {
	ISQLDDLStatementParser
}

type SQLTriggerStatementParser struct {
	*SQLStatementParser
}

func NewTriggerStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLTriggerStatementParser {
	return NewTriggerStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTriggerStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLTriggerStatementParser {
	return NewTriggerStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTriggerStatementParserByExprParser(exprParser ISQLExprParser) *SQLTriggerStatementParser {
	x := new(SQLTriggerStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (x *SQLTriggerStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	return nil
}
func (x *SQLTriggerStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	return nil
}
func (sp *SQLTriggerStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(TRIGGER, true)

	x := trigger.NewDropTriggerStatement(sp.DBType())
	name := ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Type ------------------------

type ISQLTypeStatementParser interface {
	ISQLDDLStatementParser
}

type SQLTypeStatementParser struct {
	*SQLStatementParser
}

func NewTypeStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLTypeStatementParser {
	return NewTypeStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTypeStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLTypeStatementParser {
	return NewTypeStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTypeStatementParserByExprParser(exprParser ISQLExprParser) *SQLTypeStatementParser {
	x := new(SQLTypeStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLTypeStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLTypeStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLTypeStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ Type Body ------------------------

type ISQLTypeBodyStatementParser interface {
	ISQLDDLStatementParser
}

type SQLTypeBodyStatementParser struct {
	*SQLStatementParser
}

func NewTypeBodyStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLTypeBodyStatementParser {
	return NewTypeBodyStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTypeBodyStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLTypeBodyStatementParser {
	return NewTypeBodyStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTypeBodyStatementParserByExprParser(exprParser ISQLExprParser) *SQLTypeBodyStatementParser {
	x := new(SQLTypeBodyStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLTypeBodyStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLTypeBodyStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *SQLTypeBodyStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ User ------------------------

type ISQLUserStatementParser interface {
	ISQLDDLStatementParser
}

type SQLUserStatementParser struct {
	*SQLStatementParser
}

func NewUserStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLUserStatementParser {
	return NewUserStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewUserStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLUserStatementParser {
	return NewUserStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewUserStatementParserByExprParser(exprParser ISQLExprParser) *SQLUserStatementParser {
	x := new(SQLUserStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (x *SQLUserStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	return nil
}
func (x *SQLUserStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	return nil
}
func (x *SQLUserStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	return nil
}

// ------------------------ View ------------------------

type ISQLViewStatementParser interface {
	ISQLDDLStatementParser
}

type SQLViewStatementParser struct {
	*SQLStatementParser
}

func NewViewStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLViewStatementParser {
	return NewViewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewViewStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLViewStatementParser {
	return NewViewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewViewStatementParserByExprParser(exprParser ISQLExprParser) *SQLViewStatementParser {
	x := new(SQLViewStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

/**
 *
 */
func (sp *SQLViewStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(ALTER) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(VIEW, true)

	x := view.NewAlterViewStatement(sp.DBType())

	return x
}

/**
 *
 */
func (sp *SQLViewStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(CREATE) {
		return nil
	}
	x := view.NewCreateViewStatement(sp.DBType())

	recursive := sp.AcceptAndNextToken(RECURSIVE)
	x.Recursive = recursive

	sp.AcceptAndNextTokenWithError(VIEW, true)

	// OF
	if sp.AcceptAndNextToken(OF) {

	}

	elements := ParseViewElements(sp.ExprParser())
	x.AddElements(elements)

	sp.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)

	sp.AcceptAndNextTokenWithError(AS, true)
	subQuery := ParseSelectQuery(sp.ExprParser())
	x.SetSubQuery(subQuery)

	return x
}

func (sp *SQLViewStatementParser) ParseViewElement() view2.ISQLViewElement {
	if sp.ExprParser().IsParseTableConstraint() {
		return sp.ExprParser().ParseTableConstraint(sp.ExprParser())

	} else if IsIdentifier(sp.Kind()) {

	}
	return nil
}

/**
 * DROP VIEW <table name> <drop behavior>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20view%20statement
 */
func (sp *SQLViewStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(VIEW, true)

	x := view.NewDropViewStatement(sp.DBType())

	name := ParseName(sp.ExprParser())
	x.AddName(name)

	behavior := ParseDropBehavior(sp)
	x.Behavior = behavior

	return x
}

func ParseDropBehavior(sp ISQLStatementParser) statement.SQLDropBehavior {
	if sp.AcceptAndNextToken(RESTRICT) {
		return statement.RESTRICT

	} else if sp.AcceptAndNextToken(CASCADE) {
		return statement.CASCADE
	}
	return ""
}

// --------------------------- Set Statement ---------------------------
type ISQLSetStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}
type SQLSetVariableAssignmentStatementParser struct {
	*SQLStatementParser
}

func NewSetVariableAssignmentStatementParserByExprParser(exprParser ISQLExprParser) *SQLSetVariableAssignmentStatementParser {
	x := new(SQLSetVariableAssignmentStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLSetVariableAssignmentStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

/**
 *
 */
type SQLSetCharacterSetStatementParser struct {
	*SQLStatementParser
}

func NewSetCharacterSetStatementParserByExprParser(exprParser ISQLExprParser) *SQLSetCharacterSetStatementParser {
	x := new(SQLSetCharacterSetStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLSetCharacterSetStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

/**
 *
 */
type SQLSeCharsetStatementParser struct {
	*SQLStatementParser
}

func NewSetCharsetStatementParserByExprParser(exprParser ISQLExprParser) *SQLSeCharsetStatementParser {
	x := new(SQLSeCharsetStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}

func (sp *SQLSeCharsetStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

/**
 *
 */
type SQLSetNamesStatementParser struct {
	*SQLStatementParser
}

func NewSetNamesStatementParserByExprParser(exprParser ISQLExprParser) *SQLSetNamesStatementParser {
	x := new(SQLSetNamesStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLSetNamesStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// --------------------------- Show Statement ---------------------------
type ISQLShowStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLShowCreateDatabaseParser struct {
	*SQLStatementParser
}

func NewShowCreateDatabaseBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowCreateDatabaseParser {
	return NewShowCreateDatabaseByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateDatabaseByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowCreateDatabaseParser {
	return NewShowCreateDatabaseByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateDatabaseByExprParser(exprParser ISQLExprParser) *SQLShowCreateDatabaseParser {
	x := new(SQLShowCreateDatabaseParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowCreateDatabaseParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLShowCreateEventParser struct {
	*SQLStatementParser
}

func NewShowCreateEventBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowCreateEventParser {
	return NewShowCreateEventByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateEventByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowCreateEventParser {
	return NewShowCreateEventByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateEventByExprParser(exprParser ISQLExprParser) *SQLShowCreateEventParser {
	x := new(SQLShowCreateEventParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowCreateEventParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLShowCreateFunctionParser struct {
	*SQLStatementParser
}

func NewShowCreateFunctionBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowCreateFunctionParser {
	return NewShowCreateFunctionByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateFunctionByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowCreateFunctionParser {
	return NewShowCreateFunctionByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateFunctionByExprParser(exprParser ISQLExprParser) *SQLShowCreateFunctionParser {
	x := new(SQLShowCreateFunctionParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowCreateFunctionParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLShowCreateProcedureParser struct {
	*SQLStatementParser
}

func NewShowCreateProcedureBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowCreateProcedureParser {
	return NewShowCreateProcedureByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateProcedureByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowCreateProcedureParser {
	return NewShowCreateProcedureByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateProcedureByExprParser(exprParser ISQLExprParser) *SQLShowCreateProcedureParser {
	x := new(SQLShowCreateProcedureParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowCreateProcedureParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLShowCreateTableParser struct {
	*SQLStatementParser
}

func NewShowCreateTableBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowCreateTableParser {
	return NewShowCreateTableByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateTableByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowCreateTableParser {
	return NewShowCreateTableByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateTableByExprParser(exprParser ISQLExprParser) *SQLShowCreateTableParser {
	x := new(SQLShowCreateTableParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowCreateTableParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLShowCreateTriggerParser struct {
	*SQLStatementParser
}

func NewShowCreateTriggerBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowCreateTriggerParser {
	return NewShowCreateTriggerByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateTriggerByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowCreateTriggerParser {
	return NewShowCreateTriggerByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateTriggerByExprParser(exprParser ISQLExprParser) *SQLShowCreateTriggerParser {
	x := new(SQLShowCreateTriggerParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowCreateTriggerParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLShowCreateViewParser struct {
	*SQLStatementParser
}

func NewShowCreateViewBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowCreateViewParser {
	return NewShowCreateViewByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateViewByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowCreateViewParser {
	return NewShowCreateViewByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateViewByExprParser(exprParser ISQLExprParser) *SQLShowCreateViewParser {
	x := new(SQLShowCreateViewParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowCreateViewParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLShowDatabasesParser struct {
	*SQLStatementParser
}

func NewShowDatabasesBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLShowDatabasesParser {
	return NewShowDatabasesByLexer(NewLexer(sql), dbType, config)
}

func NewShowDatabasesByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLShowDatabasesParser {
	return NewShowDatabasesByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowDatabasesByExprParser(exprParser ISQLExprParser) *SQLShowDatabasesParser {
	x := new(SQLShowDatabasesParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLShowDatabasesParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}


// --------------------------- EXPLAIN Statement Start ---------------------------
type ISQLExplainStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLDescStatementParser struct {
	*SQLStatementParser
}

func NewDescStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLDescStatementParser {
	return NewDescStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDescStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLDescStatementParser {
	return NewDescStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewDescStatementParserByExprParser(exprParser ISQLExprParser) *SQLDescStatementParser {
	x := new(SQLDescStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLDescStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}



type SQLDescribeStatementParser struct {
	*SQLStatementParser
}

func NewDescribeStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLDescribeStatementParser {
	return NewDescribeStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDescribeStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLDescribeStatementParser {
	return NewDescribeStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewDescribeStatementParserByExprParser(exprParser ISQLExprParser) *SQLDescribeStatementParser {
	x := new(SQLDescribeStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLDescribeStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}

type SQLExplainStatementParser struct {
	*SQLStatementParser
}

func NewExplainStatementParserBySQL(sql string, dbType db.Type, config *SQLParseConfig) *SQLExplainStatementParser {
	return NewExplainStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewExplainStatementParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLExplainStatementParser {
	return NewExplainStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewExplainStatementParserByExprParser(exprParser ISQLExprParser) *SQLExplainStatementParser {
	x := new(SQLExplainStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLExplainStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}
// --------------------------- EXPLAIN Statement End ---------------------------

// --------------------------- HELP Statement Start ---------------------------
type ISQLHelpStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}
type SQLHelpStatementParser struct {
	*SQLStatementParser
}
func NewHelpStatementParserByExprParser(exprParser ISQLExprParser) *SQLHelpStatementParser {
	x := new(SQLHelpStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLHelpStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}
// --------------------------- HELP Statement End ---------------------------

// --------------------------- USE Statement Start ---------------------------
type ISQLUseStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}
type SQLUseStatementParser struct {
	*SQLStatementParser
}
func NewUseStatementParserByExprParser(exprParser ISQLExprParser) *SQLUseStatementParser {
	x := new(SQLUseStatementParser)
	x.SQLStatementParser = NewStatementParserByExprParser(exprParser)
	return x
}
func (sp *SQLUseStatementParser) Parse() statement.ISQLStatement {
	panic(sp.UnSupport())
}
// --------------------------- USE Statement End ---------------------------