package tests

import (
	"auth-service/config"
	"auth-service/users/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.POST("/login", handlers.Login)
	t.Run("Success Login : 200 ", func(t *testing.T) {
		roleCredentials := cfg.User.SuperAdmin
		m := models.LoginDto{
			Username: roleCredentials.Username,
			Password: roleCredentials.Password,
		}

		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// retrive respond body
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		// check token exists
		token := responseBody["token"]
		assert.NotEmpty(t, token)

		// check role match
		role := responseBody["role"]
		assert.Equal(t, roleCredentials.Role, role)
	})

	t.Run("Failed Login with invalid credentials : 401", func(t *testing.T) {
		roleCredentials := cfg.User.Admins
		m := models.LoginDto{
			Username: roleCredentials.Username,
			Password: "wrong-password",
		}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Failed Login with missing credentials : 400", func(t *testing.T) {
		m := models.LoginDto{}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
