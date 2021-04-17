package sequence

import "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"

/**
 *  [ { INCREMENT BY | START WITH } integer
  | { MAXVALUE integer | NOMAXVALUE }
  | { MINVALUE integer | NOMINVALUE }
  | { CYCLE | NOCYCLE }
  | { CACHE integer | NOCACHE }
  | { ORDER | NOORDER }
  | { KEEP | NOKEEP }
  | { SCALE {EXTEND | NOEXTEND} | NOSCALE }
  | { SESSION | GLOBAL }
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-SEQUENCE.html#GUID-E9C78A8C-615A-4757-B2A8-5E6EFB130571
 */

/**
 * INCREMENT BY integer
 */
type SQLIncrementBySequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}

func NewIncrementBySequenceOption() *SQLIncrementBySequenceOption {
	x := new(SQLIncrementBySequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIncrementBySequenceOption) Value() expr.ISQLExpr {
	return x.value
}
func (x *SQLIncrementBySequenceOption) SetValue(value expr.ISQLExpr)  {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.value = value
}
/**
 * START WITH integer
 */
type SQLStartWithSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewStartWithSequenceOption() *SQLStartWithSequenceOption {
	x := new(SQLStartWithSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLStartWithSequenceOption) Value() expr.ISQLExpr {
	return x.value
}
func (x *SQLStartWithSequenceOption) SetValue(value expr.ISQLExpr)  {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.value = value
}

/**
 * MAXVALUE integer
 */
type SQLMaxValueSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewMaxValueSequenceOption() *SQLMaxValueSequenceOption {
	x := new(SQLMaxValueSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLMaxValueSequenceOption) Value() expr.ISQLExpr {
	return x.value
}
func (x *SQLMaxValueSequenceOption) SetValue(value expr.ISQLExpr)  {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.value = value
}
/**
 * NOMAXVALUE
 */
type SQLNoMaxValueSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewNoMaxValueSequenceOption() *SQLNoMaxValueSequenceOption {
	x := new(SQLNoMaxValueSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * MINVALUE integer
 */
type SQLMinValueSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewMinValueSequenceOption() *SQLMinValueSequenceOption {
	x := new(SQLMinValueSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLMinValueSequenceOption) Value() expr.ISQLExpr {
	return x.value
}
func (x *SQLMinValueSequenceOption) SetValue(value expr.ISQLExpr)  {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.value = value
}
/**
 * NOMINVALUE
 */
type SQLNoMinValueSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewNoMinValueSequenceOption() *SQLNoMinValueSequenceOption {
	x := new(SQLNoMinValueSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * CYCLE
 */
type SQLCycleSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewCycleSequenceOption() *SQLCycleSequenceOption {
	x := new(SQLCycleSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * NOCYCLE
 */
type SQLNoCycleSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewNoCycleSequenceOption() *SQLNoCycleSequenceOption {
	x := new(SQLNoCycleSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * CACHE integer
 */
type SQLCacheSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewCacheSequenceOption() *SQLCacheSequenceOption {
	x := new(SQLCacheSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLCacheSequenceOption) Value() expr.ISQLExpr {
	return x.value
}
func (x *SQLCacheSequenceOption) SetValue(value expr.ISQLExpr)  {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.value = value
}

/**
 * NOCACHE
 */
type SQLNoCacheSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewNoCacheSequenceOption() *SQLNoCacheSequenceOption {
	x := new(SQLNoCacheSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * ORDER
 */
type SQLOrderSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewOrderSequenceOption() *SQLOrderSequenceOption {
	x := new(SQLOrderSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * NOORDER
 */
type SQLNoOrderSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewNoOrderSequenceOption() *SQLNoOrderSequenceOption {
	x := new(SQLNoOrderSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * KEEP
 */
type SQLKeepSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewKeepSequenceOption() *SQLKeepSequenceOption {
	x := new(SQLKeepSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * NOKEEP
 */
type SQLNoKeepSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewNoKeepSequenceOption() *SQLNoKeepSequenceOption {
	x := new(SQLNoKeepSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * SCALE {EXTEND | NOEXTEND}
 */
type SQLScaleSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewScaleSequenceOption() *SQLScaleSequenceOption {
	x := new(SQLScaleSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLScaleSequenceOption) Value() expr.ISQLExpr {
	return x.value
}
func (x *SQLScaleSequenceOption) SetValue(value expr.ISQLExpr)  {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.value = value
}

/**
 * NOSCALE
 */
type SQLNoScaleSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewNoScaleSequenceOption() *SQLNoScaleSequenceOption {
	x := new(SQLNoScaleSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * SESSION
 */
type SQLSessionSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewSessionSequenceOption() *SQLSessionSequenceOption {
	x := new(SQLSessionSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * GLOBAL
 */
type SQLGlobalSequenceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}
func NewGlobalSequenceOption() *SQLGlobalSequenceOption {
	x := new(SQLGlobalSequenceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}