package condition

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/operator"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
)

type ISQLCondition interface {
	operator.ISQLOperator
}

type AbstractSQLCondition struct {
	*expr.AbstractSQLExpr
	paren bool
}

func NewAbstractSQLCondition() *AbstractSQLCondition {
	x := new(AbstractSQLCondition)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *AbstractSQLCondition) Paren() bool {
	return x.paren
}
func (x *AbstractSQLCondition) SetParen(paren bool) {
	x.paren = paren
}

/**
 * expr IS [NOT] {TRUE | FALSE | UNKNOWN | NULL}
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#boolean%20test
 *
 * expr IS [NOT] {TRUE | FALSE | UNKNOWN | NULL | NAN | INFINITE}
 * https://dev.mysql.com/doc/refman/8.0/en/expressions.html

 * expr IS [NOT] {NULL | NAN | INFINITE}
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Floating-Point-Conditions.html#GUID-D7707649-2C93-4553-BF78-F461F17A634E
 */
type SQLIsConditionValue string

const (
	TRUE     SQLIsConditionValue = "TRUE"
	FALSE                        = "FALSE"
	UNKNOWN                      = "UNKNOWN"
	NULL                         = "NULL"
	NAN                          = "NAN"
	INFINITE                     = "INFINITE"
	EMPTY                        = "EMPTY"
)

type SQLIsCondition struct {
	*AbstractSQLCondition
	expr  expr.ISQLExpr
	Not   bool
	Value SQLIsConditionValue
}

func NewIsCondition() *SQLIsCondition {
	x := new(SQLIsCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}
func (x *SQLIsCondition) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.expr {
		x.SetExpr(target)
		return true
	}
	return false
}
func (x *SQLIsCondition) Clone() ast.ISQLObject {
	panic("implement me")
}
func (x *SQLIsCondition) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLIsCondition) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

/**
 * nested_table IS [ NOT ] A SET
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Multiset-Conditions.html#GUID-E8164A15-715A-40A0-944D-26DF4C84DE3F
 */
type SQLIsASetCondition struct {
	*AbstractSQLCondition
	expr expr.ISQLExpr
	Not  bool
}

func NewIsASetCondition(expr expr.ISQLExpr) *SQLIsASetCondition {
	x := new(SQLIsASetCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}
func (x *SQLIsASetCondition) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLIsASetCondition) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

/**
 * char1 [ NOT ] { LIKE | LIKEC | LIKE2 | LIKE4 } char2 [ ESCAPE esc_char ]
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Pattern-matching-Conditions.html#GUID-3FA7F5AB-AC64-4200-8F90-294101428C26
 */
type LikeType string

const (
	LIKE  LikeType = "LIKE"
	LIKEC          = "LIKEC"
	LIKE2          = "LIKE2"
	LIKE4          = "LIKE4"
)

type SQLLikeCondition struct {
	*AbstractSQLCondition

	expr expr.ISQLExpr

	Not  bool
	Like LikeType

	pattern expr.ISQLExpr

	escape expr.ISQLExpr
}

func NewLikeCondition() *SQLLikeCondition {
	x := new(SQLLikeCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}
func (x *SQLLikeCondition) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.expr {
		x.SetExpr(target)
		return true
	}
	if source == x.pattern {
		x.SetPattern(target)
		return true
	}
	if source == x.escape {
		x.SetEscape(target)
		return true
	}
	return false
}

func (x *SQLLikeCondition) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLLikeCondition) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}
func (x *SQLLikeCondition) Pattern() expr.ISQLExpr {
	return x.pattern
}
func (x *SQLLikeCondition) SetPattern(pattern expr.ISQLExpr) {
	pattern.SetParent(x)
	x.pattern = pattern
}
func (x *SQLLikeCondition) Escape() expr.ISQLExpr {
	return x.escape
}
func (x *SQLLikeCondition) SetEscape(escape expr.ISQLExpr) {
	escape.SetParent(x)
	x.escape = escape
}

/**
 * REGEXP_LIKE(source_char, pattern [, match_param ])
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Pattern-matching-Conditions.html#GUID-D2124F3A-C6E4-4CCA-A40E-2FFCABFD8E19
 */
type SQLRegexpLikeCondition struct {
	*AbstractSQLCondition
	arguments []expr.ISQLExpr
}

