package kind

type Kind string

const (
	// A
	ALTER Kind = "ALTER"
	// B
	// C
	CREATE Kind = "CREATE"
	// D
	DELETE Kind = "DELETE"
	DROP   Kind = "DROP"
	// E
	// F
	// G
	// H
	// I
	INSERT Kind = "INSERT"
	// J
	// K
	// L
	// M
	// N
	// O
	// P
	// Q
	// R

	// S
	SELECT Kind = "SELECT"
	// T
	// U
	UPDATE Kind = "UPDATE"
	// V
	// W
	// X
	// Y
	// Z

	// Constructors symbols
	QUOTE_DOUBLE  = "\""
	QUOTE_REVERSE = "`"

	IDENTIFIER_DOUBLE_QUOTE
	IDENTIFIER_REVERSE_QUOTE

	LITERAL_TEXT
	LITERAL_INTEGER
	LITERAL_FLOATING_POINT
	LITERAL_DATETIME
	LITERAL_INTERVAL

	SEMI = ";"
	EOF
)
