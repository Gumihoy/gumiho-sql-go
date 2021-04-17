package delete

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Delete_OrderBy_0(t *testing.T) {
	sourceSQL := `DELETE FROM somelog WHERE user = 'jcole'`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM somelog
WHERE user = 'jcole'`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}

}