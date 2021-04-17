package index

import "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"

/**
 *
 * {col_name [(length)] | (expr)} [ASC | DESC]
 * https://dev.mysql.com/doc/refman/8.0/en/create-index.html
 *
 * index_expr [ ASC | DESC ]
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-INDEX.html#GUID-1F89BBC0-825F-4215-AF71-7588E31D8BFE
 */
type SQLIndexColumn struct {
	*expr.AbstractSQLExpr
	expr expr.ISQLExpr
	len  expr.ISQLExpr
}

func NewViewColumn() *SQLIndexColumn {
	x := new(SQLIndexColumn)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIndexColumn) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLIndexColumn) SetExpr(expr expr.ISQLExpr) {
	if expr == nil {
		return
	}
	expr.SetParent(x)
	x.expr = expr
}
func (x *SQLIndexColumn) Len() expr.ISQLExpr {
	return x.len
}
func (x *SQLIndexColumn) SetLen(len expr.ISQLExpr) {
	if len == nil {
		return
	}
	len.SetParent(x)
	x.len = len
}
