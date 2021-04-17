package index

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_Creating_an_Index_0(t *testing.T) {
	sourceSQL := `CREATE INDEX ord_customer_ix
   ON orders (customer_id);`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE INDEX ord_customer_ix
   ON orders (customer_id);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Creating_Compressing_an_Index_1(t *testing.T) {
	sourceSQL := `CREATE INDEX ord_customer_ix_demo 
   ON orders (customer_id, sales_rep_id)
   COMPRESS 1;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE INDEX ord_customer_ix_demo 
   ON orders (customer_id, sales_rep_id)
   COMPRESS 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Creating_an_Index_in_NOLOGGING_Mode_2(t *testing.T) {
	sourceSQL := `CREATE INDEX ord_customer_ix_demo
   ON orders (order_mode)
   NOSORT
   NOLOGGING;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE INDEX ord_customer_ix_demo
   ON orders (order_mode)
   NOSORT
   NOLOGGING;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Creating_a_Cluster_Index_3(t *testing.T) {
	sourceSQL := `CREATE INDEX idx_personnel ON CLUSTER personnel;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE INDEX idx_personnel ON CLUSTER personnel;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



func TestFormat_Drop(t *testing.T) {
	sourceSQL := `DROP TYPE person_t FORCE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TYPE person_t FORCE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
