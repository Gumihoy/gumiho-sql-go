package procedure

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER PROCEDURE proc_name [characteristic ...]

characteristic: {
    COMMENT 'string'
  | LANGUAGE SQL
  | { CONTAINS SQL | NO SQL | READS SQL DATA | MODIFIES SQL DATA }
  | SQL SECURITY { DEFINER | INVOKER }
}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-procedure.html
 */
type SQLAlterProcedureStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLName
	actions []expr.ISQLExpr
}
func NewAlterProcedureStatement(dbType db.Type) *SQLAlterProcedureStatement {
	x := new(SQLAlterProcedureStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLAlterProcedureStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLAlterProcedureStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterProcedureStatement) Actions() []expr.ISQLExpr {
	return x.actions
}
func (x *SQLAlterProcedureStatement) Action(i int) expr.ISQLExpr {
	return x.actions[i]
}
func (x *SQLAlterProcedureStatement) AddAction(action view.ISQLViewElement) {
	if action == nil {
		return
	}
	action.SetParent(x)
	x.actions = append(x.actions, action)
}
func (x *SQLAlterProcedureStatement) AddActions(actions []view.ISQLViewElement) {
	if actions == nil || len(actions) == 0 {
		return
	}
	for _, action := range actions {
		x.AddAction(action)
	}
}