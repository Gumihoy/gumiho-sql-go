package translation

import (
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/visitor"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"github.com/Gumihoy/gumiho-sql-go/sql/db/mysql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db/postgresql"
	"github.com/Gumihoy/gumiho-sql-go/translation/config"
	"github.com/Gumihoy/gumiho-sql-go/translation/result"
	visitor2 "github.com/Gumihoy/gumiho-sql-go/translation/visitor"
	mysql2 "github.com/Gumihoy/gumiho-sql-go/translation/visitor/oracle/mysql"
	"strings"
)

/**
 *
 */
func OracleToMySQL(sourceSQL string) *result.SQLTransformResult {
	return OracleToMySQLWithConfig(sourceSQL, nil)
}

func OracleToMySQLWithConfig(sourceSQL string, transformConfig *config.SQLTransformConfig) *result.SQLTransformResult {
	if transformConfig == nil {
		transformConfig = config.NewSQLTransformConfig()
	}
	transformConfig.SourceDBType = db.Oracle
	transformConfig.TargetDBType = db.MySQL
	transformConfig.TargetDBVersion = mysql.Version_5_6
	return Translate(sourceSQL, transformConfig)
}

/**
 *
 */
func OracleToEDB(sourceSQL string) *result.SQLTransformResult {
	return OracleToEDBWithConfig(sourceSQL, nil)
}

func OracleToEDBWithConfig(sourceSQL string, transformConfig *config.SQLTransformConfig) *result.SQLTransformResult {
	if transformConfig == nil {
		transformConfig = config.NewSQLTransformConfig()
	}
	transformConfig.SourceDBType = db.Oracle
	transformConfig.TargetDBType = db.EDB
	return Translate(sourceSQL, transformConfig)
}

/**
 *
 */
func OracleToPostgreSQL(sourceSQL string) *result.SQLTransformResult {
	return OracleToPostgreSQLWithConfig(sourceSQL, nil)
}

func OracleToPostgreSQLWithConfig(sourceSQL string, transformConfig *config.SQLTransformConfig) *result.SQLTransformResult {
	if transformConfig == nil {
		transformConfig = config.NewSQLTransformConfig()
	}
	transformConfig.SourceDBType = db.Oracle
	transformConfig.TargetDBType = db.PostgreSQL
	return Translate(sourceSQL, transformConfig)
}

/**
 *
 */
func MySQLToOracle(sourceSQL string) *result.SQLTransformResult {
	return MySQLToOracleWithConfig(sourceSQL, nil)
}

func MySQLToOracleWithConfig(sourceSQL string, transformConfig *config.SQLTransformConfig) *result.SQLTransformResult {
	if transformConfig == nil {
		transformConfig = config.NewSQLTransformConfig()
	}
	transformConfig.SourceDBType = db.MySQL
	transformConfig.TargetDBType = db.Oracle
	return Translate(sourceSQL, transformConfig)
}

func Translate(sourceSQL string, transformConfig *config.SQLTransformConfig) *result.SQLTransformResult {

	sourceDBType := transformConfig.SourceDBType
	targetDBType := transformConfig.TargetDBType

	transformResult := result.NewSQLTransformResult(sourceSQL)

	stmts := sql.ParseStatementsByDBType(sourceSQL, sourceDBType)

	transformConfig.Stmts = stmts
	// buffers := newStringBuilder()

	// outputConfig := newSQLOutputConfig()

	visitors := CreateASTTranslateVisitors(transformConfig)

	for i := 0; i < len(stmts); i++ {

		transformConfig.Index = i
		stmt := stmts[i]
		stmt.SetDBType(sourceDBType)
		stmt.SetTargetDBType(targetDBType)

		for _, v := range visitors {

			newStmt := stmts[i]
			if stmt != newStmt {
				i = i - 1
				break
			}

			visitor.Accept(v, stmt)
			transformResult.AddChanges(v.Changes())
			transformResult.AddWarnnings(v.Warnnings())
			transformResult.AddErrors(v.Errors())
		}
	}

	// output
	var builders strings.Builder
	for i, stmt := range stmts {

		if i != len(stmts)-1 {
			stmt.SetAfterSemi(true)
		}

		var builder strings.Builder
		outputVisitor := sql.CreateASTOutputVisitor(&builder, targetDBType, transformConfig.SQLOutputConfig)
		visitor.Accept(outputVisitor, stmt)

		if builder.Len() > 0 {
			if i != len(stmts)-1 {
				outputVisitor.WriteLn()
			}
			builders.WriteString(builder.String())
		}

		transformResult.AddObjectResult(result.NewSQLObjectResult(stmt.ObjectType(), builder.String()))
	}
	transformResult.TargetSQL = builders.String()

	// config.stmtList.clear()

	return transformResult
}

