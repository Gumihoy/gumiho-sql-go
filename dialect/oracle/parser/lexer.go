package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type OracleLexer struct {
	*parser.SQLLexer
}

func NewLexer(sql string) *OracleLexer {
	var x OracleLexer
	x.SQLLexer = parser.NewLexer(sql)
	return &x
}


// func (lexer *OracleLexer) IsSQLIdentifierStart() bool {
// 	return !parser.IsWhitespace(lexer.Ch()) &&
// 		lexer.Ch() != '"' &&
// 		lexer.Ch() != '\'' &&
// 		lexer.Ch() != '`'
// }
//
// func (lexer *OracleLexer) IsSQLIdentifierPart() bool {
// 	return lexer.Ch() != '.' &&
// 		!parser.IsWhitespace(lexer.Ch())
// }

func (lexer *OracleLexer) ScanDoubleQuota() *parser.Kind {

	for {
		lexer.ScanString()
		if lexer.Ch() == '"' {

		}
	}
	return parser.IDENTIFIER_DOUBLE_QUOTE
}
