package explain

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// -------------------------------------------------- DESC ----------------------------------------------------------------------

func TestFormat_DESC_0(t *testing.T) {
	sourceSQL := `DESC City;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DESC City;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


// -------------------------------------------------- DESCRIBE ----------------------------------------------------------------------

func TestFormat_DESCRIBE_0(t *testing.T) {
	sourceSQL := `DESCRIBE City;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DESCRIBE City;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


// -------------------------------------------------- EXPLAIN ----------------------------------------------------------------------
func TestFormat_EXPLAIN_0(t *testing.T) {
	sourceSQL := ` EXPLAIN ANALYZE SELECT * FROM t1 JOIN t2 ON (t1.c1 = t2.c2)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `EXPLAIN ANALYZE
	SELECT *
	FROM t1 JOIN t2 ON (t1.c1 = t2.c2)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
