package handlers

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	authPost := "user=postgres dbname=postgres sslmode=disable password=+++++++"
	db, err = sql.Open("postgres", authPost)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
