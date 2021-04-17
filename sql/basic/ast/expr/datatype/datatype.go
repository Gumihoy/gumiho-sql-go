package datatype

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#data%20type
type ISQLDataType interface {
	expr.ISQLExpr
}

type abstractSQLDataType struct {
	*expr.AbstractSQLExpr
}

func NewAbstractSQLDataType(dbType db.Type) *abstractSQLDataType {
	x := new(abstractSQLDataType)
	x.AbstractSQLExpr = expr.NewAbstractExprWithDBType(dbType)
	return x
}

type SQLDataType struct {
	*abstractSQLDataType

	name expr.ISQLName

	Paren      bool
	precisions []expr.ISQLExpr
}

func NewDataTypeWithName(name string, dbType db.Type) *SQLDataType {
	x := new(SQLDataType)
	x.abstractSQLDataType = NewAbstractSQLDataType(dbType)
	x.SetName(expr.OfIdentifier(name))
	return x
}

func NewDataTypeWithSQLName(name expr.ISQLName, dbType db.Type) *SQLDataType {
	x := new(SQLDataType)
	x.abstractSQLDataType = NewAbstractSQLDataType(dbType)
	x.SetName(name)
	return x
}

func (x *SQLDataType) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDataType) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

func (x *SQLDataType) Precisions() []expr.ISQLExpr {
	return x.precisions
}

func (x *SQLDataType) AddPrecision(precision expr.ISQLExpr) {
	if precision == nil {
		return
	}
	precision.SetParent(x)
	x.precisions = append(x.precisions, precision)
}

/**
 * INTERVAL YEAR [ (year_precision) ]
 * INTERVAL YEAR [ (year_precision) ] TO MONTH
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#interval%20type
 *
 * INTERVAL YEAR [ (year_precision) ] TO MONTH
	| INTERVAL DAY [ (day_precision) ] TO SECOND [ (fractional_seconds_precision) ]
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Data-Types.html#GUID-A3C0D836-BADB-44E5-A5D4-265BA5968483
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

type SQLIntervalDataType struct {
	*abstractSQLDataType

	start *SQLIntervalDataTypeField
	end   *SQLIntervalDataTypeField
}

func NewIntervalDataType(dbType db.Type) *SQLIntervalDataType {
	x := new(SQLIntervalDataType)
	x.abstractSQLDataType = NewAbstractSQLDataType(dbType)
	return x
}

func (x *SQLIntervalDataType) Start() *SQLIntervalDataTypeField {
	return x.start
}

func (x *SQLIntervalDataType) SetStart(start *SQLIntervalDataTypeField) {
	x.start = start
}

func (x *SQLIntervalDataType) End() *SQLIntervalDataTypeField {
	return x.end
}

func (x *SQLIntervalDataType) SetEnd(end *SQLIntervalDataTypeField) {
	x.end = end
}

/**
 *  YEAR | MONTH | DAY | HOUR | MINUTE [ <left paren> <interval leading field precision> <right paren> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#start%20field
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#interval%20qualifier
 */
type SQLIntervalDataTypeField struct {
	*expr.AbstractSQLExpr
	unit       SQLIntervalUnit
	precisions []expr.ISQLExpr
}

func NewIntervalDataTypeFieldWitUnit(unit SQLIntervalUnit) *SQLIntervalDataTypeField {
	x := new(SQLIntervalDataTypeField)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.unit = unit
	return x
}

func NewIntervalDataTypeFieldWitUnitAndPrecisions(unit SQLIntervalUnit, precisions []expr.ISQLExpr) *SQLIntervalDataTypeField {
	x := new(SQLIntervalDataTypeField)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.unit = unit
	if precisions != nil {
		for _, precision := range precisions {
			x.AddPrecision(precision)
		}
	}
	return x
}

func (x *SQLIntervalDataTypeField) Unit() SQLIntervalUnit {
	return x.unit
}

func (x *SQLIntervalDataTypeField) SetEnd(unit SQLIntervalUnit) {
	x.unit = unit
}

func (x *SQLIntervalDataTypeField) Precisions() []expr.ISQLExpr {
	return x.precisions
}

func (x *SQLIntervalDataTypeField) AddPrecision(precision expr.ISQLExpr) {
	if precision == nil {
		return
	}
	precision.SetParent(x)
	x.precisions = append(x.precisions, precision)
}

