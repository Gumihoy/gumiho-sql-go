package factory

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/parser"
	"gumihoy.com/sql/basic/visitor"
	"gumihoy.com/sql/config"
	"gumihoy.com/sql/db"
	mariadbParser "gumihoy.com/sql/dialect/mariadb/parser"
	mariadbVisitor "gumihoy.com/sql/dialect/mariadb/visitor"
	mysqlParser "gumihoy.com/sql/dialect/mysql/parser"
	mysqlVisitor "gumihoy.com/sql/dialect/mysql/visitor"
	oracleParser "gumihoy.com/sql/dialect/oracle/parser"
	oracleVisitor "gumihoy.com/sql/dialect/oracle/visitor"
	postgresqlParser "gumihoy.com/sql/dialect/postgresql/parser"
	postgresqlVisitor "gumihoy.com/sql/dialect/postgresql/visitor"
	"strings"
)

func ParseStatementsByDBType(sql string, dbType db.DBType) []ast.ISQLObject {
	return ParseStatementsByConfig(sql, parser.NewConfig(dbType))
}

func ParseStatementsByConfig(sql string, config parser.Config) []ast.ISQLObject {
	p := createStatementParser(sql, config)
	return parser.ParseStatements(p)
}

func createStatementParser(sql string, config parser.Config) parser.ISQLStatementParser {
	switch config.DBType {
	case db.MariaDB:
		return mariadbParser.NewStatementParserBySQL(sql)
	case db.MySQL:
		return mysqlParser.NewStatementParserBySQL(sql)
	case db.Oracle:
		return oracleParser.NewStatementParserBySQL(sql)
	case db.PostgreSQL:
		return postgresqlParser.NewStatementParserBySQL(sql)
	default:
		return parser.NewStatementParserBySQL(sql)
	}
}

func CreateASTOutputVisitor(builder *strings.Builder, dbType db.DBType, config config.Output) visitor.ISQLOutputVisitor {
	switch dbType {
	case db.MariaDB:
		return mariadbVisitor.NewOutputVisitor(builder, config)
	case db.MySQL:
		return mysqlVisitor.NewOutputVisitor(builder, config)
	case db.Oracle:
		return oracleVisitor.NewOutputVisitor(builder, config)
	case db.PostgreSQL:
		return postgresqlVisitor.NewOutputVisitor(builder, config)

	default:
		return visitor.NewOutputVisitor(builder, config)
	}
}

// func CreateTranslateASTOutputVisitor(builder *strings.Builder, sourceDBType db.DBType, targetDBType db.DBType, config config.Output) visitor.ISQLOutputVisitor {
// 	if sourceDBType == "" {
// 		sourceDBType = db.SQL
// 	}
// 	if targetDBType == "" {
// 		targetDBType = db.SQL
// 	}
//
// 	switch sourceDBType {
// 	case db.Oracle:
// 		switch targetDBType {
// 		case db.MySQL:
// 			// return newOracle2MySQLASTOutputVisitor(out, config)
// 			//                    case MariaDB:
// 			//                        return new Oracle2MariaDBASTOutputVisitor(out, config)
// 		case db.EDB:
// 			// return newOracle2EDBASTOutputVisitor(out, config)
// 		case db.PostgreSQL:
// 			// return newOracle2PostgreSQLASTOutputVisitor(out, config)
// 		default:
// 			return oracleVisitor.NewOutputVisitor(builder, config)
// 		}
// 	// case db.MySQL:
// 	// 	switch targetDBType {
// 	// 	//                    case Oracle:
// 	// 	//                        return new MySQL2OracleASTOutputVisitor(out, config)
// 	// 	default:
// 	// 		return new MySQLASTOutputVisitor(out)
// 	// 	}
// 	// case db.MariaDB:
// 	// 	switch targetDBType {
// 	// 	default:
// 	// 		//                        return new MariaDBSQLASTOutputVisitor(out)
// 	// 	}
// 	// case db.EDB:
// 	// 	switch targetDBType {
// 	// 	default:
// 	// 		return new EDBASTOutputVisitor(out)
// 	// 	}
// 	// case db.PostgreSQL:
// 	// 	switch targetDBType {
// 	// 	default:
// 	// 		return new PostgreSQLASTOutputVisitor(out)
// 	// 	}
// 	default:
// 		return visitor.NewOutputVisitor(builder, config)
// 	}
// 	return nil
// }

