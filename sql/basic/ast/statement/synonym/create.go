package synonym

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE [ OR REPLACE ] [ EDITIONABLE | NONEDITIONABLE ]
   [ PUBLIC ] SYNONYM
   [ schema. ] synonym
   [ SHARING = { METADATA | NONE } ]
   FOR [ schema. ] object [ @ dblink ] ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-SYNONYM.html#GUID-A806C82F-1171-478E-A910-F9C6C42739B2
 */
type SQLCreateSynonymStatement struct {
	*statement.AbstractSQLStatement
	OrReplace bool

	Public        bool
	name          expr.ISQLName
	sharingClause *expr.SQLAssignExpr
	forName       expr.ISQLName
}

func NewCreateSynonymStatement(dbType db.Type) *SQLCreateSynonymStatement {
	x := new(SQLCreateSynonymStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCreateSynonymStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateSynonymStatement) SetName(name expr.ISQLName)  {
	name.SetParent(x)
	x.name = name
}
func (x *SQLCreateSynonymStatement) ForName() expr.ISQLName {
	return x.forName
}
func (x *SQLCreateSynonymStatement) SetForName(forName expr.ISQLName)  {
	forName.SetParent(x)
	x.forName = forName
}