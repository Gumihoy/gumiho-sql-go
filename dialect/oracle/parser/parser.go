package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type OracleParser struct {
	*parser.SQLParser
}

func NewParserBySQL(sql string) *OracleParser {
	var x OracleParser
	x.SQLParser = parser.NewParserBySQL(sql)
	return &x
}

func NewParserByLexer(lexer parser.ISQLLexer) *OracleParser {
	return &OracleParser{parser.NewParserByLexer(lexer)}
}
