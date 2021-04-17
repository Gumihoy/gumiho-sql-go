package show

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * 	SHOW {BINARY | MASTER} LOGS
	SHOW BINLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]
	SHOW CHARACTER SET [like_or_where]
	SHOW COLLATION [like_or_where]
	SHOW [FULL] COLUMNS FROM tbl_name [FROM db_name] [like_or_where]
	SHOW CREATE DATABASE db_name
	SHOW CREATE EVENT event_name
	SHOW CREATE FUNCTION func_name
	SHOW CREATE PROCEDURE proc_name
	SHOW CREATE TABLE tbl_name
	SHOW CREATE TRIGGER trigger_name
	SHOW CREATE VIEW view_name
	SHOW DATABASES [like_or_where]
	SHOW ENGINE engine_name {STATUS | MUTEX}
	SHOW [STORAGE] ENGINES
	SHOW ERRORS [LIMIT [offset,] row_count]
	SHOW EVENTS
	SHOW FUNCTION CODE func_name
	SHOW FUNCTION STATUS [like_or_where]
	SHOW GRANTS FOR user
	SHOW INDEX FROM tbl_name [FROM db_name]
	SHOW MASTER STATUS
	SHOW OPEN TABLES [FROM db_name] [like_or_where]
	SHOW PLUGINS
	SHOW PROCEDURE CODE proc_name
	SHOW PROCEDURE STATUS [like_or_where]
	SHOW PRIVILEGES
	SHOW [FULL] PROCESSLIST
	SHOW PROFILE [types] [FOR QUERY n] [OFFSET n] [LIMIT n]
	SHOW PROFILES
	SHOW RELAYLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]
	SHOW {REPLICAS | SLAVE HOSTS}
	SHOW {REPLICA | SLAVE} STATUS [FOR CHANNEL channel]
	SHOW [GLOBAL | SESSION] STATUS [like_or_where]
	SHOW TABLE STATUS [FROM db_name] [like_or_where]
	SHOW [FULL] TABLES [FROM db_name] [like_or_where]
	SHOW TRIGGERS [FROM db_name] [like_or_where]
	SHOW [GLOBAL | SESSION] VARIABLES [like_or_where]
	SHOW WARNINGS [LIMIT [offset,] row_count]

 * https://dev.mysql.com/doc/refman/8.0/en/show.html

 *
 */

/**
 * SHOW BINARY LOGS
 */
type SQLShowBinaryLogsStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowBinaryLogsStatement(dbType db.Type) *SQLShowBinaryLogsStatement {
	x := new(SQLShowBinaryLogsStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

/**
 * SHOW MASTER LOGS
 */
type SQLShowMasterLogsStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowMasterLogsStatement(dbType db.Type) *SQLShowMasterLogsStatement {
	x := new(SQLShowMasterLogsStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

/**
 * SHOW BINLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]
 */
type SQLShowBinlogEventsStatement struct {
	*statement.AbstractSQLStatement
	in expr.ISQLExpr

	pos         expr.ISQLExpr
	limitClause select_.ISQLLimitClause
}

/**
 *
 * https://dev.mysql.com/doc/refman/8.0/en/show.html
 */
type SQLShowCreateDatabaseStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowCreateDatabaseStatement(dbType db.Type) *SQLShowCreateDatabaseStatement {
	x := new(SQLShowCreateDatabaseStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowCreateDatabaseStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLShowCreateDatabaseStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 * https://dev.mysql.com/doc/refman/8.0/en/show.html
 */
type SQLShowCreateEventStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowCreateEventStatement(dbType db.Type) *SQLShowCreateEventStatement {
	x := new(SQLShowCreateEventStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowCreateEventStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLShowCreateEventStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

type SQLShowCreateFunctionStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowCreateFunctionStatement(dbType db.Type) *SQLShowCreateFunctionStatement {
	x := new(SQLShowCreateFunctionStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowCreateFunctionStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLShowCreateFunctionStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

type SQLShowCreateProcedureStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowCreateProcedureStatement(dbType db.Type) *SQLShowCreateProcedureStatement {
	x := new(SQLShowCreateProcedureStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowCreateProcedureStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLShowCreateProcedureStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

type SQLShowCreateTableStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowCreateTableStatement(dbType db.Type) *SQLShowCreateTableStatement {
	x := new(SQLShowCreateTableStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowCreateTableStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLShowCreateTableStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

type SQLShowCreateTriggerStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowCreateTriggerStatement(dbType db.Type) *SQLShowCreateTriggerStatement {
	x := new(SQLShowCreateTriggerStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowCreateTriggerStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLShowCreateTriggerStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

type SQLShowCreateViewStatement struct {
	*statement.AbstractSQLStatement
	name expr.ISQLName
}

func NewShowCreateViewStatement(dbType db.Type) *SQLShowCreateViewStatement {
	x := new(SQLShowCreateViewStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowCreateViewStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLShowCreateViewStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

/**
 *
 * https://dev.mysql.com/doc/refman/8.0/en/show.html
 */
type SQLShowDatabaseStatement struct {
	*statement.AbstractSQLStatement
	expr expr.ISQLExpr
}

func NewShowDatabaseStatement(dbType db.Type) *SQLShowDatabaseStatement {
	x := new(SQLShowDatabaseStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLShowDatabaseStatement) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLShowDatabaseStatement) SetExpr(expr expr.ISQLExpr) {
	if expr == nil {
		return
	}
	expr.SetParent(x)
	x.expr = expr
}

/**
 * LIKE pattern
 */
type SQLLikeExpr struct {
	*expr.AbstractSQLExpr
	pattern expr.ISQLExpr
}

func (x *SQLLikeExpr) Pattern() expr.ISQLExpr {
	return x.pattern
}
func (x *SQLLikeExpr) SetPattern(pattern expr.ISQLExpr) {
	if pattern == nil {
		return
	}
	pattern.SetParent(x)
	x.pattern = pattern
}

/**
 * SHOW EVENTS
 */
type SQLShowEventsStatement struct {
	*statement.AbstractSQLStatement
}

func NewShowEventsStatement(dbType db.Type) *SQLShowEventsStatement {
	x := new(SQLShowEventsStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

/**
 * SHOW TABLES
 */
type SQLShowTablesStatement struct {
	*statement.AbstractSQLStatement
}

func NewShowTablesStatement(dbType db.Type) *SQLShowTablesStatement {
	x := new(SQLShowTablesStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
