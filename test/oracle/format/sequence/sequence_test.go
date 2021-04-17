package select_

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// -------------------------- Alter --------------------------
func TestFormat_Alter_0(t *testing.T) {
	sourceSQL := `ALTER SEQUENCE customers_seq 
   MAXVALUE 1500;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER SEQUENCE customers_seq
	MAXVALUE 1500;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Alter_1(t *testing.T) {
	sourceSQL := `ALTER SEQUENCE customers_seq 
   CYCLE
   CACHE 5;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER SEQUENCE customers_seq
	CYCLE
	CACHE 5;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// -------------------------- Create --------------------------
func TestFormat_Create_0(t *testing.T) {
	sourceSQL := `CREATE SEQUENCE customers_seq
 START WITH     1000
 INCREMENT BY   1
 NOCACHE
 NOCYCLE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE SEQUENCE customers_seq
	START WITH 1000
	INCREMENT BY 1
	NOCACHE
	NOCYCLE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// -------------------------- Drop --------------------------
func TestFormat_Drop_0(t *testing.T) {
	sourceSQL := `DROP SEQUENCE oe.customers_seq;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP SEQUENCE oe.customers_seq;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
