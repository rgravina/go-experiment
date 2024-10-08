package users

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-experiment/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	db.Init()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/users")
	users := GetUsers(c)
	if assert.NoError(t, users) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(
			t, rec.Body.String(), `"name":"Rocky Racoon"`,
		)
	}
}
