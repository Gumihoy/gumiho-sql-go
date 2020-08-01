package parser

import (
	"fmt"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	le := NewLexer("select * from table")
	fmt.Println(le)
}
func TestIsSQLIdentifierPart(t *testing.T) {
	//fmt.Println(issql('@'))
}
