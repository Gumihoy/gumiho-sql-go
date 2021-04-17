package comment

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_CommentOnColumn_0(t *testing.T) {
	sourceSQL := `COMMENT ON COLUMN employees.job_id IS 'abbreviated job title';`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `COMMENT ON COLUMN employees.job_id IS 'abbreviated job title';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_CommentOnColumn_1(t *testing.T) {
	sourceSQL := `COMMENT ON COLUMN employees.job_id IS '';`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `COMMENT ON COLUMN employees.job_id IS '';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CommentOnTable_0(t *testing.T) {
	sourceSQL := `COMMENT ON TABLE employees.job_id IS 'abbreviated job title';`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `COMMENT ON TABLE employees.job_id IS 'abbreviated job title';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}