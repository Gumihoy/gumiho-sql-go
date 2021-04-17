package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/condition"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/operator"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/sequence"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type OracleExprParser struct {
	*parser.SQLExprParser
}

func NewExprParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *OracleExprParser {
	return NewExprParserByLexer(NewLexer(sql), dbType, config)
}

func NewExprParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *OracleExprParser {
	var x OracleExprParser
	x.SQLExprParser = parser.NewExprParserByLexer(lexer, dbType, config)
	return &x
}


func (x *OracleExprParser) CreateSQLCommentOnAuditPolicyStatementParser() parser.ISQLCommentStatementParser {
	return nil
}
func (x *OracleExprParser) CreateSQLCommentOnColumnStatementParser() parser.ISQLCommentStatementParser {
	return NewCommentOnColumnStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLCommentOnEditionStatementParser() parser.ISQLCommentStatementParser {
	return nil
}
func (x *OracleExprParser) CreateSQLCommentOnIndextypeStatementParser() parser.ISQLCommentStatementParser {
	return nil
}
func (x *OracleExprParser) CreateSQLCommentOnMaterializedViewStatementParser() parser.ISQLCommentStatementParser {
	return NewCommentOnMaterializedViewStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLCommentOnMiningModelStatementParser() parser.ISQLCommentStatementParser {
	return nil
}
func (x *OracleExprParser) CreateSQLCommentOnOperatorStatementParser() parser.ISQLCommentStatementParser {
	return nil
}
func (x *OracleExprParser) CreateSQLCommentOnTableStatementParser() parser.ISQLCommentStatementParser {
	return NewCommentOnTableStatementParserByExprParser(x)
}

func (x *OracleExprParser) CreateSQLDatabaseStatementParser() parser.ISQLDatabaseStatementParser {
	return NewDatabaseStatementParserByExprParser(x)
}

func (x *OracleExprParser) CreateSQLFunctionStatementParser() parser.ISQLFunctionStatementParser {
	return NewFunctionStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLIndexStatementParser() parser.ISQLIndexStatementParser {
	return NewIndexStatementParserByExprParser(x)
}

func (x *OracleExprParser) CreateSQLPackageStatementParser() parser.ISQLPackageStatementParser {
	return NewPackageStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLPackageBodyStatementParser() parser.ISQLPackageBodyStatementParser {
	return NewPackageBodyStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLProcedureStatementParser() parser.ISQLProcedureStatementParser {
	return NewProcedureStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLRoleStatementParser() parser.ISQLRoleStatementParser {
	return NewRoleStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLSchemaStatementParser() parser.ISQLSchemaStatementParser {
	return NewSchemaStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLSequenceStatementParser() parser.ISQLSequenceStatementParser {
	return NewSequenceStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLSynonymStatementParser() parser.ISQLSynonymStatementParser {
	return NewSynonymStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLTableStatementParser() parser.ISQLTableStatementParser {
	return NewTableStatementParserByExprParser(x)
}

func (x *OracleExprParser) CreateSQLTriggerStatementParser() parser.ISQLTriggerStatementParser {
	return NewTriggerStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLTypeStatementParser() parser.ISQLTypeStatementParser {
	return NewTypeStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLTypeBodyStatementParser() parser.ISQLTypeBodyStatementParser {
	return NewTypeBodyStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLUserStatementParser() parser.ISQLUserStatementParser {
	return NewUserStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLViewStatementParser() parser.ISQLViewStatementParser {
	return NewViewStatementParserByExprParser(x)
}

func (x *OracleExprParser) CreateSQLDeleteStatementParser() parser.ISQLDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLInsertStatementParser() parser.ISQLInsertStatementParser {
	return NewInsertStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLSelectStatementParser() parser.ISQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(x)
}
func (x *OracleExprParser) CreateSQLUpdateSStatementParser() parser.ISQLUpdateStatementParser {
	return NewUpdateStatementParserByExprParser(x)
}

func (x *OracleExprParser) CreateSQLExplainStatementParser() parser.ISQLExplainStatementParser {
	return NewExplainStatementParserByExprParser(x)
}


/**
 * DELETE [ hint ]
   [ FROM ]
   { dml_table_expression_clause
   | ONLY (dml_table_expression_clause)
   } [ t_alias ]
     [ where_clause ]
     [ returning_clause ]
     [error_logging_clause];
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/DELETE.html#GUID-156845A5-B626-412B-9F95-8869B988ABD7
 */
func (sp *OracleExprParser) ParseDeleteStatement(child parser.ISQLExprParser) statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DELETE) {
		return nil
	}
	x := statement.NewDeleteStatement(sp.DBType())

	from := sp.AcceptAndNextToken(parser.FROM)
	x.From = from

	tableReference := parser.ParseTableReference(child)
	x.SetTableReference(tableReference)

	whereClause := parser.ParseWhereClause(child)
	x.SetWhereClause(whereClause)

	returningClause := parser.ParseIReturningClause(child)
	x.SetReturningClause(returningClause)

	return x
}

func (sp *OracleExprParser) ParseNameRest(child parser.ISQLExprParser, owner expr.ISQLName) expr.ISQLName {
	switch sp.Kind() {
	case parser.SYMB_DOT:
		parser.NextTokenByParser(sp)
		right := child.ParseIdentifier(child)
		name := expr.NewNameWithOwnerAndName(owner, right)
		newName := sp.ParseNameRest(child, name)
		return newName

	case parser.SYMB_AT:
		parser.NextTokenByParser(sp)
		dbLink := parser.ParseName(child)
		name := expr.NewDBLinkExpr(owner, dbLink)
		return name

	}
	return owner
}

/**
 * OuterJoin: (+)
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Joins.html#GUID-29A4584C-0741-4E6A-A89B-DCFAA222994A
 */
func (self *OracleExprParser) ParseOuterJoinExprRest(x expr.ISQLExpr) (expr.ISQLExpr, bool) {
	mark := self.Mark()
	if !self.AcceptAndNextToken(parser.SYMB_LEFT_PAREN) {
		return x, false
	}

	if !self.AcceptAndNextToken(parser.SYMB_PLUS) {
		self.ResetWithMark(mark)
		return x, false
	}

	self.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

	return expr.NewOuterJoinExpr(x), true
}

/**
 * =>
 */
func (this *OracleExprParser) ParseCallExpr(child parser.ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {
	if !this.AcceptAndNextToken(parser.SYMB_EQUAL_GREATER_THAN) {
		return name
	}

	value := parser.ParseExpr(child)
	return expr.NewCallExprWithNameAndValue(name, value)
}

func (sp *OracleExprParser) ParseExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if sp.Accept(parser.EOF) ||
		sp.Accept(parser.SYMB_SEMI) {
		return left
	}

	left = sp.ParseMultiplicativeOperatorExprRest(child, left)
	left = sp.ParseAdditiveOperatorExprRest(child, left)
	left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	left = sp.ParseLowPriorityComparisonOperatorExprRest(child, left)
	left = sp.ParseNotExprRest(child, left)
	left = sp.ParseAndOperatorExprRest(child, left)
	left = sp.ParseOrOperatorExprRest(child, left)
	return left

}

/**
* E: T op T op T ...
* T: primaryExpr
* OP: *, /
* https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/About-SQL-Operators.html#GUID-FEF44762-F45C-41D9-B380-F6A61AD25338
 */
func (sp *OracleExprParser) ParseMultiplicativeOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := parser.ParsePrimaryExpr(child)
	return sp.ParseMultiplicativeOperatorExprRest(child, left)
}

func (sp *OracleExprParser) ParseMultiplicativeOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := sp.Token().Kind
	switch kind {
	case parser.SYMB_STAR:
		parser.NextTokenByParser(sp)
		right := parser.ParsePrimaryExpr(child)
		left = operator.NewBinaryOperator(left, operator.MULTIPLY, right)
		return sp.ParseMultiplicativeOperatorExprRest(child, left)
	case parser.SYMB_SLASH:
		parser.NextTokenByParser(sp)
		right := parser.ParsePrimaryExpr(child)
		left = operator.NewBinaryOperator(left, operator.DIVIDE, right)
		return sp.ParseMultiplicativeOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T (+/-/||) T (+/-/||) T ...
* T: X (*, /) X
* +, -, ||
* https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/About-SQL-Operators.html#GUID-FEF44762-F45C-41D9-B380-F6A61AD25338
 */
func (sp *OracleExprParser) ParseAdditiveOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseMultiplicativeOperatorExpr(child)
	return sp.ParseAdditiveOperatorExprRest(child, left)
}

func (sp *OracleExprParser) ParseAdditiveOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch sp.Kind() {
	case parser.SYMB_PLUS:
		parser.NextTokenByParser(sp)
		right := sp.ParseMultiplicativeOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.PLUS, right)
		return sp.ParseAdditiveOperatorExprRest(child, left)
	case parser.SYMB_MINUS:
		parser.NextTokenByParser(sp)
		right := sp.ParseMultiplicativeOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MINUS, right)
		return sp.ParseAdditiveOperatorExprRest(child, left)
	case parser.SYMB_LOGICAL_OR:
		parser.NextTokenByParser(sp)
		right := sp.ParseMultiplicativeOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.CONCAT, right)
		return sp.ParseAdditiveOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X | X
* op: =, !=, <>, ~=, <, >, <=, >=,
* https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/About-SQL-Conditions.html#GUID-65B103FE-C00C-46A3-8173-A731DBF62C80
 */
func (sp *OracleExprParser) ParseHighPriorityComparisonOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseAdditiveOperatorExpr(child)
	return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
}

func (sp *OracleExprParser) ParseHighPriorityComparisonOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch sp.Kind() {
	case parser.SYMB_EQUAL:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.EQ, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_EXCLAMATION_EQUAL:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.NOT_EQ1, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_LESS_THAN_GREATER_THAN:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.NOT_EQ2, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_XOR_EQUAL:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.NOT_EQ3, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_NOT_EQUAL:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.NOT_EQ4, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_LESS_THAN:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.LESS_THAN, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_LESS_THAN_EQUAL:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.LESS_THAN_EQ, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_GREATER_THAN:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	case parser.SYMB_GREATER_THAN_EQUAL:
		parser.NextTokenByParser(sp)
		right := sp.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN_EQ, right)
		return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X | X
* op: IS [NOT] NULL, LIKE, [NOT] BETWEEN, [NOT] IN, EXISTS, IS OF type
* https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/About-SQL-Conditions.html#GUID-65B103FE-C00C-46A3-8173-A731DBF62C80
 */
func (sp *OracleExprParser) ParseLowPriorityComparisonOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseHighPriorityComparisonOperatorExpr(child)
	return sp.ParseLowPriorityComparisonOperatorExprRest(child, left)
}

func (sp *OracleExprParser) ParseLowPriorityComparisonOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if sp.Accept(parser.IS) {

		return sp.ParseIsConditionRest(child, left)

	} else if sp.AcceptAndNextToken(parser.NOT) {

		if sp.IsParseLikeCondition() {

			return sp.ParseLikeConditionRest(child, true, left)

		} else if sp.Accept(parser.BETWEEN) {

			return sp.ParseBetweenConditionRest(child, true, left)

		} else if sp.Accept(parser.IN) {

			return sp.ParseInConditionRest(child, true, left)
		}

	} else if sp.IsParseLikeCondition() {

		return sp.ParseLikeConditionRest(child, false, left)

	} else if sp.Accept(parser.BETWEEN) {

		return sp.ParseBetweenConditionRest(child, false, left)

	} else if sp.Accept(parser.IN) {

		return sp.ParseInConditionRest(child, false, left)

	} else if sp.Accept(parser.EXISTS) {

	}
	return left
}

/**
 * expr IS [NOT] {TRUE | FALSE | UNKNOWN | NULL | EMPTY | NAN | INFINITE}
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Floating-Point-Conditions.html#GUID-D7707649-2C93-4553-BF78-F461F17A634E
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Multiset-Conditions.html#GUID-EED3C932-8A77-4841-BCC0-CD524F1E65A1
 */
func (sp *OracleExprParser) ParseIsConditionRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(parser.IS) {
		return left
	}
	x := condition.NewIsCondition()
	x.SetExpr(left)

	x.Not = sp.AcceptAndNextToken(parser.NOT)

	if sp.AcceptAndNextToken(parser.NULL) {

		x.Value = condition.NULL

	} else if sp.AcceptAndNextToken(parser.NAN) {

		x.Value = condition.NAN

	} else if sp.AcceptAndNextToken(parser.INFINITE) {

		x.Value = condition.INFINITE

	} else if sp.AcceptAndNextToken(parser.EMPTY) {

		x.Value = condition.EMPTY

	} else {
		panic(sp.UnSupport())
	}

	return x
}

/**
 * LIKE | LIKEC | LIKE2 | LIKE4
 */
func (sp *OracleExprParser) IsParseLikeCondition() bool {
	return sp.Accept(parser.LIKE) ||
		sp.Accept(parser.LIKEC) ||
		sp.Accept(parser.LIKE2) ||
		sp.Accept(parser.LIKE4)
}
func (sp *OracleExprParser) ParseLikeConditionRest(child parser.ISQLExprParser, not bool, left expr.ISQLExpr) expr.ISQLExpr {
	if !sp.IsParseLikeCondition() {
		return left
	}

	var like condition.LikeType
	if sp.AcceptAndNextToken(parser.LIKE) {
		like = condition.LIKE

	} else if sp.AcceptAndNextToken(parser.LIKEC) {
		like = condition.LIKEC

	} else if sp.AcceptAndNextToken(parser.LIKE2) {
		like = condition.LIKE2

	} else if sp.AcceptAndNextToken(parser.LIKE4) {
		like = condition.LIKE4

	}

	x := condition.NewLikeCondition()
	x.SetExpr(left)
	x.Not = not
	x.Like = like

	pattern := parser.ParseExpr(child)
	x.SetPattern(pattern)

	if sp.AcceptAndNextToken(parser.ESCAPE) {
		escape := parser.ParseExpr(child)
		x.SetEscape(escape)
	}
	return x
}

/**
* E: NOT T
* T:
* op: NOT
* https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/About-SQL-Conditions.html#GUID-65B103FE-C00C-46A3-8173-A731DBF62C80
*/
func (sp *OracleExprParser) ParseNotExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseLowPriorityComparisonOperatorExpr(child)
	return sp.ParseNotExprRest(child, left)
}

func (sp *OracleExprParser) ParseNotExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch left.(type) {
	case *expr.SQLUnQuotedIdentifier:
		left, _ := left.(*expr.SQLUnQuotedIdentifier)
		if parser.NOT.Equals(left.Name()) {
			x := condition.NewNotExpr()
			condition := parser.ParseExpr(child)
			x.SetCondition(condition)
			return x
		}
	}
	return left
}

/**
* E: T op T op T
* T: X comparisonOperator X
* op: AND
* https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/About-SQL-Conditions.html#GUID-65B103FE-C00C-46A3-8173-A731DBF62C80
 */
func (sp *OracleExprParser) ParseAndOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseNotExpr(child)
	return sp.ParseAndOperatorExprRest(child, left)
}

func (sp *OracleExprParser) ParseAndOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if sp.AcceptAndNextToken(parser.AND) {
		right := sp.ParseNotExpr(child)
		left = operator.NewBinaryOperator(left, operator.AND, right)
		return sp.ParseAndOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T
* T: X AND X
* op: OR
* https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/About-SQL-Conditions.html#GUID-65B103FE-C00C-46A3-8173-A731DBF62C80
 */
func (sp *OracleExprParser) ParseOrOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseAndOperatorExpr(child)
	return sp.ParseOrOperatorExprRest(child, left)
}

func (sp *OracleExprParser) ParseOrOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if sp.AcceptAndNextToken(parser.OR) {
		right := sp.ParseAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.OR, right)
		return sp.ParseOrOperatorExprRest(child, left)
	}
	return left
}

