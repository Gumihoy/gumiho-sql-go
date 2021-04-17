package typebody

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

/**
 * CREATE [ OR REPLACE ] [ EDITIONABLE | NONEDITIONABLE ]
   TYPE BODY [ schema. ] type_name  { IS | AS }
   { subprog_decl_in_type
   | map_order_func_declaration
   }
     [, { subprog_decl_in_type
        | map_order_func_declaration
        }
     ]...
END;
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/CREATE-TYPE-BODY-statement.html#GUID-B468D6FB-75ED-436B-80E4-8460E4551AE0
 */
type SQLCreateTypeBodyStatement struct {
	*statement.AbstractSQLStatement
	OrReplace bool
	name      expr.ISQLName

}

func NewCreateTypeBodyStatement(dbType db.Type) *SQLCreateTypeBodyStatement {
	x := new(SQLCreateTypeBodyStatement)
	x.AbstractSQLStatement = statement.NewAbstractSQLStatementWithDBType(dbType)
	return x
}

func (x *SQLCreateTypeBodyStatement) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCreateTypeBodyStatement) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}
