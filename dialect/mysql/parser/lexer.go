package parser

import (
	"gumihoy.com/sql/basic/parser"
)

type MySQLLexer struct {
	*parser.SQLLexer
}

func NewLexer(sql string) *MySQLLexer {
	return &MySQLLexer{parser.NewLexer(sql)}
}



// https://dev.mysql.com/doc/refman/8.0/en/identifiers.html
// ASCII: [0-9,a-z,A-Z$_] (basic Latin letters, digits 0-9, dollar, underscore)
// Extended: U+0080 .. U+FFFF
func (lexer *MySQLLexer) IsSQLIdentifierStart() bool {
	return (lexer.Ch() >= 'a' && lexer.Ch() <= 'z') ||
		(lexer.Ch() >= 'A' && lexer.Ch() <= 'Z') ||
		(lexer.Ch() >= '0' && lexer.Ch() <= '9') ||
		(lexer.Ch() == '$') ||
		(lexer.Ch() == '_') ||
		(lexer.Ch() >= 0x0080 && lexer.Ch() <= 0xFFFF)
}

func (lexer *MySQLLexer) IsSQLIdentifierPart() bool {
	return (lexer.Ch() >= 'a' && lexer.Ch() <= 'z') ||
		(lexer.Ch() >= 'A' && lexer.Ch() <= 'Z') ||
		(lexer.Ch() >= '0' && lexer.Ch() <= '9') ||
		(lexer.Ch() == '$') ||
		(lexer.Ch() == '_') ||
		(lexer.Ch() >= 0x0080 && lexer.Ch() <= 0xFFFF)
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

var KindMap = make(map[string]*parser.Kind)

func init() {
	// KindMap["ACCESSIBLE"] = parser.ACCESSIBLE
	// KindMap["ADD"] = parser.ADD
	KindMap["ALL"] = parser.ALL
	KindMap["ALTER"] = parser.ALTER
	// KindMap["ANALYZE"] = parser.ANALYZE
	KindMap["AND"] = parser.AND
	KindMap["AS"] = parser.AS
	KindMap["ASC"] = parser.ASC
	// KindMap["ASENSITIVE"] = parser.ASENSITIVE


	// KindMap["BEFORE"] = parser.BEFORE
	// KindMap["BETWEEN"] = parser.BETWEEN
	// KindMap["BIGINT"] = parser.BIGINT
	// KindMap["BINARY"] = parser.BINARY
	// KindMap["BLOB"] = parser.BLOB
	// KindMap["BOTH"] = parser.BOTH
	KindMap["BY"] = parser.BY


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
	// KindMap["CROSS"] = parser.CROSS
	// KindMap["CUBE"] = parser.CUBE
	// KindMap["CUME_DIST"] = parser.CUME_DIST
	// KindMap["CURRENT_DATE"] = parser.CURRENT_DATE
	// KindMap["CURRENT_TIME"] = parser.CURRENT_TIME
	// KindMap["CURRENT_TIMESTAMP"] = parser.CURRENT_TIMESTAMP
	// KindMap["CURRENT_USER"] = parser.CURRENT_USER
	// KindMap["CURSOR"] = parser.CURSOR


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


	// KindMap["FALSE"] = parser.FALSE
	// KindMap["FETCH"] = parser.FETCH
	// KindMap["FIRST_VALUE"] = parser.FIRST_VALUE
	// KindMap["FLOAT"] = parser.FLOAT
	// KindMap["FLOAT4"] = parser.FLOAT4
	// KindMap["FLOAT8"] = parser.FLOAT8
	// KindMap["FOR"] = parser.FOR
	// KindMap["FORCE"] = parser.FORCE
	// KindMap["FOREIGN"] = parser.FOREIGN
	// KindMap["FROM"] = parser.FROM
	// KindMap["FULLTEXT"] = parser.FULLTEXT
	// KindMap["FUNCTION"] = parser.FUNCTION
	// KindMap["GENERATED"] = parser.GENERATED
	// KindMap["GET"] = parser.GET
	// KindMap["GRANT"] = parser.GRANT
	// KindMap["GROUP"] = parser.GROUP
	// KindMap["GROUPING"] = parser.GROUPING
	// KindMap["GROUPS"] = parser.GROUPS
	// KindMap["HAVING"] = parser.HAVING
	// KindMap["HIGH_PRIORITY"] = parser.HIGH_PRIORITY
	// KindMap["HOUR_MICROSECOND"] = parser.HOUR_MICROSECOND
	// KindMap["HOUR_MINUTE"] = parser.HOUR_MINUTE


	// HOUR_SECOND
	// IF
	// IGNORE
	// IN
	// INDEX
	// INFILE
	// INNER
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
	// INTERVAL
	// INTO
	// IO_AFTER_GTIDS
	// IO_BEFORE_GTIDS
	// IS
	// ITERATE
	// JOIN
	// JSON_TABLE
	// KEY
	// KEYS
	// KILL
	// LAG
	// LAST_VALUE
	// LATERAL
	// LEAD
	// LEADING
	// LEAVE
	// LEFT
	// LIKE
	// LIMIT
	// LINEAR
	// LINES
	// LOAD
	// LOCALTIME
	// LOCALTIMESTAMP
	// LOCK
	// LONG
	// LONGBLOB
	// LONGTEXT
	// LOOP
	// LOW_PRIORITY
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
	// NATURAL
	// NOT
	// NO_WRITE_TO_BINLOG
	// NTH_VALUE
	// NTILE
	// NULL
	// NUMERIC
	// OF
	// ON
	// OPTIMIZE
	// OPTIMIZER_COSTS
	// OPTION
	// OPTIONALLY
	// OR
	// ORDER
	// OUT
	// OUTER
	// OUTFILE
	// OVER
	// PARTITION
	// PERCENT_RANK
	// PRECISION
	// PRIMARY
	// PROCEDURE
	// PURGE
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
	// RIGHT
	// RLIKE
	// ROW
	// ROWS
	// ROW_NUMBER
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
	// STRAIGHT_JOIN
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
	// TRUE
	// UNDO
	// UNION
	// UNIQUE
	// UNLOCK
	// UNSIGNED
	// UPDATE
	// USAGE
	// USE
	// USING
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