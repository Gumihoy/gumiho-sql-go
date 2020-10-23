package parser

import (
	"gumihoy.com/sql/basic/parser"
	mysqlParser "gumihoy.com/sql/dialect/mysql/parser"
)

type MariaDBExprParser struct {
	*mysqlParser.MySQLExprParser
}

func NewExprParserBySQL(sql string) *MariaDBExprParser {
	return NewExprParserByLexer(NewLexer(sql))
}

func NewExprParserByLexer(lexer parser.ISQLLexer) *MariaDBExprParser {
	x := new(MariaDBExprParser)
	x.MySQLExprParser = mysqlParser.NewExprParserByLexer(lexer)
	return x
}