// func CreateASTTranslateVisitors(config *translation.SQLTransformConfig) []*translationVisitor.SQLASTTransformVisitor {
// 	if config == nil {
// 		panic("Translate config is null.")
// 	}
// 	sourceDBType := config.SourceDBType
// 	targetDBType := config.TargetDBType
// 	targetDBVersion := config.TargetDBVersion
// 	if sourceDBType == "" {
// 		panic("source database type is null.")
// 	}
// 	if targetDBType == "" {
// 		panic("target database type is null.")
// 	}
// 	if targetDBVersion == "" {
// 		panic("target database version is null.")
// 	}
//
// 	var visitors []*translationVisitor.SQLASTTransformVisitor
// 	switch sourceDBType {
// 	case db.Oracle:
// 		switch targetDBType {
// 		case db.Oracle:
// 			// visitors.add(newOracleSQLRemovePropertyASTVisitor(config))
// 			break
// 		// case EDB:
// 		// 	visitors.add(newOracleSQLRemovePropertyASTVisitor(config))
// 		// 	visitors.add(newOracle2EDBMixRemoveDoubleQuotesASTVisitor(config))
// 		// switch (targetDBVersion) {
// 		// 	case VERSION_9_6:
// 		// 		visitors.add(new
// 		// 		Oracle2EDBVersion9_6ASTTransformVisitor(config))
// 		// 		break
// 		// 	case VERSION_10:
// 		// 		visitors.add(new
// 		// 		Oracle2EDBVersion10ASTTransformVisitor(config))
// 		// 		break
// 		// 	}
// 		// break
// 		case db.PostgreSQL:
// 			// visitors.add(newOracleSQLRemovePropertyASTVisitor(config))
// 			// visitors.add(newOracle2EDBMixRemoveDoubleQuotesASTVisitor(config))
// 			switch targetDBVersion {
// 			case postgresql.Version_9_6:
// 				// visitors.add(newOracle2PostgreSQLVersion9_6ASTTransformVisitor(config))
// 				break
// 			case postgresql.Version_10:
// 				// visitors.add(newOracle2PostgreSQLVersion10ASTTransformVisitor(config))
// 				break
// 			}
// 			break
// 		case db.MySQL:
// 			// visitors.add(newOracleSQLRemovePropertyASTVisitor(config))
//
// 			switch targetDBVersion {
// 			case mysql.Version_5_6:
// 			case mysql.Version_5_7:
// 				// visitors.add(newSQLWithClauseSubQueryTranslateAndRemoveASTVisitor(config))
// 				// visitors.add(newSQLCreateViewSubQueryTableRefToCreateViewASTVisitor(config))
// 				break
// 			default:
// 				break
// 			}
// 			// visitors.add(newSQLOuterJoinToJoinASTVisitor(config))
// 			// visitors.add(newSQLRowNumToLimitASTVisitor(config))
// 			// visitors.add(newOracleSQLIntersectOrMinusToJoinASTVisitor(config))
// 			// visitors.add(newSQLBindVarToVarASTVisitor(config))
//
// 			switch targetDBVersion {
// 			case mysql.Version_5_6:
// 				// visitors.add(newOracle2MySQLVersion5_6ASTTransformVisitor(config))
// 				break
// 			case mysql.Version_5_7:
// 				// visitors.add(newOracle2MySQLVersion5_7ASTTransformVisitor(config))
// 				break
// 			case mysql.Version_8_0:
// 				// visitors.add(newOracle2MySQLVersion8_0ASTTransformVisitor(config))
// 				break
// 			default:
// 				break
// 			}
//
// 			// visitors.add(newSQLRenameColumnASTVisitor(config))
// 			// visitors.add(newSQLRenameObjectNameASTVisitor(config))
// 			//
// 			// visitors.add(newSQLDefaultClauseToTriggerAndRemoveASTVisitor(config))
//
// 			if config.IsRemoveSchema {
// 				// visitors.add(newSQLRemoveSchemaASTVisitor(config))
// 			}
//
// 			// visitors.add(newSQLOptimizationASTVisitor(config))
// 			// visitors.add(newSQLAddReverseQuotesASTVisitor(config))
// 			break
// 		}
// 		break
//
// 	default:
// 		panic("")
// 	}
// 	return visitors
// }
