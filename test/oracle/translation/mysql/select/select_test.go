package select_

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/translation"
	"strings"
	"testing"
)

func TestTranslation_Simple_Query_0(t *testing.T) {
	sourceSQL := `SELECT * 
   FROM "EMPLOYEES"
   WHERE department_id = 30 group by last_name
   ORDER BY last_name ;`
	result := translation.OracleToMySQL(sourceSQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(result.TargetSQL)

	targetSQL := `SELECT *
FROM employees
WHERE department_id = 30
GROUP BY last_name
ORDER BY last_name;`

	if !strings.EqualFold(result.TargetSQL, targetSQL) {
		t.Error()
	}
}
