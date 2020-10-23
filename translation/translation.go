package translation

import (
	"gumihoy.com/sql/basic/ast"
	"gumihoy.com/sql/basic/ast/expr"
	"gumihoy.com/sql/basic/ast/expr/datatype"
	"gumihoy.com/sql/db"
	"gumihoy.com/sql/factory"
	"strings"
)

type DoubleQuoteActionType int

const (
	NONE DoubleQuoteActionType = iota
)

func (x *SQLTransformConfig) FindTableMapping(tableName expr.ISQLName) *TableMapping {
	lowerFullName := tableName.StringName()
	lowerName := tableName.SimpleStringName()

	mapping := x.TableMappings[lowerFullName]
	if mapping != nil {
		return mapping
	}

	return x.TableMappings[lowerName]
}

func (x *SQLTransformConfig) AddTableMapping(tableMapping *TableMapping) {
	if tableMapping == nil || len(tableMapping.name) == 0 {
		return
	}
	x.TableMappings[tableMapping.name] = tableMapping
}

type SQLTransformConfig struct {
	SourceDBType    db.DBType
	TargetDBType    db.DBType
	TargetDBVersion db.Version

	/**
	 * 表名、字段名移除引号
	 */
	doubleQuoteAction DoubleQuoteActionType

	/**
	 * 对象名(表、视图、PL/SQL等)、字段名移除schema
	 */
	IsRemoveSchema bool

	// Map<String, TableMapping> tableMappings = new LinkedHashMap<>()
	TableMappings map[string]*TableMapping

	// 当前遍历下标
	Index int
	// 当前 sql 的 stmt LIST
	// public List<ISQLObject> stmtList = new ArrayList<>()
	Stmts []ast.ISQLObject

	/**
	 * CREATE index ON table 映射关系
	 */
	// 	private final ConcurrentHashMap<String, String> INDEX_TABLE_MAP = new ConcurrentHashMap<>()
	//
	// 	public String getIndexTable(String index) {
	// 	return INDEX_TABLE_MAP.get(index)
	// }
	//
	// 	public void setIndexTable(String index, String table) {
	// 	INDEX_TABLE_MAP.put(index, table)
	// }
}

func NewSQLTransformConfig() *SQLTransformConfig {
	var x SQLTransformConfig

	return &x
}

type TableMapping struct {
	owner       string
	name        string
	targetOwner string
	targetName  string
	/**
	 * 修改字段
	 */
	// public final Set<ColumnMapping> columnMappings = new LinkedHashSet<>()
	columnMappings []*ColumnMapping
	/**
	 * 移除字段
	 */
	// public final Set<String> removeColumns = new LinkedHashSet<>()
	removeColumns []string
	/**
	 * 添加字段
	 */
	// public final Set<ColumnMapping> addColumnMappings = new LinkedHashSet<>()
	addColumnMappings []*ColumnMapping
}

func NewTableMapping(name string, targetName string, columnMappings ...*ColumnMapping) *TableMapping {
	var x TableMapping
	x.name = name
	x.targetName = targetName

	for _, columnMapping := range columnMappings {
		x.AddColumnMapping(columnMapping)
	}

	return &x
}

func (x *TableMapping) findColumnMapping(columnName string) *ColumnMapping {
	for _, columnMapping := range x.columnMappings {
		if strings.EqualFold(columnMapping.name, columnName) {
			return columnMapping
		}
	}

	return nil
}

func (x *TableMapping) AddColumnMapping(columnMapping *ColumnMapping) {
	if columnMapping == nil {
		return
	}
	// x.columnMappings.add(columnMapping)
}

func (x *TableMapping) IsRemoveColumn(columnName string) bool {
	for _, removeColumn := range x.removeColumns {
		if strings.EqualFold(removeColumn, columnName) {
			return true
		}
	}

	return false
}

func (x *TableMapping) AddRemoveColumn(columnName string) {
	x.removeColumns = append(x.removeColumns, columnName)
}

func (x *TableMapping) addAddColumnMapping(columnMapping *ColumnMapping) {
	if columnMapping == nil {
		return
	}
	x.addColumnMappings = append(x.addColumnMappings, columnMapping)
}

type ColumnMapping struct {
	name         string
	targetName   string
	dataType     datatype.ISQLDataType
	defaultValue expr.ISQLExpr
}

func NewColumnMappingWithNameAndTargetName(name string, targetName string) *ColumnMapping {
	var x ColumnMapping
	x.name = name
	x.targetName = targetName
	return &x
}

func NewColumnMappingWith(name string, targetName string, dataType datatype.ISQLDataType) *ColumnMapping {
	var x ColumnMapping
	x.name = name
	x.targetName = targetName
	x.dataType = dataType
	return &x
}

func NewColumnMapping(name string, targetName string, dataType datatype.ISQLDataType, defaultValue expr.ISQLExpr) *ColumnMapping {
	var x ColumnMapping
	x.name = name
	x.targetName = targetName
	x.dataType = dataType
	x.defaultValue = defaultValue
	return &x
}

