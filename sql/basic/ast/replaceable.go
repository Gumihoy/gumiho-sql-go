package ast

type ISQLReplaceable interface {
	ReplaceChild(source ISQLObject, target ISQLObject) bool
}
