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

func TestResetPassword(t *testing.T) {

	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.POST("/reset-password", handlers.ResetPassword)
	e.POST("/login", handlers.Login)
	roleCredentials := cfg.User.SuperAdmin

	token, err := getResetPasswordToken(cfg, roleCredentials.Email)
	if err != nil {
		panic(err)
	}

	t.Run("Success Reset Password ; 201", func(t *testing.T) {
		newPassword := "newPassword12345"
		m := models.ResetPassword{
			Token:    token,
			Password: newPassword,
		}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		// Test login with new Password

		mL := models.LoginDto{
			Username: roleCredentials.Username,
			Password: newPassword,
		}

		payloadmL, _ := json.Marshal(mL)
		reqmL, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(payloadmL))
		reqmL.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		wmL := httptest.NewRecorder()
		e.ServeHTTP(wmL, reqmL)

		assert.Equal(t, http.StatusOK, wmL.Code)

		// Change password back
		newPassword = roleCredentials.Password
		m = models.ResetPassword{
			Token:    token,
			Password: newPassword,
		}
		payload, _ = json.Marshal(m)
		req, _ = http.NewRequest("POST", "/reset-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w = httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Failed Reset Password; Same password ; 422", func(t *testing.T) {
		m := models.ResetPassword{
			Token:    token,
			Password: roleCredentials.Password,
		}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	})

	t.Run("Failed Reset Password; Invalid token ; 401", func(t *testing.T) {
		m := models.ResetPassword{
			Token:    "SOME WRONG TOKEN",
			Password: roleCredentials.Password,
		}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)

	})

	t.Run("Failed Reset Password; Invalid password format ; 400", func(t *testing.T) {
		m := models.ResetPassword{
			Token:    token,
			Password: "short",
		}
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/reset-password", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
