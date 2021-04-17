package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/variable"
)

func AcceptVariableExpr(visitor ISQLVisitor, x *variable.SQLVariableExpr)  {
	visitor.VisitVariableExpr(visitor, x)
}
func AcceptAtVariableExpr(visitor ISQLVisitor, x *variable.SQLAtVariableExpr)  {
	visitor.VisitAtVariableExpr(visitor, x)
}
func AcceptAtAtVariableExpr(visitor ISQLVisitor, x *variable.SQLAtAtVariableExpr)  {
	visitor.VisitAtAtVariableExpr(visitor, x)
}

func AcceptBindVariableExpr(visitor ISQLVisitor, x *variable.SQLBindVariableExpr)  {
	if visitor.VisitBindVariableExpr(visitor, x) {
		Accept(visitor, x.Name())
	}

}