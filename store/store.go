package store

import (
	"database/sql"
	"fmt"
	"snapfood/model"
	"strings"
)

func Insert(db *sql.DB, data []model.Data) error {
	values := ""
	for _,d := range data {
		values += "('"
		values += d.FirstName
		values += "','"
		values += d.LastName
		values += "','"
		values += d.Number
		values += "'),"
	}

	values = strings.TrimSuffix(values, ",")
	fmt.Println("INSERT INTO data (first, last, number) VALUES " + values)

	_, err := db.Exec("INSERT INTO data (first, last, number) VALUES " + values + ";")

	return err
}
