package operator

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"testing"
)

// https://dev.mysql.com/doc/refman/8.0/en/non-typed-operators.html

// https://dev.mysql.com/doc/refman/8.0/en/bit-functions.html#operator_bitwise-and
func Test_Operator_Bit_0(t *testing.T) {
	sourceSQL := "SELECT 127 | 128, 128 << 2, BIT_COUNT(15);"
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(formatSQL)
}

func Test_Operator_Bit_1(t *testing.T) {
	sourceSQL := "SELECT BIT_COUNT('15');"
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(formatSQL)
}

func Test_Operator_Greater_Than_0(t *testing.T) {
	sourceSQL := "SELECT BIT_COUNT('15');"
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(formatSQL)
}

func Test_Operator_Right_Shift_0(t *testing.T) {
	sourceSQL := "SELECT 4 >> 2;"
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(formatSQL)
}