func NewRegexpLikeCondition() *SQLRegexpLikeCondition {
	x := new(SQLRegexpLikeCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}
func (x *SQLRegexpLikeCondition) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	for i, argument := range x.arguments {
		if source == argument {
			return x.SetArgument(i, argument)
		}
	}
	return false
}
func (x *SQLRegexpLikeCondition) Clone() ast.ISQLObject {
	panic("implement me")
}
func (x *SQLRegexpLikeCondition) Arguments() []expr.ISQLExpr {
	return x.arguments
}
func (x *SQLRegexpLikeCondition) SetArgument(index int, argument expr.ISQLExpr) bool {
	if argument == nil {
		return false
	}
	argument.SetParent(x)
	x.arguments[index] = argument
	return true
}
func (x *SQLRegexpLikeCondition) AddArgument(argument expr.ISQLExpr) {
	if argument == nil {
		return
	}
	argument.SetParent(x)
	x.arguments = append(x.arguments, argument)
}

/**
 * NULL condition
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Compound-Conditions.html#GUID-D2A245F5-8071-4DF7-886E-A46F3D13AC80
 */
type SQLNullCondition struct {
	*AbstractSQLCondition
	expr expr.ISQLExpr
}

func NewNullCondition() *SQLNullCondition {
	x := new(SQLNullCondition)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

func (x *SQLNullCondition) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.expr {
		x.SetExpr(target)
		return true
	}
	return false
}
func (x *SQLNullCondition) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLNullCondition) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLNullCondition) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

/**
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/XML-Conditions.html#GUID-37EF3738-5751-4888-9397-50EAD8360D6D
 */
// EQUALS_PATH Condition
// EQUALS_PATH (column, path_string [, correlation_integer ])
type SQLEqualsPathCondition struct {
	*AbstractSQLCondition
}

func NewEqualsPathCondition() *SQLEqualsPathCondition {
	x := new(SQLEqualsPathCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}

// UNDER_PATH Condition
// UNDER_PATH (column [, levels ], path_string [, correlation_integer ] )
type SQLUnderPathCondition struct {
	*AbstractSQLCondition
}

func NewUnderPathCondition() *SQLUnderPathCondition {
	x := new(SQLUnderPathCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}

/**
 * NOT expr
 * https://dev.mysql.com/doc/refman/8.0/en/expressions.html
 *
 * NOT condition
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Compound-Conditions.html#GUID-D2A245F5-8071-4DF7-886E-A46F3D13AC80
 */
type SQLNotExpr struct {
	*AbstractSQLCondition
	condition expr.ISQLExpr
}

func NewNotExpr() *SQLNotExpr {
	x := new(SQLNotExpr)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}
func (x *SQLNotExpr) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.condition {
		x.SetCondition(target)
		return true
	}
	return false
}
func (x *SQLNotExpr) Clone() ast.ISQLObject {
	panic("implement me")
}

func (x *SQLNotExpr) Condition() expr.ISQLExpr {
	return x.condition
}
func (x *SQLNotExpr) SetCondition(condition expr.ISQLExpr) {
	condition.SetParent(x)
	x.condition = condition
}

/**
 * expr1 [NOT] BETWEEN expr2 AND expr3
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/BETWEEN-Condition.html#GUID-868A7C9D-EDF9-44E7-91B5-C3F69E503CCB
 */
type SQLBetweenCondition struct {
	*AbstractSQLCondition
	expr    expr.ISQLExpr
	Not     bool
	between expr.ISQLExpr
	and     expr.ISQLExpr
}

func NewBetweenCondition() *SQLBetweenCondition {
	x := new(SQLBetweenCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}
func (x *SQLBetweenCondition) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.expr {
		x.SetExpr(target)
		return true
	}
	if source == x.between {
		x.SetBetween(target)
		return true
	}
	if source == x.and {
		x.SetAnd(target)
		return true
	}
	return false
}

func (x *SQLBetweenCondition) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLBetweenCondition) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}
func (x *SQLBetweenCondition) Between() expr.ISQLExpr {
	return x.between
}
func (x *SQLBetweenCondition) SetBetween(between expr.ISQLExpr) {
	between.SetParent(x)
	x.between = between
}
func (x *SQLBetweenCondition) And() expr.ISQLExpr {
	return x.and
}
func (x *SQLBetweenCondition) SetAnd(and expr.ISQLExpr) {
	and.SetParent(x)
	x.and = and
}

