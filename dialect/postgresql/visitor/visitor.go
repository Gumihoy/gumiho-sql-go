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
	*visitor.SQLVisitorAdapter
}

func NewSQLVisitorAdapter() *PostgreSQLVisitorAdapter {
	return NewSQLVisitorAdapterWithVisitorAdapter(visitor.NewVisitorAdapter())
}

func NewSQLVisitorAdapterWithVisitorAdapter(adapter *visitor.SQLVisitorAdapter) *PostgreSQLVisitorAdapter {
	x := new(PostgreSQLVisitorAdapter)
	x.SQLVisitorAdapter = adapter
	return x
}

type PostgreSQLOutputVisitor struct {
	*visitor.SQLOutputVisitor
}

func NewOutputVisitor(builder *strings.Builder, config config.Output) *PostgreSQLOutputVisitor {
	x := new(PostgreSQLOutputVisitor)
	x.SQLOutputVisitor = visitor.NewOutputVisitor(builder, config)
	return x
}
