package visitor

import (
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	"strings"
)

type IMySQLVisitor interface {
	visitor.ISQLVisitor
}

type MySQLVisitorAdapter struct {
	*visitor.SQLVisitorAdapter
}

func NewVisitorAdapter() *MySQLVisitorAdapter {
	return NewVisitorAdapterWithVisitorAdapter(visitor.NewVisitorAdapter())
}

func NewVisitorAdapterWithVisitorAdapter(adapter *visitor.SQLVisitorAdapter) *MySQLVisitorAdapter {
	return &MySQLVisitorAdapter{adapter}
}

type MySQLOutputVisitor struct {
	*visitor.SQLOutputVisitor
}

func NewOutputVisitor(builder *strings.Builder, config config.Output) *MySQLOutputVisitor {
	x := new(MySQLOutputVisitor)
	x.SQLOutputVisitor = visitor.NewOutputVisitor(builder, config)
	return x
}




