package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	exprTable "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/database"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/schema"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	mysqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/parser"
)

type TiDBStatementParser struct {
	*mysqlParser.MySQLStatementParser
}

func NewStatementParserBySQL(sourceSQL string, dbType db.Type, config *parser.SQLParseConfig) *TiDBStatementParser {
	return NewStatementParserByLexer(NewLexer(sourceSQL), dbType, config)
}

func NewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}
func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBStatementParser {
	x := new(TiDBStatementParser)
	x.MySQLStatementParser = mysqlParser.NewStatementParserByExprParser(exprParser)
	return x
}

// ------------------------------------------------------------- DML -------------------------------------------------------------
// ------------------------ Select ------------------------

type TiDBSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func NewSelectStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBSelectStatementParser {
	x := new(TiDBSelectStatementParser)
	x.SQLSelectStatementParser = parser.NewSelectStatementParserByExprParser(exprParser)
	return x
}

func (x *TiDBSelectStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.WITH) && !x.Accept(parser.SELECT) {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(x.DBType(), query)
}

// ------------------------ Delete ------------------------

type TiDBDeleteStatementParser struct {
	*parser.SQLDeleteStatementParser
}

func NewDeleteStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBDeleteStatementParser {
	return NewDeleteStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDeleteStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDeleteStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBDeleteStatementParser {
	x := new(TiDBDeleteStatementParser)
	x.SQLDeleteStatementParser = parser.NewDeleteStatementParserByExprParser(exprParser)
	return x
}

func (x *TiDBDeleteStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.DELETE) {
		return nil
	}
	return nil
}

// ------------------------ Insert ------------------------

type TiDBInsertStatementParser struct {
	*parser.SQLInsertStatementParser
}

func (x *TiDBInsertStatementParser) Parse() statement.ISQLStatement {
	if !x.Accept(parser.INSERT) {
		return nil
	}

	return nil
}

func NewInsertStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBInsertStatementParser {
	return NewInsertStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewInsertStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBInsertStatementParser {
	return NewInsertStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewInsertStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBInsertStatementParser {
	x := new(TiDBInsertStatementParser)
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
type TiDBDatabaseStatementParser struct {
	*mysqlParser.MySQLDatabaseStatementParser
}

func NewDatabaseStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBDatabaseStatementParser {
	return NewDatabaseStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewDatabaseStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBDatabaseStatementParser {
	return NewDatabaseStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewDatabaseStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBDatabaseStatementParser {
	x := new(TiDBDatabaseStatementParser)
	x.SQLDatabaseStatementParser = parser.NewDatabaseStatementParserByExprParser(exprParser)
	return x
}

func (p *TiDBDatabaseStatementParser) ParseAlter() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.ALTER, true)
	p.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewAlterDatabaseStatement(p.DBType())

	return x
}

//
func (p *TiDBDatabaseStatementParser) ParseCreate() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.CREATE, true)
	p.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewCreateDatabaseStatement(p.DBType())

	ifNotExists := parser.ParseIfNotExists(p.ExprParser())
	x.IfNotExists = ifNotExists

	name := parser.ParseName(p.ExprParser())
	x.SetName(name)

	return x

}
func (p *TiDBDatabaseStatementParser) ParseDrop() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.DROP, true)
	p.AcceptAndNextTokenWithError(parser.DATABASE, true)

	x := database.NewDropDatabaseSStatement(p.DBType())

	ifExists := parser.ParseIfExists(p.ExprParser())
	x.IfExists = ifExists

	name := parser.ParseName(p.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ Schema ------------------------
type TiDBSchemaStatementParser struct {
	*mysqlParser.MySQLSchemaStatementParser
}

func NewSchemaStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBSchemaStatementParser {
	return NewSchemaStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSchemaStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBSchemaStatementParser {
	return NewSchemaStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSchemaStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBSchemaStatementParser {
	x := new(TiDBSchemaStatementParser)
	x.MySQLSchemaStatementParser = mysqlParser.NewSchemaStatementParserByExprParser(exprParser)
	return x
}

func (x *TiDBSchemaStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(parser.ALTER) {
		return nil
	}

	return nil
}

//
func (x *TiDBSchemaStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(parser.CREATE) {
		return nil
	}

	return nil
}
func (p *TiDBSchemaStatementParser) ParseDrop() statement.ISQLStatement {

	p.AcceptAndNextTokenWithError(parser.DROP, true)
	p.AcceptAndNextTokenWithError(parser.SCHEMA, true)

	x := schema.NewDropSchemaStatement(p.DBType())

	ifExists := parser.ParseIfExists(p.ExprParser())
	x.IfExists = ifExists

	name := parser.ParseName(p.ExprParser())
	x.SetName(name)

	return x
}

// ------------------------ OnTable ------------------------
type TiDBTableStatementParser struct {
	*mysqlParser.MySQLTableStatementParser
}

func NewTableStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBTableStatementParser {
	return NewTableStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTableStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBTableStatementParser {
	return NewTableStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTableStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBTableStatementParser {
	x := new(TiDBTableStatementParser)
	x.MySQLTableStatementParser = mysqlParser.NewTableStatementParserByExprParser(exprParser)
	return x
}

/**
 * ALTER TABLE tbl_name
    [alter_option [, alter_option] ...]
    [partition_options]
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 */
func (sp *TiDBTableStatementParser) ParseAlter() statement.ISQLStatement {

	sp.AcceptAndNextTokenWithError(parser.ALTER, true)
	sp.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := table.NewAlterTableStatement(sp.DBType())
	name := parser.ParseName(sp.ExprParser())
	x.SetName(name)

	for {
		action := sp.ExprParser().ParseAlterTableAction(sp.ExprParser())
		x.AddAction(action)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) || action == nil {
			break
		}
	}

	return x
}

//
// https://dev.mysql.com/doc/refman/8.0/en/create-table.html
func (sp *TiDBTableStatementParser) ParseCreate() statement.ISQLStatement {

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
func (sp *TiDBTableStatementParser) IsParseOption() bool {
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
func (sp *TiDBTableStatementParser) ParseOption() expr.ISQLExpr {
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

func (p *TiDBTableStatementParser) ParseDrop() statement.ISQLStatement {

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



// ------------------------ View ------------------------
type TiDBViewStatementParser struct {
	*mysqlParser.MySQLViewStatementParser
}
func NewViewStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBViewStatementParser {
	return NewViewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewViewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBViewStatementParser {
	return NewViewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewViewStatementParserByExprParser(exprParser parser.ISQLExprParser) *TiDBViewStatementParser {
	x := new(TiDBViewStatementParser)
	x.SQLViewStatementParser = parser.NewViewStatementParserByExprParser(exprParser)
	return x
}

func (x *TiDBViewStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.AcceptAndNextToken(parser.ALTER) {
		return nil
	}

	return nil
}
func (sp *TiDBViewStatementParser) ParseCreate() statement.ISQLStatement {
	if !sp.Accept(parser.CREATE) {
		return nil
	}

	return nil
}

/**
 * DROP VIEW [IF EXISTS] view_name [, view_name] ... [RESTRICT | CASCADE]
 * https://dev.mysql.com/doc/refman/8.0/en/drop-view.html
 */
func (sp *TiDBViewStatementParser) ParseDrop() statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}
	sp.AcceptAndNextTokenWithError(parser.VIEW, true)

	x := view.NewDropViewStatement(sp.DBType())

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	for  {
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