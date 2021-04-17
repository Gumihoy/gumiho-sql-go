package comment

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_SharpComment_0(t *testing.T) {
	sourceSQL := `SELECT 1+1;     # This comment continues to the end of line`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + 1; # This comment continues to the end of line`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_MinusComment_0(t *testing.T) {
	sourceSQL := ` SELECT 1+1;     -- This comment continues to the end of line
select 1+2;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + 1;
-- This comment continues to the end of line
SELECT 1 + 2;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_MultiLineComment_0(t *testing.T) {
	sourceSQL := ` SELECT 1 /* this is an in-line comment */ + 1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 /* this is an in-line comment */ + 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_MultiLineComment_1(t *testing.T) {
	sourceSQL := `SELECT 1+
/*
this is a
multiple-line comment
*/
1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + /*
this is a
multiple-line comment
*/ 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_MultiLineComment_2(t *testing.T) {
	sourceSQL := `SELECT 1+
/*
this is a
multiple-line comment
*/
1 /* end */, /** af */2;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + /*
this is a
multiple-line comment
*/ 1 /* end */, /** af */ 2;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
