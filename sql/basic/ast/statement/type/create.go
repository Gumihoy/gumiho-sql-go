package type_

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE [OR REPLACE] [ EDITIONABLE | NONEDITIONABLE ] TYPE [ schema. ] type_name [ FORCE ] [ OID 'object_identifier' ]
[ sharing_clause ] [ default_collation_clause ] { [ invoker_rights_clause ] |  [ accessible_by_clause ] }...
  { object_base_type_def | object_subtype_def } ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/CREATE-TYPE-statement.html#GUID-389D603D-FBD0-452A-8414-240BBBC57034
 */
type SQLCreateTypeStatement struct {
	*statement.AbstractSQLStatement
	OrReplace bool
	name expr.ISQLName
	Force bool

}
func NewCreateTypeStatement(dbType db.Type) *SQLCreateTypeStatement {
	x := new(SQLCreateTypeStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCreateTypeStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLCreateTypeStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLCreateTypeStatement) ObjectType() db.SQLObjectType {
	return db.TYPE
}
func (x *SQLCreateTypeStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateTypeStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
