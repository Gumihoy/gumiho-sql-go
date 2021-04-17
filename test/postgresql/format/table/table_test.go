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
	sourceSQL := `ALTER TABLE distributors ADD COLUMN address varchar(30);`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors
	ADD COLUMN address varchar(30);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_AlterTable_AddColumn_1(t *testing.T) {
	sourceSQL := `ALTER TABLE measurements
  ADD COLUMN mtime timestamp with time zone DEFAULT now();`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

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
func TestFormat_AlterTable_AddColumn_2(t *testing.T) {
	sourceSQL := `ALTER TABLE transactions
  ADD COLUMN status varchar(30) DEFAULT 'old',
  ALTER COLUMN status SET default 'current';`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE transactions
  ADD COLUMN status varchar(30) DEFAULT 'old',
  ALTER COLUMN status SET default 'current';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_AlterTable_DropColumn_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors DROP COLUMN address RESTRICT;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors DROP COLUMN address RESTRICT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_AlterColumn_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors
    ALTER COLUMN address TYPE varchar(80),
    ALTER COLUMN name TYPE varchar(100);`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors
    ALTER COLUMN address TYPE varchar(80),
    ALTER COLUMN name TYPE varchar(100);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_AlterColumn_1(t *testing.T) {
	sourceSQL := `ALTER TABLE foo
    ALTER COLUMN foo_timestamp SET DATA TYPE timestamp with time zone
    USING
        timestamp with time zone 'epoch' + foo_timestamp * interval '1 second';`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE foo
    ALTER COLUMN foo_timestamp SET DATA TYPE timestamp with time zone
    USING
        timestamp with time zone 'epoch' + foo_timestamp * interval '1 second';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_AlterColumn_2(t *testing.T) {
	sourceSQL := `ALTER TABLE foo
    ALTER COLUMN foo_timestamp DROP DEFAULT,
    ALTER COLUMN foo_timestamp TYPE timestamp with time zone
    USING
        timestamp with time zone 'epoch' + foo_timestamp * interval '1 second',
    ALTER COLUMN foo_timestamp SET DEFAULT now();`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE foo
    ALTER COLUMN foo_timestamp DROP DEFAULT,
    ALTER COLUMN foo_timestamp TYPE timestamp with time zone
    USING
        timestamp with time zone 'epoch' + foo_timestamp * interval '1 second',
    ALTER COLUMN foo_timestamp SET DEFAULT now();`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_RenameColumn_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors RENAME COLUMN address TO city;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors RENAME COLUMN address TO city;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_RenameTable_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors RENAME TO suppliers;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors RENAME TO suppliers;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_RenameConstraint_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors RENAME CONSTRAINT zipchk TO zip_check;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors RENAME CONSTRAINT zipchk TO zip_check;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_ColumnSetNotNull_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors ALTER COLUMN street SET NOT NULL;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors ALTER COLUMN street SET NOT NULL;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_ColumnDropNotNull_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors ALTER COLUMN street DROP NOT NULL;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors ALTER COLUMN street DROP NOT NULL;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_AddCheckTableConstraint_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5);ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5) NO INHERIT;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5);ALTER TABLE distributors ADD CONSTRAINT zipchk CHECK (char_length(zipcode) = 5) NO INHERIT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_AlterTable_DropTableConstraint_0(t *testing.T) {
	sourceSQL := `ALTER TABLE distributors DROP CONSTRAINT zipchk;ALTER TABLE ONLY distributors DROP CONSTRAINT zipchk;`
	formatSQL := sql.Format(sourceSQL, db.PostgreSQL)

	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `ALTER TABLE distributors
	DROP CONSTRAINT zipchk;
ALTER TABLE ONLY distributors
	DROP CONSTRAINT zipchk;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
// ------------------------------ Create --------------------------------------------------------------------------

// ------------------------------ Drop --------------------------------------------------------------------------