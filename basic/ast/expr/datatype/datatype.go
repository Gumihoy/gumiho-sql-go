package datatype

import "gumihoy.com/sql/basic/ast/expr"

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#data%20type
type ISQLDataType interface {
	expr.ISQLExpr
}

type SQLDataType struct {
	expr.SQLExpr
}

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

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#path-resolved%20user-defined%20type%20name
type SQLUserDefinedDataType struct {
	name expr.ISQLName
}
