package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/server"
)

func AcceptHostOption(visitor ISQLVisitor, x *server.SQLHostOption) {
	if visitor.VisitHostOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptDatabaseOption(visitor ISQLVisitor, x *server.SQLDatabaseOption) {
	if visitor.VisitDatabaseOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptUserOption(visitor ISQLVisitor, x *server.SQLUserOption) {
	if visitor.VisitUserOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptPasswordOption(visitor ISQLVisitor, x *server.SQLPasswordOption) {
	if visitor.VisitPasswordOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptSocketOption(visitor ISQLVisitor, x *server.SQLSocketOption) {
	if visitor.VisitSocketOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptOwnerOption(visitor ISQLVisitor, x *server.SQLOwnerOption) {
	if visitor.VisitOwnerOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptPortOption(visitor ISQLVisitor, x *server.SQLPortOption) {
	if visitor.VisitPortOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}

