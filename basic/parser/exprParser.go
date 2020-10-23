package parser

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/expr"
	"gumihoy.com/sql/basic/ast/expr/datatype"
	"gumihoy.com/sql/basic/ast/expr/literal"
	"gumihoy.com/sql/basic/ast/expr/operator"
	"gumihoy.com/sql/basic/ast/expr/select"
	"gumihoy.com/sql/basic/ast/expr/variable"
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/basic/ast/statement/table"
)

type ISQLExprParser interface {
	ISQLParser

	CreateSQLSelectStatementParser() ISQLSelectStatementParser

	ParseIdentifier() expr.ISQLIdentifier
	ParseNameWithOwner(owner expr.ISQLName) expr.ISQLName

	ParseExpr() expr.ISQLExpr
	ParsePrimaryExpr() expr.ISQLExpr
	ParseExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr
}

type SQLExprParser struct {
	*SQLParser
}

func NewExprParserBySQL(sql string) *SQLExprParser {
	return NewExprParserByLexer(NewLexer(sql))
}

func NewExprParserByLexer(lexer ISQLLexer) *SQLExprParser {
	return &SQLExprParser{&SQLParser{lexer}}
}

func (x *SQLExprParser) ParseStatements() []ast.ISQLObject {
	return x.ParseStatementsWithParent(nil)
}

func (x *SQLExprParser) ParseStatementsWithParent(parent ast.ISQLObject) []ast.ISQLObject {
	var stmts []ast.ISQLObject
	for {
		if x.Accept(EOF) {
			break
		}

		if x.AcceptAndNextToken(SYMB_SEMI) {
			if len(stmts) > 0 {
				lastStmt := stmts[len(stmts)-1]
				lastStmt.SetAfterSemi(true)
			}
			continue
		}

		if x.AcceptAndNextToken(SYMB_SLASH) {
			continue
		}

		stmt := ParseStatement(x)
		if stmt == nil {
			break
		}
		stmt.SetParent(parent)
		stmts = append(stmts, stmt)
	}
	return stmts

}

func ParseStatement(x ISQLExprParser) ast.ISQLObject {

	// DDL
	// if x.Accept(ALTER) {
	// 	return parseAlterStatement()
	// }
	// if x.Accept(ANALYZE) {
	// 	return parseAnalyzeStatement()
	// }
	// if x.Accept(ASSOCIATE) {
	// 	panic("UnSupport .")
	// 	return nil
	// }
	// if x.Accept(AUDIT) {
	// 	panic("UnSupport .")
	// 	return nil
	// }
	// if x.Accept(COMMENT) {
	// 	return parseCommentStatement()
	// }
	// if x.Accept(CREATE)) {
	// 	return parseCreateStatement()
	// }
	// if x.Accept(DROP)) {
	// 	return parseDropStatement()
	// }
	// if x.Accept(RENAME)) {
	// 	return null
	// }

	// DML
	// if x.Accept(CALL)) {
	// 	return parseCallStatement()
	// }
	// if x.Accept(DELETE)) {
	// 	return parseDeleteStatement()
	// }
	// if x.Accept(EXPLAIN)) {
	// 	return parseExplainStatement()
	// }
	// if x.Accept(INSERT)) {
	// 	return parseInsertStatement()
	// }
	// if x.Accept(LOCK)) {
	// 	return parseLockStatement()
	// }
	// if x.Accept(WITH)) {
	// 	return parseWithStatement()
	// }
	if x.Accept(SELECT) {
		return ParseSelectStatement(x)
	}
	// if x.Accept(LPAREN)) {
	// 	return parseLParenStatement()
	// }
	// if x.Accept(UPDATE)) {
	// 	return parseUpdateStatement()
	// }

	// FCL
	// if x.Accept(CASE)) {
	// 	return parseCaseStatement()
	// }
	// if x.Accept(CLOSE)) {
	// 	return parseCloseStatement()
	// }
	// if x.Accept(CONTINUE)) {
	// 	return parseContinueStatement()
	// }
	// if x.Accept(EXIT)) {
	// 	return parseExitStatement()
	// }
	// if x.Accept(FETCH)) {
	// 	return parseFetchStatement()
	// }
	// if x.Accept(SQLToken.TokenKind.FOR)) {
	// 	return parseForLoopStatement()
	// }
	// if x.Accept(GOTO)) {
	// 	return parseGotoStatement()
	// }
	// if x.Accept(IF)) {
	// 	return parseIfStatement()
	// }
	// if x.Accept(LOOP)) {
	// 	return parseLoopStatement()
	// }
	// if x.Accept(NULL)) {
	// 	return parseNullStatement()
	// }
	// if x.Accept(OPEN)) {
	// 	return parseOpenStatement()
	// }
	// if x.Accept(PIPE)) {
	// 	return parsePipeRowStatement()
	// }
	// if x.Accept(RAISE)) {
	// 	return parseRaiseStatement()
	// }
	// if x.Accept(REPEAT)) {
	// 	return parseRepeatStatement()
	// }
	// if x.Accept(RETURN)) {
	// 	return parseReturnStatement()
	// }
	// if x.Accept(WHILE)) {
	// 	return parseWhileLoopStatement()
	// }

	// TC
	// if x.Accept(COMMIT) {
	// 	return parseCommitStatement()
	// }
	// if x.Accept(ROLLBACK) {
	// 	return parseRollbackStatement()
	// }
	// if x.Accept(SAVEPOINT) {
	// 	return null
	// }

	// SC
	// if x.Accept(SET) {
	// 	return parseSetStatement()
	// }

	// make := x.make()
	// expr := x.ParseExpr()
	// if x.Accept(SQLToken.TokenKind.ASSIGN) {
	// 	// this.reset(make)
	// 	return parseAssignmentStatement()
	// }
	//
	//
	// if x.Accept(LOOP) {
	// 	// x.lexer.Reset()
	// 	return parseLoopStatement()
	//
	// }
	// if x.Accept(WHILE) {
	// 	// this.reset(make)
	// 	return parseWhileStatement()
	// }
	// this.reset(make)
	return nil
}

