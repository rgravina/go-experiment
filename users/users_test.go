package users

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api")
	users := GetUsers(c)

	if assert.NoError(t, users) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(
			t,
			`{"name":"Rocky Racoon"}
`, rec.Body.String())
	}
}
