package expr

import "gumihoy.com/sql/basic/ast"

type ISQLExpr interface {
	ast.ISQLObject
}

type ISQLName interface {
	Name() string
}

type SQLIdentifier struct {
}
