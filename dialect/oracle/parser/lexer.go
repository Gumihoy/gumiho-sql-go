package parser

import (
	"fmt"
	"gumihoy.com/sql/basic/parser"
)

type Lexer struct {
	parser.Lexer
}

func (lexer *Lexer) scanBQuota() {
	fmt.Println("[ORACLE]scanBQuota...")
}
