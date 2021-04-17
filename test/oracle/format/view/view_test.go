package view

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// ------------------------------- Alter -------------------------
func TestFormat_AlterView_0(t *testing.T) {
	sourceSQL := `ALTER VIEW customer_ro
    COMPILE;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER VIEW customer_ro COMPILE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// ------------------------------- Create -------------------------

func TestFormat_Creating_a_View_0(t *testing.T) {
	sourceSQL := `CREATE VIEW emp_view AS 
   SELECT last_name, salary*12 annual_salary
   FROM employees 
   WHERE department_id = 20;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE hash_products 
    ( product_id          NUMBER(6)   PRIMARY KEY
    , product_name        VARCHAR2(50) 
    , product_description VARCHAR2(2000) 
    , category_id         NUMBER(2) 
    , weight_class        NUMBER(1) 
    , warranty_period     INTERVAL YEAR TO MONTH 
    , supplier_id         NUMBER(6) 
    , product_status      VARCHAR2(20) 
    , list_price          NUMBER(8,2) 
    , min_price           NUMBER(8,2) 
    , catalog_url         VARCHAR2(50) 
    , CONSTRAINT          product_status_lov_demo 
                          CHECK (product_status in ('orderable' 
                                                  ,'planned' 
                                                  ,'under development' 
                                                  ,'obsolete') 
 ) ) 
 PARTITION BY HASH (product_id) 
 PARTITIONS 4 
 STORE IN (tbs_01, tbs_02, tbs_03, tbs_04); `

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


// ------------------------------- Drop -------------------------
func TestFormat_DropView_0(t *testing.T) {
	sourceSQL := `DROP VIEW emp_view;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP VIEW emp_view;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_DropView_1(t *testing.T) {
	sourceSQL := `DROP VIEW emp_view CASCADE CONSTRAINTS;`
	formatSQL := sql.Format(sourceSQL, db.Oracle)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP VIEW emp_view CASCADE CONSTRAINTS;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
