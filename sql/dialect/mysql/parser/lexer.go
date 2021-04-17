package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
)

type MySQLLexer struct {
	*parser.SQLLexer
}

func NewLexer(sourceSQL string) *MySQLLexer {
	x := new(MySQLLexer)
	x.SQLLexer = parser.NewLexer(sourceSQL)
	return x
}

func (lexer *MySQLLexer) ScanLeftBrace() *parser.Kind {
	lexer.ScanChar()
	return parser.SYMB_LERT_BRACE
}

func (lexer *MySQLLexer) ScanRightBrace() *parser.Kind {
	lexer.ScanChar()
	return parser.SYMB_RIGHT_BRACE
}

// https://dev.mysql.com/doc/refman/8.0/en/identifiers.html
// ASCII: [0-9,a-z,A-Z$_] (basic Latin letters, digits 0-9, dollar, underscore)
// Extended: U+0080 .. U+FFFF
func (lexer *MySQLLexer) IsSQLIdentifierStartWithCh(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		(ch == '$') ||
		(ch == '_') ||
		(ch >= 0x0080 && ch <= 0xFFFF && ch != '（' && ch != '）' && ch != '，')
}
func (lexer *MySQLLexer) IsSQLIdentifierPartWithCh(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9') ||
		(ch == '$') ||
		(ch == '_') ||
		(ch >= 0x0080 && ch <= 0xFFFF && ch != '（' && ch != '）' && ch != '，')
}

// ""
// "xxx"" xxx"  => xxx" xxx
// "xxx" " xxx"  => xxx xxx
func (lexer *MySQLLexer) ScanDoubleQuota() *parser.Kind {
	for {
		lexer.ScanString()
		if lexer.Ch() == '"' {
			lexer.PutCharWithCh('\\')
			lexer.PutCharWithCh('"')
			continue
		}
		if parser.IsWhitespace(lexer.Ch()) {
			lexer.ScanChar()
		}
		if lexer.Ch() == '"' || lexer.Ch() == '\'' {
			continue
		}
		break
	}

	return parser.IDENTIFIER_DOUBLE_QUOTE
}

// ``
func (lexer *MySQLLexer) ScanReverseQuota() *parser.Kind {
	lexer.ScanString()
	return parser.IDENTIFIER_REVERSE_QUOTE
}


/*
 * /
 * /*
 * /*+
 */
func (lexer *MySQLLexer) ScanSlash() *parser.Kind {
	kind := parser.SYMB_SLASH

	lexer.ScanChar()

	if lexer.Ch() == '*' {
		lexer.ScanChar()
		kind = parser.COMMENT_MULTI_LINE

		if lexer.Ch() == '+' {
			lexer.ScanChar()
			kind = parser.COMMENT_MULTI_LINE_HINT
		}
	}

	return kind
}

/**
 * #
 */
func (lexer *MySQLLexer) ScanSharp(child parser.ISQLLexer) *parser.Kind {
	lexer.ScanChar()
	return parser.COMMENT_SHARP
}



var KindMap = make(map[string]*parser.Kind)

