package type_

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type SQLAlterTypeStatement struct {
	*statement.AbstractSQLStatement
}
func (x *SQLAlterTypeStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLAlterTypeStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLAlterTypeStatement) ObjectType() db.SQLObjectType {
	return db.TYPE
}
