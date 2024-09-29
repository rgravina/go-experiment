package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-experiment/web"
	"net/http"
)

type User struct {
	Name string `json:"name" xml:"name"`
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	web.RegisterHandlers(e)
	u := &User{
		Name: "Rocky Racoon",
	}
	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
