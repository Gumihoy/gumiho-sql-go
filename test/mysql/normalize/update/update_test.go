package update

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestNormalize_0(t *testing.T) {
	sourceSQL := `update dual set name = 'name', code = 1 where id BETWEEN 1 and 100 and cd = abs(1)`
	normalizeSQL := sql.NormalizeSQL(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(normalizeSQL)

	targetSQL := `UPDATE dual
SET name = ?, code = ?
WHERE id BETWEEN ? AND ? AND cd = abs(?)`

	if !strings.EqualFold(normalizeSQL, targetSQL) {
		t.Error()
	}
}