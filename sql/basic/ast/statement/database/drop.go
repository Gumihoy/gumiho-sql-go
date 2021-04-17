package database

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP DATABASE [IF EXISTS] db_name
 * https://dev.mysql.com/doc/refman/8.0/en/drop-database.html
 *
 * DROP DATABASE ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DROP-DATABASE.html#GUID-4FFC1AF5-538D-4882-8979-7A9957492A23
 */
type SQLDropDatabaseStatement struct {
	*statement.AbstractSQLStatement

	IfExists  bool
	name expr.ISQLName
}

func NewDropDatabaseSStatement(dbType db.Type) *SQLDropDatabaseStatement {
	x := new(SQLDropDatabaseStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLDropDatabaseStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropDatabaseStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}
