package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type OracleParser struct {
	*parser.SQLParser
}

func NewParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleParser {
	var x OracleParser
	x.SQLParser = parser.NewParserBySQL(sql, dbType, config)
	return &x
}

func NewParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleParser {
	return &OracleParser{parser.NewParserByLexer(lexer, dbType, config)}
}
