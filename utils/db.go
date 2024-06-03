package utils

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=postgres dbname=mciu_db sslmode=disable password=6222"
	var err error
	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

func CloseDB() {
	DB.Close()
}
