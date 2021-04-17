package parser

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	le := NewLexer("select * from table")
	fmt.Println(le)
}
func TestIsSQLIdentifierPart(t *testing.T) {
	fmt.Println(strconv.ParseFloat("+6.34", 64))
}
