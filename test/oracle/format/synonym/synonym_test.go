package synonym

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// -------------------------- Alter --------------------------
func TestFormat_Alter_0(t *testing.T) {
	sourceSQL := `ALTER SYNONYM offices COMPILE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER SYNONYM offices COMPILE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Alter_1(t *testing.T) {
	sourceSQL := `ALTER PUBLIC SYNONYM emp_table COMPILE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER PUBLIC SYNONYM emp_table COMPILE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Alter_2(t *testing.T) {
	sourceSQL := `ALTER SYNONYM offices NONEDITIONABLE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER SYNONYM offices NONEDITIONABLE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// -------------------------- Create --------------------------
func TestFormat_Create_0(t *testing.T) {
	sourceSQL := `CREATE SYNONYM offices 
   FOR hr.locations;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE SYNONYM offices FOR hr.locations;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Create_1(t *testing.T) {
	sourceSQL := `CREATE PUBLIC SYNONYM emp_table 
   FOR hr.employees@remote.us.example.com;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE PUBLIC SYNONYM emp_table FOR hr.employees@remote.us.example.com;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Create_2(t *testing.T) {
	sourceSQL := `CREATE PUBLIC SYNONYM customers FOR oe.customers;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE PUBLIC SYNONYM customers FOR oe.customers;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// -------------------------- Drop --------------------------
func TestFormat_Drop_0(t *testing.T) {
	sourceSQL := `DROP PUBLIC SYNONYM customers;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP PUBLIC SYNONYM customers;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Drop_1(t *testing.T) {
	sourceSQL := `DROP public SYNONYM oe.customers_seq force;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP PUBLIC SYNONYM oe.customers_seq FORCE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}