func ParseSelectStatement(x ISQLExprParser) statement.ISQLStatement {
	parser := x.CreateSQLSelectStatementParser()
	return parser.Parse()
}

func (x *SQLExprParser) CreateSQLSelectStatementParser() ISQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(x)
}


/**
 * E: T.T.T
 * T: ID
 */
func ParseName(x ISQLExprParser) expr.ISQLName {
	owner := x.ParseIdentifier()
	return x.ParseNameWithOwner(owner)
}

func (x *SQLExprParser) ParseIdentifier() expr.ISQLIdentifier {
	var name expr.ISQLIdentifier
	switch x.Token().Kind {
	case IDENTIFIER:
		name = expr.NewUnQuotedIdentifier(x.StringValue())
		NextTokenByParser(x)
		break
	case IDENTIFIER_DOUBLE_QUOTE:
		name = expr.NewDoubleQuotedIdentifier(x.StringValue())
		NextTokenByParser(x)
		break
	case IDENTIFIER_REVERSE_QUOTE:
		name = expr.NewReverseQuotedIdentifier(x.StringValue())
		NextTokenByParser(x)
		break
		// case SYMB_STAR:
		// 	name = expr.NewAllColumnExpr()
		// 	break
	}

	return name
}

func (x *SQLExprParser) ParseNameWithOwner(owner expr.ISQLName) expr.ISQLName {
	switch x.Token().Kind {
	case SYMB_DOT:
		NextTokenByParser(x)
		right := x.ParseIdentifier()
		name := expr.NewNameWithOwnerAndName(owner, right)
		newName := x.ParseNameWithOwner(name)
		return newName
	}
	return owner
}

/**
 * expr
 */
func (x *SQLExprParser) ParseExpr() expr.ISQLExpr {
	left := x.ParsePrimaryExpr()
	if left == nil {
		return left
	}
	return x.ParseExprWithLeft(left)
}

