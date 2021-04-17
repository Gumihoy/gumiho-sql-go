package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	mysqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/parser"
)

type TiDBParser struct {
	*mysqlParser.MySQLParser
}

func NewParserBySQL(sourceSQL string, dbType db.Type, config *parser.SQLParseConfig) *TiDBParser {
	return NewParserByLexer(NewLexer(sourceSQL), dbType, config)
}

func NewParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBParser {
	x := new(TiDBParser)
	x.MySQLParser = mysqlParser.NewParserByLexer(lexer, dbType, config)
	return x
}
