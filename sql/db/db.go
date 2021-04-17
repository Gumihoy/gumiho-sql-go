package db

type Type string

const (
	SQL        Type = "SQL"
	MySQL      Type = "MySQL"
	MariaDB    Type = "MariaDB"
	Oracle     Type = "Oracle"
	PostgreSQL Type = "PostgreSQL"
	EDB        Type = "EDB"
	TiDB       Type = "TiDB"
	Hive       Type = "Hive"
)

type Version string

type SQLObjectType string

const (
	DATABASE SQLObjectType = "DATABASE"
	EXPLAIN  SQLObjectType = "EXPLAIN"
	FUNCTION SQLObjectType = "FUNCTION"
	INDEX    SQLObjectType = "INDEX"

	PROCEDURE SQLObjectType = "PROCEDURE"

	DELETE SQLObjectType = "DELETE"
	INSERT SQLObjectType = "INSERT"
	SELECT SQLObjectType = "SELECT"
	UPDATE SQLObjectType = "UPDATE"

	SEQUENCE SQLObjectType = "SEQUENCE"
	SYNONYM  SQLObjectType = "SYNONYM"

	TABLE     = "TABLE"
	TRIGGER   = "TRIGGER"
	TYPE      = "TYPE"
	TYPE_BODY = "TYPE BODY"
	VIEW      = "VIEW"
)
