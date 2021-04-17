package schema

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"reflect"
)

/**
 * CREATE SCHEMA <schema name clause> [ <schema character set or path> ] [ <schema element> ... ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#schema%20definition
 *
 * CREATE SCHEMA [IF NOT EXISTS] db_name
//    [create_option] ...
 * https://dev.mysql.com/doc/refman/8.0/en/create-database.html
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/CREATE-SCHEMA.html#GUID-2D154F9C-9E2B-4A09-B658-2EA5B99AC838
 */
type SQLCreateSchemaStatement struct {
	*statement.AbstractSQLStatement
	IfNotExists       bool
	name              expr.ISQLName
	AuthorizationName expr.ISQLName
	options           []expr.ISQLExpr
}

func NewCreateSchemaStatement(dbType db.Type) *SQLCreateSchemaStatement {
	x := new(SQLCreateSchemaStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	x.options = make([]expr.ISQLExpr, 0, 10)
	return x
}

func (x *SQLCreateSchemaStatement) replace(source expr.ISQLExpr, target expr.ISQLExpr) bool {

	if _, ok := target.(expr.ISQLName); reflect.DeepEqual(source, x.Name) && ok {
		x.SetName(target.(expr.ISQLName))
		return true
	}

	return false
}

func (x *SQLCreateSchemaStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLCreateSchemaStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

func (x *SQLCreateSchemaStatement) Options() []expr.ISQLExpr {
	return x.options
}

func (x *SQLCreateSchemaStatement) AddOption(option expr.ISQLName) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
