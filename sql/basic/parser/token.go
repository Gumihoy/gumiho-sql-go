package parser

import (
	"strconv"
	"strings"
)

type SQLToken struct {
	Kind       *Kind
	Start, End int
	Line, Col  int
}

func NewToken(kind *Kind, start int, end int, Line int, Col int) *SQLToken {
	x := new(SQLToken)
	x.Kind = kind
	x.Start, x.End = start, end
	x.Line, x.Col = Line, Col
	return x
}

func (token *SQLToken) UnSupport() string {
	return "line " + strconv.Itoa(token.Line) + ", col " + strconv.Itoa(token.Col) + " UnSupport " + token.Kind.String()
}

func (token *SQLToken) Error() string {
	return "line " + strconv.Itoa(token.Line) + ", col " + strconv.Itoa(token.Col) + " Error"
}

func (token *SQLToken) String() string {
	return "Start " + strconv.Itoa(token.Start) +
		", End " + strconv.Itoa(token.End) +
		token.Kind.String()
}

type Kind struct {
	name string
}

func NewKind(name string) *Kind {
	return &Kind{name: name}
}

func (kind *Kind) Equals(t string) bool {
	return strings.EqualFold(kind.name, t)
}
func (kind *Kind) String() string {
	return kind.name
}

var (
	// A
	ADD            = NewKind("ADD")
	ADMIN          = NewKind("ADMIN")
	ALL            = NewKind("ALL")
	ALGORITHM      = NewKind("ALGORITHM")
	ALTER          = NewKind("ALTER")
	ALWAYS         = NewKind("ALWAYS")
	ANALYTIC       = NewKind("ANALYTIC")
	ANALYZE        = NewKind("ANALYZE")
	AND            = NewKind("AND")
	APPLY          = NewKind("APPLY")
	AS             = NewKind("AS")
	ASC            = NewKind("ASC")
	ATTACH         = NewKind("ATTACH")
	AUTO_INCREMENT = NewKind("AUTO_INCREMENT")
	AVG_ROW_LENGTH = NewKind("AVG_ROW_LENGTH")

	// B
	BETWEEN = NewKind("BETWEEN")
	BITMAP  = NewKind("BITMAP")
	BLOCK   = NewKind("BLOCK")
	BODY    = NewKind("BODY")
	BULK    = NewKind("BULK")
	BTREE   = NewKind("BTREE")
	BY      = NewKind("BY")

	// C

	CACHE    = NewKind("CACHE")
	CASCADE  = NewKind("CASCADE")
	CASCADED = NewKind("CASCADED")
	CAST     = NewKind("CAST")

	CHANGE    = NewKind("CHANGE")
	CHARACTER = NewKind("CHARACTER")
	CHARSET   = NewKind("CHARSET")

	CHECK = NewKind("CHECK")

	CHECKSUM = NewKind("CHECKSUM")
	CLUSTER  = NewKind("CLUSTER")

	COALESCE = NewKind("COALESCE")
	COLLATE  = NewKind("COLLATE")
	COLLECT  = NewKind("COLLECT")
	COLUMN   = NewKind("COLUMN")
	COLUMNS  = NewKind("COLUMNS")

	COMMENT     = NewKind("COMMENT")
	COMPILE     = NewKind("COMPILE")
	COMPRESSION = NewKind("COMPRESSION")

	CONNECT         = NewKind("CONNECT")
	CONNECT_BY_ROOT = NewKind("CONNECT_BY_ROOT")
	CONNECTION      = NewKind("CONNECTION")
	CONSTRAINT      = NewKind("CONSTRAINT")
	CONSTRAINTS     = NewKind("CONSTRAINTS")
	CONTAINS        = NewKind("CONTAINS")

	CREATE       = NewKind("CREATE")
	CROSS        = NewKind("CROSS")
	CURRENT      = NewKind("CURRENT")
	CURRENT_ROLE = NewKind("CURRENT_ROLE")
	CURRENT_USER = NewKind("CURRENT_USER")
	CYCLE        = NewKind("CYCLE")

	// D
	DATA      = NewKind("DATA")
	DATABASE  = NewKind("DATABASE")
	DATABASES = NewKind("DATABASES")
	DATE      = NewKind("DATE")
	DATETIME  = NewKind("DATETIME")
	DAY       = NewKind("DAY")

	DEFAULT = NewKind("DEFAULT")

	DELAY_KEY_WRITE = NewKind("DELAY_KEY_WRITE")
	DELAYED         = NewKind("DELAYED")
	DELETE          = NewKind("DELETE")
	DEFINER         = NewKind("DEFINER")

	DESC      = NewKind("DESC")
	DESCRIBE  = NewKind("DESCRIBE")
	DETACH    = NewKind("DETACH")
	DIRECTORY = NewKind("DIRECTORY")
	DISABLE   = NewKind("DISABLE")
	DISCARD   = NewKind("DISCARD")
	DISTINCT  = NewKind("DISTINCT")

	DIV = NewKind("DIV")

	DROP      = NewKind("DROP")
	DUPLICATE = NewKind("DUPLICATE")

	// E

	EDITIONABLE = NewKind("EDITIONABLE")
	ENABLE      = NewKind("ENABLE")
	ENCRYPTION  = NewKind("ENCRYPTION")
	ENGINE      = NewKind("ENGINE")
	EMPTY       = NewKind("EMPTY")
	EXCEPT      = NewKind("EXCEPT")
	EXCHANGE    = NewKind("EXCHANGE")
	EXISTS      = NewKind("EXISTS")
	EXPLAIN     = NewKind("EXPLAIN")
	EXTENDED    = NewKind("EXTENDED")
	ESCAPE      = NewKind("ESCAPE")
	EVENT       = NewKind("EVENT")

	// F
	FALSE    = NewKind("FALSE")
	FETCH    = NewKind("FETCH")
	FOR      = NewKind("FOR")
	FORCE    = NewKind("FORCE")
	FOREIGN  = NewKind("FOREIGN")
	FORMAT   = NewKind("FORMAT")
	FROM     = NewKind("FROM")
	FUNCTION = NewKind("FUNCTION")
	FULL     = NewKind("FULL")
	FULLTEXT = NewKind("FULLTEXT")

	// G
	GENERATED = NewKind("GENERATED")
	GLOBAL    = NewKind("GLOBAL")
	GROUP     = NewKind("GROUP")

	// H
	HASH          = NewKind("HASH")
	HAVING        = NewKind("HAVING")
	HELP          = NewKind("HELP")
	HIGH_PRIORITY = NewKind("HIGH_PRIORITY")
	HOST          = NewKind("HOST")
	HOUR          = NewKind("HOUR")

	// I
	IDENTIFIED = NewKind("IDENTIFIED")
	IDENTITY   = NewKind("IDENTITY")
	IF         = NewKind("IF")
	IMPORT     = NewKind("IMPORT")
	IN         = NewKind("IN")

	INCREMENT     = NewKind("INCREMENT")
	INDEX         = NewKind("INDEX")
	INHERIT       = NewKind("INHERIT")
	INFINITE      = NewKind("INFINITE")
	INNER         = NewKind("INNER")
	INOUT         = NewKind("INOUT")
	INTERVAL      = NewKind("INTERVAL")
	INTO          = NewKind("INTO")
	INSERT        = NewKind("INSERT")
	INSERT_METHOD = NewKind("INSERT_METHOD")
	INTERSECT     = NewKind("INTERSECT")
	INVISIBLE     = NewKind("INVISIBLE")
	IGNORE        = NewKind("IGNORE")
	IS            = NewKind("IS")

	// J
	JOIN = NewKind("JOIN")

	// K
	KEEP           = NewKind("KEEP")
	KEY            = NewKind("KEY")
	KEY_BLOCK_SIZE = NewKind("KEY_BLOCK_SIZE")

	// L
	LANGUAGE     = NewKind("LANGUAGE")
	LATERAL      = NewKind("LATERAL")
	LEFT         = NewKind("LEFT")
	LESS         = NewKind("LESS")
	LINEAR       = NewKind("LINEAR")
	LIMIT        = NewKind("LIMIT")
	LIKE         = NewKind("LIKE")
	LIKEC        = NewKind("LIKEC")
	LIKE2        = NewKind("LIKE2")
	LIKE4        = NewKind("LIKE4")
	LIST         = NewKind("LIST")
	LOCAL        = NewKind("LOCAL")
	LOCK         = NewKind("LOCK")
	LOCKED       = NewKind("LOCKED")
	LOW_PRIORITY = NewKind("LOW_PRIORITY")

	// M
	MATERIALIZED = NewKind("MATERIALIZED")
	MAX_ROWS     = NewKind("MAX_ROWS")
	MAXVALUE     = NewKind("MAXVALUE")

	MIN_ROWS = NewKind("MIN_ROWS")
	MINUS    = NewKind("MINUS")
	MINUTE   = NewKind("MINUTE")
	MINVALUE = NewKind("MINVALUE")
	MOD      = NewKind("MOD")
	MODE     = NewKind("MODE")
	MODEL    = NewKind("MODEL")
	MODIFIES = NewKind("MODIFIES")
	MODIFY   = NewKind("MODIFY")
	MONTH    = NewKind("MONTH")

	// N
	NAMES   = NewKind("NAMES")
	NAN     = NewKind("NAN")
	NATURAL = NewKind("NATURAL")
	NO      = NewKind("NO")

	NOCACHE = NewKind("NOCACHE")
	NOCYCLE = NewKind("NOCYCLE")
	NOKEEP  = NewKind("NOKEEP")

	NOMAXVALUE     = NewKind("NOMAXVALUE")
	NOMINVALUE     = NewKind("NOMINVALUE")
	NONEDITIONABLE = NewKind("NONEDITIONABLE")
	NOORDER        = NewKind("NOORDER")
	NORELY         = NewKind("NORELY")
	NOSCALE        = NewKind("NOSCALE")
	NOT            = NewKind("NOT")
	NOWAIT         = NewKind("NOWAIT")

	NULL = NewKind("NULL")

	// O
	OF       = NewKind("OF")
	OFFSET   = NewKind("OFFSET")
	OJ       = NewKind("OJ")
	ON       = NewKind("ON")
	ONLY     = NewKind("ONLY")
	OPTIMIZE = NewKind("OPTIMIZE")
	OPTION   = NewKind("OPTION")
	OPTIONS  = NewKind("OPTIONS")
	OR       = NewKind("OR")
	ORDER    = NewKind("ORDER")
	OUT      = NewKind("OUT")
	OUTER    = NewKind("OUTER")
	OWNER    = NewKind("OWNER")

	// P
	PACKAGE      = NewKind("PACKAGE")
	PACK_KEYS    = NewKind("PACK_KEYS")
	PARTITION    = NewKind("PARTITION")
	PARTITIONS   = NewKind("PARTITIONS")
	PASSWORD     = NewKind("PASSWORD")
	PERSIST      = NewKind("PERSIST")
	PERSIST_ONLY = NewKind("PERSIST_ONLY")
	PLAN         = NewKind("PLAN")
	PORT         = NewKind("PORT")
	PRIMARY      = NewKind("PRIMARY")
	PRIOR        = NewKind("PRIOR")
	PROCEDURE    = NewKind("PROCEDURE")
	PUBLIC       = NewKind("PUBLIC")

	// Q
	QUICK = NewKind("QUICK")

	// R
	RANDOM = NewKind("RANDOM")
	RANGE  = NewKind("RANGE")
	READ   = NewKind("READ")
	READS  = NewKind("READS")

	REBUILD    = NewKind("REBUILD")
	RECURSIVE  = NewKind("RECURSIVE")
	REF        = NewKind("REF")
	REFERENCES = NewKind("REFERENCES")

	REGEXP      = NewKind("REGEXP")
	REGEXP_LIKE = NewKind("REGEXP_LIKE")
	RELY        = NewKind("RELY")
	REMOVE      = NewKind("REMOVE")
	RENAME      = NewKind("RENAME")
	REPAIR      = NewKind("REPAIR")
	REPLACE     = NewKind("REPLACE")
	REPLICA     = NewKind("REPLICA")
	RESET       = NewKind("RESET")
	RESTRICT    = NewKind("RESTRICT")
	RESULT      = NewKind("RESULT")
	RETURN      = NewKind("RETURN")
	RETURNING   = NewKind("RETURNING")
	RETURNS     = NewKind("RETURNS")

	RIGHT      = NewKind("RIGHT")
	ROLE       = NewKind("ROLE")
	ROW_FORMAT = NewKind("ROW_FORMAT")

	// S
	SAMPLE   = NewKind("SAMPLE")
	SCALE    = NewKind("SCALE")
	SCOPE    = NewKind("SCOPE")
	SCHEMA   = NewKind("SCHEMA")
	SECOND   = NewKind("SECOND")
	SECURITY = NewKind("SECURITY")
	SEED     = NewKind("SEED")
	SELECT   = NewKind("SELECT")
	SEQUENCE = NewKind("SEQUENCE")

	SERVER  = NewKind("SERVER")
	SESSION = NewKind("SESSION")
	SET     = NewKind("SET")

	SHARE   = NewKind("SHARE")
	SHARING = NewKind("SHARING")
	SHOW    = NewKind("SHOW")
	SKIP    = NewKind("SKIP")
	SOCKET  = NewKind("SOCKET")
	SORT    = NewKind("SORT")
	SPATIAL = NewKind("SPATIAL")
	SQL     = NewKind("SQL")
	START   = NewKind("START")

	STATEMENT_ID       = NewKind("STATEMENT_ID")
	STATS_AUTO_RECALC  = NewKind("STATS_AUTO_RECALC")
	STATS_PERSISTENT   = NewKind("STATS_PERSISTENT")
	STATS_SAMPLE_PAGES = NewKind("STATS_SAMPLE_PAGES")
	STORAGE            = NewKind("STORAGE")
	STRAIGHT_JOIN      = NewKind("STRAIGHT_JOIN")

	SUBPARTITION  = NewKind("SUBPARTITION")
	SUBPARTITIONS = NewKind("SUBPARTITIONS")

	SYNONYM = NewKind("SYNONYM")

	// T
	TABLE      = NewKind("TABLE")
	TABLESPACE = NewKind("TABLESPACE")
	TEMPORARY  = NewKind("TEMPORARY")
	TIME       = NewKind("TIME")
	TIMESTAMP  = NewKind("TIMESTAMP")
	THAN       = NewKind("THAN")
	TO         = NewKind("TO")
	TRIGGER    = NewKind("TRIGGER")
	TRUE       = NewKind("TRUE")
	TRUNCATE   = NewKind("TRUNCATE")
	TYPE       = NewKind("TYPE")

	// U
	UNION   = NewKind("UNION")
	UNIQUE  = NewKind("UNIQUE")
	UNKNOWN = NewKind("UNKNOWN")
	UNNEST  = NewKind("UNNEST")
	UPDATE  = NewKind("UPDATE")
	USE     = NewKind("USE")
	USER    = NewKind("USER")
	USING   = NewKind("USING")

	// V
	VALIDATE = NewKind("VALIDATE")
	VALUE    = NewKind("VALUE")
	VALUES   = NewKind("VALUES")
	VIEW     = NewKind("VIEW")
	VISIBLE  = NewKind("VISIBLE")

	// W
	WAIT = NewKind("WAIT")

	WHERE   = NewKind("WHERE")
	WITH    = NewKind("WITH")
	WRAPPER = NewKind("WRAPPER")
	WRITE   = NewKind("WRITE")

	// X
	// Y
	YEAR = NewKind("YEAR")
	// Z

	// Identifier
	IDENTIFIER               = NewKind("IDENTIFIER")
	IDENTIFIER_DOUBLE_QUOTE  = NewKind("IDENTIFIER_DOUBLE_QUOTE")
	IDENTIFIER_REVERSE_QUOTE = NewKind("IDENTIFIER_REVERSE_QUOTE")

	// Literal
	LITERAL_STRING         = NewKind("LITERAL_STRING")
	LITERAL_INTEGER        = NewKind("LITERAL_INTEGER")
	LITERAL_FLOATING_POINT = NewKind("LITERAL_FLOATING_POINT")
	LITERAL_HEXADECIMAL_X  = NewKind("LITERAL_HEXADECIMAL_X")
	LITERAL_HEXADECIMAL_0X = NewKind("LITERAL_HEXADECIMAL_0X")
	LITERAL_DATETIME       = NewKind("LITERAL_DATETIME")
	LITERAL_INTERVAL       = NewKind("LITERAL_INTERVAL")
	LITERAL_NULL           = NewKind("LITERAL_NULL")

	// Operators
	// Operators. Assigns
	// :=
	SYMB_COLON_EQUAL = NewKind(":=")
	// +=
	SYMB_PLUS_EQUAL = NewKind("+=")
	// -=
	SYMB_MINUS_EQUAL = NewKind("-=")
	// *=
	SYMB_MULT_EQUAL = NewKind("*=")
	// /=
	SYMB_DIV_EQUAL = NewKind("/=")

	// %=
	SYMB_MOD_EQUAL = NewKind("%=")

	// &=
	SYMB_AND_EQUAL = NewKind("INTERSECT")

	// ^=
	SYMB_XOR_EQUAL = NewKind("^=")

	// |=
	SYMB_OR_EQUAL = NewKind("|=")

	// Operators. Arithmetics
	// +
	SYMB_PLUS = NewKind("+")

	// -
	SYMB_MINUS = NewKind("-")
	// *
	SYMB_STAR = NewKind("*")

	// /
	SYMB_SLASH = NewKind("/")

	// %
	SYMB_PERCENT = NewKind("%")

	// --
	SYMB_MINUSMINUS = NewKind("--")

	// Operators. Comparation
	// =
	SYMB_EQUAL = NewKind("=")

	// =>
	SYMB_EQUAL_GREATER_THAN = NewKind("=>")

	// >
	SYMB_GREATER_THAN = NewKind(">")

	// >>
	SYMB_GREATER_THAN_GREATER_THAN = NewKind(">>")
	// >=
	SYMB_GREATER_THAN_EQUAL = NewKind(">=")

	// <
	SYMB_LESS_THAN = NewKind("<")
	// <<
	SYMB_LESS_THAN_LESS_THAN = NewKind("<<")
	// <=
	SYMB_LESS_THAN_EQUAL = NewKind("<=")

	// !
	SYMB_EXCLAMATION = NewKind("!")
	// !=
	SYMB_EXCLAMATION_EQUAL = NewKind("!=")
	// <>
	SYMB_LESS_THAN_GREATER_THAN = NewKind("<>")

	// <=>
	SYMB_LESS_THAN_EQUAL_GREATER_THAN = NewKind("<>")

	// ~=
	SYMB_NOT_EQUAL = NewKind("~=")

	// Operators. Bit
	// ~
	SYMB_BIT_NOT = NewKind("~")
	// |
	SYMB_BIT_OR = NewKind("|")
	// &
	SYMB_BIT_AND = NewKind("&")
	// ^
	SYMB_BIT_XOR = NewKind("^")

	// Constructors symbols
	// .
	SYMB_DOT = NewKind(".")
	// @
	SYMB_AT       = NewKind("@")
	SYMB_QUESTION = NewKind("?")

	// (
	SYMB_LEFT_PAREN = NewKind("(")
	// )
	SYMB_RIGHT_PAREN = NewKind(")")
	// [
	SYMB_LERT_BRACKET = NewKind("[")
	// ]
	SYMB_RIGHT_BRACKET = NewKind("]")
	// {
	SYMB_LERT_BRACE = NewKind("{")
	// }
	SYMB_RIGHT_BRACE = NewKind("}")

	// ,
	SYMB_COMMA = NewKind(",")
	// ;
	SYMB_SEMI = NewKind(";")

	// :
	SYMB_COLON = NewKind(":")
	// _
	SYMB_INTRODUCER = NewKind("_")

	// ||
	SYMB_LOGICAL_OR = NewKind("||")
	// &&
	SYMB_LOGICAL_AND = NewKind("&&")

	// Comment
	COMMENT_MINUS           = NewKind("-- comment")
	COMMENT_MULTI_LINE      = NewKind("/* comment */")
	COMMENT_SHARP           = NewKind("# comment")
	COMMENT_MULTI_LINE_HINT = NewKind("/*+ hint */")

	EOF = NewKind("EOF")
)
