package visitor

import "gumihoy.com/sql/basic/visitor"

type IMariaDBVisitor interface {
	visitor.ISQLVisitor
}

type MariaDBVisitorAdapter struct {
	IMariaDBVisitor
}

type MariaDBOutputVisitor struct {
	*MariaDBVisitorAdapter
}

func NewOutputVisitor() *MariaDBOutputVisitor {
	x := new(MariaDBOutputVisitor)
	return x
}
