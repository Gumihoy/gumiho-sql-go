package parser

import (
	"fmt"
	"gumihoy.com/sql/basic/parser"
)

type PostgreSQLLexer struct {
	*parser.SQLLexer
}

func NewLexer(sql string) *PostgreSQLLexer {
	return &PostgreSQLLexer{parser.NewLexer(sql)}
}

func (lexer *PostgreSQLLexer) scanBQuota() {
	fmt.Println("[ORACLE]scanBQuota...")
}
