package datatype

import (
	"fmt"
	"gumihoy.com/sql/dbtype"
	"gumihoy.com/sql/format"
	"testing"
)

func Test_DataType_0(t *testing.T) {
	sql := "select \"133      xxx\" , `x` from dual where ;"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}
