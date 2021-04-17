package variable

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
)

/**
 * ?
 */
type SQLVariableExpr struct {
	*expr.AbstractSQLExpr
}

func NewVariableExpr() *SQLVariableExpr {
	var x SQLVariableExpr
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return &x
}

func (x *SQLVariableExpr) Value() string {
	return "?"
}

/**
 * @name
 * https://dev.mysql.com/doc/refman/8.0/en/user-variables.html
 */
type SQLAtVariableExpr struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewAtVariableExpr() *SQLAtVariableExpr {
	x := new(SQLAtVariableExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func NewAtVariableExprWithName(name expr.ISQLName) *SQLAtVariableExpr {
	x := new(SQLAtVariableExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetName(name)
	return x
}

func (x *SQLAtVariableExpr) Name() expr.ISQLName {
	return x.name
}
func (x *SQLAtVariableExpr) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * @@name
 * @@kind.name
 * kind name
 * https://dev.mysql.com/doc/refman/8.0/en/user-variables.html
 */
type SQLAtAtVariableKind string

const (
	GLOBAL       SQLAtAtVariableKind = "GLOBAL"
	LOCAL                            = "LOCAL"
	PERSIST                          = "PERSIST"
	PERSIST_ONLY                     = "PERSIST_ONLY"
	SESSION                          = "SESSION"
)
/**
 * @@name
 * @@kind.name
 * kind name
 * https://dev.mysql.com/doc/refman/8.0/en/user-variables.html
 */
type SQLAtAtVariableExpr struct {
	*expr.AbstractSQLExpr
	Kind    SQLAtAtVariableKind
	HasAtAT bool
	name    expr.ISQLName
}

func NewAtAtVariableExpr() *SQLAtAtVariableExpr {
	x := new(SQLAtAtVariableExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func NewAtAtVariableExprWithName(name expr.ISQLName) *SQLAtAtVariableExpr {
	x := new(SQLAtAtVariableExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetName(name)
	return x
}

func (x *SQLAtAtVariableExpr) Name() expr.ISQLName {
	return x.name
}
func (x *SQLAtAtVariableExpr) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * :name
 * :owner.name
 * :1, :2
 */
type SQLBindVariableExpr struct {
	*expr.AbstractSQLExpr
	name expr.ISQLExpr
}

func NewBindVariableExpr() *SQLBindVariableExpr {
	x := new(SQLBindVariableExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func NewBindVariableExprWithName(name expr.ISQLExpr) *SQLBindVariableExpr {
	x := new(SQLBindVariableExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetName(name)
	return x
}

func (x *SQLBindVariableExpr) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLBindVariableExpr) SetName(name expr.ISQLExpr) {
	name.SetParent(x)
	x.name = name
}
