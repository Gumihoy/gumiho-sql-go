package config

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
)

type DoubleQuoteActionType int

const (
	NONE DoubleQuoteActionType = iota
)

/**
 *
 */
type SQLTransformConfig struct {
	*config.SQLOutputConfig

	SourceDBType    db.Type
	TargetDBType    db.Type
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
	Stmts []statement.ISQLStatement

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
	x := new(SQLTransformConfig)
	x.SQLOutputConfig = config.NewOutputConfig()
	return x
}

func (x *SQLTransformConfig) FindTableMapping(tableName expr.ISQLName) *TableMapping {
	lowerFullName := tableName.StringName()
	lowerName := tableName.StringName()

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

/**
 *
 */
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
