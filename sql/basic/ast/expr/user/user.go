package user

import "github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr"

/**
 * ''@'' auth_option
 *
 * https://dev.mysql.com/doc/refman/8.0/en/create-user.html
 */
type SQLUserName struct {
	*expr.AbstractSQLExpr
	name   expr.ISQLExpr
	host   expr.ISQLExpr
	option expr.ISQLExpr
}

func NewUserName() *SQLUserName {
	x := new(SQLUserName)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func NewUserNameWithNameAndHost(name expr.ISQLExpr, host expr.ISQLExpr) *SQLUserName {
	x := new(SQLUserName)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	x.SetName(name)
	x.SetHost(host)
	return x
}
func (x *SQLUserName) StringName() string {
	return ""
}
func (x *SQLUserName) Name() expr.ISQLExpr {
	return x.name
}
func (x *SQLUserName) SetName(name expr.ISQLExpr) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLUserName) Host() expr.ISQLExpr {
	return x.host
}
func (x *SQLUserName) SetHost(host expr.ISQLExpr) {
	host.SetParent(x)
	x.host = host
}
func (x *SQLUserName) Option() expr.ISQLExpr {
	return x.option
}
func (x *SQLUserName) SetOption(option expr.ISQLExpr) {
	if option == nil {
		return
	}
	option.SetParent(x)
	x.option = option
}

/**
 * IDENTIFIED BY 'auth_string'
 */
type SQLIdentifiedByAuthOption struct {
	*expr.AbstractSQLExpr
	auth expr.ISQLExpr
}

