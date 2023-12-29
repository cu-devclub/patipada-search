package tests

import (
	"auth-service/config"
	"auth-service/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.GET("/authorize", handlers.Authorize)
	t.Run("Success Authorize  : 200 & true", func(t *testing.T) {
		token, err := jwt.CreateToken(cfg.User.SuperAdmin.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			panic(err)
		}
		path := fmt.Sprintf("/authorize?requiredRole=%s", "super-admin")
		req, _ := http.NewRequest("GET", path, nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// retrive respond body
		var responseBody map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		result := responseBody["result"]
		assert.Equal(t, true, result)
	})

	t.Run("Success Authorize : 200 & false ", func(t *testing.T) {
		token, err := jwt.CreateToken(cfg.User.Users.Username, cfg.User.Users.Role)
		if err != nil {
			panic(err)
		}
		path := fmt.Sprintf("/authorize?requiredRole=%s", "super-admin")
		req, _ := http.NewRequest("GET", path, nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// retrive respond body
		var responseBody map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		result := responseBody["result"]
		assert.Equal(t, false, result)
	})

	t.Run("Failed  : 400 => Invalid require role  ", func(t *testing.T) {
		token, err := jwt.CreateToken(cfg.User.Users.Username, cfg.User.Users.Role)
		if err != nil {
			panic(err)
		}
		path := fmt.Sprintf("/authorize?requiredRole=%s", "superadmin")
		req, _ := http.NewRequest("GET", path, nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)

	})
	t.Run("Fail Authorize : 400 Missing token", func(t *testing.T) {
		path := fmt.Sprintf("/authorize?requiredRole=%s", "super-admin")
		req, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Fail Authorize : 401 Invalid token", func(t *testing.T) {
		token, err := jwt.CreateToken(cfg.User.SuperAdmin.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			panic(err)
		}
		path := fmt.Sprintf("/authorize?requiredRole=%s", "super-admin")
		req, _ := http.NewRequest("GET", path, nil)
		req.Header.Set("Authorization", token+"Malformed")
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
