package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	mysqlVisitor "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/visitor"
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

func NewOutputVisitor(builder *strings.Builder, config *config.SQLOutputConfig) *MariaDBOutputVisitor {
	x := new(MariaDBOutputVisitor)
	x.MySQLOutputVisitor = mysqlVisitor.NewOutputVisitor(builder, config)
	return x
}
