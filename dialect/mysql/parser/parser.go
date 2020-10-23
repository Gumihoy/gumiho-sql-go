package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type MySQLParser struct {
	*parser.SQLParser
}

func NewParserBySQL(sql string) *MySQLParser {
	return &MySQLParser{parser.NewParserBySQL(sql)}
}

func NewParserByLexer(lexer parser.ISQLLexer) *MySQLParser {
	return &MySQLParser{parser.NewParserByLexer(lexer)}
}