/**
 * EXISTS (subquery)
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/EXISTS-Condition.html#GUID-20259A83-C42B-4E0D-8DF4-9A2A66ACA8E7
 */
type SQLExistsCondition struct {
	*AbstractSQLCondition
	subQuery select_.ISQLSelectQuery
}

func NewExistsCondition(subQuery select_.ISQLSelectQuery) *SQLExistsCondition {
	x := new(SQLExistsCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	x.SetSubQuery(subQuery)
	return x
}

func (x *SQLExistsCondition) SubQuery() select_.ISQLSelectQuery {
	return x.subQuery
}
func (x *SQLExistsCondition) SetSubQuery(subQuery select_.ISQLSelectQuery) {
	subQuery.SetParent(x)
	x.subQuery = subQuery
}

/**
 * { expr [ NOT ] IN ({ expression_list | subquery })
 | ( expr [, expr ]... ) [ NOT ] IN ({ expression_list [, expression_list ]...| subquery})
}
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/IN-Condition.html#GUID-C7961CB3-8F60-47E0-96EB-BDCF5DB1317C
 */
type SQLInCondition struct {
	*AbstractSQLCondition
	expr   expr.ISQLExpr
	Not    bool
	values []expr.ISQLExpr
}

func NewInCondition() *SQLInCondition {
	x := new(SQLInCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}

func NewInConditionWithExpr(expr expr.ISQLExpr) *SQLInCondition {
	x := new(SQLInCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	x.SetExpr(expr)
	return x
}

func (x *SQLInCondition) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	if source == x.expr {
		x.SetExpr(target)
		return true
	}
	for i, value := range x.values {
		if source == value {
			return x.SetValue(i, target)
		}
	}
	return false
}

func (x *SQLInCondition) Clone() ast.ISQLObject {
	clone := NewInCondition()

	exprClone := x.expr.Clone()
	clone.SetExpr(exprClone)

	for _, child := range x.Values() {
		childClone := child.Clone()
		clone.AddValue(childClone)
	}
	return clone
}

func (x *SQLInCondition) Expr() expr.ISQLExpr {
	return x.expr
}
func (x *SQLInCondition) SetExpr(expr expr.ISQLExpr) {
	expr.SetParent(x)
	x.expr = expr
}

func (x *SQLInCondition) Values() []expr.ISQLExpr {
	return x.values
}
func (x *SQLInCondition) Value(i int) expr.ISQLExpr {
	return x.values[i]
}
func (x *SQLInCondition) SetValue(i int, value expr.ISQLExpr) bool {
	if value == nil {
		return false
	}
	value.SetParent(x)
	x.values[i] = value
	return true
}
func (x *SQLInCondition) AddValue(value expr.ISQLExpr) bool {
	if value == nil {
		return false
	}
	value.SetParent(x)
	x.values = append(x.values, value)
	return true
}

/**
 * expr IS [ NOT ] OF [ TYPE ]
   ([ ONLY ] [ schema. ] type
      [, [ ONLY ] [ schema. ] type ]...
   )
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/IS-OF-type-Condition.html#GUID-7254E4C7-0194-4C1F-A3B2-2CFB0AD907CD
 */
type SQLIsOfCondition struct {
	*AbstractSQLCondition
	expr      expr.ISQLName
	Not       bool
	Type      bool
	arguments []*SQLIsOfConditionArgument
}

type SQLIsOfConditionArgument struct {
	*expr.AbstractSQLExpr
}

/**
 * CURRENT OF <cursor name>
 */
type SQLCurrentOfCondition struct {
	*AbstractSQLCondition
	name expr.ISQLName
}

func NewCurrentOfCondition() *SQLCurrentOfCondition {
	x := new(SQLCurrentOfCondition)
	x.AbstractSQLCondition = NewAbstractSQLCondition()
	return x
}

func (x *SQLCurrentOfCondition) Name() expr.ISQLName {
	return x.name
}
func (x *SQLCurrentOfCondition) SetName(name expr.ISQLName) {
	name.SetParent(x)
	x.name = name
}
