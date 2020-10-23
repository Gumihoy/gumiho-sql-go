package visitor

import (
	"gumihoy.com/sql/basic"
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/expr"
	"gumihoy.com/sql/basic/ast/expr/comment"
	"gumihoy.com/sql/basic/ast/expr/literal"
	"gumihoy.com/sql/basic/ast/expr/operator"
	"gumihoy.com/sql/basic/ast/expr/select"
	"gumihoy.com/sql/basic/ast/expr/variable"
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/config"
	"strings"
)

type ISQLVisitor interface {
	BeforeVisit(x ast.ISQLObject)
	AfterVisit(x ast.ISQLObject)

	// ---------------------------- Comment Start ------------------------------------
	VisitMultiLineCommentExpr(x *comment.SQLMultiLineCommentExpr) bool
	VisitMinusCommentExpr(x *comment.SQLMinusCommentExpr) bool
	VisitSharpCommentExpr(x *comment.SQLSharpCommentExpr) bool

	VisitMultiLineHintExpr(x *comment.SQLMultiLineHintExpr) bool
	VisitMinusHintExpr(x *comment.SQLMinusHintExpr) bool
	// ---------------------------- Comment End ------------------------------------

	// ---------------------------- Literal Start ------------------------------------
	// visitNQStringLiteral(x SQLNQStringLiteral) bool
	// visitNStringLiteral(x SQLNStringLiteral) bool
	// visitQStringLiteral(x SQLQStringLiteral) bool
	VisitStringLiteral(x *literal.SQLStringLiteral) bool
	VisitCharacterStringLiteral(x *literal.SQLCharacterStringLiteral) bool

	// visitBinaryDoubleLiteral(x SQLBinaryDoubleLiteral) bool
	// visitBinaryFloatLiteral(x SQLBinaryFloatLiteral) bool
	// visitDecimalLiteral(x SQLDecimalLiteral) bool
	// visitFloatingPointLiteral(x SQLFloatingPointLiteral) bool
	VisitIntegerLiteral(x *literal.SQLIntegerLiteral) bool
	VisitFloatingPointLiteral(x *literal.SQLFloatingPointLiteral) bool
	// visitBitValueLiteral(x literal.SQLBitValueLiteral) bool
	// visitHexaDecimalLiteral(x SQLHexaDecimalLiteral) bool
	//
	VisitDateLiteral(x *literal.SQLDateLiteral) bool
	VisitTimeLiteral(x *literal.SQLTimeLiteral) bool
	VisitTimestampLiteral(x *literal.SQLTimestampLiteral) bool

	VisitBooleanLiteral(x *literal.SQLBooleanLiteral) bool
	// visitIntervalLiteral(x SQLIntervalLiteral) bool
	// ---------------------------- Literal End ------------------------------------

	// ---------------------------- Identifier Start ------------------------------------
	VisitIdentifier(x *expr.SQLUnQuotedIdentifier) bool
	VisitDoubleQuotedIdentifier(x *expr.SQLDoubleQuotedIdentifier) bool
	VisitReverseQuotedIdentifier(x *expr.SQLReverseQuotedIdentifier) bool

	VisitName(x *expr.SQLName) bool

	// ---------------------------- Identifier End ------------------------------------

	// ---------------------------- Variable Start ------------------------------------
	VisitVariableExpr(x *variable.SQLVariableExpr) bool
	VisitBindVariableExpr(x *variable.SQLBindVariableExpr) bool

	// ---------------------------- Variable End ------------------------------------

	// ---------------------------- Operator Start ------------------------------------
	VisitUnaryOperatorExpr(x *operator.SQLUnaryOperatorExpr) bool
	VisitBinaryOperatorExpr(x *operator.SQLBinaryOperatorExpr) bool

	// ---------------------------- Operator End ------------------------------------

	// ---------------------------- Expr Start ------------------------------------
	VisitDBLinkExpr(x *expr.SQLDBLinkExpr) bool
	VisitAllColumnExpr(x *expr.SQLAllColumnExpr) bool

	// ---------------------------- Expr End ------------------------------------

	// ---------------------------- Select Start ------------------------------------
	VisitSelectQuery(x *select_.SQLSelectQuery) bool
	VisitParenSelectQuery(x *select_.SQLParenSelectQuery) bool
	VisitSelectUnionQuery(x *select_.SQLSelectUnionQuery) bool

	VisitSelectElement(x *select_.SQLSelectElement) bool
	VisitSelectTargetElement(x *select_.SQLSelectTargetElement) bool

	VisitFromClause(x *select_.SQLFromClause) bool
	VisitTableReference(x *select_.SQLTableReference) bool
	VisitSubQueryTableReference(x *select_.SQLSubQueryTableReference) bool
	VisitJoinTableReference(x *select_.SQLJoinTableReference) bool
	VisitJoinOnCondition(x *select_.SQLJoinOnCondition) bool
	VisitJoinUsingCondition(x *select_.SQLJoinUsingCondition) bool

	VisitWhereClause(x *select_.SQLWhereClause) bool
	VisitGroupByHavingClause(x *select_.SQLGroupByHavingClause) bool
	VisitHavingGroupByClause(x *select_.SQLHavingGroupByClause) bool
	VisitOrderByClause(x *select_.SQLOrderByClause) bool

	VisitLimitOffsetClause(x *select_.SQLLimitOffsetClause) bool
	VisitOffsetFetchClause(x *select_.SQLOffsetFetchClause) bool

	VisitForUpdateClause(x *select_.SQLForUpdateClause) bool

	// ---------------------------- Select End ------------------------------------

	// ---------------------------- Statement Start ------------------------------------

	VisitSelectStatement(x *statement.SQLSelectStatement) bool

	// ---------------------------- Statement End ------------------------------------
}

