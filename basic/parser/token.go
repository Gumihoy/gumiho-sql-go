package parser

import (
	"gumihoy.com/sql/basic/parser/kind"
	"strconv"
)

var KindMap = make(map[string]kind.Kind)

func init() {
	KindMap[string(kind.ALTER)] = kind.SELECT
	KindMap[string(kind.SELECT)] = kind.SELECT
	KindMap[string(kind.CREATE)] = kind.CREATE
	KindMap[string(kind.UPDATE)] = kind.UPDATE
}

type Token struct {
	Kind       kind.Kind
	start, end int
	line, col  int
}

func (token *Token) String() string {

	return "start " + strconv.Itoa(token.start) +
		", end " + strconv.Itoa(token.end) +
		string(token.Kind)
}

func NewToken(start int, end int, kind kind.Kind) *Token {
	return &Token{start: start, end: end, Kind: kind}
}

func NewTokenByLexer(lexer *Lexer, kind kind.Kind) *Token {
	return &Token{start: lexer.line, end: lexer.col, Kind: kind}
}
