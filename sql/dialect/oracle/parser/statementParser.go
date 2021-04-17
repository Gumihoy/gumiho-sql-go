package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	index2 "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/index"
	exprTable "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	exprUser "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/user"
	exprView "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/comment"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/database"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/index"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/role"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/schema"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/sequence"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/synonym"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/trigger"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/type"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/typebody"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/user"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type OracleStatementParser struct {
	*parser.SQLStatementParser
}

func NewStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleStatementParser {
	return NewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleStatementParser {
	var x OracleStatementParser
	x.SQLStatementParser = parser.NewStatementParserByExprParser(exprParser)
	return &x
}

// ------------------------------------------------------------- DML -------------------------------------------------------------
// ------------------------ Select ------------------------

type OracleSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func NewSelectStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleSelectStatementParser {
	x := new(OracleSelectStatementParser)
	x.SQLSelectStatementParser = parser.NewSelectStatementParserByExprParser(exprParser)
	return x
}

func (x *OracleSelectStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.WITH) && !x.Accept(parser.SELECT) {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(x.DBType(), query)
}

// ------------------------ Delete ------------------------

type OracleDeleteStatementParser struct {
	*parser.SQLDeleteStatementParser
}

func NewDeleteStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleDeleteStatementParser {
	return NewDeleteStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDeleteStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDeleteStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleDeleteStatementParser {
	x := new(OracleDeleteStatementParser)
	x.SQLDeleteStatementParser = parser.NewDeleteStatementParserByExprParser(exprParser)
	return x
}

func (x *OracleDeleteStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.DELETE) {
		return nil
	}
	return nil
}

// ------------------------ Insert ------------------------

type OracleInsertStatementParser struct {
	*parser.SQLInsertStatementParser
}

func (x *OracleInsertStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.INSERT) {
		return nil
	}
	return nil
}

func NewInsertStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleInsertStatementParser {
	return NewInsertStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewInsertStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleInsertStatementParser {
	return NewInsertStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewInsertStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleInsertStatementParser {
	x := new(OracleInsertStatementParser)
	x.SQLInsertStatementParser = parser.NewInsertStatementParserByExprParser(exprParser)
	return x
}

// ------------------------ Update ------------------------
type OracleUpdateStatementParser struct {
	*parser.SQLUpdateStatementParser
}

func (x *OracleUpdateStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.UPDATE) {
		return nil
	}
	return nil
}

func NewUpdateStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleUpdateStatementParser {
	return NewUpdateStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewUpdateStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleUpdateStatementParser {
	return NewUpdateStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewUpdateStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleUpdateStatementParser {
	x := new(OracleUpdateStatementParser)
	x.SQLUpdateStatementParser = parser.NewUpdateStatementParserByExprParser(exprParser)
	return x
}

// ------------------------ COMMENT ------------------------

type OracleCommentOnColumnStatementParser struct {
	*parser.SQLCommentOnColumnStatementParser
}

func NewCommentOnColumnStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleCommentOnColumnStatementParser {
	x := new(OracleCommentOnColumnStatementParser)
	x.SQLCommentOnColumnStatementParser = parser.NewCommentOnColumnStatementParserByExprParser(exprParser)
	return x
}

func (sp *OracleCommentOnColumnStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.COMMENT, true)
	sp.AcceptAndNextTokenWithError(parser.ON, true)
	sp.AcceptAndNextTokenWithError(parser.COLUMN, true)

	x := comment.NewCommentOnColumnStatement(sp.DBType())
	name := parser.ParsePrimaryExpr(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.IS, true)
	comment := parser.ParsePrimaryExpr(sp.ExprParser())
	x.SetComment(comment)

	return x
}

type OracleCommentOnMaterializedViewStatementParser struct {
	*parser.SQLCommentOnMaterializedViewStatementParser
}

func NewCommentOnMaterializedViewStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleCommentOnMaterializedViewStatementParser {
	x := new(OracleCommentOnMaterializedViewStatementParser)
	x.SQLCommentOnMaterializedViewStatementParser = parser.NewCommentOnMaterializedViewStatementParserByExprParser(exprParser)
	return x
}

func (sp *OracleCommentOnMaterializedViewStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.COMMENT, true)
	sp.AcceptAndNextTokenWithError(parser.ON, true)
	sp.AcceptAndNextTokenWithError(parser.COLUMN, true)

	x := comment.NewCommentOnMaterializedViewStatement(sp.DBType())
	name := parser.ParsePrimaryExpr(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.IS, true)
	comment := parser.ParsePrimaryExpr(sp.ExprParser())
	x.SetComment(comment)

	return x
}

type OracleCommentOnTableStatementParser struct {
	*parser.SQLCommentOnTableStatementParser
}

func NewCommentOnTableStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleCommentOnTableStatementParser {
	x := new(OracleCommentOnTableStatementParser)
	x.SQLCommentOnTableStatementParser = parser.NewCommentOnTableStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleCommentOnTableStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.COMMENT, true)
	sp.AcceptAndNextTokenWithError(parser.ON, true)
	sp.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := comment.NewCommentOnTableStatement(sp.DBType())
	name := parser.ParsePrimaryExpr(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.IS, true)
	comment := parser.ParsePrimaryExpr(sp.ExprParser())
	x.SetComment(comment)

	return x
}

