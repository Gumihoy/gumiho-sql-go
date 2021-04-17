package datatype

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)
// ------------------------------- Alter -------------------------
func TestFormat_AlterDatabase_0(t *testing.T) {
	sourceSQL := `ALTER DATABASE mydb READ ONLY = 1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER DATABASE mydb READ ONLY = 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterDatabase_1(t *testing.T) {
	sourceSQL := `ALTER DATABASE mydb DEFAULT CHARACTER SET = 'utf8mb4';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER DATABASE mydb DEFAULT CHARACTER SET = 'utf8mb4';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// ------------------------------- Create -------------------------
func Test_CreateDatabase_0(t *testing.T) {
	sourceSQL := `CREATE DATABASE test`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE DATABASE test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_CreateDatabase_1(t *testing.T) {
	sourceSQL := `CREATE DATABASE IF NOT EXISTS test`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE DATABASE IF NOT EXISTS test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


// ------------------------------- Drop -------------------------
func Test_DropDatabase_0(t *testing.T) {
	sourceSQL := `DROP DATABASE test`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP DATABASE test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_DropDatabase_1(t *testing.T) {
	sourceSQL := `DROP DATABASE IF EXISTS test`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP DATABASE IF EXISTS test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
