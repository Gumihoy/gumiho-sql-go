package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	mysqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/parser"
)

type MariaDBStatementParser struct {
	*mysqlParser.MySQLStatementParser
}

func NewStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MariaDBStatementParser {
	return NewStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MariaDBStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *MariaDBStatementParser {
	var x MariaDBStatementParser
	x.MySQLStatementParser = mysqlParser.NewStatementParserByExprParser(exprParser)
	return &x
}

type MariaDBSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func (x *MariaDBSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != parser.SELECT {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(x.DBType(), query)
}

func NewSelectStatementParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MariaDBSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql), dbType, config)
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MariaDBSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer, dbType, config))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *MariaDBSelectStatementParser {
	return &MariaDBSelectStatementParser{parser.NewSelectStatementParserByExprParser(exprParser)}
}
