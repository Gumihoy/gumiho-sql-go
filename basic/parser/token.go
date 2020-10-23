package parser

import (
	"strconv"
	"strings"
)

type Token struct {
	Kind       *Kind
	Start, End int
	Line, Col  int
}

func (token *Token) UnSupport() string {
	return "line " + strconv.Itoa(token.Line) + ", col " + strconv.Itoa(token.Col) + " UnSupport " + token.Kind.String()
}

func (token *Token) Error() string {
	return "line " + strconv.Itoa(token.Line) + ", col " + strconv.Itoa(token.Col) + " Error."
}

func (token *Token) String() string {

	return "Start " + strconv.Itoa(token.Start) +
		", End " + strconv.Itoa(token.End) +
		token.Kind.String()
}

func NewToken(start int, end int, kind *Kind) *Token {
	return &Token{Start: start, End: end, Kind: kind}
}

func NewTokenByLexer(lexer *SQLLexer, kind *Kind) *Token {
	return &Token{Start: lexer.line, End: lexer.col, Kind: kind}
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
	ALL   = NewKind("ALL")
	ALTER = NewKind("ALTER")
	APPLY = NewKind("APPLY")
	AS    = NewKind("AS")
	ASC   = NewKind("ASC")

	// B
	BULK = NewKind("BULK")
	BY   = NewKind("BY")

	// C
	COLLECT = NewKind("COLLECT")
	CREATE  = NewKind("CREATE")
	CROSS   = NewKind("CROSS")
	CURRENT = NewKind("CURRENT")

	// D
	DATE     = NewKind("DATE")
	DELETE   = NewKind("DELETE")
	DESC     = NewKind("DESC")
	DISTINCT = NewKind("DISTINCT")
	DIV      = NewKind("DIV")
	DROP     = NewKind("DROP")

	// E
	EXCEPT = NewKind("EXCEPT")

	// F
	FETCH = NewKind("FETCH")
	FOR   = NewKind("FOR")
	FROM  = NewKind("FROM")
	FULL  = NewKind("FULL")

	// G
	GROUP = NewKind("GROUP")

	// H
	HAVING = NewKind("HAVING")
	// I
	IN        = NewKind("IN")
	INNER     = NewKind("INNER")
	INTO      = NewKind("INTO")
	INSERT    = NewKind("INSERT")
	INTERSECT = NewKind("INTERSECT")

	// J
	JOIN = NewKind("JOIN")

	// K
	KEY = NewKind("KEY")

	// L
	LATERAL = NewKind("LATERAL")
	LEFT    = NewKind("LEFT")
	LIMIT   = NewKind("LIMIT")
	LIKE    = NewKind("LIKE")
	LIKE2   = NewKind("LIKE2")
	LIKE4   = NewKind("LIKE4")
	LOCK    = NewKind("LOCK")
	LOCKED  = NewKind("LOCKED")

	// M
	MINUS = NewKind("MINUS")
	MOD   = NewKind("MOD")
	MODE  = NewKind("MODE")

	// N
	NATURAL = NewKind("NATURAL")
	NO      = NewKind("NO")
	NOWAIT  = NewKind("NOWAIT")

	// O
	OF     = NewKind("INTERSECT")
	OFFSET = NewKind("INTERSECT")
	ONLY   = NewKind("INTERSECT")
	ORDER  = NewKind("INTERSECT")
	OUTER  = NewKind("INTERSECT")

	// P
	// Q
	// R
	RETURN    = NewKind("RETURN")
	RETURNING = NewKind("RETURNING")
	RIGHT     = NewKind("RIGHT")

	// S
	SELECT = NewKind("SELECT")
	SET    = NewKind("SET")
	SHARE  = NewKind("SHARE")
	SKIP   = NewKind("SKIP")
	SORT   = NewKind("SORT")

	// T
	TABLE     = NewKind("TABLE")
	TIME      = NewKind("TIME")
	TIMESTAMP = NewKind("TIMESTAMP")

	// U

	UNION  = NewKind("UNION")
	UNNEST = NewKind("UNNEST")
	UPDATE = NewKind("UPDATE")
	USING  = NewKind("USING")

	// V

	// W

	WAIT  = NewKind("WAIT")
	WHERE = NewKind("WHERE")
	WITH  = NewKind("WITH")

	// X
	// Y
	// Z

	// Identifier
	IDENTIFIER               = NewKind("")
	IDENTIFIER_DOUBLE_QUOTE  = NewKind("\"")
	IDENTIFIER_REVERSE_QUOTE = NewKind("`")

	// Literal
	LITERAL_STRING         = NewKind("LITERAL_STRING")
	LITERAL_INTEGER        = NewKind("LITERAL_INTEGER")
	LITERAL_FLOATING_POINT = NewKind("LITERAL_FLOATING_POINT")
	LITERAL_DATETIME       = NewKind("LITERAL_DATETIME")
	LITERAL_INTERVAL       = NewKind("LITERAL_INTERVAL")
	LITERAL_BOOLEAN_TRUE   = NewKind("LITERAL_BOOLEAN_TRUE")
	LITERAL_BOOLEAN_FALSE  = NewKind("LITERAL_BOOLEAN_FALSE")
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
	SYMB_LESS_GREATER_THAN = NewKind("<>")
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
	SYMB_RIGHT_PAREN = NewKind("(")
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

	// logical
	AND = NewKind("AND")
	OR  = NewKind("OR")
	// ||
	SYMB_LOGICAL_OR = NewKind("||")

	COMMENT = NewKind("")
	EOF     = NewKind("")
)