// ------------------------------------------------------ DDL --------------------------------------------------
// --------------------------- Sequence --------------------------------------------------
func (sp *OracleExprParser) ParseSequenceOption(child parser.ISQLExprParser) expr.ISQLExpr {
	option := sp.SQLExprParser.ParseSequenceOption(child)
	if option != nil {
		return option
	}
	if sp.AcceptAndNextToken(parser.NOMAXVALUE) {
		return sequence.NewNoMaxValueSequenceOption()

	} else if sp.AcceptAndNextToken(parser.NOMINVALUE) {
		return sequence.NewNoMinValueSequenceOption()

	} else if sp.AcceptAndNextToken(parser.NOCYCLE) {
		return sequence.NewNoCycleSequenceOption()

	} else if sp.AcceptAndNextToken(parser.CACHE) {
		x := sequence.NewCacheSequenceOption()
		value := parser.ParseExpr(child)
		x.SetValue(value)
		return x

	} else if sp.AcceptAndNextToken(parser.NOCACHE) {
		return sequence.NewNoCacheSequenceOption()

	} else if sp.AcceptAndNextToken(parser.ORDER) {
		return sequence.NewOrderSequenceOption()

	} else if sp.AcceptAndNextToken(parser.NOORDER) {
		return sequence.NewNoOrderSequenceOption()

	} else if sp.AcceptAndNextToken(parser.KEEP) {
		return sequence.NewKeepSequenceOption()

	} else if sp.AcceptAndNextToken(parser.NOKEEP) {
		return sequence.NewNoKeepSequenceOption()

	} else if sp.AcceptAndNextToken(parser.SCALE) {
		return sequence.NewScaleSequenceOption()

	} else if sp.AcceptAndNextToken(parser.NOSCALE) {
		return sequence.NewNoScaleSequenceOption()

	} else if sp.AcceptAndNextToken(parser.SESSION) {
		return sequence.NewSessionSequenceOption()

	} else if sp.AcceptAndNextToken(parser.GLOBAL) {
		return sequence.NewGlobalSequenceOption()

	}
	return nil
}