func init() {

	// A
	// KindMap["ACCESSIBLE"] = parser.ACCESSIBLE
	// KindMap["ADD"] = parser.ADD
	KindMap["ALL"] = parser.ALL
	KindMap["ALTER"] = parser.ALTER
	// KindMap["ANALYZE"] = parser.ANALYZE
	KindMap["AND"] = parser.AND
	KindMap["AS"] = parser.AS
	KindMap["ASC"] = parser.ASC
	// KindMap["ASENSITIVE"] = parser.ASENSITIVE

	// B
	// KindMap["BEFORE"] = parser.BEFORE
	// KindMap["BETWEEN"] = parser.BETWEEN
	// KindMap["BIGINT"] = parser.BIGINT
	// KindMap["BINARY"] = parser.BINARY
	// KindMap["BLOB"] = parser.BLOB
	// KindMap["BOTH"] = parser.BOTH
	KindMap["BY"] = parser.BY

	// C
	// KindMap["CALL"] = parser.CALL
	// KindMap["CASCADE"] = parser.CASCADE
	// KindMap["CASE"] = parser.CASE
	// KindMap["CHANGE"] = parser.CHANGE
	// KindMap["CHAR"] = parser.CHAR
	// KindMap["CHARACTER"] = parser.CHARACTER
	// KindMap["CHECK"] = parser.CHECK
	// KindMap["COLLATE"] = parser.COLLATE
	// KindMap["COLUMN"] = parser.COLUMN
	// KindMap["CONDITION"] = parser.CONDITION
	// KindMap["CONSTRAINT"] = parser.CONSTRAINT
	// KindMap["CONTINUE"] = parser.CONTINUE
	// KindMap["CONVERT"] = parser.CONVERT
	// KindMap["CREATE"] = parser.CREATE
	KindMap["CROSS"] = parser.CROSS
	// KindMap["CUBE"] = parser.CUBE
	// KindMap["CUME_DIST"] = parser.CUME_DIST
	// KindMap["CURRENT_DATE"] = parser.CURRENT_DATE
	// KindMap["CURRENT_TIME"] = parser.CURRENT_TIME
	// KindMap["CURRENT_TIMESTAMP"] = parser.CURRENT_TIMESTAMP
	// KindMap["CURRENT_USER"] = parser.CURRENT_USER
	// KindMap["CURSOR"] = parser.CURSOR

	// D
	// KindMap["DATABASE"] = parser.DATABASE
	// KindMap["DATABASES"] = parser.DATABASES
	// KindMap["DAY_HOUR"] = parser.DAY_HOUR
	// KindMap["DAY_MICROSECOND"] = parser.DAY_MICROSECOND
	// KindMap["DAY_MINUTE"] = parser.DAY_MINUTE
	// KindMap["DAY_SECOND"] = parser.DAY_SECOND
	// KindMap["DEC"] = parser.DEC
	// KindMap["DECIMAL"] = parser.DECIMAL
	// KindMap["DECLARE"] = parser.DECLARE
	// KindMap["DEFAULT"] = parser.DEFAULT
	// KindMap["DELAYED"] = parser.DELAYED
	// KindMap["DELETE"] = parser.DELETE
	// KindMap["DENSE_RANK"] = parser.DENSE_RANK
	// KindMap["DESC"] = parser.DESC
	// KindMap["DESCRIBE"] = parser.DESCRIBE
	// KindMap["DETERMINISTIC"] = parser.DETERMINISTIC
	// KindMap["DISTINCT"] = parser.DISTINCT
	// KindMap["DISTINCTROW"] = parser.DISTINCTROW
	// KindMap["DIV"] = parser.DIV
	// KindMap["DOUBLE"] = parser.DOUBLE
	// KindMap["DROP"] = parser.DROP
	// KindMap["DUAL"] = parser.DUAL

	// E
	// KindMap["EACH"] = parser.EACH
	// KindMap["ELSE"] = parser.ELSE
	// KindMap["ELSEIF"] = parser.ELSEIF
	// KindMap["EMPTY"] = parser.EMPTY
	// KindMap["ENCLOSED"] = parser.ENCLOSED
	// KindMap["ESCAPED"] = parser.ESCAPED
	// KindMap["EXCEPT"] = parser.EXCEPT
	// KindMap["EXISTS"] = parser.EXISTS
	// KindMap["EXIT"] = parser.EXIT
	// KindMap["EXPLAIN"] = parser.EXPLAIN

	// F
	KindMap["FALSE"] = parser.FALSE
	// KindMap["FETCH"] = parser.FETCH
	// KindMap["FIRST_VALUE"] = parser.FIRST_VALUE
	// KindMap["FLOAT"] = parser.FLOAT
	// KindMap["FLOAT4"] = parser.FLOAT4
	// KindMap["FLOAT8"] = parser.FLOAT8
	KindMap["FOR"] = parser.FOR
	// KindMap["FORCE"] = parser.FORCE
	// KindMap["FOREIGN"] = parser.FOREIGN
	// KindMap["FROM"] = parser.FROM
	// KindMap["FULLTEXT"] = parser.FULLTEXT
	// KindMap["FUNCTION"] = parser.FUNCTION

	// G

	// KindMap["GENERATED"] = parser.GENERATED
	// KindMap["GET"] = parser.GET
	KindMap["GLOBAL"] = parser.GLOBAL
	// KindMap["GRANT"] = parser.GRANT
	KindMap["GROUP"] = parser.GROUP
	// KindMap["GROUPING"] = parser.GROUPING
	// KindMap["GROUPS"] = parser.GROUPS

	// H
	KindMap["HAVING"] = parser.HAVING
	// KindMap["HIGH_PRIORITY"] = parser.HIGH_PRIORITY
	// KindMap["HOUR_MICROSECOND"] = parser.HOUR_MICROSECOND
	// KindMap["HOUR_MINUTE"] = parser.HOUR_MINUTE
	// HOUR_SECOND

	// I
	KindMap["IF"] = parser.IF
	// IGNORE
	KindMap["IN"] = parser.IN
	KindMap["INDEX"] = parser.INDEX
	// INFILE
	KindMap["INNER"] = parser.INNER
	// INOUT
	// INSENSITIVE
	// INSERT
	// INT
	// INT1
	// INT2
	// INT3
	// INT4
	// INT8
	// INTEGER
	KindMap["INTERVAL"] = parser.INTERVAL
	// INTO
	// IO_AFTER_GTIDS
	// IO_BEFORE_GTIDS
	// IS
	// ITERATE

	// J
	KindMap["JOIN"] = parser.JOIN
	// JSON_TABLE

	// K
	// KEY
	// KEYS
	// KILL

	// L
	// LAG
	// LAST_VALUE
	// LATERAL
	// LEAD
	// LEADING
	// LEAVE

	KindMap["LEFT"] = parser.LEFT
	// LIKE
	KindMap["LIMIT"] = parser.LIMIT
	// LINEAR
	// LINES
	// LOAD
	// LOCALTIME
	// LOCALTIMESTAMP
	// LOCK
	KindMap["LOCK"] = parser.LOCK
	// LONG
	// LONGBLOB
	// LONGTEXT
	// LOOP
	// LOW_PRIORITY

	// M
	// MASTER_BIND
	// MASTER_SSL_VERIFY_SERVER_CERT
	// MATCH
	// MAXVALUE
	// MEDIUMBLOB
	// MEDIUMINT
	// MEDIUMTEXT
	// MIDDLEINT
	// MINUTE_MICROSECOND
	// MINUTE_SECOND
	// MOD
	// MODIFIES

	// N
	KindMap["NATURAL"] = parser.NATURAL
	// NOT
	// NO_WRITE_TO_BINLOG
	// NTH_VALUE
	// NTILE
	// NULL
	// NUMERIC

	// O
	// OF
	KindMap["ON"] = parser.ON
	// OPTIMIZE
	// OPTIMIZER_COSTS
	// OPTION
	// OPTIONALLY

	KindMap["OR"] = parser.OR
	KindMap["ORDER"] = parser.ORDER
	// OUT
	// OUTER
	// OUTFILE
	// OVER

	// P
	KindMap["PARTITION"] = parser.PARTITION
	// PERCENT_RANK
	// PRECISION
	// PRIMARY
	// PROCEDURE
	// PURGE

	// R
	// RANGE
	// RANK
	// READ
	// READS
	// READ_WRITE
	// REAL
	// RECURSIVE
	// REFERENCES
	// REGEXP
	// RELEASE
	// RENAME
	// REPEAT
	// REPLACE
	// REQUIRE
	// RESIGNAL
	// RESTRICT
	// RETURN
	// REVOKE
	KindMap["RIGHT"] = parser.RIGHT
	// RLIKE
	// ROW
	// ROWS
	// ROW_NUMBER

	// S
	// SCHEMA
	// SCHEMAS
	// SECOND_MICROSECOND

	KindMap["SELECT"] = parser.SELECT

	// SENSITIVE
	// SEPARATOR
	KindMap["SET"] = parser.SET
	// SHOW
	// SIGNAL
	// SMALLINT
	// SPATIAL
	// SPECIFIC
	// SQL
	// SQLEXCEPTION
	// SQLSTATE
	// SQLWARNING
	// SQL_BIG_RESULT
	// SQL_CALC_FOUND_ROWS
	// SQL_SMALL_RESULT
	// SSL
	// STARTING
	// STORED
	KindMap["STRAIGHT_JOIN"] = parser.STRAIGHT_JOIN
	// SYSTEM
	// TABLE
	// TERMINATED
	// THEN
	// TINYBLOB
	// TINYINT
	// TINYTEXT
	// TO
	// TRAILING
	// TRIGGER
	KindMap["TRUE"] = parser.TRUE
	// UNDO
	// UNION
	KindMap["UNION"] = parser.UNION
	// UNIQUE
	// UNLOCK
	// UNSIGNED
	// UPDATE
	// USAGE
	// USE
	KindMap["USING"] = parser.USING
	// UTC_DATE
	// UTC_TIME
	// UTC_TIMESTAMP
	// VALUES
	// VARBINARY
	// VARCHAR
	// VARCHARACTER
	// VARYING
	// VIRTUAL

	// KindMap["WHEN"] = parser.WHEN
	KindMap["WHERE"] = parser.WHERE
	// KindMap["WHILE"] = parser.WHILE
	// KindMap["WINDOW"] = parser.WINDOW
	KindMap["WITH"] = parser.WITH
	// KindMap["WRITE"] = parser.WRITE
	// KindMap["XOR"] = parser.XOR
	// KindMap["YEAR_MONTH"] = parser.YEAR_MONTH
	// KindMap["ZEROFILL"] = parser.ZEROFILL
}

func (lexer *MySQLLexer) GetKindMap() map[string]*parser.Kind {
	return KindMap
}
