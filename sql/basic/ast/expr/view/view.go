package view

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
)

/**
 * ADD out_of_line_constraint
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLAddTableConstraintAlterViewAction struct {
	*expr.AbstractSQLExpr
	tableConstraint table.ISQLTableConstraint
}

func NewAddTableConstraintAlterViewAction() *SQLAddTableConstraintAlterViewAction {
	x := new(SQLAddTableConstraintAlterViewAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLAddTableConstraintAlterViewAction) TableConstraint() table.ISQLTableConstraint {
	return x.tableConstraint
}

func (x *SQLAddTableConstraintAlterViewAction) SetTableConstraint(tableConstraint table.ISQLTableConstraint) {
	tableConstraint.SetParent(x)
	x.tableConstraint = tableConstraint
}


/**
 * MODIFY CONSTRAINT constraint
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLModifyConstraintAlterViewAction struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewModifyConstraintAlterViewAction() *SQLModifyConstraintAlterViewAction {
	x := new(SQLModifyConstraintAlterViewAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLModifyConstraintAlterViewAction) Name() expr.ISQLName {
	return x.name
}

func (x *SQLModifyConstraintAlterViewAction) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * DROP CONSTRAINT name
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLDropConstraintTableConstraintAlterViewAction struct {
	*expr.AbstractSQLExpr
	IfExists bool
	name     expr.ISQLName
}

func NewDropConstraintTableConstraintAlterViewAction() *SQLDropConstraintTableConstraintAlterViewAction {
	x := new(SQLDropConstraintTableConstraintAlterViewAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropConstraintTableConstraintAlterViewAction) Name() expr.ISQLName {
	return x.name
}

func (x *SQLDropConstraintTableConstraintAlterViewAction) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}

/**
 * DROP PRIMARY KEY
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLDropPrimaryKeyTableConstraintAlterViewAction struct {
	*expr.AbstractSQLExpr
}

func NewDropPrimaryKeyTableConstraintAlterViewAction() *SQLDropPrimaryKeyTableConstraintAlterViewAction {
	x := new(SQLDropPrimaryKeyTableConstraintAlterViewAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * drop_constraint_clause: DROP UNIQUE (column [, column ]...)
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/ALTER-VIEW.html#GUID-0DEDE960-B481-4B55-8027-EA9E4C863625
 */
type SQLDropUniqueTableConstraintAlterViewAction struct {
	*expr.AbstractSQLExpr
	names []expr.ISQLName
}

func NewDropUniqueAlterViewAction() *SQLDropUniqueTableConstraintAlterViewAction {
	x := new(SQLDropUniqueTableConstraintAlterViewAction)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLDropUniqueTableConstraintAlterViewAction) Names() []expr.ISQLName {
	return x.names
}
func (x *SQLDropUniqueTableConstraintAlterViewAction) Name(i int) expr.ISQLName {
	return x.names[i]
}
func (x *SQLDropUniqueTableConstraintAlterViewAction) AddName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.names = append(x.names, name)
}



type ISQLViewElement interface {
	expr.ISQLExpr
}

/**
 * name [DataType]
 *
 * { alias [ VISIBLE | INVISIBLE ] [ inline_constraint... ]
      | out_of_line_constraint}
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/CREATE-VIEW.html#GUID-61D2D2B4-DACC-4C7C-89EB-7E50D9594D30
 */
type SQLViewColumn struct {
	*expr.AbstractSQLExpr
	name     expr.ISQLName
	dataType datatype.ISQLDataType

	// option or column constraint
	options []expr.ISQLExpr
}

func NewViewColumn() *SQLViewColumn {
	x := new(SQLViewColumn)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLViewColumn) Name() expr.ISQLName {
	return x.name
}
func (x *SQLViewColumn) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLViewColumn) DataType() datatype.ISQLDataType {
	return x.dataType
}
func (x *SQLViewColumn) SetDataType(datatype datatype.ISQLDataType) {
	if datatype == nil {
		return
	}
	datatype.SetParent(x)
	x.dataType = datatype
}
func (x *SQLViewColumn) Options() []expr.ISQLExpr {
	return x.options
}
func (x *SQLViewColumn) Option(i int) expr.ISQLExpr {
	return x.options[i]
}
func (x *SQLViewColumn) AddOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.options = append(x.options, option)
}
