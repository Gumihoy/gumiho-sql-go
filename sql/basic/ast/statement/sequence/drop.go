package sequence

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP SEQUENCE <sequence generator name> <drop behavior>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20sequence%20generator%20statement
 *
 * DROP SEQUENCE [ schema. ] sequence_name ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DROP-SEQUENCE.html#GUID-32B640EE-47C9-46A7-9746-6125BAF8FF8C
 */
type SQLDropSequenceStatement struct {
	*statement.AbstractSQLStatement
	name         expr.ISQLName
	DropBehavior statement.SQLDropBehavior
}

func NewDropSequenceStatement(dbType db.Type) *SQLDropSequenceStatement {
	x := new(SQLDropSequenceStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropSequenceStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropSequenceStatement) SetName(name expr.ISQLName ) {
	name.SetParent(x)
	x.name = name
}
