package update

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func Test_Simple_Update_0(t *testing.T) {
	sourceSQL := `UPDATE t1 SET col1 = col1 + 1;`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `UPDATE t1
SET col1 = col1 + 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Simple_Update_1(t *testing.T) {
	sourceSQL := `UPDATE t1 SET col1 = col1 + 1, col2 = col1;`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `UPDATE t1
SET col1 = col1 + 1, col2 = col1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Where_Clause_0(t *testing.T) {
	sourceSQL := `UPDATE items,month SET items.price=month.price
WHERE items.id=month.id;`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `UPDATE items, month
SET items.price = month.price
WHERE items.id = month.id;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_SubQuery_0(t *testing.T) {
	sourceSQL := `UPDATE items,
       (SELECT id FROM items
        WHERE id IN
            (SELECT id FROM items
             WHERE retail / wholesale >= 1.3 AND quantity < 100))
        AS discounted
SET items.retail = items.retail * 0.9
WHERE items.id = discounted.id;`
	formatSQL := sql.Format(sourceSQL, db.SQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `UPDATE items, (
	SELECT id
	FROM items
	WHERE id IN (
		SELECT id
		FROM items
		WHERE retail / wholesale <= 1.3 AND quantity < 100
	)
) AS discounted
SET items.retail = items.retail * 0.9
WHERE items.id = discounted.id;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

