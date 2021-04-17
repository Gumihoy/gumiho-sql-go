package select_

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Where_Clause_0(t *testing.T) {
	sourceSQL := `SELECT 1 + 1 FROM DUAL`
	formatSQL := sql.Format(sourceSQL, db.TiDB)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + 1
FROM DUAL`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}