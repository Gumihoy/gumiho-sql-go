package view

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER
    [ALGORITHM = {UNDEFINED | MERGE | TEMPTABLE}]
    [DEFINER = user]
    [SQL SECURITY { DEFINER | INVOKER }]
    VIEW view_name [(column_list)]
    AS select_statement
    [WITH [CASCADED | LOCAL] CHECK OPTION]
 * https://dev.mysql.com/doc/refman/8.0/en/alter-view.html
 *
 * ALTER VIEW [ schema. ] view
  { ADD out_of_line_constraint
  | MODIFY CONSTRAINT constraint
      { RELY | NORELY }
  | DROP { CONSTRAINT constraint
         | PRIMARY KEY
         | UNIQUE (column [, column ]...)
         }
  | COMPILE
  | { READ ONLY | READ WRITE }
  | { EDITIONABLE | NONEDITIONABLE }
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLWithCheckOptionKind string

const (
	CASCADED SQLWithCheckOptionKind = "CASCADED"
	LOCAL                           = "LOCAL"
)

type SQLAlterViewStatement struct {
	*statement.AbstractSQLStatement

	algorithmExpr *expr.SQLAssignExpr
	definerExpr   *expr.SQLAssignExpr

	name     expr.ISQLName
	elements []view.ISQLViewElement

	// MySQL
	subQuery            select_.ISQLSelectQuery
	WithCheckOptionKind SQLWithCheckOptionKind

	// Oracle
	actions []expr.ISQLExpr
}

func NewAlterViewStatement(dbType db.Type) *SQLAlterViewStatement {
	x := new(SQLAlterViewStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLAlterViewStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}
func (x *SQLAlterViewStatement) Clone() ast.ISQLObject {
	panic("implement me")
}
func (x *SQLAlterViewStatement) ObjectType() db.SQLObjectType {
	return db.VIEW
}
func (x *SQLAlterViewStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLAlterViewStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterViewStatement) Elements() []view.ISQLViewElement {
	return x.elements
}
func (x *SQLAlterViewStatement) AddElement(element view.ISQLViewElement) {
	if element == nil {
		return
	}
	element.SetParent(x)
	x.elements = append(x.elements, element)
}
func (x *SQLAlterViewStatement) AddElements(elements []view.ISQLViewElement) {
	if elements == nil || len(elements) == 0 {
		return
	}
	for _, element := range elements {
		x.AddElement(element)
	}
}

func (x *SQLAlterViewStatement) SubQuery() select_.ISQLSelectQuery {
	return x.subQuery
}

func (x *SQLAlterViewStatement) SetSubQuery(subQuery select_.ISQLSelectQuery) {
	if subQuery == nil {
		return
	}
	subQuery.SetParent(x)
	x.subQuery = subQuery
}


func (x *SQLAlterViewStatement) Actions() []expr.ISQLExpr {
	return x.actions
}
func (x *SQLAlterViewStatement) Action(i int) expr.ISQLExpr {
	return x.actions[i]
}
func (x *SQLAlterViewStatement) AddAction(action view.ISQLViewElement) {
	if action == nil {
		return
	}
	action.SetParent(x)
	x.actions = append(x.actions, action)
}
func (x *SQLAlterViewStatement) AddActions(actions []view.ISQLViewElement) {
	if actions == nil || len(actions) == 0 {
		return
	}
	for _, action := range actions {
		x.AddAction(action)
	}
}