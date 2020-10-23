package visitor

import (
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	mysqlVisitor "gumihoy.com/sql/dialect/mysql/visitor"
	"strings"
)

type IMariaDBVisitor interface {
	visitor.ISQLVisitor
}

type MariaDBVisitorAdapter struct {
	*mysqlVisitor.MySQLVisitorAdapter
}

func NewVisitorAdapter() *MariaDBVisitorAdapter {
	return &MariaDBVisitorAdapter{mysqlVisitor.NewVisitorAdapter()}
}

type MariaDBOutputVisitor struct {
	*mysqlVisitor.MySQLOutputVisitor
}

func NewOutputVisitor(builder *strings.Builder, config config.Output) *MariaDBOutputVisitor {
	x := new(MariaDBOutputVisitor)
	x.MySQLOutputVisitor = mysqlVisitor.NewOutputVisitor(builder, config)
	return x
}
