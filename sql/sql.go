package sql

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	mariadbParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mariadb/parser"
	mariadbVisitor "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mariadb/visitor"
	mysqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/parser"
	mysqlVisitor "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/visitor"
	oracleParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/oracle/parser"
	oracleVisitor "github.com/Gumihoy/gumiho-sql-go/sql/dialect/oracle/visitor"
	postgresqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/postgresql/parser"
	postgresqlVisitor "github.com/Gumihoy/gumiho-sql-go/sql/dialect/postgresql/visitor"
	tidbParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/tidb/parser"
	tidbVisitor "github.com/Gumihoy/gumiho-sql-go/sql/dialect/tidb/visitor"
	"log"
	"strings"
)

func Format(sourceSQL string, dbType db.Type) string {
	stmts := ParseStatementsByDBType(sourceSQL, dbType)
	return ToStringWithStatements(stmts, dbType)
}

func FormatWithConfig(sourceSQL string, dbType db.Type, config *config.SQLFormatConfig) string {
	if config == nil {
		panic("config is nil.")
	}
	stmts := ParseStatementsByConfig(sourceSQL, dbType, config.SQLOutputConfig.SQLParseConfig)
	return toStringWithStatementsAndConfig(stmts, dbType, config.SQLOutputConfig)
}

func ReplaceInParent(source expr.ISQLExpr, target expr.ISQLExpr) bool {
	return source.Parent().ReplaceChild(source, target)
}

func IsSameSQL(expectedSQL string, actualSQL string, dbType db.Type) bool {
	expectedNormalizeSQL := NormalizeSQL(expectedSQL, dbType)
	actualNormalizeSQL := NormalizeSQL(actualSQL, dbType)

	if strings.EqualFold(expectedNormalizeSQL, actualNormalizeSQL) {
		return true
	}

	return false
}

func NormalizeSQL(sourceSQL string, dbType db.Type) string {
	stmts := ParseStatementsByDBType(sourceSQL, dbType)

	normalizeVisitor := CreateNormalizeVisitor(dbType)
	for _, stmt := range stmts {
		visitor.Accept(normalizeVisitor, stmt)
	}

	return ToStringWithStatements(stmts, dbType)
}

func ToStringWithObject(object ast.ISQLObject, dbType db.Type) string {
	return "toSQLString(object)"
}

func ToStringWithObjects(objects []ast.ISQLObject, dbType db.Type) string {
	return toStringWithObjectsAndConfig(objects, dbType, config.NewOutputConfig())
}

func toStringWithObjectsAndConfig(objects []ast.ISQLObject, dbType db.Type, config *config.SQLOutputConfig) string {
	var builder strings.Builder
	visitor := CreateASTOutputVisitor(&builder, dbType, config)
	OutputSQLObjectVisitor(visitor, objects)
	return builder.String()
}

func ToStringWithStatements(stmts []statement.ISQLStatement, dbType db.Type) string {
	return toStringWithStatementsAndConfig(stmts, dbType, config.NewOutputConfig())
}
func toStringWithStatementsAndConfig(stmts []statement.ISQLStatement, dbType db.Type, config *config.SQLOutputConfig) string {
	var builder strings.Builder
	visitor := CreateASTOutputVisitor(&builder, dbType, config)
	OutputSQLStatementVisitor(visitor, stmts)
	return builder.String()
}

func OutputSQLObjectVisitor(v visitor.ISQLOutputVisitor, objects []ast.ISQLObject) {
	if objects == nil {
		log.Println("objects is nil")
		return
	}

	for i, object := range objects {
		if i != len(objects)-1 {
			object.SetAfterSemi(true)
		}
		if i != 0 {
			v.WriteLn()
		}
		visitor.Accept(v, object)
	}
}
func OutputSQLStatementVisitor(v visitor.ISQLOutputVisitor, objects []statement.ISQLStatement) {
	if objects == nil {
		log.Println("objects is nil")
		return
	}

	for i, object := range objects {
		if i != len(objects)-1 {
			object.SetAfterSemi(true)
		}
		if i != 0 {
			v.WriteLn()
		}
		visitor.Accept(v, object)
	}
}

func ParseStatementsByDBType(sourceSQL string, dbType db.Type) []statement.ISQLStatement {
	return ParseStatementsByConfig(sourceSQL, dbType, parser.NewParseConfig())
}

func ParseStatementsByConfig(sourceSQL string, dbType db.Type, config *parser.SQLParseConfig) []statement.ISQLStatement {
	p := createStatementParser(sourceSQL, dbType, config)
	return parser.ParseStatements(p)
}

func createStatementParser(sourceSQL string, dbType db.Type, config *parser.SQLParseConfig) parser.ISQLStatementParser {
	switch dbType {
	case db.MariaDB:
		return mariadbParser.NewStatementParserBySQL(sourceSQL, dbType, config)
	case db.MySQL:
		return mysqlParser.NewStatementParserBySQL(sourceSQL, dbType, config)
	case db.Oracle:
		return oracleParser.NewStatementParserBySQL(sourceSQL, dbType, config)
	case db.PostgreSQL:
		return postgresqlParser.NewStatementParserBySQL(sourceSQL, dbType, config)
	case db.TiDB:
		return tidbParser.NewStatementParserBySQL(sourceSQL, dbType, config)
	default:
		return parser.NewStatementParserBySQL(sourceSQL, dbType, config)
	}
}

func CreateASTOutputVisitor(builder *strings.Builder, dbType db.Type, config *config.SQLOutputConfig) visitor.ISQLOutputVisitor {
	switch dbType {
	case db.MariaDB:
		return mariadbVisitor.NewOutputVisitor(builder, config)
	case db.MySQL:
		return mysqlVisitor.NewOutputVisitor(builder, config)
	case db.Oracle:
		return oracleVisitor.NewOutputVisitor(builder, config)
	case db.PostgreSQL:
		return postgresqlVisitor.NewOutputVisitor(builder, config)
	case db.TiDB:
		return tidbVisitor.NewOutputVisitor(builder, config)
	default:
		return visitor.NewOutputVisitor(builder, config)
	}
}
