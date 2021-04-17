package show

import (
	"fmt"
	"github.com/Gumihoy/gumiho-sql-go/sql"
	"github.com/Gumihoy/gumiho-sql-go/sql/db"
	"strings"
	"testing"
)

func TestFormat_SetVariableAssignment_0(t *testing.T) {
	sourceSQL := `SET @name = 43;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET @name = 43;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_SetVariableAssignment_1(t *testing.T) {
	sourceSQL := `SET @total_tax = (SELECT SUM(tax) FROM taxable_transactions);`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET @total_tax = (
	SELECT SUM(tax)
	FROM taxable_transactions
);`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_2(t *testing.T) {
	sourceSQL := `SET GLOBAL max_connections = 1000;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET GLOBAL max_connections = 1000;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_SetVariableAssignment_3(t *testing.T) {
	sourceSQL := `SET @@GLOBAL.max_connections = 1000;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET @@GLOBAL.max_connections = 1000;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_SetVariableAssignment_4(t *testing.T) {
	sourceSQL := `SET SESSION sql_mode = 'TRADITIONAL';
SET LOCAL sql_mode = 'TRADITIONAL';
SET @@SESSION.sql_mode = 'TRADITIONAL';
SET @@LOCAL.sql_mode = 'TRADITIONAL';
SET @@sql_mode = 'TRADITIONAL';
SET sql_mode = 'TRADITIONAL';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET SESSION sql_mode = 'TRADITIONAL';
SET LOCAL sql_mode = 'TRADITIONAL';
SET @@SESSION.sql_mode = 'TRADITIONAL';
SET @@LOCAL.sql_mode = 'TRADITIONAL';
SET @@sql_mode = 'TRADITIONAL';
SET sql_mode = 'TRADITIONAL';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_5(t *testing.T) {
	sourceSQL := `SET PERSIST max_connections = 1000;
SET @@PERSIST.max_connections = 1000;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET PERSIST max_connections = 1000;
SET @@PERSIST.max_connections = 1000;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_6(t *testing.T) {
	sourceSQL := `SET PERSIST_ONLY back_log = 100;
SET @@PERSIST_ONLY.back_log = 100;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET PERSIST_ONLY back_log = 100;
SET @@PERSIST_ONLY.back_log = 100;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_7(t *testing.T) {
	sourceSQL := `SET @@SESSION.max_join_size = DEFAULT;
SET @@SESSION.max_join_size = @@GLOBAL.max_join_size;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET @@SESSION.max_join_size = DEFAULT;
SET @@SESSION.max_join_size = @@GLOBAL.max_join_size;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_8(t *testing.T) {
	sourceSQL := `SET @x = 1, SESSION sql_mode = '';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET @x = 1, SESSION sql_mode = '';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_9(t *testing.T) {
	sourceSQL := `SET GLOBAL sort_buffer_size = 1000000, SESSION sort_buffer_size = 1000000;
SET @@GLOBAL.sort_buffer_size = 1000000, @@LOCAL.sort_buffer_size = 1000000;
SET GLOBAL max_connections = 1000, sort_buffer_size = 1000000;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET GLOBAL sort_buffer_size = 1000000, SESSION sort_buffer_size = 1000000;
SET @@GLOBAL.sort_buffer_size = 1000000, @@LOCAL.sort_buffer_size = 1000000;
SET GLOBAL max_connections = 1000, sort_buffer_size = 1000000;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_10(t *testing.T) {
	sourceSQL := `SET @@GLOBAL.sort_buffer_size = 50000, sort_buffer_size = 1000000;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET @@GLOBAL.sort_buffer_size = 50000, sort_buffer_size = 1000000;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetVariableAssignment_11(t *testing.T) {
	sourceSQL := `SELECT @@GLOBAL.sql_mode, @@SESSION.sql_mode, @@sql_mode;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SELECT @@GLOBAL.sql_mode, @@SESSION.sql_mode, @@sql_mode;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}






func TestFormat_SetCharacterSet_0(t *testing.T) {
	sourceSQL := `SET CHARACTER SET 'charset_name';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET CHARACTER SET 'charset_name';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_SetCharacterSet_1(t *testing.T) {
	sourceSQL := `SET CHARACTER SET DEFAULT;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET CHARACTER SET DEFAULT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}

func TestFormat_SetCharset_0(t *testing.T) {
	sourceSQL := `SET CHARSET 'charset_name';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET CHARSET 'charset_name';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_SetCharset_1(t *testing.T) {
	sourceSQL := `SET CHARSET DEFAULT;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET CHARSET DEFAULT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}


func TestFormat_SetNames_0(t *testing.T) {
	sourceSQL := `SET NAMES 'charset_name';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET NAMES 'charset_name';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_SetNames_1(t *testing.T) {
	sourceSQL := `SET NAMES 'charset_name' COLLATE 'collation_name';`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET NAMES 'charset_name' COLLATE 'collation_name';`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}
func TestFormat_SetNames_2(t *testing.T) {
	sourceSQL := `SET NAMES DEFAULT;`
	formatSQL := sql.Format(sourceSQL, db.MySQL)
	fmt.Println(sourceSQL)
	fmt.Println("----------------------")
	fmt.Println(formatSQL)

	targetSQL := `SET NAMES DEFAULT;`

	if !strings.EqualFold(formatSQL, targetSQL) {
		t.Error()
	}
}