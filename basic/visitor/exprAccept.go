package visitor

import (
	"gumihoy.com/sql/basic/ast/expr"
)


func AcceptIdentifier(visitor ISQLVisitor, x *expr.SQLUnQuotedIdentifier) {
	visitor.VisitIdentifier(x)
}


func AcceptDoubleQuotedIdentifier(visitor ISQLVisitor, x *expr.SQLDoubleQuotedIdentifier) {
	visitor.VisitDoubleQuotedIdentifier(x)
}

func AcceptReverseQuotedIdentifier(visitor ISQLVisitor, x *expr.SQLReverseQuotedIdentifier) {
	visitor.VisitReverseQuotedIdentifier(x)
}

func AcceptName(visitor ISQLVisitor, x *expr.SQLName) {
	visitor.VisitName(x)
}

func AcceptAllColumnExpr(visitor ISQLVisitor, x *expr.SQLAllColumnExpr) {
	visitor.VisitAllColumnExpr(x)
}