// https://dev.mysql.com/doc/refman/8.0/en/data-types.html
// https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Data-Types.html#GUID-A3C0D836-BADB-44E5-A5D4-265BA5968483
const (
	// Character DataTypes
	CHAR      string = "CHAR"
	VARCHAR2         = "VARCHAR2"
	NCHAR            = "NCHAR"
	NVARCHAR2        = "NVARCHAR2"

	// Number DataTypes
	NUMBER        = "NUMBER"
	FLOAT         = "FLOAT"
	BINARY_FLOAT  = "BINARY_FLOAT"
	BINARY_DOUBLE = "BINARY_DOUBLE"

	// Long and Raw DataTypes
	LONG     = "LONG"
	LONG_RAW = "LONG_RAW"
	RAW      = "RAW"

	// DateTime DataTypes
	DATE = "DATE"

	// Large Object DataTypes
	BLOB  = "BLOB"
	CLOB  = "CLOB"
	NCLOB = "NCLOB"
	BFILE = "BFILE"

	// ROWID DataTypes
	ROWID  = "ROWID"
	UROWID = "UROWID"

	// ANSI Supported DataTypes
	CHARACTER         = "CHARACTER"
	CHARACTER_VARYING = "CHARACTER VARYING"

	CHAR_VARYING  = "CHAR VARYING"
	NCHAR_VARYING = "NCHAR VARYING"

	VARCHAR = "VARCHAR"

	NATIONAL_CHARACTER         = "NATIONAL CHARACTER"
	NATIONAL_CHAR              = "NATIONAL CHAR"
	NATIONAL_CHARACTER_VARYING = "NATIONAL CHARACTER VARYING"
	NATIONAL_CHAR_VARYING      = "NATIONAL CHAR VARYING"

	NUMERIC = "NUMERIC"
	DECIMAL = "DECIMAL"
	DEC     = "DEC"

	INTEGER  = "INTEGER"
	INT      = "INT"
	SMALLINT = "SMALLINT"

	DOUBLE_PRECISION = "DOUBLE PRECISION"
	REAL             = "REAL"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#character%20string%20type
//
// CHARACTER [ <left paren> <length> <right paren> ]
// CHAR [ <left paren> <length> <right paren> ]
// CHARACTER VARYING <left paren> <length> <right paren>
// CHAR VARYING <left paren> <length> <right paren>
// VARCHAR <left paren> <length> <right paren>
// CHARACTER LARGE OBJECT [ <left paren> <large object length> <right paren> ]
// CHAR LARGE OBJECT [ <left paren> <large object length> <right paren> ]
// CLOB [ <left paren> <large object length> <right paren> ]
type SQLCharacterStringDataType struct {
}

type SQLCharacterDataType struct {
}

type SQLCharDataType struct {
}

type SQLCharacterVaryingDataType struct {
}
type SQLCharVaryingDataType struct {
}

type SQLVarCharDataType struct {
}

type SQLCharacterLargeObjectDataType struct {
}

type SQLCharLargeObjectDataType struct {
}

type SQLClobDataType struct {
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#numeric%20type
// NUMERIC [ <left paren> <precision> [ <comma> <scale> ] <right paren> ]
// DECIMAL [ <left paren> <precision> [ <comma> <scale> ] <right paren> ]
// DEC [ <left paren> <precision> [ <comma> <scale> ] <right paren> ]
// SMALLINT
// INTEGER
// INT
// BIGINT
// FLOAT [ <left paren> <precision> <right paren> ]
// REAL
// DOUBLE PRECISION
type SQLNumericDataType struct {
}

type SQLDecimalDataType struct {
}
type SQLDecDataType struct {
}
type SQLSmallIntDataType struct {
}
type SQLIntegerDataType struct {
}
type SQLIntDataType struct {
}
type SQLBigIntDataType struct {
}
type SQLFloatDataType struct {
}
type SQLDoublePrecisionDataType struct {
}

/**
 *
 */
type SQLDateDataType struct {
	*abstractSQLDataType
}

func NewDateDataType(dbType db.Type) *SQLDateDataType {
	x := new(SQLDateDataType)
	x.abstractSQLDataType = NewAbstractSQLDataType(dbType)
	return x
}

func (dt *SQLDateDataType) String() string {
	return "DATE"
}

type SQLDateTimeDataType struct {
	*abstractSQLDataType
}

func NewDateTimeDataType(dbType db.Type) *SQLDateTimeDataType {
	x := new(SQLDateTimeDataType)
	x.abstractSQLDataType = NewAbstractSQLDataType(dbType)
	return x
}

func (dt *SQLDateTimeDataType) String() string {
	return "DATETIME"
}

type SQLTimeDataType struct {
	*abstractSQLDataType
}

func NewTimeDataType(dbType db.Type) *SQLTimeDataType {
	x := new(SQLTimeDataType)
	x.abstractSQLDataType = NewAbstractSQLDataType(dbType)
	return x
}

func (dt *SQLTimeDataType) String() string {
	return "TIME"
}

type SQLTimestampDataType struct {
	*abstractSQLDataType
}

func NewTimestampDataType(dbType db.Type) *SQLTimestampDataType {
	x := new(SQLTimestampDataType)
	x.abstractSQLDataType = NewAbstractSQLDataType(dbType)
	return x
}

func (dt *SQLTimestampDataType) String() string {
	return "TIMESTAMP"
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#path-resolved%20user-defined%20type%20name
type SQLUserDefinedDataType struct {
	name expr.ISQLName
}
