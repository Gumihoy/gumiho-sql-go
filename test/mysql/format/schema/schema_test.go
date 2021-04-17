package datatype

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)






func Test_DropSchema_0(t *testing.T) {
	sourceSQL := `DROP SCHEMA test`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SCHEMA test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_DropSchema_1(t *testing.T) {
	sourceSQL := `DROP SCHEMA IF EXISTS test`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SCHEMA IF EXISTS test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
