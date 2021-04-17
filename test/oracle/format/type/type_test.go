package type_

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Drop(t *testing.T) {
	sourceSQL := `DROP TYPE person_t FORCE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TYPE person_t FORCE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
