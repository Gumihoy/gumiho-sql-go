package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/translation/config"
	"github.com/Gumihoy/gumiho-sql-go/translation/result"
)

type ISQLASTTransformVisitor interface {
	visitor.ISQLVisitor
	Config() *config.SQLTransformConfig
	Changes() []*result.SQLTransformChange
	Warnnings() []*result.SQLTransformWarnning
	Errors() []*result.SQLTransformError
}

type SQLASTTransformVisitor struct {
	*visitor.SQLVisitorAdapter
	config    *config.SQLTransformConfig
	changes   []*result.SQLTransformChange
	warnnings []*result.SQLTransformWarnning
	errors    []*result.SQLTransformError
}

func NewSQLASTTransformVisitor(config *config.SQLTransformConfig) *SQLASTTransformVisitor {
	x := new(SQLASTTransformVisitor)
	x.SQLVisitorAdapter = visitor.NewVisitorAdapter()
	x.config = config
	return x
}
func (x *SQLASTTransformVisitor) Config() *config.SQLTransformConfig {
	return x.config
}
func (x *SQLASTTransformVisitor) Changes() []*result.SQLTransformChange {
	return x.changes
}
func (x *SQLASTTransformVisitor) Warnnings() []*result.SQLTransformWarnning {
	return x.warnnings
}
func (x *SQLASTTransformVisitor) Errors() []*result.SQLTransformError {
	return x.errors
}
