package role

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_CreateUser_0(t *testing.T) {
	sourceSQL := `CREATE USER 'jeffrey'@'localhost' IDENTIFIED BY 'password';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE USER 'jeffrey'@'localhost' IDENTIFIED BY 'password';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_CreateUser_1(t *testing.T) {
	sourceSQL := `CREATE USER 'jeffrey'@'localhost'
  IDENTIFIED BY 'new_password' PASSWORD EXPIRE;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE USER 'jeffrey'@'localhost'
  IDENTIFIED BY 'new_password' PASSWORD EXPIRE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateUser_11(t *testing.T) {
	sourceSQL := `CREATE USER 'joe'@'10.0.0.1' DEFAULT ROLE administrator, developer;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE USER 'joe'@'10.0.0.1' DEFAULT ROLE administrator, developer;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func TestFormat_DropRole_0(t *testing.T) {
	sourceSQL := `DROP USER 'administrator', 'developer';
DROP USER 'webapp'@'localhost';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP USER 'administrator', 'developer';
DROP USER 'webapp'@'localhost';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
