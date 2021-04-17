package visitor

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
	ADD             = NewKeywordCaseSensitive("ADD")
	ADMIN           = Keyword{"", ""}
	AFTER           = Keyword{"", ""}
	AGAINST         = Keyword{"", ""}
	AGGREGATE       = Keyword{"", ""}
	ALGORITHM       = Keyword{"", ""}
	ALL             = Keyword{"", ""}
	ALTER           = NewKeywordCaseSensitive("ALTER")
	ALWAYS          = Keyword{"", ""}
	ANALYSE         = Keyword{"", ""}
	ANALYZE         = NewKeywordCaseSensitive("ANALYZE")
	AND             = NewKeywordCaseSensitive("AND")
	ANY             = Keyword{"", ""}
	ARRAY           = Keyword{"", ""}
	AS              = NewKeywordCaseSensitive("AS")
	ASC             = Keyword{"", ""}
	ASCII           = Keyword{"", ""}
	ASENSITIVE      = Keyword{"", ""}
	AT              = Keyword{"", ""}
	ATTRIBUTE       = Keyword{"", ""}
	AUDIT           = NewKeywordCaseSensitive("AUDIT")
	AUTOEXTEND_SIZE = Keyword{"", ""}
	AUTO_INCREMENT  = NewKeywordCaseSensitive("AUTO_INCREMENT")
	AVG             = Keyword{"", ""}
	AVG_ROW_LENGTH  = Keyword{"", ""}

	// B

	BACKUP  = Keyword{"", ""}
	BEFORE  = Keyword{"", ""}
	BEGIN   = Keyword{"", ""}
	BETWEEN = NewKeywordCaseSensitive("BETWEEN")
	BIGINT  = Keyword{"", ""}
	BINARY  = Keyword{"", ""}
	BINLOG  = Keyword{"", ""}
	BIT     = Keyword{"", ""}
	BLOB    = Keyword{"", ""}
	BLOCK   = NewKeywordCaseSensitive("BLOCK")
	BODY    = NewKeywordCaseSensitive("BODY")
	BOOL    = Keyword{"", ""}
	BOOLEAN = Keyword{"", ""}
	BOTH    = Keyword{"", ""}
	BTREE   = Keyword{"", ""}
	BUCKETS = Keyword{"", ""}
	BY      = NewKeywordCaseSensitive("BY")
	BYTE    = Keyword{"", ""}

	// C

	CACHE        = NewKeywordCaseSensitive("CACHE")
	CALL         = Keyword{"", ""}
	CASCADE      = Keyword{"", ""}
	CASCADED     = Keyword{"", ""}
	CASE         = Keyword{"", ""}
	CATALOG_NAME = Keyword{"", ""}
	CHAIN        = Keyword{"", ""}
	CHANGE       = Keyword{"", ""}
	CHANGED      = Keyword{"", ""}
	CHANNEL      = Keyword{"", ""}
	CHAR         = Keyword{"", ""}
	CHARACTER    = NewKeywordCaseSensitive("CHARACTER")
	CHARSET      = NewKeywordCaseSensitive("CHARSET")
	CHECK        = NewKeywordCaseSensitive("CHECK")
	CHECKSUM     = Keyword{"", ""}
	CIPHER       = Keyword{"", ""}
	CLASS_ORIGIN = Keyword{"", ""}
	CLIENT       = Keyword{"", ""}
	CLONE        = Keyword{"", ""}
	CLOSE        = Keyword{"", ""}
	CLUSTER      = NewKeywordCaseSensitive("CLUSTER")

	COALESCE      = Keyword{"", ""}
	CODE          = Keyword{"", ""}
	COLLATE       = Keyword{"", ""}
	COLLATION     = Keyword{"", ""}
	COLUMN        = NewKeywordCaseSensitive("COLUMN")
	COLUMNS       = NewKeywordCaseSensitive("COLUMNS")
	COLUMN_FORMAT = Keyword{"", ""}
	COLUMN_NAME   = Keyword{"", ""}
	COMMENT       = NewKeywordCaseSensitive("COMMENT")
	COMMIT        = Keyword{"", ""}
	COMMITTED     = Keyword{"", ""}

	COMPACT    = Keyword{"", ""}
	COMPILE    = NewKeywordCaseSensitive("COMPILE")
	COMPLETION = Keyword{"", ""}

	COMPONENT          = Keyword{"", ""}
	COMPRESSED         = Keyword{"", ""}
	COMPRESSION        = Keyword{"", ""}
	CONCURRENT         = Keyword{"", ""}
	CONDITION          = Keyword{"", ""}
	CONNECT            = NewKeywordCaseSensitive("CONNECT")
	CONNECTION         = NewKeywordCaseSensitive("CONNECTION")
	CONSISTENT         = Keyword{"", ""}
	CONSTRAINT         = NewKeywordCaseSensitive("CONSTRAINT")
	CONSTRAINT_CATALOG = Keyword{"", ""}
	CONSTRAINT_NAME    = Keyword{"", ""}
	CONSTRAINT_SCHEMA  = Keyword{"", ""}
	CONTAINS           = Keyword{"", ""}
	CONTEXT            = Keyword{"", ""}
	CONTINUE           = Keyword{"", ""}
	CONVERT            = Keyword{"", ""}
	CPU                = Keyword{"", ""}
	CREATE             = NewKeywordCaseSensitive("CREATE")
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
	CYCLE              = NewKeywordCaseSensitive("CYCLE")

	// D

	DATA            = NewKeywordCaseSensitive("DATA")
	DATABASE        = NewKeywordCaseSensitive("DATABASE")
	DATABASES       = Keyword{"", ""}
	DATAFILE        = Keyword{"", ""}
	DATE            = NewKeywordCaseSensitive("DATE")
	DATETIME        = NewKeywordCaseSensitive("DATETIME")
	DAY             = NewKeywordCaseSensitive("DAY")
	DAY_HOUR        = Keyword{"", ""}
	DAY_MICROSECOND = Keyword{"", ""}
	DAY_MINUTE      = Keyword{"", ""}
	DAY_SECOND      = Keyword{"", ""}

	DEALLOCATE      = Keyword{"", ""}
	DEC             = Keyword{"", ""}
	DECIMAL         = Keyword{"", ""}
	DECLARE         = Keyword{"", ""}
	DEFAULT         = NewKeywordCaseSensitive("DEFAULT")
	DEFAULT_AUTH    = Keyword{"", ""}
	DEFINER         = Keyword{"", ""}
	DEFINITION      = Keyword{"", ""}
	DELAYED         = Keyword{"", ""}
	DELAY_KEY_WRITE = Keyword{"", ""}
	DELETE          = NewKeywordCaseSensitive("DELETE")
	DENSE_RANK      = Keyword{"", ""}
	DESC            = NewKeywordCaseSensitive("DESC")
	DESCRIBE        = NewKeywordCaseSensitive("DESCRIBE")
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

	DROP = NewKeywordCaseSensitive("DROP")

	DUAL      = Keyword{"", ""}
	DUMPFILE  = Keyword{"", ""}
	DUPLICATE = NewKeywordCaseSensitive("DUPLICATE")
	DYNAMIC   = Keyword{"", ""}

	// E
	EACH        = Keyword{"", ""}
	EDITION     = NewKeywordCaseSensitive("EDITION")
	EDITIONABLE = NewKeywordCaseSensitive("EDITIONABLE")

	ELSE             = Keyword{"", ""}
	ELSEIF           = Keyword{"", ""}
	EMPTY            = Keyword{"", ""}
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

	ESCAPE  = NewKeywordCaseSensitive("ESCAPE")
	ESCAPED = Keyword{"", ""}

	EVENT  = NewKeywordCaseSensitive("EVENT")
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
	EXPLAIN     = NewKeywordCaseSensitive("EXPLAIN")
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
	FOR       = NewKeywordCaseSensitive("FOR")
	FORCE     = NewKeywordCaseSensitive("FORCE")
	FOREIGN   = NewKeywordCaseSensitive("FOREIGN")
	FORMAT    = Keyword{"", ""}
	FOUND     = Keyword{"", ""}

	FROM = NewKeywordCaseSensitive("FROM")

	FULL     = NewKeywordCaseSensitive("FULL")
	FULLTEXT = Keyword{"", ""}
	FUNCTION = NewKeywordCaseSensitive("FUNCTION")

	// G

	GENERAL               = Keyword{"", ""}
	GENERATED             = Keyword{"", ""}
	GEOMCOLLECTION        = Keyword{"", ""}
	GEOMETRY              = Keyword{"", ""}
	GEOMETRYCOLLECTION    = Keyword{"", ""}
	GET                   = Keyword{"", ""}
	GET_FORMAT            = Keyword{"", ""}
	GET_MASTER_PUBLIC_KEY = Keyword{"", ""}
	GLOBAL                = NewKeywordCaseSensitive("GLOBAL")

	GRANT             = Keyword{"", ""}
	GRANTS            = Keyword{"", ""}
	GROUP             = NewKeywordCaseSensitive("GROUP")
	GROUPING          = Keyword{"", ""}
	GROUPS            = Keyword{"", ""}
	GROUP_REPLICATION = Keyword{"", ""}

	// H

	HANDLER = Keyword{"", ""}

	HASH   = NewKeywordCaseSensitive("HASH")
	HAVING = NewKeywordCaseSensitive("HAVING")

	HELP = NewKeywordCaseSensitive("HELP")

	HIGH_PRIORITY = Keyword{"", ""}
	HISTOGRAM     = Keyword{"", ""}
	HISTORY       = Keyword{"", ""}

	HOST             = NewKeywordCaseSensitive("HOST")
	HOSTS            = Keyword{"", ""}
	HOUR             = Keyword{"", ""}
	HOUR_MICROSECOND = Keyword{"", ""}
	HOUR_MINUTE      = Keyword{"", ""}
	HOUR_SECOND      = Keyword{"", ""}

	// I

	IDENTIFIED = NewKeywordCaseSensitive("IDENTIFIED")

	IF            = NewKeywordCaseSensitive("IF")
	IF_EXISTS     = NewKeywordCaseSensitive("IF EXISTS")
	IF_NOT_EXISTS = NewKeywordCaseSensitive("IF NOT EXISTS")

	IGNORE            = NewKeywordCaseSensitive("IGNORE")
	IGNORE_SERVER_IDS = Keyword{"", ""}

	IMPORT        = Keyword{"", ""}
	IN            = NewKeywordCaseSensitive("IN")
	INACTIVE      = Keyword{"", ""}
	INCREMENT     = NewKeywordCaseSensitive("INCREMENT")
	INDEX         = NewKeywordCaseSensitive("INDEX")
	INDEXESV      = Keyword{"", ""}
	INDEXTYPE     = NewKeywordCaseSensitive("INDEXTYPE")
	INFILE        = Keyword{"", ""}
	INITIAL_SIZE  = Keyword{"", ""}
	INNER         = Keyword{"", ""}
	INOUT         = Keyword{"", ""}
	INSENSITIVE   = Keyword{"", ""}
	INSERT        = NewKeywordCaseSensitive("INSERT")
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

	INTO      = NewKeywordCaseSensitive("INTO")
	INVISIBLE = NewKeywordCaseSensitive("INVISIBLE")
	INVOKER   = Keyword{"", ""}

	IO              = Keyword{"", ""}
	IO_AFTER_GTIDS  = Keyword{"", ""}
	IO_BEFORE_GTIDS = Keyword{"", ""}
	IO_THREAD       = Keyword{"", ""}
	IPC             = Keyword{"", ""}
	IS              = NewKeywordCaseSensitive("IS")
	ISOLATION       = Keyword{"", ""}
	ISSUER          = Keyword{"", ""}
	ITERATE         = Keyword{"", ""}

	// J

	JOIN       = Keyword{"", ""}
	JSON       = Keyword{"", ""}
	JSON_TABLE = Keyword{"", ""}
	JSON_VALUE = Keyword{"", ""}

	// K
	KEEP           = NewKeywordCaseSensitive("KEEP")
	KEY            = NewKeywordCaseSensitive("KEY")
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
	LESS           = NewKeywordCaseSensitive("LESS")
	LEVEL          = Keyword{"", ""}
	LIKE           = NewKeywordCaseSensitive("LIKE")
	LIMIT          = NewKeywordCaseSensitive("LIMIT")
	LINEAR         = NewKeywordCaseSensitive("LINEAR")
	LINES          = Keyword{"", ""}
	LINESTRING     = Keyword{"", ""}
	LIST           = NewKeywordCaseSensitive("LIST")
	LOAD           = Keyword{"", ""}
	LOCAL          = Keyword{"", ""}
	LOCALTIME      = Keyword{"", ""}
	LOCALTIMESTAMP = Keyword{"", ""}
	LOCK           = NewKeywordCaseSensitive("LOCK")
	LOCKED         = Keyword{"", ""}
	LOCKS          = Keyword{"", ""}
	LOGFILE        = Keyword{"", ""}
	LOGS           = Keyword{"", ""}
	LONG           = Keyword{"", ""}
	LONGBLOB       = Keyword{"", ""}
	LONGTEXT       = Keyword{"", ""}
	LOOP           = Keyword{"", ""}
	LOW_PRIORITY   = NewKeywordCaseSensitive("LOW_PRIORITY")

	// M
	MATERIALIZED                  = NewKeywordCaseSensitive("MATERIALIZED")
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
	MAXVALUE                      = NewKeywordCaseSensitive("MAXVALUE")

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

	MESSAGE_TEXT = Keyword{"", ""}
	MICROSECOND  = Keyword{"", ""}
	MIDDLEINT    = Keyword{"", ""}
	MIGRATE      = Keyword{"", ""}

	MINUTE             = Keyword{"", ""}
	MINUTE_MICROSECOND = Keyword{"", ""}
	MINUTE_SECOND      = Keyword{"", ""}
	MINVALUE           = NewKeywordCaseSensitive("MINVALUE")
	MIN_ROWS           = Keyword{"", ""}
	MOD                = Keyword{"", ""}
	MODE               = NewKeywordCaseSensitive("MODE")
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
	NAMES      = NewKeywordCaseSensitive("NAMES")
	NATIONAL   = Keyword{"", ""}
	NATURAL    = Keyword{"", ""}
	NCHAR      = Keyword{"", ""}
	NDB        = Keyword{"", ""}
	NDBCLUSTER = Keyword{"", ""}
	NESTED     = Keyword{"", ""}

	NETWORK_NAMESPACE = Keyword{"", ""}
	NEVER             = Keyword{"", ""}
	NEW               = Keyword{"", ""}
	NEXTV             = Keyword{"", ""}
	NO                = NewKeywordCaseSensitive("NO")
	NOCACHE           = NewKeywordCaseSensitive("NOCACHE")

	NOCYCLE            = NewKeywordCaseSensitive("NOCYCLE")
	NODEGROUP          = Keyword{"", ""}
	NOKEEP             = NewKeywordCaseSensitive("NOKEEP")
	NOMAXVALUE         = NewKeywordCaseSensitive("NOMAXVALUE")
	NOMINVALUE         = NewKeywordCaseSensitive("NOMINVALUE")
	NONE               = Keyword{"", ""}
	NONEDITIONABLE     = NewKeywordCaseSensitive("NONEDITIONABLE")
	NOORDER            = NewKeywordCaseSensitive("NOORDER")
	NOSCALE            = NewKeywordCaseSensitive("NOSCALE")
	NOT                = NewKeywordCaseSensitive("NOT")
	NOWAIT             = Keyword{"", ""}
	NO_WAIT            = Keyword{"", ""}
	NO_WRITE_TO_BINLOG = Keyword{"", ""}

	NTH_VALUE = Keyword{"", ""}
	NTILE     = Keyword{"", ""}
	NULL      = NewKeywordCaseSensitive("NULL")
	NULLS     = Keyword{"", ""}
	NUMBER    = Keyword{"", ""}
	NUMERIC   = Keyword{"", ""}
	NVARCHAR  = Keyword{"", ""}

	// O

	OF              = NewKeywordCaseSensitive("OF")
	OFF             = Keyword{"", ""}
	OFFSET          = NewKeywordCaseSensitive("OFFSET")
	OJ              = NewKeywordCaseSensitive("OJ")
	OLD             = Keyword{"", ""}
	ON              = NewKeywordCaseSensitive("ON")
	ONE             = Keyword{"", ""}
	ONLY            = NewKeywordCaseSensitive("ONLY")
	OPEN            = Keyword{"", ""}
	OPERATOR        = NewKeywordCaseSensitive("OPERATOR")
	OPTIMIZE        = Keyword{"", ""}
	OPTIMIZER_COSTS = Keyword{"", ""}
	OPTION          = Keyword{"", ""}
	OPTIONAL        = Keyword{"", ""}
	OPTIONALLY      = Keyword{"", ""}
	OPTIONS         = NewKeywordCaseSensitive("OPTIONS")
	OR              = NewKeywordCaseSensitive("OR")
	ORDER           = NewKeywordCaseSensitive("ORDER")
	ORDINALITY      = Keyword{"", ""}
	ORGANIZATION    = Keyword{"", ""}
	OTHERS          = Keyword{"", ""}
	OUT             = Keyword{"", ""}
	OUTER           = Keyword{"", ""}
	OUTFILE         = Keyword{"", ""}
	OVER            = Keyword{"", ""}
	OWNER           = NewKeywordCaseSensitive("OWNER")

	// P
	PACKAGE            = NewKeywordCaseSensitive("PACKAGE")
	PACK_KEYS          = Keyword{"", ""}
	PAGE               = Keyword{"", ""}
	PARSER             = Keyword{"", ""}
	PARTIAL            = Keyword{"", ""}
	PARTITION          = NewKeywordCaseSensitive("PARTITION")
	PARTITIONING       = Keyword{"", ""}
	PARTITIONS         = NewKeywordCaseSensitive("PARTITIONS")
	PASSWORD           = NewKeywordCaseSensitive("PASSWORD")
	PASSWORD_LOCK_TIME = Keyword{"", ""}
	PATH               = Keyword{"", ""}
	PERCENT_RANK       = Keyword{"", ""}
	PERSIST            = Keyword{"", ""}
	PERSIST_ONLY       = Keyword{"", ""}
	PHASE              = Keyword{"", ""}
	PLAIN              = NewKeywordCaseSensitive("PLAIN")
	PLUGIN             = Keyword{"", ""}
	PLUGINS            = Keyword{"", ""}
	PLUGIN_DIR         = Keyword{"", ""}

	POINT = Keyword{"", ""}

	POLICY  = NewKeywordCaseSensitive("POLICY")
	POLYGON = Keyword{"", ""}

	PORT = NewKeywordCaseSensitive("PORT")

	PRECEDES = Keyword{"", ""}

	PRECEDING = Keyword{"", ""}

	PRECISION = Keyword{"", ""}

	PREPARE = Keyword{"", ""}

	PRESERVE = Keyword{"", ""}

	PREV = Keyword{"", ""}

	PRIMARY = NewKeywordCaseSensitive("PRIMARY")

	PRIVILEGES = Keyword{"", ""}

	PRIVILEGE_CHECKS_USER = Keyword{"", ""}
	PROCEDURE             = NewKeywordCaseSensitive("PROCEDURE")
	PROCESS               = Keyword{"", ""}
	PROCESSLIST           = Keyword{"", ""}
	PROFILE               = Keyword{"", ""}
	PROFILES              = Keyword{"", ""}
	PROXY                 = Keyword{"", ""}
	PUBLIC                = NewKeywordCaseSensitive("PUBLIC")
	PURGE                 = Keyword{"", ""}

	// Q

	QUARTER = Keyword{"", ""}

	QUERY = Keyword{"", ""}

	QUICK = NewKeywordCaseSensitive("QUICK")

	// R

	RANDOM = NewKeywordCaseSensitive("RANDOM")
	RANGE  = NewKeywordCaseSensitive("RANGE")
	RANK   = Keyword{"", ""}

	READ       = Keyword{"", ""}
	READS      = Keyword{"", ""}
	READ_ONLY  = Keyword{"", ""}
	READ_WRITE = Keyword{"", ""}
	REAL       = Keyword{"", ""}

	REBUILD          = Keyword{"", ""}
	RECOVER          = Keyword{"", ""}
	RECURSIVE        = NewKeywordCaseSensitive("RECURSIVE")
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

	REPLACE = NewKeywordCaseSensitive("REPLACE")

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

	ROLE = NewKeywordCaseSensitive("ROLE")

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
	SAMPLE    = NewKeywordCaseSensitive("SAMPLE")
	SAVEPOINT = Keyword{"", ""}
	SCALE     = NewKeywordCaseSensitive("SCALE")
	SCHEDULE  = Keyword{"", ""}

	SCHEMA = NewKeywordCaseSensitive("SCHEMA")

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

	SEED   = NewKeywordCaseSensitive("SEED")
	SELECT = NewKeywordCaseSensitive("SELECT")

	SENSITIVE = Keyword{"", ""}

	SEPARATOR = Keyword{"", ""}
	SEQUENCE  = NewKeywordCaseSensitive("SEQUENCE")

	SERIAL = Keyword{"", ""}

	SERIALIZABLE = Keyword{"", ""}

	SERVER  = NewKeywordCaseSensitive("SERVER")
	SESSION = NewKeywordCaseSensitive("SESSION")

	SET = NewKeywordCaseSensitive("SET")

	SHARE = NewKeywordCaseSensitive("SHARE")

	SHOW = NewKeywordCaseSensitive("SHOW")

	SHUTDOWN = Keyword{"", ""}

	STATEMENT_ID = NewKeywordCaseSensitive("STATEMENT_ID")
	SIGNAL       = Keyword{"", ""}

	SIGNED = Keyword{"", ""}

	SIMPLE = Keyword{"", ""}

	SKIP = Keyword{"", ""}

	SLAVE = Keyword{"", ""}

	SLOW = Keyword{"", ""}

	SMALLINT = Keyword{"", ""}

	SNAPSHOT = Keyword{"", ""}

	SOCKET = NewKeywordCaseSensitive("SOCKET")

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

	START = NewKeywordCaseSensitive("START")

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

	SUBPARTITION  = NewKeywordCaseSensitive("SUBPARTITION")
	SUBPARTITIONS = NewKeywordCaseSensitive("SUBPARTITIONS")

	SUPER   = Keyword{"", ""}
	SUSPEND = Keyword{"", ""}

	SWAPS = Keyword{"", ""}

	SWITCHES = Keyword{"", ""}

	SYNONYM = NewKeywordCaseSensitive("SYNONYM")
	SYSTEM  = Keyword{"", ""}

	// T

	TABLE = NewKeywordCaseSensitive("TABLE")

	TABLES = Keyword{"", ""}

	TABLESPACE = Keyword{"", ""}

	TABLE_CHECKSUM = Keyword{"", ""}

	TABLE_NAME = Keyword{"", ""}

	TEMPORARY = NewKeywordCaseSensitive("TEMPORARY")

	TEMPTABLE = Keyword{"", ""}

	TERMINATED = Keyword{"", ""}

	TEXT = Keyword{"", ""}

	THAN = NewKeywordCaseSensitive("THAN")

	THEN = Keyword{"", ""}

	THREAD_PRIORITY = Keyword{"", ""}

	TIES = Keyword{"", ""}

	TIME      = NewKeywordCaseSensitive("TIME")
	TIMESTAMP = NewKeywordCaseSensitive("TIMESTAMP")

	TIMESTAMPADD = Keyword{"", ""}

	TIMESTAMPDIFF = Keyword{"", ""}

	TINYBLOB = Keyword{"", ""}

	TINYINT = Keyword{"", ""}

	TINYTEXT = Keyword{"", ""}

	TLS = Keyword{"", ""}

	TO = Keyword{"", ""}

	TRAILING = Keyword{"", ""}

	TRANSACTION = Keyword{"", ""}

	TRIGGER = NewKeywordCaseSensitive("TRIGGER")

	TRIGGERS = Keyword{"", ""}

	TRUE = Keyword{"", ""}

	TRUNCATE = Keyword{"", ""}

	TYPE = NewKeywordCaseSensitive("TYPE")

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

	UNIQUE = NewKeywordCaseSensitive("UNIQUE")

	UNKNOWN = Keyword{"", ""}

	UNLOCK = Keyword{"", ""}

	UNSIGNED = Keyword{"", ""}

	UNTIL  = Keyword{"", ""}
	UPDATE = NewKeywordCaseSensitive("UPDATE")

	UPGRADE = Keyword{"", ""}

	USAGE = Keyword{"", ""}

	USE = NewKeywordCaseSensitive("USE")

	USER = NewKeywordCaseSensitive("USER")

	USER_RESOURCES = Keyword{"", ""}

	USE_FRM = Keyword{"", ""}

	USING         = NewKeywordCaseSensitive("USING")
	UTC_DATE      = Keyword{"", ""}
	UTC_TIME      = Keyword{"", ""}
	UTC_TIMESTAMP = Keyword{"", ""}

	// V

	VALIDATION   = Keyword{"", ""}
	VALUE        = NewKeywordCaseSensitive("VALUE")
	VALUES       = NewKeywordCaseSensitive("VALUES")
	VARBINARY    = Keyword{"", ""}
	VARCHAR      = Keyword{"", ""}
	VARCHARACTER = Keyword{"", ""}
	VARIABLES    = Keyword{"", ""}
	VARYING      = Keyword{"", ""}
	VCPU         = Keyword{"", ""}
	VIEW         = NewKeywordCaseSensitive("VIEW")
	VIRTUAL      = Keyword{"", ""}
	VISIBLE      = NewKeywordCaseSensitive("VISIBLE")

	// W

	WAIT          = Keyword{"", ""}
	WARNINGS      = Keyword{"", ""}
	WEEK          = Keyword{"", ""}
	WEIGHT_STRING = Keyword{"", ""}
	WHEN          = Keyword{"", ""}
	WHERE         = NewKeywordCaseSensitive("WHERE")
	WHILE         = Keyword{"", ""}
	WINDOW        = Keyword{"", ""}
	WITH          = NewKeywordCaseSensitive("WITH")
	WITHOUT       = Keyword{"", ""}
	WORK          = Keyword{"", ""}
	WRAPPER       = NewKeywordCaseSensitive("WRAPPER")
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

	// Operators
	// Operators. Assigns
	// :=
	SYMB_COLON_EQUAL = NewKeyword(":=")
	// +=
	SYMB_PLUS_EQUAL = NewKeyword("+=")
	// -=
	SYMB_MINUS_EQUAL = NewKeyword("-=")
	// *=
	SYMB_MULT_EQUAL = NewKeyword("*=")
	// /=
	SYMB_DIV_EQUAL = NewKeyword("/=")

	// %=
	SYMB_MOD_EQUAL = NewKeyword("%=")

	// &=
	SYMB_AND_EQUAL = NewKeyword("INTERSECT")

	// ^=
	SYMB_XOR_EQUAL = NewKeyword("^=")

	// |=
	SYMB_OR_EQUAL = NewKeyword("|=")

	// Operators. Arithmetics
	// +
	SYMB_PLUS = NewKeyword("+")

	// -
	SYMB_MINUS = NewKeyword("-")
	// *
	SYMB_STAR = NewKeyword("*")

	// /
	SYMB_SLASH = NewKeyword("/")

	// %
	SYMB_PERCENT = NewKeyword("%")

	// --
	SYMB_MINUSMINUS = NewKeyword("--")

	// Operators. Comparation
	// =
	SYMB_EQUAL = NewKeyword("=")

	// =>
	SYMB_EQUAL_GREATER_THAN = NewKeyword("=>")

	// >
	SYMB_GREATER_THAN = NewKeyword(">")

	// >>
	SYMB_GREATER_THAN_GREATER_THAN = NewKeyword(">>")
	// >=
	SYMB_GREATER_THAN_EQUAL = NewKeyword(">=")

	// <
	SYMB_LESS_THAN = NewKeyword("<")
	// <<
	SYMB_LESS_THAN_LESS_THAN = NewKeyword("<<")
	// <=
	SYMB_LESS_THAN_EQUAL = NewKeyword("<=")

	// !
	SYMB_EXCLAMATION = NewKeyword("!")
	// !=
	SYMB_EXCLAMATION_EQUAL = NewKeyword("!=")
	// <>
	SYMB_LESS_THAN_GREATER_THAN = NewKeyword("<>")

	// <=>
	SYMB_LESS_THAN_EQUAL_GREATER_THAN = NewKeyword("<>")

	// ~=
	SYMB_NOT_EQUAL = NewKeyword("~=")

	// Operators. Bit
	// ~
	SYMB_BIT_NOT = NewKeyword("~")
	// |
	SYMB_BIT_OR = NewKeyword("|")
	// &
	SYMB_BIT_AND = NewKeyword("&")
	// ^
	SYMB_BIT_XOR = NewKeyword("^")

	// Constructors symbols
	// .
	SYMB_DOT = NewKeyword(".")
	// @
	SYMB_AT = NewKeyword("@")
	// #
	SYMB_SHARP = NewKeyword("#")
	// ?
	SYMB_QUESTION = NewKeyword("?")

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

	// ,
	SYMB_COMMA = NewKeyword(",")
	// ;
	SYMB_SEMI = NewKeyword(";")

	// :
	SYMB_COLON = NewKeyword(":")
	// _
	SYMB_INTRODUCER = NewKeyword("_")

	// ||
	SYMB_LOGICAL_OR = NewKeyword("||")
	// &&
	SYMB_LOGICAL_AND = NewKeyword("&&")

	SYMB_SINGLE_QUOTE = NewKeyword("'")
)
