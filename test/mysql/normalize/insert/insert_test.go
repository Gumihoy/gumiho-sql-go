package insert

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestNormalize_0(t *testing.T) {
	sourceSQL := `insert into dual(id, name) values(1, 'name');`
	normalizeSQL := sql.NormalizeSQL(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(normalizeSQL)

	targetSQL := `INSERT INTO dual (id, name )
VALUES (?, ?);`

	if !strings.EqualFold(normalizeSQL, targetSQL) {
		t.Error()
	}
}



