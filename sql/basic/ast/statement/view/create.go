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
 * CREATE [ RECURSIVE ] VIEW <table name> <view specification> AS <query expression>
         [ WITH [ <levels clause> ] CHECK OPTION ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#view%20definition
 *
 * CREATE
    [OR REPLACE]
    [ALGORITHM = {UNDEFINED | MERGE | TEMPTABLE}]
    [DEFINER = user]
    [SQL SECURITY { DEFINER | INVOKER }]
    VIEW view_name [(column_list)]
    AS select_statement
    [WITH [CASCADED | LOCAL] CHECK OPTION]
 * https://dev.mysql.com/doc/refman/8.0/en/create-view.html
 *
 * CREATE [OR REPLACE]
  [[NO] FORCE]
  [ EDITIONING | EDITIONABLE [ EDITIONING ] | NONEDITIONABLE ]
  VIEW [schema.] view
  [ SHARING = { METADATA | DATA | EXTENDED DATA | NONE } ]
  [ ( { alias [ VISIBLE | INVISIBLE ] [ inline_constraint... ]
      | out_of_line_constraint
      }
        [, { alias [ VISIBLE | INVISIBLE ] [ inline_constraint...]
           | out_of_line_constraint
           }
        ]
    )
  | object_view_clause
  | XMLType_view_clause
  ]
  [ DEFAULT COLLATION collation_name ]
  [ BEQUEATH { CURRENT_USER | DEFINER } ]
  AS subquery [ subquery_restriction_clause ]
  [ CONTAINER_MAP | CONTAINERS_DEFAULT ] ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-VIEW.html#GUID-61D2D2B4-DACC-4C7C-89EB-7E50D9594D30
 */
type SQLCreateViewStatement struct {
	*statement.AbstractSQLStatement
	Recursive bool

	// MySQL
	OrReplace     bool
	algorithmExpr *expr.SQLArrayExpr
	definerExpr   *expr.SQLArrayExpr

	name     expr.ISQLName
	elements []view.ISQLViewElement

	subQuery select_.ISQLSelectQuery
}

func NewCreateViewStatement(dbType db.Type) *SQLCreateViewStatement {
	x := new(SQLCreateViewStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCreateViewStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLCreateViewStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLCreateViewStatement) ObjectType() db.SQLObjectType {
	return db.VIEW
}
func (x *SQLCreateViewStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLCreateViewStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLCreateViewStatement) Elements() []view.ISQLViewElement {
	return x.elements
}
func (x *SQLCreateViewStatement) AddElement(element view.ISQLViewElement) {
	if element == nil {
		return
	}
	element.SetParent(x)
	x.elements = append(x.elements, element)
}
func (x *SQLCreateViewStatement) AddElements(elements []view.ISQLViewElement) {
	if elements == nil || len(elements) == 0 {
		return
	}
	for _, element := range elements {
		x.AddElement(element)
	}
}

func (x *SQLCreateViewStatement) SubQuery() select_.ISQLSelectQuery {
	return x.subQuery
}

func (x *SQLCreateViewStatement) SetSubQuery(subQuery select_.ISQLSelectQuery) {
	if subQuery == nil {
		return
	}
	subQuery.SetParent(x)
	x.subQuery = subQuery
}
