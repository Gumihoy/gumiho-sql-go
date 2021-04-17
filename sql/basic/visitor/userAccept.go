package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/user"
)

func AcceptUserName(visitor ISQLVisitor, x *user.SQLUserName) {
	if visitor.VisitUserName(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Host())
	}
}
func AcceptIdentifiedByAuthOption(visitor ISQLVisitor, x *user.SQLIdentifiedByAuthOption) {
	if visitor.VisitIdentifiedByAuthOption(visitor, x) {
	}
}
func AcceptIdentifiedByPasswordAuthOption(visitor ISQLVisitor, x *user.SQLIdentifiedByPasswordAuthOption) {
	if visitor.VisitIdentifiedByPasswordAuthOption(visitor, x) {
	}
}
func AcceptIdentifiedByRandomPasswordAuthOption(visitor ISQLVisitor, x *user.SQLIdentifiedByRandomPasswordAuthOption) {
	if visitor.VisitIdentifiedByRandomPasswordAuthOption(visitor, x) {
	}
}
func AcceptIdentifiedWithAuthOption(visitor ISQLVisitor, x *user.SQLIdentifiedWithAuthOption) {
	if visitor.VisitIdentifiedWithAuthOption(visitor, x) {
	}
}
func AcceptIdentifiedWithByAuthOption(visitor ISQLVisitor, x *user.SQLIdentifiedWithByAuthOption) {
	if visitor.VisitIdentifiedWithByAuthOption(visitor, x) {
	}
}
func AcceptIdentifiedWithByRandomPasswordAuthOption(visitor ISQLVisitor, x *user.SQLIdentifiedWithByRandomPasswordAuthOption) {
	if visitor.VisitIdentifiedWithByRandomPasswordAuthOption(visitor, x) {
	}
}
func AcceptIdentifiedWithAsAuthOption(visitor ISQLVisitor, x *user.SQLIdentifiedWithAsAuthOption) {
	if visitor.VisitIdentifiedWithAsAuthOption(visitor, x) {
	}
}

