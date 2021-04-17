package explain

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_USE_0(t *testing.T) {
	sourceSQL := `USE db1;
SELECT COUNT(*) FROM mytable;   # selects from db1.mytable
USE db2;
SELECT COUNT(*) FROM mytable;   # selects from db2.mytable`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `USE db1;
SELECT COUNT(*)
FROM mytable;
# selects from db1.mytable
USE db2;
SELECT COUNT(*)
FROM mytable; # selects from db2.mytable`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}