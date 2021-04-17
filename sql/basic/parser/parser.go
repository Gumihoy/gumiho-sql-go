package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strconv"
)

type SQLParseConfig struct {
	SkipComment bool
}

func NewParseConfig() *SQLParseConfig {
	x := new(SQLParseConfig)
	x.SkipComment = true
	return x
}

type ISQLParser interface {
	DBType() db.Type

	Config() *SQLParseConfig
	SetConfig(config *SQLParseConfig)

	Lexer() ISQLLexer
	Token() *SQLToken
	Kind() *Kind

	Mark() *Mark
	Reset() bool
	ResetWithMark(mark *Mark) bool

	Accept(kind *Kind) bool
	AcceptWithError(kind *Kind, error bool) bool

	AcceptAndNextToken(kind *Kind) bool
	AcceptAndNextTokenWithError(kind *Kind, error bool) bool

	StringValue() string

	Comments() []ast.ISQLComment
	ClearComments()

	SyntaxError() string
	UnSupport() string
}

type SQLParser struct {
	dbType db.Type
	lexer  ISQLLexer
	config *SQLParseConfig
}

func NewParserBySQL(sourceSQL string, dbType db.Type, config *SQLParseConfig) *SQLParser {
	return NewParserByLexer(NewLexer(sourceSQL), dbType, config)
}

func NewParserByLexer(lexer ISQLLexer, dbType db.Type, config *SQLParseConfig) *SQLParser {
	x := new(SQLParser)
	x.lexer = lexer
	x.dbType = dbType
	x.config = config
	return x
}

func (sp *SQLParser) DBType() db.Type {
	return sp.dbType
}

func (sp *SQLParser) Lexer() ISQLLexer {
	return sp.lexer
}

func (sp *SQLParser) SetLexer(lexer ISQLLexer) {
	sp.lexer = lexer
}

func (sp *SQLParser) Config() *SQLParseConfig {
	return sp.config
}

func (sp *SQLParser) SetConfig(config *SQLParseConfig) {
	sp.config = config
}

func (sp *SQLParser) Token() *SQLToken {
	return sp.lexer.Token()
}

func (sp *SQLParser) Kind() *Kind {
	return sp.Token().Kind
}

func (sp *SQLParser) Mark() *Mark {
	return sp.lexer.Mark()
}

func (sp *SQLParser) Reset() bool {
	return sp.lexer.Reset()
}

func (sp *SQLParser) ResetWithMark(mark *Mark) bool {
	return sp.lexer.ResetWithMark(mark)
}

func (sp *SQLParser) Accept(kind *Kind) bool {
	return sp.AcceptWithError(kind, false)
}

func (sp *SQLParser) AcceptWithError(kind *Kind, error bool) bool {
	if sp.Kind() == kind ||
		(sp.Kind() == IDENTIFIER &&
			kind.Equals(sp.StringValue())) {
		return true
	}

	if error {
		sp.acceptError(kind)
	}
	return false
}

func (sp *SQLParser) AcceptAndNextToken(kind *Kind) bool {
	return sp.AcceptAndNextTokenWithError(kind, false)
}

func (sp *SQLParser) AcceptAndNextTokenWithError(kind *Kind, error bool) bool {
	if sp.Kind() == kind ||
		(sp.Kind() == IDENTIFIER &&
			kind.Equals(sp.StringValue())) {
		NextTokenByParser(sp)
		return true
	}
	if error {
		sp.acceptError(kind)
	}
	return false
}

func (sp *SQLParser) acceptError(kind *Kind) {
	if sp.Kind() == IDENTIFIER {
		panic("Syntax Error: " + sp.Token().Error() +
			". expected " + kind.String() + ", actual " + sp.Kind().String() + ": " + sp.StringValue() +
			", for the right syntax to use near: \"" + sp.lexer.UseNearErrorMsg())
	} else {
		panic("Syntax Error: " + sp.Token().Error() +
			". expected " + kind.String() + ", actual " + sp.Kind().String() +
			", for the right syntax to use near: \"" + sp.lexer.UseNearErrorMsg())
	}
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
// 		panic("Syntax Error: expected " + parser.SQLToken().Error() + ", actual " + parser.StringValue())
// 	}
// 	return false
// }

func NextTokenByParser(parser ISQLParser) {
	NextToken(parser.Lexer())
}

func (sp *SQLParser) StringValue() string {
	return sp.lexer.StringValue()
}

func (sp *SQLParser) IntegerValue() int64 {
	value, error := strconv.ParseInt(sp.StringValue(), 10, 64)
	if error != nil {
		panic("")
	}
	return value
}

func (sp *SQLParser) Comments() []ast.ISQLComment {
	return sp.Lexer().Comments()
}
func (sp *SQLParser) ClearComments() {
	sp.Lexer().ClearComments()
}

func (sp *SQLParser) SyntaxError() string {
	panic("Syntax Error: " + sp.Token().Error() + ", SQL: \"" + sp.lexer.UseNearErrorMsg() + "\"")
}
func (sp *SQLParser) UnSupport() string {
	return sp.lexer.UnSupport()
}

func IsIdentifier(kind *Kind) bool {
	if kind == IDENTIFIER ||
		kind == IDENTIFIER_DOUBLE_QUOTE ||
		kind == IDENTIFIER_REVERSE_QUOTE {
		return true
	}
	return false
}

func AddBeforeComments(sp ISQLParser, object ast.ISQLObject) {
	if sp == nil || object == nil {
		return
	}
	object.AddBeforeComments(sp.Lexer().Comments())
	sp.Lexer().ClearComments()
}
func AddAfterComments(sp ISQLParser, object ast.ISQLObject) {
	if sp == nil || object == nil {
		return
	}
	object.AddAfterComments(sp.Lexer().Comments())
	sp.Lexer().ClearComments()
}
