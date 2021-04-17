package expr

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
)

type SQLSetQuantifier string

const (
	DISTINCT SQLSetQuantifier = "DISTINCT"
	ALL                       = "ALL"
)

type ISQLExpr interface {
	ast.ISQLObject
}

type AbstractSQLExpr struct {
	*ast.SQLObject
}

func NewAbstractExpr() *AbstractSQLExpr {
	x := new(AbstractSQLExpr)
	x.SQLObject = ast.NewObject()
	return x
}

func NewAbstractExprWithDBType(dbType db.Type) *AbstractSQLExpr {
	x := new(AbstractSQLExpr)
	x.SQLObject = ast.NewObjectWithDBType(dbType)
	return x
}

type ISQLName interface {
	ISQLExpr
	StringName() string
}

type ISQLIdentifier interface {
	ISQLName

	Equal(t string) bool
	EqualIgnoreCase(t string) bool
}

type AbstractSQLIdentifier struct {
	*AbstractSQLExpr
	name string
}

func NewAbstractSQLIdentifier() *AbstractSQLIdentifier {
	x := new(AbstractSQLIdentifier)
	x.AbstractSQLExpr = NewAbstractExpr()
	return x
}

func (x *AbstractSQLIdentifier) StringName() string {
	return x.name
}

func (x *AbstractSQLIdentifier) Name() string {
	return x.name
}

func (x *AbstractSQLIdentifier) SetName(name string) {
	if strings.TrimSpace(name) == "" {
		panic("name is nil.")
	}

	x.name = removeQuoted(name)
}

func (x *AbstractSQLIdentifier) Equal(t string) bool {
	return x.name == t
}

func (x *AbstractSQLIdentifier) EqualIgnoreCase(t string) bool {
	return strings.EqualFold(x.name, t)
}

func (x *AbstractSQLIdentifier) String() string {
	return x.name
}

func removeQuoted(name string) string {
	if len(name) < 2 {
		return name
	}

	len := len(name)

	if (name[0] == '"' && name[len-1] == '"') ||
		(name[0] == '`' && name[len-1] == '`') {
		name = name[1 : len-1]
	}

	return name
}

/**
 * xx.xxx
 */
type SQLName struct {
	*AbstractSQLExpr
	owner ISQLExpr
	name  ISQLIdentifier
}

func NewName(owner string, name string) *SQLName {
	return NewNameWithOwner(NewUnQuotedIdentifier(owner), name)
}

func NewNameWithOwner(owner ISQLExpr, name string) *SQLName {
	return NewNameWithOwnerAndName(owner, NewUnQuotedIdentifier(name))
}

func NewNameWithOwnerAndName(owner ISQLExpr, name ISQLIdentifier) *SQLName {
	x := new(SQLName)
	x.AbstractSQLExpr = NewAbstractExpr()
	x.owner = owner
	x.name = name
	return x
}

func (x *SQLName) StringName() string {
	return x.name.StringName()
}

func (x *SQLName) Owner() ISQLExpr {
	return x.owner
}

func (x *SQLName) SetOwner(owner ISQLExpr) {
	owner.SetParent(x)
	x.owner = owner
}

func (x *SQLName) Name() ISQLIdentifier {
	return x.name
}

func (x *SQLName) SetName(name ISQLIdentifier) {
	name.SetParent(x)
	x.owner = name
}

func (x *SQLName) String() string {
	return x.owner.String() + "." + x.name.String()
}

func OfIdentifier(name string) ISQLIdentifier {
	if name == "" {
		panic("name is nil.")
	}

	name = strings.TrimSpace(name)
	if name[0] == '"' && name[len(name)-1] == '"' {
		return NewDoubleQuotedIdentifier(name)
	} else if name[0] == '`' && name[len(name)-1] == '`' {
		return NewReverseQuotedIdentifier(name)
	} else {
		return NewUnQuotedIdentifier(name)
	}
}

/**
 * name
 */
type SQLUnQuotedIdentifier struct {
	*AbstractSQLIdentifier
}

func NewUnQuotedIdentifier(name string) *SQLUnQuotedIdentifier {
	x := new(SQLUnQuotedIdentifier)
	x.AbstractSQLIdentifier = NewAbstractSQLIdentifier()
	x.SetName(name)
	return x
}

func (x *SQLUnQuotedIdentifier) String() string {
	return x.name
}

/**
 * "name"
 */
type SQLDoubleQuotedIdentifier struct {
	*AbstractSQLIdentifier
}

