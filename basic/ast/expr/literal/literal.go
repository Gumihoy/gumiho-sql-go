package literal

import (
	"gumihoy.com/sql/basic/ast/expr"
	"strconv"
)

type ISQLLiteral interface {
	expr.ISQLExpr
}


/**
 * 'xxx'
 */
type SQLStringLiteral struct {
	*expr.SQLExpr
	value string
}

func NewStringLiteral(value string) *SQLStringLiteral {
	var x SQLStringLiteral
	x.SQLExpr = expr.NewExpr()
	x.value = value
	return &x
}

func (x *SQLStringLiteral) Value() string {
	return x.value
}

type SQLCharacterStringLiteral struct {
	*expr.SQLExpr
	charset string
	value   string
}

func NewCharacterStringLiteral(charset string, value string) *SQLCharacterStringLiteral {
	var x SQLCharacterStringLiteral
	x.SQLExpr = expr.NewExpr()
	x.charset = charset
	x.value = value
	return &x
}
func (x *SQLCharacterStringLiteral) Charset() string {
	return x.charset
}
func (x *SQLCharacterStringLiteral) SetCharset(charset string) {
	x.charset = charset
}
func (x *SQLCharacterStringLiteral) Value() string {
	return x.value
}
func (x *SQLCharacterStringLiteral) SetValue(value string) {
	x.value = value
}





type ISQLNumericLiteral struct {
	ISQLLiteral
}
type AbstractSQLNumericLiteral struct {
	*expr.SQLExpr
	value string
}

func newAbstractNumericLiteral(value string) *AbstractSQLNumericLiteral {
	var x AbstractSQLNumericLiteral
	x.SQLExpr = expr.NewExpr()
	x.value = value
	return &x
}
func (x *AbstractSQLNumericLiteral) StringValue() string {
	return x.value
}

type SQLIntegerLiteral struct {
	*AbstractSQLNumericLiteral
}

func NewIntegerLiteral(value int64) *SQLIntegerLiteral {
	var x SQLIntegerLiteral
	x.AbstractSQLNumericLiteral = newAbstractNumericLiteral(string(value))
	return &x
}

func NewIntegerLiteralWithString(value string) *SQLIntegerLiteral {
	var x SQLIntegerLiteral
	x.AbstractSQLNumericLiteral = newAbstractNumericLiteral(value)
	return &x
}

func (x *SQLIntegerLiteral) Value() (int64, error) {
	return strconv.ParseInt(x.value, 10, 64)
}

type SQLFloatingPointLiteral struct {
	*AbstractSQLNumericLiteral
}

func NewFloatingPointLiteral(value float64) *SQLFloatingPointLiteral {
	var x SQLFloatingPointLiteral
	x.AbstractSQLNumericLiteral = newAbstractNumericLiteral(strconv.FormatFloat(value, 'E', -1, 64))
	return &x
}
func NewFloatingPointLiteralWith(value string) *SQLFloatingPointLiteral {
	var x SQLFloatingPointLiteral
	x.AbstractSQLNumericLiteral = newAbstractNumericLiteral(value)
	return &x
}

func (x *SQLFloatingPointLiteral) Value() (float64, error) {
	return strconv.ParseFloat(x.value, 64)
}




type AbstractSQLDateTimeLiteral struct {
	*expr.SQLExpr
	value expr.ISQLExpr
}

func newAbstractDateTimeLiteral(value expr.ISQLExpr) *AbstractSQLDateTimeLiteral {
	var x AbstractSQLDateTimeLiteral
	x.SQLExpr = expr.NewExpr()
	x.value = value
	return &x
}
func (x *AbstractSQLDateTimeLiteral) Value() expr.ISQLExpr {
	return x.value
}

func (x *AbstractSQLDateTimeLiteral) SetValue(value expr.ISQLExpr) {
	value.SetParent(x)
	x.value = value
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#datetime%20literal
// https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Literals.html#GUID-8F4B3F82-8821-4071-84D6-FBBA21C05AC1
//
// DATE '1998-12-25'
type SQLDateLiteral struct {
	*AbstractSQLDateTimeLiteral
}

func NewDateLiteral(value expr.ISQLExpr) *SQLDateLiteral {
	var x SQLDateLiteral
	x.AbstractSQLDateTimeLiteral = newAbstractDateTimeLiteral(value)
	return &x
}

// TIME '1998-12-25'
type SQLTimeLiteral struct {
	*AbstractSQLDateTimeLiteral
}

func NewTimeLiteral(value expr.ISQLExpr) *SQLTimeLiteral {
	var x SQLTimeLiteral
	x.AbstractSQLDateTimeLiteral = newAbstractDateTimeLiteral(value)
	return &x
}

// TIMESTAMP ''
type SQLTimestampLiteral struct {
	*AbstractSQLDateTimeLiteral
}

func NewTimestampLiteral(value expr.ISQLExpr) *SQLTimestampLiteral {
	var x SQLTimestampLiteral
	x.AbstractSQLDateTimeLiteral = newAbstractDateTimeLiteral(value)
	return &x
}



// https://dev.mysql.com/doc/refman/8.0/en/hexadecimal-literals.html
// X'01AF'
// X'01af'
// x'01AF'
// x'01af'
// 0x01AF
// 0x01af
type SQLHexadecimalLiteral struct {

}


// https://dev.mysql.com/doc/refman/8.0/en/bit-value-literals.html
// b'01'
// B'01'
// 0b01
type SQLBitValueLiteral struct {

}




// https://dev.mysql.com/doc/refman/8.0/en/boolean-literals.html
// true / false
type SQLBooleanLiteral struct {
	*expr.SQLExpr
	value bool
}

func NewBooleanLiteral(value bool) *SQLBooleanLiteral {
	x := new(SQLBooleanLiteral)
	x.SQLExpr = expr.NewExpr()
	x.value = value
	return x
}

func (x *SQLBooleanLiteral) Value() bool {
	return x.value
}
func (x *SQLBooleanLiteral) StringValue() string {
	return strconv.FormatBool(x.value)
}
