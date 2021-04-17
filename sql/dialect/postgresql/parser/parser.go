package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type PostgreSQLParser struct {
	*parser.SQLParser
}

func NewParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLParser {
	return NewParserByLexer(NewLexer(sql), dbType, config)
}

func NewParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLParser {
	return &PostgreSQLParser{parser.NewParserByLexer(lexer, dbType, config)}
}
