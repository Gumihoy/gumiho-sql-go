package server

import "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"

/**
 *  HOST character-literal
  | DATABASE character-literal
  | USER character-literal
  | PASSWORD character-literal
  | SOCKET character-literal
  | OWNER character-literal
  | PORT numeric-literal
 * https://dev.mysql.com/doc/refman/8.0/en/create-server.html
 */
type ISQLViewOption interface {
	expr.ISQLExpr
}

/**
 *  HOST character-literal
 */
type SQLHostOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}

func NewHostOption() *SQLHostOption {
	x := new(SQLHostOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLHostOption) Value() expr.ISQLExpr  {
	return x.value
}
func (x *SQLHostOption) SetValue(value expr.ISQLExpr)   {
	value.SetParent(x)
	x.value = value
}

/**
 * DATABASE character-literal
 */
type SQLDatabaseOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewDatabaseOption() *SQLDatabaseOption {
	x := new(SQLDatabaseOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDatabaseOption) Value() expr.ISQLExpr  {
	return x.value
}
func (x *SQLDatabaseOption) SetValue(value expr.ISQLExpr)   {
	value.SetParent(x)
	x.value = value
}
/**
 * USER character-literal
 */
type SQLUserOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewUserOption() *SQLUserOption {
	x := new(SQLUserOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLUserOption) Value() expr.ISQLExpr  {
	return x.value
}
func (x *SQLUserOption) SetValue(value expr.ISQLExpr)   {
	value.SetParent(x)
	x.value = value
}
/**
 * PASSWORD character-literal
 */
type SQLPasswordOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewPasswordOption() *SQLPasswordOption {
	x := new(SQLPasswordOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLPasswordOption) Value() expr.ISQLExpr  {
	return x.value
}
func (x *SQLPasswordOption) SetValue(value expr.ISQLExpr)   {
	value.SetParent(x)
	x.value = value
}
/**
 * SOCKET character-literal
 */
type SQLSocketOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewSocketOption() *SQLSocketOption {
	x := new(SQLSocketOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLSocketOption) Value() expr.ISQLExpr  {
	return x.value
}
func (x *SQLSocketOption) SetValue(value expr.ISQLExpr)   {
	value.SetParent(x)
	x.value = value
}
/**
 * OWNER character-literal
 */
type SQLOwnerOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewOwnerOption() *SQLOwnerOption {
	x := new(SQLOwnerOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLOwnerOption) Value() expr.ISQLExpr  {
	return x.value
}
func (x *SQLOwnerOption) SetValue(value expr.ISQLExpr)   {
	value.SetParent(x)
	x.value = value
}
/**
 * PORT numeric-literal
 */
type SQLPortOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewPortOption() *SQLPortOption {
	x := new(SQLPortOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLPortOption) Value() expr.ISQLExpr  {
	return x.value
}
func (x *SQLPortOption) SetValue(value expr.ISQLExpr)   {
	value.SetParent(x)
	x.value = value
}