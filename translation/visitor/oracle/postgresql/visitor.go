package visitor

import "github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"

type Oracle2PostgreSQLASTTransformVisitor struct {
	*visitor.SQLVisitorAdapter
}

