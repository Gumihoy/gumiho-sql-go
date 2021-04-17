package typebody

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Drop(t *testing.T) {
	sourceSQL := `DROP TYPE BODY data_typ1;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TYPE BODY data_typ1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
