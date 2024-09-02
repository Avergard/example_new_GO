package handlers

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres dbname=postgres sslmode=disable password=Yanelox46"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
