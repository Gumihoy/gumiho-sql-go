package view

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP VIEW <table name> <drop behavior>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20view%20statement
 *
 * DROP VIEW [IF EXISTS]
    view_name [, view_name] ...
    [RESTRICT | CASCADE]
 * https://dev.mysql.com/doc/refman/8.0/en/drop-view.html
 *
 * DROP VIEW [ schema. ] view [ CASCADE CONSTRAINTS ] ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DROP-VIEW.html#GUID-1A1BD841-66B9-47E4-896F-D36E075AE296
 */
type SQLDropViewStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	names    []expr.ISQLName
	Behavior statement.SQLDropBehavior
}

func NewDropViewStatement(dbType db.Type) *SQLDropViewStatement {
	x := new(SQLDropViewStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropViewStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLDropViewStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLDropViewStatement) ObjectType() db.SQLObjectType {
	return db.VIEW
}
func (x *SQLDropViewStatement) Names() []expr.ISQLName {
	return x.names
}
func (x *SQLDropViewStatement) Name(i int) expr.ISQLName {
	return x.names[i]
}
func (x *SQLDropViewStatement) AddName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}
