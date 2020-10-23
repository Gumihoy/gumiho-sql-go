package visitor

import "gumihoy.com/sql/basic/visitor"

type SQLASTTransformVisitor struct {
	*visitor.SQLVisitorAdapter
}