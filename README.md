# SQL Go


## Format

``` go
func TestFormat(t *testing.T) {
	sql := "select *, idaaaaaaaaaeeeeeee from dual where id = ?;"
	formatSQL := format.Format(sql, db.MySQL)
	fmt.Println(formatSQL)
}
```