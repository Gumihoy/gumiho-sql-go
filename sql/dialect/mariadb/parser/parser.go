package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	mysqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/parser"
)

type MariaDBParser struct {
	*mysqlParser.MySQLParser
}

func NewParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MariaDBParser {
	return &MariaDBParser{mysqlParser.NewParserBySQL(sql, dbType, config)}
}

func NewParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MariaDBParser {
	return &MariaDBParser{mysqlParser.NewParserByLexer(lexer, dbType, config)}
}
