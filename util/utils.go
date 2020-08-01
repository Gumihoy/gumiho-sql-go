package util

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/parser"
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	"gumihoy.com/sql/dbtype"
	mysqlVisitor "gumihoy.com/sql/dialect/mysql/visitor"
	oracleVisitor "gumihoy.com/sql/dialect/oracle/visitor"
	postgresqlVisitor "gumihoy.com/sql/dialect/postgresql/visitor"
	"strings"
)

func Format(sql string, dbType dbtype.DBType) string {
	statements := ParseStatementsByDBType(sql, dbType)
	return toStringWithObjects(statements, dbType)
}

func toStringWithObject(object ast.ISQLObject, dbType dbtype.DBType) string {
	return "toSQLString(object)"
}

func toStringWithObjects(objects []ast.ISQLObject, dbType dbtype.DBType) string {
	return toStringWithObjectsAndConfig(objects, dbType, *config.NewOutput())
}

func toStringWithObjectsAndConfig(objects []ast.ISQLObject, dbType dbtype.DBType, config config.Output) string {
	var builder strings.Builder
	visitor := createASTOutputVisitor(builder, dbType, config)
	outputVisitor(objects, visitor)
	return builder.String()
}

func outputVisitor(objects []ast.ISQLObject, visitor visitor.ISQLOutputVisitor) {
	for i, object := range objects {
		if i > 0 {
			preStmt := objects[i-1]
			if !preStmt.IsAfterSemi() {
				visitor.Write(parser.SEMI)
			}
			visitor.WriteLn()
		}
		object.Accept(visitor)
	}
}

func createASTOutputVisitor(builder strings.Builder, dbType dbtype.DBType, config config.Output) visitor.ISQLOutputVisitor {
	switch dbType {
	case dbtype.Oracle:

		return oracleVisitor.NewOutputVisitor(builder, config)
	case dbtype.PostgreSQL:
		return postgresqlVisitor.NewOutputVisitor(builder, config)
	case dbtype.MySQL:
		return mysqlVisitor.NewOutputVisitor(builder, config)
	default:
		return visitor.NewOutputVisitor(builder, config)
	}
}
