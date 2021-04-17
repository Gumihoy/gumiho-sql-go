# SQL Go

## String SQL to Abstract Syntax Tree

## Format String SQL 

``` go
func TestFormat(t *testing.T) {
	sourceSQL := "select id, name from dual where id = ?;"
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(formatSQL)
}
```

## Normalize String SQL
``` go
func TestFormat(t *testing.T) {
	sourceSQL := "select id, name from dual where id = 1;"
	normalizeSQL := sql.NormalizeSQL(sourceSQL, db.MySQL)
	fmt.Println(normalizeSQL)
}
```

## Translation
### Oracle To MySQL

#### DataType

| Oracle数据类型 | 转换规则 |
| :-----------: | :----------- |
| 字符型     |
| CHAR/CHARACTER(x) | 0 CHAR <br/> 0 < x <= 255 CHAR(x) <br/> 255 < x <= 2000 VARCHAR(x)  |
| NCHAR (x) | 0 NCHAR <br/> 0<x<=255  NCHAR(x) <br/> 255<x <=2000 NVARCHAR(x)  |
| VARCHAR2(x) | 0<x<=2000 VARCHAR(x) <br/> 2000<x<=4000 TEXT  |
| VARCHAR(x) | 0<x<=2000 VARCHAR(x) <br/> 2000<x<=4000 TEXT  |
| RAW(x) | 1<x<=127 BINARY(x * 2) <br/> 127<x<=2000 VARBINARY(x * 2)   |
| NVARCHAR2(x) | NVARCHAR(x)  |
| LONG | LONGTEXT |
| LONGRAW | LONGBLOB |
| BLOB | LONGBLOB |
| CLOB | LONGTEXT |
| NCLOB | LONGTEXT |
| 数字型    |
| BINARY_FLOAT | FLOAT  |
| BINARY_DOUBLE | DOUBLE |
| NUMBER | DECIAML(40,30) |
| NUMBER(a) | DECIMAL(a) |
| NUMBER(*/a, b) | DECIMAL(38/a,b) |
| DEC | 同NUMBER规则
| DECIMAL | 同NUMBER规则 |
| NUMERIC | 同NUMBER规则 |
| FLOAT | DOUBLE |
| FLOAT(x) | 0<=x<=24 FLOAT <br/> x>24 DOUBLE
| INT | DECIMAL (38) |
| INTEGER | DECIMAL (38) |
| SMALLINT | DECIMAL (38) |
| REAL | DOUBLE |
| 日期型 |
| DATE | DATETIME
| TIMESTAMP (x) | DATATIME(x),x默认是6
| TIMESTAMP (x) WITH (LOCAL) TIME ZONE | DATATIME(x),x默认是6
| INTERVAL YEAR (x) TO MONTH | 不支持
| INTERVAL DAY (x) TO SECOND (y) | 不支持
| 其他型 |
| BFILE | 不支持 |
| ROWID | CHAR(18) |
| UROWID | VARCHAR(4000) |
| UROWID(x) | VARCHAR(x) |
| UDT | 不支持 |
| SYS.* | 不支持 |
| XMLTYPE | 不支持 |
| SDO_* | 不支持 |
| ORD* | 不支持 |
| SI_* | 不支持 |

### Oracle To EDB
#### DataType
Oracle数据类型 | 转换规则 |
:-----------: | :----------- |
| 字符型     |
CHAR (x CHAR) <br/>CHAR(x BYTE) | CHAR(x) |
NCHAR(x CHAR) <br/>NCHAR(x BYTE) | NCHAR(x)
VARCHAR(x CHAR)<br/>VARCHAR(x BYTE) | VARCHAR(x)
RAW(x) |  默认BYTEA，设置参数后转OID
NVARCHAR2 (x [CHAR | BYTE]) | NVARCHAR2(x)
CHARACTER (x) | CHARACTER (x)
LONG | TEXT
BLOB | 默认BYTEA，设置参数后转OID
CLOB | TEXT
NCLOB | TEXT
LONG RAW | 默认BYTEA，设置参数后转OID
| 数字型 |
PLS_INTEGER | INTEGER
BINARY_INTEGER | INTEGER
SIMPLE_INTEGER | INTEGER
BINARY_FLOAT | DOUBLE_PRECISION
SIMPLE_FLOAT | REAL
BINARY_DOUBLE | DOUBLE_PRECISION
SIMPLE_DOUBLE | DOUBLE_PRECISION
NUMBER (x) | NUMBER(x)
DEC (x) | DECIMAL (x)
DECIMAL (x) | DECIMAL (x)
NUMERIC (x) | DECIMAL (x)
DOUBLE PRECISION | DOUBLE PRECISION
FLOAT(x) | 1<=x<=24 转换为 REAL <br/> 5<=x<=53 转换为 DOUBLE PRECISION
INT | INTEGER
INTEGER | INTEGER
SMALLINT | INTEGER
REAL | REAL
| 日期型 |
DATE | 目前转为TIMESTAMP(0)，等Polar-O 10-30版本修复date精度问题，不再转换
TIME (x) [WITH/WITHOUT TIME ZONE] | TIME (x) [WITH/WITHOUT TIME ZONE]
TIMESTAMP (x) | TIMESTAMP (x)
TIMESTAMP (x) WITH TIME ZONE | TIMESTAMP (x) WITH TIME ZONE
TIMESTAMP (x) WITH LOCAL TIME ZONE | TIMESTAMP (x)，数据迁移时需要考虑时区问题
INTERVAL YEAR (x) TO MONTH | INTERVAL YEAR TO MONTH
INTERVAL DAY (x) TO SECOND (y) | INTERVAL DAY TO SECOND(y)
TIME_UNCONSTRAINED | TIME(9)
TIME_TZ_UNCONSTRAINED | TIME(9) WITH TIME ZONE
TIMESTAMP_UNCONSTRAINED | TIMESTAMP(9)
TIMESTAMP_TZ_UNCONSTRAINED | TIMESTAMP(9) WITH TIME ZONE
YMINTERVAL_UNCONSTRAINED | INTERVAL YEAR TO MONTH
DSINTERVAL_UNCONSTRAINED | INTERVAL DAY TO SECOND (9)
TIMESTAMP_LTZ_UNCONSTRAINED | TIMESTAMP(9) WITH TIME ZONE
| 其他 | 
BFILE | 不支持
ROWID | CHAR(18)
UROWID | VARCHAR2(4000)
UROWID(x) | VARCHAR2(x)
BOOLEAN | BOOLEAN
UDT | UDT （自定义TYPE）
SYS.* | 不支持
XMLTYPE | XMLTYPE/XML
SDO_* | 不支持
ORD* | 不支持
SI_* | 不支持

