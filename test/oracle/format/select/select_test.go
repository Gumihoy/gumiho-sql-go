package select_

import (
	"fmt"
	"gumihoy.com/sql/dbtype"
	"gumihoy.com/sql/format"
	"testing"
)

func TestFormat_1(t *testing.T) {
	sql := "select *, \"idaaaaaaaaaeeeeeee\", id, id, idaffffffffffffffffffffffffff from dual where id = ?;"
	formatSQL := format.Format(sql, db.Oracle)
	fmt.Println(formatSQL)
}


