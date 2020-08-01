package visitor

import (
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	"strings"
)

type IMySQLASTVisitor interface {
	visitor.ISQLVisitor
}

type MySQLASTVisitorAdapter struct {
	IMySQLASTVisitor
}

type MySQLASTOutputVisitor struct {
	*MySQLASTVisitorAdapter
	*visitor.SQLOutputVisitor
}

func NewOutputVisitor(builder strings.Builder, config config.Output) *MySQLASTOutputVisitor {
	x := new(MySQLASTOutputVisitor)
	x.Builder = builder
	x.Config = config
	return x
}
