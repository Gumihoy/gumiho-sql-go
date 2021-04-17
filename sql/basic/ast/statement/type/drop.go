package type_

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP TYPE [ schema. ] type_name [ FORCE | VALIDATE ] ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/DROP-TYPE-statement.html#GUID-EDB83D12-26B5-42D1-9DB6-BD8AAB6490EC
 */
type SQLDropTypeStatement struct {
	*statement.AbstractSQLStatement
	name     expr.ISQLName
	Behavior statement.SQLDropBehavior
}

func NewDropTypeStatement(dbType db.Type) *SQLDropTypeStatement {
	x := new(SQLDropTypeStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropTypeStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLDropTypeStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLDropTypeStatement) ObjectType() db.SQLObjectType {
	return db.TYPE
}
func (x *SQLDropTypeStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropTypeStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
