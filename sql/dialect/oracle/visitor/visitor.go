package visitor

import (
	exprSequence "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/sequence"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	"strings"
)

type IOracleVisitor interface {
	visitor.ISQLVisitor
}

type OracleVisitorAdapter struct {
	*visitor.SQLVisitorAdapter
}

func NewVisitorAdapter() *OracleVisitorAdapter {
	x := new(OracleVisitorAdapter)
	x.SQLVisitorAdapter = visitor.NewVisitorAdapter()
	return x
}

type OracleOutputVisitor struct {
	*visitor.SQLOutputVisitor
	*OracleVisitorAdapter
}

func NewOutputVisitor(builder *strings.Builder, config *config.SQLOutputConfig) *OracleOutputVisitor {
	x := new(OracleOutputVisitor)
	x.OracleVisitorAdapter = NewVisitorAdapter()
	x.SQLOutputVisitor = visitor.NewOutputVisitor(builder, config)
	return x
}

// ---------------------------- Sequence Start ------------------------------------

func (v *OracleOutputVisitor) VisitNoMaxValueSequenceOption(child visitor.ISQLVisitor, x *exprSequence.SQLNoMaxValueSequenceOption) bool {
	v.WriteKeyword(visitor.NOMAXVALUE)
	return false
}

func (v *OracleOutputVisitor) VisitNoMinValueSequenceOption(child visitor.ISQLVisitor, x *exprSequence.SQLNoMinValueSequenceOption) bool {
	v.WriteKeyword(visitor.NOMINVALUE)
	return false
}

// ---------------------------- Sequence Start ------------------------------------

// ---------------------------- Statement Start ------------------------------------
func (v *OracleOutputVisitor) VisitExplainStatement(child visitor.ISQLVisitor, x *statement.SQLExplainStatement) bool {
	v.WriteKeyword(visitor.EXPLAIN)
	v.WriteSpaceAfterKeyword(visitor.PLAIN)

	v.IncrementIndent()
	if x.SetStatementIdValueExpr() != nil {
		v.WriteKeyword(visitor.SET)
		v.WriteSpaceAfterKeyword(visitor.STATEMENT_ID)
		v.WriteSpaceAfterKeyword(visitor.SYMB_EQUAL)
		visitor.WriteSpaceAfterAccept(child, x.SetStatementIdValueExpr())
	}

	if x.IntoTable() != nil {
		v.WriteLnAfterKeyword(visitor.INTO)
		visitor.WriteSpaceAfterAccept(child, x.IntoTable())
	}

	v.WriteLnAfterKeyword(visitor.FOR)
	visitor.WriteSpaceAfterAccept(child, x.Stmt())

	v.DecrementIndent()
	return false
}

// ---------------------------- Statement End ------------------------------------
