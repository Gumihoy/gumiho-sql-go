package visitor

import "gumihoy.com/sql/basic/ast/expr/literal"

func AcceptStringLiteral(visitor ISQLVisitor, x *literal.SQLStringLiteral) {
	visitor.VisitStringLiteral(x)
}

func AcceptCharacterStringLiteral(visitor ISQLVisitor, x *literal.SQLCharacterStringLiteral) {
	visitor.VisitCharacterStringLiteral(x)
}


func AcceptIntegerLiteral(visitor ISQLVisitor, x *literal.SQLIntegerLiteral) {
	visitor.VisitIntegerLiteral(x)
}

func AcceptFloatingPointLiteral(visitor ISQLVisitor, x *literal.SQLFloatingPointLiteral) {
	visitor.VisitFloatingPointLiteral(x)
}




func AcceptDateLiteral(visitor ISQLVisitor, x *literal.SQLDateLiteral) {
	if visitor.VisitDateLiteral(x) {
		Accept(visitor, x.Value())
	}
}

func AcceptTimeLiteral(visitor ISQLVisitor, x *literal.SQLTimeLiteral) {
	if visitor.VisitTimeLiteral(x) {
		Accept(visitor, x.Value())
	}
}

func AcceptTimestampLiteral(visitor ISQLVisitor, x *literal.SQLTimestampLiteral) {
	if visitor.VisitTimestampLiteral(x) {
		Accept(visitor, x.Value())
	}
}

func AcceptBooleanLiteral(visitor ISQLVisitor, x *literal.SQLBooleanLiteral) {
	visitor.VisitBooleanLiteral(x)
}

