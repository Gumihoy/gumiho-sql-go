package parser

import (
	"gumihoy.com/sql/basic/parser"
	mysqlParser "gumihoy.com/sql/dialect/mysql/parser"
)

type MariaDBParser struct {
	*mysqlParser.MySQLParser
}

func NewParserBySQL(sql string) *MariaDBParser {
	return &MariaDBParser{mysqlParser.NewParserBySQL(sql)}
}

func NewParserByLexer(lexer parser.ISQLLexer) *MariaDBParser {
	return &MariaDBParser{mysqlParser.NewParserByLexer(lexer)}
}
