package identifier

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func Test_Identifier_1(t *testing.T) {
	sourceSQL := "select \"133      xxx\" , `x` from dual where id = ?;"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := "SELECT \"133      xxx\", `x`\n" +
		"FROM dual\n" +
		"WHERE id = ?;"

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
