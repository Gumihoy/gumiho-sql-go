package ast

import (
	"gumihoy.com/sql/db"
	"reflect"
)

type ISQLObject interface {
	DBType() db.DBType
	SetDBType(dbType db.DBType)

	TargetDBType() db.DBType
	SetTargetDBType(dbType db.DBType)

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
}

type SQLObject struct {
	dbType       db.DBType
	targetDBType db.DBType

	line, col        int
	startPos, endPos int

	formatLine, formatCol        int
	formatStartPos, formatEndPos int

	parent ISQLObject
	semi   bool
}

func NewObject() *SQLObject {
	x := new(SQLObject)
	return x
}

func (x *SQLObject) DBType() db.DBType {
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
func (x *SQLObject) SetDBType(dbType db.DBType) {
	x.dbType = dbType
}

func (x *SQLObject) TargetDBType() db.DBType {
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
func (x *SQLObject) SetTargetDBType(targetDBType db.DBType) {
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

// func (x *SQLObject) String() string {
// 	return ""
// }

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

type ISQLReplaceable interface {
	Replace(source ISQLObject, target ISQLObject) bool
}

type SQLReplaceable struct {
}

func (x *SQLReplaceable) Replace(source ISQLObject, target ISQLObject) bool {
	panic("implement me")
}

func (x *SQLReplaceable) ReplaceInList(exprs []interface{}, source ISQLObject, target ISQLObject, parent ISQLObject) bool {
	if exprs == nil {
		return false
	}

	if target == nil {
		for i := len(exprs) - 1; i >= 0; i-- {
			if source == exprs[i] {
				// delete(exprs, i)
				return true
			}
		}
		return false
	}

	for i := 0; i < len(exprs); i++ {
		if exprs[i] == source {
			// target.setParent(parent);
			// exprList.set(i, target);
			return true
		}
	}
	return false
}
