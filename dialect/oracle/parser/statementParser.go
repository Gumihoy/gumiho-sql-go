package parser

import (
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/basic/parser"
)



type OracleStatementParser struct {
	*parser.SQLStatementParser
}

func NewStatementParserBySQL(sql string) *OracleStatementParser {
	return NewStatementParserByLexer(NewLexer(sql))
}

func NewStatementParserByLexer(lexer parser.ISQLLexer) *OracleStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer))
}


func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleStatementParser {
	var x OracleStatementParser
	x.SQLStatementParser = parser.NewStatementParserByExprParser(exprParser)
	return &x
}


type OracleSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func (x *OracleSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != parser.SELECT {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewSelectStatementParserBySQL(sql string) *OracleSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql))
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer) *OracleSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *OracleSelectStatementParser {
	return &OracleSelectStatementParser{parser.NewSelectStatementParserByExprParser(exprParser)}
}







