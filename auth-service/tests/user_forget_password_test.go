package tests

import (
	"auth-service/config"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestForgetPassword(t *testing.T) {
	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.POST("/forget-password/:email", handlers.ForgetPassword)

	t.Run("Success Forget Password (request reset password link through email)", func(t *testing.T) {
		path := fmt.Sprintf("/forget-password/%s", cfg.Email.ReceiverTestEmail)
		req, _ := http.NewRequest("POST", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Failed Forget Password: invalid email : 400", func(t *testing.T) {
		path := fmt.Sprintf("/forget-password/%s", "some-invalid-email")
		req, _ := http.NewRequest("POST", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Failed Forget Password: email not found : 404", func(t *testing.T) {
		path := fmt.Sprintf("/forget-password/%s", "notfoundemail@example.com")
		req, _ := http.NewRequest("POST", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
