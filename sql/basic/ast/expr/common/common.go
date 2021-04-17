package common

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
)

/**
 * (select)
 */
type SQLSubQueryExpr struct {
	*expr.AbstractSQLExpr
	Paren bool
	query select_.ISQLSelectQuery
}

func NewSubQueryExpr() *SQLSubQueryExpr {
	x := new(SQLSubQueryExpr)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x;
}

func (x *SQLSubQueryExpr) Query() select_.ISQLSelectQuery {
	return x.query
}
func (x *SQLSubQueryExpr) SetQuery(query select_.ISQLSelectQuery) {
	if query == nil {
		return
	}
	query.SetParent(x)
	x.query = query
}

/**
 * COLLATE <collation name>
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#collate%20clause
 */
type SQLCollateClause struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}

func NewCollateClause() *SQLCollateClause {
	x := new(SQLCollateClause)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x;
}

func (x *SQLCollateClause) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCollateClause) SetName(name expr.ISQLName) {
	if name == nil {
		return
	}
	name.SetParent(x)
	x.name = name
}

// Oracle
// 13 PL/SQL Language Elements

/**
 * 13.1 ACCESSIBLE BY Clause
 * ACCESSIBLE BY ( accessor [, accessor ]... )
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/ACCESSIBLE-BY-clause.html#GUID-9720619C-9862-4123-96E7-3E85F240FF36
 */
type SQLAccessibleByClause struct {
	*expr.AbstractSQLExpr
	accessors []*SQLAccessibleByClauseAccessor
}

type unitKind string

const (
	FUNCTION  unitKind = "FUNCTION"
	PROCEDURE          = "PROCEDURE"
	PACKAGE            = "PACKAGE"
	TRIGGER            = "TRIGGER"
	TYPE               = "TYPE"
)

/**
 *  [ unit_kind ] [schema.]unit_name
 */
type SQLAccessibleByClauseAccessor struct {
	*expr.AbstractSQLExpr
	UnitKind unitKind
	name     expr.ISQLName
}


/**
 * 13.2 AGGREGATE Clause
 * AGGREGATE USING [ schema. ] implementation_type
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/AGGREGATE-clause.html#GUID-2ED21240-E45A-4982-B674-CF0E1BE0985B
 */
type SQLAggregateClause struct {
	*expr.AbstractSQLExpr
	name expr.ISQLName
}


/**
 * 13.4 AUTONOMOUS_TRANSACTION Pragma
 * AGGREGATE USING [ schema. ] implementation_type
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/AGGREGATE-clause.html#GUID-2ED21240-E45A-4982-B674-CF0E1BE0985B
 */
type SQLAutonomousTransactionPragma struct {
	*expr.AbstractSQLExpr
}


/**
 * LANGUAGE JAVA NAME string
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/call-specification.html#GUID-C5F117AE-E9A2-499B-BA6A-35D072575BAD
 */
type SQLJavaDeclaration struct {
	*expr.AbstractSQLExpr
	name expr.ISQLExpr
}

/**
 * {LANGUAGE C | EXTERNAL } { [ NAME name ] LIBRARY lib_name | LIBRARY lib_name [ NAME name ] }
   [ AGENT IN ( argument[, argument ]... ) ]
   [ WITH CONTEXT ]
   [ PARAMETERS ( external_parameter[, external_parameter ]... ) ]
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/call-specification.html#GUID-C5F117AE-E9A2-499B-BA6A-35D072575BAD
 */
type ISQLCDeclaration interface {
	expr.ISQLExpr
}
type abstractSQLCDeclaration struct {
	*expr.AbstractSQLExpr
}

func NewAbstractSQLCDeclaration() *abstractSQLCDeclaration {
	x := new(abstractSQLCDeclaration)
	return x
}

type SQLCDeclaration struct {
	*expr.AbstractSQLExpr
}

func New() *SQLCDeclaration {
	x := new(SQLCDeclaration)

	return x
}
