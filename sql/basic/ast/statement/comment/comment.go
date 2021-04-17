package comment

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 *  COMMENT ON AUDIT POLICY policy IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnAuditPolicyStatement struct {
	*statement.AbstractSQLStatement
	name  expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnAuditPolicyStatement(dbType db.Type) *SQLCommentOnAuditPolicyStatement {
	x := new(SQLCommentOnAuditPolicyStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnAuditPolicyStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnAuditPolicyStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnAuditPolicyStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnAuditPolicyStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}

/**
 *  COMMENT ON COLUMN [ schema. ] { table. | view. | materialized_view. } column IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnColumnStatement struct {
	*statement.AbstractSQLStatement
	name  expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnColumnStatement(dbType db.Type) *SQLCommentOnColumnStatement {
	x := new(SQLCommentOnColumnStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnColumnStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnColumnStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnColumnStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnColumnStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}

/**
 *  COMMENT ON  EDITION edition_name IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnEditionStatement struct {
	*statement.AbstractSQLStatement
	name  expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnEditionStatement(dbType db.Type) *SQLCommentOnEditionStatement {
	x := new(SQLCommentOnEditionStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnEditionStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnEditionStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnEditionStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnEditionStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}

/**
 *  COMMENT ON INDEXTYPE [ schema. ] indextype IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnIndextypeStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnIndextypeStatement(dbType db.Type) *SQLCommentOnIndextypeStatement {
	x := new(SQLCommentOnIndextypeStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnIndextypeStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnIndextypeStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnIndextypeStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnIndextypeStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}

/**
 *  COMMENT ON MATERIALIZED VIEW materialized_view IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnMaterializedViewStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnMaterializedViewStatement(dbType db.Type) *SQLCommentOnMaterializedViewStatement {
	x := new(SQLCommentOnMaterializedViewStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnMaterializedViewStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnMaterializedViewStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnMaterializedViewStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnMaterializedViewStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}

/**
 *  COMMENT ON MINING MODEL [ schema. ] model IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnMiningModelStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnMiningModelStatement(dbType db.Type) *SQLCommentOnMiningModelStatement {
	x := new(SQLCommentOnMiningModelStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnMiningModelStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnMiningModelStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnMiningModelStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnMiningModelStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}

/**
 *  COMMENT ON OPERATOR [ schema. ] operator IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnOperatorStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnOperatorStatement(dbType db.Type) *SQLCommentOnOperatorStatement {
	x := new(SQLCommentOnOperatorStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnOperatorStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnOperatorStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnOperatorStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnOperatorStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}

/**
 *  COMMENT ON TABLE [ schema. ] { table | view } IS string
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMENT.html#GUID-65F447C4-6914-4823-9691-F15D52DB74D7
 */
type SQLCommentOnTableStatement struct {
	*statement.AbstractSQLStatement
	name    expr.ISQLExpr
	comment expr.ISQLExpr
}

func NewCommentOnTableStatement(dbType db.Type) *SQLCommentOnTableStatement {
	x := new(SQLCommentOnTableStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLCommentOnTableStatement) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLCommentOnTableStatement) SetName(name expr.ISQLExpr) bool {
	if name == nil {
		return false
	}
	name.SetParent(x)
	x.name = name
	return true
}
func (x *SQLCommentOnTableStatement) Comment() expr.ISQLExpr {
	return x.comment
}
func (x *SQLCommentOnTableStatement) SetComment(comment expr.ISQLExpr) bool {
	if comment == nil {
		return false
	}
	comment.SetParent(x)
	x.comment = comment
	return true
}