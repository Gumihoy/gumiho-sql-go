package comment

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Single_Line_Comments_0(t *testing.T) {
	sourceSQL := `-- DELETE FROM employees WHERE comm_pct IS NULL
select * from dual;`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `-- DELETE FROM employees WHERE comm_pct IS NULL
SELECT *
FROM dual;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Single_Line_Comments_1(t *testing.T) {
	sourceSQL := `select -- a
 a  from dual;`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT -- a
	a
FROM dual;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Multiline_Comments_0(t *testing.T) {
	sourceSQL := `/*
  IF 2 + 2 = 4 THEN
    some_condition := TRUE;
  -- We expect this THEN to always be performed
  END IF;
*/`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `/*
  IF 2 + 2 = 4 THEN
    some_condition := TRUE;
  -- We expect this THEN to always be performed
  END IF;
*/`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Multiline_Comments_1(t *testing.T) {
	sourceSQL := `select /* a */
 a  from dual;`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT /* a */ a
FROM dual;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
