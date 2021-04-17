package trigger

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type SQLCreateTriggerStatement struct {
	*statement.AbstractSQLStatement
}
func (x *SQLCreateTriggerStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLCreateTriggerStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLCreateTriggerStatement) ObjectType() db.SQLObjectType {
	return db.TRIGGER
}
