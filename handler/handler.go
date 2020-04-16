package handler

import (
	"database/sql"
	"net/http"
	"snapfood/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type API struct {
	DB *sql.DB
}

func (a API) Run() {
	e := echo.New()
	e.GET("/api/:n", a.retrieve)

	e.Logger.Fatal(e.Start(":8080"))
}

func (a API) retrieve(c echo.Context) error {
	number := c.Param("n")

	num,_ := strconv.Atoi(number)
	rows, _ := a.DB.Query("SELECT * from data OFFSET $1 LIMIT 100;", num*100)

	result := make([]model.Data, 0)

	defer rows.Close()
	for  rows.Next() {
		var data model.Data
		rows.Scan(&data.FirstName, &data.LastName, &data.Number)
		result = append(result, data)
	}

	return c.JSON(http.StatusOK, result)
}