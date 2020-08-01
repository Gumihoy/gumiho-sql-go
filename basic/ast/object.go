package ast

import "gumihoy.com/sql/basic/visitor"

type ISQLObject interface {
	Accept(visitor visitor.ISQLVisitor)
	IsAfterSemi() bool
	SetAfterSemi(semi bool)
}

type SQLObject struct {
	Semi bool
}

func (s *SQLObject) Accept(visitor visitor.ISQLVisitor) {
}

func (s *SQLObject) IsAfterSemi() bool {
	return s.Semi
}

func (s *SQLObject) SetAfterSemi(semi bool) {
	s.Semi = semi
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
				//delete(exprs, i)
				return true
			}
		}
		return false
	}

	for i := 0; i < len(exprs); i++ {
		if exprs[i] == source {
			//target.setParent(parent);
			//exprList.set(i, target);
			return true
		}
	}
	return false
}
