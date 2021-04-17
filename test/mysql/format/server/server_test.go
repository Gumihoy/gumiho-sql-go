package server

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// ------------------------------- Alter -------------------------
func TestFormat_Alter_0(t *testing.T) {
	sourceSQL := `ALTER SERVER s OPTIONS (USER 'sally');`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER SERVER s OPTIONS (USER 'sally');`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
// ------------------------------- Create -------------------------
func TestFormat_Create_0(t *testing.T) {
	sourceSQL := `CREATE SERVER s
FOREIGN DATA WRAPPER mysql
OPTIONS (USER 'Remote', HOST '198.51.100.106', DATABASE 'test');`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE SERVER s
FOREIGN DATA WRAPPER mysql
OPTIONS (USER 'Remote', HOST '198.51.100.106', DATABASE 'test');`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
// ------------------------------- Drop -------------------------
func TestFormat_Drop_0(t *testing.T) {
	sourceSQL := `DROP SERVER server_name`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SERVER server_name`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Drop_1(t *testing.T) {
	sourceSQL := `DROP SERVER IF EXISTS server_name`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SERVER IF EXISTS server_name`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
