package format

import (
	"fmt"
	"gumihoy.com/sql/db"
	"testing"
)

func TestFormat(t *testing.T) {
	sql := "select *, idaaaaaaaaaeeeeeee,id, id, idaffffffffffffffffffffffffff from dual where id = ?;"
	formatSQL := Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}
