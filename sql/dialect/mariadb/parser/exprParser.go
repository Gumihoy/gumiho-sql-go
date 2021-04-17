package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	mysqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/parser"
)

type MariaDBExprParser struct {
	*mysqlParser.MySQLExprParser
}

func NewExprParserBySQL(sourceSQL string, dbType db.Type, config *parser.SQLParseConfig) *MariaDBExprParser {
	return NewExprParserByLexer(NewLexer(sourceSQL), dbType, config)
}

func NewExprParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MariaDBExprParser {
	x := new(MariaDBExprParser)
	x.MySQLExprParser = mysqlParser.NewExprParserByLexer(lexer, dbType, config)
	return x
}
