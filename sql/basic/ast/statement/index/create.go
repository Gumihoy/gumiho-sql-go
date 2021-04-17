package index

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/index"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 *
 *
 * CREATE [UNIQUE | FULLTEXT | SPATIAL] INDEX index_name
    [index_type]
    ON tbl_name (key_part,...)
    [index_option]
    [algorithm_option | lock_option] ...
 * https://dev.mysql.com/doc/refman/8.0/en/create-index.html
 *
 * CREATE [ UNIQUE | BITMAP ] INDEX [ schema. ] index
  ON { cluster_index_clause
     | table_index_clause
     | bitmap_join_index_clause
     }
[ USABLE | UNUSABLE ]
[ { DEFERRED | IMMEDIATE } INVALIDATION ] ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-INDEX.html#GUID-1F89BBC0-825F-4215-AF71-7588E31D8BFE
 */
type SQLCreateIndexStatement struct {
	*statement.AbstractSQLStatement

	name expr.ISQLName

	Cluster bool
	onName  expr.ISQLName
	alias   expr.ISQLName
	columns []*index.SQLIndexColumn

	fromClause  select_.SQLFromClause
	whereClause *select_.SQLWhereClause

	options []expr.ISQLExpr
}

func NewCreateIndexStatement(dbType db.Type) *SQLCreateIndexStatement {
	x := new(SQLCreateIndexStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLCreateIndexStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateIndexStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

func (x *SQLCreateIndexStatement) OnName() expr.ISQLName {
	return x.onName
}
func (x *SQLCreateIndexStatement) SetOnName(onTable expr.ISQLName) {
	if onTable == nil {
		return
	}
	onTable.SetParent(x)
	x.onName = onTable
}
func (x *SQLCreateIndexStatement) Columns() []*index.SQLIndexColumn{
	return x.columns
}
func (x *SQLCreateIndexStatement) Column(i int) *index.SQLIndexColumn {
	return x.columns[i]
}
func (x *SQLCreateIndexStatement) AddColumn(column *index.SQLIndexColumn) {
	if column == nil {
		return
	}
	column.SetParent(x)
	x.columns = append(x.columns, column)
}


func (x *SQLCreateIndexStatement) Options() []expr.ISQLExpr {
	return x.options
}
func (x *SQLCreateIndexStatement) Option(i int) expr.ISQLExpr {
	return x.options[i]
}
func (x *SQLCreateIndexStatement) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
