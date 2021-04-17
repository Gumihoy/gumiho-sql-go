package database

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE {DATABASE | SCHEMA} [IF NOT EXISTS] db_name
    [create_option] ...

create_option: [DEFAULT] {
    CHARACTER SET [=] charset_name
  | COLLATE [=] collation_name
  | ENCRYPTION [=] {'Y' | 'N'}
}
 * https://dev.mysql.com/doc/refman/8.0/en/create-database.html
 *
 * CREATE DATABASE [ database ]
  { USER SYS IDENTIFIED BY password
  | USER SYSTEM IDENTIFIED BY password
  | CONTROLFILE REUSE
  | MAXDATAFILES integer
  | MAXINSTANCES integer
  | CHARACTER SET charset
  | NATIONAL CHARACTER SET charset
  | SET DEFAULT
      { BIGFILE | SMALLFILE } TABLESPACE
  | database_logging_clauses
  | tablespace_clauses
  | set_time_zone_clause
  | [ BIGFILE | SMALLFILE ] USER_DATA TABLESPACE tablespace_name
      DATAFILE datafile_tempfile_spec [, datafile_tempfile_spec ]...
  | enable_pluggable_database
  }... ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-DATABASE.html#GUID-ECE717DF-F116-4151-927C-2E51BB9DD39C
 */
type SQLCreateDatabaseStatement struct {
	*statement.AbstractSQLStatement

	IfNotExists bool
	name        expr.ISQLName
	options     []expr.ISQLExpr
}

func NewCreateDatabaseStatement(dbType db.Type) *SQLCreateDatabaseStatement {
	x := new(SQLCreateDatabaseStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	x.options = make([]expr.ISQLExpr, 0, 10)
	return x
}

func (x *SQLCreateDatabaseStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLCreateDatabaseStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

func (x *SQLCreateDatabaseStatement) Options() []expr.ISQLExpr {
	return x.options
}

func (x *SQLCreateDatabaseStatement) AddOption(option expr.ISQLName) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
