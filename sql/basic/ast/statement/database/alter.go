package database

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER DATABASE [db_name]
    alter_option ...
alter_option: {
    [DEFAULT] CHARACTER SET [=] charset_name
  | [DEFAULT] COLLATE [=] collation_name
  | [DEFAULT] ENCRYPTION [=] {'Y' | 'N'}
  | READ ONLY [=] {DEFAULT | 0 | 1}
}
 * https://dev.mysql.com/doc/refman/8.0/en/alter-database.html
 *
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-DATABASE.html#GUID-8069872F-E680-4511-ADD8-A4E30AF67986
 */
type SQLAlterDatabaseStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLExpr
	actions []expr.ISQLExpr
}

func NewAlterDatabaseStatement(dbType db.Type) *SQLAlterDatabaseStatement {
	x := new(SQLAlterDatabaseStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLAlterDatabaseStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLAlterDatabaseStatement) SetName(name expr.ISQLExpr) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterDatabaseStatement) Actions() []expr.ISQLExpr {
	return x.actions
}
func (x *SQLAlterDatabaseStatement) Action(i int) expr.ISQLExpr {
	return x.actions[i]
}
func (x *SQLAlterDatabaseStatement) AddAction(action view.ISQLViewElement) {
	if action == nil {
		return
	}
	action.SetParent(x)
	x.actions = append(x.actions, action)
}
func (x *SQLAlterDatabaseStatement) AddActions(actions []view.ISQLViewElement) {
	if actions == nil || len(actions) == 0 {
		return
	}
	for _, action := range actions {
		x.AddAction(action)
	}
}