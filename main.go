package main

import (
	"net/http"

	"test-routes/api"

	"github.com/labstack/echo/v4"
)

var srv http.Handler

func main() {

	e := echo.New()
	e.GET("/", api.HealthCheck)
	e.GET("/books", api.Book)
	srv = e

	e.Start(":8080")
}

func Ф(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
