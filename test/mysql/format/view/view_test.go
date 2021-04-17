package index

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// ------------------------------- Alter -------------------------
func TestFormat_AlterView_0(t *testing.T) {
	sourceSQL := `ALTER VIEW v1 as select * from dual;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER VIEW v1
AS
SELECT *
FROM dual;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


// ------------------------------- Create -------------------------
func TestFormat_CreateView_Simple_0(t *testing.T) {
	sourceSQL := `CREATE VIEW test.v AS SELECT * FROM t;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE VIEW test.v
AS
SELECT *
FROM t;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_CreateView_Simple_1(t *testing.T) {
	sourceSQL := `CREATE VIEW v_today (today) AS SELECT CURRENT_DATE;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE VIEW v_today (
	today
)
AS
SELECT CURRENT_DATE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// ------------------------------- Drop -------------------------
func TestFormat_DropView_0(t *testing.T) {
	sourceSQL := `DROP VIEW IF EXISTS index1, index2 RESTRICT`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP VIEW IF EXISTS index1, index2 RESTRICT`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
