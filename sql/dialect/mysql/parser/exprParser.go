package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/function"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/literal"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/operator"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/select"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/user"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/variable"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/view"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
)

var complexFunctionNameMap = make(map[string]bool)
var nonParametricFunctionNameMap = make(map[string]bool)

func init() {
	complexFunctionNameMap["CAST"] = true
}

type MySQLExprParser struct {
	*parser.SQLExprParser
}

func NewExprParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *MySQLExprParser {
	return NewExprParserByLexer(parser.NewLexer(sql), dbType, config)
}

func NewExprParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *MySQLExprParser {
	return NewExprParserByExprParser(parser.NewExprParserByLexer(lexer, dbType, config))
}

func NewExprParserByExprParser(parser *parser.SQLExprParser) *MySQLExprParser {
	x := new(MySQLExprParser)
	x.SQLExprParser = parser
	return x
}

func (x *MySQLExprParser) CreateSQLDatabaseStatementParser() parser.ISQLDatabaseStatementParser {
	return NewDatabaseStatementParserByExprParser(x)
}

func (x *MySQLExprParser) CreateSQLFunctionStatementParser() parser.ISQLFunctionStatementParser {
	return NewFunctionStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLIndexStatementParser() parser.ISQLIndexStatementParser {
	return NewIndexStatementParserByExprParser(x)
}

func (x *MySQLExprParser) CreateSQLPackageStatementParser() parser.ISQLPackageStatementParser {
	return NewPackageStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLPackageBodyStatementParser() parser.ISQLPackageBodyStatementParser {
	return NewPackageBodyStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLProcedureStatementParser() parser.ISQLProcedureStatementParser {
	return NewProcedureStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLRoleStatementParser() parser.ISQLRoleStatementParser {
	return NewRoleStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLSchemaStatementParser() parser.ISQLSchemaStatementParser {
	return NewSchemaStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLSequenceStatementParser() parser.ISQLSequenceStatementParser {
	return NewSequenceStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLServerStatementParser() parser.ISQLServerStatementParser {
	return NewServerStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLSynonymStatementParser() parser.ISQLSynonymStatementParser {
	return NewSynonymStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLTableStatementParser() parser.ISQLTableStatementParser {
	return NewTableStatementParserByExprParser(x)
}

func (x *MySQLExprParser) CreateSQLTriggerStatementParser() parser.ISQLTriggerStatementParser {
	return NewTriggerStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLTypeStatementParser() parser.ISQLTypeStatementParser {
	return parser.NewTypeStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLTypeBodyStatementParser() parser.ISQLTypeBodyStatementParser {
	return parser.NewTypeBodyStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLUserStatementParser() parser.ISQLUserStatementParser {
	return NewUserStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLViewStatementParser() parser.ISQLViewStatementParser {
	return NewViewStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLDeleteStatementParser() parser.ISQLDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLInsertStatementParser() parser.ISQLInsertStatementParser {
	return NewInsertStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLSelectStatementParser() parser.ISQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLUpdateSStatementParser() parser.ISQLUpdateStatementParser {
	return NewUpdateStatementParserByExprParser(x)
}

func (x *MySQLExprParser) CreateSQLSetVariableAssignmentStatementParser() parser.ISQLSetStatementParser {
	return NewSetVariableAssignmentStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLSetCharacterSetStatementParser() parser.ISQLSetStatementParser {
	return NewSetCharacterSetStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLSetCharsetStatementParser() parser.ISQLSetStatementParser {
	return NewSetCharsetStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLSetNamesStatementParser() parser.ISQLSetStatementParser {
	return NewSetNamesStatementParserByExprParser(x)
}

func (x *MySQLExprParser) CreateSQLShowCreateDatabaseStatementParser() parser.ISQLShowStatementParser {
	return NewShowCreateDatabaseByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLShowCreateEventStatementParser() parser.ISQLShowStatementParser {
	return NewShowCreateEventByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLShowCreateFunctionStatementParser() parser.ISQLShowStatementParser {
	return NewShowCreateFunctionByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLShowCreateProcedureStatementParser() parser.ISQLShowStatementParser {
	return NewShowCreateProcedureByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLShowCreateTableStatementParser() parser.ISQLShowStatementParser {
	return NewShowCreateTableByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLShowCreateTriggerStatementParser() parser.ISQLShowStatementParser {
	return NewShowCreateTriggerByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLShowCreateViewStatementParser() parser.ISQLShowStatementParser {
	return NewShowCreateViewByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLShowDatabasesStatementParser() parser.ISQLShowStatementParser {
	return parser.NewShowDatabasesByExprParser(x)
}

func (x *MySQLExprParser) CreateSQLDescStatementParser() parser.ISQLExplainStatementParser {
	return NewDescStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLDescribeStatementParser() parser.ISQLExplainStatementParser {
	return NewDescribeStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLExplainStatementParser() parser.ISQLExplainStatementParser {
	return NewExplainStatementParserByExprParser(x)
}

func (x *MySQLExprParser) CreateSQLHelpStatementParser() parser.ISQLHelpStatementParser {
	return NewHelpStatementParserByExprParser(x)
}
func (x *MySQLExprParser) CreateSQLUseStatementParser() parser.ISQLUseStatementParser {
	return NewUseStatementParserByExprParser(x)
}


/**
 * DELETE [LOW_PRIORITY] [QUICK] [IGNORE] FROM tbl_name [[AS] tbl_alias]
    [PARTITION (partition_name [, partition_name] ...)]
    [WHERE where_condition]
    [ORDER BY ...]
    [LIMIT row_count]
 * https://dev.mysql.com/doc/refman/8.0/en/delete.html
 */
func (sp *MySQLExprParser) ParseDeleteStatement(child parser.ISQLExprParser) statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.DELETE) {
		return nil
	}
	x := statement.NewDeleteStatement(sp.DBType())

	lowPriority := sp.AcceptAndNextToken(parser.LOW_PRIORITY)
	x.LowPriority = lowPriority

	quick := sp.AcceptAndNextToken(parser.QUICK)
	x.Quick = quick

	ignore := sp.AcceptAndNextToken(parser.IGNORE)
	x.Ignore = ignore

	if !sp.Accept(parser.FROM) && parser.IsIdentifier(sp.Kind()) {
		for {
			table := parser.ParseName(child)
			x.AddTable(table)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
	}

	sp.AcceptAndNextTokenWithError(parser.FROM, true)
	tableReference := parser.ParseTableReference(child)
	x.SetTableReference(tableReference)

	if sp.AcceptAndNextToken(parser.USING) {
		usingTableReference := parser.ParseTableReference(child)
		x.SetUsingTableReference(usingTableReference)
	}

	whereClause := parser.ParseWhereClause(child)
	x.SetWhereClause(whereClause)

	orderByClause := parser.ParseOrderByClause(child)
	x.SetOrderByClause(orderByClause)

	limitClause := parser.ParseLimitClause(child)
	x.SetLimitClause(limitClause)

	return x
}

/**
 * INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    { {VALUES | VALUE} (value_list) [, (value_list)] ...
      |
      VALUES row_constructor_list
    }
    [AS row_alias[(col_alias [, col_alias] ...)]]
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | DELAYED | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [AS row_alias[(col_alias [, col_alias] ...)]]
    SET assignment_list
    [ON DUPLICATE KEY UPDATE assignment_list]

INSERT [LOW_PRIORITY | HIGH_PRIORITY] [IGNORE]
    [INTO] tbl_name
    [PARTITION (partition_name [, partition_name] ...)]
    [(col_name [, col_name] ...)]
    [AS row_alias[(col_alias [, col_alias] ...)]]
    {SELECT ... | TABLE table_name}
    [ON DUPLICATE KEY UPDATE assignment_list]
 * https://dev.mysql.com/doc/refman/8.0/en/insert.html
 */
func (sp *MySQLExprParser) ParseInsertStatement(child parser.ISQLExprParser) statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.INSERT) {
		return nil
	}

	x := statement.NewInsertStatement(sp.DBType())

	if sp.AcceptAndNextToken(parser.LOW_PRIORITY) {
		x.PriorityKind = statement.LOW_PRIORITY

	} else if sp.AcceptAndNextToken(parser.DELAYED) {
		x.PriorityKind = statement.DELAYED

	} else if sp.AcceptAndNextToken(parser.HIGH_PRIORITY) {
		x.PriorityKind = statement.HIGH_PRIORITY
	}

	if sp.AcceptAndNextToken(parser.IGNORE) {
		x.Ignore = true
	}

	into := sp.AcceptAndNextToken(parser.INTO)
	x.Into = into

	tableReference := sp.ParseTableReference(child)
	x.SetTableReference(tableReference)

	if sp.AcceptAndNextToken(parser.SYMB_LEFT_PAREN) {

		for {
			column := child.ParseIdentifier(child)
			x.AddColumn(column)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}

		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

	} else if sp.AcceptAndNextToken(parser.DEFAULT) {
		sp.AcceptAndNextTokenWithError(parser.VALUES, true)
	}

	if sp.Accept(parser.VALUE) || sp.Accept(parser.VALUES) {
		if sp.AcceptAndNextToken(parser.VALUE) {
			x.ValueKind = statement.VALUE

		} else if sp.AcceptAndNextToken(parser.VALUES) {
			x.ValueKind = statement.VALUES
		}

		for {
			value := parser.ParseExpr(child)
			x.AddValue(value)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
	}

	if sp.AcceptAndNextToken(parser.AS) {
		panic(sp.UnSupport())
	}

	if parser.IsSelect(sp) {
		subQuery := parser.ParseSelectQuery(child)
		x.SetSubQuery(subQuery)
	}

	if sp.AcceptAndNextToken(parser.TABLE) {
		parser.ParseName(child)

	}

	if sp.AcceptAndNextToken(parser.ON) {
		// ON DUPLICATE KEY UPDATE
		sp.AcceptAndNextTokenWithError(parser.DUPLICATE, true)
		sp.AcceptAndNextTokenWithError(parser.KEY, true)
		sp.AcceptAndNextTokenWithError(parser.UPDATE, true)
		for {
			assignment := parser.ParseExpr(child)
			x.AddUpdateAssignment(assignment)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
	}

	return x
}

/**
 * UPDATE [LOW_PRIORITY] [IGNORE] table_reference
    SET assignment_list
    [WHERE where_condition]
    [ORDER BY ...]
    [LIMIT row_count]
 * https://dev.mysql.com/doc/refman/8.0/en/update.html
 */
func (sp *MySQLExprParser) ParseUpdateStatement(child parser.ISQLExprParser) statement.ISQLStatement {
	if !sp.AcceptAndNextToken(parser.UPDATE) {
		return nil
	}
	x := statement.NewUpdateStatement(sp.DBType())

	if sp.AcceptAndNextToken(parser.LOW_PRIORITY) {
		x.LowPriority = true
	}
	if sp.AcceptAndNextToken(parser.IGNORE) {
		x.Ignore = true
	}

	tableReference := parser.ParseTableReference(child)
	x.SetTableReference(tableReference)

	sp.AcceptAndNextTokenWithError(parser.SET, true)
	for {
		assignment := parser.ParseExpr(child)
		x.AddAssignment(assignment)
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	whereClause := parser.ParseWhereClause(sp)
	x.SetWhereClause(whereClause)

	orderByClause := parser.ParseOrderByClause(sp)
	x.SetOrderByClause(orderByClause)

	limitClause := parser.ParseLimitClause(sp)
	x.SetLimitClause(limitClause)

	return x
}

func (sp MySQLExprParser) ParseIdentifier(child parser.ISQLExprParser) expr.ISQLIdentifier {
	var name expr.ISQLIdentifier
	switch sp.Kind() {
	case parser.IDENTIFIER:
		name = expr.NewUnQuotedIdentifier(sp.StringValue())
		parser.NextTokenByParser(child)

	case parser.IDENTIFIER_DOUBLE_QUOTE:
		name = expr.NewDoubleQuotedIdentifier(sp.StringValue())
		parser.NextTokenByParser(child)

	case parser.IDENTIFIER_REVERSE_QUOTE:
		name = expr.NewReverseQuotedIdentifier(sp.StringValue())
		parser.NextTokenByParser(sp)

	case parser.SYMB_STAR:
		name = expr.NewAllColumnExpr()
		parser.NextTokenByParser(sp)

	case parser.LITERAL_STRING:
		name = literal.NewStringLiteral(sp.StringValue())
		parser.NextTokenByParser(sp)
	}

	return name
}

func (sp *MySQLExprParser) ParseNameRest(child parser.ISQLExprParser, owner expr.ISQLName) expr.ISQLName {
	switch sp.Kind() {
	case parser.SYMB_DOT:
		parser.NextTokenByParser(sp)
		right := child.ParseIdentifier(child)
		name := expr.NewNameWithOwnerAndName(owner, right)
		newName := sp.ParseNameRest(child, name)
		return newName
	case parser.SYMB_AT:
		parser.NextTokenByParser(sp)
		host := parser.ParseExpr(child)
		name := user.NewUserNameWithNameAndHost(owner, host)
		return name

	}
	return owner
}

/**
 * @name
 * @@name
 */
func (sp *MySQLExprParser) ParseAtExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	if !sp.AcceptAndNextToken(parser.SYMB_AT) {
		return nil
	}

	if sp.AcceptAndNextToken(parser.SYMB_AT) {
		if sp.AcceptAndNextToken(parser.GLOBAL) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_DOT, true)
			x := variable.NewAtAtVariableExpr()

			x.Kind = variable.GLOBAL
			x.HasAtAT = true

			name := parser.ParseName(child)
			x.SetName(name)

			return x

		} else if sp.AcceptAndNextToken(parser.LOCAL) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_DOT, true)
			x := variable.NewAtAtVariableExpr()

			x.Kind = variable.LOCAL
			x.HasAtAT = true

			name := parser.ParseName(child)
			x.SetName(name)

			return x

		} else if sp.AcceptAndNextToken(parser.PERSIST) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_DOT, true)
			x := variable.NewAtAtVariableExpr()

			x.Kind = variable.PERSIST
			x.HasAtAT = true

			name := parser.ParseName(child)
			x.SetName(name)

			return x

		} else if sp.AcceptAndNextToken(parser.PERSIST_ONLY) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_DOT, true)
			x := variable.NewAtAtVariableExpr()

			x.Kind = variable.PERSIST_ONLY
			x.HasAtAT = true

			name := parser.ParseName(child)
			x.SetName(name)
			return x

		} else if sp.AcceptAndNextToken(parser.SESSION) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_DOT, true)
			x := variable.NewAtAtVariableExpr()

			x.Kind = variable.SESSION
			x.HasAtAT = true

			name := parser.ParseName(child)
			x.SetName(name)

			return x
		} else if parser.IsIdentifier(sp.Kind()) {
			x := variable.NewAtAtVariableExpr()

			x.HasAtAT = true

			name := parser.ParseName(child)
			x.SetName(name)
			return x
		}

	} else if parser.IsIdentifier(sp.Kind()) {
		x := variable.NewAtVariableExpr()
		name := parser.ParseName(child)
		x.SetName(name)
		return x
	}

	panic(sp.UnSupport())
}

func (sep *MySQLExprParser) IsComplexFunction(name string) bool {
	return complexFunctionNameMap[strings.ToUpper(name)]
}

func (sep *MySQLExprParser) ParseComplexFunction(child parser.ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {
	method := sep.SQLExprParser.ParseComplexFunction(child, name)

	if method != name {
		return method
	}

	switch name.(type) {
	case *expr.SQLUnQuotedIdentifier:
		identifier := name.(*expr.SQLUnQuotedIdentifier)
		if identifier.EqualIgnoreCase("CAST") {
			return sep.ParseCastFunction(child, identifier)
		}
	}

	return name
}

func (sep *MySQLExprParser) ParseCastFunction(child parser.ISQLExprParser, name *expr.SQLUnQuotedIdentifier) expr.ISQLExpr {
	if !name.EqualIgnoreCase("CAST") {
		return name
	}

	left := parser.ParseExpr(child)
	sep.AcceptAndNextToken(parser.AS)
	dataType := parser.ParseDataType(sep)

	argument := function.NewCastFunctionArgument(left, dataType)

	method := function.NewMethodInvocation(name, argument)

	return method
}

func (sep *MySQLExprParser) IsNonParametricFunction(name string) bool {
	return nonParametricFunctionNameMap[strings.ToUpper(name)]
}

func (sep *MySQLExprParser) ParseNonParametricFunction(child parser.ISQLExprParser, name expr.ISQLExpr) expr.ISQLExpr {
	return name
}

func (sp *MySQLExprParser) ParseExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	if sp.Accept(parser.EOF) ||
		sp.Accept(parser.SYMB_SEMI) {
		return left
	}

	left = sp.ParseBitXorOperatorExprRest(child, left)
	left = sp.ParseMultiplicativeOperatorExprRest(child, left)
	left = sp.ParseAdditiveOperatorExprRest(child, left)
	left = sp.ParseShiftOperatorExprRest(child, left)
	left = sp.ParseBitAndOperatorExprRest(child, left)
	left = sp.ParseBitOrOperatorExprRest(child, left)
	left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
	left = sp.ParseLowPriorityComparisonOperatorExprRest(child, left)
	left = sp.ParseAndOperatorExprRest(child, left)
	left = sp.ParseXOROperatorExprRest(child, left)
	left = sp.ParseOrOperatorExprRest(child, left)
	return left

}

/**
 * E: T ^ T ^ T
 * T: primaryExpr
 * https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (this *MySQLExprParser) ParseBitXorOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := parser.ParsePrimaryExpr(child)
	return this.ParseBitXorOperatorExprRest(child, left)
}
func (this *MySQLExprParser) ParseBitXorOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch this.Kind() {
	case parser.SYMB_BIT_XOR:
		parser.NextTokenByParser(this)
		right := this.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return this.ParseBitXorOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X ^ X
* OP: *, /, DIV, %, MOD
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *MySQLExprParser) ParseMultiplicativeOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := x.ParseBitXorOperatorExpr(child)
	return x.ParseMultiplicativeOperatorExprRest(child, left)
}

func (x *MySQLExprParser) ParseMultiplicativeOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_STAR:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MULTIPLY, right)
		return x.ParseMultiplicativeOperatorExprRest(child, left)
	case parser.SYMB_SLASH:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.DIVIDE, right)
		return x.ParseMultiplicativeOperatorExprRest(child, left)
	case parser.SYMB_PERCENT:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MODULO, right)
		return x.ParseMultiplicativeOperatorExprRest(child, left)
	case parser.DIV:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.DIV, right)
		return x.ParseMultiplicativeOperatorExprRest(child, left)
	case parser.MOD:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MOD, right)
		return x.ParseMultiplicativeOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T (+/-) T (+/-) T ...
* T: X (*, /, DIV, %, MOD) X
* +, -
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *MySQLExprParser) ParseAdditiveOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := x.ParseMultiplicativeOperatorExpr(child)
	return x.ParseAdditiveOperatorExprRest(child, left)
}

func (x *MySQLExprParser) ParseAdditiveOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_PLUS:
		parser.NextTokenByParser(x)
		right := x.ParseMultiplicativeOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.PLUS, right)
		return x.ParseAdditiveOperatorExprRest(child, left)
	case parser.SYMB_MINUS:
		parser.NextTokenByParser(x)
		right := x.ParseMultiplicativeOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.MINUS, right)
		return x.ParseAdditiveOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T
* T: X +/- X
* op: <<, >>
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *MySQLExprParser) ParseShiftOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := x.ParseAdditiveOperatorExpr(child)
	return x.ParseShiftOperatorExprRest(child, left)
}

func (x *MySQLExprParser) ParseShiftOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_LESS_THAN_LESS_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return x.ParseShiftOperatorExprRest(child, left)
	case parser.SYMB_GREATER_THAN_GREATER_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseAdditiveOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_RIGHT, right)
		return x.ParseShiftOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T & T & T
* T: X (<<,>> ) X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *MySQLExprParser) ParseBitAndOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := x.ParseShiftOperatorExpr(child)
	return x.ParseBitAndOperatorExprRest(child, left)
}

func (x *MySQLExprParser) ParseBitAndOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_BIT_AND:
		parser.NextTokenByParser(x)
		right := x.ParseShiftOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.BIT_AND, right)
		return x.ParseBitAndOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T | T | T
* T: X & X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *MySQLExprParser) ParseBitOrOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := x.ParseBitAndOperatorExpr(child)
	return x.ParseBitOrOperatorExprRest(child, left)
}

func (x *MySQLExprParser) ParseBitOrOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_BIT_OR:
		parser.NextTokenByParser(x)
		right := x.ParseBitAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.BIT_OR, right)
		return x.ParseBitOrOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X | X
* op: = (comparison), <=>, >=, >, <=, <, <>, !=
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *MySQLExprParser) ParseHighPriorityComparisonOperatorExpr(child parser.ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseBitOrOperatorExpr(child)
	return sp.ParseHighPriorityComparisonOperatorExprRest(child, left)
}

func (sp *MySQLExprParser) ParseHighPriorityComparisonOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	if sp.AcceptAndNextToken(parser.SYMB_EQUAL) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.EQ, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(parser.SYMB_LESS_THAN) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.LESS_THAN, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(parser.SYMB_LESS_THAN_EQUAL) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.LESS_THAN_EQ, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(parser.SYMB_GREATER_THAN) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	} else if sp.AcceptAndNextToken(parser.SYMB_GREATER_THAN_EQUAL) {

		right := sp.ParseBitOrOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN_EQ, right)
		left = sp.ParseHighPriorityComparisonOperatorExprRest(child, left)

	}

	return left
}

/**
* E: T op T op T ...
* T: X | X
* op: IS [NOT] NULL, LIKE, [NOT] BETWEEN, [NOT] IN
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (sp *MySQLExprParser) ParseLowPriorityComparisonOperatorExpr(child parser.ISQLExprParser, ) expr.ISQLExpr {
	left := sp.ParseHighPriorityComparisonOperatorExpr(child)
	return sp.ParseLowPriorityComparisonOperatorExprRest(child, left)
}

func (sp *MySQLExprParser) ParseLowPriorityComparisonOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {

	if sp.Accept(parser.IS) {

		return sp.ParseIsConditionRest(child, left)

	} else if sp.AcceptAndNextToken(parser.NOT) {

		if sp.IsParseLikeCondition() {

			return sp.ParseLikeConditionRest(child, true, left)

		} else if sp.Accept(parser.REGEXP) {

			return sp.ParseRegexpConditionRest(child, true, left)

		} else if sp.Accept(parser.BETWEEN) {

			return sp.ParseBetweenConditionRest(child, true, left)

		} else if sp.Accept(parser.IN) {

			return sp.ParseInConditionRest(child, true, left)
		}

	} else if sp.IsParseLikeCondition() {

		return sp.ParseLikeConditionRest(child, false, left)

	} else if sp.Accept(parser.REGEXP) {

		return sp.ParseRegexpConditionRest(child, false, left)

	} else if sp.Accept(parser.BETWEEN) {

		return sp.ParseBetweenConditionRest(child, false, left)

	} else if sp.Accept(parser.IN) {

		return sp.ParseInConditionRest(child, false, left)
	}

	return left
}

/**
* E: T op T op T
* T: X comparisonOperator X
* op: AND, &&
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (sp *MySQLExprParser) ParseAndOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := sp.ParseLowPriorityComparisonOperatorExpr(child)
	return sp.ParseAndOperatorExprRest(child, left)
}

func (x *MySQLExprParser) ParseAndOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.AND:
		parser.NextTokenByParser(x)
		right := x.ParseLowPriorityComparisonOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.AND, right)
		return x.ParseAndOperatorExprRest(child, left)

	case parser.SYMB_LOGICAL_AND:
		parser.NextTokenByParser(x)
		right := x.ParseLowPriorityComparisonOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.LOGICAL_AND, right)
		return x.ParseAndOperatorExprRest(child, left)

	}
	return left
}

/**
* E: T XOR T XOR T
* T: X (AND, &&) X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *MySQLExprParser) ParseXorOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := x.ParseAndOperatorExpr(child)
	return x.ParseXorOperatorExprRest(child, left)
}

func (x *MySQLExprParser) ParseXorOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_LESS_THAN_LESS_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return x.ParseXorOperatorExprRest(child, left)
	case parser.SYMB_GREATER_THAN_GREATER_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseAndOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.SHIFT_RIGHT, right)
		return x.ParseXorOperatorExprRest(child, left)
	}
	return left
}

/**
* E: T op T op T
* T: X XOR X
* op: OR, ||
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *MySQLExprParser) ParseOrOperatorExpr(child parser.ISQLExprParser) expr.ISQLExpr {
	left := x.ParseXorOperatorExpr(child)
	return x.ParseOrOperatorExprRest(child, left)
}

func (this *MySQLExprParser) ParseOrOperatorExprRest(child parser.ISQLExprParser, left expr.ISQLExpr) expr.ISQLExpr {
	switch this.Kind() {
	case parser.OR:
		parser.NextTokenByParser(this)
		right := this.ParseXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.OR, right)
		return this.ParseOrOperatorExprRest(child, left)
	case parser.SYMB_LOGICAL_OR:
		parser.NextTokenByParser(this)
		right := this.ParseXorOperatorExpr(child)
		left = operator.NewBinaryOperator(left, operator.CONCAT, right)
		return this.ParseOrOperatorExprRest(child, left)
	}
	return left
}

// ------------------------------------------------------- DDL ---------------------------------------------------------------------------------------------------
// --------------------------- OnTable --------------------------------------------------
func (sp MySQLExprParser) ParseTableElement(child parser.ISQLExprParser) table.ISQLTableElement {

	if sp.IsParseTableConstraint() {

		return child.ParseTableConstraint(child)

	} else if sp.Accept(parser.LIKE) {

		return parser.ParseLikeClause(child)

	} else if sp.Accept(parser.INDEX) {

		return sp.ParseIndexTableElement(child)

	} else if sp.Accept(parser.KEY) {

		return sp.ParseKeyTableElement(child)

	} else if sp.Accept(parser.FULLTEXT) {

		return parser.ParseLikeClause(child)

	} else if sp.Accept(parser.SPATIAL) {

		return parser.ParseLikeClause(child)

	} else if parser.IsIdentifier(sp.Kind()) {

		return child.ParseTableColumn(child)
	}

	return nil
}

/**
  INDEX [index_name] [index_type] (key_part,...)
      [index_option] ...
 */
func (sp *MySQLExprParser) ParseIndexTableElement(child parser.ISQLExprParser) *table.SQLIndexTableElement {
	if !sp.AcceptAndNextToken(parser.INDEX) {
		return nil
	}
	return nil
}

/**
 *  KEY [index_name] [index_type] (key_part,...) [index_option] ...
 */
func (sp *MySQLExprParser) ParseKeyTableElement(child parser.ISQLExprParser) *table.SQLKeyTableElement {
	if !sp.AcceptAndNextToken(parser.INDEX) {
		return nil
	}
	return nil
}

/**
 * {FULLTEXT | SPATIAL} [INDEX | KEY] [index_name] (key_part,...)
      [index_option] ...
 */
func (sp *MySQLExprParser) ParseFulltextIndexTableElement(child parser.ISQLExprParser) *table.SQLFulltextIndexTableElement {
	if !sp.AcceptAndNextToken(parser.FULLTEXT) {
		return nil
	}
	return nil
}
func (sp *MySQLExprParser) ParseFulltextKeyTableElement(child parser.ISQLExprParser) *table.SQLFulltextKeyTableElement {
	if !sp.AcceptAndNextToken(parser.SPATIAL) {
		return nil
	}

	if sp.AcceptAndNextToken(parser.INDEX) {

	} else if sp.AcceptAndNextToken(parser.KEY) {

	} else {

		panic(sp.SyntaxError())
	}

	return nil
}

func (sp *MySQLExprParser) ParseTableColumnOption(child parser.ISQLExprParser) (expr.ISQLExpr, bool) {
	return nil, false
}
func (sp *MySQLExprParser) ParseColumnConstraint(child parser.ISQLExprParser) (table.ISQLColumnConstraint, bool) {
	x, ok := sp.SQLExprParser.ParseColumnConstraint(child)
	if ok {
		return x, ok
	}

	if sp.Accept(parser.COMMENT) {
		return sp.ParseCommentExpr(child), true
	}

	if sp.Accept(parser.DEFAULT) {
		return sp.ParseDefaultClause(child), true
	}

	if sp.AcceptAndNextToken(parser.AUTO_INCREMENT) {
		return table.NewAutoIncrementExpr(), true
	}

	if sp.AcceptAndNextToken(parser.VISIBLE) {
		return table.NewVisibleExpr(), true
	}

	if sp.AcceptAndNextToken(parser.VISIBLE) {
		return table.NewInvisibleExpr(), true
	}

	return x, ok
}

func (sp *MySQLExprParser) ParseTableConstraint(child parser.ISQLExprParser) table.ISQLTableConstraint {
	x := sp.SQLExprParser.ParseTableConstraint(child)
	if x != nil {
		return x
	}

	return nil
}

func (sp *MySQLExprParser) ParseTableConstraintRest(child parser.ISQLExprParser, name expr.ISQLName) (table.ISQLTableConstraint, bool) {
	if sp.AcceptAndNextToken(parser.INDEX) {
		x := table.NewUniqueIndexTableConstraint()
		parser.ParseName(child)

		if sp.AcceptAndNextToken(parser.USING) {
			if sp.AcceptAndNextToken(parser.BTREE) {

			} else if sp.AcceptAndNextToken(parser.HASH) {

			} else {
				panic(sp.UnSupport())
			}
		}

		sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
		for {

			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

		return x, true

	} else if sp.AcceptAndNextToken(parser.KEY) {
		x := table.NewUniqueIndexTableConstraint()
		parser.ParseName(child)
		if sp.AcceptAndNextToken(parser.USING) {
			if sp.AcceptAndNextToken(parser.BTREE) {

			} else if sp.AcceptAndNextToken(parser.HASH) {

			} else {
				panic(sp.UnSupport())
			}
		}

		return x, true

	} else if sp.AcceptAndNextToken(parser.FULLTEXT) {

		x := table.NewUniqueIndexTableConstraint()
		return x, true

	} else if sp.AcceptAndNextToken(parser.SPATIAL) {

		x := table.NewUniqueIndexTableConstraint()
		return x, true

	}
	return nil, false
}

func (sp *MySQLExprParser) ParsePartitionBy(child parser.ISQLExprParser) table.ISQLPartitionBy {
	if !child.AcceptAndNextToken(parser.PARTITION) {
		return nil
	}

	child.AcceptAndNextTokenWithError(parser.BY, true)

	linear := child.AcceptAndNextToken(parser.LINEAR)

	if sp.AcceptAndNextToken(parser.HASH) {
		x := table.NewPartitionByHash()
		x.Linear = linear

		sp.ParsePartitionByRest(child, x)

		return x

	} else if sp.AcceptAndNextToken(parser.KEY) {
		x := table.NewPartitionByKey()
		x.Linear = linear

		if sp.AcceptAndNextToken(parser.ALGORITHM) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
			value := parser.ParseExpr(child)
			x.SetAlgorithmValue(value)
		}

		sp.ParsePartitionByRest(child, x)
		return x

	} else if sp.AcceptAndNextToken(parser.RANGE) {
		x := table.NewPartitionByRange()

		columns := sp.AcceptAndNextToken(parser.COLUMNS)
		x.HasColumns = columns

		sp.ParsePartitionByRest(child, x)

		return x

	} else if sp.AcceptAndNextToken(parser.LIST) {
		x := table.NewPartitionByList()

		columns := sp.AcceptAndNextToken(parser.COLUMNS)
		x.HasColumns = columns

		sp.ParsePartitionByRest(child, x)
		return x
	}

	panic(child.SyntaxError())
}

// func (sp *TiDBExprParser) ParsePartitionByRest(child parser.ISQLExprParser, x table.ISQLPartitionBy) {
// }

func (sp *MySQLExprParser) ParseSubPartitionBy(child parser.ISQLExprParser) table.ISQLSubPartitionBy {
	if !sp.AcceptAndNextToken(parser.SUBPARTITION) {
		return nil
	}

	sp.AcceptAndNextTokenWithError(parser.BY, true)

	linear := sp.AcceptAndNextToken(parser.LINEAR)

	if sp.AcceptAndNextToken(parser.HASH) {
		x := table.NewSubPartitionByHash()
		x.Linear = linear

		sp.ParseSubPartitionByRest(child, x)

		return x

	} else if sp.AcceptAndNextToken(parser.KEY) {
		x := table.NewSubPartitionByKey()
		x.Linear = linear

		if sp.AcceptAndNextToken(parser.ALGORITHM) {
			sp.AcceptAndNextTokenWithError(parser.SYMB_EQUAL, true)
			value := parser.ParseExpr(child)
			x.SetAlgorithmValue(value)
		}

		sp.ParseSubPartitionByRest(child, x)

		return x

	}

	panic(sp.UnSupport())
}

/**
 * PARTITION partition_name
        [VALUES
            {LESS THAN {(expr | value_list) | MAXVALUE}
            |
            IN (value_list)}]
        [[STORAGE] ENGINE [=] engine_name]
        [COMMENT [=] 'string' ]
        [DATA DIRECTORY [=] 'data_dir']
        [INDEX DIRECTORY [=] 'index_dir']
        [MAX_ROWS [=] max_number_of_rows]
        [MIN_ROWS [=] min_number_of_rows]
        [TABLESPACE [=] tablespace_name]
        [(subpartition_definition [, subpartition_definition] ...)]
 */
func (sp *MySQLExprParser) ParsePartitionDefinition(child parser.ISQLExprParser) *table.SQLPartitionDefinition {
	if !sp.AcceptAndNextToken(parser.PARTITION) {
		return nil
	}

	x := table.NewPartitionDefinition()
	name := parser.ParseName(child)
	x.SetName(name)

	values := sp.ParsePartitionValues(child)
	x.SetValues(values)

	for {
		option, ok := child.ParsePartitionDefinitionOption(child)
		if !ok {
			break
		}
		x.AddOption(option)
	}

	if sp.AcceptAndNextToken(parser.SYMB_LEFT_PAREN) {
		for {
			subPartitionDefinition := sp.ParseSubPartitionDefinition(child)
			x.AddSubpartitionDefinition(subPartitionDefinition)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_BRACE, true)
	}

	return x
}

func (sp *MySQLExprParser) ParsePartitionDefinitionOption(child parser.ISQLExprParser) (expr.ISQLExpr, bool) {
	if sp.Accept(parser.STORAGE) {
		sp.AcceptAndNextToken(parser.STORAGE)
		sp.AcceptAndNextTokenWithError(parser.ENGINE, true)

		sp.AcceptAndNextToken(parser.SYMB_EQUAL)

		parser.ParseExpr(child)

		return nil, true

	} else if sp.Accept(parser.DATA) {

		sp.AcceptAndNextToken(parser.DATA)
		sp.AcceptAndNextTokenWithError(parser.DIRECTORY, true)

		sp.AcceptAndNextToken(parser.SYMB_EQUAL)

		parser.ParseExpr(child)

		return nil, true

	} else if sp.Accept(parser.INDEX) {

		sp.AcceptAndNextToken(parser.INDEX)
		sp.AcceptAndNextTokenWithError(parser.DIRECTORY, true)

		sp.AcceptAndNextToken(parser.SYMB_EQUAL)

		parser.ParseExpr(child)

	} else if sp.Accept(parser.ENGINE) ||
		sp.Accept(parser.COMMENT) ||
		sp.Accept(parser.ENGINE) ||
		sp.Accept(parser.MAX_ROWS) || sp.Accept(parser.MIN_ROWS) || sp.Accept(parser.TABLESPACE) {

		return parser.ParseAssignExpr(child), true
	}
	return nil, false
}

/**
 * SUBPARTITION logical_name
        [[STORAGE] ENGINE [=] engine_name]
        [COMMENT [=] 'string' ]
        [DATA DIRECTORY [=] 'data_dir']
        [INDEX DIRECTORY [=] 'index_dir']
        [MAX_ROWS [=] max_number_of_rows]
        [MIN_ROWS [=] min_number_of_rows]
        [TABLESPACE [=] tablespace_name]
 */
func (sp *MySQLExprParser) ParseSubPartitionDefinition(child parser.ISQLExprParser) *table.SQLSubPartitionDefinition {
	if !sp.AcceptAndNextToken(parser.SUBPARTITION) {
		return nil
	}

	x := table.NewSubPartitionDefinition()
	name := parser.ParseName(child)
	x.SetName(name)

	return x
}

func (sp *MySQLExprParser) ParseAlterTableAction(child parser.ISQLExprParser) table.ISQLAlterTableAction {
	x := sp.SQLExprParser.ParseAlterTableAction(child)
	if x != nil {
		return x
	}

	if sp.AcceptAndNextToken(parser.CHANGE) {

	} else if sp.Accept(parser.RENAME) {

	} else if sp.Accept(parser.DISCARD) {

	} else if sp.Accept(parser.IMPORT) {

	} else if sp.Accept(parser.TRUNCATE) {

	} else if sp.Accept(parser.COALESCE) {

	} else if sp.Accept(parser.EXCHANGE) {

	} else if sp.AcceptAndNextToken(parser.ANALYZE) {
		sp.AcceptAndNextTokenWithError(parser.PARTITION, true)

		for {
			parser.ParseExpr(child)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}

	} else if sp.Accept(parser.CHECK) {

	} else if sp.Accept(parser.OPTIMIZE) {

	} else if sp.Accept(parser.REBUILD) {

	} else if sp.Accept(parser.REPAIR) {

	} else if sp.Accept(parser.REMOVE) {

	} else if sp.Accept(parser.ALGORITHM) {
		return parser.ParseAssignExpr(child)

	} else if sp.Accept(parser.DEFAULT) {

	}

	return nil
}
func (sp *MySQLExprParser) ParseAddAlterTableAction(child parser.ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(parser.ADD) {
		return nil
	}

	hasConstraint := sp.AcceptAndNextToken(parser.CONSTRAINT)
	var name expr.ISQLName
	if hasConstraint {
		name = parser.ParseName(child)
	}

	if sp.Accept(parser.INDEX) {

		return nil

	} else if sp.Accept(parser.KEY) {
		x := table.NewAddKeyAlterTableAction()

		return x

	} else if sp.IsParseTableConstraint() {

		x := table.NewAddTableConstraintAlterTableAction()

		tableConstraint := sp.ParseTableConstraint(child)
		if tableConstraint == nil {
			panic(sp.SyntaxError())
		}
		tableConstraint.SetName(name)

		x.SetTableConstraint(tableConstraint)

	} else if sp.Accept(parser.UNIQUE) {

	} else if sp.Accept(parser.COLUMN) ||
		parser.IsIdentifier(sp.Kind()) {
		return sp.ParseAddColumnAlterTableAction(child)

	} else if sp.AcceptAndNextToken(parser.PARTITION) {
		x := table.NewAddPartitionAlterTableAction()
		sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)
		partition := child.ParsePartitionDefinition(child)
		x.SetPartition(partition)
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
		return x
	}

	panic(sp.UnSupport())
}
func (sp *MySQLExprParser) ParseAddColumnAlterTableAction(child parser.ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.Accept(parser.COLUMN) && !parser.IsIdentifier(sp.Kind()) {
		return nil
	}

	x := table.NewAddColumnAlterTableAction()

	hasColumn := sp.AcceptAndNextToken(parser.COLUMN)
	x.HasColumn = hasColumn

	paren := sp.AcceptAndNextToken(parser.SYMB_LEFT_PAREN)

	column := child.ParseTableColumn(child)
	x.AddColumn(column)

	for paren {
		if !sp.AcceptAndNextToken(parser.SYMB_COMMA) || column == nil {
			break
		}
		column := child.ParseTableColumn(child)
		x.AddColumn(column)
	}

	if paren {
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
	}
	return x
}

func (sp *MySQLExprParser) ParseDropAlterTableAction(child parser.ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}

	if sp.AcceptAndNextToken(parser.CONSTRAINT) {
		x := table.NewDropTableConstraintAlterTableAction()

		name := parser.ParseName(child)
		x.SetName(name)

		return x

	} else if sp.AcceptAndNextToken(parser.INDEX) {
		x := table.NewDropIndexAlterTableAction()

		name := parser.ParseName(child)
		x.SetName(name)

		return x

	} else if sp.AcceptAndNextToken(parser.KEY) {

		x := table.NewDropKeyAlterTableAction()

		name := parser.ParseName(child)
		x.SetName(name)

		return x

	} else if sp.AcceptAndNextToken(parser.PRIMARY) {
		sp.AcceptAndNextToken(parser.KEY)

		return table.NewDropPrimaryKeyTableConstraintAlterTableAction()

	} else if sp.AcceptAndNextToken(parser.FOREIGN) {
		sp.AcceptAndNextToken(parser.KEY)

		x := table.NewDropForeignKeyTableConstraintAlterTableAction()

		name := parser.ParseName(child)
		x.SetName(name)

		return x

	} else if sp.AcceptAndNextToken(parser.CHECK) {
		x := table.NewDropCheckTableConstraintAlterTableAction()

		name := parser.ParseName(child)
		x.SetName(name)

		return x

	} else if sp.AcceptAndNextToken(parser.PARTITION) {
		for {
			parser.ParseName(child)
			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}

	} else if sp.Accept(parser.COLUMN) || parser.IsIdentifier(sp.Kind()) {
		x := table.NewDropColumnAlterTableAction()

		hasColumn := sp.AcceptAndNextToken(parser.COLUMN)
		x.HasColumn = hasColumn

		column := parser.ParseName(child)
		x.SetColumn(column)
		return x
	}

	panic(sp.UnSupport())
}

// ---------------------- VIEW

func (sp *MySQLExprParser) ParseViewColumn(child parser.ISQLExprParser) *view.SQLViewColumn {
	x := view.NewViewColumn()
	name := parser.ParseName(child)
	x.SetName(name)

	dataType := parser.ParseDataType(child)
	x.SetDataType(dataType)

	for {
		option, ok := child.ParseViewColumnOption(child)
		if !ok {
			break
		}
		x.AddOption(option)
	}

	return x
}
func (sp *MySQLExprParser) ParseViewColumnOption(child parser.ISQLExprParser) (expr.ISQLExpr, bool) {
	if child.IsParseColumnConstraint() {

		return child.ParseColumnConstraint(child)
	}

	return nil, false
}

// ------------------------------------------------------- DML ---------------------------------------------------------------------------------------------------
// ---------------------- SELECT
/**
 WITH [RECURSIVE]
        cte_name [(col_name [, col_name] ...)] AS (subquery)
        [, cte_name [(col_name [, col_name] ...)] AS (subquery)] ...
 * https://dev.mysql.com/doc/refman/8.0/en/with.html
 */
func (x *MySQLExprParser) ParseWithClause(child parser.ISQLExprParser) select_.ISQLWithClause {

	if !x.AcceptAndNextToken(parser.WITH) {
		return nil
	}

	withClause := select_.NewWithClause()

	recursive := x.AcceptAndNextToken(parser.RECURSIVE)
	withClause.Recursive = recursive

	for {

		withClause.AddFactoringClause(x.ParseFactoringClause())

		if !x.AcceptAndNextToken(parser.SYMB_COMMA) {
			break
		}
	}

	return withClause
}

func (sp *MySQLExprParser) ParseFactoringClause() select_.ISQLFactoringClause {
	x := select_.NewSubQueryFactoringClause()

	name := sp.ParseIdentifier(sp)
	x.SetName(name)

	if sp.AcceptAndNextToken(parser.SYMB_LEFT_PAREN) {

		for {

			column := sp.ParseIdentifier(sp)
			x.AddColumn(column)

			if !sp.AcceptAndNextToken(parser.SYMB_COMMA) {
				break
			}
		}

		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
	}

	sp.AcceptAndNextTokenWithError(parser.AS, true)

	sp.AcceptAndNextTokenWithError(parser.SYMB_LEFT_PAREN, true)

	subQuery := parser.ParseSelectQuery(sp)
	x.SetSubQuery(subQuery)

	sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)

	return x
}

func (sp MySQLExprParser) ParseLockClause(child parser.ISQLExprParser) select_.ISQLLockClause {
	if sp.Accept(parser.FOR) {

		return sp.ParseForUpdate(child)

	} else if sp.Accept(parser.LOCK) {

		return sp.ParseLockInShareModeClause(child)
	}
	return nil
}

/**
* FOR { UPDATE | SHARE } [ OF table_name [, ...] ] [ NOWAIT | SKIP LOCKED ] [...]
*/
func (sp *MySQLExprParser) ParseForUpdate(child parser.ISQLExprParser) select_.ISQLLockClause {
	if !sp.AcceptAndNextToken(parser.FOR) {
		return nil
	}

	if sp.AcceptAndNextToken(parser.UPDATE) {

		x := select_.NewForUpdateClause()
		sp.ParseForUpdateRest(child, x.AbstractSQLLockForClause)
		return x

	} else if sp.AcceptAndNextToken(parser.SHARE) {

		x := select_.NewForShareClause()
		sp.ParseForUpdateRest(child, x.AbstractSQLLockForClause)
		return x

	} else {

		panic(sp.UnSupport())
	}

}
