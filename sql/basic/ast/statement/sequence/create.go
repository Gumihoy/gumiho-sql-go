package sequence

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE SEQUENCE <sequence generator name> [ <sequence generator options> ]
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#sequence%20generator%20definition
 *
 * CREATE SEQUENCE [ schema. ] sequence
  [ SHARING = { METADATA | DATA | NONE } ]
  [ { INCREMENT BY | START WITH } integer
  | { MAXVALUE integer | NOMAXVALUE }
  | { MINVALUE integer | NOMINVALUE }
  | { CYCLE | NOCYCLE }
  | { CACHE integer | NOCACHE }
  | { ORDER | NOORDER }
  | { KEEP | NOKEEP }
  | { SCALE {EXTEND | NOEXTEND} | NOSCALE }
  | { SESSION | GLOBAL }
  ]...
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-SEQUENCE.html#GUID-E9C78A8C-615A-4757-B2A8-5E6EFB130571
 */
type SQLCreateSequenceStatement struct {
	*statement.AbstractSQLStatement
	name        expr.ISQLName
	shareClause *expr.SQLAssignExpr
	options     []expr.ISQLExpr
}

func NewCreateSequenceStatement(dbType db.Type) *SQLCreateSequenceStatement {
	x := new(SQLCreateSequenceStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCreateSequenceStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateSequenceStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLCreateSequenceStatement) Options() []expr.ISQLExpr {
	return x.options
}

func (x *SQLCreateSequenceStatement) Option(i int) expr.ISQLExpr {
	return x.options[i]
}

func (x *SQLCreateSequenceStatement) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}