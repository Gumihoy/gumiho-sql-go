package explain

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_HELP_0(t *testing.T) {
	sourceSQL := `HELP 'contents'`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `HELP 'contents'`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_HELP_1(t *testing.T) {
	sourceSQL := `HELP 'data types';
HELP 'ascii;'
HELP 'create table';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `HELP 'data types';
HELP 'ascii;';
HELP 'create table';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

