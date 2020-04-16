package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" //adding dialect for postgres
)

func New() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=snapfood password=postgres sslmode=disable")
	if err != nil {
		log.Printf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}
