package literal

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"strconv"
	"strings"
)

type ISQLLiteral interface {
	expr.ISQLExpr
}

/**
 * 'xxx'
 */
type SQLStringLiteral struct {
	*expr.AbstractSQLExpr
	value string
}

func NewStringLiteral(value string) *SQLStringLiteral {
	var x SQLStringLiteral
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.value = value
	return &x
}
func (x *SQLStringLiteral) StringName() string {
	return x.value
}

func (x *SQLStringLiteral) Equal(t string) bool {
	return x.value == t
}

func (x *SQLStringLiteral) EqualIgnoreCase(t string) bool {
	return strings.EqualFold(x.value, t)
}
func (x *SQLStringLiteral) Value() string {
	return x.value
}

/**
 * _latin1'string'、_binary'string'、_utf8'string'、N'some text'、n'some text'
 * https://dev.mysql.com/doc/refman/8.0/en/string-literals.html
 *
 * q'some text'、Q'some text'、
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Literals.html#GUID-1824CBAA-6E16-4921-B2A6-112FB02248DA
 *
 */
type SQLCharacterStringLiteral struct {
	*expr.AbstractSQLExpr
	charset string
	value   string
}

func NewCharacterStringLiteral(charset string, value string) *SQLCharacterStringLiteral {
	var x SQLCharacterStringLiteral
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
	*expr.AbstractSQLExpr
	value string
}

func newAbstractNumericLiteral(value string) *AbstractSQLNumericLiteral {
	var x AbstractSQLNumericLiteral
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}

func newAbstractDateTimeLiteral(value expr.ISQLExpr) *AbstractSQLDateTimeLiteral {
	var x AbstractSQLDateTimeLiteral
	x.AbstractSQLExpr = expr.NewAbstractExpr()
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
	*expr.AbstractSQLExpr
	hexadecimal SQLHexadecimal
	value       string
}

func NewHexadecimalLiteral(hexadecimal SQLHexadecimal, value string) *SQLHexadecimalLiteral {
	x := new(SQLHexadecimalLiteral)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetHexadecimal(hexadecimal)
	x.SetValue(value)
	return x
}

func (self *SQLHexadecimalLiteral) Hexadecimal() SQLHexadecimal {
	return self.hexadecimal
}

func (self *SQLHexadecimalLiteral) SetHexadecimal(hexadecimal SQLHexadecimal) {
	self.hexadecimal = hexadecimal
}

func (self *SQLHexadecimalLiteral) Value() string {
	return self.value
}

func (self *SQLHexadecimalLiteral) SetValue(value string) {
	self.value = value
}

type SQLHexadecimal string

const (
	X_UPPER SQLHexadecimal = "X"
	X_LOWER                = "x"
	ZERO_X  SQLHexadecimal = "0x"
)

// https://dev.mysql.com/doc/refman/8.0/en/bit-value-literals.html
// b'01'
// B'01'
// 0b01
type SQLBitValueLiteral struct {
}

// https://dev.mysql.com/doc/refman/8.0/en/boolean-literals.html
// true / false
type SQLBooleanLiteral struct {
	*expr.AbstractSQLExpr
	value bool
}

func NewBooleanLiteral(value bool) *SQLBooleanLiteral {
	x := new(SQLBooleanLiteral)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.value = value
	return x
}

func (x *SQLBooleanLiteral) Value() bool {
	return x.value
}
func (x *SQLBooleanLiteral) StringValue() string {
	return strconv.FormatBool(x.value)
}

/**
 * INTERVAL [ <sign> ] <interval string> <interval qualifier>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#interval%20literal
 *
 * INTERVAL expr unit
 * https://dev.mysql.com/doc/refman/8.0/en/expressions.html
 *
 *
 *
 *
 *
 * INTERVAL YEAR TO MONTH
 * INTERVAL 'integer [- integer ]' { YEAR | MONTH } [ (precision) ] [ TO { YEAR | MONTH } ]
 *
 * INTERVAL DAY TO SECOND
 * INTERVAL '{ integer | integer time_expr | time_expr }' { { DAY | HOUR | MINUTE } [ (leading_precision) ]
| SECOND [ (leading_precision [, fractional_seconds_precision ]) ]
}
[ TO { DAY | HOUR | MINUTE | SECOND [ (fractional_seconds_precision) ] } ]
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Literals.html#GUID-DC8D1DAD-7D04-45EA-9546-82810CD09A1B
 */
type SQLIntervalUnit string

const (
	YEAR   SQLIntervalUnit = "YEAR"
	MONTH                  = "MONTH"
	DAY                    = "DAY"
	HOUR                   = "HOUR"
	MINUTE                 = "MINUTE"
	SECOND                 = "SECOND"
)

type SQLIntervalLiteral struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
	start *SQLIntervalLiteralField
	end   *SQLIntervalLiteralField
}

func NewIntervalLiteral() *SQLIntervalLiteral {
	x := new(SQLIntervalLiteral)

	return x
}
func (x *SQLIntervalLiteral) Value() *SQLIntervalLiteralField {
	return x.start
}

func (x *SQLIntervalLiteral) SetValue(value expr.ISQLExpr) {
	value.SetParent(x)
	x.value = value
}
func (x *SQLIntervalLiteral) Start() *SQLIntervalLiteralField {
	return x.start
}

func (x *SQLIntervalLiteral) SetStart(start *SQLIntervalLiteralField) {
	x.start = start
}

func (x *SQLIntervalLiteral) End() *SQLIntervalLiteralField {
	return x.end
}

func (x *SQLIntervalLiteral) SetEnd(end *SQLIntervalLiteralField) {
	x.end = end
}

/**
 *  YEAR | MONTH | DAY | HOUR | MINUTE [ <left paren> <interval leading field precision> <right paren> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#start%20field
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#interval%20qualifier
 */
type SQLIntervalLiteralField struct {
	*expr.AbstractSQLExpr
	unit       SQLIntervalUnit
	precisions []expr.ISQLExpr
}

func NewIntervalLiteralFieldWitUnit(unit SQLIntervalUnit) *SQLIntervalLiteralField {
	x := new(SQLIntervalLiteralField)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.unit = unit
	return x
}

func NewIntervalLiteralFieldWitUnitAndPrecisions(unit SQLIntervalUnit, precisions []expr.ISQLExpr) *SQLIntervalLiteralField {
	x := new(SQLIntervalLiteralField)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.unit = unit
	if precisions != nil {
		for _, precision := range precisions {
			x.AddPrecision(precision)
		}
	}
	return x
}

func (x *SQLIntervalLiteralField) Unit() SQLIntervalUnit {
	return x.unit
}

func (x *SQLIntervalLiteralField) SetEnd(unit SQLIntervalUnit) {
	x.unit = unit
}

func (x *SQLIntervalLiteralField) Precisions() []expr.ISQLExpr {
	return x.precisions
}

func (x *SQLIntervalLiteralField) AddPrecision(precision expr.ISQLExpr) {
	if precision == nil {
		return
	}
	precision.SetParent(x)
	x.precisions = append(x.precisions, precision)
}
