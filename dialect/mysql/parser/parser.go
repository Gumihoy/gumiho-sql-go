package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type Parser struct {
	*parser.Parser
}

func NewParserBySQL(sql string) parser.IParser {
	parser := Parser{parser.NewParserByLexer(NewLexer(sql))}
	return parser
}
