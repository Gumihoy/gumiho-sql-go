package select_

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)



func TestFormat_With_Clause_0(t *testing.T) {
	sourceSQL := `WITH
  cte1 AS (SELECT a, b FROM table1),
  cte2 AS (SELECT c, d FROM table2)
SELECT b, d FROM cte1 JOIN cte2
WHERE cte1.a = cte2.c;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `WITH
	cte1 AS (
		SELECT a, b
		FROM table1
	),
	cte2 AS (
		SELECT c, d
		FROM table2
	)
SELECT b, d
FROM cte1 JOIN cte2
WHERE cte1.a = cte2.c;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}

}

func TestFormat_With_Clause_1(t *testing.T) {
	sourceSQL := `WITH cte (col1, col2) AS
(
  SELECT 1, 2
  UNION ALL
  SELECT 3, 4
)
SELECT col1, col2 FROM cte;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `WITH
	cte (col1, col2) AS (
		SELECT 1, 2
		UNION ALL
		SELECT 3, 4
	)
SELECT col1, col2
FROM cte;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_With_Clause_2(t *testing.T) {
	sourceSQL := `WITH cte AS
(
  SELECT 1 AS col1, 2 AS col2
  UNION ALL
  SELECT 3, 4
)
SELECT col1, col2 FROM cte;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `WITH
	cte AS (
		SELECT 1 AS col1, 2 AS col2
		UNION ALL
		SELECT 3, 4
	)
SELECT col1, col2
FROM cte;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_With_Clause_3(t *testing.T) {
	sourceSQL := `WITH cte1 AS (SELECT 1)
SELECT * FROM (WITH cte2 AS (SELECT 2) SELECT * FROM cte2 JOIN cte1) AS dt;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `WITH
	cte1 AS (
		SELECT 1
	)
SELECT *
FROM (
	WITH
		cte2 AS (
			SELECT 2
		)
	SELECT *
	FROM cte2 JOIN cte1
) AS dt;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func TestFormat_With_Clause_4(t *testing.T) {
	sourceSQL := `WITH RECURSIVE cte (n) AS
(
  SELECT 1
  UNION ALL
  SELECT n + 1 FROM cte WHERE n < 5
)
SELECT * FROM cte;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `WITH RECURSIVE
	cte (n) AS (
		SELECT 1
		UNION ALL
		SELECT n + 1
		FROM cte
		WHERE n < 5
	)
SELECT *
FROM cte;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_With_Clause_5(t *testing.T) {
	sourceSQL := `WITH RECURSIVE cte AS
(
  SELECT 1 AS n, 'abc' AS str
  UNION ALL
  SELECT n + 1, CONCAT(str, str) FROM cte WHERE n < 3
)
SELECT * FROM cte;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `WITH RECURSIVE
	cte AS (
		SELECT 1 AS n, 'abc' AS str
		UNION ALL
		SELECT n + 1, CONCAT(str, str)
		FROM cte
		WHERE n < 3
	)
SELECT *
FROM cte;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_With_Clause_6(t *testing.T) {
	sourceSQL := `WITH RECURSIVE cte AS
(
  SELECT 1 AS n, CAST('abc' AS CHAR(20)) AS str
  UNION ALL
  SELECT n + 1, CONCAT(str, str) FROM cte WHERE n < 3
)
SELECT * FROM cte;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `WITH RECURSIVE
	cte AS (
		SELECT 1 AS n, CAST('abc' AS CHAR(20)) AS str
		UNION ALL
		SELECT n + 1, CONCAT(str, str)
		FROM cte
		WHERE n < 3
	)
SELECT *
FROM cte;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Simple_Query_0(t *testing.T) {
	sourceSQL := `SELECT 1 + 1`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + 1`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Simple_Query_1(t *testing.T) {
	sourceSQL := `SELECT 1 + 1 FROM DUAL`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + 1
FROM DUAL`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Where_Clause_0(t *testing.T) {
	sourceSQL := `SELECT 1 + 1 FROM DUAL`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1 + 1
FROM DUAL`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_GroupBy_Clause_0(t *testing.T) {
	sourceSQL := `SELECT COUNT(col1) AS col2 FROM t GROUP BY col2 HAVING col2 = 2;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT COUNT(col1) AS col2
FROM t
GROUP BY col2 HAVING col2 = 2;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_GroupBy_Clause_1(t *testing.T) {
	sourceSQL := `SELECT user, MAX(salary) FROM users
  GROUP BY user HAVING MAX(salary) > 10;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT user, MAX(salary)
FROM users
GROUP BY user HAVING MAX(salary) > 10;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_OrderBy_Clause_0(t *testing.T) {
	sourceSQL := `SELECT college, region, seed FROM tournament
  ORDER BY region, seed;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT college, region, seed
FROM tournament
ORDER BY region, seed;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_OrderBy_Clause_1(t *testing.T) {
	sourceSQL := `SELECT a, b, COUNT(c) AS t FROM test_table GROUP BY a,b ORDER BY a,t DESC;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT a, b, COUNT(c) AS t
FROM test_table
GROUP BY a, b
ORDER BY a, t DESC;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Limit_Clause_0(t *testing.T) {
	sourceSQL := `SELECT * FROM tbl LIMIT 5,10;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM tbl
LIMIT 5, 10;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Limit_Clause_1(t *testing.T) {
	sourceSQL := `SELECT * FROM tbl LIMIT 95,18446744073709551615;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM tbl
LIMIT 95, 18446744073709551615;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Limit_Clause_2(t *testing.T) {
	sourceSQL := `SELECT * FROM tbl LIMIT 95;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM tbl
LIMIT 95;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Limit_Clause_3(t *testing.T) {
	sourceSQL := `SELECT * FROM tbl LIMIT 18446744073709551615 OFFSET 95;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM tbl
LIMIT 18446744073709551615 OFFSET 95;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_ForUpdate_Clause_0(t *testing.T) {
	sourceSQL := `SELECT * FROM t1, t2 FOR SHARE OF t1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM t1, t2
FOR SHARE OF t1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_ForUpdate_Clause_1(t *testing.T) {
	sourceSQL := `SELECT * FROM t1, t2 FOR UPDATE OF t2;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM t1, t2
FOR UPDATE OF t2;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_Into_0(t *testing.T) {
	sourceSQL := `SELECT * INTO @myvar FROM t1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT * INTO @myvar FROM t1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Into_1(t *testing.T) {
	sourceSQL := `SELECT * FROM t1 INTO @myvar FOR UPDATE;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT * FROM t1 INTO @myvar FOR UPDATE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Into_2(t *testing.T) {
	sourceSQL := `SELECT * FROM t1 FOR UPDATE INTO @myvar;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT * FROM t1 FOR UPDATE INTO @myvar;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func TestFormat_Join_Query_0(t *testing.T) {
	sourceSQL := `SELECT t1.name, t2.salary FROM employee AS t1, info AS t2
  WHERE t1.name = t2.name;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT t1.name, t2.salary
FROM employee AS t1, info AS t2
WHERE t1.name = t2.name;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Join_Query_1(t *testing.T) {
	sourceSQL := `SELECT * FROM t1 LEFT JOIN (t2, t3, t4)
                 ON (t2.a = t1.a AND t3.b = t1.b AND t4.c = t1.c)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM t1 LEFT JOIN (t2, t3, t4) ON (t2.a = t1.a AND t3.b = t1.b AND t4.c = t1.c)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Join_Query_2(t *testing.T) {
	sourceSQL := `SELECT * FROM t1 LEFT JOIN (t2 CROSS JOIN t3 CROSS JOIN t4)
                 ON (t2.a = t1.a AND t3.b = t1.b AND t4.c = t1.c)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT *
FROM t1 LEFT JOIN (t2 CROSS JOIN t3 CROSS JOIN t4) ON (t2.a = t1.a AND t3.b = t1.b AND t4.c = t1.c)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Join_Query_3(t *testing.T) {
	sourceSQL := `SELECT t1.name, t2.salary
  FROM employee AS t1 INNER JOIN info AS t2 ON t1.name = t2.name;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT t1.name, t2.salary
FROM employee AS t1 INNER JOIN info AS t2 ON t1.name = t2.name;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_Join_Query_4(t *testing.T) {
	sourceSQL := `SELECT left_tbl.*
  FROM left_tbl LEFT JOIN right_tbl ON left_tbl.id = right_tbl.id
  WHERE right_tbl.id IS NULL;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT left_tbl.*
FROM left_tbl LEFT JOIN right_tbl ON left_tbl.id = right_tbl.id
WHERE right_tbl.id IS NULL;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func TestFormat_UNION_Clause_0(t *testing.T) {
	sourceSQL := `SELECT 1, 2 UNION SELECT 'a', 'b';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1, 2
UNION
SELECT 'a', 'b';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_UNION_Clause_1(t *testing.T) {
	sourceSQL := `SELECT REPEAT('a',1) UNION SELECT REPEAT('b',20);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT REPEAT('a',1) UNION SELECT REPEAT('b',20);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_UNION_Clause_2(t *testing.T) {
	sourceSQL := `(SELECT 1 FOR UPDATE) UNION (SELECT 1 FOR UPDATE);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `(SELECT 1 FOR UPDATE) UNION (SELECT 1 FOR UPDATE);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_UNION_Clause_3(t *testing.T) {
	sourceSQL := `SELECT 1 FOR UPDATE UNION ALL SELECT 1 FOR UPDATE;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1
FOR UPDATE
UNION ALL
SELECT 1
FOR UPDATE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_UNION_Clause_4(t *testing.T) {
	sourceSQL := `SELECT 1 FOR UPDATE UNION DISTINCT SELECT 1 FOR UPDATE;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT 1
FOR UPDATE
UNION DISTINCT
SELECT 1
FOR UPDATE;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}