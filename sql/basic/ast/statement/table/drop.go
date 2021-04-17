package table

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP [TEMPORARY] TABLE [IF EXISTS] tbl_name [, tbl_name] ... [RESTRICT | CASCADE]
 * https://dev.mysql.com/doc/refman/8.0/en/drop-table.html
 */
type SQLDropTableStatement struct {
	*statement.AbstractSQLStatement

	Temporary bool
	IfExists  bool
	names []expr.ISQLName
	option expr.ISQLExpr
}

func NewDropTableStatement(dbType db.Type) *SQLDropTableStatement {
	x := new(SQLDropTableStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	x.names = make([]expr.ISQLName, 0, 10)
	return x
}
func (x *SQLDropTableStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLDropTableStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLDropTableStatement) ObjectType() db.SQLObjectType {
	return db.TABLE
}
func (x *SQLDropTableStatement) Name() expr.ISQLName {
	if len(x.names) == 0 {
		return nil
	}
	return x.names[0]
}

func (x *SQLDropTableStatement) Names() []expr.ISQLName {
	return x.names
}

func (x *SQLDropTableStatement) AddName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}


func (x *SQLDropTableStatement) Option() expr.ISQLExpr {
	return x.option
}

func (x *SQLDropTableStatement) SetOption(option expr.ISQLExpr) {
	option.SetParent(x)
	x.option = option
}