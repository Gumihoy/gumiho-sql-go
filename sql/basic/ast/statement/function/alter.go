package function

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER FUNCTION func_name [characteristic ...]

characteristic: {
    COMMENT 'string'
  | LANGUAGE SQL
  | { CONTAINS SQL | NO SQL | READS SQL DATA | MODIFIES SQL DATA }
  | SQL SECURITY { DEFINER | INVOKER }
}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-function.html
 */
type SQLAlterFunctionStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLName
	actions []expr.ISQLExpr
}

func NewAlterFunctionStatement(dbType db.Type) *SQLAlterFunctionStatement {
	x := new(SQLAlterFunctionStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLAlterFunctionStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLAlterFunctionStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterFunctionStatement) Actions() []expr.ISQLExpr {
	return x.actions
}
func (x *SQLAlterFunctionStatement) Action(i int) expr.ISQLExpr {
	return x.actions[i]
}
func (x *SQLAlterFunctionStatement) AddAction(action view.ISQLViewElement) {
	if action == nil {
		return
	}
	action.SetParent(x)
	x.actions = append(x.actions, action)
}
func (x *SQLAlterFunctionStatement) AddActions(actions []view.ISQLViewElement) {
	if actions == nil || len(actions) == 0 {
		return
	}
	for _, action := range actions {
		x.AddAction(action)
	}
}
