# README

The third assignment of the course IDBS-Spring20-Fudan

To run the test, you must modify the following code in *library.go*

```go
func (lib *Library) ConnectDBLocal() {
	lib.User = "Your username"
	lib.Password = "Your password"
	lib.DBName = "Your database name"
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@Your location such as tcp(127.0.0.1:3306)/%s", lib.User, lib.Password, lib.DBName))
	if err != nil {
		panic(err)
	}
	lib.db = db
	//	defer db.Close()
}
```

Then run the following instruction

```
go test -v
```