func (p *SQLExprParser) ParsePrimaryExpr() expr.ISQLExpr {
	if p.Kind() == SYMB_SEMI ||
		p.Kind() == EOF {
		return nil
	}

	var x expr.ISQLExpr

	if p.Accept(IDENTIFIER) {
		x = expr.NewUnQuotedIdentifier(p.StringValue())
		NextTokenByParser(p)
		return x
	}
	if p.Accept(IDENTIFIER_DOUBLE_QUOTE) {
		x = expr.NewDoubleQuotedIdentifier(p.StringValue())
		NextTokenByParser(p)
		return x
	}
	if p.Accept(IDENTIFIER_REVERSE_QUOTE) {
		x = expr.NewReverseQuotedIdentifier(p.StringValue())
		NextTokenByParser(p)
		return x
	}

	if p.Accept(LITERAL_STRING) {
		x = literal.NewStringLiteral(p.StringValue())
		NextTokenByParser(p)
		return x
	}
	if p.Accept(LITERAL_INTEGER) {
		x = literal.NewIntegerLiteralWithString(p.StringValue())
		NextTokenByParser(p)
		return x
	}
	if p.Accept(LITERAL_FLOATING_POINT) {
		x = literal.NewFloatingPointLiteralWith(p.StringValue())
		NextTokenByParser(p)
		return x
	}
	if p.Accept(DATE) {
		NextTokenByParser(p)
		x = literal.NewDateLiteral(p.ParseExpr())
		return x
	}
	if p.Accept(TIME) {
		NextTokenByParser(p)
		x = literal.NewTimeLiteral(p.ParseExpr())
		return x
	}
	if p.Accept(TIMESTAMP) {
		NextTokenByParser(p)
		x = literal.NewTimestampLiteral(p.ParseExpr())
		return x
	}

	if p.Kind() == LITERAL_BOOLEAN_TRUE {
		x = literal.NewBooleanLiteral(true)
		NextTokenByParser(p)
		return x
	}

	if p.Kind() == LITERAL_BOOLEAN_FALSE {
		x = literal.NewBooleanLiteral(false)
		NextTokenByParser(p)
		return x
	}

	if p.Accept(SYMB_STAR) {
		x = expr.NewAllColumnExpr()
		NextTokenByParser(p)
		return x
	}
	if p.Accept(SYMB_QUESTION) {
		x = variable.NewVariableExpr()
		NextTokenByParser(p)
		return x
	}

	if p.Accept(SYMB_LEFT_PAREN) {
		NextTokenByParser(p)
		x = p.ParseExpr()

		if p.Accept(SYMB_COMMA) {
			listExpr := expr.NewListExpr()
			listExpr.AddElement(x)
			for {
				NextTokenByParser(p)
				listExpr.AddElement(p.ParseExpr())
				if !p.Accept(SYMB_COMMA) {
					break
				}
			}
			x = listExpr
		}

		switch x.(type) {
		case *operator.SQLUnaryOperatorExpr:
			x.(*operator.SQLUnaryOperatorExpr).Paren = true

		case *operator.SQLBinaryOperatorExpr:
			x.(*operator.SQLBinaryOperatorExpr).Paren = true

		case *expr.SQLListExpr:
		default:
			listExpr := expr.NewListExpr()
			listExpr.AddElement(x)
			x = listExpr
		}

		p.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		return x
	}

	return x
}

func (x *SQLExprParser) ParseExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	if x.Token().Kind == EOF || x.Token().Kind == SYMB_SEMI {
		return left
	}

	left = x.ParseBitXorOperatorExprWithLeft(left)
	left = x.ParseMultiplicativeOperatorExprWithLeft(left)
	left = x.ParseAdditiveOperatorExprWithLeft(left)
	left = x.ParseShiftOperatorExprWithLeft(left)
	left = x.ParseBitAndOperatorExprWithLeft(left)
	left = x.ParseBitOrOperatorExprWithLeft(left)
	left = x.ParseComparisonOperatorExprWithLeft(left)
	left = x.ParseAndOperatorExprWithLeft(left)
	left = x.ParseXorOperatorExprWithLeft(left)
	left = x.ParseOrOperatorExprWithLeft(left)
	return left

}

