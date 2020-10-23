package table

import (
	"gumihoy.com/sql/basic/ast/expr"
	"gumihoy.com/sql/basic/ast/statement"
)

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20table%20statement
type SQLAlterTableStatement struct {
	statement.SQLStatement
	name    expr.ISQLName
	actions []ISQLAlterTableAction
}

func NewAlterTableStatement(name expr.ISQLName, actions ...ISQLAlterTableAction) *SQLAlterTableStatement {
	x := new(SQLAlterTableStatement)

	x.SetName(name)

	for _, action := range actions {
		action.SetParent(x)
	}
	x.actions = actions

	return x
}

func (x *SQLAlterTableStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLAlterTableStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

func (x *SQLAlterTableStatement) addAction(action ISQLAlterTableAction) {
	action.SetParent(x)
	x.actions = append(x.actions, action)
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20table%20action
type ISQLAlterTableAction interface {
	expr.ISQLExpr
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#add%20column%20definition
type SQLAddColumnAction struct {
	expr.SQLExpr
	hasColumn bool
	column    SQLTableColumn
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20column%20definition
type SQLAlterColumnAction struct {
	expr.SQLExpr
	hasColumn bool
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20column%20definition
type SQLDropColumnAction struct {
	expr.SQLExpr
	hasColumn bool
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#add%20table%20constraint%20definition
type SQLAddTableConstraintAction struct {
	expr.SQLExpr
	constraint SQLTableConstraint
}

// https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#drop%20table%20constraint%20definition
type SQLDropTableConstraintAction struct {
	expr.SQLExpr
}
