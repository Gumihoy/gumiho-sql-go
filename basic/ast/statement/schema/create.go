package schema

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/expr"
	"reflect"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#schema%20definition

// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/CREATE-SCHEMA.html#GUID-2D154F9C-9E2B-4A09-B658-2EA5B99AC838

// https://dev.mysql.com/doc/refman/8.0/en/create-database.html
// CREATE SCHEMA [IF NOT EXISTS] db_name
//    [create_option] ...
type SQLCreateSchemaStatement struct {
	IfNotExists       bool
	Name              expr.ISQLName
	AuthorizationName expr.ISQLName
	Elements          []ast.SQLObject
}

func NewCreateSchemaStatement() *SQLCreateSchemaStatement {
	return nil
}

func (x *SQLCreateSchemaStatement) replace(source expr.ISQLExpr, target expr.ISQLExpr) bool {

	if _, ok := target.(expr.ISQLName); reflect.DeepEqual(source, x.Name) && ok {
		x.Name = target.(expr.ISQLName)
		return true
	}

	return false
}

func (x *SQLCreateSchemaStatement) addElement(element ast.SQLObject) {
	x.Elements = append(x.Elements, element)
}
