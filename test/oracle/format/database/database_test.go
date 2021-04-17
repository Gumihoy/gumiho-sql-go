package datatype

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)


func Test_CreateDatabase_0(t *testing.T) {
	sourceSQL := `CREATE DATABASE test`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

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
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE DATABASE IF NOT EXISTS test`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



func Test_DropDatabase_0(t *testing.T) {
	sourceSQL := `DROP DATABASE`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP DATABASE`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_DropDatabase_1(t *testing.T) {
	sourceSQL := `DROP DATABASE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP DATABASE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