func OracleToMySQL(sql string) *SQLTransformResult {
	config := NewSQLTransformConfig()
	return OracleToMySQLWithConfig(sql, config)
}

func OracleToMySQLWithConfig(sql string, config *SQLTransformConfig) *SQLTransformResult {
	if config == nil {
		config = NewSQLTransformConfig()
	}
	config.SourceDBType = db.Oracle
	config.TargetDBType = db.MySQL
	return Translate(sql, config)
}

func OracleToEDB(sql string) *SQLTransformResult {
	config := NewSQLTransformConfig()
	return OracleToEDBWithConfig(sql, config)
}

func OracleToEDBWithConfig(sql string, config *SQLTransformConfig) *SQLTransformResult {
	if config == nil {
		config = NewSQLTransformConfig()
	}
	config.SourceDBType = db.Oracle
	config.TargetDBType = db.EDB
	return Translate(sql, config)
}

func oracleToPostgreSQL(sql string) *SQLTransformResult {
	config := NewSQLTransformConfig()
	return OracleToPostgreSQLWithConfig(sql, config)
}

func OracleToPostgreSQLWithConfig(sql string, config *SQLTransformConfig) *SQLTransformResult {
	if config == nil {
		config = NewSQLTransformConfig()
	}
	config.SourceDBType = db.Oracle
	config.TargetDBType = db.PostgreSQL
	return Translate(sql, config)
}

func MySQLToOracle(sql string) *SQLTransformResult {
	config := NewSQLTransformConfig()
	return MySQLToOracleWithConfig(sql, config)
}

func MySQLToOracleWithConfig(sql string, config *SQLTransformConfig) *SQLTransformResult {
	if config == nil {
		config = NewSQLTransformConfig()
	}
	config.SourceDBType = db.MySQL
	config.TargetDBType = db.Oracle
	return Translate(sql, config)
}

func Translate(sql string, config *SQLTransformConfig) *SQLTransformResult {

	sourceDBType := config.SourceDBType
	targetDBType := config.TargetDBType

	result := NewSQLTransformResult(sql)

	stmts := factory.ParseStatementsByDBType(sql, sourceDBType)

	config.Stmts = stmts
	// buffers := newStringBuilder()

	// outputConfig := newSQLOutputConfig()

	// transformVisitors := factory.CreateASTTranslateVisitors(config)

	for i := 0; i < len(stmts); i++ {

		config.Index = i
		stmt := stmts[i]
		stmt.SetDBType(sourceDBType)
		stmt.SetTargetDBType(targetDBType)

		// for _, visitor := range transformVisitors {
		//
		// 	newStmt := stmts[i]
		// 	if stmt != newStmt {
		// 		i = i - 1
		// 		break
		// 	}

		// accept.Accept(visitor, stmt)

		// result.changes.addAll(visitor.getChanges())
		// result.warnnings.addAll(visitor.getWarnnings())
		// result.errors.addAll(visitor.getErrors())
		// }
	}

	// output
	for i := 0; i < len(stmts); i++ {
		stmt := stmts[i]

		if i != len(stmts)-1 {
			stmt.SetAfterSemi(true)
		}

		// var builder strings.Builder

		// outputVisitor := factory.CreateASTOutputVisitor(&builder, outputConfig, sourceDBType, targetDBType)
		// SQLUtils.outputVisitor(stmt, outputVisitor)

		// if (buffer.length() > 0) {
		// 	if (i != stmtList.size()-1) {
		// 		outputVisitor.println()
		// 	}
		// 	buffers.append(buffer)
		// }

		// result.addResult(SQLTransformResult.SQLResult.of(stmt.getObjectType(), buffer.toString()))
	}

	// result.targetSql = buffers.toString()

	// config.stmtList.clear()

	return result
}

type SQLTransformResult struct {
	sourceSql string
	targetSql string
	// results []*SQLResult

	//
	// public final Set<SQLTransformError> errors = new HashSet<SQLTransformError>()
	// public final Set<SQLTransformChange> changes = new HashSet<SQLTransformChange>()
	// public final Set<SQLTransformWarnning> warnnings = new HashSet<SQLTransformWarnning>()

	createTableCount int

	insertCount            int
	deleteCount            int
	selectCount            int
	updateCount            int
	createTypeCount        int
	createPackageCount     int
	createPackageBodyCount int
	createFunctionCount    int
	createProcedureCount   int
}

func NewSQLTransformResult(sourceSql string) *SQLTransformResult {
	var x SQLTransformResult
	x.sourceSql = sourceSql
	return &x
}

func (x *SQLTransformResult) AddResult(result *SQLResult) {
	if result == nil {
		return
	}
	// results.add(result)
}

type SQLResult struct {
	objectType db.ObjectType
	targetSql  string
}

func NewSQLResult(objectType db.ObjectType, targetSql string) *SQLResult {
	var x SQLResult
	x.objectType = objectType
	x.targetSql = targetSql
	return &x
}
