package table

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * ALTER TABLE <table name> <alter table action>
 *
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#alter%20table%20statement
 *
 * ALTER TABLE tbl_name
    [alter_option [, alter_option] ...]
    [partition_options]
 * https://dev.mysql.com/doc/refman/8.0/en/alter-table.html
 *
 *
 * ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
    action [, ... ]
ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
    RENAME [ COLUMN ] column_name TO new_column_name
ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
    RENAME CONSTRAINT constraint_name TO new_constraint_name
ALTER TABLE [ IF EXISTS ] name
    RENAME TO new_name
ALTER TABLE [ IF EXISTS ] name
    SET SCHEMA new_schema
ALTER TABLE ALL IN TABLESPACE name [ OWNED BY role_name [, ... ] ]
    SET TABLESPACE new_tablespace [ NOWAIT ]
ALTER TABLE [ IF EXISTS ] name
    ATTACH PARTITION partition_name { FOR VALUES partition_bound_spec | DEFAULT }
ALTER TABLE [ IF EXISTS ] name
    DETACH PARTITION partition_name
 * https://www.postgresql.org/docs/devel/sql-altertable.html
 *
 * ALTER TABLE [ schema. ] table
  [ memoptimize_read_clause ]
  [ alter_table_properties
  | column_clauses
  | constraint_clauses
  | alter_table_partitioning
  | alter_external_table
  | move_table_clause
  | modify_to_partitioned
  | modify_opaque_type
  | blockchain_table_clauses
  ]
  [ enable_disable_clause
  | { ENABLE | DISABLE }
    { TABLE LOCK | ALL TRIGGERS | CONTAINER_MAP | CONTAINERS_DEFAULT }
  ] ...
 * https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-TABLE.html#GUID-552E7373-BF93-477D-9DA3-B2C9386F2877
 */
type SQLAlterTableStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	Only     bool
	name     expr.ISQLName
	actions  []table.ISQLAlterTableAction
}

func NewAlterTableStatement(dbType db.Type) *SQLAlterTableStatement {
	x := new(SQLAlterTableStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLAlterTableStatement) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	panic("implement me")
}

func (x *SQLAlterTableStatement) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLAlterTableStatement) ObjectType() db.SQLObjectType {
	return db.TABLE
}

func (x *SQLAlterTableStatement) Name() expr.ISQLName {
	return x.name
}

func (x *SQLAlterTableStatement) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLAlterTableStatement) Actions() []table.ISQLAlterTableAction {
	return x.actions
}
func (x *SQLAlterTableStatement) Action(i int) table.ISQLAlterTableAction {
	return x.actions[i]
}
func (x *SQLAlterTableStatement) AddAction(action table.ISQLAlterTableAction) {
	if action == nil {
		return
	}
	action.SetParent(x)
	x.actions = append(x.actions, action)
}
