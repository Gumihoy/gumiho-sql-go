package trigger

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type SQLAlterTriggerStatement struct {
	*statement.AbstractSQLStatement
}
func (x *SQLAlterTriggerStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLAlterTriggerStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLAlterTriggerStatement) ObjectType() db.SQLObjectType {
	return db.TRIGGER
}
