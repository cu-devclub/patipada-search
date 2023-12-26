package tests

import (
	"auth-service/config"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestVerifyResetToken(t *testing.T) {

	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.GET("/verify-reset-token/:token", handlers.VerifyResetToken)
	token, err := getResetPasswordToken(cfg, cfg.Email.ReceiverTestEmail)
	if err != nil {
		panic(err)
	}
	// token := "56fc42cd361b467666c8174a72573ddb3b53753adbb912f1b758a20fd0c9e5c3"
	t.Run("Success Verify Reset Token: 200 & result true", func(t *testing.T) {
		path := fmt.Sprintf("/verify-reset-token/%s", token)
		req, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		//Get the response body
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}
		result := responseBody["result"]

		assert.Equal(t, true, result)
	})

	t.Run("Success Verify Reset Token: 200 & result false (wrong token)", func(t *testing.T) {
		path := fmt.Sprintf("/verify-reset-token/%s", " ")
		req, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		//Get the response body
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}
		result := responseBody["result"]

		assert.Equal(t, false, result)
	})

	t.Run("Success Verify Reset Token: 200 & result false (expire token)", func(t *testing.T) {
		// ! Need to create new token and wait 15 minute to test this expire token
		oldToken := ""
		path := fmt.Sprintf("/verify-reset-token/%s", oldToken)
		req, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		//Get the response body
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}
		result := responseBody["result"]

		assert.Equal(t, false, result)
	})
	t.Run("Failed Verify Reset Token: invalid token : 404", func(t *testing.T) {
		path := fmt.Sprintf("/verify-reset-token/%s", "")
		req, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

}
