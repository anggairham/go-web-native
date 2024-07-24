# 

* jalankan
```bash
go run .
go get -u github.com/go-sql-driver/mysql
```

## Trouble
### Windows showing security when running go run .
* dikarenkan penulisan 
```go
http.ListenAndServe(":8080",nil)
// ubah menjadi
http.ListenAndServe("127.0.0.1:8080",nil)
```