/**
 * E: T ^ T ^ T
 * T: primaryExpr
 * https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *SQLExprParser) ParseBitXorOperatorExpr() expr.ISQLExpr {
	left := x.ParsePrimaryExpr()
	return x.ParseBitXorOperatorExprWithLeft(left)
}
func (x *SQLExprParser) ParseBitXorOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_BIT_XOR:
		NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseBitXorOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X ^ X
* OP: *, /, DIV, %, MOD
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseMultiplicativeOperatorExpr() expr.ISQLExpr {
	left := x.ParseBitXorOperatorExpr()
	return x.ParseMultiplicativeOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseMultiplicativeOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_STAR:
		NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case SYMB_SLASH:
		NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case SYMB_PERCENT:
		NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case DIV:
		NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case MOD:
		NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T (+/-) T (+/-) T ...
* T: X (*, /, DIV, %, MOD) X
* +, -
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseAdditiveOperatorExpr() expr.ISQLExpr {
	left := x.ParseMultiplicativeOperatorExpr()
	return x.ParseAdditiveOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseAdditiveOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_PLUS:
		NextTokenByParser(x)
		right := x.ParseMultiplicativeOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.PLUS, right)
		return x.ParseAdditiveOperatorExprWithLeft(left)
	case SYMB_MINUS:
		NextTokenByParser(x)
		right := x.ParseMultiplicativeOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.MINUS, right)
		return x.ParseAdditiveOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T op T op T
* T: X +/- X
* op: <<, >>
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseShiftOperatorExpr() expr.ISQLExpr {
	left := x.ParseAdditiveOperatorExpr()
	return x.ParseShiftOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseShiftOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_LESS_THAN_LESS_THAN:
		NextTokenByParser(x)
		right := x.ParseAdditiveOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return x.ParseShiftOperatorExprWithLeft(left)
	case SYMB_GREATER_THAN_GREATER_THAN:
		NextTokenByParser(x)
		right := x.ParseAdditiveOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.SHIFT_RIGHT, right)
		return x.ParseShiftOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T & T & T
* T: X (<<,>> ) X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseBitAndOperatorExpr() expr.ISQLExpr {
	left := x.ParseShiftOperatorExpr()
	return x.ParseBitAndOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseBitAndOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_BIT_AND:
		NextTokenByParser(x)
		right := x.ParseShiftOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_AND, right)
		return x.ParseBitAndOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T | T | T
* T: X & X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseBitOrOperatorExpr() expr.ISQLExpr {
	left := x.ParseBitAndOperatorExpr()
	return x.ParseBitOrOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseBitOrOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_BIT_OR:
		NextTokenByParser(x)
		right := x.ParseBitAndOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_OR, right)
		return x.ParseBitOrOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T op T op T ...
* T: X | X
* op: = (comparison), <=>, >=, >, <=, <, <>, !=, IS, [not] LIKE, REGEXP, IN,
* [not] BETWEEN, CASE, WHEN, THEN, ELSE
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseComparisonOperatorExpr() expr.ISQLExpr {
	left := x.ParseBitOrOperatorExpr()
	return x.ParseComparisonOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseComparisonOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_EQUAL:
		NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.EQ, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case SYMB_LESS_THAN:
		NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.LESS_THAN, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case SYMB_LESS_THAN_EQUAL:
		NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.LESS_THAN_EQ, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case SYMB_GREATER_THAN:
		NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case SYMB_GREATER_THAN_EQUAL:
		NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN_EQ, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T op T op T
* T: X comparisonOperator X
* op: AND, &&
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseAndOperatorExpr() expr.ISQLExpr {
	left := x.ParseComparisonOperatorExpr()
	return x.ParseAndOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseAndOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case AND:
		NextTokenByParser(x)
		right := x.ParseComparisonOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.AND, right)
		return x.ParseAndOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T XOR T XOR T
* T: X (AND, &&) X
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseXorOperatorExpr() expr.ISQLExpr {
	left := x.ParseAndOperatorExpr()
	return x.ParseXorOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseXorOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case SYMB_LESS_THAN_LESS_THAN:
		NextTokenByParser(x)
		right := x.ParseAndOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return x.ParseXorOperatorExprWithLeft(left)
	case SYMB_GREATER_THAN_GREATER_THAN:
		NextTokenByParser(x)
		right := x.ParseAndOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.SHIFT_RIGHT, right)
		return x.ParseXorOperatorExprWithLeft(left)
	}
	return left
}

/**
* E: T op T op T
* T: X XOR X
* op: OR, ||
* https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
*/
func (x *SQLExprParser) ParseOrOperatorExpr() expr.ISQLExpr {
	left := x.ParseXorOperatorExpr()
	return x.ParseOrOperatorExprWithLeft(left)
}

func (x *SQLExprParser) ParseOrOperatorExprWithLeft(left expr.ISQLExpr) expr.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case OR:
		NextTokenByParser(x)
		right := x.ParseXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.OR, right)
		return x.ParseOrOperatorExprWithLeft(left)
	case SYMB_LOGICAL_OR:
		NextTokenByParser(x)
		right := x.ParseXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.CONCAT, right)
		return x.ParseOrOperatorExprWithLeft(left)
	}
	return left
}

/**
* (
* 	TableElements
* )
*/
func ParseTableElements(parser ISQLExprParser, tableElements []table.ISQLTableElement, parent ast.ISQLObject) {
	accept := parser.AcceptAndNextToken(SYMB_LEFT_PAREN)
	if !accept {
		return
	}

	for {

		tableElement := ParseTableElement(parser)
		//            tableElement.addBeforeComment()
		if tableElement != nil {
			tableElement.SetParent(parent)
		}
		tableElements = append(tableElements, tableElement)

		if !parser.Accept(SYMB_COMMA) {
			break
		}

		NextTokenByParser(parser)
	}

	parser.AcceptWithError(SYMB_RIGHT_PAREN, true)
}

func ParseTableElement(parser ISQLExprParser) table.ISQLTableElement {

	var x table.ISQLTableElement = nil

	if isTableConstraint() {
		x = parseTableConstraint(parser)

	} else if parser.Accept(LIKE) {
		x = parseLikeClause(parser)

	} else if IsIdentifier(parser.Kind()) {
		x = ParseColumn(parser)
	}

	return x
}

func parseTableConstraint(parser ISQLExprParser) table.ISQLTableConstraint {
	return nil
}

func parseLikeClause(parser ISQLExprParser) *table.SQLTableLikeClause {
	return nil
}

/**
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#column%20definition
 */
func ParseColumn(parser ISQLExprParser) *table.SQLTableColumn {

	if !IsIdentifier(parser.Kind()) {
		return nil
	}

	x := table.NewColumn()

	name := ParseName(parser)
	x.SetName(name)

	dataType := parseDataType(parser)
	x.SetDataType(dataType)

	if parser.AcceptAndNextToken(SORT) {
		// x.setSort(true)
	}

	// SQLCollateOptionExpr collateOptionExpr = parseCollateClause()
	// x.setCollateClause(collateOptionExpr)

	// SQLVisibleType visibleType = parseVisibleType()
	// x.setVisible(visibleType)

	// x.setDefaultClause(parseDefaultClause())

	return x
}

