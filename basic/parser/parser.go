package parser

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/statement"
	"gumihoy.com/sql/basic/parser/kind"
	"gumihoy.com/sql/dbtype"
	"strconv"
)

type Config struct {
	DBType dbtype.DBType
}

func NewConfig(dbType dbtype.DBType) Config {
	var c Config
	c.DBType = dbType
	return c
}

type IParser interface {
	Lexer() ILexer
	ParseCreateStatement() statement.ISQLStatement
	ParseDropStatement() statement.ISQLStatement

	ParseDeleteStatement() statement.ISQLStatement
	ParseInsertStatement() statement.ISQLStatement
	ParseUpdateStatement() statement.ISQLStatement
	ParseSelectStatement() statement.ISQLStatement
}

func ParseStatements(parser IParser) []ast.ISQLObject {
	var statements []ast.ISQLObject

	NextToken(parser.Lexer())
	token := parser.Lexer().Token()
	for {
		k := token.Kind
		if k == kind.EOF {
			break
		}

		if k == kind.SEMI {
			if len(statements) > 0 {
				stmt := statements[len(statements)]
				stmt.SetAfterSemi(true)
			}
			continue
		}
		stmt := ParseStatement(parser)
		if stmt != nil {
			statements = append(statements, stmt)
		}
	}
	return statements
}

func ParseStatement(parser IParser) statement.ISQLStatement {
	token := parser.Lexer().Token()
	switch token.Kind {
	case kind.ALTER:
		return nil
	case kind.CREATE:
		return nil
	case kind.DELETE:
		return nil
	case kind.INSERT:
		return nil
	case kind.UPDATE:
		return nil
	case kind.SELECT:
		return parser.ParseSelectStatement()
	default:
		panic("TODO. line " + strconv.Itoa(token.line) + ", col" + strconv.Itoa(token.col) + " UnSupport " + string(token.Kind))
	}
	return nil
}

type Parser struct {
	lexer ILexer
}

func (parser *Parser) Lexer() ILexer {
	return parser.lexer
}

func (parser *Parser) ParseCreateStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *Parser) ParseDropStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *Parser) ParseDeleteStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *Parser) ParseInsertStatement() statement.ISQLStatement {
	panic("implement me")
}

func (parser *Parser) ParseUpdateStatement() statement.ISQLStatement {
	panic("implement me")
}

func NewParserBySQL(sql string) *Parser {
	return NewParserByLexer(NewLexer(sql))
}

func NewParserByLexer(lexer ILexer) *Parser {
	return &Parser{lexer}
}

func (parser *Parser) ParseSelectStatement() statement.ISQLStatement {
	return nil
}

func ParseExpr() {

}
