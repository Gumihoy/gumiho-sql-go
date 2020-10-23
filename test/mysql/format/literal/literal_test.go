package literal

import (
	"fmt"
	"gumihoy.com/sql/dbtype"
	"gumihoy.com/sql/format"
	"testing"
)

func Test_Literal_String_0(t *testing.T) {
	sql := "SELECT 'hello', '\"hello\"', '\"\"hello\"\"', 'hel''lo', '\\'hello';"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}

func Test_Literal_String_1(t *testing.T) {
	sql := "SELECT 'This\\nIs\\nFour\\nLines';"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(sql, formatSQL)
}

func Test_Literal_String_2(t *testing.T) {
	sql := "SELECT 'disappearing\\ backslash'"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}



func Test_Literal_Numeric_0(t *testing.T) {
	sql := "SELECT (-11-5)"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}
func Test_Literal_Numeric_1(t *testing.T) {
	sql := "SELECT 1.2E3, 1.2E-3, -1.2E3, -1.2E-3"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}



func Test_Literal_DATE_0(t *testing.T) {
	sql := "SELECT DATE '2015-07-21'"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}

func Test_Literal_TIME_0(t *testing.T) {
	sql := "SELECT TIME ''"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}

func Test_Literal_TIMESTAMP_0(t *testing.T) {
	sql := "SELECT TIMESTAMP ''"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}





func Test_Literal_BOOLEAN_0(t *testing.T) {
	sql := "SELECT TRUE, true, FALSE, false, aaaaaaaaaaaaaaaaaaaaaaaa;"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}




