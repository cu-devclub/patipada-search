package tests

import (
	"auth-service/config"
	"auth-service/jwt"
	"auth-service/users/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestChangePassword(t *testing.T) {
	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.POST("/change-password", handlers.ChangePassword)
	t.Run("Change Password Success", func(t *testing.T) {
		roleCredentials := cfg.User.SuperAdmin
		m := models.ChangePassword{
			OldPassword: roleCredentials.Password,
			NewPassword: "test-new-password",
		}

		token, err := jwt.CreateToken(roleCredentials.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/change-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		req.Header.Set("Authorization", token)             // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		//Change password back 
		m = models.ChangePassword{
			OldPassword: m.NewPassword,
			NewPassword: roleCredentials.Password,
		}
		payload, _ = json.Marshal(m)
		req, _ = http.NewRequest("POST", "/change-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		req.Header.Set("Authorization", token)             // Set the Content-Type header
		w = httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Change Password Fail - Invalid Old Password", func(t *testing.T) {
		roleCredentials := cfg.User.SuperAdmin
		m := models.ChangePassword{
			OldPassword: "test-invalid-password",
			NewPassword: "test-new-password",
		}

		token, err := jwt.CreateToken(roleCredentials.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/change-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		req.Header.Set("Authorization", token)             // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