func CreateASTTranslateVisitors(transformConfig *config.SQLTransformConfig) []visitor2.ISQLASTTransformVisitor {
	if transformConfig == nil {
		panic("Translate config is null.")
	}
	sourceDBType := transformConfig.SourceDBType
	targetDBType := transformConfig.TargetDBType
	targetDBVersion := transformConfig.TargetDBVersion
	if sourceDBType == "" {
		panic("source database type is null.")
	}
	if targetDBType == "" {
		panic("target database type is null.")
	}
	if targetDBVersion == "" {
		panic("target database version is null.")
	}

	var visitors []visitor2.ISQLASTTransformVisitor
	switch sourceDBType {
	case db.Oracle:
		switch targetDBType {
		case db.Oracle:
			// visitors.add(newOracleSQLRemovePropertyASTVisitor(config))
			break
		// case EDB:
		// 	visitors.add(newOracleSQLRemovePropertyASTVisitor(config))
		// 	visitors.add(newOracle2EDBMixRemoveDoubleQuotesASTVisitor(config))
		// switch (targetDBVersion) {
		// 	case VERSION_9_6:
		// 		visitors.add(new
		// 		Oracle2EDBVersion9_6ASTTransformVisitor(config))
		// 		break
		// 	case VERSION_10:
		// 		visitors.add(new
		// 		Oracle2EDBVersion10ASTTransformVisitor(config))
		// 		break
		// 	}
		// break
		case db.PostgreSQL:
			// visitors.add(newOracleSQLRemovePropertyASTVisitor(config))
			// visitors.add(newOracle2EDBMixRemoveDoubleQuotesASTVisitor(config))
			switch targetDBVersion {
			case postgresql.Version_9_6:
				// visitors.add(newOracle2PostgreSQLVersion9_6ASTTransformVisitor(config))
				break
			case postgresql.Version_10:
				// visitors.add(newOracle2PostgreSQLVersion10ASTTransformVisitor(config))
				break
			}
			break
		case db.MySQL:
			switch targetDBVersion {
			case mysql.Version_5_6:
				visitors = append(visitors, mysql2.NewOracle2MySQLV5_6ASTTransformVisitor(transformConfig))
			case mysql.Version_5_7:
				visitors = append(visitors, mysql2.NewOracle2MySQLV5_7ASTTransformVisitor(transformConfig))
			// visitors.add(newSQLWithClauseSubQueryTranslateAndRemoveASTVisitor(config))
			// visitors.add(newSQLCreateViewSubQueryTableRefToCreateViewASTVisitor(config))
			case mysql.Version_8_0:
				visitors = append(visitors, mysql2.NewOracle2MySQLV8_0ASTTransformVisitor(transformConfig))
			default:
				break
			}
			// visitors.add(newSQLOuterJoinToJoinASTVisitor(config))
			// visitors.add(newSQLRowNumToLimitASTVisitor(config))
			// visitors.add(newOracleSQLIntersectOrMinusToJoinASTVisitor(config))
			// visitors.add(newSQLBindVarToVarASTVisitor(config))

			switch targetDBVersion {
			case mysql.Version_5_6:
				// visitors.add(newOracle2MySQLVersion5_6ASTTransformVisitor(config))
				break
			case mysql.Version_5_7:
				// visitors.add(newOracle2MySQLVersion5_7ASTTransformVisitor(config))
				break
			case mysql.Version_8_0:
				// visitors.add(newOracle2MySQLVersion8_0ASTTransformVisitor(config))
				break
			default:
				break
			}

			// visitors.add(newSQLRenameColumnASTVisitor(config))
			// visitors.add(newSQLRenameObjectNameASTVisitor(config))
			//
			// visitors.add(newSQLDefaultClauseToTriggerAndRemoveASTVisitor(config))

			if transformConfig.IsRemoveSchema {
				// visitors.add(newSQLRemoveSchemaASTVisitor(config))
			}

			// visitors.add(newSQLOptimizationASTVisitor(config))
			// visitors.add(newSQLAddReverseQuotesASTVisitor(config))
			break
		}
		break

	default:
		panic("")
	}
	return visitors
}