// --------------------------- OnTable --------------------------------------------------
func (sp *OracleExprParser) IsParseColumnConstraint() bool {
	if sp.Accept(parser.CONSTRAINT) ||
		sp.Accept(parser.NOT) ||
		sp.Accept(parser.NULL) ||
		sp.Accept(parser.UNIQUE) ||
		sp.Accept(parser.PRIMARY) ||
		sp.Accept(parser.FOREIGN) ||
		sp.Accept(parser.CHECK) ||
		sp.Accept(parser.SCOPE) ||
		sp.Accept(parser.WITH) {
		return true
	}
	return false
}

func (sp *OracleExprParser) ParseColumnConstraint(child parser.ISQLExprParser) (table.ISQLColumnConstraint, bool) {
	if !sp.IsParseColumnConstraint() {
		return nil, false
	}

	x, ok := sp.SQLExprParser.ParseColumnConstraint(child)
	if ok {
		return x, true
	}

	return nil, false
}

func (sp *OracleExprParser) IsParseTableConstraint() bool {
	if sp.Accept(parser.CONSTRAINT) ||
		sp.Accept(parser.UNIQUE) ||
		sp.Accept(parser.PRIMARY) ||
		sp.Accept(parser.FOREIGN) ||
		sp.Accept(parser.CHECK) ||
		sp.Accept(parser.SCOPE) ||
		sp.Accept(parser.REF) {
		return true
	}
	return false
}

