package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// variabel global
var DB *sql.DB

func ConnectDb() {
	// connect mysql db
	db, err := sql.Open("mysql", "angga:password@/go_web_native")
	if err != nil {
		panic(err)
	}
	log.Println("Database connected")
	// isi variabel global
	DB = db
}
