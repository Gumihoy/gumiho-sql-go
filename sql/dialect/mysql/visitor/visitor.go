package visitor

import (
	exprTable "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	statementUser "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/user"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	"strings"
)

type IMySQLVisitor interface {
	visitor.ISQLVisitor
}

type MySQLVisitorAdapter struct {
	*visitor.SQLVisitorAdapter
}

func NewVisitorAdapter() *MySQLVisitorAdapter {
	return NewVisitorAdapterWithVisitorAdapter(visitor.NewVisitorAdapter())
}

func NewVisitorAdapterWithVisitorAdapter(adapter *visitor.SQLVisitorAdapter) *MySQLVisitorAdapter {
	x := new(MySQLVisitorAdapter)
	x.SQLVisitorAdapter = adapter
	return x
}

type MySQLOutputVisitor struct {
	*MySQLVisitorAdapter
	*visitor.SQLOutputVisitor
}

func NewOutputVisitor(builder *strings.Builder, config *config.SQLOutputConfig) *MySQLOutputVisitor {
	x := new(MySQLOutputVisitor)
	x.MySQLVisitorAdapter = NewVisitorAdapter()
	x.SQLOutputVisitor = visitor.NewOutputVisitor(builder, config)
	return x
}

// ---------------------------- Table Start ------------------------------------
func (v *MySQLOutputVisitor) VisitAddPartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLAddPartitionAlterTableAction) bool {
	v.WriteKeyword(visitor.ADD)
	v.WriteSpaceAfterKeyword(visitor.PARTITION)
	visitor.WriteSpaceAndIndentLnAfterAccept(child, true, x.Partition())
	return false
}
func (v *MySQLOutputVisitor) VisitDropPartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLDropPartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitDiscardPartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLDiscardPartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitImportPartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLImportPartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitTruncatePartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLTruncatePartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitCoalescePartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLCoalescePartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitReorganizePartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLReorganizePartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitExchangePartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLExchangePartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitAnalyzePartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLAnalyzePartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitCheckPartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLCheckPartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitOptimizePartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLOptimizePartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitRebuildPartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLRebuildPartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitRepairPartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLRepairPartitionAlterTableAction) bool {
	return false
}
func (v *MySQLOutputVisitor) VisitRemovePartitionAlterTableAction(child visitor.ISQLVisitor, x *exprTable.SQLRemovePartitionAlterTableAction) bool {
	return false
}

// ---------------------------- Table End ------------------------------------

func (v *MySQLOutputVisitor) VisitCreateUserStatement(child visitor.ISQLVisitor, x *statementUser.SQLCreateUserStatement) bool {
	v.WriteKeyword(visitor.CREATE)
	v.WriteSpaceAfterKeyword(visitor.USER)

	for i := 0; i < len(x.Names()); i++ {
		if i != 0 {
			v.WriteKeyword(visitor.SYMB_COMMA)
		}
		visitor.WriteSpaceAfterAccept(child, x.Name(i))
	}

	if len(x.Roles()) > 0 {
		v.WriteSpaceAfterKeyword(visitor.DEFAULT)
		v.WriteSpaceAfterKeyword(visitor.ROLE)
		for i := 0; i < len(x.Roles()); i++ {
			if i != 0 {
				v.WriteKeyword(visitor.SYMB_COMMA)
			}
			visitor.WriteSpaceAfterAccept(child, x.Role(i))
		}
	}

	return false
}
