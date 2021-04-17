package select_

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// -------------------------- Alter --------------------------

// -------------------------- Create --------------------------
func TestFormat_Create_0(t *testing.T) {
	sourceSQL := `DROP SEQUENCE oe.customers_seq;`
	formatSQL := sql.Format(sourceSQL, db.SQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SEQUENCE oe.customers_seq;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// -------------------------- Drop --------------------------
func TestFormat_Drop_0(t *testing.T) {
	sourceSQL := `DROP SEQUENCE oe.customers_seq;`
	formatSQL := sql.Format(sourceSQL, db.SQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SEQUENCE oe.customers_seq;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Drop_1(t *testing.T) {
	sourceSQL := `DROP SEQUENCE oe.customers_seq CASCADE;`
	formatSQL := sql.Format(sourceSQL, db.SQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SEQUENCE oe.customers_seq CASCADE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