func (sp *OracleExprParser) ParseTableConstraint(child parser.ISQLExprParser) table.ISQLTableConstraint {
	if !child.IsParseTableConstraint() {
		return nil
	}

	x := sp.SQLExprParser.ParseTableConstraint(child)
	if x != nil {
		return x
	}

	return nil
}

// --------------- VIEW

func (sp *OracleExprParser) ParseViewColumnOption(child parser.ISQLExprParser) (expr.ISQLExpr, bool) {
	if sp.Accept(parser.VISIBLE) {

	} else if sp.Accept(parser.INVISIBLE) {

	} else if child.IsParseColumnConstraint() {
		return child.ParseColumnConstraint(child)
	}

	return nil, false
}

// --------------------------- DML --------------------------------------------------
// --------------- SELECT

func (sp *OracleExprParser) IsParserSubAvFactoringClauseRest() bool {
	return sp.Accept(parser.ANALYTIC)
}

func (sp *OracleExprParser) ParseTableReference(child parser.ISQLExprParser) select_.ISQLTableReference {
	if !parser.IsIdentifier(sp.Kind()) {
		return nil
	}

	x := select_.NewTableReference()

	name := parser.ParseName(sp)
	x.SetName(name)

	partitionExtensionClause := sp.ParsePartitionExtensionClause(sp)
	x.SetPartitionExtensionClause(partitionExtensionClause)

	sameClause := sp.ParseSameClause(sp)
	x.SetSampleClause(sameClause)

	as := false
	if sp.AcceptAndNextToken(parser.AS) {
		as = true
	}
	x.SetAs(as)

	alias := child.ParseIdentifier(sp)
	x.SetAlias(alias)

	if as && alias == nil {
		panic("")
	}

	return x
}
func (sp *OracleExprParser) ParseSameClause(child parser.ISQLExprParser) *select_.SQLSampleClause {
	if !sp.AcceptAndNextToken(parser.SAMPLE) {
		return nil
	}
	x := select_.NewSampleClause()
	sp.AcceptAndNextToken(parser.BLOCK)

	sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
	percent := parser.ParseExpr(child)
	sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
	x.SetPercent(percent)

	if sp.AcceptAndNextToken(parser.SEED) {
		sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
		seedValue := parser.ParseExpr(child)
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
		x.SetSeedValue(seedValue)
	}

	return x
}
func (sp *OracleExprParser) ParsePartitionExtensionClause(child parser.ISQLExprParser) expr.ISQLExpr {
	if !sp.Accept(parser.PARTITION) && !sp.Accept(parser.SUBPARTITION) {
		return nil
	}

	if sp.AcceptAndNextToken(parser.PARTITION) {

		if sp.AcceptAndNextToken(parser.FOR) {
			x := select_.NewPartitionForClause()
			sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
			for {
				name := parser.ParseExpr(child)
				x.AddName(name)
				if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
					break
				}
			}
			sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
			return x
		}

		x := select_.NewPartitionClause()
		sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
		for {
			name := parser.ParseExpr(child)
			x.AddName(name)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

		return x
	} else if sp.AcceptAndNextToken(parser.SUBPARTITION) {

		if sp.AcceptAndNextToken(parser.FOR) {
			x := select_.NewSubPartitionForClause()

			sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
			for {
				name := parser.ParseExpr(child)
				x.AddName(name)
				if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
					break
				}
			}
			sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

			return x
		}

		x := select_.NewSubPartitionForClause()

		sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
		for {
			name := parser.ParseExpr(child)
			x.AddName(name)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

		return x

	}

	return nil
}

