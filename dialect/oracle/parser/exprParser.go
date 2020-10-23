package parser

import (
	exprAST "gumihoy.com/sql/basic/ast/expr"
	"gumihoy.com/sql/basic/ast/expr/operator"
	"gumihoy.com/sql/basic/parser"
)

type OracleExprParser struct {
	*parser.SQLExprParser
}

func NewExprParserBySQL(sql string) *OracleExprParser {
	var x OracleExprParser
	x.SQLExprParser = parser.NewExprParserBySQL(sql)
	return &x
}

func NewExprParserByLexer(lexer parser.ISQLLexer) *OracleExprParser {
	var x OracleExprParser
	x.SQLExprParser = parser.NewExprParserByLexer(lexer)
	return &x
}

func (x *OracleExprParser) ParseExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	if x.Token().Kind == parser.EOF ||
		x.Token().Kind == parser.SYMB_SEMI {
		return left
	}

	expr := x.ParseBitXorOperatorExprWithLeft(left)
	expr = x.ParseMultiplicativeOperatorExprWithLeft(left)
	expr = x.ParseAdditiveOperatorExprWithLeft(left)
	expr = x.ParseShiftOperatorExprWithLeft(left)
	expr = x.ParseBitAndOperatorExprWithLeft(left)
	expr = x.ParseBitOrOperatorExprWithLeft(left)
	expr = x.ParseComparisonOperatorExprWithLeft(left)
	expr = x.ParseAndOperatorExprWithLeft(left)
	expr = x.ParseXorOperatorExprWithLeft(left)
	expr = x.ParseOrOperatorExprWithLeft(left)
	return expr

}

/**
 * E: T ^ T ^ T
 * T: primaryExpr
 * https://dev.mysql.com/doc/refman/8.0/en/operator-precedence.html
 */
func (x *OracleExprParser) ParseBitXorOperatorExpr() exprAST.ISQLExpr {
	left := x.ParsePrimaryExpr()
	return x.ParseBitXorOperatorExprWithLeft(left)
}
func (x *OracleExprParser) ParseBitXorOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_BIT_XOR:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseMultiplicativeOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseBitXorOperatorExpr()
	return x.ParseMultiplicativeOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseMultiplicativeOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_STAR:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case parser.SYMB_SLASH:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case parser.SYMB_PERCENT:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case parser.DIV:
		parser.NextTokenByParser(x)
		right := x.ParseBitXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.BIT_XOR, right)
		return x.ParseMultiplicativeOperatorExprWithLeft(left)
	case parser.MOD:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseAdditiveOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseMultiplicativeOperatorExpr()
	return x.ParseAdditiveOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseAdditiveOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_PLUS:
		parser.NextTokenByParser(x)
		right := x.ParseMultiplicativeOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.PLUS, right)
		return x.ParseAdditiveOperatorExprWithLeft(left)
	case parser.SYMB_MINUS:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseShiftOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseAdditiveOperatorExpr()
	return x.ParseShiftOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseShiftOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_LESS_THAN_LESS_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseAdditiveOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return x.ParseShiftOperatorExprWithLeft(left)
	case parser.SYMB_GREATER_THAN_GREATER_THAN:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseBitAndOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseShiftOperatorExpr()
	return x.ParseBitAndOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseBitAndOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_BIT_AND:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseBitOrOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseBitAndOperatorExpr()
	return x.ParseBitOrOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseBitOrOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_BIT_OR:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseComparisonOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseBitOrOperatorExpr()
	return x.ParseComparisonOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseComparisonOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_EQUAL:
		parser.NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.EQ, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case parser.SYMB_LESS_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.LESS_THAN, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case parser.SYMB_LESS_THAN_EQUAL:
		parser.NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.LESS_THAN_EQ, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case parser.SYMB_GREATER_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseBitOrOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.GREATER_THAN, right)
		return x.ParseComparisonOperatorExprWithLeft(left)
	case parser.SYMB_GREATER_THAN_EQUAL:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseAndOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseComparisonOperatorExpr()
	return x.ParseAndOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseAndOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.AND:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseXorOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseAndOperatorExpr()
	return x.ParseXorOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseXorOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.SYMB_LESS_THAN_LESS_THAN:
		parser.NextTokenByParser(x)
		right := x.ParseAndOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.SHIFT_LEFT, right)
		return x.ParseXorOperatorExprWithLeft(left)
	case parser.SYMB_GREATER_THAN_GREATER_THAN:
		parser.NextTokenByParser(x)
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
func (x *OracleExprParser) ParseOrOperatorExpr() exprAST.ISQLExpr {
	left := x.ParseXorOperatorExpr()
	return x.ParseOrOperatorExprWithLeft(left)
}

func (x *OracleExprParser) ParseOrOperatorExprWithLeft(left exprAST.ISQLExpr) exprAST.ISQLExpr {
	kind := x.Token().Kind
	switch kind {
	case parser.OR:
		parser.NextTokenByParser(x)
		right := x.ParseXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.OR, right)
		return x.ParseOrOperatorExprWithLeft(left)
	case parser.SYMB_LOGICAL_OR:
		parser.NextTokenByParser(x)
		right := x.ParseXorOperatorExpr()
		left = operator.NewBinaryOperator(left, operator.CONCAT, right)
		return x.ParseOrOperatorExprWithLeft(left)
	}
	return left
}
