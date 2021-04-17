package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
)

type PostgreSQLExprParser struct {
	*parser.SQLExprParser
}

func NewExprParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLExprParser {
	return NewExprParserByLexer(parser.NewLexer(sql), dbType, config)
}

func NewExprParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *PostgreSQLExprParser {
	return &PostgreSQLExprParser{parser.NewExprParserByLexer(lexer, dbType, config)}
}

func (x *PostgreSQLExprParser) CreateSQLTableStatementParser() parser.ISQLTableStatementParser {
	return NewTableStatementParserByExprParser(x)
}

// ---------------------- Table
func (sp *PostgreSQLExprParser) ParseAlterTableAction(child parser.ISQLExprParser) table.ISQLAlterTableAction {
	if sp.Accept(parser.ADD) {
		return child.ParseAddAlterTableAction(child)

	} else if sp.Accept(parser.ALTER) {
		return child.ParseAlterAlterTableAction(child)

	} else if sp.Accept(parser.DROP) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.RENAME) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.ATTACH) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.DETACH) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.VALIDATE) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.DISABLE) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.ENABLE) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.FORCE) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.NO) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.CLUSTER) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.SET) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.RESET) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.INHERIT) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.OWNER) {
		return child.ParseDropAlterTableAction(child)

	} else if sp.Accept(parser.REPLICA) {
		return child.ParseDropAlterTableAction(child)

	}

	return nil
}
func (sp *PostgreSQLExprParser) ParseAddColumnAlterTableAction(child parser.ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.Accept(parser.COLUMN) && !parser.IsIdentifier(sp.Kind()) {
		return nil
	}

	x := table.NewAddColumnAlterTableAction()

	hasColumn := sp.AcceptAndNextToken(parser.COLUMN)
	x.HasColumn = hasColumn

	paren := sp.AcceptAndNextToken(parser.SYMB_LEFT_PAREN)
	column := child.ParseTableColumn(child)
	x.AddColumn(column)
	if paren {
		sp.AcceptAndNextTokenWithError(parser.SYMB_RIGHT_PAREN, true)
	}
	return x
}

func (sp *PostgreSQLExprParser) ParseDropAlterTableAction(child parser.ISQLExprParser) table.ISQLAlterTableAction {
	if !sp.AcceptAndNextToken(parser.DROP) {
		return nil
	}

	if sp.AcceptAndNextToken(parser.CONSTRAINT) {
		x := table.NewDropTableConstraintAlterTableAction()
		ifExists := parser.ParseIfExists(child)
		x.IfExists = ifExists

		name := parser.ParseName(child)
		x.SetName(name)

		return x

	} else if sp.Accept(parser.COLUMN) || sp.Accept(parser.IF) || parser.IsIdentifier(sp.Kind()) {
		x := table.NewDropColumnAlterTableAction()

		hasColumn := sp.AcceptAndNextToken(parser.COLUMN)
		x.HasColumn = hasColumn

		ifExists := parser.ParseIfExists(child)
		x.IfExists = ifExists

		column := parser.ParseName(child)
		x.SetColumn(column)
		return x
	}

	panic(sp.UnSupport())
}

// ---------------------- VIEW
