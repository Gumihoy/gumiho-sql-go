package config

import "github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"

type SQLFormatConfig struct {
	*SQLOutputConfig
}
func NewFormatConfig() *SQLFormatConfig {
	x := new(SQLFormatConfig)
	x.SQLOutputConfig = NewOutputConfig()
	return x
}

type SQLOutputConfig struct {
	*parser.SQLParseConfig
	LowerCase     bool
	FillAfterSemi bool
}

func NewOutputConfig() *SQLOutputConfig {
	x := new(SQLOutputConfig)
	x.SQLParseConfig = parser.NewParseConfig()
	return x
}
