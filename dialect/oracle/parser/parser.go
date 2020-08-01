package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type Parser struct {
	*parser.Parser
}

func NewParserBySQL(sql string) Parser {
	var parser Parser

	return parser
}

func ParseExpr() {

}
