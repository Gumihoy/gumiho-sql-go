package literal

import (
	"gumihoy.com/sql/basic/ast/expr"
)

type ISQLLiteral interface {
	expr.ISQLExpr
	Value() interface{}
}

type SQLDateTimeLiteral struct {
	ISQLLiteral
	value bool
}

func (x *SQLDateTimeLiteral) Value() interface{} {
	return x.value
}

type SQLNumericLiteral struct {
	ISQLLiteral
	value string
}

func (x *SQLNumericLiteral) Value() string {
	return x.value
}

type SQLStringLiteral struct {
	ISQLLiteral
	value string
}

func (x *SQLStringLiteral) Value() string {
	return x.value
}

type SQLBooleanLiteral struct {
	ISQLLiteral
	value bool
}

func NewBooleanLiteral(value bool) *SQLBooleanLiteral {
	x := new(SQLBooleanLiteral)
	x.value = value
	return x
}

func (x *SQLBooleanLiteral) Value() bool {
	return x.value
}
