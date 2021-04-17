package ast

type ISQLCloneable interface {
	Clone() ISQLObject
}