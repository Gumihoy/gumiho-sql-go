package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
)

func AcceptMinusComment(visitor ISQLVisitor, x *ast.SQLMinusComment)  {
	if visitor.VisitMinusComment(visitor, x) {
	}
}
func AcceptMultiLineComment(visitor ISQLVisitor, x *ast.SQLMultiLineComment)  {
	if visitor.VisitMultiLineComment(visitor, x) {
	}
}
func AcceptSharpComment(visitor ISQLVisitor, x *ast.SQLSharpComment)  {
	if visitor.VisitSharpComment(visitor, x) {
	}
}
