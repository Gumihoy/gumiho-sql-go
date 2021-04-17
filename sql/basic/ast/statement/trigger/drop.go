package trigger

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP TRIGGER [IF EXISTS] [schema_name.]trigger_name
 * https://dev.mysql.com/doc/refman/8.0/en/drop-trigger.html
 *
 * DROP TRIGGER [ schema. ] trigger ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/DROP-TRIGGER-statement.html#GUID-C664FDA9-656B-49D4-A86D-D08B615137E9
 */
type SQLDropTriggerStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	name     expr.ISQLName
}

func NewDropTriggerStatement(dbType db.Type) *SQLDropTriggerStatement {
	x := new(SQLDropTriggerStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropTriggerStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLDropTriggerStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLDropTriggerStatement) ObjectType() db.SQLObjectType {
	return db.TRIGGER
}
func (x *SQLDropTriggerStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropTriggerStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
