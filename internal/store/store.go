package store

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/1995parham-teaching/record-appender/internal/model"
)

type Data interface {
	Insert(data []model.Data) error
	Retrieve(limit int, offset int) ([]model.Data, error)
}

type SQLData struct {
	DB *sql.DB
}

func (s SQLData) Retrieve(limit int, offset int) ([]model.Data, error) {
	rows, err := s.DB.Query("SELECT * from data OFFSET $1 LIMIT $2;", offset, limit)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	result := make([]model.Data, 0)

	for rows.Next() {
		var data model.Data
		if err := rows.Scan(&data.FirstName, &data.LastName, &data.Number); err != nil {
			return nil, err
		}

		result = append(result, data)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s SQLData) Insert(data []model.Data) error {
	query := "INSERT INTO data (first, last, number) VALUES "
	parameters := make([]interface{}, 0)

	// create a bulk query for inserting all data.
	for i, d := range data {
		// nolint: gomnd
		query += fmt.Sprintf("($%d, $%d, $%d),", 3*i+1, 3*i+2, 3*i+3)

		parameters = append(parameters, d.FirstName, d.LastName, d.Number)
	}

	query = strings.TrimSuffix(query, ",")

	_, err := s.DB.Exec(query, parameters...)

	return err
}
