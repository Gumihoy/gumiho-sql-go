package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/datatype"
)


func AcceptDataType(visitor ISQLVisitor, x *datatype.SQLDataType) {
	if visitor.VisitDataType(visitor, x) {
		Accept(visitor, x.Name())

		for _, child := range x.Precisions() {
			Accept(visitor, child)
		}
	}
}

func AcceptIntervalDataType(visitor ISQLVisitor, x *datatype.SQLIntervalDataType) {
	if visitor.VisitIntervalDataType(visitor, x) {
		Accept(visitor, x.Start())
		Accept(visitor, x.End())
	}
}
func AcceptIntervalDataTypeField(visitor ISQLVisitor, x *datatype.SQLIntervalDataTypeField) {
	if visitor.VisitIntervalDataTypeField(visitor, x) {

		for _, child := range x.Precisions() {
			Accept(visitor, child)
		}
	}
}


func AcceptDateDataType(visitor ISQLVisitor, x *datatype.SQLDateDataType) {
	if visitor.VisitDateDataType(visitor, x) {

	}
}
func AcceptDateTimeDataType(visitor ISQLVisitor, x *datatype.SQLDateTimeDataType) {
	if visitor.VisitDateTimeDataType(visitor, x) {

	}
}
func AcceptTimeDataType(visitor ISQLVisitor, x *datatype.SQLTimeDataType) {
	if visitor.VisitTimeDataType(visitor, x) {

	}
}
func AcceptTimestampDataType(visitor ISQLVisitor, x *datatype.SQLTimestampDataType) {
	if visitor.VisitTimestampDataType(visitor, x) {

	}
}