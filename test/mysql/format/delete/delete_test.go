package delete

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Delete_OrderBy_0(t *testing.T) {
	sourceSQL := `DELETE FROM somelog WHERE user = 'jcole'
ORDER BY timestamp_column LIMIT 1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM somelog
WHERE user = 'jcole'
ORDER BY timestamp_column
LIMIT 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}

}



func Test_Delete_MultiTable_0(t *testing.T) {
	sourceSQL := `DELETE t1, t2 FROM t1 INNER JOIN t2 INNER JOIN t3
WHERE t1.id=t2.id AND t2.id=t3.id;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE t1, t2 FROM t1 INNER JOIN t2 INNER JOIN t3
WHERE t1.id = t2.id AND t2.id = t3.id;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}

}

func Test_Delete_MultiTable_1(t *testing.T) {
	sourceSQL := `DELETE FROM t1, t2 USING t1 INNER JOIN t2 INNER JOIN t3
WHERE t1.id=t2.id AND t2.id=t3.id;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM t1, t2 USING t1 INNER JOIN t2 INNER JOIN t3
WHERE t1.id = t2.id AND t2.id = t3.id;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}

}