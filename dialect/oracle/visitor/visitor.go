package visitor

import (
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	"strings"
)

type IOracleVisitor interface {
	visitor.ISQLVisitor
}

type OracleVisitorAdapter struct {
	*visitor.SQLVisitorAdapter
}

type OracleOutputVisitor struct {
	*visitor.SQLOutputVisitor
}

func NewOutputVisitor(builder *strings.Builder, config config.Output) *OracleOutputVisitor {
	x := new(OracleOutputVisitor)
	x.SQLOutputVisitor = visitor.NewOutputVisitor(builder, config)
	return x
}
