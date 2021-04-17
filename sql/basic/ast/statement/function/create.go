package function

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE <SQL-invoked function>
 * { <function specification> | <method specification designator> } <routine body>
 <function specification> 	::=
		FUNCTION <schema qualified routine name>
         <SQL parameter declaration list> <returns clause> <routine characteristics> [ <dispatch clause> ]
<method specification designator>    ::=
         SPECIFIC METHOD <specific method name>
     |     [ INSTANCE | STATIC | CONSTRUCTOR ] METHOD <method name> <SQL parameter declaration list>
         [ <returns clause> ] FOR <schema-resolved user-defined type name>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#schema%20function
 *
 *
 * CREATE
    [DEFINER = user]
    FUNCTION sp_name ([func_parameter[,...]])
    RETURNS type
    [characteristic ...] routine_body
 * https://dev.mysql.com/doc/refman/8.0/en/create-procedure.html
 *
 *
* CREATE [ OR REPLACE ]
[ EDITIONABLE | NONEDITIONABLE ]
FUNCTION [ schema. ] function_name
  [ ( parameter_declaration [, parameter_declaration]... ) ] RETURN datatype
[ sharing_clause ]
  [ { invoker_rights_clause
    | accessible_by_clause
    | default_collation_clause
    | deterministic_clause
    | parallel_enable_clause
    | result_cache_clause
    | aggragate_clause
    | pipelined_clause
       }...
  ]
{ IS | AS } { [ declare_section ] body
              | call_spec
            }
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/CREATE-FUNCTION-statement.html#GUID-B71BC5BD-B87C-4054-AAA5-213E856651F2
 */
type SQLCreateFunctionStatement struct {
	*statement.AbstractSQLStatement

	// MySQL
	definerExpr *expr.SQLAssignExpr

	// Oracle
	OrReplace bool

	name expr.ISQLName
	parameters []*statement.SQLParameterDeclaration
	returnType datatype.ISQLDataType

	// Oracle



}

func NewCreateFunctionStatement(dbType db.Type) *SQLCreateFunctionStatement {
	x := new(SQLCreateFunctionStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLCreateFunctionStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateFunctionStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLCreateFunctionStatement) Parameters() []*statement.SQLParameterDeclaration {
	return x.parameters
}
func (x *SQLCreateFunctionStatement) Parameter(i int) *statement.SQLParameterDeclaration {
	return x.parameters[i]
}
func (x *SQLCreateFunctionStatement) AddParameter(parameter *statement.SQLParameterDeclaration) {
	if parameter == nil {
		return
	}
	parameter.SetParent(x)
	x.parameters = append(x.parameters, parameter)
}
