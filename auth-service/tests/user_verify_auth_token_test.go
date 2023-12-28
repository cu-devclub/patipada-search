package tests

import (
	"auth-service/config"
	"auth-service/jwt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserVerifyAuthToken(t *testing.T) {
	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.GET("/verify-token", handlers.VerifyToken)
	t.Run("Success Verify Token : 200 ", func(t *testing.T) {
		token, err := jwt.CreateToken(cfg.User.SuperAdmin.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			panic(err)
		}
		req, _ := http.NewRequest("GET", "/verify-token", nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Fail Verify Token : 400 Missing token", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/verify-token", nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Fail Verify Token : 401 Invalid token", func(t *testing.T) {
		token, err := jwt.CreateToken(cfg.User.SuperAdmin.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			panic(err)
		}
		req, _ := http.NewRequest("GET", "/verify-token", nil)
		req.Header.Set("Authorization", token+"Malformed")
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
