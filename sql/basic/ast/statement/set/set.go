package set

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)
/**
 * SET DEFAULT ROLE
    {NONE | ALL | role [, role ] ...}
    TO user [, user ] ...
 * https://dev.mysql.com/doc/refman/8.0/en/set-default-role.html
 */
type SQLSetDefaultRoleStatement struct {
	*statement.AbstractSQLStatement

}











/**
 * SET variable = expr [, variable = expr] ...
 * https://dev.mysql.com/doc/refman/8.0/en/set-variable.html
 */
type SQLSetVariableAssignmentStatement struct {
	*statement.AbstractSQLStatement
	elements []*expr.SQLAssignExpr
}

func NewSetVariableAssignmentStatement(dbType db.Type) *SQLSetVariableAssignmentStatement {
	x := new(SQLSetVariableAssignmentStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLSetVariableAssignmentStatement) Elements() []*expr.SQLAssignExpr {
	return x.elements
}
func (x *SQLSetVariableAssignmentStatement) Element(i int) *expr.SQLAssignExpr {
	return x.elements[i]
}
func (x *SQLSetVariableAssignmentStatement) AddElement(element *expr.SQLAssignExpr) {
	if element == nil {
		return
	}
	element.SetParent(x)
	x.elements = append(x.elements, element)
}

/**
 * SET {CHARACTER SET | CHARSET}
    {'charset_name' | DEFAULT}
 * https://dev.mysql.com/doc/refman/8.0/en/set-character-set.html
 */
type SQLSetCharacterSetStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLExpr
}

func NewSetCharacterSetStatement(dbType db.Type) *SQLSetCharacterSetStatement {
	x := new(SQLSetCharacterSetStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLSetCharacterSetStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLSetCharacterSetStatement) SetName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

type SQLSetCharsetStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLExpr
}

func NewSetCharsetStatement(dbType db.Type) *SQLSetCharsetStatement {
	x := new(SQLSetCharsetStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLSetCharsetStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLSetCharsetStatement) SetName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 * SET NAMES {'charset_name' [COLLATE 'collation_name'] | DEFAULT}
 * https://dev.mysql.com/doc/refman/8.0/en/set-names.html
 */
type SQLSetNamesStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLExpr
}

func NewSetNamesStatement(dbType db.Type) *SQLSetNamesStatement {
	x := new(SQLSetNamesStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLSetNamesStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLSetNamesStatement) SetName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}






