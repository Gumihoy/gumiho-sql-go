package delete

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestNormalize_0(t *testing.T) {
	sourceSQL := `delete from dual where id = 1 and name like 'name' and code in (1, 2, 3);`
	normalizeSQL := sql.NormalizeSQL(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(normalizeSQL)

	targetSQL := `DELETE FROM dual
WHERE id = ? AND name LIKE ? AND code IN (?);`

	if !strings.EqualFold(normalizeSQL, targetSQL) {
		t.Error()
	}
}
