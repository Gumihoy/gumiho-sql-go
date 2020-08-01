package util

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/parser"
	"gumihoy.com/sql/dbtype"
	mysqlParser "gumihoy.com/sql/dialect/mysql/parser"
	oracleParser "gumihoy.com/sql/dialect/oracle/parser"
)

func ParseStatementsByDBType(sql string, dbType dbtype.DBType) []ast.ISQLObject {
	return ParseStatementsByConfig(sql, parser.NewConfig(dbType))
}

func ParseStatementsByConfig(sql string, config parser.Config) []ast.ISQLObject {
	p := createParser(sql, config)
	return parser.ParseStatements(p)
}

func createParser(sql string, config parser.Config) parser.IParser {
	switch config.DBType {
	case dbtype.MySQL:
		return mysqlParser.NewParserBySQL(sql)
	case dbtype.Oracle:
		return oracleParser.NewParserBySQL(sql)
	default:
		return parser.NewParserBySQL(sql)
	}
}
