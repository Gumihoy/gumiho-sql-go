package config

type Lexer struct {
	KeepComment bool
}

func NewLexer() *Lexer {
	l := new(Lexer)
	return l
}

type Output struct {
	*Lexer
	LowerCase      bool
	PrintAfterSemi bool
}

func NewOutput() *Output {
	o := new(Output)
	return o
}
