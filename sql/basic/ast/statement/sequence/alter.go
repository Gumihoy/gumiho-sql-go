package sequence

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER SEQUENCE [ schema. ] sequence
  { INCREMENT BY integer
  | { MAXVALUE integer | NOMAXVALUE }
  | { MINVALUE integer | NOMINVALUE }
  | { CYCLE | NOCYCLE }
  | { CACHE integer | NOCACHE }
  | { ORDER | NOORDER }
  | { KEEP | NOKEEP }
  | { SCALE {EXTEND | NOEXTEND} | NOSCALE }
  | { SESSION | GLOBAL }
  } ...
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-SEQUENCE.html#GUID-A6468B63-E7C9-4EF0-B048-82FE2449B26D
 */
type SQLAlterSequenceStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
	options     []expr.ISQLExpr
}
func NewAlterSequenceStatement(dbType db.Type) *SQLAlterSequenceStatement {
	x := new(SQLAlterSequenceStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLAlterSequenceStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLAlterSequenceStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterSequenceStatement) Options() []expr.ISQLExpr {
	return x.options
}

func (x *SQLAlterSequenceStatement) Option(i int) expr.ISQLExpr {
	return x.options[i]
}

func (x *SQLAlterSequenceStatement) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}