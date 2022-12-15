package handler

import (
	"net/http"
	"strconv"

	"github.com/1995parham-teaching/record-appender/store"
	"github.com/labstack/echo/v4"
)

const BulkSize = 100

type API struct {
	Store store.Data
}

func (a API) Run() {
	e := echo.New()
	e.GET("/api/:n", a.retrieve)

	e.Logger.Fatal(e.Start(":8080"))
}

func (a API) retrieve(c echo.Context) error {
	number := c.Param("n")

	num, err := strconv.Atoi(number)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := a.Store.Retrieve(BulkSize, num*BulkSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
