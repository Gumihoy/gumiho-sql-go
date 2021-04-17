package delete

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func Test_Deleting_Rows_0(t *testing.T) {
	sourceSQL := `DELETE FROM product_descriptions
   WHERE language_id = 'AR';`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM product_descriptions
WHERE language_id = 'AR';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Deleting_Rows_1(t *testing.T) {
	sourceSQL := `DELETE FROM employees
   WHERE job_id = 'SA_REP'
   AND commission_pct < .2;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM employees
WHERE job_id = 'SA_REP' AND commission_pct < 2;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Deleting_Rows_2(t *testing.T) {
	sourceSQL := `DELETE FROM (SELECT * FROM employees)
   WHERE job_id = 'SA_REP'
   AND commission_pct < .2;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM (
	SELECT *
	FROM employees
)
WHERE job_id = 'SA_REP' AND commission_pct < 2;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func Test_Deleting_Rows_from_a_Remote_Database_0(t *testing.T) {
	sourceSQL := `DELETE FROM hr.locations@remote
   WHERE location_id > 3000;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM hr.locations@remote
WHERE location_id > 3000;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func Test_Deleting_Rows_from_a_Partition_0(t *testing.T) {
	sourceSQL := `DELETE FROM sales PARTITION (sales_q1_1998)
   WHERE amount_sold > 1000;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM sales PARTITION(sales_q1_1998)
WHERE amount_sold > 1000;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Using_the_RETURNING_Clause_0(t *testing.T) {
	sourceSQL := `DELETE FROM employees
   WHERE job_id = 'SA_REP' 
   AND hire_date + TO_YMINTERVAL('01-00') < SYSDATE 
   RETURNING salary INTO :bnd1;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE FROM employees
   WHERE job_id = 'SA_REP' 
   AND hire_date + TO_YMINTERVAL('01-00') < SYSDATE 
   RETURNING salary INTO :bnd1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func Test_Deleting_Data_from_a_Table_0(t *testing.T) {
	sourceSQL := `DELETE product_price_history pp 
WHERE  (product_id, currency_code, effective_from_date) 
   IN (SELECT product_id, currency_code, Max(effective_from_date) 
       FROM   product_price_history 
       GROUP BY product_id, currency_code);`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DELETE product_price_history pp
WHERE (product_id, currency_code, effective_from_date) IN (
	SELECT product_id, currency_code, Max(effective_from_date)
	FROM product_price_history
	GROUP BY product_id, currency_code
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}