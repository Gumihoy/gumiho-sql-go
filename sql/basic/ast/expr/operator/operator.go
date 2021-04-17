package operator

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
)

// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/About-SQL-Operators.html#GUID-6A0C265F-3A7E-4E1C-8F79-8C6BCA26CFBA
type ISQLOperator interface {
	expr.ISQLExpr
	Paren() bool
	SetParen(paren bool)
}

type SQLUnaryOperator string

const (
	PRIOR           = "PRIOR"
	CONNECT_BY_ROOT = "CONNECT BY ROOT"
	COLLATE         = "COLLATE"
)

// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/About-SQL-Operators.html#GUID-6A0C265F-3A7E-4E1C-8F79-8C6BCA26CFBA
// ( )
type SQLUnaryOperatorExpr struct {
	*expr.AbstractSQLExpr
	paren    bool
	Operator SQLUnaryOperator
	operand  expr.ISQLExpr
}

func NewUnaryOperatorExpr(operator SQLUnaryOperator, operand  expr.ISQLExpr) *SQLUnaryOperatorExpr  {
	var x SQLUnaryOperatorExpr
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.Operator = operator
	x.operand = operand
	return &x
}
func (x *SQLUnaryOperatorExpr) Paren() bool {
	return x.paren
}
func (x *SQLUnaryOperatorExpr) SetParen(paren bool) {
	x.paren = paren
}
func (x *SQLUnaryOperatorExpr) Operand() expr.ISQLExpr {
	return x.operand
}
func (x *SQLUnaryOperatorExpr) SetOperand(operand expr.ISQLExpr)  {
	if ast.IsNil(operand) {
		return
	}
	operand.SetParent(x)
	x.operand = operand
}


type SQLBinaryOperator string

const (
	BIT_NOT       = "!"
	BIT_OR        = "|"
	BIT_AND       = "&"
	BIT_XOR       = "^"
	BIT_INVERSION = "~"

	MULTIPLY = "*"
	DIVIDE   = "/"
	MODULO   = "%"
	DIV      = "DIV"
	MOD      = "MOD"

	PLUS  = "+"
	MINUS = "-"

	CONCAT = "||"

	EQ = "="

	NOT_EQ1 = "!="
	NOT_EQ2 = "<>"
	NOT_EQ3 = "^="
	NOT_EQ4 = "~="

	LESS_THAN       = "<"
	LESS_THAN_EQ    = "<="
	GREATER_THAN    = ">"
	GREATER_THAN_EQ = "<="

	SHIFT_RIGHT = ">>"
	SHIFT_LEFT  = "<<"

	NOT = "NOT"
	AND = "AND"
	OR  = "OR"

	LOGICAL_AND = "&&"

)

// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/About-SQL-Operators.html#GUID-6A0C265F-3A7E-4E1C-8F79-8C6BCA26CFBA
type SQLBinaryOperatorExpr struct {
	*expr.AbstractSQLExpr

	paren    bool
	left     expr.ISQLExpr
	Operator SQLBinaryOperator
	right    expr.ISQLExpr
}

func NewBinaryOperator(left expr.ISQLExpr, operator SQLBinaryOperator, right expr.ISQLExpr) *SQLBinaryOperatorExpr {
	return NewBinaryOperatorWithParen(false, left, operator, right)
}

func NewBinaryOperatorWithParen(paren bool, left expr.ISQLExpr, operator SQLBinaryOperator, right expr.ISQLExpr) *SQLBinaryOperatorExpr {
	var x  SQLBinaryOperatorExpr
	x.AbstractSQLExpr = expr.NewAbstractExpr()

	x.paren = paren
	x.SetLeft(left)
	x.Operator = operator
	x.SetRight(right)
	return &x
}

func (x *SQLBinaryOperatorExpr) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.left {
		x.SetLeft(target)
		return true
	}
	if source == x.right {
		x.SetRight(target)
		return true
	}
	return false
}

func (x *SQLBinaryOperatorExpr) Paren() bool {
	return x.paren
}
func (x *SQLBinaryOperatorExpr) SetParen(paren bool) {
	x.paren = paren
}
func (x *SQLBinaryOperatorExpr) Left() expr.ISQLExpr {
	return x.left
}
func (x *SQLBinaryOperatorExpr) SetLeft(left expr.ISQLExpr) {
	if ast.IsNil(left) {
		return
	}
	left.SetParent(x)
	x.left = left
}

func (x *SQLBinaryOperatorExpr) Right() expr.ISQLExpr {
	return x.right
}
func (x *SQLBinaryOperatorExpr) SetRight(right expr.ISQLExpr) {
	if ast.IsNil(right) {
		return
	}
	right.SetParent(x)
	x.right = right
}


type PurposeType string

const (
	COMPARISON = "COMPARISON"
)

func (x *SQLBinaryOperatorExpr) Type() {
	switch x.Operator {
	case EQ, NOT_EQ1, NOT_EQ2, NOT_EQ3, NOT_EQ4:
		return
	}
}

// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/Pattern-matching-Conditions.html#GUID-D2124F3A-C6E4-4CCA-A40E-2FFCABFD8E19
// https://dev.mysql.com/doc/refman/8.0/en/string-comparison-functions.html#operator_like
type SQLRegexpLikeOperator struct {
	expr.AbstractSQLExpr
	arguments []expr.ISQLExpr
}

// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/Null-Conditions.html#GUID-657F2BA6-5687-4A00-8C2F-57515FD2DAEB
type SQLNullOperator struct {
	expr.AbstractSQLExpr
	expr expr.ISQLExpr
}

// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/Null-Conditions.html#GUID-657F2BA6-5687-4A00-8C2F-57515FD2DAEB
type SQLNotNullOperator struct {
	expr.AbstractSQLExpr
	expr expr.ISQLExpr
}

// https://dev.mysql.com/doc/refman/8.0/en/comparison-operators.html#operator_in
// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/IN-Condition.html#GUID-C7961CB3-8F60-47E0-96EB-BDCF5DB1317C
// expr IN (value,...)
type SQLInOperator struct {
	expr.AbstractSQLExpr
	expr   expr.ISQLExpr
	values []expr.ISQLExpr
}

func (x *SQLInOperator) Expr() expr.ISQLExpr {
	return x.expr
}

func (x *SQLInOperator) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

func (x *SQLInOperator) Values() []expr.ISQLExpr {
	return x.values
}

func (x *SQLInOperator) AddValue(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.values = append(x.values, expr)
}
