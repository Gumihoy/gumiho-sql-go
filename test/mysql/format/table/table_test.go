package table

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

// ------------------------------ Alter --------------------------------------------------------------------------
func TestFormat_AlterTable_AddColumn_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 ADD COLUMN c2 INT, ADD COLUMN c3 INT, ALGORITHM=INSTANT;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1
	ADD COLUMN c2 INT,
	ADD COLUMN c3 INT,
	ALGORITHM = INSTANT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_AddIndex_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 ADD index index_name `
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 ADD COLUMN c2 INT, ADD COLUMN c3 INT, ALGORITHM = INSTANT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_AddKey_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 ADD COLUMN c2 INT, ADD COLUMN c3 INT, ALGORITHM=INSTANT;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 ADD COLUMN c2 INT, ADD COLUMN c3 INT, ALGORITHM = INSTANT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}




func TestFormat_AlterTable_DropColumn_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t2 DROP COLUMN c, DROP COLUMN d;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t2 DROP COLUMN c, DROP COLUMN d;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_DropIndex_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 DROP INDEX index1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 DROP INDEX index1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_DropKey_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 DROP KEY index1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 DROP KEY index1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_DropConstraint_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 DROP CONSTRAINT constraint1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 DROP CONSTRAINT constraint1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_DropPrimaryKey_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 DROP PRIMARY KEY;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 DROP PRIMARY KEY;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_DropForeignKey_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 DROP FOREIGN KEY foreignName;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 DROP FOREIGN KEY foreignName;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_DropCheck_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 DROP CHECK check1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 DROP CHECK check1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_Partition_0(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 ADD PARTITION (PARTITION p3 VALUES LESS THAN (2002));`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 ADD PARTITION (PARTITION p3 VALUES LESS THAN (2002));`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_Partition_7(t *testing.T) {
	sourceSQL := `ALTER TABLE t1 ANALYZE PARTITION p1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE t1 ANALYZE PARTITION p1`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}



// ------------------------------ Create --------------------------------------------------------------------------

func TestFormat_CreateTable_DataType_0(t *testing.T) {
	sourceSQL := `CREATE TABLE t (c CHAR(20) CHARACTER SET utf8 COLLATE utf8_bin);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE t (c CHAR(20) CHARACTER SET utf8 COLLATE utf8_bin);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_Simple_0(t *testing.T) {
	sourceSQL := `CREATE TABLE test (
					col1 NUMBER(5,2),
					col2 FLOAT(5)
				)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE test (
	col1 NUMBER(5, 2),
	col2 FLOAT(5)
)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TEMPORARY_0(t *testing.T) {
	sourceSQL := `CREATE TEMPORARY TABLE test (
					col1 NUMBER(5,2),
					col2 FLOAT(5)
				)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TEMPORARY TABLE test (
	col1 NUMBER(5, 2),
	col2 FLOAT(5)
)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_IfNotExists_0(t *testing.T) {
	sourceSQL := `CREATE TABLE IF NOT EXISTS test (
					col1 NUMBER(5,2),
					col2 FLOAT(5)
				)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE IF NOT EXISTS test (
	col1 NUMBER(5, 2),
	col2 FLOAT(5)
)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TableColumn_0(t *testing.T) {
	sourceSQL := `CREATE TABLE IF NOT EXISTS test (
					col1 NUMBER(5,2),
					col2 FLOAT(5)
				)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE IF NOT EXISTS test (
	col1 NUMBER(5, 2),
	col2 FLOAT(5)
)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TableConstraint_0(t *testing.T) {
	sourceSQL := `CREATE TABLE IF NOT EXISTS test (
					col1 NUMBER(5,2),
					col2 FLOAT(5)
				)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE IF NOT EXISTS test (
	col1 NUMBER(5, 2),
	col2 FLOAT(5)
)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_CreateTable_TableConstraint_1(t *testing.T) {
	sourceSQL := `CREATE TABLE test (blob_col BLOB, INDEX(blob_col(10)));`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE test (blob_col BLOB, INDEX(blob_col(10)));`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_ColumnConstraint_0(t *testing.T) {
	sourceSQL := `CREATE TABLE IF NOT EXISTS test (
					col1 NUMBER(5,2) NOT NULL DEFAULT '0' VISIBLE AUTO_INCREMENT PRIMARY KEY,
					col2 NUMBER(5,2) NULL DEFAULT '1' UNIQUE,
					col2 FLOAT(5)
				)`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE IF NOT EXISTS test (
	col1 NUMBER(5, 2),
	col2 FLOAT(5)
)`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_LikeClause_0(t *testing.T) {
	sourceSQL := `CREATE TABLE new_tbl LIKE orig_tbl;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE new_tbl LIKE orig_tbl;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_CreateTable_LikeClause_1(t *testing.T) {
	sourceSQL := `CREATE TABLE new_tbl (LIKE orig_tbl);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE new_tbl (
	LIKE orig_tbl
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_SubQuery_0(t *testing.T) {
	sourceSQL := `CREATE TABLE new_tbl AS SELECT * FROM orig_tbl;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE new_tbl AS SELECT * FROM orig_tbl;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_0(t *testing.T) {
	sourceSQL := `CREATE TABLE t1 (col1 INT, col2 CHAR(5))
    PARTITION BY HASH(col1) PARTITIONS 1 SUBPARTITION BY HASH(col1) SUBPARTITIONS 1;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE t1 (
	col1 INT,
	col2 CHAR(5)
)
PARTITION BY HASH (col1)
PARTITIONS 1
SUBPARTITION BY HASH (
	col1
 )
SUBPARTITIONS 1;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_1(t *testing.T) {
	sourceSQL := `CREATE TABLE tk (col1 INT, col2 CHAR(5), col3 DATE)
    PARTITION BY KEY(col3)
    PARTITIONS 4;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE tk (
	col1 INT,
	col2 CHAR(5),
	col3 DATE
)
PARTITION BY KEY (col3)
PARTITIONS 4;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_2(t *testing.T) {
	sourceSQL := `CREATE TABLE tk (col1 INT, col2 CHAR(5), col3 DATE)
    PARTITION BY LINEAR KEY(col3)
    PARTITIONS 5;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE tk (
	col1 INT,
	col2 CHAR(5),
	col3 DATE
)
PARTITION BY KEY (col3)
PARTITIONS 5;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_3(t *testing.T) {
	sourceSQL := `CREATE TABLE t1 (
    year_col  INT,
    some_data INT
)
PARTITION BY RANGE (year_col) (
    PARTITION p0 VALUES LESS THAN (1991),
    PARTITION p1 VALUES LESS THAN (1995),
    PARTITION p2 VALUES LESS THAN (1999),
    PARTITION p3 VALUES LESS THAN (2002),
    PARTITION p4 VALUES LESS THAN (2006),
    PARTITION p5 VALUES LESS THAN MAXVALUE
);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE t1 (
	year_col INT,
	some_data INT
)
PARTITION BY RANGE (year_col) (
	PARTITION p0 VALUES LESS THAN (1991),
	PARTITION p1 VALUES LESS THAN (1995),
	PARTITION p2 VALUES LESS THAN (1999),
	PARTITION p3 VALUES LESS THAN (2002),
	PARTITION p4 VALUES LESS THAN (2006),
	PARTITION p5 VALUES LESS THAN MAXVALUE
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_4(t *testing.T) {
	sourceSQL := `CREATE TABLE rc (
    a INT NOT NULL,
    b INT NOT NULL
)
PARTITION BY RANGE COLUMNS(a,b) (
    PARTITION p0 VALUES LESS THAN (10,5),
    PARTITION p1 VALUES LESS THAN (20,10),
    PARTITION p2 VALUES LESS THAN (50,MAXVALUE),
    PARTITION p3 VALUES LESS THAN (65,MAXVALUE),
    PARTITION p4 VALUES LESS THAN (MAXVALUE,MAXVALUE)
);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE rc (
	a INT NOT NULL,
	b INT NOT NULL
)
PARTITION BY RANGE (a, b) (
	PARTITION p0 VALUES LESS THAN (10, 5),
	PARTITION p1 VALUES LESS THAN (20, 10),
	PARTITION p2 VALUES LESS THAN (50, MAXVALUE),
	PARTITION p3 VALUES LESS THAN (65, MAXVALUE),
	PARTITION p4 VALUES LESS THAN (MAXVALUE, MAXVALUE)
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_5(t *testing.T) {
	sourceSQL := `CREATE TABLE client_firms (
    id   INT,
    name VARCHAR(35)
)
PARTITION BY LIST (id) (
    PARTITION r0 VALUES IN (1, 5, 9, 13, 17, 21),
    PARTITION r1 VALUES IN (2, 6, 10, 14, 18, 22),
    PARTITION r2 VALUES IN (3, 7, 11, 15, 19, 23),
    PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE client_firms (
    id   INT,
    name VARCHAR(35)
)
PARTITION BY LIST (id) (
    PARTITION r0 VALUES IN (1, 5, 9, 13, 17, 21),
    PARTITION r1 VALUES IN (2, 6, 10, 14, 18, 22),
    PARTITION r2 VALUES IN (3, 7, 11, 15, 19, 23),
    PARTITION r3 VALUES IN (4, 8, 12, 16, 20, 24)
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_6(t *testing.T) {
	sourceSQL := `CREATE TABLE lc (
    a INT NULL,
    b INT NULL
)
PARTITION BY LIST COLUMNS(a,b) (
    PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
    PARTITION p1 VALUES IN( (0,1), (0,2), (0,3), (1,1), (1,2) ),
    PARTITION p2 VALUES IN( (1,0), (2,0), (2,1), (3,0), (3,1) ),
    PARTITION p3 VALUES IN( (1,3), (2,2), (2,3), (3,2), (3,3) )
);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE lc (
    a INT NULL,
    b INT NULL
)
PARTITION BY LIST COLUMNS(a,b) (
    PARTITION p0 VALUES IN( (0,0), (NULL,NULL) ),
    PARTITION p1 VALUES IN( (0,1), (0,2), (0,3), (1,1), (1,2) ),
    PARTITION p2 VALUES IN( (1,0), (2,0), (2,1), (3,0), (3,1) ),
    PARTITION p3 VALUES IN( (1,3), (2,2), (2,3), (3,2), (3,3) )
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_TablePartitioning_7(t *testing.T) {
	sourceSQL := `CREATE TABLE th (id INT, name VARCHAR(30), adate DATE)
PARTITION BY LIST(YEAR(adate))
(
  PARTITION p1999 VALUES IN (1995, 1999, 2003)
    DATA DIRECTORY = '/var/appdata/95/data'
    INDEX DIRECTORY = '/var/appdata/95/idx',
  PARTITION p2000 VALUES IN (1996, 2000, 2004)
    DATA DIRECTORY = '/var/appdata/96/data'
    INDEX DIRECTORY = '/var/appdata/96/idx',
  PARTITION p2001 VALUES IN (1997, 2001, 2005)
    DATA DIRECTORY = '/var/appdata/97/data'
    INDEX DIRECTORY = '/var/appdata/97/idx',
  PARTITION p2002 VALUES IN (1998, 2002, 2006)
    DATA DIRECTORY = '/var/appdata/98/data'
    INDEX DIRECTORY = '/var/appdata/98/idx'
);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE th (id INT, name VARCHAR(30), adate DATE)
PARTITION BY LIST(YEAR(adate))
(
  PARTITION p1999 VALUES IN (1995, 1999, 2003)
    DATA DIRECTORY = '/var/appdata/95/data'
    INDEX DIRECTORY = '/var/appdata/95/idx',
  PARTITION p2000 VALUES IN (1996, 2000, 2004)
    DATA DIRECTORY = '/var/appdata/96/data'
    INDEX DIRECTORY = '/var/appdata/96/idx',
  PARTITION p2001 VALUES IN (1997, 2001, 2005)
    DATA DIRECTORY = '/var/appdata/97/data'
    INDEX DIRECTORY = '/var/appdata/97/idx',
  PARTITION p2002 VALUES IN (1998, 2002, 2006)
    DATA DIRECTORY = '/var/appdata/98/data'
    INDEX DIRECTORY = '/var/appdata/98/idx'
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_CreateTable_Simple_3(t *testing.T) {
	sourceSQL := "CREATE TABLE `objects`(" +
		"	`id` bigint(20) NOT NULL AUTO_INCREMENT," +
		" 	`dbid` bigint(20) DEFAULT NULL," +
		" 	`file_md5` varchar(32) DEFAULT NULL," +
		" 	`owner` varchar(32) DEFAULT NULL," +
		" 	`object_type` varchar(32) DEFAULT NULL," +
		" 	`object_name` varchar(128) DEFAULT NULL," +
		" 	`object_ddl` longtext," +
		" 	`islet` int(11) DEFAULT NULL," +
		" 	`tables` text," +
		" 	`volume` bigint(20) DEFAULT '0'," +
		" 	`lob` bigint(20) DEFAULT '0'," +
		" 	`num_rows` bigint(20) DEFAULT '0'," +
		" 	`tags` text," +
		" 	`partitioned` varchar(3) DEFAULT NULL," +
		" 	`where_clause` text," +
		" 	`tablespace_name` varchar(32) DEFAULT NULL," +
		" 	`param1` varchar(1000) DEFAULT NULL," +
		" 	`param2` varchar(1000) DEFAULT NULL," +
		" 	`param3` varchar(1000) DEFAULT NULL," +
		" 	`param4` varchar(1000) DEFAULT NULL," +
		" 	`md5` char(32) DEFAULT NULL," +
		" 	PRIMARY KEY (`id`)," +
		" 	KEY `ind_dbid` (`dbid`)," +
		" 	KEY `objects_owner_index` (`owner`)," +
		" 	KEY `objects_object_name_index` (`object_name`)," +
		" 	KEY `md5_idx` (`file_md5`)" +
		" ) ENGINE=InnoDB AUTO_INCREMENT=42244440 DEFAULT CHARSET=utf8;"
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `CREATE TABLE t1 (
	year_col INT,
	some_data INT
)
PARTITION BY RANGE (year_col) (
	PARTITION p0 VALUES LESS THAN (1991),
	PARTITION p1 VALUES LESS THAN (1995),
	PARTITION p2 VALUES LESS THAN (1999),
	PARTITION p3 VALUES LESS THAN (2002),
	PARTITION p4 VALUES LESS THAN (2006),
	PARTITION p5 VALUES LESS THAN MAXVALUE
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

// ------------------------------ Drop --------------------------------------------------------------------------

func TestFormat_DropTable_0(t *testing.T) {
	sourceSQL := `DROP TABLE test1`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TABLE test1`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_DropTable_1(t *testing.T) {
	sourceSQL := `DROP TABLE IF EXISTS test1`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TABLE IF EXISTS test1`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_DropTable_2(t *testing.T) {
	sourceSQL := `DROP TABLE IF EXISTS test1, test2`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TABLE IF EXISTS test1, test2`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_DropTable_3(t *testing.T) {
	sourceSQL := `DROP TEMPORARY TABLE IF EXISTS test1`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TEMPORARY TABLE IF EXISTS test1`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_DropTable_4(t *testing.T) {
	sourceSQL := `DROP TEMPORARY TABLE IF EXISTS test1, test2 RESTRICT`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TEMPORARY TABLE IF EXISTS test1, test2 RESTRICT`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_DropTable_5(t *testing.T) {
	sourceSQL := `DROP TEMPORARY TABLE IF EXISTS test1, test2 CASCADE`
	formatSQL := sql.Format(sourceSQL, db.MySQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `DROP TEMPORARY TABLE IF EXISTS test1, test2 CASCADE`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
