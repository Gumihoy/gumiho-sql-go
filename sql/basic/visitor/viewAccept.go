package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
)

func AcceptViewColumn(visitor ISQLVisitor, x *view.SQLViewColumn) {
	if visitor.VisitViewColumn(visitor, x) {
		Accept(visitor, x.Name())
	}
}