func parseDataType(parser ISQLExprParser) datatype.ISQLDataType {
	return nil
}

func isTableConstraint() bool {
	return false
}

/**
 *  Select
 */
func ParseSelectQuery(parser ISQLExprParser) select_.ISQLSelectQuery {
	x := ParseQueryBlock(parser)
	if x == nil {
		return nil
	}

	orderByClause := ParseOrderBy(parser)
	x.SetOrderByClause(orderByClause)

	limitClause := ParseLimit(parser)
	x.SetLimitClause(limitClause)

	lockClause := ParseLockClause(parser)
	x.SetLockClause(lockClause)

	return x
}

func ParseQueryBlock(parser ISQLExprParser) select_.ISQLSelectQuery {
	x := parsePrimaryQueryBlock(parser)
	x = parseUnionQueryBlock(parser, x)
	return x
}

func parsePrimaryQueryBlock(parser ISQLExprParser) select_.ISQLSelectQuery {
	if parser.AcceptAndNextToken(SYMB_LEFT_PAREN) {

		selectQuery := ParseSelectQuery(parser)
		query := select_.NewParenSelectQuery(selectQuery)
		parser.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		return query
	}

	// withClause := parseWith()
	if !parser.AcceptAndNextToken(SELECT) {
		return nil
	}

	query := select_.NewSelectQuery()
	// query.setWithClause(withClause)

	// SQLSetQuantifier setQuantifier := parseSetQuantifier()
	// query.setSetQuantifier(setQuantifier)

	ParseSelectElements(parser, query)

	ParseSelectTargetElements(parser, query)

	fromClause := parseFrom(parser)
	query.SetFromClause(fromClause)

	whereClause := parseWhere(parser)
	query.SetWhereClause(whereClause)

	groupByClause := parseGroupBy(parser)
	query.SetGroupByClause(groupByClause)
	//
	// windowClause := parseWindow()
	// query.setWindowClause(windowClause)

	return query
}

