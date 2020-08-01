package dbtype

type DBType string

const (
	SQL        DBType = "SQL"
	MySQL      DBType = "MySQL"
	MariaDB           = "MariaDB"
	Oracle            = "Oracle"
	PostgreSQL        = "PostgreSQL"
)
