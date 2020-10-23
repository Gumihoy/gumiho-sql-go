package parser

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/statement"
)

type ISQLStatementParser interface {
	ISQLParser
	ExprParser() ISQLExprParser
	SetExprParser(exprParser ISQLExprParser)
}

type SQLStatementParser struct {
	*SQLParser
	exprParser ISQLExprParser
}

func (x *SQLStatementParser) ExprParser() ISQLExprParser {
	return x.exprParser
}

func (x *SQLStatementParser) SetExprParser(exprParser ISQLExprParser) {
	x.exprParser = exprParser
}

func NewStatementParserBySQL(sql string) *SQLStatementParser {
	return NewStatementParserByLexer(NewLexer(sql))
}

func NewStatementParserByLexer(lexer ISQLLexer) *SQLStatementParser {
	return NewStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewStatementParserByExprParser(exprParser ISQLExprParser) *SQLStatementParser {
	var x SQLStatementParser
	x.SQLParser = NewParserByLexer(exprParser.Lexer())
	x.exprParser = exprParser
	return &x
}

func ParseStatements(parser ISQLStatementParser) []ast.ISQLObject {
	return ParseStatementsWithParent(parser,nil)
}

func ParseStatementsWithParent(x ISQLStatementParser, parent ast.ISQLObject) []ast.ISQLObject {
	NextTokenByParser(x)
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

		stmt := ParseStatement(x.ExprParser())
		if stmt == nil {
			break
		}
		stmt.SetParent(parent)
		stmts = append(stmts, stmt)
	}
	return stmts
}





// ------------------------------------------------------------- DML -------------------------------------------------------------

// ------------------------ Select ------------------------

type ISQLSelectStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLSelectStatementParser struct {
	*SQLStatementParser
}

