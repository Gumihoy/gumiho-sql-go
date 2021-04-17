package index

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)
// ------------------------------- Drop -------------------------

func TestFormat_CreateIndex_0(t *testing.T) {
	sourceSQL := `CREATE INDEX idx1 ON t1 ((col1 + col2));`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 ADD COLUMN c2 INT, ADD COLUMN c3 INT, ALGORITHM = INSTANT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
// ------------------------------- Alter -------------------------



// ------------------------------- Drop -------------------------
func TestFormat_DropIndex_0(t *testing.T) {
	sourceSQL := "DROP INDEX `PRIMARY` ON t;"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := "DROP INDEX `PRIMARY` ON t;"

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_DropIndex_1(t *testing.T) {
	sourceSQL := "DROP INDEX `PRIMARY` ON t ALGORITHM = DEFAULT;"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := "DROP INDEX `PRIMARY` ON t ALGORITHM = DEFAULT;"

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}