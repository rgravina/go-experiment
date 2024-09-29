package main

import (
	"github.com/labstack/echo/v4"
	"go-experiment/web"
	"net/http"
)

func main() {
	e := echo.New()
	web.RegisterHandlers(e)
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