func ParseSelectElements(parser ISQLExprParser, parent *select_.SQLSelectQuery) {
	for {
		element := ParseSelectElement(parser)
		// addBeforeComments(item)
		parent.AddSelectElement(element)

		if !parser.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
}

func ParseSelectTargetElements(parser ISQLExprParser, parent *select_.SQLSelectQuery) {
	if !parser.AcceptAndNextToken(INTO) {
		return
	}

	for {
		expr := parser.ParseExpr()
		parent.AddSelectTargetElement(expr)
		if !parser.AcceptAndNextToken(SYMB_COMMA) {
			break
		}
	}
}

func ParseSelectElement(parser ISQLExprParser) *select_.SQLSelectElement {
	expr := parser.ParseExpr()

	if parser.Accept(SYMB_COMMA) || parser.Accept(FROM) || parser.Accept(BULK) {
		return select_.NewSelectElementWithExpr(expr)
	}

	as := parser.AcceptAndNextToken(AS)

	alias := parser.ParseExpr()
	if as && alias == nil {
		panic("alias is null. " + parser.Token().Error())
	}

	return select_.NewSelectElementWithAlias(expr, as, alias)
}

func parseUnionQueryBlock(parser ISQLExprParser, left select_.ISQLSelectQuery) select_.ISQLSelectQuery {

	var operator select_.SQLUnionOperator
	var right select_.ISQLSelectQuery
	if parser.AcceptAndNextToken(UNION) {

		operator = select_.UNION

		if parser.AcceptAndNextToken(ALL) {

			operator = select_.UNION_ALL

		} else if parser.AcceptAndNextToken(DISTINCT) {

			operator = select_.UNION_DISTINCT
		}

		right = ParseQueryBlock(parser)

	} else if parser.AcceptAndNextToken(MINUS) {

		operator = select_.MINUS
		right = ParseQueryBlock(parser)

	} else if parser.AcceptAndNextToken(EXCEPT) {
		operator = select_.EXCEPT

		if parser.AcceptAndNextToken(ALL) {
			operator = select_.EXCEPT_ALL

		} else if parser.AcceptAndNextToken(DISTINCT) {
			operator = select_.EXCEPT_DISTINCT
		}

		right = ParseQueryBlock(parser)

	} else if parser.AcceptAndNextToken(INTERSECT) {
		operator = select_.INTERSECT

		if parser.AcceptAndNextToken(ALL) {
			operator = select_.INTERSECT_ALL

		} else if parser.AcceptAndNextToken(DISTINCT) {
			operator = select_.INTERSECT_DISTINCT
		}

		right = ParseQueryBlock(parser)
	}

	if operator != "" {
		unionQuery := select_.NewSelectUnionQuery(left, operator, right)
		return parseUnionQueryBlock(parser, unionQuery)
	}

	return left
}

func parseFrom(parser ISQLExprParser) *select_.SQLFromClause {
	if !parser.AcceptAndNextToken(FROM) {
		return nil
	}
	tableReference := ParseTableReference(parser)

	return select_.NewFromClause(tableReference)
}

/**
     * E: T (JOIN, COMM) T JOIN T
     * T: primary
     */
func ParseTableReference(parser ISQLExprParser) select_.ISQLTableReference {
	tableReference := ParsePrimaryTableReference(parser)
	return parseJoinTableReference(parser, tableReference)
}

/**
     * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#table%20primary
     */
func ParsePrimaryTableReference(parser ISQLExprParser) select_.ISQLTableReference {

	var tableReference select_.ISQLTableReference

	if parser.AcceptAndNextToken(SYMB_LEFT_PAREN) {
		tableReference = ParseTableReference(parser)
		parser.AcceptAndNextTokenWithError(SYMB_RIGHT_PAREN, true)
		as := false
		if parser.AcceptAndNextToken(AS) {
			as = true
		}

		alias := parser.ParseIdentifier()
		if as && alias == nil {
			panic("TODO....")
		}

		tableReference.SetParen(true)
		tableReference.SetAs(as)
		tableReference.SetAlias(alias)

	} else if parser.Accept(WITH) || parser.Accept(SELECT) {
		subQuery := ParseSelectQuery(parser)

		tableReference = select_.NewSubQueryTableReference(subQuery)

	} else if parser.Accept(ONLY) {

	} else if parser.Accept(TABLE) {
		// tableReference = parserTableFunctionTableReference()

	} else if parser.AcceptAndNextToken(UNNEST) {

	} else if IsIdentifier(parser.Kind()) {

		name := ParseName(parser)
		// partitionExtensionClause := parsePartitionExtensionClause()

		as := false
		if parser.AcceptAndNextToken(AS) {
			as = true
		}

		alias := parser.ParseIdentifier()

		if as && alias == nil {
			panic("")
		}

		tableReference = select_.NewTableReferenceWithAlias(name, as, alias)
	}

	return tableReference
}

/**
* Join , comma
*/
func parseJoinTableReference(parser ISQLExprParser, left select_.ISQLTableReference) select_.ISQLTableReference {

	if parser.Accept(SYMB_RIGHT_PAREN) || parser.Accept(WHERE) {
		return left
	}

	if left == nil {
		panic("TableReference is nil.")
	}

	var joinType select_.SQLJoinType
	if parser.AcceptAndNextToken(SYMB_COMMA) {

		joinType = select_.COMMA

	} else if parser.AcceptAndNextToken(JOIN) {

		joinType = select_.JOIN

	} else if parser.AcceptAndNextToken(INNER) {

		parser.AcceptAndNextTokenWithError(JOIN, true)

		joinType = select_.INNER_JOIN

	} else if parser.AcceptAndNextToken(CROSS) {

		if parser.AcceptAndNextToken(JOIN) {
			joinType = select_.CROSS_JOIN
		} else if parser.AcceptAndNextToken(APPLY) {

			joinType = select_.CROSS_APPLY

		} else {
			panic("")
		}

	} else if parser.AcceptAndNextToken(LEFT) {

		joinType = select_.LEFT_JOIN
		if parser.AcceptAndNextToken(OUTER) {
			joinType = select_.LEFT_OUTER_JOIN
		}

		parser.AcceptAndNextTokenWithError(JOIN, true)

	} else if parser.AcceptAndNextToken(RIGHT) {

		joinType = select_.RIGHT_JOIN
		if parser.AcceptAndNextToken(OUTER) {
			joinType = select_.RIGHT_OUTER_JOIN
		}

		parser.AcceptAndNextTokenWithError(JOIN, true)

	} else if parser.AcceptAndNextToken(FULL) {

		joinType = select_.FULL_JOIN
		if parser.AcceptAndNextToken(OUTER) {
			joinType = select_.FULL_OUTER_JOIN
		}

		parser.AcceptAndNextTokenWithError(JOIN, true)

	} else if parser.AcceptAndNextToken(NATURAL) {

		joinType = select_.NATURAL_JOIN

		if parser.AcceptAndNextToken(INNER) {

			joinType = select_.NATURAL_INNER_JOIN
		} else if parser.AcceptAndNextToken(LEFT) {

			joinType = select_.NATURAL_LEFT_JOIN

			if parser.AcceptAndNextToken(OUTER) {
				joinType = select_.NATURAL_LEFT_OUTER_JOIN
			}

		} else if parser.AcceptAndNextToken(RIGHT) {

			joinType = select_.NATURAL_RIGHT_JOIN
			if parser.AcceptAndNextToken(OUTER) {
				joinType = select_.NATURAL_RIGHT_OUTER_JOIN
			}

		} else if parser.AcceptAndNextToken(FULL) {

			joinType = select_.NATURAL_FULL_JOIN
			if parser.AcceptAndNextToken(OUTER) {
				joinType = select_.NATURAL_FULL_OUTER_JOIN
			}
		}

		parser.AcceptAndNextTokenWithError(JOIN, true)

	} else if parser.AcceptAndNextToken(OUTER) {

		if parser.AcceptAndNextToken(APPLY) {
			joinType = select_.OUTER_APPLY

		} else {
			panic("")
		}

	}

	if joinType != "" {
		right := ParsePrimaryTableReference(parser)
		joinTableReference := select_.NewJoinTableReference(left, joinType, right)
		// parseJoinConditions(joinTableReference)
		return parseJoinTableReference(parser, joinTableReference)
	}

	return left
}

func parseWhere(parser ISQLExprParser) *select_.SQLWhereClause {
	if !parser.AcceptAndNextToken(WHERE) {
		return nil
	}

	var condition expr.ISQLExpr = nil
	if parser.AcceptAndNextToken(CURRENT) {

		if parser.AcceptAndNextToken(OF) {
			// TODO.
			// name := ParseName(parser)
			// condition = new SQLCurrentOfClause(name)
		} else {
			panic("")
		}

	} else {
		condition = parser.ParseExpr()
	}

	return select_.NewWhereClause(condition)
}

func parseGroupBy(parser ISQLExprParser) select_.ISQLGroupByClause {
	var x select_.ISQLGroupByClause = nil
	if parser.AcceptAndNextToken(GROUP) {
		parser.AcceptAndNextTokenWithError(BY, true)

		// x = select_.NewGroupByHavingClause()

		// SQLSetQuantifier setQuantifier = parseSetQuantifier()
		// x.setQuantifier(setQuantifier)

		for {
			element := parseGroupByElement(parser)
			x.AddElement(element)
			if !parser.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

		// SQLHavingClause havingClause = parseHavingClause()
		// x.setHavingClause(havingClause)

	} else if parser.AcceptAndNextToken(HAVING) {
		// x = new SQLGroupByClause()
		// x.setOrder(false)

		// SQLHavingClause havingClause = parseHavingClause()
		// x.setHavingClause(havingClause)

		parser.AcceptAndNextTokenWithError(GROUP, true)
		parser.AcceptAndNextTokenWithError(BY, true)

		// setQuantifier := parseSetQuantifier()
		// x.setQuantifier(setQuantifier)

		for {
			element := parseGroupByElement(parser)
			x.AddElement(element)

			if !parser.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
	}

	return x
}

func parseGroupByElement(parser ISQLExprParser) *select_.SQLGroupByElement {
	expr := parser.ParseExpr()
	return select_.NewGroupByElement(expr)
}

func ParseOrderBy(parser ISQLExprParser) *select_.SQLOrderByClause {
	if !parser.AcceptAndNextToken(ORDER) {
		return nil
	}

	parser.AcceptAndNextToken(BY)

	x := select_.NewOrderByClause()
	for {
		element := ParseOrderByElement(parser)
		x.AddElement(element)

		if parser.Kind() != SYMB_COMMA {
			break
		}
		NextTokenByParser(parser)
	}

	return x
}

func ParseOrderByElement(parser ISQLExprParser) *select_.SQLOrderByElement {
	key := parser.ParseExpr()

	var specification select_.SQLOrderingSpecification
	if parser.AcceptAndNextToken(ASC) {
		specification = select_.ASC
	} else if parser.AcceptAndNextToken(DESC) {
		specification = select_.ASC
	} else if parser.AcceptAndNextToken(USING) {

	}

	x := select_.NewOrderByElementWithSpecification(key, specification)

	return x
}

func ParseLimit(parser ISQLExprParser) select_.ISQLLimitClause {
	return nil
}

/**
* LIMIT row_count
* LIMIT offset, row_count
* LIMIT row_count OFFSET offset
*/
func parseLimitOffsetClause(parser ISQLExprParser) *select_.SQLLimitOffsetClause {
	if !parser.AcceptAndNextToken(LIMIT) {
		return nil
	}

	expr1 := parser.ParseExpr()
	offset := false
	var offsetExpr expr.ISQLExpr = nil
	var countExpr expr.ISQLExpr = nil

	if parser.AcceptAndNextToken(SYMB_COMMA) {
		offsetExpr = expr1
		countExpr = parser.ParseExpr()
	} else if parser.AcceptAndNextToken(OFFSET) {

		offset = true
		countExpr = expr1
		offsetExpr = parser.ParseExpr()
	} else {

		countExpr = expr1
	}

	return select_.NewLimitOffsetClauseWithOffset(offset, offsetExpr, countExpr)
}

func parseOffsetFetchClause(parser ISQLExprParser) *select_.SQLOffsetFetchClause {
	if !parser.Accept(OFFSET) && !parser.Accept(FETCH) {
		return nil
	}

	x := select_.NewOffsetFetchClause()

	if parser.AcceptAndNextToken(OFFSET) {

		offsetExpr := parser.ParseExpr()
		x.SetOffsetExpr(offsetExpr)
	}

	if parser.AcceptAndNextToken(FETCH) {

		countExpr := parser.ParseExpr()
		x.SetCountExpr(countExpr)
	}

	return x
}

func ParseLockClause(parser ISQLExprParser) select_.ISQLLockClause {
	var x select_.ISQLLockClause = nil
	if parser.Accept(FOR) {
		x = parseForUpdate(parser)
	} else if parser.AcceptAndNextToken(LOCK) {

		parser.AcceptAndNextTokenWithError(IN, true)

		if parser.AcceptAndNextToken(SHARE) {

			if parser.AcceptAndNextToken(MODE) {

				// x = new SQLLockInShareModeClause()
			} else {
				// todo.
				panic("")
			}

		} else {
			// todo.
			panic("")
		}

	}

	return x
}

/**
* FOR { UPDATE | NO KEY UPDATE | SHARE | KEY SHARE } [ OF table_name [, ...] ] [ NOWAIT | SKIP LOCKED| WAIT integer  ] [...]
*/
func parseForUpdate(parser ISQLExprParser) *select_.SQLForUpdateClause {

	if !parser.AcceptAndNextToken(FOR) {
		return nil
	}

	x := select_.NewForUpdateClause()
	// SQLForUpdateClause.SQLForType forType = null
	if parser.AcceptAndNextToken(UPDATE) {

		// forType = SQLForUpdateClause.SQLForType.UPDATE

	} else if parser.AcceptAndNextToken(NO) {

		if parser.AcceptAndNextToken(KEY) {
			parser.AcceptAndNextTokenWithError(UPDATE, true)
			// forType = SQLForUpdateClause.SQLForType.NO_KEY_UPDATE
		} else {
			// TODO.
		}

	} else if parser.AcceptAndNextToken(SHARE) {

		// forType = SQLForUpdateClause.SQLForType.SHARE

	} else if parser.AcceptAndNextToken(KEY) {

		if parser.AcceptAndNextToken(SHARE) {

			// forType = SQLForUpdateClause.SQLForType.KEY_SHARE

		} else {
			// TODO
		}

	} else {
		// TODO
		// throw new SQLParserException("TODO")
	}
	// x.setForType(forType)

	if parser.AcceptAndNextToken(OF) {

		for {
			// x.addColumn(parseName())
			if !parser.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
	}

	// SQLForUpdateClause.SQLForOption option = null
	if parser.AcceptAndNextToken(NOWAIT) {

		// option = new SQLForUpdateClause.SQLForNoWaitOption()

	} else if parser.AcceptAndNextToken(SKIP) {

		if parser.AcceptAndNextToken(LOCKED) {

			// option = new SQLForUpdateClause.SQLForSkipLockedOption()

		} else {
			// todo.
			panic("")
		}

	} else if parser.AcceptAndNextToken(WAIT) {

		// expr := parser.parseExpr()
		// option = new SQLForUpdateClause.SQLForWaitOption(expr)
	}
	// x.setForOption(option)

	return x
}

func parseReturningClause(parser ISQLExprParser) select_.ISQLReturningClause {
	if !parser.Accept(RETURN) && !parser.Accept(RETURNING) {
		return nil
	}

	var x select_.ISQLReturningClause = nil
	if parser.AcceptAndNextToken(RETURN) {
		// x = new ISQLReturningClause.SQLReturnIntoClause()

	} else if parser.AcceptAndNextToken(RETURNING) {
		// x = new ISQLReturningClause.SQLReturningIntoClause()

	}

	//  RETURN INTO data_item [, data_item ]...
	//  RETURN BULK COLLECT INTO data_item [, data_item ]...
	if parser.AcceptAndNextToken(BULK) {
		parser.AcceptAndNextTokenWithError(COLLECT, true)
		// x.setBulkCollect(true)
	}

	if !parser.AcceptAndNextToken(INTO) {
		for {
			// returningItem := parser.ParseExpr()
			// x.addReturningItem(returningItem)
			if !parser.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}
	}

	if parser.AcceptAndNextToken(INTO) {

		for {
			// ISQLExpr intoItem = parseExpr()
			// x.addIntoItem(intoItem)
			if !parser.AcceptAndNextToken(SYMB_COMMA) {
				break
			}
		}

	}

	return x
}
