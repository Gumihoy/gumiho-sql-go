package procedure

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Drop(t *testing.T) {
	sourceSQL := `DROP FUNCTION oe.SecondMax; `
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP FUNCTION oe.SecondMax;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
