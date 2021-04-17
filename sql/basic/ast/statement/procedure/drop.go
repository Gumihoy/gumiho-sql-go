package procedure

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP {PROCEDURE | FUNCTION} [IF EXISTS] sp_name
 * https://dev.mysql.com/doc/refman/8.0/en/drop-procedure.html

 * DROP PROCEDURE [ schema. ] procedure ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/DROP-PROCEDURE-statement.html#GUID-A66E96E6-A582-4A59-B97E-8E6EA52EC292
 */
type SQLDropProcedureStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	name     expr.ISQLName
}
func NewDropProcedureStatement(dbType db.Type) *SQLDropProcedureStatement {
	x := new(SQLDropProcedureStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLDropProcedureStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropProcedureStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}