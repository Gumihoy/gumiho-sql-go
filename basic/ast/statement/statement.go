package statement

import "gumihoy.com/sql/basic/ast"

type ISQLStatement interface {
	ast.ISQLObject
}

type SQLStatement struct {
}

type SQLSelectStatement struct {
}
