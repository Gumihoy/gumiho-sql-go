package db

type DBType string

const (
	SQL        DBType = "SQL"
	MySQL      DBType = "MySQL"
	MariaDB           = "MariaDB"
	Oracle            = "Oracle"
	PostgreSQL        = "PostgreSQL"
	EDB               = "EDB"
)

type Version string

type ObjectType string

const (
	TABLE ObjectType = "TABLE"
)
