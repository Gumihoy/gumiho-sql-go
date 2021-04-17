package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
)

type MariaDBLexer struct {
	*parser.SQLLexer
}

func NewLexer(sql string) *MariaDBLexer {
	var l MariaDBLexer
	l.SQLLexer = parser.NewLexer(sql)
	return &l
}
