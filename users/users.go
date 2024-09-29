package users

import (
	"github.com/labstack/echo/v4"
	"go-experiment/db"
	"go-experiment/model"
	"net/http"
)

type User struct {
	Name string `json:"name" xml:"name"`
}

func GetUsers(c echo.Context) error {
	database := db.DbManager()
	var users []model.User
	database.Find(&users)
	return c.JSON(http.StatusOK, users)
}
