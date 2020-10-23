package parser

import (
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/db"
	"strconv"
)

type Config struct {
	DBType db.DBType
}

func NewConfig(dbType db.DBType) Config {
	var c Config
	c.DBType = dbType
	return c
}

type ISQLParser interface {
	Lexer() ISQLLexer
	Token() *Token
	Kind() *Kind

	Accept(kind *Kind) bool
	AcceptWithError(kind *Kind, error bool) bool

	AcceptAndNextToken(kind *Kind) bool
	AcceptAndNextTokenWithError(kind *Kind, error bool) bool

	StringValue() string
}

type SQLParser struct {
	lexer ISQLLexer
}

func (parser *SQLParser) Lexer() ISQLLexer {
	return parser.lexer
}

func (parser *SQLParser) SetLexer(lexer ISQLLexer) {
	parser.lexer = lexer
}

func (parser *SQLParser) Token() *Token {
	return parser.lexer.Token()
}

func (parser *SQLParser) Kind() *Kind {
	return parser.Token().Kind
}

func (x *SQLParser) Accept(kind *Kind) bool {
	return x.AcceptWithError(kind, false)
}

func (x *SQLParser) AcceptWithError(kind *Kind, error bool) bool {
	if x.Kind() == kind ||
		(x.Kind() == IDENTIFIER &&
			kind.Equals(x.StringValue())) {
		return true
	}

	if error {
		panic("Syntax Error: expected " + x.Token().Error() + ", actual " + x.StringValue())
	}
	return false
}

func (x *SQLParser) AcceptAndNextToken(kind *Kind) bool {
	return x.AcceptAndNextTokenWithError(kind, false)
}

func (x *SQLParser) AcceptAndNextTokenWithError(kind *Kind, error bool) bool {
	if x.Kind() == kind ||
		(x.Kind() == IDENTIFIER &&
			kind.Equals(x.StringValue())) {
		NextTokenByParser(x)
		return true
	}
	if error {
		panic("Syntax Error: expected " + x.Token().Error() + ", actual " + x.StringValue())
	}
	return false
}

// func AcceptAndNextToken(parser ISQLParser, kind *Kind) bool {
// 	return AcceptAndNextTokenWithError(parser, kind, false)
// }
// func AcceptAndNextTokenWithError(parser ISQLParser, kind *Kind, error bool) bool {
// 	if parser.Kind() == kind ||
// 		(parser.Kind() == IDENTIFIER &&
// 			kind.Equals(parser.StringValue())) {
// 		NextTokenByParser(parser)
// 		return true
// 	}
// 	if error {
// 		panic("Syntax Error: expected " + parser.Token().Error() + ", actual " + parser.StringValue())
// 	}
// 	return false
// }

func NextTokenByParser(parser ISQLParser) {
	NextToken(parser.Lexer())
}

func (parser *SQLParser) StringValue() string {
	return parser.lexer.StringValue()
}

func (parser *SQLParser) IntegerValue() int64 {
	value, error := strconv.ParseInt(parser.StringValue(), 10, 64)
	if error != nil {
		panic("")
	}
	return value
}

func (parser *SQLParser) ParseCreateStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *SQLParser) ParseDropStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *SQLParser) ParseDeleteStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *SQLParser) ParseInsertStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *SQLParser) ParseUpdateStatement() statement.ISQLStatement {
	panic("implement me")
}

func IsIdentifier(kind *Kind) bool {
	if kind == IDENTIFIER ||
		kind == IDENTIFIER_DOUBLE_QUOTE ||
		kind == IDENTIFIER_REVERSE_QUOTE {
		return true
	}
	return false
}

func NewParserBySQL(sql string) *SQLParser {
	return NewParserByLexer(NewLexer(sql))
}

func NewParserByLexer(lexer ISQLLexer) *SQLParser {
	return &SQLParser{lexer}
}