func (x *SQLSelectStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != SELECT {
		return nil
	}
	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewSelectStatementParserBySQL(sql string) *SQLSelectStatementParser {
	return NewSelectStatementParserByLexer(NewLexer(sql))
}

func NewSelectStatementParserByLexer(lexer ISQLLexer) *SQLSelectStatementParser {
	return NewSelectStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewSelectStatementParserByExprParser(exprParser ISQLExprParser) *SQLSelectStatementParser {
	return &SQLSelectStatementParser{NewStatementParserByExprParser(exprParser)}
}


// ------------------------ Delete ------------------------

type ISQLDeleteStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLDeleteStatementParser struct {
	*SQLStatementParser
}

func (x *SQLDeleteStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != SELECT {
		return nil
	}
	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewDeleteStatementParserBySQL(sql string) *SQLDeleteStatementParser {
	return NewDeleteStatementParserByLexer(NewLexer(sql))
}

func NewDeleteStatementParserByLexer(lexer ISQLLexer) *SQLDeleteStatementParser {
	return NewDeleteStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewDeleteStatementParserByExprParser(exprParser ISQLExprParser) *SQLDeleteStatementParser {
	return &SQLDeleteStatementParser{NewStatementParserByExprParser(exprParser)}
}


// ------------------------ Insert ------------------------
type ISQLInsertStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLInsertStatementParser struct {
	*SQLStatementParser
}

func (x *SQLInsertStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != SELECT {
		return nil
	}
	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewInsertStatementParserBySQL(sql string) *SQLInsertStatementParser {
	return NewInsertStatementParserByLexer(NewLexer(sql))
}

func NewInsertStatementParserByLexer(lexer ISQLLexer) *SQLInsertStatementParser {
	return NewInsertStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewInsertStatementParserByExprParser(exprParser ISQLExprParser) *SQLInsertStatementParser {
	return &SQLInsertStatementParser{NewStatementParserByExprParser(exprParser)}
}


// ------------------------ Update ------------------------

type ISQLUpdateStatementParser interface {
	ISQLStatementParser
	Parse() statement.ISQLStatement
}

type SQLUpdateStatementParser struct {
	*SQLStatementParser
}

func (x *SQLUpdateStatementParser) Parse() statement.ISQLStatement {
	if x.Kind() != SELECT {
		return nil
	}
	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

func NewUpdateStatementParserBySQL(sql string) *SQLUpdateStatementParser {
	return NewUpdateStatementParserByLexer(NewLexer(sql))
}

func NewUpdateStatementParserByLexer(lexer ISQLLexer) *SQLUpdateStatementParser {
	return NewUpdateStatementParserByExprParser(NewExprParserByLexer(lexer))
}

func NewUpdateStatementParserByExprParser(exprParser ISQLExprParser) *SQLUpdateStatementParser {
	return &SQLUpdateStatementParser{NewStatementParserByExprParser(exprParser)}
}





// ------------------------------------------------------------- DDL -------------------------------------------------------------

type ISQLDDLStatementParser interface {
	ISQLStatementParser
	ParseAlter() statement.ISQLStatement
	ParseCreate() statement.ISQLStatement
	ParseDrop() statement.ISQLStatement
}

// ------------------------ Database ------------------------

type ISQLDatabaseStatementParser interface {
	ISQLDDLStatementParser
}

type SQLDatabaseStatementParser struct {
	*SQLStatementParser
}

func (x *SQLDatabaseStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLDatabaseStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLDatabaseStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}




// ------------------------ Function ------------------------

type ISQLFunctionStatementParser interface {
	ISQLDDLStatementParser
}

type SQLFunctionStatementParser struct {
	*SQLStatementParser
}

func (x *SQLFunctionStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLFunctionStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLFunctionStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ Index ------------------------

type ISQLIndexStatementParser interface {
	ISQLDDLStatementParser
}

type SQLIndexStatementParser struct {
	*SQLStatementParser
}

func (x *SQLIndexStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLIndexStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLIndexStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ Package ------------------------

type ISQLPackageStatementParser interface {
	ISQLDDLStatementParser
}

type SQLPackageStatementParser struct {
	*SQLStatementParser
}

func (x *SQLPackageStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLPackageStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLPackageStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ Procedure ------------------------

type ISQLProcedureStatementParser interface {
	ISQLDDLStatementParser
}

type SQLProcedureStatementParser struct {
	*SQLStatementParser
}

func (x *SQLProcedureStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLProcedureStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLProcedureStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ Role ------------------------

type ISQLRoleStatementParser interface {
	ISQLDDLStatementParser
}

type SQLRoleStatementParser struct {
	*SQLStatementParser
}

func (x *SQLRoleStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLRoleStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLRoleStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ Schema ------------------------

type ISQLSchemaStatementParser interface {
	ISQLDDLStatementParser
}

type SQLSchemaStatementParser struct {
	*SQLStatementParser
}

func (x *SQLSchemaStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLSchemaStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLSchemaStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}




// ------------------------ Sequence ------------------------

type ISQLSequenceStatementParser interface {
	ISQLDDLStatementParser
}

type SQLSequenceStatementParser struct {
	*SQLStatementParser
}

func (x *SQLSequenceStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLSequenceStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLSequenceStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}








// ------------------------ Synonym ------------------------

type ISQLSynonymStatementParser interface {
	ISQLDDLStatementParser
}

type SQLSynonymStatementParser struct {
	*SQLStatementParser
}

func (x *SQLSynonymStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLSynonymStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLSynonymStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}




// ------------------------ Table ------------------------

type ISQLTableStatementParser interface {
	ISQLDDLStatementParser
}

type SQLTableStatementParser struct {
	*SQLStatementParser
}

func (x *SQLTableStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLTableStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLTableStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ Trigger ------------------------

type ISQLTriggerStatementParser interface {
	ISQLDDLStatementParser
}

type SQLTriggerStatementParser struct {
	*SQLStatementParser
}

func (x *SQLTriggerStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLTriggerStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLTriggerStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ Type ------------------------

type ISQLTypeStatementParser interface {
	ISQLDDLStatementParser
}

type SQLTypeStatementParser struct {
	*SQLStatementParser
}

func (x *SQLTypeStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLTypeStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLTypeStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}



// ------------------------ User ------------------------

type ISQLUserStatementParser interface {
	ISQLDDLStatementParser
}

type SQLUserStatementParser struct {
	*SQLStatementParser
}

func (x *SQLUserStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLUserStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLUserStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}

// ------------------------ View ------------------------

type ISQLViewStatementParser interface {
	ISQLDDLStatementParser
}

type SQLViewStatementParser struct {
	*SQLStatementParser
}

func (x *SQLViewStatementParser) ParseAlter() statement.ISQLStatement {
	if !x.Accept(ALTER) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLViewStatementParser) ParseCreate() statement.ISQLStatement {
	if !x.Accept(CREATE) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}
func (x *SQLViewStatementParser) ParseDrop() statement.ISQLStatement {
	if !x.Accept(DROP) {
		return nil
	}

	query := ParseSelectQuery(x.ExprParser())
	return statement.NewSelectStatement(query)
}


