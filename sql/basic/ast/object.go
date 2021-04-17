package ast

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"reflect"
)

func IsNil(x ISQLObject) bool {
	if x == nil {
		return true
	}
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(x)
	return vi.IsNil()
}

type ISQLObject interface {
	ISQLReplaceable
	ISQLCloneable

	DBType() db.Type
	SetDBType(dbType db.Type)

	TargetDBType() db.Type
	SetTargetDBType(dbType db.Type)

	Line() int
	SetLine(line int)
	Col() int
	SetCol(col int)
	StartPos() int
	SetStartPos(startPos int)
	EndPos() int
	SetEndPos(endPos int)

	FormatLine() int
	SetFormatLine(formatLine int)
	FormatCol() int
	SetFormatCol(formatCol int)
	FormatStartPos() int
	SetFormatStartPos(formatStartPos int)
	FormatEndPos() int
	SetFormatEndPos(formatEndPos int)

	Parent() ISQLObject
	SetParent(parent ISQLObject)
	IsAfterSemi() bool
	SetAfterSemi(semi bool)

	String() string

	BeforeComments() []ISQLComment
	BeforeComment(i int) ISQLComment
	AddBeforeComment(comment ISQLComment)
	AddBeforeComments(comments []ISQLComment)

	AfterComments() []ISQLComment
	AfterComment(i int) ISQLComment
	AddAfterComment(comment ISQLComment)
	AddAfterComments(comments []ISQLComment)
}

type SQLObject struct {

	dbType       db.Type
	targetDBType db.Type

	line, col        int
	startPos, endPos int

	formatLine, formatCol        int
	formatStartPos, formatEndPos int

	parent ISQLObject
	semi   bool

	beforeComments []ISQLComment
	afterComments  []ISQLComment
}

func NewObject() *SQLObject {
	x := new(SQLObject)
	return x
}

func NewObjectWithDBType(dbType db.Type) *SQLObject {
	x := new(SQLObject)

	x.dbType = dbType
	return x
}

func (x *SQLObject) DBType() db.Type {
	if x.dbType != "" {
		return x.dbType
	}

	parent := x.parent
	for {
		if parent == nil {
			return x.dbType
		}

		if parent.DBType() != "" {
			x.dbType = parent.DBType()
			return x.dbType
		}
		parent = parent.Parent()
	}

}
func (x *SQLObject) SetDBType(dbType db.Type) {
	x.dbType = dbType
}

func (x *SQLObject) TargetDBType() db.Type {
	if x.targetDBType != "" {
		return x.targetDBType
	}

	parent := x.parent
	for {
		if parent == nil {
			return x.targetDBType
		}

		if parent.TargetDBType() != "" {
			x.targetDBType = parent.TargetDBType()
			return x.targetDBType
		}
		parent = parent.Parent()
	}
}
func (x *SQLObject) SetTargetDBType(targetDBType db.Type) {
	x.targetDBType = targetDBType
}

func (x *SQLObject) Line() int                { return x.line }
func (x *SQLObject) SetLine(line int)         { x.line = line }
func (x *SQLObject) Col() int                 { return x.col }
func (x *SQLObject) SetCol(col int)           { x.col = col }
func (x *SQLObject) StartPos() int            { return x.startPos }
func (x *SQLObject) SetStartPos(startPos int) { x.startPos = startPos }
func (x *SQLObject) EndPos() int              { return x.endPos }
func (x *SQLObject) SetEndPos(endPos int)     { x.endPos = endPos }

func (x *SQLObject) FormatLine() int                      { return x.formatLine }
func (x *SQLObject) SetFormatLine(formatLine int)         { x.formatLine = formatLine }
func (x *SQLObject) FormatCol() int                       { return x.formatCol }
func (x *SQLObject) SetFormatCol(formatCol int)           { x.formatCol = formatCol }
func (x *SQLObject) FormatStartPos() int                  { return x.formatStartPos }
func (x *SQLObject) SetFormatStartPos(formatStartPos int) { x.formatStartPos = formatStartPos }
func (x *SQLObject) FormatEndPos() int                    { return x.formatEndPos }
func (x *SQLObject) SetFormatEndPos(formatEndPos int)     { x.formatEndPos = formatEndPos }

func (x *SQLObject) Parent() ISQLObject {
	return x.parent
}

func (x *SQLObject) SetParent(parent ISQLObject) {
	x.parent = parent
}

