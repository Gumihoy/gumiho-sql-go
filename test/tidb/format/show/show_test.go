package show

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_ShowCreateDatabase_0(t *testing.T) {
	sourceSQL := `show create database d`
	formatSQL := sql.Format(sourceSQL, db.TiDB)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SHOW CREATE DATABASE d`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_ShowCreateEvent_0(t *testing.T) {
	sourceSQL := `show create event d`
	formatSQL := sql.Format(sourceSQL, db.TiDB)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SHOW CREATE event d`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
