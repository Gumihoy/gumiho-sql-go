package oracle

import (
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	visitor2 "github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/translation/config"
	"github.com/Gumihoy/gumiho-sql-go/translation/visitor"
)

type OracleASTTransformVisitor struct {
	*visitor.SQLASTTransformVisitor
}

func NewOracleASTTransformVisitor(config *config.SQLTransformConfig) *OracleASTTransformVisitor {
	x := new(OracleASTTransformVisitor)
	x.SQLASTTransformVisitor = visitor.NewSQLASTTransformVisitor(config)
	return x
}

func (v *OracleASTTransformVisitor) VisitDoubleQuotedIdentifier(child visitor2.ISQLVisitor, x *expr.SQLDoubleQuotedIdentifier) bool {
	if !IsKeepDoubleQuote(x.Name()) {
		target := expr.NewUnQuotedIdentifier(x.Name())
		sql.ReplaceInParent(x, target)
	}
	return false
}