type SQLVisitorAdapter struct {
}

func NewVisitorAdapter() *SQLVisitorAdapter {
	return &SQLVisitorAdapter{}
}

func (v *SQLVisitorAdapter) BeforeVisit(x ast.ISQLObject) {

}

func (v *SQLVisitorAdapter) VisitMultiLineCommentExpr(x *comment.SQLMultiLineCommentExpr) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitMinusCommentExpr(x *comment.SQLMinusCommentExpr) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitSharpCommentExpr(x *comment.SQLSharpCommentExpr) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitMultiLineHintExpr(x *comment.SQLMultiLineHintExpr) bool {
	panic("implement me")
}

func (v *SQLVisitorAdapter) VisitMinusHintExpr(x *comment.SQLMinusHintExpr) bool {
	panic("implement me")
}

// ---------------------------- Literal Start ------------------------------------

func (v *SQLVisitorAdapter) VisitStringLiteral(x *literal.SQLStringLiteral) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitCharacterStringLiteral(x *literal.SQLCharacterStringLiteral) bool { return true }

func (v *SQLVisitorAdapter) VisitIntegerLiteral(x *literal.SQLIntegerLiteral) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitFloatingPointLiteral(x *literal.SQLFloatingPointLiteral) bool { return true }

func (v *SQLVisitorAdapter) VisitDateLiteral(x *literal.SQLDateLiteral) bool           { return true }
func (v *SQLVisitorAdapter) VisitTimeLiteral(x *literal.SQLTimeLiteral) bool           { return true }
func (v *SQLVisitorAdapter) VisitTimestampLiteral(x *literal.SQLTimestampLiteral) bool { return true }
func (v *SQLVisitorAdapter) VisitBooleanLiteral(x *literal.SQLBooleanLiteral) bool { return true }

// ---------------------------- Literal End ------------------------------------

// ---------------------------- Identifier Start ------------------------------------
func (v *SQLVisitorAdapter) VisitIdentifier(x *expr.SQLUnQuotedIdentifier) bool                   { return true }
func (v *SQLVisitorAdapter) VisitDoubleQuotedIdentifier(x *expr.SQLDoubleQuotedIdentifier) bool   { return true }
func (v *SQLVisitorAdapter) VisitReverseQuotedIdentifier(x *expr.SQLReverseQuotedIdentifier) bool { return true }

func (v *SQLVisitorAdapter) VisitName(x *expr.SQLName) bool {
	return true
}

// ---------------------------- Identifier End ------------------------------------

// ---------------------------- Variable Start ------------------------------------
func (v *SQLVisitorAdapter) VisitVariableExpr(x *variable.SQLVariableExpr) bool         { return true }
func (v *SQLVisitorAdapter) VisitBindVariableExpr(x *variable.SQLBindVariableExpr) bool { return true }

// ---------------------------- Variable End ------------------------------------

// ---------------------------- Operator Start ------------------------------------
func (v *SQLVisitorAdapter) VisitUnaryOperatorExpr(x *operator.SQLUnaryOperatorExpr) bool   { return true }
func (v *SQLVisitorAdapter) VisitBinaryOperatorExpr(x *operator.SQLBinaryOperatorExpr) bool { return true }

// ---------------------------- Operator End ------------------------------------