func (x *SQLObject) IsAfterSemi() bool {
	return x.semi
}

func (x *SQLObject) SetAfterSemi(semi bool) {
	x.semi = semi
}

func (x *SQLObject) ReplaceChild(source ISQLObject, target ISQLObject) bool {
	sourceParentType := reflect.TypeOf(source.Parent())
	panic(sourceParentType)
}
func (x *SQLObject) Clone() ISQLObject {
	panic("implement me")
}

func (x *SQLObject) String() string {
	return ""
}


func (x *SQLObject) BeforeComments() []ISQLComment {
	return x.beforeComments
}
func (x *SQLObject) BeforeComment(i int) ISQLComment {
	return x.beforeComments[i]
}
func (x *SQLObject) AddBeforeComment(comment ISQLComment) {
	if comment == nil {
		return
	}
	x.beforeComments = append(x.beforeComments, comment)
}
func (x *SQLObject) AddBeforeComments(comments []ISQLComment) {
	if comments == nil || len(comments) == 0 {
		return
	}
	for _, comment := range comments {
		x.AddBeforeComment(comment)
	}
}
func (x *SQLObject) AfterComments() []ISQLComment {
	return x.afterComments
}
func (x *SQLObject) AfterComment(i int) ISQLComment {
	return x.afterComments[i]
}
func (x *SQLObject) AddAfterComment(comment ISQLComment) {
	x.afterComments = append(x.afterComments, comment)
}
func (x *SQLObject) AddAfterComments(comments []ISQLComment) {
	if comments == nil || len(comments) == 0 {
		return
	}
	for _, comment := range comments {
		x.AddAfterComment(comment)
	}
}






/**
 * --
 * /*
 */
type ISQLComment interface {
	ISQLObject
	Comment() string
	SetComment(comment string)
}

type abstractSQLComment struct {
	*SQLObject
	comment string
}

func NewAbstractSQLComment() *abstractSQLComment {
	x := new(abstractSQLComment)
	x.SQLObject = NewObject()
	return x
}
func (x *abstractSQLComment) Comment() string {
	return x.comment
}
func (x *abstractSQLComment) SetComment(comment string) {
	x.comment = comment
}

/**
 * --
 *
 * https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/plsql-language-fundamentals.html#GUID-068B0807-E244-4D0B-BA1B-47929CE626AF
 */
type SQLMinusComment struct {
	*abstractSQLComment
}

func NewMinusComment() *SQLMinusComment {
	x := new(SQLMinusComment)
	x.abstractSQLComment = NewAbstractSQLComment()
	return x
}
func NewMinusCommentWithComment(comment string) *SQLMinusComment {
	x := new(SQLMinusComment)
	x.abstractSQLComment = NewAbstractSQLComment()
	x.SetComment(comment)
	return x
}

// /* */
// https://docs.oracle.com/en/database/oracle/oracle-database/18/lnpls/plsql-language-fundamentals.html#GUID-2722A49F-C233-4DDF-B236-10EE3DD6B79B
type SQLMultiLineComment struct {
	*abstractSQLComment
}

func NewMultiLineComment() *SQLMultiLineComment {
	x := new(SQLMultiLineComment)
	x.abstractSQLComment = NewAbstractSQLComment()
	return x
}
func NewMultiLineCommentWithComment(comment string) *SQLMultiLineComment {
	x := new(SQLMultiLineComment)
	x.abstractSQLComment = NewAbstractSQLComment()
	x.SetComment(comment)
	return x
}

/**
 * #
 */
type SQLSharpComment struct {
	*abstractSQLComment
}
func NewSharpComment() *SQLSharpComment {
	x := new(SQLSharpComment)
	x.abstractSQLComment = NewAbstractSQLComment()
	return x
}
func NewSharpCommentWithComment(comment string) *SQLSharpComment {
	x := new(SQLSharpComment)
	x.abstractSQLComment = NewAbstractSQLComment()
	x.SetComment(comment)
	return x
}

// /* */
/**
 * /*+
 * https://dev.mysql.com/doc/refman/8.0/en/optimizer-hints.html
 */
type SQLMultiLineHint struct {
	*abstractSQLComment
}

func NewMultiLineHint() *SQLMultiLineHint {
	x := new(SQLMultiLineHint)
	x.abstractSQLComment = NewAbstractSQLComment()
	return x
}

type SQLMinusHint struct {
	*abstractSQLComment
}
