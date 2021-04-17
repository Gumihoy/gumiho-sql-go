package typebody

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP TYPE BODY [ schema. ] type_name ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/DROP-TYPE-BODY-statement.html#GUID-4668B4DD-213D-452A-8706-F27D36C03D3C
 */
type SQLDropTypeBodyStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewDropTypeStatement(dbType db.Type) *SQLDropTypeBodyStatement {
	x := new(SQLDropTypeBodyStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLDropTypeBodyStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropTypeBodyStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}