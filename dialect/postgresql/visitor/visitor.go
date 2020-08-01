package visitor

import (
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	"strings"
)

type IPostgreSQLVisitor interface {
	visitor.ISQLVisitor
}

type PostgreSQLVisitorAdapter struct {
	IPostgreSQLVisitor
}

type PostgreSQLOutputVisitor struct {
	*PostgreSQLVisitorAdapter
	*visitor.SQLOutputVisitor
}

func NewOutputVisitor(builder strings.Builder, config config.Output) *PostgreSQLOutputVisitor {
	x := new(PostgreSQLOutputVisitor)
	x.Builder = builder
	x.Config = config
	return x
}
