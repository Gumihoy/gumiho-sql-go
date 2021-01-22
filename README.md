# SQL Go

## String SQL to Abstract Syntax Tree

## String SQL Format

``` go
func TestFormat(t *testing.T) {
	sql := "select *, idaaaaaaaaaeeeeeee from dual where id = ?;"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}
```