// ---------------------------- Expr Start ------------------------------------
func (v *SQLVisitorAdapter) VisitDBLinkExpr(x *expr.SQLDBLinkExpr) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitAllColumnExpr(x *expr.SQLAllColumnExpr) bool {
	return true
}

// ---------------------------- Expr End ------------------------------------

// ---------------------------- Select Start ------------------------------------
func (v *SQLVisitorAdapter) VisitSelectQuery(x *select_.SQLSelectQuery) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitParenSelectQuery(x *select_.SQLParenSelectQuery) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSelectUnionQuery(x *select_.SQLSelectUnionQuery) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitSelectElement(x *select_.SQLSelectElement) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSelectTargetElement(x *select_.SQLSelectTargetElement) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitFromClause(x *select_.SQLFromClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitTableReference(x *select_.SQLTableReference) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitSubQueryTableReference(x *select_.SQLSubQueryTableReference) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitJoinTableReference(x *select_.SQLJoinTableReference) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitJoinOnCondition(x *select_.SQLJoinOnCondition) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitJoinUsingCondition(x *select_.SQLJoinUsingCondition) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitWhereClause(x *select_.SQLWhereClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitGroupByHavingClause(x *select_.SQLGroupByHavingClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitHavingGroupByClause(x *select_.SQLHavingGroupByClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitOrderByClause(x *select_.SQLOrderByClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitLimitOffsetClause(x *select_.SQLLimitOffsetClause) bool {
	return true
}
func (v *SQLVisitorAdapter) VisitOffsetFetchClause(x *select_.SQLOffsetFetchClause) bool {
	return true
}

func (v *SQLVisitorAdapter) VisitForUpdateClause(x *select_.SQLForUpdateClause) bool {
	return true
}

// ---------------------------- Select End ------------------------------------

// ---------------------------- Statement Start ------------------------------------

func (v *SQLVisitorAdapter) VisitSelectStatement(x *statement.SQLSelectStatement) bool {
	return true
}

// ---------------------------- Statement End ------------------------------------

func (v *SQLVisitorAdapter) AfterVisit(x ast.ISQLObject) {

}

type ISQLOutputVisitor interface {
	ISQLVisitor

	IsLowerCase() bool

	Write(s string)
	WriteSpace()
	WriteKeyword(keyword *basic.Keyword)
	WriteLn()
}

type SQLOutputVisitor struct {
	*SQLVisitorAdapter

	Builder   *strings.Builder
	Indent    int
	Line, Pos int

	SelectElementLimit int
	Config    config.Output
}

func NewOutputVisitor(builder *strings.Builder, config config.Output) *SQLOutputVisitor {
	var x SQLOutputVisitor
	x.SQLVisitorAdapter = NewVisitorAdapter()
	x.Builder = builder
	x.Indent = 0
	x.Line, x.Pos = 1, 0

	x.SelectElementLimit = 80
	x.Config = config
	return &x
}

func (v *SQLOutputVisitor) IsLowerCase() bool {
	return v.Config.LowerCase
}

func (v *SQLOutputVisitor) IncrementIndent() {
	v.Indent++
}

func (v *SQLOutputVisitor) DecrementIndent() {
	v.Indent--
}

func (v *SQLOutputVisitor) IncrementLine() {
	v.Line++
	v.Pos = 0
}

func (v *SQLOutputVisitor) Write(s string) {
	v.Builder.WriteString(s)
	v.Pos += len(s)
}

func (v *SQLOutputVisitor) WriteSpace() {
	v.Builder.WriteString(" ")
	v.Pos++
}
func (v *SQLOutputVisitor) WriteSpaceAfterValue(s string) {
	v.WriteSpace()
	if v.IsLowerCase() {
		v.Builder.WriteString(strings.ToLower(s))
	} else {
		v.Builder.WriteString(strings.ToUpper(s))
	}
}
func (v *SQLOutputVisitor) WriteSpaceAfterAccept(x ast.ISQLObject) {
	if ast.IsNil(x) {
		return
	}
	v.Builder.WriteString(" ")
	Accept(v, x)
}

func (v *SQLOutputVisitor) WriteKeyword(keyword *basic.Keyword) {
	if v.IsLowerCase() {
		v.Builder.WriteString(keyword.Lower)
	} else {
		v.Builder.WriteString(keyword.Upper)
	}
}

func (v *SQLOutputVisitor) WriteLn() {
	v.Builder.WriteByte('\n')
	v.IncrementLine()
	v.WriteIndent()
}
func (v *SQLOutputVisitor) WriteLnAfterKeyword(keyword *basic.Keyword) {
	if keyword == nil {
		return
	}
	v.WriteLn()
	v.WriteKeyword(keyword)
}

func (v *SQLOutputVisitor) WriteLnAfterAccept(x ast.ISQLObject) {
	if ast.IsNil(x) {
		return
	}
	v.WriteLn()
	Accept(v, x)
}

func (v *SQLOutputVisitor) WriteIndent() {
	for i := 0; i < v.Indent; i++ {
		v.Builder.WriteByte('\t')
	}
}

// ---------------------------- Literal Start ------------------------------------

func (v *SQLOutputVisitor) VisitStringLiteral(x *literal.SQLStringLiteral) bool {
	v.Write("'")
	v.Write(x.Value())
	v.Write("'")
	return false
}
func (v *SQLOutputVisitor) VisitCharacterStringLiteral(x *literal.SQLCharacterStringLiteral) bool {

	v.Write(x.Charset())
	v.Write("'")
	v.Write(x.Value())
	v.Write("'")

	return false
}
func (v *SQLOutputVisitor) VisitIntegerLiteral(x *literal.SQLIntegerLiteral) bool {
	v.Write(x.StringValue())
	return false
}
func (v *SQLOutputVisitor) VisitFloatingPointLiteral(x *literal.SQLFloatingPointLiteral) bool {
	v.Write(x.StringValue())
	return false
}
func (v *SQLOutputVisitor) VisitDateLiteral(x *literal.SQLDateLiteral) bool {
	v.Write("DATE")
	v.WriteSpaceAfterAccept(x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitTimeLiteral(x *literal.SQLTimeLiteral) bool {
	v.Write("TIME")
	v.WriteSpaceAfterAccept(x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitTimestampLiteral(x *literal.SQLTimestampLiteral) bool {
	v.Write("TIMESTAMP")
	v.WriteSpaceAfterAccept(x.Value())
	return false
}
func (v *SQLOutputVisitor) VisitBooleanLiteral(x *literal.SQLBooleanLiteral) bool {
	if x.Value() {
		v.Write("true")
	} else {
		v.Write("false")
	}
	return false
}


// ---------------------------- Literal End ------------------------------------

// ---------------------------- Identifier Start ------------------------------------
func (v *SQLOutputVisitor) VisitIdentifier(x *expr.SQLUnQuotedIdentifier) bool {
	v.Write(x.StringName())
	return false
}
func (v *SQLOutputVisitor) VisitDoubleQuotedIdentifier(x *expr.SQLDoubleQuotedIdentifier) bool {
	v.Write("\"")
	v.Write(x.StringName())
	v.Write("\"")
	return false
}
func (v *SQLOutputVisitor) VisitReverseQuotedIdentifier(x *expr.SQLReverseQuotedIdentifier) bool {
	v.Write("`")
	v.Write(x.StringName())
	v.Write("`")
	return false
}

func (v *SQLOutputVisitor) VisitName(x *expr.SQLName) bool {
	return true
}

// ---------------------------- Identifier End ------------------------------------

// ---------------------------- Variable Start ------------------------------------
func (v *SQLOutputVisitor) VisitVariableExpr(x *variable.SQLVariableExpr) bool {
	v.Write("?")
	return false
}
func (v *SQLOutputVisitor) VisitBindVariableExpr(x *variable.SQLBindVariableExpr) bool {
	v.Write(":")
	Accept(v, x.Name())
	return false
}

// ---------------------------- Variable End ------------------------------------

// ---------------------------- Operator Start ------------------------------------
func (v *SQLOutputVisitor) VisitUnaryOperatorExpr(x *operator.SQLUnaryOperatorExpr) bool {
	switch x.Operator {
	default:
		v.Write(string(x.Operator))
		v.WriteSpaceAfterAccept(x.Operand())
	}
	return false
}
func (v *SQLOutputVisitor) VisitBinaryOperatorExpr(x *operator.SQLBinaryOperatorExpr) bool {
	if x.Paren {
		v.Write("(")
	}
	Accept(v, x.Left())
	v.WriteSpaceAfterValue(string(x.Operator))
	v.WriteSpaceAfterAccept(x.Right())
	if x.Paren {
		v.Write(")")
	}
	return false
}

// ---------------------------- Operator End ------------------------------------

// ---------------------------- Expr Start ------------------------------------
func (v *SQLOutputVisitor) VisitDBLinkExpr(x *expr.SQLDBLinkExpr) bool {
	return true
}
func (v *SQLOutputVisitor) VisitAllColumnExpr(x *expr.SQLAllColumnExpr) bool {
	v.Write("*")
	return false
}

// ---------------------------- Expr End ------------------------------------

// ---------------------------- Select Start ------------------------------------
func (v *SQLOutputVisitor) VisitSelectQuery(x *select_.SQLSelectQuery) bool {
	v.WriteKeyword(basic.SELECT)

	v.WriteSelectElements(x.SelectElements())

	v.WriteLnAfterAccept(x.FromClause())

	v.WriteLnAfterAccept(x.WhereClause())

	v.WriteLnAfterAccept(x.GroupByClause())

	return false
}

func (v *SQLOutputVisitor) WriteSelectElements(x []*select_.SQLSelectElement) {
	start := v.Builder.Len()
	end := v.Builder.Len()
	v.IncrementIndent()
	for i := 0; i < len(x); i++ {
		if i == 0 {
			v.WriteSpace()
		}
		if i != 0 {
			v.Write(",")
			if  end - start > v.SelectElementLimit {
				v.WriteLn()
				start = v.Builder.Len()
			} else {
				v.WriteSpace()
			}
		}

		child := x[i]
		Accept(v, child)

		end = v.Builder.Len()
	}
	v.DecrementIndent()
}

func (v *SQLOutputVisitor) VisitParenSelectQuery(x *select_.SQLParenSelectQuery) bool {
	v.Write("(")

	v.IncrementIndent()
	v.WriteLn()

	Accept(v, x.SubQuery())

	v.DecrementIndent()
	v.WriteLn()
	v.Write(")")

	v.WriteLnAfterAccept(x.OrderByClause())
	v.WriteLnAfterAccept(x.LimitClause())

	return false
}
func (v *SQLOutputVisitor) VisitSelectUnionQuery(x *select_.SQLSelectUnionQuery) bool {

	Accept(v, x.Left())

	v.Write(string(x.Operator))

	Accept(v, x.Right())

	v.WriteLnAfterAccept(x.OrderByClause())
	v.WriteLnAfterAccept(x.LimitClause())

	return false
}

func (v *SQLOutputVisitor) VisitSelectElement(x *select_.SQLSelectElement) bool {
	Accept(v, x.Expr())

	if x.As {
		v.WriteSpaceAfterValue("AS")
	}

	v.WriteSpaceAfterAccept(x.Alias())

	return false
}
func (v *SQLOutputVisitor) VisitSelectTargetElement(x *select_.SQLSelectTargetElement) bool {

	return false
}

func (v *SQLOutputVisitor) VisitFromClause(x *select_.SQLFromClause) bool {
	v.WriteKeyword(basic.FROM)
	v.WriteSpaceAfterAccept(x.TableReference())
	return false
}
func (v *SQLOutputVisitor) VisitTableReference(x *select_.SQLTableReference) bool {
	Accept(v, x.Name())
	return false
}
func (v *SQLOutputVisitor) VisitSubQueryTableReference(x *select_.SQLSubQueryTableReference) bool {
	return true
}
func (v *SQLOutputVisitor) VisitJoinTableReference(x *select_.SQLJoinTableReference) bool {
	return true
}
func (v *SQLOutputVisitor) VisitJoinOnCondition(x *select_.SQLJoinOnCondition) bool {
	return true
}
func (v *SQLOutputVisitor) VisitJoinUsingCondition(x *select_.SQLJoinUsingCondition) bool {
	return true
}

func (v *SQLOutputVisitor) VisitWhereClause(x *select_.SQLWhereClause) bool {
	v.WriteKeyword(basic.WHERE)
	v.WriteSpaceAfterAccept(x.Condition())
	return false
}
func (v *SQLOutputVisitor) VisitGroupByHavingClause(x *select_.SQLGroupByHavingClause) bool {
	return true
}
func (v *SQLOutputVisitor) VisitHavingGroupByClause(x *select_.SQLHavingGroupByClause) bool {
	return true
}
func (v *SQLOutputVisitor) VisitOrderByClause(x *select_.SQLOrderByClause) bool {
	return true
}

func (v *SQLOutputVisitor) VisitLimitOffsetClause(x *select_.SQLLimitOffsetClause) bool {
	return true
}
func (v *SQLOutputVisitor) VisitOffsetFetchClause(x *select_.SQLOffsetFetchClause) bool {
	return true
}

func (v *SQLOutputVisitor) VisitForUpdateClause(x *select_.SQLForUpdateClause) bool {
	return true
}

// ---------------------------- Select End ------------------------------------
