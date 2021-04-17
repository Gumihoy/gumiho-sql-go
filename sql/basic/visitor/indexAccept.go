package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/index"
)

func AcceptIndexColumn(visitor ISQLVisitor, x *index.SQLIndexColumn) {
	visitor.VisitIndexColumn(visitor, x)
}
