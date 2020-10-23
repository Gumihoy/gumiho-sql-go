package variable

import (
	"gumihoy.com/sql/basic/ast/expr"
)

/**
 * ?
 */
type SQLVariableExpr struct {
	*expr.SQLExpr
}

func NewVariableExpr() *SQLVariableExpr {
	var x SQLVariableExpr
	x.SQLExpr = expr.NewExpr()
	return &x
}


func (x *SQLVariableExpr) Value() string {
	return "?"
}

/**
 * :name
 */
type SQLBindVariableExpr struct {
	*expr.SQLExpr
	name expr.ISQLIdentifier
}

func NewBindVariableExpr(name string) *SQLBindVariableExpr {
	return NewBindVariableExprWithName(expr.OfIdentifier(name))
}


func NewBindVariableExprWithName(name expr.ISQLIdentifier) *SQLBindVariableExpr {
	var x SQLBindVariableExpr
	x.SetName(name)
	return &x
}

func (x *SQLBindVariableExpr) Name() expr.ISQLIdentifier {
	return x.name
}
func (x *SQLBindVariableExpr) SetName(name expr.ISQLIdentifier) {
	name.SetParent(x)
	x.name = name
}