func (sp *OracleExprParser) ParseOffsetFetchClause(child parser.ISQLExprParser) select_.ISQLLimitClause {
	if !sp.Accept(parser.OFFSET) && !sp.Accept(parser.FETCH) {
		return nil
	}

	x := select_.NewOffsetFetchClause()

	if sp.AcceptAndNextToken(parser.OFFSET) {

		offsetExpr := parser.ParseExpr(child)
		x.SetOffsetExpr(offsetExpr)
	}

	if sp.AcceptAndNextToken(parser.FETCH) {

		countExpr := parser.ParseExpr(child)
		x.SetCountExpr(countExpr)
	}

	return x
}
func (sp OracleExprParser) ParseLockClause(child parser.ISQLExprParser) select_.ISQLLockClause {
	if sp.Accept(parser.FOR) {

		return sp.ParseForUpdate(child)
	}
	return nil
}

/**
* FOR UPDATE [ OF table_name [, ...] ] [ NOWAIT | SKIP LOCKED | WAIT expr ] [...]
*/
func (sp *OracleExprParser) ParseForUpdate(child parser.ISQLExprParser) select_.ISQLLockClause {
	if !sp.AcceptAndNextToken(parser.FOR) {
		return nil
	}

	if sp.AcceptAndNextToken(parser.UPDATE) {

		x := select_.NewForUpdateClause()
		sp.ParseForUpdateRest(child, x.AbstractSQLLockForClause)
		return x

	} else {

		panic(sp.UnSupport())
	}

}

func (sp *OracleExprParser) ParseForUpdateRest(child parser.ISQLExprParser, x *select_.AbstractSQLLockForClause) {
	if sp.AcceptAndNextToken(parser.OF) {
		for {
			x.AddTable(parser.ParseName(child))
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
	}

	if sp.AcceptAndNextToken(parser.NOWAIT) {

		x.SetWaitExpr(select_.NewLockForNoWaitExpr())

	} else if sp.AcceptAndNextToken(parser.SKIP) {

		sp.AcceptAndNextTokenWithError(parser.LOCKED, true)
		x.SetWaitExpr(select_.NewLockForSkipLockedExpr())

	} else if sp.AcceptAndNextToken(parser.WAIT) {

		value := parser.ParseExpr(child)
		x.SetWaitExpr(select_.NewLockForWaitExprWithValue(value))
	}

}
