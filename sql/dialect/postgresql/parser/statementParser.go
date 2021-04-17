package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type PostgreSQLStatementParser struct {
	*parser.SQLStatementParser
}

func NewStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLStatementParser {
	return NewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *PostgreSQLStatementParser {
	var x PostgreSQLStatementParser
	x.SQLStatementParser = parser.NewStatementParserByExprParser(exprParser)
	return &x
}

// DML

type PostgreSQLSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func (x *PostgreSQLSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != parser.SELECT {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(x.DBType(), query)
}

func NewSelectStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *PostgreSQLSelectStatementParser {
	return &PostgreSQLSelectStatementParser{parser.NewSelectStatementParserByExprParser(exprParser)}
}

// ------------------------------------------------------------- DDL -------------------------------------------------------------
// ------------------------ Table ------------------------
type PostgreSQLTableStatementParser struct {
	*parser.SQLTableStatementParser
}

func NewTableStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLTableStatementParser {
	return NewTableStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewTableStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLTableStatementParser {
	return NewTableStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewTableStatementParserByExprParser(exprParser parser.ISQLExprParser) *PostgreSQLTableStatementParser {
	x := new(PostgreSQLTableStatementParser)
	x.SQLTableStatementParser = parser.NewTableStatementParserByExprParser(exprParser)
	return x
}

/**
 * ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
    action [, ... ]
ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
    RENAME [ COLUMN ] column_name TO new_column_name
ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
    RENAME CONSTRAINT constraint_name TO new_constraint_name
ALTER TABLE [ IF EXISTS ] name
    RENAME TO new_name
ALTER TABLE [ IF EXISTS ] name
    SET SCHEMA new_schema
ALTER TABLE ALL IN TABLESPACE name [ OWNED BY role_name [, ... ] ]
    SET TABLESPACE new_tablespace [ NOWAIT ]
ALTER TABLE [ IF EXISTS ] name
    ATTACH PARTITION partition_name { FOR VALUES partition_bound_spec | DEFAULT }
ALTER TABLE [ IF EXISTS ] name
    DETACH PARTITION partition_name

 * https://www.postgresql.org/docs/devel/sql-altertable.html
 */
func (sp *PostgreSQLTableStatementParser) ParseAlter() statement.ISQLStatement {

	sp.AcceptAndNextTokenWithError(parser.ALTER, true)
	sp.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := table.NewAlterTableStatement(sp.DBType())

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	only := sp.AcceptAndNextToken(parser.ONLY)
	x.Only = only

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
func (sp *PostgreSQLTableStatementParser) ParseCreate() statement.ISQLStatement {

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
func (sp *PostgreSQLTableStatementParser) IsParseOption() bool {
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
func (sp *PostgreSQLTableStatementParser) ParseOption() expr.ISQLExpr {
	return nil
}

/**
 *
 */
func (sp *PostgreSQLTableStatementParser) ParseDrop() statement.ISQLStatement {

	sp.AcceptAndNextTokenWithError(parser.DROP, true)
	temporary := sp.AcceptAndNextToken(parser.TEMPORARY)
	sp.AcceptAndNextTokenWithError(parser.TABLE, true)

	x := table.NewDropTableStatement(sp.DBType())
	x.Temporary = temporary

	ifExists := parser.ParseIfExists(sp.ExprParser())
	x.IfExists = ifExists

	for {
		name := parser.ParseName(sp.ExprParser())
		if name == nil {
			panic(sp.SyntaxError())
		}
		x.AddName(name)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	return x
}
