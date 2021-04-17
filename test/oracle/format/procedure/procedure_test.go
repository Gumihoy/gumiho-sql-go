package procedure

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Drop(t *testing.T) {
	sourceSQL := `DROP PROCEDURE hr.remove_emp; `
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP PROCEDURE hr.remove_emp;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
