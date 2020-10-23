package parser

import (
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/basic/parser"
)

type PostgreSQLStatementParser struct {
	*parser.SQLStatementParser
}

func NewStatementParserBySQL(sql string) *PostgreSQLStatementParser {
	return NewStatementParserByLexer(NewLexer(sql))
}

func NewStatementParserByLexer(lexer parser.ISQLLexer) *PostgreSQLStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *PostgreSQLStatementParser {
	var x PostgreSQLStatementParser
	x.SQLStatementParser = parser.NewStatementParserByExprParser(exprParser)
	return &x
}

// DML

type PostgreSQLSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func (x *PostgreSQLSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != parser.SELECT {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewSelectStatementParserBySQL(sql string) *PostgreSQLSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql))
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer) *PostgreSQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *PostgreSQLSelectStatementParser {
	return &PostgreSQLSelectStatementParser{parser.NewSelectStatementParserByExprParser(exprParser)}
}