func NewDoubleQuotedIdentifier(name string) *SQLDoubleQuotedIdentifier {
	x := new(SQLDoubleQuotedIdentifier)
	x.AbstractSQLIdentifier = NewAbstractSQLIdentifier()
	x.SetName(name)
	return x
}

func (x *SQLDoubleQuotedIdentifier) String() string {
	return "\"" + x.name + "\""
}

/**
 * `name`
 */
type SQLReverseQuotedIdentifier struct {
	*AbstractSQLIdentifier
}

func NewReverseQuotedIdentifier(name string) *SQLReverseQuotedIdentifier {
	x := new(SQLReverseQuotedIdentifier)
	x.AbstractSQLIdentifier = NewAbstractSQLIdentifier()
	x.SetName(name)
	return x
}

func (x *SQLReverseQuotedIdentifier) String() string {
	return "`" + x.name + "`"
}

/**
 * xx.xx@xx.xx
 */
type SQLDBLinkExpr struct {
	*AbstractSQLExpr

	name   ISQLExpr
	dbLink ISQLName

	fullName string
}

func NewDBLinkExpr(name ISQLExpr, dbLink ISQLName) *SQLDBLinkExpr {
	x := new(SQLDBLinkExpr)
	x.AbstractSQLExpr = NewAbstractExpr()
	x.SetName(name)
	x.SetDBLink(dbLink)
	return x
}
func (x *SQLDBLinkExpr) StringName() string {
	switch x.name.(type) {
	case ISQLName:
		return x.name.(ISQLName).StringName() + "@" + x.dbLink.StringName()
	}
	return ""
}

func (x *SQLDBLinkExpr) SimpleStringName() string {
	return x.StringName()
}
func (x *SQLDBLinkExpr) Name() ISQLExpr {
	return x.name
}

func (x *SQLDBLinkExpr) SetName(name ISQLExpr) {
	name.SetParent(x)
	x.name = name
}

func (x *SQLDBLinkExpr) DBLink() ISQLName {
	return x.dbLink
}
func (x *SQLDBLinkExpr) SetDBLink(dbLink ISQLName) {
	dbLink.SetParent(x)
	x.dbLink = dbLink
}

type SQLAllColumnExpr struct {
	*AbstractSQLIdentifier
}

// *
func NewAllColumnExpr() *SQLAllColumnExpr {
	x := new(SQLAllColumnExpr)
	x.AbstractSQLIdentifier = NewAbstractSQLIdentifier()
	return x
}

// func (x *SQLAllColumnExpr) Name() string {
// 	return "*"
// }
//
// func (x *SQLAllColumnExpr) SimpleName() string {
// 	return "*"
// }

// null
type SQLNullExpr struct {
	*AbstractSQLExpr
}

func NewNullExpr() *SQLNullExpr {
	var x SQLNullExpr
	x.AbstractSQLExpr = NewAbstractExpr()
	return &x
}

// ( expr, expr... )
type SQLListExpr struct {
	*AbstractSQLExpr
	elements []ISQLExpr
}

func NewListExpr(elements ...ISQLExpr) *SQLListExpr {
	var x SQLListExpr
	x.AbstractSQLExpr = NewAbstractExpr()
	for _, element := range elements {
		x.AddElement(element)
	}
	return &x
}

func (x *SQLListExpr) ReplaceChild(source ast.ISQLObject, target ast.ISQLObject) bool {
	for i, element := range x.elements {
		if element == source {
			x.elements[i] = target
			return true
		}
	}
	return false
}

func (x *SQLListExpr) Elements() []ISQLExpr {
	return x.elements
}
func (x *SQLListExpr) Element(i int) ISQLExpr {
	return x.elements[i]
}
func (x *SQLListExpr) AddElement(element ISQLExpr) {
	element.SetParent(x)
	x.elements = append(x.elements, element)
}

/**
 *
 * https://ronsavage.github.io/SQL/sql-2003-2.bnf.html#case%20expression
 */
type SQLCaseExpr struct {
	*AbstractSQLExpr
}

// id(+)
//
// https://docs.oracle.com/en/database/oracle/oracle-database/18/sqlrf/Joins.html#GUID-29A4584C-0741-4E6A-A89B-DCFAA222994A
type SQLOuterJoinExpr struct {
	*AbstractSQLExpr
	name ISQLExpr
}

