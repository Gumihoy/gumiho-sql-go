package insert

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func Test_Insert_Values_0(t *testing.T) {
	sourceSQL := `INSERT INTO tbl_name (col1,col2) VALUES(15,col1*2);`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `INSERT INTO tbl_name (col1, col2 )
VALUES (15, col1 * 2);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



