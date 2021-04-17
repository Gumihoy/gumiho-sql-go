package select_

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

/**
 * https://www.postgresql.org/docs/devel/sql-select.html
 */



func TestFormat_Simple_Query_0(t *testing.T) {
	sourceSQL := `SELECT f.title, f.did, d.name, f.date_prod, f.kind
    FROM distributors d, films f
    WHERE f.did = d.did`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT f.title, f.did, d.name, f.date_prod, f.kind
FROM distributors d, films f
WHERE f.did = d.did`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func TestFormat_GroupBy_Clause_0(t *testing.T) {
	sourceSQL := `SELECT kind, sum(len) AS total FROM films GROUP BY kind;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT kind, sum(len) AS total
FROM films
GROUP BY kind;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_GroupBy_Clause_1(t *testing.T) {
	sourceSQL := `SELECT kind, sum(len) AS total
    FROM films
    GROUP BY kind
    HAVING sum(len) < interval '5 hours';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT kind, sum(len) AS total
    FROM films
    GROUP BY kind
    HAVING sum(len) < interval '5 hours';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_OrderBy_Clause_0(t *testing.T) {
	sourceSQL := `SELECT kind, sum(len) AS total FROM films GROUP BY kind;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT kind, sum(len) AS total
FROM films
GROUP BY kind;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}