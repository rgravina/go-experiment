package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name string `json:"name" xml:"name"`
}

var u = &User{Name: "Rocky Racoon"}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, *u)
}
