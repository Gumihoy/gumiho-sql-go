package role

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_CreateRole_0(t *testing.T) {
	sourceSQL := `CREATE ROLE 'administrator', 'developer';
CREATE ROLE 'webapp'@'localhost';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE ROLE 'administrator', 'developer';
CREATE ROLE 'webapp'@'localhost';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



func TestFormat_DropRole_0(t *testing.T) {
	sourceSQL := `DROP ROLE 'administrator', 'developer';
DROP ROLE 'webapp'@'localhost';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP ROLE 'administrator', 'developer';
DROP ROLE 'webapp'@'localhost';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
