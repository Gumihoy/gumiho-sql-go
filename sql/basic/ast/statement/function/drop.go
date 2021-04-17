package function

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * DROP FUNCTION [IF EXISTS] sp_name
 * https://dev.mysql.com/doc/refman/8.0/en/drop-procedure.html;

 * DROP FUNCTION [ schema. ] function_name ;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/img_text/drop_function.html
 */
type SQLDropFunctionStatement struct {
	*statement.AbstractSQLStatement
	IfExists bool
	name     expr.ISQLName
}

func NewDropFunctionStatement(dbType db.Type) *SQLDropFunctionStatement {
	x := new(SQLDropFunctionStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}
func (x *SQLDropFunctionStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLDropFunctionStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
