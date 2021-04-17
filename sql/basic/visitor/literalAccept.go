package visitor

import "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/literal"

func AcceptStringLiteral(visitor ISQLVisitor, x *literal.SQLStringLiteral) {
	visitor.VisitStringLiteral(visitor, x)
}

func AcceptCharacterStringLiteral(visitor ISQLVisitor, x *literal.SQLCharacterStringLiteral) {
	visitor.VisitCharacterStringLiteral(visitor, x)
}


func AcceptIntegerLiteral(visitor ISQLVisitor, x *literal.SQLIntegerLiteral) {
	visitor.VisitIntegerLiteral(visitor, x)
}

func AcceptFloatingPointLiteral(visitor ISQLVisitor, x *literal.SQLFloatingPointLiteral) {
	visitor.VisitFloatingPointLiteral(visitor, x)
}
func AcceptHexadecimalLiteral(visitor ISQLVisitor, x *literal.SQLHexadecimalLiteral) {
	visitor.VisitHexadecimalLiteral(visitor, x)
}



func AcceptDateLiteral(visitor ISQLVisitor, x *literal.SQLDateLiteral) {
	if visitor.VisitDateLiteral(visitor, x) {
		Accept(visitor, x.Value())
	}
}

func AcceptTimeLiteral(visitor ISQLVisitor, x *literal.SQLTimeLiteral) {
	if visitor.VisitTimeLiteral(visitor, x) {
		Accept(visitor, x.Value())
	}
}

func AcceptTimestampLiteral(visitor ISQLVisitor, x *literal.SQLTimestampLiteral) {
	if visitor.VisitTimestampLiteral(visitor, x) {
		Accept(visitor, x.Value())
	}
}

func AcceptBooleanLiteral(visitor ISQLVisitor, x *literal.SQLBooleanLiteral) {
	visitor.VisitBooleanLiteral(visitor, x)
}

