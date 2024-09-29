package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-experiment/users"
	"go-experiment/web"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	web.RegisterHandlers(e)
	e.GET("/api/users", users.GetUsers)
	e.Logger.Fatal(e.Start(":8080"))
}
