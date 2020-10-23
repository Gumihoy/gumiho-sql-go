package parser

import (
	"fmt"
	"testing"
)

func TestParseExpr(t *testing.T) {
	le := NewLexer("select * from table")
	fmt.Println(le)
}
