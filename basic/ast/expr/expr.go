package expr

import (
	"gumihoy.com/sql/basic/ast"
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

type SQLExpr struct {
	*ast.SQLObject
}

func NewExpr() *SQLExpr {
	x := new(SQLExpr)
	x.SQLObject = ast.NewObject()
	return x
}

type ISQLName interface {
	ISQLExpr
	StringName() string
	SimpleStringName() string
}

type SQLName struct {
	*SQLExpr
	owner ISQLName
	name  ISQLIdentifier

	fullName string
}

func NewName(owner string, name string) *SQLName {
	return NewNameWithOwner(NewUnQuotedIdentifier(owner), name)
}

func NewNameWithOwner(owner ISQLName, name string) *SQLName {
	return NewNameWithOwnerAndName(owner, NewUnQuotedIdentifier(name))
}

func NewNameWithOwnerAndName(owner ISQLName, name ISQLIdentifier) *SQLName {
	x := new(SQLName)
	x.SQLExpr = NewExpr()
	x.owner = owner
	x.name = name
	return x
}

func (x *SQLName) StringName() string {
	if x.fullName == "" {
		x.fullName = x.owner.StringName() + "." + x.name.StringName()
	}
	return x.fullName
}

func (x *SQLName) SimpleStringName() string {
	return x.name.StringName()
}

type ISQLIdentifier interface {
	ISQLName
}

type AbstractSQLIdentifier struct {
	*SQLExpr
	name string
}

const (
	DOUBLE_QUOTE  = "\""
	REVERSE_QUOTE = "`"
)

func OfIdentifier(name string) ISQLIdentifier {
	if name == "" {
		panic("name is nil.")
	}

	name = strings.TrimSpace(name)
	if name[0] == '"' && name[len(name)-1] == '"' {
		return nil
	} else if name[0] == '"' && name[len(name)-1] == '"' {
		return nil
	} else {
		return nil
	}
}

type SQLUnQuotedIdentifier struct {
	*SQLExpr
	name string
}

func NewUnQuotedIdentifier(name string) *SQLUnQuotedIdentifier {
	var x SQLUnQuotedIdentifier
	x.SQLExpr = NewExpr()
	x.SetName(name)
	return &x
}

func (x *SQLUnQuotedIdentifier) StringName() string {
	return x.name
}
func (x *SQLUnQuotedIdentifier) SimpleStringName() string {
	return x.name
}

func (x *SQLUnQuotedIdentifier) SetName(name string) {
	if strings.TrimSpace(name) == "" {
		panic("name is nil.")
	}

	x.name = name
}

type SQLDoubleQuotedIdentifier struct {
	*SQLExpr
	name string
}

func NewDoubleQuotedIdentifier(name string) *SQLDoubleQuotedIdentifier {
	var x SQLDoubleQuotedIdentifier
	x.SQLExpr = NewExpr()
	x.SetName(name)
	return &x
}

func (x *SQLDoubleQuotedIdentifier) StringName() string {
	return x.name
}
func (x *SQLDoubleQuotedIdentifier) SimpleStringName() string {
	return x.name
}

func (x *SQLDoubleQuotedIdentifier) SetName(name string) {
	length := len(name)
	if length > 1 &&
		((name[0] == '"' && name[len(name)-1] != '"') ||
			(name[0] != '"' && name[len(name)-1] == '"')) {
		panic("name:" + name + " is error.")
	}
	if length > 1 && name[0] == '"' && name[len(name)-1] == '"' {
		name = name[1 : len(name)-1]
	}
	x.name = name
}

type SQLReverseQuotedIdentifier struct {
	*SQLExpr
	name string
}

func NewReverseQuotedIdentifier(name string) *SQLReverseQuotedIdentifier {
	var x SQLReverseQuotedIdentifier
	x.SQLExpr = NewExpr()
	x.SetName(name)
	return &x
}

func (x *SQLReverseQuotedIdentifier) StringName() string {
	return x.name
}
func (x *SQLReverseQuotedIdentifier) SimpleStringName() string {
	return x.name
}

func (x *SQLReverseQuotedIdentifier) SetName(name string) {
	length := len(name)
	if length > 1 &&
		((name[0] == '`' && name[len(name)-1] != '`') ||
			(name[0] != '`' && name[len(name)-1] == '`')) {
		panic("name:" + name + " is error.")
	}
	if length > 1 && name[0] == '`' && name[len(name)-1] == '`' {
		name = name[1 : len(name)-1]
	}
	x.name = name
}

type SQLDBLinkExpr struct {
	*SQLExpr

	owner ISQLName
	name  ISQLName

	fullName string
}

func NewDBLinkExpr(owner ISQLName, name ISQLName) *SQLDBLinkExpr {
	x := new(SQLDBLinkExpr)
	x.SQLExpr = NewExpr()
	x.owner = owner
	x.name = name
	return x
}
func (x *SQLDBLinkExpr) StringName() string {
	if x.fullName == "" {
		x.fullName = x.owner.StringName() + "@" + x.name.StringName()
	}
	return x.fullName
}
func (x *SQLDBLinkExpr) SimpleStringName() string {
	return x.StringName()
}
func (x *SQLDBLinkExpr) Owner() ISQLName {
	return x.owner
}
func (x *SQLDBLinkExpr) SetOwner(owner ISQLName) {
	owner.SetParent(x)
	x.owner = owner
}
func (x *SQLDBLinkExpr) Name() ISQLName {
	return x.name
}

func (x *SQLDBLinkExpr) SetName(name ISQLName) {
	name.SetParent(x)
	x.name = name
}

type SQLAllColumnExpr struct {
	*SQLExpr
}



// *
func NewAllColumnExpr() *SQLAllColumnExpr {
	x := new(SQLAllColumnExpr)
	x.SQLExpr = NewExpr()
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
	*SQLExpr
}

func NewNullExpr() *SQLNullExpr  {
	var x SQLNullExpr
	x.SQLExpr = NewExpr()
	return &x
}



// ( expr, expr... )
type SQLListExpr struct {
	*SQLExpr
	elements []ISQLExpr
}

func NewListExpr(elements ...ISQLExpr) *SQLListExpr {
	var x SQLListExpr
	x.SQLExpr = NewExpr()
	for _, element := range elements {
		x.AddElement(element)
	}
	return &x
}

func (x *SQLListExpr) Elements() []ISQLExpr {
	return x.elements
}

func (x *SQLListExpr) AddElement(element ISQLExpr) {
	element.SetParent(x)
	x.elements = append(x.elements, element)
}

type R struct {
}