// ------------------------------------------------------------- DDL -------------------------------------------------------------

// ------------------------ Database ------------------------
type OracleDatabaseStatementParser struct {
	*parser.SQLDatabaseStatementParser
}

func NewDatabaseStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleDatabaseStatementParser {
	return NewDatabaseStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDatabaseStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleDatabaseStatementParser {
	return NewDatabaseStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDatabaseStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleDatabaseStatementParser {
	x := new(OracleDatabaseStatementParser)
	x.SQLDatabaseStatementParser = parser.NewDatabaseStatementParserByExprParser(exprParser)
	return x
}

func (p *OracleDatabaseStatementParser) ParseAlter() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.ALTER, true)
	p.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewAlterDatabaseStatement(p.DBType())

	return x
}

//
func (p *OracleDatabaseStatementParser) ParseCreate() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.CREATE, true)
	p.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewCreateDatabaseStatement(p.DBType())

	ifNotExists := parser.ParseIfNotExists(p.ExprParser())
	x.IfNotExists = ifNotExists

	name := parser.ParseName(p.ExprParser())
	x.SetName(name)

	return x

}
func (sp *OracleDatabaseStatementParser) ParseDrop() statement.ISQLStatement {

	sp.AcceptAndNextTokenWithError(parser.DROP, true)
	sp.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewDropDatabaseSStatement(sp.DBType())

	return x
}

// ------------------------ Function ------------------------

type OracleFunctionStatementParser struct {
	*parser.SQLFunctionStatementParser
}

func NewFunctionStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleFunctionStatementParser {
	return NewFunctionStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewFunctionStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleFunctionStatementParser {
	return NewFunctionStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewFunctionStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleFunctionStatementParser {
	x := new(OracleFunctionStatementParser)
	x.SQLFunctionStatementParser = parser.NewFunctionStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleFunctionStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (x *OracleFunctionStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	return nil
}

// func (sp *OracleFunctionStatementParser) ParseDrop() statement.ISQLStatement {
// 	if !sp.AcceptAndNextToken(parser.DROP) {
// 		return nil
// 	}
//
// 	sp.AcceptAndNextTokenWithError(parser.FUNCTION, true)
// 	x := function.NewDropFunctionStatement(sp.DBType())
// 	ifExists := ParseIfExists(sp.ExprParser())
// 	x.IfExists = ifExists
//
// 	name := ParseName(sp.ExprParser())
// 	x.SetName(name)
// 	return x
// }

// ------------------------ Index ------------------------

type OracleIndexStatementParser struct {
	*parser.SQLIndexStatementParser
}

func NewIndexStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleIndexStatementParser {
	return NewIndexStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewIndexStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleIndexStatementParser {
	return NewIndexStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewIndexStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleIndexStatementParser {
	x := new(OracleIndexStatementParser)
	x.SQLIndexStatementParser = parser.NewIndexStatementParserByExprParser(exprParser)
	return x
}
func (x *OracleIndexStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (sp *OracleIndexStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	x := index.NewCreateIndexStatement(sp.DBType())

	if sp.AcceptAndNextToken(parser.UNIQUE) {

	} else if sp.AcceptAndNextToken(parser.BITMAP) {

	}

	sp.AcceptAndNextTokenWithError(parser.INDEX, true)
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.ON, true)

	cluster := sp.AcceptAndNextToken(parser.CLUSTER)
	x.Cluster = cluster

	onName := parser.ParseName(sp.ExprParser())
	x.SetOnName(onName)

	if sp.AcceptAndNextToken(parser.SYMB_LEFT_PAREN) {
		for {
			column := sp.ParseIndexColumn()
			x.AddColumn(column)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
	}

	return x
}

func (sp *OracleIndexStatementParser) ParseIndexColumn() *index2.SQLIndexColumn {
	x := index2.NewViewColumn()

	expr := parser.ParseExpr(sp.ExprParser())
	x.SetExpr(expr)

	if sp.AcceptAndNextToken(parser.ASC) {

	} else if sp.AcceptAndNextToken(parser.DESC) {

	}

	return x
}

// func (x *OracleIndexStatementParser) ParseDrop() statement.ISQLStatement {
// 	if !x.AcceptAndNextToken(parser.DROP) {
// 		return nil
// 	}
//
// 	return nil
// }

// ------------------------ Package ------------------------

type OraclePackageStatementParser struct {
	*parser.SQLPackageStatementParser
}

func NewPackageStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OraclePackageStatementParser {
	return NewPackageStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewPackageStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OraclePackageStatementParser {
	return NewPackageStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewPackageStatementParserByExprParser(exprParser parser.ISQLExprParser) *OraclePackageStatementParser {
	x := new(OraclePackageStatementParser)
	x.SQLPackageStatementParser = parser.NewPackageStatementParserByExprParser(exprParser)
	return x
}
func (x *OraclePackageStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (x *OraclePackageStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	return nil
}
func (x *OraclePackageStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.DROP) {
		return nil
	}

	return nil
}

// ------------------------ Package Body ------------------------

type OraclePackageBodyStatementParser struct {
	*parser.SQLPackageBodyStatementParser
}

func NewPackageBodyStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OraclePackageBodyStatementParser {
	return NewPackageBodyStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewPackageBodyStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OraclePackageBodyStatementParser {
	return NewPackageBodyStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewPackageBodyStatementParserByExprParser(exprParser parser.ISQLExprParser) *OraclePackageBodyStatementParser {
	x := new(OraclePackageBodyStatementParser)
	x.SQLPackageBodyStatementParser = parser.NewPackageBodyStatementParserByExprParser(exprParser)
	return x
}
func (x *OraclePackageBodyStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (x *OraclePackageBodyStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	return nil
}
func (x *OraclePackageBodyStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.DROP) {
		return nil
	}

	return nil
}

// ------------------------ Procedure ------------------------

type OracleProcedureStatementParser struct {
	*parser.SQLProcedureStatementParser
}

func NewProcedureStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleProcedureStatementParser {
	return NewProcedureStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewProcedureStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleProcedureStatementParser {
	return NewProcedureStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewProcedureStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleProcedureStatementParser {
	x := new(OracleProcedureStatementParser)
	x.SQLProcedureStatementParser = parser.NewProcedureStatementParserByExprParser(exprParser)
	return x
}
func (x *OracleProcedureStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (x *OracleProcedureStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	return nil
}

// func (sp *OracleProcedureStatementParser) ParseDrop() statement.ISQLStatement {
// 	if !sp.AcceptAndNextToken(parser.DROP) {
// 		return nil
// 	}
// 	sp.AcceptAndNextTokenWithError(PROCEDURE, true)
//
// 	x := procedure.NewDropProcedureStatement(sp.DBType())
//
// 	ifExists := ParseIfExists(sp.ExprParser())
// 	x.IfExists = ifExists
//
// 	name := ParseName(sp.ExprParser())
// 	x.SetName(name)
//
// 	return x
// }

// ------------------------ Role ------------------------

type OracleRoleStatementParser struct {
	*parser.SQLRoleStatementParser
}

func NewRoleStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleRoleStatementParser {
	x := new(OracleRoleStatementParser)
	x.SQLRoleStatementParser = parser.NewRoleStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleRoleStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.ROLE, true)
	x := role.NewAlterRoleStatement(sp.DBType())

	return x
}

/**
 * CREATE ROLE [IF NOT EXISTS] role [, role ] ...
 * https://dev.mysql.com/doc/refman/8.0/en/create-role.html
 */
func (sp *OracleRoleStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.ROLE, true)

	x := role.NewCreateRoleStatement(sp.DBType())
	ifNotExists := parser.ParseIfNotExists(sp.ExprParser())
	x.IfNotExists = ifNotExists

	for {
		name := parser.ParseName(sp.ExprParser())
		x.AddName(name)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	return x
}

/**
 * DROP ROLE [IF EXISTS] role [, role ] ...
 */
func (sp *OracleRoleStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.ROLE, true)

	x := role.NewDropRoleStatement(sp.DBType())
	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	for {
		name := parser.ParseName(sp.ExprParser())
		x.AddName(name)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	return x
}

// ------------------------ Schema ------------------------
type OracleSchemaStatementParser struct {
	*parser.SQLSchemaStatementParser
}

func NewSchemaStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleSchemaStatementParser {
	return NewSchemaStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSchemaStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleSchemaStatementParser {
	return NewSchemaStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSchemaStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleSchemaStatementParser {
	x := new(OracleSchemaStatementParser)
	x.SQLSchemaStatementParser = parser.NewSchemaStatementParserByExprParser(exprParser)
	return x
}

func (x *OracleSchemaStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(parser.ALTER) {
		return nil
	}

	return nil
}

//
func (x *OracleSchemaStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(parser.CREATE) {
		return nil
	}

	return nil
}
func (p *OracleSchemaStatementParser) ParseDrop() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.DROP, true)
	p.AcceptAndNextTokenWithError(parser.SCHEMA, true)

	x := schema.NewDropSchemaStatement(p.DBType())

	ifExists := parser.ParseIfExists(p.ExprParser())
	x.IfExists = ifExists

	name := parser.ParseName(p.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Sequence ------------------------
type OracleSequenceStatementParser struct {
	*parser.SQLSequenceStatementParser
}

func NewSequenceStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleSequenceStatementParser {
	x := new(OracleSequenceStatementParser)
	x.SQLSequenceStatementParser = parser.NewSequenceStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleSequenceStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}
	x := sequence.NewAlterSequenceStatement(sp.DBType())

	sp.AcceptAndNextTokenWithError(parser.SEQUENCE, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	for {
		option := sp.ExprParser().ParseSequenceOption(sp.ExprParser())
		if option == nil {
			break
		}
		x.AddOption(option)
	}
	return x
}
func (sp *OracleSequenceStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	x := sequence.NewCreateSequenceStatement(sp.DBType())

	sp.AcceptAndNextTokenWithError(parser.SEQUENCE, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	if sp.Accept(parser.SHARING) {
		parser.ParseAssignExpr(sp.ExprParser())
	}

	for {
		option := sp.ExprParser().ParseSequenceOption(sp.ExprParser())
		if option == nil {
			break
		}
		x.AddOption(option)
	}

	return x
}
func (sp *OracleSequenceStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	x := sequence.NewDropSequenceStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Synonym ------------------------
type OracleSynonymStatementParser struct {
	*parser.SQLSynonymStatementParser
}

func NewSynonymStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleSynonymStatementParser {
	return NewSynonymStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSynonymStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleSynonymStatementParser {
	return NewSynonymStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSynonymStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleSynonymStatementParser {
	x := new(OracleSynonymStatementParser)
	x.SQLSynonymStatementParser = parser.NewSynonymStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleSynonymStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}
	x := synonym.NewAlterSynonymStatement(sp.DBType())

	public := sp.AcceptAndNextToken(parser.PUBLIC)
	x.Public = public

	sp.AcceptAndNextTokenWithError(parser.SYNONYM, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	action := sp.ExprParser().ParseAlterSynonymAction(sp.ExprParser())
	x.SetAction(action)
	return x
}

func (sp *OracleSynonymStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	x := synonym.NewCreateSynonymStatement(sp.DBType())

	orReplace := parser.ParseOrReplace(sp.ExprParser())
	x.OrReplace = orReplace

	public := sp.AcceptAndNextToken(parser.PUBLIC)
	x.Public = public

	sp.AcceptAndNextTokenWithError(parser.SYNONYM, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	if sp.Accept(parser.SHARING) {
		parser.ParseAssignExpr(sp.ExprParser())
	}

	sp.AcceptAndNextTokenWithError(parser.FOR, true)
	forName := parser.ParseName(sp.ExprParser())
	x.SetForName(forName)

	return x
}
func (sp *OracleSynonymStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	x := synonym.NewDropSynonymStatement(sp.DBType())

	public := sp.AcceptAndNextToken(parser.PUBLIC)
	x.Public = public

	sp.AcceptAndNextTokenWithError(parser.SYNONYM, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	force := sp.AcceptAndNextToken(parser.FORCE)
	x.Force = force
	return x
}

// ------------------------ OnTable ------------------------
type OracleTableStatementParser struct {
	*parser.SQLTableStatementParser
}

func NewTableStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleTableStatementParser {
	return NewTableStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTableStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleTableStatementParser {
	return NewTableStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTableStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleTableStatementParser {
	x := new(OracleTableStatementParser)
	x.SQLTableStatementParser = parser.NewTableStatementParserByExprParser(exprParser)
	return x
}

func (x *OracleTableStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(parser.ALTER) {
		return nil
	}

	return nil
}

//
// https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-TABLE.html
func (sp *OracleTableStatementParser) ParseCreate() statement.ISQLStatement {

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	temporary := sp.AcceptAndNextToken(parser.TEMPORARY)
	sp.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := table.NewCreateTableStatement(sp.DBType())

	if temporary {
		x.TableScope = table.TEMPORARY
	}

	ifNotExists := parser.ParseIfNotExists(sp.ExprParser())
	x.IfNotExists = ifNotExists

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	// (TableElements)
	elements := parser.ParseTableElements(sp.ExprParser())
	x.AddElements(elements)

	if sp.Accept(parser.LIKE) {
		x.Paren = false
		likeClause := parser.ParseLikeClause(sp.ExprParser())
		x.AddElement(likeClause)
	}

	partitionBy := sp.ExprParser().ParsePartitionBy(sp.ExprParser())
	x.SetPartitionBy(partitionBy)

	as := sp.AcceptAndNextToken(parser.AS)
	x.As = as
	if as {
		subQuery := parser.ParseSelectQuery(sp.ExprParser())
		x.SetSubQuery(subQuery)
	}

	return x
}

func (p *OracleTableStatementParser) ParseDrop() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.DROP, true)
	temporary := p.AcceptAndNextToken(parser.TEMPORARY)
	p.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := table.NewDropTableStatement(p.DBType())
	x.Temporary = temporary

	ifExists := parser.ParseIfExists(p.ExprParser())
	x.IfExists = ifExists

	for {
		name := parser.ParseName(p.ExprParser())
		if name == nil {
			panic(p.SyntaxError())
		}
		x.AddName(name)
		if !p.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	if p.AcceptAndNextToken(parser.RESTRICT) {

		x.SetOption(exprTable.NewDropTableStatementRestrictOption())

	} else if p.AcceptAndNextToken(parser.CASCADE) {

		x.SetOption(exprTable.NewDropTableStatementCascadeOption())

	}

	return x
}

// ------------------------ Trigger ------------------------

type OracleTriggerStatementParser struct {
	*parser.SQLTriggerStatementParser
}

func NewTriggerStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleTriggerStatementParser {
	return NewTriggerStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTriggerStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleTriggerStatementParser {
	return NewTriggerStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTriggerStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleTriggerStatementParser {
	x := new(OracleTriggerStatementParser)
	x.SQLTriggerStatementParser = parser.NewTriggerStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleTriggerStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (x *OracleTriggerStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	return nil
}
func (sp *OracleTriggerStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.TRIGGER, true)

	x := trigger.NewDropTriggerStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Type ------------------------

type OracleTypeStatementParser struct {
	*parser.SQLTypeStatementParser
}

func NewTypeStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleTypeStatementParser {
	return NewTypeStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTypeStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleTypeStatementParser {
	return NewTypeStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTypeStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleTypeStatementParser {
	x := new(OracleTypeStatementParser)
	x.SQLTypeStatementParser = parser.NewTypeStatementParserByExprParser(exprParser)
	return x
}
func (x *OracleTypeStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (x *OracleTypeStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	return nil
}
func (sp *OracleTypeStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.TYPE, true)

	x := type_.NewDropTypeStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	behavior := sp.ParseDropBehavior()
	x.Behavior = behavior
	return x
}

func (sp *OracleTypeStatementParser) ParseDropBehavior() statement.SQLDropBehavior {
	if sp.AcceptAndNextToken(parser.FORCE) {
		return statement.FORCE

	} else if sp.AcceptAndNextToken(parser.VALIDATE) {
		return statement.VALIDATE
	}
	return ""
}

// ------------------------ Type Body ------------------------
type OracleTypeBodyStatementParser struct {
	*parser.SQLTypeBodyStatementParser
}

func NewTypeBodyStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleTypeBodyStatementParser {
	return NewTypeBodyStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTypeBodyStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleTypeBodyStatementParser {
	return NewTypeBodyStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTypeBodyStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleTypeBodyStatementParser {
	x := new(OracleTypeBodyStatementParser)
	x.SQLTypeBodyStatementParser = parser.NewTypeBodyStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleTypeBodyStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *OracleTypeBodyStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	x := typebody.NewCreateTypeBodyStatement(sp.DBType())

	orReplace := parser.ParseOrReplace(sp.ExprParser())
	x.OrReplace = orReplace

	sp.AcceptAndNextTokenWithError(parser.TYPE, true)
	sp.AcceptAndNextTokenWithError(parser.BODY, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}
func (sp *OracleTypeBodyStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.TYPE, true)
	sp.AcceptAndNextTokenWithError(parser.BODY, true)

	x := typebody.NewDropTypeStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ user ------------------------

type OracleUserStatementParser struct {
	*parser.SQLUserStatementParser
}

func NewUserStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleUserStatementParser {
	x := new(OracleUserStatementParser)
	x.SQLUserStatementParser = parser.NewUserStatementParserByExprParser(exprParser)
	return x
}
func (x *OracleUserStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (sp *OracleUserStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.USER, true)

	x := user.NewCreateUserStatement(sp.DBType())

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	for {
		name := sp.ParseCreateUserName()
		x.AddName(name)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	sp.AcceptAndNextTokenWithError(parser.DEFAULT, true)
	sp.AcceptAndNextTokenWithError(parser.ROLE, true)
	for {
		role := parser.ParseName(sp.ExprParser())
		x.AddRole(role)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	return x
}

func (sp *OracleUserStatementParser) ParseCreateUserName() expr.ISQLExpr {
	var x *exprUser.SQLUserName
	name := parser.ParseName(sp.ExprParser())
	switch name.(type) {
	case *exprUser.SQLUserName:
		x = name.(*exprUser.SQLUserName)
	default:
		x = exprUser.NewUserName()
	}

	if sp.AcceptAndNextToken(parser.IDENTIFIED) {
		if sp.AcceptAndNextToken(parser.BY) {
			if sp.AcceptAndNextToken(parser.RANDOM) {
				sp.AcceptAndNextTokenWithError(parser.PASSWORD, true)

				return x
			} else {
				parser.ParseExpr(sp.ExprParser())
				return x
			}

		} else if sp.AcceptAndNextToken(parser.WITH) {
			parser.ParseExpr(sp.ExprParser())

			if sp.AcceptAndNextToken(parser.BY) {
				if sp.AcceptAndNextToken(parser.RANDOM) {
					sp.AcceptAndNextTokenWithError(parser.PASSWORD, true)
					return x
				} else {
					parser.ParseExpr(sp.ExprParser())
					return x
				}
			} else if sp.AcceptAndNextToken(parser.AS) {
				parser.ParseExpr(sp.ExprParser())
			}

			return x
		}

		panic(sp.UnSupport())
	}

	return x
}

func (sp *OracleUserStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.USER, true)

	x := user.NewDropRoleStatement(sp.DBType())
	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	for {
		name := parser.ParseName(sp.ExprParser())
		x.AddName(name)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}
	return x
}

// ------------------------ View ------------------------
type OracleViewStatementParser struct {
	*parser.SQLViewStatementParser
}

func NewViewStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleViewStatementParser {
	return NewViewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewViewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleViewStatementParser {
	return NewViewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewViewStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleViewStatementParser {
	x := new(OracleViewStatementParser)
	x.SQLViewStatementParser = parser.NewViewStatementParserByExprParser(exprParser)
	return x
}

func (sp *OracleViewStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	x := view.NewAlterViewStatement(sp.DBType())

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	action := sp.ParseAlterViewAction()
	x.AddAction(action)

	return x
}

func (sp *OracleViewStatementParser) ParseAlterViewAction() expr.ISQLExpr {
	if sp.AcceptAndNextToken(parser.ADD) {
		if sp.ExprParser().IsParseTableConstraint() {
			action := exprView.NewAddTableConstraintAlterViewAction()
			tableConstraint := sp.ExprParser().ParseTableConstraint(sp.ExprParser())
			action.SetTableConstraint(tableConstraint)

			return action
		} else {
			panic(sp.UnSupport())
		}

	} else if sp.AcceptAndNextToken(parser.MODIFY) {
		if sp.AcceptAndNextToken(parser.CONSTRAINT) {
			action := exprView.NewModifyConstraintAlterViewAction()

			name := parser.ParseName(sp.ExprParser())
			action.SetName(name)

			if sp.AcceptAndNextToken(parser.RELY) {

			} else if sp.AcceptAndNextToken(parser.NORELY) {

			} else {
				panic(sp.UnSupport())
			}

			return action
		} else {
			panic(sp.UnSupport())
		}

	} else if sp.AcceptAndNextToken(parser.DROP) {
		if sp.AcceptAndNextToken(parser.CONSTRAINT) {

		} else if sp.AcceptAndNextToken(parser.PRIMARY) {
			sp.AcceptAndNextTokenWithError(parser.KEY, true)

			return exprView.NewDropPrimaryKeyTableConstraintAlterViewAction()

		} else if sp.AcceptAndNextToken(parser.UNIQUE) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
			action := exprView.NewDropUniqueAlterViewAction()

			for {
				name := parser.ParseName(sp.ExprParser())
				action.AddName(name)
				if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
					break
				}
			}

			sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

			return action

		} else {
			panic(sp.UnSupport())
		}

	} else if sp.AcceptAndNextToken(parser.COMPILE) {
		return statement.NewCompileExpr()

	} else if sp.AcceptAndNextToken(parser.READ) {
		if sp.AcceptAndNextToken(parser.ONLY) {
			return statement.NewReadOnlyExpr()

		} else if sp.AcceptAndNextToken(parser.WRITE) {
			return statement.NewReadWriteExpr()

		} else {
			panic(sp.UnSupport())
		}

	} else if sp.AcceptAndNextToken(parser.EDITIONABLE) {
		return statement.NewEditionAbleExpr()

	} else if sp.AcceptAndNextToken(parser.NONEDITIONABLE) {
		return statement.NewNonEditionAbleExpr()
	}
	return nil
}

func (sp *OracleViewStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	x := view.NewCreateViewStatement(sp.DBType())

	orReplace := parser.ParseOrReplace(sp.ExprParser())
	x.OrReplace = orReplace

	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	if sp.AcceptAndNextToken(parser.OF) {

	}

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	if sp.Accept(parser.SHARING) {
		parser.ParseAssignExpr(sp.ExprParser())
	}

	elements := parser.ParseViewElements(sp.ExprParser())
	x.AddElements(elements)

	sp.AcceptAndNextTokenWithError(parser.AS, true)
	subQuery := parser.ParseSelectQuery(sp.ExprParser())
	x.SetSubQuery(subQuery)

	return x
}

/**
 * DROP VIEW [ schema. ] view [ CASCADE CONSTRAINTS ] ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DROP-VIEW.html#GUID-1A1BD841-66B9-47E4-896F-D36E075AE296
 */
func (sp *OracleViewStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	x := view.NewDropViewStatement(sp.DBType())

	name := parser.ParseName(sp.ExprParser())
	x.AddName(name)

	if sp.AcceptAndNextToken(parser.CASCADE) {
		sp.AcceptAndNextTokenWithError(parser.CONSTRAINTS, true)
		x.Behavior = statement.CASCADE_CONSTRAINTS
	}

	return x
}

// --------------------------- EXPLAIN Statement Start ---------------------------
type OracleExplainStatementParser struct {
	*parser.SQLExplainStatementParser
}

func NewExplainStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleExplainStatementParser {
	x := new(OracleExplainStatementParser)
	x.SQLExplainStatementParser = parser.NewExplainStatementParserByExprParser(exprParser)
	return x
}
func (sp *OracleExplainStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.EXPLAIN, true)
	sp.AcceptAndNextTokenWithError(parser.PLAN, true)

	x := statement.NewExplainStatement(sp.DBType())

	if sp.AcceptAndNextToken(parser.SET) {
		sp.AcceptAndNextTokenWithError(parser.STATEMENT_ID, true)
		sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
		parser.ParseExpr(sp.ExprParser())

	}

	if sp.AcceptAndNextToken(parser.INTO) {
		intoTable := parser.ParseName(sp.ExprParser())
		x.SetIntoTable(intoTable)
	}

	sp.AcceptAndNextTokenWithError(parser.FOR, true)
	stmt := parser.ParseStatement(sp.ExprParser())
	x.SetStmt(stmt)

	return x
}

// --------------------------- EXPLAIN Statement End ---------------------------