func NewOuterJoinExpr(name ISQLExpr) *SQLOuterJoinExpr {
	x := new(SQLOuterJoinExpr)
	x.AbstractSQLExpr = NewAbstractExpr()
	x.SetName(name)
	return x
}

func (x *SQLOuterJoinExpr) Name() ISQLExpr {
	return x.name
}

func (x *SQLOuterJoinExpr) SetName(name ISQLExpr) {
	name.SetParent(x)
	x.name = name
}

/**
 * name => value
 *
 */
type SQLCallExpr struct {
	*AbstractSQLExpr
	name  ISQLExpr
	value ISQLExpr
}

func NewCallExpr() *SQLCallExpr {
	x := new(SQLCallExpr)
	return x
}

func NewCallExprWithNameAndValue(name ISQLExpr, value ISQLExpr) *SQLCallExpr {
	x := new(SQLCallExpr)
	x.SetName(name)
	x.SetValue(value)
	return x
}

func (x *SQLCallExpr) Name() ISQLExpr {
	return x.name
}

func (x *SQLCallExpr) SetName(name ISQLExpr) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLCallExpr) Value() ISQLExpr {
	return x.value
}

func (x *SQLCallExpr) SetValue(value ISQLExpr) {
	value.SetParent(x)
	x.value = value
}

/**
 *
 * https://www.postgresql.org/docs/devel/sql-expressions.html#SQL-SYNTAX-ARRAY-CONSTRUCTORS
 */
type SQLArrayExpr struct {
	*AbstractSQLExpr
	name      ISQLExpr
	arguments []ISQLExpr
}

func NewLArrayExpr() *SQLArrayExpr {
	x := new(SQLArrayExpr)
	x.AbstractSQLExpr = NewAbstractExpr()
	return x
}
func (x *SQLArrayExpr) Name() ISQLExpr {
	return x.name
}

func (x *SQLArrayExpr) SetName(name ISQLExpr) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLArrayExpr) Arguments() []ISQLExpr {
	return x.arguments
}

func (x *SQLArrayExpr) SetArguments(argument ISQLExpr) {
	if argument == nil {
		return
	}
	argument.SetParent(x)
	x.arguments = append(x.arguments, argument)
}

/**
 *  AUTO_INCREMENT [=] value
  | AVG_ROW_LENGTH [=] value
  | [DEFAULT] CHARACTER SET [=] charset_name
  | CHECKSUM [=] {0 | 1}
  | [DEFAULT] COLLATE [=] collation_name
  | COMMENT [=] 'string'
  | COMPRESSION [=] {'ZLIB' | 'LZ4' | 'NONE'}
  | CONNECTION [=] 'connect_string'
  | {DATA | INDEX} DIRECTORY [=] 'absolute path to directory'
  | DELAY_KEY_WRITE [=] {0 | 1}
  | ENCRYPTION [=] {'Y' | 'N'}
  | ENGINE [=] engine_name
  | INSERT_METHOD [=] { NO | FIRST | LAST }
  | KEY_BLOCK_SIZE [=] value
  | MAX_ROWS [=] value
  | MIN_ROWS [=] value
  | PACK_KEYS [=] {0 | 1 | DEFAULT}
  | PASSWORD [=] 'string'
  | ROW_FORMAT [=] {DEFAULT | DYNAMIC | FIXED | COMPRESSED | REDUNDANT | COMPACT}
  | STATS_AUTO_RECALC [=] {DEFAULT | 0 | 1}
  | STATS_PERSISTENT [=] {DEFAULT | 0 | 1}
  | STATS_SAMPLE_PAGES [=] value
  | TABLESPACE tablespace_name [STORAGE {DISK | MEMORY}]
  | UNION [=] (tbl_name[,tbl_name]...)
 *
 * https://dev.mysql.com/doc/refman/5.7/en/create-table.html
 */
type SQLAssignExpr struct {
	*AbstractSQLExpr
	name  ISQLExpr
	Equal bool
	value ISQLExpr
}

func NewAssignExpr() *SQLAssignExpr {
	x := new(SQLAssignExpr)
	x.AbstractSQLExpr = NewAbstractExpr()
	return x
}
func (x *SQLAssignExpr) Name() ISQLExpr {
	return x.name
}

func (x *SQLAssignExpr) SetName(name ISQLExpr) {
	name.SetParent(x)
	x.name = name
}
func (x *SQLAssignExpr) Value() ISQLExpr {
	return x.value
}
func (x *SQLAssignExpr) SetValue(value ISQLExpr) {
	value.SetParent(x)
	x.value = value
}
