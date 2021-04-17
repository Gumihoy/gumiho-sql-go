package synonym

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)


/**
 * DROP [PUBLIC] SYNONYM [ schema. ] synonym [FORCE] ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DROP-SYNONYM.html#GUID-C7293D40-83B8-4E60-9E90-CB907F2CA6C7
 */
type SQLDropSynonymStatement struct {
	*statement.AbstractSQLStatement
	Public bool
	name expr.ISQLName
	Force bool
}

func NewDropSynonymStatement(dbType db.Type) *SQLDropSynonymStatement {
	x:=new(SQLDropSynonymStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropSynonymStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropSynonymStatement) SetName(name expr.ISQLName)  {
	name.SetParent(x)
	x.name = name
}
