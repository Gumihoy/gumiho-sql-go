package function

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_CreateFunction_0(t *testing.T) {
	sourceSQL := ``
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := ``

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}

}


func TestFormat_DropFunction_0(t *testing.T) {
	sourceSQL := "drop function f"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP FUNCTION f`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}

}