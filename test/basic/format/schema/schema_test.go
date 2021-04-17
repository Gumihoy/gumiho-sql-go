package datatype

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)


func TestFormat_CreateSchema_0(t *testing.T) {
	sourceSQL := `create SCHEMA test`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE SCHEMA test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



func Test_DropSchema_0(t *testing.T) {
	sourceSQL := `DROP SCHEMA test`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SCHEMA test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_DropSchema_1(t *testing.T) {
	sourceSQL := `DROP SCHEMA test CASCADE `
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SCHEMA test CASCADE `

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
