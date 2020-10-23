package basic

import "strings"

// https://dev.mysql.com/doc/refman/8.0/en/keywords.html
// https://docs.oracle.com/en/database/oracle/oracle-database/19/sqlrf/Oracle-SQL-Reserved-Words.html#GUID-55C49D1E-BE08-4C50-A9DD-8593EB925612
type Keyword struct {
	Lower string
	Upper string
}
func NewKeyword(value string) *Keyword {
	var x Keyword
	x.Lower, x.Upper = value, value
	return &x
}

func NewKeywordCaseSensitive(value string) *Keyword {
	var x Keyword
	x.Lower, x.Upper = strings.ToLower(value), strings.ToUpper(value)
	return &x
}


var (
	// A

	ACCESSIBLE      = NewKeywordCaseSensitive("ACCESSIBLE")
	ACCOUNT         = Keyword{"", ""}
	ACTION          = Keyword{"", ""}
	ACTIVE          = Keyword{"", ""}
	ADD             = Keyword{"", ""}
	ADMIN           = Keyword{"", ""}
	AFTER           = Keyword{"", ""}
	AGAINST         = Keyword{"", ""}
	AGGREGATE       = Keyword{"", ""}
	ALGORITHM       = Keyword{"", ""}
	ALL             = Keyword{"", ""}
	ALTER           = Keyword{"", ""}
	ALWAYS          = Keyword{"", ""}
	ANALYSE         = Keyword{"", ""}
	ANALYZE         = Keyword{"", ""}
	AND             = Keyword{"", ""}
	ANY             = Keyword{"", ""}
	ARRAY           = Keyword{"", ""}
	AS              = Keyword{"", ""}
	ASC             = Keyword{"", ""}
	ASCII           = Keyword{"", ""}
	ASENSITIVE      = Keyword{"", ""}
	AT              = Keyword{"", ""}
	ATTRIBUTE       = Keyword{"", ""}
	AUTOEXTEND_SIZE = Keyword{"", ""}
	AUTO_INCREMENT  = Keyword{"", ""}
	AVG             = Keyword{"", ""}
	AVG_ROW_LENGTH  = Keyword{"", ""}

	// B

	BACKUP  = Keyword{"", ""}
	BEFORE  = Keyword{"", ""}
	BEGIN   = Keyword{"", ""}
	BETWEEN = Keyword{"", ""}
	BIGINT  = Keyword{"", ""}
	BINARY  = Keyword{"", ""}
	BINLOG  = Keyword{"", ""}
	BIT     = Keyword{"", ""}
	BLOB    = Keyword{"", ""}
	BLOCK   = Keyword{"", ""}
	BOOL    = Keyword{"", ""}
	BOOLEAN = Keyword{"", ""}
	BOTH    = Keyword{"", ""}
	BTREE   = Keyword{"", ""}
	BUCKETS = Keyword{"", ""}
	BY      = Keyword{"", ""}
	BYTE    = Keyword{"", ""}

	// C
	CACHE              = Keyword{"", ""}
	CALL               = Keyword{"", ""}
	CASCADE            = Keyword{"", ""}
	CASCADED           = Keyword{"", ""}
	CASE               = Keyword{"", ""}
	CATALOG_NAME       = Keyword{"", ""}
	CHAIN              = Keyword{"", ""}
	CHANGE             = Keyword{"", ""}
	CHANGED            = Keyword{"", ""}
	CHANNEL            = Keyword{"", ""}
	CHAR               = Keyword{"", ""}
	CHARACTER          = Keyword{"", ""}
	CHARSET            = Keyword{"", ""}
	CHECK              = Keyword{"", ""}
	CHECKSUM           = Keyword{"", ""}
	CIPHER             = Keyword{"", ""}
	CLASS_ORIGIN       = Keyword{"", ""}
	CLIENT             = Keyword{"", ""}
	CLONE              = Keyword{"", ""}
	CLOSE              = Keyword{"", ""}
	COALESCE           = Keyword{"", ""}
	CODE               = Keyword{"", ""}
	COLLATE            = Keyword{"", ""}
	COLLATION          = Keyword{"", ""}
	COLUMN             = Keyword{"", ""}
	COLUMNS            = Keyword{"", ""}
	COLUMN_FORMAT      = Keyword{"", ""}
	COLUMN_NAME        = Keyword{"", ""}
	COMMENT            = Keyword{"", ""}
	COMMIT             = Keyword{"", ""}
	COMMITTED          = Keyword{"", ""}
	COMPACT            = Keyword{"", ""}
	COMPLETION         = Keyword{"", ""}
	COMPONENT          = Keyword{"", ""}
	COMPRESSED         = Keyword{"", ""}
	COMPRESSION        = Keyword{"", ""}
	CONCURRENT         = Keyword{"", ""}
	CONDITION          = Keyword{"", ""}
	CONNECTION         = Keyword{"", ""}
	CONSISTENT         = Keyword{"", ""}
	CONSTRAINT         = Keyword{"", ""}
	CONSTRAINT_CATALOG = Keyword{"", ""}
	CONSTRAINT_NAME    = Keyword{"", ""}
	CONSTRAINT_SCHEMA  = Keyword{"", ""}
	CONTAINS           = Keyword{"", ""}
	CONTEXT            = Keyword{"", ""}
	CONTINUE           = Keyword{"", ""}
	CONVERT            = Keyword{"", ""}
	CPU                = Keyword{"", ""}
	CREATE             = Keyword{"", ""}
	CROSS              = Keyword{"", ""}
	CUBE               = Keyword{"", ""}
	CUME_DIST          = Keyword{"", ""}
	CURRENT            = Keyword{"", ""}
	CURRENT_DATE       = Keyword{"", ""}
	CURRENT_TIME       = Keyword{"", ""}
	CURRENT_TIMESTAMP  = Keyword{"", ""}
	CURRENT_USER       = Keyword{"", ""}
	CURSOR             = Keyword{"", ""}
	CURSOR_NAME        = Keyword{"", ""}

	// D

	DATA            = Keyword{"", ""}
	DATABASE        = Keyword{"", ""}
	DATABASES       = Keyword{"", ""}
	DATAFILE        = Keyword{"", ""}
	DATE            = Keyword{"", ""}
	DATETIME        = Keyword{"", ""}
	DAY             = Keyword{"", ""}
	DAY_HOUR        = Keyword{"", ""}
	DAY_MICROSECOND = Keyword{"", ""}
	DAY_MINUTE      = Keyword{"", ""}
	DAY_SECOND      = Keyword{"", ""}

	DEALLOCATE      = Keyword{"", ""}
	DEC             = Keyword{"", ""}
	DECIMAL         = Keyword{"", ""}
	DECLARE         = Keyword{"", ""}
	DEFAULT         = Keyword{"", ""}
	DEFAULT_AUTH    = Keyword{"", ""}
	DEFINER         = Keyword{"", ""}
	DEFINITION      = Keyword{"", ""}
	DELAYED         = Keyword{"", ""}
	DELAY_KEY_WRITE = Keyword{"", ""}
	DELETE          = Keyword{"", ""}
	DENSE_RANK      = Keyword{"", ""}
	DESC            = Keyword{"", ""}
	DESCRIBE        = Keyword{"", ""}
	DESCRIPTION     = Keyword{"", ""}
	DES_KEY_FILE    = Keyword{"", ""}
	DETERMINISTIC   = Keyword{"", ""}

	DIAGNOSTICS = Keyword{"", ""}
	DIRECTORY   = Keyword{"", ""}
	DISABLE     = Keyword{"", ""}
	DISCARD     = Keyword{"", ""}
	DISK        = Keyword{"", ""}
	DISTINCT    = Keyword{"", ""}
	DISTINCTROW = Keyword{"", ""}
	DIV         = Keyword{"", ""}

	DO     = Keyword{"", ""}
	DOUBLE = Keyword{"", ""}

	DROP = Keyword{"", ""}

	DUAL      = Keyword{"", ""}
	DUMPFILE  = Keyword{"", ""}
	DUPLICATE = Keyword{"", ""}
	DYNAMIC   = Keyword{"", ""}

	// E

	EACH = Keyword{"", ""}

	ELSE   = Keyword{"", ""}
	ELSEIF = Keyword{"", ""}

	EMPTY = Keyword{"", ""}

	ENABLE           = Keyword{"", ""}
	ENCLOSED         = Keyword{"", ""}
	ENCRYPTION       = Keyword{"", ""}
	END              = Keyword{"", ""}
	ENDS             = Keyword{"", ""}
	ENFORCED         = Keyword{"", ""}
	ENGINE           = Keyword{"", ""}
	ENGINES          = Keyword{"", ""}
	ENGINE_ATTRIBUTE = Keyword{"", ""}
	ENUM             = Keyword{"", ""}

	ERROR  = Keyword{"", ""}
	ERRORS = Keyword{"", ""}

	ESCAPE  = Keyword{"", ""}
	ESCAPED = Keyword{"", ""}

	EVENT  = Keyword{"", ""}
	EVENTS = Keyword{"", ""}
	EVERY  = Keyword{"", ""}

	EXCEPT   = Keyword{"", ""}
	EXCHANGE = Keyword{"", ""}
	EXCLUDE  = Keyword{"", ""}
	EXECUTE  = Keyword{"", ""}

	EXISTS      = Keyword{"", ""}
	EXIT        = Keyword{"", ""}
	EXPANSION   = Keyword{"", ""}
	EXPIRE      = Keyword{"", ""}
	EXPLAIN     = Keyword{"", ""}
	EXPORT      = Keyword{"", ""}
	EXTENDED    = Keyword{"", ""}
	EXTENT_SIZE = Keyword{"", ""}

	// F

	FAILED_LOGIN_ATTEMPTS = Keyword{"", ""}
	FALSE                 = Keyword{"", ""}
	FAST                  = Keyword{"", ""}
	FAULTS                = Keyword{"", ""}

	FETCH           = Keyword{"", ""}
	FIELDS          = Keyword{"", ""}
	FILE            = Keyword{"", ""}
	FILE_BLOCK_SIZE = Keyword{"", ""}
	FILTER          = Keyword{"", ""}
	FIRST           = Keyword{"", ""}
	FIRST_VALUE     = Keyword{"", ""}
	FIXED           = Keyword{"", ""}

	FLOAT  = Keyword{"", ""}
	FLOAT4 = Keyword{"", ""}
	FLOAT8 = Keyword{"", ""}
	FLUSH  = Keyword{"", ""}

	FOLLOWING = Keyword{"", ""}
	FOLLOWS   = Keyword{"", ""}
	FOR       = Keyword{"", ""}
	FORCE     = Keyword{"", ""}
	FOREIGN   = Keyword{"", ""}
	FORMAT    = Keyword{"", ""}
	FOUND     = Keyword{"", ""}

	FROM = NewKeywordCaseSensitive("FROM")

	FULL     = NewKeywordCaseSensitive("FULL")
	FULLTEXT = Keyword{"", ""}
	FUNCTION = Keyword{"", ""}

	// G

	GENERAL               = Keyword{"", ""}
	GENERATED             = Keyword{"", ""}
	GEOMCOLLECTION        = Keyword{"", ""}
	GEOMETRY              = Keyword{"", ""}
	GEOMETRYCOLLECTION    = Keyword{"", ""}
	GET                   = Keyword{"", ""}
	GET_FORMAT            = Keyword{"", ""}
	GET_MASTER_PUBLIC_KEY = Keyword{"", ""}
	GLOBAL                = Keyword{"", ""}

	GRANT             = Keyword{"", ""}
	GRANTS            = Keyword{"", ""}
	GROUP             = Keyword{"", ""}
	GROUPING          = Keyword{"", ""}
	GROUPS            = Keyword{"", ""}
	GROUP_REPLICATION = Keyword{"", ""}

	// H

	HANDLER = Keyword{"", ""}

	HASH   = Keyword{"", ""}
	HAVING = Keyword{"", ""}

	HELP = Keyword{"", ""}

	HIGH_PRIORITY = Keyword{"", ""}
	HISTOGRAM     = Keyword{"", ""}
	HISTORY       = Keyword{"", ""}

	HOST             = Keyword{"", ""}
	HOSTS            = Keyword{"", ""}
	HOUR             = Keyword{"", ""}
	HOUR_MICROSECOND = Keyword{"", ""}
	HOUR_MINUTE      = Keyword{"", ""}
	HOUR_SECOND      = Keyword{"", ""}

	// I

	IDENTIFIED = Keyword{"", ""}

	IF = Keyword{"", ""}

	IGNORE            = Keyword{"", ""}
	IGNORE_SERVER_IDS = Keyword{"", ""}

	IMPORT        = Keyword{"", ""}
	IN            = Keyword{"", ""}
	INACTIVE      = Keyword{"", ""}
	INDEX         = Keyword{"", ""}
	INDEXESV      = Keyword{"", ""}
	INFILE        = Keyword{"", ""}
	INITIAL_SIZE  = Keyword{"", ""}
	INNER         = Keyword{"", ""}
	INOUT         = Keyword{"", ""}
	INSENSITIVE   = Keyword{"", ""}
	INSERT        = Keyword{"", ""}
	INSERT_METHOD = Keyword{"", ""}
	INSTALL       = Keyword{"", ""}
	INSTANCE      = Keyword{"", ""}
	INT           = Keyword{"", ""}
	INT1          = Keyword{"", ""}
	INT2          = Keyword{"", ""}
	INT3          = Keyword{"", ""}
	INT4          = Keyword{"", ""}
	INT8          = Keyword{"", ""}

	INTEGER  = Keyword{"", ""}
	INTERVAL = Keyword{"", ""}

	INTO      = Keyword{"", ""}
	INVISIBLE = Keyword{"", ""}
	INVOKER   = Keyword{"", ""}

	IO              = Keyword{"", ""}
	IO_AFTER_GTIDS  = Keyword{"", ""}
	IO_BEFORE_GTIDS = Keyword{"", ""}
	IO_THREAD       = Keyword{"", ""}
	IPC             = Keyword{"", ""}
	IS              = Keyword{"", ""}
	ISOLATION       = Keyword{"", ""}
	ISSUER          = Keyword{"", ""}
	ITERATE         = Keyword{"", ""}

	// J

	JOIN       = Keyword{"", ""}
	JSON       = Keyword{"", ""}
	JSON_TABLE = Keyword{"", ""}
	JSON_VALUE = Keyword{"", ""}

	// K

	KEY            = Keyword{"", ""}
	KEYS           = Keyword{"", ""}
	KEY_BLOCK_SIZE = Keyword{"", ""}
	KILL           = Keyword{"", ""}

	// L

	LAG            = Keyword{"", ""}
	LANGUAGE       = Keyword{"", ""}
	LAST           = Keyword{"", ""}
	LAST_VALUE     = Keyword{"", ""}
	LATERAL        = Keyword{"", ""}
	LEAD           = Keyword{"", ""}
	LEADING        = Keyword{"", ""}
	LEAVE          = Keyword{"", ""}
	LEAVES         = Keyword{"", ""}
	LEFT           = Keyword{"", ""}
	LESS           = Keyword{"", ""}
	LEVEL          = Keyword{"", ""}
	LIKE           = Keyword{"", ""}
	LIMIT          = Keyword{"", ""}
	LINEAR         = Keyword{"", ""}
	LINES          = Keyword{"", ""}
	LINESTRING     = Keyword{"", ""}
	LIST           = Keyword{"", ""}
	LOAD           = Keyword{"", ""}
	LOCAL          = Keyword{"", ""}
	LOCALTIME      = Keyword{"", ""}
	LOCALTIMESTAMP = Keyword{"", ""}
	LOCK           = Keyword{"", ""}
	LOCKED         = Keyword{"", ""}
	LOCKS          = Keyword{"", ""}
	LOGFILE        = Keyword{"", ""}
	LOGS           = Keyword{"", ""}
	LONG           = Keyword{"", ""}
	LONGBLOB       = Keyword{"", ""}
	LONGTEXT       = Keyword{"", ""}
	LOOP           = Keyword{"", ""}
	LOW_PRIORITY   = Keyword{"", ""}

	// M

	MASTER                        = Keyword{"", ""}
	MASTER_AUTO_POSITION          = Keyword{"", ""}
	MASTER_BIND                   = Keyword{"", ""}
	MASTER_COMPRESSION_ALGORITHMS = Keyword{"", ""}
	MASTER_CONNECT_RETRY          = Keyword{"", ""}
	MASTER_DELAY                  = Keyword{"", ""}
	MASTER_HEARTBEAT_PERIOD       = Keyword{"", ""}
	MASTER_HOST                   = Keyword{"", ""}
	MASTER_LOG_FILE               = Keyword{"", ""}
	MASTER_LOG_POS                = Keyword{"", ""}
	MASTER_PASSWORD               = Keyword{"", ""}
	MASTER_PORT                   = Keyword{"", ""}

	MASTER_PUBLIC_KEY_PATH = Keyword{"", ""}
	MASTER_RETRY_COUNT     = Keyword{"", ""}
	MASTER_SERVER_ID       = Keyword{"", ""}
	MASTER_SSL             = Keyword{"", ""}
	MASTER_SSL_CA          = Keyword{"", ""}
	MASTER_SSL_CAPATH      = Keyword{"", ""}
	MASTER_SSL_CERT        = Keyword{"", ""}
	MASTER_SSL_CIPHER      = Keyword{"", ""}
	MASTER_SSL_CRL         = Keyword{"", ""}

	MASTER_SSL_CRLPATH            = Keyword{"", ""}
	MASTER_SSL_KEY                = Keyword{"", ""}
	MASTER_SSL_VERIFY_SERVER_CERT = Keyword{"", ""}
	MASTER_TLS_CIPHERSUITES       = Keyword{"", ""}
	MASTER_TLS_VERSION            = Keyword{"", ""}
	MASTER_USER                   = Keyword{"", ""}
	MASTER_ZSTD_COMPRESSION_LEVEL = Keyword{"", ""}
	MATCH                         = Keyword{"", ""}
	MAXVALUE                      = Keyword{"", ""}

	MAX_CONNECTIONS_PER_HOUR = Keyword{"", ""}
	MAX_QUERIES_PER_HOUR     = Keyword{"", ""}
	MAX_ROWS                 = Keyword{"", ""}
	MAX_SIZE                 = Keyword{"", ""}
	MAX_UPDATES_PER_HOUR     = Keyword{"", ""}

	MAX_USER_CONNECTIONS = Keyword{"", ""}
	MEDIUM               = Keyword{"", ""}
	MEDIUMBLOB           = Keyword{"", ""}
	MEDIUMINT            = Keyword{"", ""}
	MEDIUMTEXT           = Keyword{"", ""}
	MEMBER               = Keyword{"", ""}
	MEMORY               = Keyword{"", ""}

	MERGE = Keyword{"", ""}

	MESSAGE_TEXT       = Keyword{"", ""}
	MICROSECOND        = Keyword{"", ""}
	MIDDLEINT          = Keyword{"", ""}
	MIGRATE            = Keyword{"", ""}
	MINUTE             = Keyword{"", ""}
	MINUTE_MICROSECOND = Keyword{"", ""}
	MINUTE_SECOND      = Keyword{"", ""}
	MIN_ROWS           = Keyword{"", ""}
	MOD                = Keyword{"", ""}
	MODE               = Keyword{"", ""}
	MODIFIES           = Keyword{"", ""}
	MODIFY             = Keyword{"", ""}
	MONTH              = Keyword{"", ""}
	MULTILINESTRING    = Keyword{"", ""}
	MULTIPOINT         = Keyword{"", ""}
	MULTIPOLYGON       = Keyword{"", ""}
	MUTEX              = Keyword{"", ""}
	MYSQL_ERRNO        = Keyword{"", ""}

	// N

	NAME       = Keyword{"", ""}
	NAMES      = Keyword{"", ""}
	NATIONAL   = Keyword{"", ""}
	NATURAL    = Keyword{"", ""}
	NCHAR      = Keyword{"", ""}
	NDB        = Keyword{"", ""}
	NDBCLUSTER = Keyword{"", ""}
	NESTED     = Keyword{"", ""}

	NETWORK_NAMESPACE  = Keyword{"", ""}
	NEVER              = Keyword{"", ""}
	NEW                = Keyword{"", ""}
	NEXTV              = Keyword{"", ""}
	NO                 = Keyword{"", ""}
	NODEGROUP          = Keyword{"", ""}
	NONE               = Keyword{"", ""}
	NOT                = Keyword{"", ""}
	NOWAIT             = Keyword{"", ""}
	NO_WAIT            = Keyword{"", ""}
	NO_WRITE_TO_BINLOG = Keyword{"", ""}
	NTH_VALUE          = Keyword{"", ""}
	NTILE              = Keyword{"", ""}
	NULL               = Keyword{"", ""}
	NULLS              = Keyword{"", ""}
	NUMBER             = Keyword{"", ""}
	NUMERIC            = Keyword{"", ""}
	NVARCHAR           = Keyword{"", ""}

	// O

	OF              = Keyword{"", ""}
	OFF             = Keyword{"", ""}
	OFFSET          = Keyword{"", ""}
	OJ              = Keyword{"", ""}
	OLD             = Keyword{"", ""}
	ON              = Keyword{"", ""}
	ONE             = Keyword{"", ""}
	ONLY            = Keyword{"", ""}
	OPEN            = Keyword{"", ""}
	OPTIMIZE        = Keyword{"", ""}
	OPTIMIZER_COSTS = Keyword{"", ""}
	OPTION          = Keyword{"", ""}
	OPTIONAL        = Keyword{"", ""}
	OPTIONALLY      = Keyword{"", ""}
	OPTIONS         = Keyword{"", ""}
	OR              = Keyword{"", ""}
	ORDER           = Keyword{"", ""}
	ORDINALITY      = Keyword{"", ""}
	ORGANIZATION    = Keyword{"", ""}
	OTHERS          = Keyword{"", ""}
	OUT             = Keyword{"", ""}
	OUTER           = Keyword{"", ""}
	OUTFILE         = Keyword{"", ""}
	OVER            = Keyword{"", ""}
	OWNER           = Keyword{"", ""}

	// P

	PACK_KEYS          = Keyword{"", ""}
	PAGE               = Keyword{"", ""}
	PARSER             = Keyword{"", ""}
	PARTIAL            = Keyword{"", ""}
	PARTITION          = Keyword{"", ""}
	PARTITIONING       = Keyword{"", ""}
	PARTITIONS         = Keyword{"", ""}
	PASSWORD           = Keyword{"", ""}
	PASSWORD_LOCK_TIME = Keyword{"", ""}
	PATH               = Keyword{"", ""}
	PERCENT_RANK       = Keyword{"", ""}
	PERSIST            = Keyword{"", ""}
	PERSIST_ONLY       = Keyword{"", ""}
	PHASE              = Keyword{"", ""}
	PLUGIN             = Keyword{"", ""}
	PLUGINS            = Keyword{"", ""}
	PLUGIN_DIR         = Keyword{"", ""}

	POINT = Keyword{"", ""}

	POLYGON = Keyword{"", ""}

	PORT = Keyword{"", ""}

	PRECEDES = Keyword{"", ""}

	PRECEDING = Keyword{"", ""}

	PRECISION = Keyword{"", ""}

	PREPARE = Keyword{"", ""}

	PRESERVE = Keyword{"", ""}

	PREV = Keyword{"", ""}

	PRIMARY = Keyword{"", ""}

	PRIVILEGES = Keyword{"", ""}

	PRIVILEGE_CHECKS_USER = Keyword{"", ""}
	PROCEDURE             = Keyword{"", ""}
	PROCESS               = Keyword{"", ""}
	PROCESSLIST           = Keyword{"", ""}
	PROFILE               = Keyword{"", ""}
	PROFILES              = Keyword{"", ""}
	PROXY                 = Keyword{"", ""}
	PURGE                 = Keyword{"", ""}

	// Q

	QUARTER = Keyword{"", ""}

	QUERY = Keyword{"", ""}

	QUICK = Keyword{"", ""}

	// R

	RANDOM = Keyword{"", ""}
	RANGE  = Keyword{"", ""}
	RANK   = Keyword{"", ""}

	READ       = Keyword{"", ""}
	READS      = Keyword{"", ""}
	READ_ONLY  = Keyword{"", ""}
	READ_WRITE = Keyword{"", ""}
	REAL       = Keyword{"", ""}

	REBUILD          = Keyword{"", ""}
	RECOVER          = Keyword{"", ""}
	RECURSIVE        = Keyword{"", ""}
	REDOFILE         = Keyword{"", ""}
	REDO_BUFFER_SIZE = Keyword{"", ""}

	REDUNDANT = Keyword{"", ""}
	REFERENCE = Keyword{"", ""}

	REFERENCES     = Keyword{"", ""}
	REGEXP         = Keyword{"", ""}
	RELAY          = Keyword{"", ""}
	RELAYLOG       = Keyword{"", ""}
	RELAY_LOG_FILE = Keyword{"", ""}
	RELAY_LOG_POS  = Keyword{"", ""}
	RELAY_THREAD   = Keyword{"", ""}
	RELEASE        = Keyword{"", ""}

	RELOAD     = Keyword{"", ""}
	REMOTE     = Keyword{"", ""}
	REMOVE     = Keyword{"", ""}
	RENAME     = Keyword{"", ""}
	REORGANIZE = Keyword{"", ""}
	REPAIR     = Keyword{"", ""}

	REPEAT = Keyword{"", ""}

	REPEATABLE = Keyword{"", ""}

	REPLACE = Keyword{"", ""}

	REPLICA = Keyword{"", ""}

	REPLICAS = Keyword{"", ""}

	REPLICATE_DO_DB    = Keyword{"", ""}
	REPLICATE_DO_TABLE = Keyword{"", ""}

	REPLICATE_IGNORE_DB = Keyword{"", ""}

	REPLICATE_IGNORE_TABLE = Keyword{"", ""}

	REPLICATE_REWRITE_DB = Keyword{"", ""}

	REPLICATE_WILD_DO_TABLE = Keyword{"", ""}

	REPLICATE_WILD_IGNORE_TABLE = Keyword{"", ""}

	REPLICATION = Keyword{"", ""}

	REQUIRE = Keyword{"", ""}

	REQUIRE_ROW_FORMAT = Keyword{"", ""}

	RESET = Keyword{"", ""}

	RESIGNAL = Keyword{"", ""}

	RESOURCE = Keyword{"", ""}

	RESPECT = Keyword{"", ""}

	RESTART = Keyword{"", ""}

	RESTORE = Keyword{"", ""}

	RESTRICT = Keyword{"", ""}

	RESUME = Keyword{"", ""}

	RETAIN = Keyword{"", ""}

	RETURN = Keyword{"", ""}

	RETURNED_SQLSTATE = Keyword{"", ""}

	RETURNING = Keyword{"", ""}

	RETURNS = Keyword{"", ""}

	REUSE = Keyword{"", ""}

	REVERSE = Keyword{"", ""}

	REVOKE = Keyword{"", ""}

	RIGHT = Keyword{"", ""}

	RLIKE = Keyword{"", ""}

	ROLE = Keyword{"", ""}

	ROLLBACK = Keyword{"", ""}

	ROLLUP = Keyword{"", ""}

	ROTATE = Keyword{"", ""}

	ROUTINE = Keyword{"", ""}

	ROW = Keyword{"", ""}

	ROWS = Keyword{"", ""}

	ROW_COUNT = Keyword{"", ""}

	ROW_FORMAT = Keyword{"", ""}

	ROW_NUMBER = Keyword{"", ""}

	RTREE = Keyword{"", ""}

	// S

	SAVEPOINT = Keyword{"", ""}

	SCHEDULE = Keyword{"", ""}

	SCHEMA = Keyword{"", ""}

	SCHEMAS = Keyword{"", ""}

	SCHEMA_NAME = Keyword{"", ""}

	SECOND = Keyword{"", ""}

	SECONDARY = Keyword{"", ""}

	SECONDARY_ENGINE = Keyword{"", ""}

	SECONDARY_ENGINE_ATTRIBUTE = Keyword{"", ""}

	SECONDARY_LOAD = Keyword{"", ""}

	SECONDARY_UNLOAD = Keyword{"", ""}

	SECOND_MICROSECOND = Keyword{"", ""}

	SECURITY = Keyword{"", ""}

	SELECT = NewKeywordCaseSensitive("SELECT")

	SENSITIVE = Keyword{"", ""}

	SEPARATOR = Keyword{"", ""}

	SERIAL = Keyword{"", ""}

	SERIALIZABLE = Keyword{"", ""}

	SERVER = Keyword{"", ""}

	SESSION = Keyword{"", ""}

	SET = Keyword{"", ""}

	SHARE = Keyword{"", ""}

	SHOW = Keyword{"", ""}

	SHUTDOWN = Keyword{"", ""}

	SIGNAL = Keyword{"", ""}

	SIGNED = Keyword{"", ""}

	SIMPLE = Keyword{"", ""}

	SKIP = Keyword{"", ""}

	SLAVE = Keyword{"", ""}

	SLOW = Keyword{"", ""}

	SMALLINT = Keyword{"", ""}

	SNAPSHOT = Keyword{"", ""}

	SOCKET = Keyword{"", ""}

	SOME = Keyword{"", ""}

	SONAME = Keyword{"", ""}

	SOUNDS = Keyword{"", ""}

	SOURCE = Keyword{"", ""}

	SPATIAL = Keyword{"", ""}

	SPECIFIC = Keyword{"", ""}

	SQL = Keyword{"", ""}

	SQLEXCEPTION = Keyword{"", ""}

	SQLSTATE = Keyword{"", ""}

	SQLWARNING = Keyword{"", ""}

	SQL_AFTER_GTIDS = Keyword{"", ""}

	SQL_AFTER_MTS_GAPS = Keyword{"", ""}

	SQL_BEFORE_GTIDS = Keyword{"", ""}

	SQL_BIG_RESULT = Keyword{"", ""}

	SQL_BUFFER_RESULT = Keyword{"", ""}

	SQL_CACHE = Keyword{"", ""}

	SQL_CALC_FOUND_ROWS = Keyword{"", ""}

	SQL_NO_CACHE = Keyword{"", ""}

	SQL_SMALL_RESULT = Keyword{"", ""}

	SQL_THREAD = Keyword{"", ""}

	SQL_TSI_DAY = Keyword{"", ""}

	SQL_TSI_HOUR = Keyword{"", ""}

	SQL_TSI_MINUTE = Keyword{"", ""}

	SQL_TSI_MONTH = Keyword{"", ""}

	SQL_TSI_QUARTER = Keyword{"", ""}

	SQL_TSI_SECOND = Keyword{"", ""}

	SQL_TSI_WEEK = Keyword{"", ""}

	SQL_TSI_YEAR = Keyword{"", ""}

	SRID = Keyword{"", ""}

	SSL = Keyword{"", ""}

	STACKED = Keyword{"", ""}

	START = Keyword{"", ""}

	STARTING = Keyword{"", ""}

	STARTS = Keyword{"", ""}

	STATS_AUTO_RECALC = Keyword{"", ""}

	STATS_PERSISTENT = Keyword{"", ""}

	STATS_SAMPLE_PAGES = Keyword{"", ""}

	STATUS = Keyword{"", ""}

	STOP = Keyword{"", ""}

	STORAGE = Keyword{"", ""}

	STORED = Keyword{"", ""}

	STRAIGHT_JOIN = Keyword{"", ""}

	STREAM = Keyword{"", ""}

	STRING = Keyword{"", ""}

	SUBCLASS_ORIGIN = Keyword{"", ""}

	SUBJECT = Keyword{"", ""}

	SUBPARTITION = Keyword{"", ""}

	SUBPARTITIONS = Keyword{"", ""}

	SUPER = Keyword{"", ""}

	SUSPEND = Keyword{"", ""}

	SWAPS = Keyword{"", ""}

	SWITCHES = Keyword{"", ""}

	SYSTEM = Keyword{"", ""}

	// T

	TABLE = Keyword{"", ""}

	TABLES = Keyword{"", ""}

	TABLESPACE = Keyword{"", ""}

	TABLE_CHECKSUM = Keyword{"", ""}

	TABLE_NAME = Keyword{"", ""}

	TEMPORARY = Keyword{"", ""}

	TEMPTABLE = Keyword{"", ""}

	TERMINATED = Keyword{"", ""}

	TEXT = Keyword{"", ""}

	THAN = Keyword{"", ""}

	THEN = Keyword{"", ""}

	THREAD_PRIORITY = Keyword{"", ""}

	TIES = Keyword{"", ""}

	TIME = Keyword{"", ""}

	TIMESTAMP = Keyword{"", ""}

	TIMESTAMPADD = Keyword{"", ""}

	TIMESTAMPDIFF = Keyword{"", ""}

	TINYBLOB = Keyword{"", ""}

	TINYINT = Keyword{"", ""}

	TINYTEXT = Keyword{"", ""}

	TLS = Keyword{"", ""}

	TO = Keyword{"", ""}

	TRAILING = Keyword{"", ""}

	TRANSACTION = Keyword{"", ""}

	TRIGGER = Keyword{"", ""}

	TRIGGERS = Keyword{"", ""}

	TRUE = Keyword{"", ""}

	TRUNCATE = Keyword{"", ""}

	TYPE = Keyword{"", ""}

	TYPES = Keyword{"", ""}

	// U

	UNBOUNDED = Keyword{"", ""}

	UNCOMMITTED = Keyword{"", ""}

	UNDEFINED = Keyword{"", ""}

	UNDO = Keyword{"", ""}

	UNDOFILE = Keyword{"", ""}

	UNDO_BUFFER_SIZE = Keyword{"", ""}

	UNICODE = Keyword{"", ""}

	UNINSTALL = Keyword{"", ""}

	UNION = Keyword{"", ""}

	UNIQUE = Keyword{"", ""}

	UNKNOWN = Keyword{"", ""}

	UNLOCK = Keyword{"", ""}

	UNSIGNED = Keyword{"", ""}

	UNTIL  = Keyword{"", ""}
	UPDATE = Keyword{"", ""}

	UPGRADE = Keyword{"", ""}

	USAGE = Keyword{"", ""}

	USE = Keyword{"", ""}

	USER = Keyword{"", ""}

	USER_RESOURCES = Keyword{"", ""}

	USE_FRM = Keyword{"", ""}

	USING         = Keyword{"", ""}
	UTC_DATE      = Keyword{"", ""}
	UTC_TIME      = Keyword{"", ""}
	UTC_TIMESTAMP = Keyword{"", ""}

	// V

	VALIDATION   = Keyword{"", ""}
	VALUE        = Keyword{"", ""}
	VALUES       = Keyword{"", ""}
	VARBINARY    = Keyword{"", ""}
	VARCHAR      = Keyword{"", ""}
	VARCHARACTER = Keyword{"", ""}
	VARIABLES    = Keyword{"", ""}
	VARYING      = Keyword{"", ""}
	VCPU         = Keyword{"", ""}
	VIEW         = Keyword{"", ""}
	VIRTUAL      = Keyword{"", ""}
	VISIBLE      = Keyword{"", ""}

	// W

	WAIT          = Keyword{"", ""}
	WARNINGS      = Keyword{"", ""}
	WEEK          = Keyword{"", ""}
	WEIGHT_STRING = Keyword{"", ""}
	WHEN          = Keyword{"", ""}
	WHERE         = NewKeywordCaseSensitive("WHERE")
	WHILE         = Keyword{"", ""}
	WINDOW        = Keyword{"", ""}
	WITH          = Keyword{"", ""}
	WITHOUT       = Keyword{"", ""}
	WORK          = Keyword{"", ""}
	WRAPPER       = Keyword{"", ""}
	WRITE         = Keyword{"", ""}

	// X
	X509 = Keyword{"", ""}
	XA   = Keyword{"", ""}
	XID  = Keyword{"", ""}
	XML  = Keyword{"", ""}
	XOR  = Keyword{"", ""}

	// Y

	YEAR       = Keyword{"", ""}
	YEAR_MONTH = Keyword{"", ""}

	// Z
	ZEROFILL = Keyword{"", ""}
	ZONE     = Keyword{"", ""}




	// Constructors symbols
	// .
	SYMB_DOT = NewKeyword(".")
	// (
	SYMB_LEFT_PAREN = NewKeyword("(")
	// )
	SYMB_RIGHT_PAREN = NewKeyword(")")
	// [
	SYMB_LERT_BRACKET = NewKeyword("[")
	// ]
	SYMB_RIGHT_BRACKET = NewKeyword("]")
	// {
	SYMB_LERT_BRACE = NewKeyword("{")
	// }
	SYMB_RIGHT_BRACE = NewKeyword("}")
	// @
	SYMB_AT = NewKeyword("@")
	// '
	SYMB_SINGLE_QUOTE = NewKeyword("'")
	// "
	SYMB_DOUBLE_QUOTE = "\""
	// `
	SYMB_REVERSE_QUOTE = "`"
	// :
	SYMB_COLON = ":"
	SYMB_INTRODUCER = "_"


	SYMB_COMMA = NewKeyword(",")
	SYMB_SEMI = NewKeyword(";")
)
