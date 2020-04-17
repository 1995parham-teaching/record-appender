package db

import (
	"database/sql"
	"log"

	"github.com/elahe-dastan/record-appender/config"
	_ "github.com/lib/pq" //adding dialect for postgres
)

func New(database config.Database) *sql.DB {
	db, err := sql.Open("postgres", database.Cstring())
	if err != nil {
		log.Printf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}
