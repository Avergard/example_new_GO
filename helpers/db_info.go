package helpers

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	authPost := "user=postgres dbname=postgres sslmode=disable password=yourPass"
	db, err = sql.Open("postgres", authPost)
	if err != nil {
		log.Fatalf("Ошибка открытия базы данных: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
}

/*вынести данные authpost В константу как у дена
Postgres поменять на postgres

*/