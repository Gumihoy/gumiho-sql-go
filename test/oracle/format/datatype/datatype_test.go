package datatype

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func Test_DataType_0(t *testing.T) {
	sourceSQL := "SELECT cast(2 as date)"
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT cast(2 AS DATE)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