### Oracle To Postgre|
#### DataType

Oracle数据类型 | 转换规则 |
:-----------: | :----------- |
| 字符型     |
CHAR(x CHAR/BYTE)	|	CHAR(x)
NCHAR(x CHAR/BYTE)	|	CHAR(x)
CHARACTER(x CHAR/BYTE)	|	CHARACTER(x)
VARCHAR2(x CHAR/BYTE)	|	VARCHAR(x)
VARCHAR(x CHAR/BYTE)	|	VARCHAR(x)
NVARCHAR2(x CHAR/BYTE)	|	VARCHAR(x)
LONG	|	TEXT
CLOB	|	TEXT
NCLOB	|	TEXT
BLOB	|	BYTEA
LONG RAW	|	BYTEA
RAW(x)	|	BYTEA （BYTEA不能指定精度）
| 数值类型	|	
PLS_INTEGER	|	INTEGER
BINARY_INTEGER	|	INTEGER
SIMPLE_INTEGER	|	INTEGER
INT	|	INTEGER
INTEGER[(x)]	|	INTEGER。PLSQL下INTEGER类型可以带精度，转换时会去掉该精度
SMALLINT	|	INTEGER
BINARY_FLOAT	|	REAL
SIMPLE_FLOAT	|	REAL
BINARY_DOUBLE	|	DOUBLE_PRECISION
SIMPLE_DOUBLE	|	DOUBLE_PRECISION
NUMBER(x)	|	DECIMAL(x)
DEC(x)	|	DECIMAL(x)
DECIMAL(x)	|	DECIMAL(x)
NUMERIC(x)	|	DECIMAL(x)，如果精度为*，则转为38
DOUBLE PRECISION	|	DOUBLE_PRECISION
FLOAT(x)	|	1<=x<=24 转换为 REAL，其他转换为 DOUBLE PRECISION
REAL	|	REAL
| 时间日期类型 |		
DATE	|	TIMESTAMP(0)
TIME(x) [WITH/WITHOUT TIME ZONE]	|	TIME(x) [WITH/WITHOUT TIME ZONE]
TIMESTAMP(x)	|	TIMESTAMP(x)
TIMESTAMP(x) WITH [LOCAL] TIME ZONE	|	TIMESTAMP(x) WITH TIME ZONE
INTERVAL YEAR(x) TO MONTH	|	INTERVAL YEAR TO MONTH
INTERVAL DAY(x) TO SECOND(y)	|	INTERVAL DAY TO SECOND(y)
TIME_UNCONSTRAINED	|	TIME(9)
TIME_TZ_UNCONSTRAINED	|	TIME(9) WITH TIME ZONE
TIMESTAMP_UNCONSTRAINED	|	TIMESTAMP(9)
TIMESTAMP_TZ_UNCONSTRAINED	|	TIMESTAMP(9) WITH TIME ZONE
YMINTERVAL_UNCONSTRAINED	|	INTERVAL YEAR TO MONTH
DSINTERVAL_UNCONSTRAINED	|	INTERVAL DAY TO SECOND(9)
TIMESTAMP_LTZ_UNCONSTRAINED	|	TIMESTAMP(9) WITH TIME ZONE
| 其他类型		
BFILE	|	不支持
ROWID	|	CHAR(18)
UROWID(x)	|	如果带精度则转换VARCHAR(x)，否则VARCHAR(4000)
UROWID(x)	|	VARCHAR(x)
BOOLEAN	|	不支持
UDT	|	UDT （自定义TYPE）
SYS.*	|	不支持
XMLTYPE	|	不支持
SDO_*	|	不支持
ORD*	|	不支持
SI_*	|	不支持
