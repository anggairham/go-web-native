# learning golang crud
* https://www.youtube.com/watch?v=YYjb_RszQYY
* https://github.com/tegarpratama/go-crud-web

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

### Scan error on column index 2, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Tim
* Solusi
```bash
# tambahkan ?parseTime=true
	db, err := sql.Open("mysql", "angga:password@/go_web_native?parseTime=true")
```

## Live-reload
* instalasi air live-reload
```bash
go install github.com/air-verse/air@latest
```
* langsung ketik perintah air pada directory root project untuk live-reload
```bash
air
```