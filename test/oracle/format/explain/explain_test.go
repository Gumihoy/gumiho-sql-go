package explain

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_EXPLAIN_0(t *testing.T) {
	sourceSQL := `EXPLAIN PLAN 
    SET STATEMENT_ID = 'Raise in Tokyo' 
    INTO plan_table 
    FOR UPDATE employees 
        SET salary = salary * 1.10 
        WHERE department_id =  
           (SELECT department_id FROM departments
               WHERE location_id = 1700); `
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `EXPLAIN PLAIN
	INTO plan_table
	FOR UPDATE employees
	SET salary = salary * 1.10
	WHERE department_id = (
		SELECT department_id
		FROM departments
		WHERE location_id = 1700
	);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_EXPLAIN_1(t *testing.T) {
	sourceSQL := `EXPLAIN PLAN FOR
  SELECT * FROM sales 
     WHERE time_id BETWEEN :h AND '01-OCT-2000';`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `EXPLAIN PLAIN
	FOR SELECT *
	FROM sales
	WHERE time_id BETWEEN time_id AND time_id;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
