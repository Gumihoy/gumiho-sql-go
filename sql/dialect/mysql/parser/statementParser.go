package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	server2 "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/server"
	exprTable "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	user2 "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/user"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/variable"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/database"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/function"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/index"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/procedure"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/role"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/schema"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/server"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/set"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/show"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/trigger"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/user"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type MySQLStatementParser struct {
	*parser.SQLStatementParser
}

func NewStatementParserBySQL(sourceSQL string, dbType db.Type, config *parser.SQLParseConfig) *MySQLStatementParser {
	return NewStatementParserByLexer(NewLexer(sourceSQL), dbType, config)
}

func NewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLStatementParser {
	x := new(MySQLStatementParser)
	x.SQLStatementParser = parser.NewStatementParserByExprParser(exprParser)
	return x
}

// ------------------------------------------------------------- DML -------------------------------------------------------------
// ------------------------ Select ------------------------

type MySQLSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func NewSelectStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSelectStatementParser {
	x := new(MySQLSelectStatementParser)
	x.SQLSelectStatementParser = parser.NewSelectStatementParserByExprParser(exprParser)
	return x
}

func (x *MySQLSelectStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.WITH) && !x.Accept(parser.SELECT) {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(x.DBType(), query)
}

// ------------------------ Delete ------------------------

type MySQLDeleteStatementParser struct {
	*parser.SQLDeleteStatementParser
}

func NewDeleteStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLDeleteStatementParser {
	return NewDeleteStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDeleteStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDeleteStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLDeleteStatementParser {
	x := new(MySQLDeleteStatementParser)
	x.SQLDeleteStatementParser = parser.NewDeleteStatementParserByExprParser(exprParser)
	return x
}

func (x *MySQLDeleteStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.DELETE) {
		return nil
	}
	return nil
}

// ------------------------ Insert ------------------------

type MySQLInsertStatementParser struct {
	*parser.SQLInsertStatementParser
}

func (x *MySQLInsertStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.INSERT) {
		return nil
	}

	return nil
}

func NewInsertStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLInsertStatementParser {
	return NewInsertStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewInsertStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLInsertStatementParser {
	return NewInsertStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewInsertStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLInsertStatementParser {
	x := new(MySQLInsertStatementParser)
	x.SQLInsertStatementParser = parser.NewInsertStatementParserByExprParser(exprParser)
	return x
}

// ------------------------ Update ------------------------
type MySQLUpdateStatementParser struct {
	*parser.SQLUpdateStatementParser
}

func (x *MySQLUpdateStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.UPDATE) {
		return nil
	}
	return nil
}

func NewUpdateStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLUpdateStatementParser {
	return NewUpdateStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewUpdateStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLUpdateStatementParser {
	return NewUpdateStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewUpdateStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLUpdateStatementParser {
	x := new(MySQLUpdateStatementParser)
	x.SQLUpdateStatementParser = parser.NewUpdateStatementParserByExprParser(exprParser)
	return x
}

// ------------------------------------------------------------- DDL -------------------------------------------------------------

// ------------------------ Database ------------------------
type MySQLDatabaseStatementParser struct {
	*parser.SQLDatabaseStatementParser
}

func NewDatabaseStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLDatabaseStatementParser {
	return NewDatabaseStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDatabaseStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLDatabaseStatementParser {
	return NewDatabaseStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDatabaseStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLDatabaseStatementParser {
	x := new(MySQLDatabaseStatementParser)
	x.SQLDatabaseStatementParser = parser.NewDatabaseStatementParserByExprParser(exprParser)
	return x
}

func (sp *MySQLDatabaseStatementParser) ParseAlter() statement.ISQLStatement {

	sp.AcceptAndNextTokenWithError(parser.ALTER, true)
	sp.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewAlterDatabaseStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	for {
		action := sp.ParseAlterDatabaseAction()
		if action == nil {
			break
		}
		x.AddAction(action)
	}

	return x
}
func (sp *MySQLDatabaseStatementParser) ParseAlterDatabaseAction() expr.ISQLExpr {
	var hasDefault bool
	if sp.AcceptAndNextToken(parser.DEFAULT) {
		hasDefault = true
	}
	if sp.AcceptAndNextToken(parser.CHARACTER) {
		sp.AcceptAndNextTokenWithError(parser.SET, true)

		x := exprTable.NewCharacterSetAssignExpr()
		x.Default = hasDefault

		equal := sp.AcceptAndNextToken(parser.SYMB_EQUAL)
		x.Equal = equal

		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)

		return x

	} else if sp.AcceptAndNextToken(parser.COLLATE) {
		sp.AcceptAndNextToken(parser.SYMB_EQUAL)

		parser.ParseExpr(sp.ExprParser())

	} else if sp.AcceptAndNextToken(parser.ENCRYPTION) {

		sp.AcceptAndNextToken(parser.SYMB_EQUAL)

		parser.ParseExpr(sp.ExprParser())

	} else if sp.AcceptAndNextToken(parser.READ) {
		sp.AcceptAndNextTokenWithError(parser.ONLY, true)
		sp.AcceptAndNextToken(parser.SYMB_EQUAL)

		parser.ParseExpr(sp.ExprParser())
	}

	if hasDefault {
		panic(sp.UnSupport())
	}

	return nil
}

//
func (p *MySQLDatabaseStatementParser) ParseCreate() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.CREATE, true)
	p.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewCreateDatabaseStatement(p.DBType())

	ifNotExists := parser.ParseIfNotExists(p.ExprParser())
	x.IfNotExists = ifNotExists

	name := parser.ParseName(p.ExprParser())
	x.SetName(name)

	return x

}
func (p *MySQLDatabaseStatementParser) ParseDrop() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.DROP, true)
	p.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewDropDatabaseSStatement(p.DBType())

	ifExists := parser.ParseIfExists(p.ExprParser())
	x.IfExists = ifExists

	name := parser.ParseName(p.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Function ------------------------

type MySQLFunctionStatementParser struct {
	*parser.SQLFunctionStatementParser
}

func NewFunctionStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLFunctionStatementParser {
	return NewFunctionStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewFunctionStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLFunctionStatementParser {
	return NewFunctionStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewFunctionStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLFunctionStatementParser {
	x := new(MySQLFunctionStatementParser)
	x.SQLFunctionStatementParser = parser.NewFunctionStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLFunctionStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.FUNCTION, true)

	x := function.NewAlterFunctionStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	for {
		action := sp.ParseAlterFunctionAction()
		if action == nil {
			break
		}
		x.AddAction(action)
	}
	return x
}
func (sp *MySQLFunctionStatementParser) ParseAlterFunctionAction() expr.ISQLExpr {
	if sp.Accept(parser.COMMENT) {

	} else if sp.AcceptAndNextToken(parser.LANGUAGE) {

	} else if sp.AcceptAndNextToken(parser.CONTAINS) {

	} else if sp.AcceptAndNextToken(parser.NO) {

	} else if sp.AcceptAndNextToken(parser.READS) {

	} else if sp.AcceptAndNextToken(parser.MODIFIES) {

	} else if sp.AcceptAndNextToken(parser.SQL) {

	}
	return nil
}

func (x *MySQLFunctionStatementParser) ParseCreate() statement.ISQLStatement {
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

type MySQLIndexStatementParser struct {
	*parser.SQLIndexStatementParser
}

func NewIndexStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLIndexStatementParser {
	return NewIndexStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewIndexStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLIndexStatementParser {
	return NewIndexStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewIndexStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLIndexStatementParser {
	x := new(MySQLIndexStatementParser)
	x.SQLIndexStatementParser = parser.NewIndexStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLIndexStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *MySQLIndexStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	x := index.NewCreateIndexStatement(sp.DBType())

	if sp.AcceptAndNextToken(parser.UNIQUE) {

	} else if sp.AcceptAndNextToken(parser.FULLTEXT) {

	} else if sp.AcceptAndNextToken(parser.SPATIAL) {

	}

	sp.AcceptAndNextTokenWithError(parser.INDEX, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.ON, true)
	onTable := parser.ParseName(sp.ExprParser())
	x.SetOnName(onTable)

	sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
	for {

		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

	return x
}

func (sp *MySQLIndexStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.INDEX, true)

	x := index.NewDropIndexStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.ON, true)
	onTable := parser.ParseName(sp.ExprParser())
	x.SetOnTable(onTable)

	for {
		option := sp.ParseDropOption()
		if option == nil {
			break
		}
		x.AddOption(option)
	}

	return x
}

func (sp *MySQLIndexStatementParser) ParseDropOption() expr.ISQLExpr {
	if sp.Accept(parser.ALGORITHM) {
		return parser.ParseAssignExpr(sp.ExprParser())

	} else if sp.Accept(parser.LOCK) {
		return parser.ParseAssignExpr(sp.ExprParser())

	}
	return nil
}

// ------------------------ Package ------------------------

type MySQLPackageStatementParser struct {
	*parser.SQLPackageStatementParser
}

func NewPackageStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLPackageStatementParser {
	return NewPackageStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewPackageStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLPackageStatementParser {
	return NewPackageStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewPackageStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLPackageStatementParser {
	x := new(MySQLPackageStatementParser)
	x.SQLPackageStatementParser = parser.NewPackageStatementParserByExprParser(exprParser)
	return x
}
func (x *MySQLPackageStatementParser) ParseAlter() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *MySQLPackageStatementParser) ParseCreate() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *MySQLPackageStatementParser) ParseDrop() statement.ISQLStatement {
	panic(x.UnSupport())
}

// ------------------------ Package Body ------------------------

type MySQLPackageBodyStatementParser struct {
	*parser.SQLPackageBodyStatementParser
}

func NewPackageBodyStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLPackageBodyStatementParser {
	return NewPackageBodyStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewPackageBodyStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLPackageBodyStatementParser {
	return NewPackageBodyStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewPackageBodyStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLPackageBodyStatementParser {
	x := new(MySQLPackageBodyStatementParser)
	x.SQLPackageBodyStatementParser = parser.NewPackageBodyStatementParserByExprParser(exprParser)
	return x
}
func (x *MySQLPackageBodyStatementParser) ParseAlter() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *MySQLPackageBodyStatementParser) ParseCreate() statement.ISQLStatement {
	panic(x.UnSupport())
}
func (x *MySQLPackageBodyStatementParser) ParseDrop() statement.ISQLStatement {
	panic(x.UnSupport())
}

// ------------------------ Procedure ------------------------

type MySQLProcedureStatementParser struct {
	*parser.SQLProcedureStatementParser
}

func NewProcedureStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLProcedureStatementParser {
	return NewProcedureStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewProcedureStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLProcedureStatementParser {
	return NewProcedureStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewProcedureStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLProcedureStatementParser {
	x := new(MySQLProcedureStatementParser)
	x.SQLProcedureStatementParser = parser.NewProcedureStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLProcedureStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.PROCEDURE, true)
	x := procedure.NewAlterProcedureStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	for {
		action := sp.ParseAlterProcedureAction()
		if action == nil {
			break
		}
		x.AddAction(action)
	}

	return x
}
func (sp *MySQLProcedureStatementParser) ParseAlterProcedureAction() expr.ISQLExpr {
	if sp.Accept(parser.COMMENT) {

	} else if sp.AcceptAndNextToken(parser.LANGUAGE) {

	} else if sp.AcceptAndNextToken(parser.CONTAINS) {

	} else if sp.AcceptAndNextToken(parser.NO) {

	} else if sp.AcceptAndNextToken(parser.READS) {

	} else if sp.AcceptAndNextToken(parser.MODIFIES) {

	} else if sp.AcceptAndNextToken(parser.SQL) {

	}
	return nil
}
func (x *MySQLProcedureStatementParser) ParseCreate() statement.ISQLStatement {
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
// 	x := procedure.NewDropIndexStatement(sp.DBType())
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

type MySQLRoleStatementParser struct {
	*parser.SQLRoleStatementParser
}

func NewRoleStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLRoleStatementParser {
	x := new(MySQLRoleStatementParser)
	x.SQLRoleStatementParser = parser.NewRoleStatementParserByExprParser(exprParser)
	return x
}

func (sp *MySQLRoleStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}

/**
 * CREATE ROLE [IF NOT EXISTS] role [, role ] ...
 * https://dev.mysql.com/doc/refman/8.0/en/create-role.html
 */
func (sp *MySQLRoleStatementParser) ParseCreate() statement.ISQLStatement {
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
func (sp *MySQLRoleStatementParser) ParseDrop() statement.ISQLStatement {
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
type MySQLSchemaStatementParser struct {
	*parser.SQLSchemaStatementParser
}

func NewSchemaStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLSchemaStatementParser {
	return NewSchemaStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSchemaStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLSchemaStatementParser {
	return NewSchemaStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSchemaStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSchemaStatementParser {
	x := new(MySQLSchemaStatementParser)
	x.SQLSchemaStatementParser = parser.NewSchemaStatementParserByExprParser(exprParser)
	return x
}

func (sp *MySQLSchemaStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}

//
func (sp *MySQLSchemaStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.SCHEMA, true)

	x := schema.NewCreateSchemaStatement(sp.DBType())

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}
func (sp *MySQLSchemaStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.SCHEMA, true)

	x := schema.NewDropSchemaStatement(sp.DBType())

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Sequence ------------------------
type MySQLSequenceStatementParser struct {
	*parser.SQLSequenceStatementParser
}

func NewSequenceStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSequenceStatementParser {
	x := new(MySQLSequenceStatementParser)
	x.SQLSequenceStatementParser = parser.NewSequenceStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLSequenceStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *MySQLSequenceStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *MySQLSequenceStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ Server ------------------------
type MySQLServerStatementParser struct {
	*parser.SQLServerStatementParser
}

func NewServerStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLServerStatementParser {
	x := new(MySQLServerStatementParser)
	x.SQLServerStatementParser = parser.NewServerStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLServerStatementParser) ParseAlter() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.ALTER, true)
	sp.AcceptAndNextTokenWithError(parser.SERVER, true)

	x := server.NewAlterServerStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.OPTIONS, true)
	sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
	for {
		option := sp.ParseServerOption()
		x.AddOption(option)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
	return x
}

func (sp *MySQLServerStatementParser) ParseCreate() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.SERVER, true)

	x := server.NewCreateServerStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	sp.AcceptAndNextTokenWithError(parser.FOREIGN, true)
	sp.AcceptAndNextTokenWithError(parser.DATA, true)
	sp.AcceptAndNextTokenWithError(parser.WRAPPER, true)
	wrapperName := parser.ParseName(sp.ExprParser())
	x.SetWrapperName(wrapperName)

	sp.AcceptAndNextTokenWithError(parser.OPTIONS, true)
	sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
	for {
		option := sp.ParseServerOption()
		x.AddOption(option)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}
	sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
	return x
}
func (sp *MySQLServerStatementParser) ParseServerOption() expr.ISQLExpr {
	if sp.AcceptAndNextToken(parser.HOST) {
		x := server2.NewHostOption()
		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)
		return x
	} else if sp.AcceptAndNextToken(parser.DATABASE) {
		x := server2.NewDatabaseOption()
		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(parser.USER) {
		x := server2.NewUserOption()
		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(parser.PASSWORD) {
		x := server2.NewPasswordOption()
		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(parser.SOCKET) {
		x := server2.NewSocketOption()
		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(parser.OWNER) {
		x := server2.NewOwnerOption()
		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(parser.PORT) {
		x := server2.NewPortOption()
		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)
		return x

	}
	return nil
}
func (sp *MySQLServerStatementParser) ParseDrop() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.DROP, true)
	sp.AcceptAndNextTokenWithError(parser.SERVER, true)

	x := server.NewDropServerStatement(sp.DBType())

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Synonym ------------------------
type MySQLSynonymStatementParser struct {
	*parser.SQLSynonymStatementParser
}

func NewSynonymStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLSynonymStatementParser {
	return NewSynonymStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSynonymStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLSynonymStatementParser {
	return NewSynonymStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSynonymStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSynonymStatementParser {
	x := new(MySQLSynonymStatementParser)
	x.SQLSynonymStatementParser = parser.NewSynonymStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLSynonymStatementParser) ParseAlter() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *MySQLSynonymStatementParser) ParseCreate() statement.ISQLStatement {
	panic(sp.UnSupport())
}
func (sp *MySQLSynonymStatementParser) ParseDrop() statement.ISQLStatement {
	panic(sp.UnSupport())
}

// ------------------------ Table ------------------------
type MySQLTableStatementParser struct {
	*parser.SQLTableStatementParser
}

func NewTableStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLTableStatementParser {
	return NewTableStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTableStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLTableStatementParser {
	return NewTableStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTableStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLTableStatementParser {
	x := new(MySQLTableStatementParser)
	x.SQLTableStatementParser = parser.NewTableStatementParserByExprParser(exprParser)
	return x
}

/**
 * ALTER TABLE tbl_name
    [alter_option [, alter_option] ...]
    [partition_options]
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
func (sp *MySQLTableStatementParser) ParseAlter() statement.ISQLStatement {

	sp.AcceptAndNextTokenWithError(parser.ALTER, true)
	sp.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := table.NewAlterTableStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	for {
		action := sp.ExprParser().ParseAlterTableAction(sp.ExprParser())
		x.AddAction(action)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	return x
}

//
// https://dev.mysql.com/doc/refman/8.0/en/create-table.html
func (sp *MySQLTableStatementParser) ParseCreate() statement.ISQLStatement {

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

	for {
		option := sp.ParseOption()
		if option == nil {
			break
		}
		x.AddOption(option)
		sp.AcceptAndNextToken(parser.SYMB_COMMA)
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

/**
 *  AUTO_INCREMENT [=] value
  | AVG_ROW_LENGTH [=] value
  | [DEFAULT] CHARACTER SET [=] charset_name
  | CHECKSUM [=] {0 | 1}
  | [DEFAULT] COLLATE [=] collation_name
  | COMMENT [=] 'string'
  | COMPRESSION [=] {'ZLIB' | 'LZ4' | 'NONE'}
  | CONNECTION [=] 'connect_string'
  | {DATA | INDEX} DIRECTORY [=] 'absolute path to directory'
  | DELAY_KEY_WRITE [=] {0 | 1}
  | ENCRYPTION [=] {'Y' | 'N'}
  | ENGINE [=] engine_name
  | INSERT_METHOD [=] { NO | FIRST | LAST }
  | KEY_BLOCK_SIZE [=] value
  | MAX_ROWS [=] value
  | MIN_ROWS [=] value
  | PACK_KEYS [=] {0 | 1 | DEFAULT}
  | PASSWORD [=] 'string'
  | ROW_FORMAT [=] {DEFAULT | DYNAMIC | FIXED | COMPRESSED | REDUNDANT | COMPACT}
  | STATS_AUTO_RECALC [=] {DEFAULT | 0 | 1}
  | STATS_PERSISTENT [=] {DEFAULT | 0 | 1}
  | STATS_SAMPLE_PAGES [=] value
  | TABLESPACE tablespace_name [STORAGE {DISK | MEMORY}]
  | UNION [=] (tbl_name[,tbl_name]...)
 */
func (sp *MySQLTableStatementParser) IsParseOption() bool {
	if sp.Accept(parser.AUTO_INCREMENT) ||
		sp.Accept(parser.AVG_ROW_LENGTH) ||
		sp.Accept(parser.DEFAULT) ||
		sp.Accept(parser.CHARACTER) ||
		sp.Accept(parser.CHECKSUM) ||
		sp.Accept(parser.COLLATE) ||
		sp.Accept(parser.COMMENT) ||
		sp.Accept(parser.COMPRESSION) ||
		sp.Accept(parser.CONNECTION) ||
		sp.Accept(parser.DATA) ||
		sp.Accept(parser.INDEX) ||
		sp.Accept(parser.DELAY_KEY_WRITE) ||
		sp.Accept(parser.ENCRYPTION) ||
		sp.Accept(parser.ENGINE) ||
		sp.Accept(parser.INSERT_METHOD) ||
		sp.Accept(parser.KEY_BLOCK_SIZE) ||
		sp.Accept(parser.MAX_ROWS) ||
		sp.Accept(parser.MIN_ROWS) ||
		sp.Accept(parser.PACK_KEYS) ||
		sp.Accept(parser.PASSWORD) ||
		sp.Accept(parser.ROW_FORMAT) ||
		sp.Accept(parser.STATS_AUTO_RECALC) ||
		sp.Accept(parser.STATS_PERSISTENT) ||
		sp.Accept(parser.STATS_SAMPLE_PAGES) ||
		sp.Accept(parser.TABLESPACE) ||
		sp.Accept(parser.UNION) {
		return true
	}
	return false
}
func (sp *MySQLTableStatementParser) ParseOption() expr.ISQLExpr {
	if !sp.IsParseOption() {
		return nil
	}
	hasDefault := sp.AcceptAndNextToken(parser.DEFAULT)

	if sp.AcceptAndNextToken(parser.CHARSET) {
		x := exprTable.NewCharsetAssignExpr()
		x.Default = hasDefault

		equal := sp.AcceptAndNextToken(parser.SYMB_EQUAL)
		x.Equal = equal

		value := parser.ParseName(sp.ExprParser())
		x.SetValue(value)

		return x

	} else if sp.AcceptAndNextToken(parser.CHARACTER) {
		sp.AcceptAndNextTokenWithError(parser.SET, true)

		x := exprTable.NewCharacterSetAssignExpr()
		x.Default = hasDefault

		equal := sp.AcceptAndNextToken(parser.SYMB_EQUAL)
		x.Equal = equal

		value := parser.ParseName(sp.ExprParser())
		x.SetValue(value)

		return x

	} else if sp.AcceptAndNextToken(parser.COLLATE) {
		x := exprTable.NewCollateAssignExpr()
		x.Default = hasDefault

		equal := sp.AcceptAndNextToken(parser.SYMB_EQUAL)
		x.Equal = equal

		value := parser.ParseName(sp.ExprParser())
		x.SetValue(value)

		return x

	} else if sp.AcceptAndNextToken(parser.DATA) {
		sp.AcceptAndNextTokenWithError(parser.DIRECTORY, true)

		x := exprTable.NewDataDirectoryAssignExpr()
		equal := sp.AcceptAndNextToken(parser.SYMB_EQUAL)
		x.Equal = equal

		value := parser.ParseName(sp.ExprParser())
		x.SetValue(value)

	} else if sp.AcceptAndNextToken(parser.INDEX) {
		sp.AcceptAndNextTokenWithError(parser.DIRECTORY, true)

		x := exprTable.NewIndexDirectoryAssignExpr()

		equal := sp.AcceptAndNextToken(parser.SYMB_EQUAL)
		x.Equal = equal

		value := parser.ParseName(sp.ExprParser())
		x.SetValue(value)
	}

	if hasDefault {
		panic(sp.UnSupport())
	}

	return parser.ParseAssignExpr(sp.ExprParser())
}

func (p *MySQLTableStatementParser) ParseDrop() statement.ISQLStatement {

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

type MySQLTriggerStatementParser struct {
	*parser.SQLTriggerStatementParser
}

func NewTriggerStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLTriggerStatementParser {
	return NewTriggerStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTriggerStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLTriggerStatementParser {
	return NewTriggerStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTriggerStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLTriggerStatementParser {
	x := new(MySQLTriggerStatementParser)
	x.SQLTriggerStatementParser = parser.NewTriggerStatementParserByExprParser(exprParser)
	return x
}
func (x *MySQLTriggerStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (x *MySQLTriggerStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.CREATE) {
		return nil
	}

	return nil
}
func (sp *MySQLTriggerStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.TRIGGER, true)

	x := trigger.NewDropTriggerStatement(sp.DBType())

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ user ------------------------

type MySQLUserStatementParser struct {
	*parser.SQLUserStatementParser
}

func NewUserStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLUserStatementParser {
	x := new(MySQLUserStatementParser)
	x.SQLUserStatementParser = parser.NewUserStatementParserByExprParser(exprParser)
	return x
}
func (x *MySQLUserStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (sp *MySQLUserStatementParser) ParseCreate() statement.ISQLStatement {
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

	// MySQL 8.0
	if sp.AcceptAndNextToken(parser.DEFAULT) {
		sp.AcceptAndNextTokenWithError(parser.ROLE, true)
		for {
			role := parser.ParseName(sp.ExprParser())
			x.AddRole(role)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
	}

	return x
}

func (sp *MySQLUserStatementParser) ParseCreateUserName() expr.ISQLExpr {
	var x *user2.SQLUserName
	name := parser.ParseName(sp.ExprParser())
	switch name.(type) {
	case *user2.SQLUserName:
		x = name.(*user2.SQLUserName)
	default:
		x = user2.NewUserName()
	}

	if sp.AcceptAndNextToken(parser.IDENTIFIED) {
		if sp.AcceptAndNextToken(parser.BY) {
			if sp.AcceptAndNextToken(parser.PASSWORD) {

				option := user2.NewIdentifiedByPasswordAuthOption()
				auth := parser.ParseExpr(sp.ExprParser())
				option.SetAuth(auth)
				x.SetOption(option)

			} else if sp.AcceptAndNextToken(parser.RANDOM) {
				sp.AcceptAndNextTokenWithError(parser.PASSWORD, true)

				option := user2.NewIdentifiedByRandomPasswordAuthOption()
				x.SetOption(option)

			} else {
				option := user2.NewIdentifiedByAuthOption()
				auth := parser.ParseExpr(sp.ExprParser())
				option.SetAuth(auth)
				x.SetOption(option)

			}

			return x

		} else if sp.AcceptAndNextToken(parser.WITH) {
			plugin := parser.ParseExpr(sp.ExprParser())

			if sp.AcceptAndNextToken(parser.BY) {
				if sp.AcceptAndNextToken(parser.RANDOM) {
					sp.AcceptAndNextTokenWithError(parser.PASSWORD, true)

					option := user2.NewIdentifiedWithByRandomPasswordAuthOption()
					option.SetPlugin(plugin)

					x.SetOption(option)

				} else {
					option := user2.NewIdentifiedWithByAuthOption()
					option.SetPlugin(plugin)

					auth := parser.ParseExpr(sp.ExprParser())
					option.SetAuth(auth)

					x.SetOption(option)
				}

			} else if sp.AcceptAndNextToken(parser.AS) {
				option := user2.NewIdentifiedWithAsAuthOption()
				option.SetPlugin(plugin)

				auth := parser.ParseExpr(sp.ExprParser())
				option.SetAuth(auth)

				x.SetOption(option)

			} else {

				panic(sp.UnSupport())
			}

			return x
		}

		panic(sp.UnSupport())
	}

	return x
}

func (sp *MySQLUserStatementParser) ParseDrop() statement.ISQLStatement {
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
type MySQLViewStatementParser struct {
	*parser.SQLViewStatementParser
}

func NewViewStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLViewStatementParser {
	return NewViewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewViewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLViewStatementParser {
	return NewViewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewViewStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLViewStatementParser {
	x := new(MySQLViewStatementParser)
	x.SQLViewStatementParser = parser.NewViewStatementParserByExprParser(exprParser)
	return x
}

func (sp *MySQLViewStatementParser) ParseAlter() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.ALTER) {
		return nil
	}
	x := view.NewAlterViewStatement(sp.DBType())

	if sp.Accept(parser.ALGORITHM) {
		parser.ParseAssignExpr(sp.ExprParser())

	} else if sp.Accept(parser.DEFINER) {
		parser.ParseAssignExpr(sp.ExprParser())

	}

	if sp.AcceptAndNextToken(parser.SQL) {
		sp.AcceptAndNextTokenWithError(parser.SECURITY, true)
	}

	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	elements := parser.ParseViewElements(sp.ExprParser())
	x.AddElements(elements)

	sp.AcceptAndNextTokenWithError(parser.AS, true)
	subQuery := parser.ParseSelectQuery(sp.ExprParser())
	x.SetSubQuery(subQuery)

	if sp.AcceptAndNextToken(parser.WITH) {
		var withCheckOptionKind view.SQLWithCheckOptionKind
		if sp.AcceptAndNextToken(parser.CASCADED) {
			withCheckOptionKind = view.CASCADED

		} else if sp.AcceptAndNextToken(parser.LOCAL) {
			withCheckOptionKind = view.LOCAL

		} else {
			panic(sp.UnSupport())
		}

		sp.AcceptAndNextTokenWithError(parser.CHECK, true)
		sp.AcceptAndNextTokenWithError(parser.OPTION, true)
		x.WithCheckOptionKind = withCheckOptionKind
	}
	return x
}

func (sp *MySQLViewStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.CREATE) {
		return nil
	}
	x := view.NewCreateViewStatement(sp.DBType())

	orReplace := parser.ParseOrReplace(sp.ExprParser())
	x.OrReplace = orReplace

	if sp.Accept(parser.ALGORITHM) {
		parser.ParseAssignExpr(sp.ExprParser())

	} else if sp.Accept(parser.DEFINER) {
		parser.ParseAssignExpr(sp.ExprParser())

	}

	if sp.AcceptAndNextToken(parser.SQL) {
		sp.AcceptAndNextTokenWithError(parser.SECURITY, true)
	}

	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	elements := parser.ParseViewElements(sp.ExprParser())
	x.AddElements(elements)

	sp.AcceptAndNextTokenWithError(parser.AS, true)
	subQuery := parser.ParseSelectQuery(sp.ExprParser())
	x.SetSubQuery(subQuery)

	return x
}

/**
 * DROP VIEW [IF EXISTS] view_name [, view_name] ... [RESTRICT | CASCADE]
 * https://dev.mysql.com/doc/refman/8.0/en/drop-view.html
 */
func (sp *MySQLViewStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	x := view.NewDropViewStatement(sp.DBType())

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	for {
		name := parser.ParseName(sp.ExprParser())
		x.AddName(name)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) || name == nil {
			break
		}
	}

	behavior := parser.ParseDropBehavior(sp)
	x.Behavior = behavior

	return x
}

// --------------------------- Set Statement ---------------------------

type MySQLSetVariableAssignmentStatementParser struct {
	*parser.SQLSetVariableAssignmentStatementParser
}

func NewSetVariableAssignmentStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSetVariableAssignmentStatementParser {
	x := new(MySQLSetVariableAssignmentStatementParser)
	x.SQLSetVariableAssignmentStatementParser = parser.NewSetVariableAssignmentStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLSetVariableAssignmentStatementParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SET) {
		return nil
	}

	x := set.NewSetVariableAssignmentStatement(sp.DBType())

	for {
		element := sp.ParseSetVariableAssignmentElement()
		x.AddElement(element)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) || element == nil {
			break
		}
	}

	return x
}

func (sp *MySQLSetVariableAssignmentStatementParser) ParseSetVariableAssignmentElement() *expr.SQLAssignExpr {
	if sp.AcceptAndNextToken(parser.GLOBAL) {
		x := expr.NewAssignExpr()

		name := variable.NewAtAtVariableExpr()
		name.Kind = variable.GLOBAL
		variableName := parser.ParseName(sp.ExprParser())
		name.SetName(variableName)
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
		x.Equal = true

		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)

		return x
	} else if sp.AcceptAndNextToken(parser.LOCAL) {
		x := expr.NewAssignExpr()

		name := variable.NewAtAtVariableExpr()
		name.Kind = variable.LOCAL
		variableName := parser.ParseName(sp.ExprParser())
		name.SetName(variableName)
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
		x.Equal = true

		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)

		return x
	} else if sp.AcceptAndNextToken(parser.PERSIST) {
		x := expr.NewAssignExpr()

		name := variable.NewAtAtVariableExpr()
		name.Kind = variable.PERSIST
		variableName := parser.ParseName(sp.ExprParser())
		name.SetName(variableName)
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
		x.Equal = true

		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)

		return x

	} else if sp.AcceptAndNextToken(parser.PERSIST_ONLY) {
		x := expr.NewAssignExpr()

		name := variable.NewAtAtVariableExpr()
		name.Kind = variable.PERSIST_ONLY
		variableName := parser.ParseName(sp.ExprParser())
		name.SetName(variableName)
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
		x.Equal = true

		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)

		return x

	} else if sp.AcceptAndNextToken(parser.SESSION) {
		x := expr.NewAssignExpr()

		name := variable.NewAtAtVariableExpr()
		name.Kind = variable.SESSION

		variableName := parser.ParseName(sp.ExprParser())
		name.SetName(variableName)
		x.SetName(name)

		sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
		x.Equal = true

		value := parser.ParseExpr(sp.ExprParser())
		x.SetValue(value)

		return x

	} else if sp.Accept(parser.SYMB_AT) || parser.IsIdentifier(sp.Kind()) {
		return parser.ParseAssignExpr(sp.ExprParser())
	}

	panic(sp.UnSupport())
}

/**
 *
 */
type MySQLSetCharacterSetStatementParser struct {
	*parser.SQLSetCharacterSetStatementParser
}

func NewSetCharacterSetStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSetCharacterSetStatementParser {
	x := new(MySQLSetCharacterSetStatementParser)
	x.SQLSetCharacterSetStatementParser = parser.NewSetCharacterSetStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLSetCharacterSetStatementParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SET) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.CHARACTER, true)
	sp.AcceptAndNextTokenWithError(parser.SET, true)

	x := set.NewSetCharacterSetStatement(sp.DBType())
	name := parser.ParseExpr(sp.ExprParser())
	x.SetName(name)

	return x
}

/**
 *
 */
type MySQLSeCharsetStatementParser struct {
	*parser.SQLSeCharsetStatementParser
}

func NewSetCharsetStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSeCharsetStatementParser {
	x := new(MySQLSeCharsetStatementParser)
	x.SQLSeCharsetStatementParser = parser.NewSetCharsetStatementParserByExprParser(exprParser)
	return x
}

func (sp *MySQLSeCharsetStatementParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SET) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.CHARSET, true)

	x := set.NewSetCharsetStatement(sp.DBType())
	name := parser.ParseExpr(sp.ExprParser())
	x.SetName(name)

	return x
}

/**
 *
 */
type MySQLSetNamesStatementParser struct {
	*parser.SQLSetNamesStatementParser
}

func NewSetNamesStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSetNamesStatementParser {
	x := new(MySQLSetNamesStatementParser)
	x.SQLSetNamesStatementParser = parser.NewSetNamesStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLSetNamesStatementParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SET) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.NAMES, true)

	x := set.NewSetNamesStatement(sp.DBType())
	name := parser.ParseExpr(sp.ExprParser())
	x.SetName(name)

	if sp.AcceptAndNextToken(parser.COLLATE) {
		parser.ParseExpr(sp.ExprParser())
	}

	return x
}

// --------------------------- Show Statement ---------------------------

type MySQLShowCreateDatabaseParser struct {
	*parser.SQLShowCreateDatabaseParser
}

func NewShowCreateDatabaseBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateDatabaseParser {
	return NewShowCreateDatabaseByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateDatabaseByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateDatabaseParser {
	return NewShowCreateDatabaseByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateDatabaseByExprParser(exprParser parser.ISQLExprParser) *MySQLShowCreateDatabaseParser {
	x := new(MySQLShowCreateDatabaseParser)
	x.SQLShowCreateDatabaseParser = parser.NewShowCreateDatabaseByExprParser(exprParser)
	return x
}
func (sp *MySQLShowCreateDatabaseParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := show.NewShowCreateDatabaseStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

type MySQLShowCreateEventParser struct {
	*parser.SQLShowCreateEventParser
}

func NewShowCreateEventBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateEventParser {
	return NewShowCreateEventByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateEventByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateEventParser {
	return NewShowCreateEventByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateEventByExprParser(exprParser parser.ISQLExprParser) *MySQLShowCreateEventParser {
	x := new(MySQLShowCreateEventParser)
	x.SQLShowCreateEventParser = parser.NewShowCreateEventByExprParser(exprParser)
	return x
}
func (sp *MySQLShowCreateEventParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.EVENT, true)

	x := show.NewShowCreateEventStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

type MySQLShowCreateFunctionParser struct {
	*parser.SQLShowCreateFunctionParser
}

func NewShowCreateFunctionBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateFunctionParser {
	return NewShowCreateFunctionByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateFunctionByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateFunctionParser {
	return NewShowCreateFunctionByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateFunctionByExprParser(exprParser parser.ISQLExprParser) *MySQLShowCreateFunctionParser {
	x := new(MySQLShowCreateFunctionParser)
	x.SQLShowCreateFunctionParser = parser.NewShowCreateFunctionByExprParser(exprParser)
	return x
}
func (sp *MySQLShowCreateFunctionParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.FUNCTION, true)

	x := show.NewShowCreateFunctionStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

type MySQLShowCreateProcedureParser struct {
	*parser.SQLShowCreateProcedureParser
}

func NewShowCreateProcedureBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateProcedureParser {
	return NewShowCreateProcedureByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateProcedureByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateProcedureParser {
	return NewShowCreateProcedureByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateProcedureByExprParser(exprParser parser.ISQLExprParser) *MySQLShowCreateProcedureParser {
	x := new(MySQLShowCreateProcedureParser)
	x.SQLShowCreateProcedureParser = parser.NewShowCreateProcedureByExprParser(exprParser)
	return x
}
func (sp *MySQLShowCreateProcedureParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.FUNCTION, true)

	x := show.NewShowCreateProcedureStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

type MySQLShowCreateTableParser struct {
	*parser.SQLShowCreateTableParser
}

func NewShowCreateTableBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateTableParser {
	return NewShowCreateTableByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateTableByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateTableParser {
	return NewShowCreateTableByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateTableByExprParser(exprParser parser.ISQLExprParser) *MySQLShowCreateTableParser {
	x := new(MySQLShowCreateTableParser)
	x.SQLShowCreateTableParser = parser.NewShowCreateTableByExprParser(exprParser)
	return x
}
func (sp *MySQLShowCreateTableParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := show.NewShowCreateProcedureStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

type MySQLShowCreateTriggerParser struct {
	*parser.SQLShowCreateTriggerParser
}

func NewShowCreateTriggerBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateTriggerParser {
	return NewShowCreateTriggerByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateTriggerByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateTriggerParser {
	return NewShowCreateTriggerByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateTriggerByExprParser(exprParser parser.ISQLExprParser) *MySQLShowCreateTriggerParser {
	x := new(MySQLShowCreateTriggerParser)
	x.SQLShowCreateTriggerParser = parser.NewShowCreateTriggerByExprParser(exprParser)
	return x
}
func (sp *MySQLShowCreateTriggerParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.TRIGGER, true)

	x := show.NewShowCreateProcedureStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

type MySQLShowCreateViewParser struct {
	*parser.SQLShowCreateViewParser
}

func NewShowCreateViewBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateViewParser {
	return NewShowCreateViewByLexer(NewLexer(sql), dbType, config)
}

func NewShowCreateViewByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowCreateViewParser {
	return NewShowCreateViewByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowCreateViewByExprParser(exprParser parser.ISQLExprParser) *MySQLShowCreateViewParser {
	x := new(MySQLShowCreateViewParser)
	x.SQLShowCreateViewParser = parser.NewShowCreateViewByExprParser(exprParser)
	return x
}
func (sp *MySQLShowCreateViewParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.CREATE, true)
	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	x := show.NewShowCreateProcedureStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	return x
}

type MySQLShowDatabasesParser struct {
	*parser.SQLShowDatabasesParser
}

func NewShowDatabaseBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowDatabasesParser {
	return NewShowDatabaseByLexer(NewLexer(sql), dbType, config)
}

func NewShowDatabaseByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLShowDatabasesParser {
	return NewShowDatabaseByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewShowDatabaseByExprParser(exprParser parser.ISQLExprParser) *MySQLShowDatabasesParser {
	x := new(MySQLShowDatabasesParser)
	x.SQLShowDatabasesParser = parser.NewShowDatabasesByExprParser(exprParser)
	return x
}
func (sp *MySQLShowDatabasesParser) Parse() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.SHOW) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.DATABASE, true)

	return nil
}

// --------------------------- EXPLAIN Statement Start ---------------------------
type MySQLDescStatementParser struct {
	*parser.SQLDescStatementParser
}

func NewDescStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLDescStatementParser {
	x := new(MySQLDescStatementParser)
	x.SQLDescStatementParser = parser.NewDescStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLDescStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.DESC, true)

	x := statement.NewDescStatement(sp.DBType())

	analyze := sp.AcceptAndNextToken(parser.ANALYZE)
	x.Analyze = analyze

	// explain_type
	var explainType expr.ISQLExpr
	if sp.Accept(parser.EXTENDED) {
		explainType = parser.ParseExpr(sp.ExprParser())

	} else if sp.Accept(parser.PARTITIONS) {
		explainType = parser.ParseExpr(sp.ExprParser())

	} else if sp.Accept(parser.FORMAT) {
		explainType = parser.ParseAssignExpr(sp.ExprParser())
	}
	x.SetExplainType(explainType)

	if sp.Accept(parser.SELECT) ||
		sp.Accept(parser.TABLE) ||
		sp.Accept(parser.DELETE) ||
		sp.Accept(parser.INSERT) ||
		sp.Accept(parser.REPLACE) ||
		sp.Accept(parser.UPDATE) {

		stmt := parser.ParseStatement(sp.ExprParser())
		x.SetStmt(stmt)

	} else if sp.AcceptAndNextToken(parser.FOR) {
		sp.AcceptAndNextTokenWithError(parser.CONNECTION, true)
		connectionId := parser.ParseExpr(sp.ExprParser())
		x.SetConnectionId(connectionId)

	} else if parser.IsIdentifier(sp.Kind()) {
		table := parser.ParseName(sp.ExprParser())
		x.SetTable(table)

	} else {

		panic(sp.UnSupport())
	}

	return x
}

type MySQLDescribeStatementParser struct {
	*parser.SQLDescribeStatementParser
}

func NewDescribeStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLDescribeStatementParser {
	x := new(MySQLDescribeStatementParser)
	x.SQLDescribeStatementParser = parser.NewDescribeStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLDescribeStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.DESCRIBE, true)

	x := statement.NewDescribeStatement(sp.DBType())
	analyze := sp.AcceptAndNextToken(parser.ANALYZE)
	x.Analyze = analyze

	// explain_type
	var explainType expr.ISQLExpr
	if sp.Accept(parser.EXTENDED) {
		explainType = parser.ParseExpr(sp.ExprParser())

	} else if sp.Accept(parser.PARTITIONS) {
		explainType = parser.ParseExpr(sp.ExprParser())

	} else if sp.Accept(parser.FORMAT) {
		explainType = parser.ParseAssignExpr(sp.ExprParser())
	}
	x.SetExplainType(explainType)

	if sp.Accept(parser.SELECT) ||
		sp.Accept(parser.TABLE) ||
		sp.Accept(parser.DELETE) ||
		sp.Accept(parser.INSERT) ||
		sp.Accept(parser.REPLACE) ||
		sp.Accept(parser.UPDATE) {

		stmt := parser.ParseStatement(sp.ExprParser())
		x.SetStmt(stmt)

	} else if sp.AcceptAndNextToken(parser.FOR) {
		sp.AcceptAndNextTokenWithError(parser.CONNECTION, true)
		connectionId := parser.ParseExpr(sp.ExprParser())
		x.SetConnectionId(connectionId)

	} else if parser.IsIdentifier(sp.Kind()) {
		table := parser.ParseName(sp.ExprParser())
		x.SetTable(table)

	} else {

		panic(sp.UnSupport())
	}

	return x
}

type MySQLExplainStatementParser struct {
	*parser.SQLExplainStatementParser
}

func NewExplainStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLExplainStatementParser {
	x := new(MySQLExplainStatementParser)
	x.SQLExplainStatementParser = parser.NewExplainStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLExplainStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.EXPLAIN, true)

	x := statement.NewExplainStatement(sp.DBType())

	analyze := sp.AcceptAndNextToken(parser.ANALYZE)
	x.Analyze = analyze

	// explain_type
	var explainType expr.ISQLExpr
	if sp.Accept(parser.EXTENDED) {
		explainType = parser.ParseExpr(sp.ExprParser())

	} else if sp.Accept(parser.PARTITIONS) {
		explainType = parser.ParseExpr(sp.ExprParser())

	} else if sp.Accept(parser.FORMAT) {
		explainType = parser.ParseAssignExpr(sp.ExprParser())
	}
	x.SetExplainType(explainType)

	if sp.Accept(parser.SELECT) ||
		sp.Accept(parser.TABLE) ||
		sp.Accept(parser.DELETE) ||
		sp.Accept(parser.INSERT) ||
		sp.Accept(parser.REPLACE) ||
		sp.Accept(parser.UPDATE) {

		stmt := parser.ParseStatement(sp.ExprParser())
		x.SetStmt(stmt)

	} else if sp.AcceptAndNextToken(parser.FOR) {
		sp.AcceptAndNextTokenWithError(parser.CONNECTION, true)
		connectionId := parser.ParseExpr(sp.ExprParser())
		x.SetConnectionId(connectionId)

	} else if parser.IsIdentifier(sp.Kind()) {
		table := parser.ParseName(sp.ExprParser())
		x.SetTable(table)

	} else {

		panic(sp.UnSupport())
	}

	return x
}

// --------------------------- EXPLAIN Statement End ---------------------------

// --------------------------- HELP Statement Start ---------------------------
type MySQLHelpStatementParser struct {
	*parser.SQLHelpStatementParser
}

func NewHelpStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLHelpStatementParser {
	x := new(MySQLHelpStatementParser)
	x.SQLHelpStatementParser = parser.NewHelpStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLHelpStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.HELP, true)

	x := statement.NewHelpStatement(sp.DBType())

	name := parser.ParseExpr(sp.ExprParser())
	x.SetName(name)

	return x
}

// --------------------------- HELP Statement End ---------------------------

// --------------------------- USE Statement Start ---------------------------
type MySQLUseStatementParser struct {
	*parser.SQLUseStatementParser
}

func NewUseStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLUseStatementParser {
	x := new(MySQLUseStatementParser)
	x.SQLUseStatementParser = parser.NewUseStatementParserByExprParser(exprParser)
	return x
}
func (sp *MySQLUseStatementParser) Parse() statement.ISQLStatement {
	sp.AcceptAndNextTokenWithError(parser.USE, true)

	x := statement.NewUseStatement(sp.DBType())

	name := parser.ParseExpr(sp.ExprParser())
	x.SetName(name)

	return x
}

// --------------------------- USE Statement End ---------------------------
