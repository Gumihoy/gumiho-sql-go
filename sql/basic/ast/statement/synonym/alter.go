package synonym

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER [ PUBLIC ] SYNONYM [ schema. ] synonym
  { EDITIONABLE | NONEDITIONABLE | COMPILE } ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-SYNONYM.html#GUID-C31B6804-6783-4A8C-B448-DF78E3FE6837
 */
type SQLAlterSynonymStatement struct {
	*statement.AbstractSQLStatement
	Public bool
	name expr.ISQLName
	action expr.ISQLExpr
}
func NewAlterSynonymStatement(dbType db.Type) *SQLAlterSynonymStatement {
	x:=new(SQLAlterSynonymStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLAlterSynonymStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLAlterSynonymStatement) SetName(name expr.ISQLName)  {
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterSynonymStatement) Action() expr.ISQLExpr {
	return x.action
}
func (x *SQLAlterSynonymStatement) SetAction(action expr.ISQLExpr)  {
	if action == nil {
		return
	}
	action.SetParent(x)
	x.action = action
}

