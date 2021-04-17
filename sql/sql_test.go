package sql

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql/config"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"testing"
)

func Test_Format_Oracle_0(t *testing.T) {
	sql := `SELECT object_type, count(*) from dba_objects
	where (object_type not like 'LOB%' and object_type not like 'TABLE%PARTITION' and object_type not like 'INDEX%PARTITION')
	and (owner, object_name) not in(select owner,table_name from dba_all_tables where nested='YES')
		and owner in (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		AND (OWNER, OBJECT_TYPE, OBJECT_NAME) NOT IN (SELECT OWNER, 'INDEX' AS OBJECT_TYPE, CONSTRAINT_NAME FROM DBA_CONSTRAINTS WHERE OWNER in (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		AND CONSTRAINT_TYPE IN('U','P'))
		AND OBJECT_NAME NOT LIKE 'SYS_IL%'
        and object_type not in('DATABASE LINK')
		AND OBJECT_NAME NOT LIKE 'MLOG$_%' 
		AND OBJECT_NAME NOT LIKE 'RUPD$_%'
        GROUP BY object_type`

	formatSQL := Format(sql, db.Oracle)
	fmt.Println(formatSQL)
}


func TestFormat_MySQL_0(t *testing.T) {
	sql := `insert
  into print_test
SELECT
  cate_id,
  seller_id,
  stat_date,
  pay_ord_amt  --不输出rownum字段，能减小结果表的输出量。
FROM (
    SELECT
      *,
      ROW_NUMBER () as rownum  --根据上游sum结果排序。
    FROM (
        SELECT
          cate_id,
          seller_id,
          stat_date,
          --重点。声明Sum的参数都是正数，所以Sum的结果是单调递增的，因此TopN能使用优化算法，只获取前100个数据。
          sum (total_fee) as pay_ord_amt
        FROM
          random_test
        WHERE
          total_fee >= 0
        GROUP
          BY cate_name,
          seller_id,
          stat_date
      ) a
    WHERE
      rownum <= 100
  );             `
	config := config.NewFormatConfig()
	config.SkipComment = true
	formatSQL := FormatWithConfig(sql, db.MySQL, config)
	fmt.Println(formatSQL)
}
