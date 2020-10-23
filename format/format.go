package format

import (
	"gumihoy.com/sql/basic"
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	"gumihoy.com/sql/db"
	"gumihoy.com/sql/factory"
	"strings"
)

func Format(sql string, dbType db.DBType) string {
	statements := factory.ParseStatementsByDBType(sql, dbType)
	return toStringWithObjects(statements, dbType)
}

func toStringWithObject(object ast.ISQLObject, dbType db.DBType) string {
	return "toSQLString(object)"
}

func toStringWithObjects(objects []ast.ISQLObject, dbType db.DBType) string {
	return toStringWithObjectsAndConfig(objects, dbType, *config.NewOutput())
}

func toStringWithObjectsAndConfig(objects []ast.ISQLObject, dbType db.DBType, config config.Output) string {
	var builder strings.Builder
	visitor := factory.CreateASTOutputVisitor(&builder, dbType, config)
	outputVisitor(objects, visitor)
	return builder.String()
}

func outputVisitor(objects []ast.ISQLObject, v visitor.ISQLOutputVisitor) {
	if objects == nil {
		panic("objects is nil")
	}
	for i, object := range objects {
		if i > 0 {
			preStmt := objects[i-1]
			if !preStmt.IsAfterSemi() {
				v.Write(basic.SYMB_SEMI.Upper)
			}
			v.WriteLn()
		}
		visitor.Accept(v, object)
	}
}

