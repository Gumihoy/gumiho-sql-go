package util

import (
	"fmt"
	"gumihoy.com/sql/basic/parser"
	"gumihoy.com/sql/dbtype"
	"testing"
)

func TestParseStatementsByDBType(t *testing.T) {
	p := createParser("select * from table", parser.NewConfig(dbtype.MySQL))
	fmt.Println(parser.ParseStatements(p))
}

func TestParseStatementsByDBType1(t *testing.T) {
	var ch rune = 96
	switch ch {
	case '"':
		fmt.Println('"')
		break
	case '`':
		fmt.Println("`")
		break
	}
}
