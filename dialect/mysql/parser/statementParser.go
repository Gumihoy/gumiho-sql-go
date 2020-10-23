package parser

import (
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/basic/parser"
)


type MySQLStatementParser struct {
	*parser.SQLStatementParser
}

func NewStatementParserBySQL(sql string) *MySQLStatementParser {
	return NewStatementParserByLexer(NewLexer(sql))
}

func NewStatementParserByLexer(lexer parser.ISQLLexer) *MySQLStatementParser {
	var x MySQLStatementParser
	x.SQLStatementParser = parser.NewStatementParserByExprParser(NewExprParserByLexer(lexer))
	return &x
}
func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLStatementParser {
	var x MySQLStatementParser
	x.SQLStatementParser = parser.NewStatementParserByExprParser(exprParser)
	return &x
}




type MySQLSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func (x *MySQLSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != parser.SELECT {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewSelectStatementParserBySQL(sql string) *MySQLSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql))
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer) *MySQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *MySQLSelectStatementParser {
	return &MySQLSelectStatementParser{parser.NewSelectStatementParserByExprParser(exprParser)}
}







