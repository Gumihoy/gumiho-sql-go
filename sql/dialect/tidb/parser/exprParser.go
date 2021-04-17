package parser

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/parser"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	mysqlParser "github.com/Gumihoy/gumiho-sql-go/sql/dialect/mysql/parser"
)

var complexFunctionNameMap = make(map[string]bool)
var nonParametricFunctionNameMap = make(map[string]bool)

func init() {
	complexFunctionNameMap["CAST"] = true
}

type TiDBExprParser struct {
	*mysqlParser.MySQLExprParser
}

func NewExprParserBySQL(sql string, dbType db.Type, config *parser.SQLParseConfig) *TiDBExprParser {
	return NewExprParserByLexer(parser.NewLexer(sql), dbType, config)
}

func NewExprParserByLexer(lexer parser.ISQLLexer, dbType db.Type, config *parser.SQLParseConfig) *TiDBExprParser {
	return NewExprParserByExprParser(mysqlParser.NewExprParserByLexer(lexer, dbType, config))
}

func NewExprParserByExprParser(parser *mysqlParser.MySQLExprParser) *TiDBExprParser {
	x := new(TiDBExprParser)
	x.MySQLExprParser = parser
	return x
}

// func (x *TiDBExprParser) CreateSQLDatabaseStatementParser() parser.ISQLDatabaseStatementParser {
// 	return NewDatabaseStatementParserByExprParser(x)
// }
//
// func (x *TiDBExprParser) CreateSQLFunctionStatementParser() parser.ISQLFunctionStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLIndexStatementParser() parser.ISQLIndexStatementParser {
// 	return nil
// }
//
// func (x *TiDBExprParser) CreateSQLPackageStatementParser() parser.ISQLPackageStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLPackageBodyStatementParser() parser.ISQLPackageBodyStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLProcedureStatementParser() parser.ISQLProcedureStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLRoleStatementParser() parser.ISQLRoleStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLSchemaStatementParser() parser.ISQLSchemaStatementParser {
// 	return NewSchemaStatementParserByExprParser(x)
// }
// func (x *TiDBExprParser) CreateSQLSequenceStatementParser() parser.ISQLSequenceStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLSynonymStatementParser() parser.ISQLSynonymStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLTableStatementParser() parser.ISQLTableStatementParser {
// 	return NewTableStatementParserByExprParser(x)
// }
//
// func (x *TiDBExprParser) CreateSQLTriggerStatementParser() parser.ISQLTriggerStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLTypeStatementParser() parser.ISQLTypeStatementParser {
// 	return nil
// }
// func (x *TiDBExprParser) CreateSQLTypeBodyStatementParser() parser.ISQLTypeBodyStatementParser {
// 	return parser.NewTypeBodyStatementParserByExprParser(x)
// }
// func (x *TiDBExprParser) CreateSQLUserStatementParser() parser.ISQLUserStatementParser {
// 	return parser.NewUserStatementParserByExprParser(x)
// }
// func (x *TiDBExprParser) CreateSQLViewStatementParser() parser.ISQLViewStatementParser {
// 	return NewViewStatementParserByExprParser(x)
// }
//
// func (x *TiDBExprParser) CreateSQLDeleteStatementParser() parser.ISQLDeleteStatementParser {
// 	return NewDeleteStatementParserByExprParser(x)
// }
// func (x *TiDBExprParser) CreateSQLInsertStatementParser() parser.ISQLInsertStatementParser {
// 	return NewInsertStatementParserByExprParser(x)
// }
// func (x *TiDBExprParser) CreateSQLSelectStatementParser() parser.ISQLSelectStatementParser {
// 	return NewSelectStatementParserByExprParser(x)
// }
// func (x *TiDBExprParser) CreateSQLUpdateStatementParser() parser.ISQLUpdateStatementParser {
// 	return NewUpdateStatementParserByExprParser(x)
// }


