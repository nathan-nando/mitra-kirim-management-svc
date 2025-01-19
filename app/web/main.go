package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	log.Printf("THIS IS MGMT SVC")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hi MGMT")
	})

	if err := e.Start(":9000"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
