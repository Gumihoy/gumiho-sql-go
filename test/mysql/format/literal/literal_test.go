package literal

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	db "github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func Test_Literal_String_0(t *testing.T) {
	sourceSQL := "SELECT 'hello', '\"hello\"', '\"\"hello\"\"', 'hel''lo', '\\'hello';"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 'hello', '"hello"', '""hello""', 'hel\'lo', '\'hello';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Literal_String_1(t *testing.T) {
	sourceSQL := "SELECT 'This\\nIs\\nFour\\nLines';"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 'This\nIs\nFour\nLines';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Literal_String_2(t *testing.T) {
	sourceSQL := "SELECT 'disappearing\\ backslash'"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 'disappearing\ backslash'`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



func Test_Literal_Numeric_0(t *testing.T) {
	sourceSQL := "SELECT (-11-5)"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT (-11 - 5)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func Test_Literal_Numeric_1(t *testing.T) {
	sourceSQL := "SELECT 1.2E3, 1.2E-3, -1.2E3, -1.2E-3"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1.2E3, 1.2E-3, -1.2E3, -1.2E-3`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



func Test_Literal_DATE_0(t *testing.T) {
	sourceSQL := "SELECT DATE '2015-07-21'"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT DATE '2015-07-21'`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Literal_TIME_0(t *testing.T) {
	sourceSQL := "SELECT TIME ''"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT TIME ''`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Literal_TIMESTAMP_0(t *testing.T) {
	sourceSQL := "SELECT TIMESTAMP ''"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT TIMESTAMP ''`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}





func Test_Literal_BOOLEAN_0(t *testing.T) {
	sourceSQL := "SELECT TRUE, true, FALSE, false, aaaaaaaaaaaaaaaaaaaaaaaa;"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT true, true, false, false, aaaaaaaaaaaaaaaaaaaaaaaa;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}




