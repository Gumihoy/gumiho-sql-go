package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/common"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
)

func AcceptIdentifier(visitor ISQLVisitor, x *expr.SQLUnQuotedIdentifier) {
	visitor.VisitIdentifier(visitor, x)
}

func AcceptDoubleQuotedIdentifier(visitor ISQLVisitor, x *expr.SQLDoubleQuotedIdentifier) {
	visitor.VisitDoubleQuotedIdentifier(visitor, x)
}

func AcceptReverseQuotedIdentifier(visitor ISQLVisitor, x *expr.SQLReverseQuotedIdentifier) {
	visitor.VisitReverseQuotedIdentifier(visitor, x)
}

func AcceptName(visitor ISQLVisitor, x *expr.SQLName) {
	if visitor.VisitName(visitor, x) {
		Accept(visitor, x.Owner())
		Accept(visitor, x.Name())
	}
}
func AcceptDBLinkExpr(visitor ISQLVisitor, x *expr.SQLDBLinkExpr) {
	if visitor.VisitDBLinkExpr(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.DBLink())
	}
}



func AcceptAllColumnExpr(visitor ISQLVisitor, x *expr.SQLAllColumnExpr) {
	visitor.VisitAllColumnExpr(visitor, x)
}
func AcceptNullExpr(visitor ISQLVisitor, x *expr.SQLNullExpr) {
	visitor.VisitNullExpr(visitor, x)
}
func AcceptListExpr(visitor ISQLVisitor, x *expr.SQLListExpr) {
	if visitor.VisitListExpr(visitor, x) {
		for _, child := range x.Elements() {
			Accept(visitor, child)
		}
	}
}
func AcceptSubQuery(visitor ISQLVisitor, x *common.SQLSubQueryExpr) {
	if visitor.VisitSubQuery(visitor, x) {
		Accept(visitor, x.Query())
	}
}

func AcceptAssignExpr(visitor ISQLVisitor, x *expr.SQLAssignExpr) {
	if visitor.VisitAssignExpr(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Value())
	}
}



func AcceptEditionAbleExpr(visitor ISQLVisitor, x *statement.SQLEditionAbleExpr) {
	if visitor.VisitEditionAbleExpr(visitor, x) {
	}
}
func AcceptNonEditionAbleExpr(visitor ISQLVisitor, x *statement.SQLNonEditionAbleExpr) {
	if visitor.VisitNonEditionAbleExpr(visitor, x) {
	}
}
func AcceptCompileExpr(visitor ISQLVisitor, x *statement.SQLCompileExpr) {
	if visitor.VisitCompileExpr(visitor, x) {
	}
}



