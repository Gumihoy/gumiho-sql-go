package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type MariaDBLexer struct {
	*parser.SQLLexer
}

func NewLexer(sql string) *MariaDBLexer {
	var l MariaDBLexer
	l.SQLLexer = parser.NewLexer(sql)
	return &l
}
