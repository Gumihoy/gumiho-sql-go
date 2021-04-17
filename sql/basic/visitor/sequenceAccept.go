package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/sequence"
)

func AcceptIncrementBySequenceOption(visitor ISQLVisitor, x *sequence.SQLIncrementBySequenceOption) {
	if visitor.VisitIncrementBySequenceOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptStartWithSequenceOption(visitor ISQLVisitor, x *sequence.SQLStartWithSequenceOption) {
	if visitor.VisitStartWithSequenceOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptMaxValueSequenceOption(visitor ISQLVisitor, x *sequence.SQLMaxValueSequenceOption) {
	if visitor.VisitMaxValueSequenceOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptNoMaxValueSequenceOption(visitor ISQLVisitor, x *sequence.SQLNoMaxValueSequenceOption) {
	if visitor.VisitNoMaxValueSequenceOption(visitor, x) {

	}
}
func AcceptMinValueSequenceOption(visitor ISQLVisitor, x *sequence.SQLMinValueSequenceOption) {
	if visitor.VisitMinValueSequenceOption(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptNoMinValueSequenceOption(visitor ISQLVisitor, x *sequence.SQLNoMinValueSequenceOption) {
	if visitor.VisitNoMinValueSequenceOption(visitor, x) {
	}
}
func AcceptCycleSequenceOption(visitor ISQLVisitor, x *sequence.SQLCycleSequenceOption) {
	if visitor.VisitCycleSequenceOption(visitor, x) {
	}
}
func AcceptNoCycleSequenceOption(visitor ISQLVisitor, x *sequence.SQLNoCycleSequenceOption) {
	if visitor.VisitNoCycleSequenceOption(visitor, x) {
	}
}
func AcceptCacheSequenceOption(visitor ISQLVisitor, x *sequence.SQLCacheSequenceOption) {
	if visitor.VisitCacheSequenceOption(visitor, x) {
	}
}
func AcceptNoCacheSequenceOption(visitor ISQLVisitor, x *sequence.SQLNoCacheSequenceOption) {
	if visitor.VisitNoCacheSequenceOption(visitor, x) {
	}
}
func AcceptOrderSequenceOption(visitor ISQLVisitor, x *sequence.SQLOrderSequenceOption) {
	if visitor.VisitOrderSequenceOption(visitor, x) {
	}
}
func AcceptNoOrderSequenceOption(visitor ISQLVisitor, x *sequence.SQLNoOrderSequenceOption) {
	if visitor.VisitNoOrderSequenceOption(visitor, x) {
	}
}
func AcceptKeepSequenceOption(visitor ISQLVisitor, x *sequence.SQLKeepSequenceOption) {
	if visitor.VisitKeepSequenceOption(visitor, x) {
	}
}
func AcceptNoKeepSequenceOption(visitor ISQLVisitor, x *sequence.SQLNoKeepSequenceOption) {
	if visitor.VisitNoKeepSequenceOption(visitor, x) {
	}
}
func AcceptScaleSequenceOption(visitor ISQLVisitor, x *sequence.SQLScaleSequenceOption) {
	if visitor.VisitScaleSequenceOption(visitor, x) {
	}
}
func AcceptNoScaleSequenceOption(visitor ISQLVisitor, x *sequence.SQLNoScaleSequenceOption) {
	if visitor.VisitNoScaleSequenceOption(visitor, x) {
	}
}
func AcceptSessionSequenceOption(visitor ISQLVisitor, x *sequence.SQLSessionSequenceOption) {
	if visitor.VisitSessionSequenceOption(visitor, x) {
	}
}
func AcceptGlobalSequenceOption(visitor ISQLVisitor, x *sequence.SQLGlobalSequenceOption) {
	if visitor.VisitGlobalSequenceOption(visitor, x) {
	}
}
