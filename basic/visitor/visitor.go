package visitor

import (
	"gumihoy.com/sql/basic/ast/expr/comment"
	"gumihoy.com/sql/config"
	"strings"
)

type ISQLVisitor interface {

	// ---------------------------- Comment Start ------------------------------------
	visitMultiLineCommentExpr(x comment.SQLMultiLineCommentExpr) bool
	visitMinusCommentExpr(x comment.SQLMinusCommentExpr) bool
	visitSharpCommentExpr(x comment.SQLSharpCommentExpr) bool

	visitMultiLineHintExpr(x comment.SQLMultiLineHintExpr) bool
	visitMinusHintExpr(x comment.SQLMinusHintExpr) bool
	// ---------------------------- Comment End ------------------------------------

	// ---------------------------- Literal Start ------------------------------------
	// visitNQStringLiteral(x SQLNQStringLiteral) bool
	//visitNStringLiteral(x SQLNStringLiteral) bool
	//visitQStringLiteral(x SQLQStringLiteral) bool
	//visitStringLiteral(x SQLStringLiteral) bool
	//visitUStringLiteral(x SQLUStringLiteral) bool

	//visitBinaryDoubleLiteral(x SQLBinaryDoubleLiteral) bool
	//visitBinaryFloatLiteral(x SQLBinaryFloatLiteral) bool
	//visitDecimalLiteral(x SQLDecimalLiteral) bool
	//visitFloatingPointLiteral(x SQLFloatingPointLiteral) bool
	//visitIntegerLiteral(x SQLIntegerLiteral) bool
	//visitBitValueLiteral(x SQLBitValueLiteral) bool
	//visitHexaDecimalLiteral(x SQLHexaDecimalLiteral) bool
	//
	//visitDateLiteral(x SQLDateLiteral) bool
	//visitTimeLiteral(x SQLTimeLiteral) bool
	//visitTimestampLiteral(x SQLTimestampLiteral) bool
	//
	//visitIntervalLiteral(x SQLIntervalLiteral) bool
	// ---------------------------- Literal End ------------------------------------

	// ---------------------------- Identifier Start ------------------------------------
	// visit(x SQLUnquotedIdentifier) bool
	// visit(x SQLDoubleQuotedIdentifier) bool
	// visit(x SQLReverseQuotedIdentifier) bool
	//
	// visit(x SQLPropertyExpr) bool
	//
	// visit(x SQLAllColumnExpr) bool

	//visitDBLinkExpr(x SQLDBLinkExpr) bool
	// ---------------------------- Identifier End ------------------------------------

}

type SQLVisitorAdapter struct {
}

func (S SQLVisitorAdapter) visitMultiLineCommentExpr(x comment.SQLMultiLineCommentExpr) bool {
	panic("implement me")
}

func (S SQLVisitorAdapter) visitMinusCommentExpr(x comment.SQLMinusCommentExpr) bool {
	panic("implement me")
}

func (S SQLVisitorAdapter) visitSharpCommentExpr(x comment.SQLSharpCommentExpr) bool {
	panic("implement me")
}

func (S SQLVisitorAdapter) visitMultiLineHintExpr(x comment.SQLMultiLineHintExpr) bool {
	panic("implement me")
}

func (S SQLVisitorAdapter) visitMinusHintExpr(x comment.SQLMinusHintExpr) bool {
	panic("implement me")
}

type ISQLOutputVisitor interface {
	ISQLVisitor
	Write(s string)
	WriteLn()
}

type SQLOutputVisitor struct {
	*SQLVisitorAdapter

	Builder   strings.Builder
	Config    config.Output
	Line, Pos int
}

func (v *SQLOutputVisitor) Write(s string) {
	v.Builder.WriteString(s)
}

func (v *SQLOutputVisitor) WriteLn() {
	v.Builder.WriteString("\n")
	v.Line++
	v.Pos = 0
}

func NewOutputVisitor(builder strings.Builder, config config.Output) *SQLOutputVisitor {
	x := new(SQLOutputVisitor)
	x.Builder = builder
	x.Config = config
	return x
}
