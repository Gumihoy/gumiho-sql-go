package procedure

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE  { <function specification> | <method specification designator> } <routine body>
 <function specification>    ::=
         FUNCTION <schema qualified routine name>
         <SQL parameter declaration list> <returns clause> <routine characteristics> [ <dispatch clause> ]

<method specification designator>    ::=
         SPECIFIC METHOD <specific method name>
     |     [ INSTANCE | STATIC | CONSTRUCTOR ] METHOD <method name> <SQL parameter declaration list>
         [ <returns clause> ] FOR <schema-resolved user-defined type name>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#schema%20procedure
 *

 * CREATE
    [DEFINER = user]
    PROCEDURE sp_name ([proc_parameter[,...]])
    [characteristic ...] routine_body
 * https://dev.mysql.com/doc/refman/8.0/en/create-procedure.html

 *
 */
type SQLCreateProcedureStatement struct {
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
func NewCreateProcedureStatement(dbType db.Type) *SQLCreateProcedureStatement {
	x := new(SQLCreateProcedureStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLCreateProcedureStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateProcedureStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
func (x *SQLCreateProcedureStatement) Parameters() []*statement.SQLParameterDeclaration {
	return x.parameters
}
func (x *SQLCreateProcedureStatement) Parameter(i int) *statement.SQLParameterDeclaration {
	return x.parameters[i]
}
func (x *SQLCreateProcedureStatement) SetParameter(parameter *statement.SQLParameterDeclaration) {
	if parameter == nil {
		return
	}
	parameter.SetParent(x)
	x.parameters = append(x.parameters, parameter)
}



