package parser

import (
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/basic/parser"
	mysqlParser "gumihoy.com/sql/dialect/mysql/parser"
)

type MariaDBStatementParser struct {
	*mysqlParser.MySQLStatementParser
}

func NewStatementParserBySQL(sql string) *MariaDBStatementParser {
	return NewStatementParserByLexer(NewLexer(sql))
}

func NewStatementParserByLexer(lexer parser.ISQLLexer) *MariaDBStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewStatementParserByExprParser(exprParser parser.ISQLExprParser) *MariaDBStatementParser {
	var x MariaDBStatementParser
	x.MySQLStatementParser = mysqlParser.NewStatementParserByExprParser(exprParser)
	return &x
}


type MariaDBSelectStatementParser struct {
	*parser.SQLSelectStatementParser
}

func (x *MariaDBSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != parser.SELECT {
		return nil
	}
	query := parser.ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewSelectStatementParserBySQL(sql string) *MariaDBSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql))
}

func NewSelectStatementParserByLexer(lexer parser.ISQLLexer) *MariaDBSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewSelectStatementParserByExprParser(exprParser parser.ISQLExprParser) *MariaDBSelectStatementParser {
	return &MariaDBSelectStatementParser{parser.NewSelectStatementParserByExprParser(exprParser)}
}







