package select_

type ISQLSelectQuery interface {
}

type AbstractSQLSelectQuery struct {
	orderByClause SQLOrderByClause
	limitClause   ISQLLimitClause
	lockClause    ISQLLockClause
}

type SQLOrderByClause struct {
}
type ISQLLimitClause interface {
}
type ISQLLockClause interface {
}
