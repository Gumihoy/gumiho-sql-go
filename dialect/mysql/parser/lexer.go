package parser

import (
	"gumihoy.com/sql/basic/parser"
	"gumihoy.com/sql/basic/parser/kind"
)

type Lexer struct {
	*parser.Lexer
}

func NewLexer(sql string) *Lexer {
	var l Lexer
	l.Lexer = parser.NewLexer(sql)
	return &l
}

func (lexer *Lexer) ScanReverseQuota() kind.Kind {
	lexer.ScanRune()
	for lexer.Ch() != '`' || lexer.IsEOF() {
		lexer.AppendValue()
		lexer.ScanRune()
	}
	return kind.QUOTE_REVERSE
}
