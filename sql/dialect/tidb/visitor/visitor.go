package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	"github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/visitor"
	"strings"
)

type ITiDBVisitor interface {
	visitor.IMySQLVisitor
}

type TiDBVisitorAdapter struct {
	*visitor.MySQLVisitorAdapter
}

func NewVisitorAdapter() *TiDBVisitorAdapter {
	return NewVisitorAdapterWithVisitorAdapter(visitor.NewVisitorAdapter())
}

func NewVisitorAdapterWithVisitorAdapter(adapter *visitor.MySQLVisitorAdapter) *TiDBVisitorAdapter {
	return &TiDBVisitorAdapter{adapter}
}

type TiDBOutputVisitor struct {
	*visitor.MySQLOutputVisitor
}

func NewOutputVisitor(builder *strings.Builder, config *config.SQLOutputConfig) *TiDBOutputVisitor {
	x := new(TiDBOutputVisitor)
	x.MySQLOutputVisitor = visitor.NewOutputVisitor(builder, config)
	return x
}




