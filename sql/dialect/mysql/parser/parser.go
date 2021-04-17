package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type MySQLParser struct {
	*parser.SQLParser
}

func NewParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLParser {
	return &MySQLParser{parser.NewParserBySQL(sql, dbType, config)}
}

func NewParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLParser {
	return &MySQLParser{parser.NewParserByLexer(lexer, dbType, config)}
}
