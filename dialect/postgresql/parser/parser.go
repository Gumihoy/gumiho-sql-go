package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type PostgreSQLParser struct {
	*parser.SQLParser
}

func NewParserBySQL(sql string) *PostgreSQLParser {
	return &PostgreSQLParser{parser.NewParserBySQL(sql)}
}

func NewParserByLexer(lexer parser.ISQLLexer) *PostgreSQLParser {
	return &PostgreSQLParser{parser.NewParserByLexer(lexer)}
}
