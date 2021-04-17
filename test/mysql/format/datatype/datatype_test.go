package datatype

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func Test_DataType_0(t *testing.T) {
	sourceSQL := `CREATE TABLE test (
					col1 NUMBER(5,2),
					col2 FLOAT(5)
				)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE test (
	col1 NUMBER(5, 2),
	col2 FLOAT(5)
)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
