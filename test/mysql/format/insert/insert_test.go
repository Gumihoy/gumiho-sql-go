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
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `INSERT INTO tbl_name (col1, col2 )
VALUES (15, col1 * 2);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}








func Test_ON_DUPLICATE_KEY_UPDATE_0(t *testing.T) {
	sourceSQL := `INSERT INTO t1 (a,b,c) VALUES (1,2,3)
  ON DUPLICATE KEY UPDATE c=c+1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `INSERT INTO t1 (a, b, c )
VALUES (1, 2, 3) ON DUPLICATE KEY UPDATE c = c + 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func Test_ON_DUPLICATE_KEY_UPDATE_1(t *testing.T) {
	sourceSQL := `INSERT INTO t1 (a,b,c) VALUES (1,2,3),(4,5,6)
  ON DUPLICATE KEY UPDATE c=VALUES(a)+VALUES(b);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `INSERT INTO t1 (a, b, c )
VALUES (1, 2, 3), (4, 5, 6) ON DUPLICATE KEY UPDATE c = VALUES(a) + VALUES(b);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}