func NewIdentifiedByAuthOption() *SQLIdentifiedByAuthOption {
	x := new(SQLIdentifiedByAuthOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIdentifiedByAuthOption) Auth() expr.ISQLExpr {
	return x.auth
}
func (x *SQLIdentifiedByAuthOption) SetAuth(auth expr.ISQLExpr) {
	if auth == nil {
		return
	}
	auth.SetParent(x)
	x.auth = auth
}

/**
 *  IDENTIFIED BY PASSWORD 'auth_string'
 */
type SQLIdentifiedByPasswordAuthOption struct {
	*expr.AbstractSQLExpr
	auth expr.ISQLExpr
}

func NewIdentifiedByPasswordAuthOption() *SQLIdentifiedByPasswordAuthOption {
	x := new(SQLIdentifiedByPasswordAuthOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIdentifiedByPasswordAuthOption) Auth() expr.ISQLExpr {
	return x.auth
}
func (x *SQLIdentifiedByPasswordAuthOption) SetAuth(auth expr.ISQLExpr) {
	if auth == nil {
		return
	}
	auth.SetParent(x)
	x.auth = auth
}

/**
 * IDENTIFIED BY RANDOM PASSWORD
 */
type SQLIdentifiedByRandomPasswordAuthOption struct {
	*expr.AbstractSQLExpr
	auth expr.ISQLExpr
}

func NewIdentifiedByRandomPasswordAuthOption() *SQLIdentifiedByRandomPasswordAuthOption {
	x := new(SQLIdentifiedByRandomPasswordAuthOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}

/**
 * IDENTIFIED WITH auth_plugin
 */
type SQLIdentifiedWithAuthOption struct {
	*expr.AbstractSQLExpr
	plugin expr.ISQLExpr
}

func NewIdentifiedWithAuthOption() *SQLIdentifiedWithAuthOption {
	x := new(SQLIdentifiedWithAuthOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIdentifiedWithAuthOption) Plugin() expr.ISQLExpr {
	return x.plugin
}
func (x *SQLIdentifiedWithAuthOption) SetPlugin(plugin expr.ISQLExpr) {
	if plugin == nil {
		return
	}
	plugin.SetParent(x)
	x.plugin = plugin
}

/**
 * IDENTIFIED WITH auth_plugin BY 'auth_string'
 */
type SQLIdentifiedWithByAuthOption struct {
	*expr.AbstractSQLExpr
	plugin expr.ISQLExpr
	auth   expr.ISQLExpr
}

func NewIdentifiedWithByAuthOption() *SQLIdentifiedWithByAuthOption {
	x := new(SQLIdentifiedWithByAuthOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIdentifiedWithByAuthOption) Plugin() expr.ISQLExpr {
	return x.plugin
}
func (x *SQLIdentifiedWithByAuthOption) SetPlugin(plugin expr.ISQLExpr) {
	if plugin == nil {
		return
	}
	plugin.SetParent(x)
	x.plugin = plugin
}
func (x *SQLIdentifiedWithByAuthOption) Auth() expr.ISQLExpr {
	return x.auth
}
func (x *SQLIdentifiedWithByAuthOption) SetAuth(auth expr.ISQLExpr) {
	if auth == nil {
		return
	}
	auth.SetParent(x)
	x.auth = auth
}

/**
 *  IDENTIFIED WITH auth_plugin BY RANDOM PASSWORD
 */
type SQLIdentifiedWithByRandomPasswordAuthOption struct {
	*expr.AbstractSQLExpr
	plugin expr.ISQLExpr
}

func NewIdentifiedWithByRandomPasswordAuthOption() *SQLIdentifiedWithByRandomPasswordAuthOption {
	x := new(SQLIdentifiedWithByRandomPasswordAuthOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIdentifiedWithByRandomPasswordAuthOption) Plugin() expr.ISQLExpr {
	return x.plugin
}
func (x *SQLIdentifiedWithByRandomPasswordAuthOption) SetPlugin(plugin expr.ISQLExpr) {
	if plugin == nil {
		return
	}
	plugin.SetParent(x)
	x.plugin = plugin
}

/**
 * IDENTIFIED WITH auth_plugin AS 'auth_string'
 */
type SQLIdentifiedWithAsAuthOption struct {
	*expr.AbstractSQLExpr
	plugin expr.ISQLExpr
	auth   expr.ISQLExpr
}

func NewIdentifiedWithAsAuthOption() *SQLIdentifiedWithAsAuthOption {
	x := new(SQLIdentifiedWithAsAuthOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *SQLIdentifiedWithAsAuthOption) Plugin() expr.ISQLExpr {
	return x.plugin
}
func (x *SQLIdentifiedWithAsAuthOption) SetPlugin(plugin expr.ISQLExpr) {
	if plugin == nil {
		return
	}
	plugin.SetParent(x)
	x.plugin = plugin
}
func (x *SQLIdentifiedWithAsAuthOption) Auth() expr.ISQLExpr {
	return x.auth
}
func (x *SQLIdentifiedWithAsAuthOption) SetAuth(auth expr.ISQLExpr) {
	if auth == nil {
		return
	}
	auth.SetParent(x)
	x.auth = auth
}

/**
 * resource_option: {
    MAX_QUERIES_PER_HOUR count
  | MAX_UPDATES_PER_HOUR count
  | MAX_CONNECTIONS_PER_HOUR count
  | MAX_USER_CONNECTIONS count
}
 * https://dev.mysql.com/doc/refman/8.0/en/create-user.html
 */
type abstractSQLResourceOption struct {
	*expr.AbstractSQLExpr
	value expr.ISQLExpr
}

func newAbstractSQLResourceOption() *abstractSQLResourceOption {
	x := new(abstractSQLResourceOption)
	x.AbstractSQLExpr = expr.NewAbstractExpr()
	return x
}
func (x *abstractSQLResourceOption) Value() expr.ISQLExpr {
	return x.value
}

func (x *abstractSQLResourceOption) SetValue(value expr.ISQLExpr) {
	if value == nil {
		return
	}
	value.SetParent(x)
	x.value = value
}

/**
 * MAX_QUERIES_PER_HOUR count
 * https://dev.mysql.com/doc/refman/8.0/en/create-user.html
 */
type SQLMaxQueriesPerHourResourceOption struct {
	*abstractSQLResourceOption
}

func NewMaxQueriesPerHourResourceOption() *SQLMaxQueriesPerHourResourceOption {
	x := new(SQLMaxQueriesPerHourResourceOption)
	x.abstractSQLResourceOption = newAbstractSQLResourceOption()
	return x
}

/**
 * password_option: {
    PASSWORD EXPIRE [DEFAULT | NEVER | INTERVAL N DAY]
  | PASSWORD HISTORY {DEFAULT | N}
  | PASSWORD REUSE INTERVAL {DEFAULT | N DAY}
  | PASSWORD REQUIRE CURRENT [DEFAULT | OPTIONAL]
  | FAILED_LOGIN_ATTEMPTS N
  | PASSWORD_LOCK_TIME {N | UNBOUNDED}
}
 * https://dev.mysql.com/doc/refman/8.0/en/create-user.html
 */

/**
 * lock_option: {
   ACCOUNT LOCK
 | ACCOUNT UNLOCK
}
 * https://dev.mysql.com/doc/refman/8.0/en/create-user.html
 */
