package config

type LexerConfig struct {
	KeepComment bool
}

func NewLexer() *LexerConfig {
	l := new(LexerConfig)
	return l
}

type Output struct {
	*LexerConfig
	LowerCase      bool
	PrintAfterSemi bool
}

func NewOutput() *Output {
	o := new(Output)
	return o
}
