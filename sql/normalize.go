package sql

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/condition"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/literal"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/variable"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

func CreateNormalizeVisitor(dbType db.Type) visitor.ISQLVisitor {
	switch dbType {
	case db.Oracle, db.MySQL, db.MariaDB, db.TiDB:
		return NewSQLNormalizeVisitor()
	}

	panic("UnSupport dbType:" + dbType)
}

/**
 *
 */
type SQLNormalizeVisitor struct {
	*visitor.SQLVisitorAdapter
}

func NewSQLNormalizeVisitor() *SQLNormalizeVisitor {
	x := new(SQLNormalizeVisitor)
	x.SQLVisitorAdapter = visitor.NewVisitorAdapter()
	return x
}
func (v *SQLNormalizeVisitor) VisitInCondition(child visitor.ISQLVisitor, x *condition.SQLInCondition) bool {
	if len(x.Values()) > 1 {
		target := condition.NewInCondition()
		target.SetExpr(x.Expr())
		target.AddValue(variable.NewVariableExpr())
		ReplaceInParent(x, target)
		return false
	}
	return true
}
func (v *SQLNormalizeVisitor) VisitStringLiteral(child visitor.ISQLVisitor, x *literal.SQLStringLiteral) bool {
	target := variable.NewVariableExpr()
	ReplaceInParent(x, target)
	return false
}

func (v *SQLNormalizeVisitor) VisitCharacterStringLiteral(child visitor.ISQLVisitor, x *literal.SQLCharacterStringLiteral) bool {
	return false
}

func (v *SQLNormalizeVisitor) VisitIntegerLiteral(child visitor.ISQLVisitor, x *literal.SQLIntegerLiteral) bool {
	target := variable.NewVariableExpr()
	ReplaceInParent(x, target)
	return false
}

func (v *SQLNormalizeVisitor) VisitFloatingPointLiteral(child visitor.ISQLVisitor, x *literal.SQLFloatingPointLiteral) bool {
	target := variable.NewVariableExpr()
	ReplaceInParent(x, target)
	return false
}