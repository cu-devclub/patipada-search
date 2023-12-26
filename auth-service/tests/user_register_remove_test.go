package tests

import (
	"auth-service/config"
	"auth-service/jwt"
	"auth-service/users/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAndRemove(t *testing.T) {
	handlers := setUpTestEnvironment()
	cfg := config.GetConfig()
	e := echo.New()
	e.POST("/register", handlers.RegisterUser)
	e.DELETE("/user/:username", handlers.RemoveUser)
	t.Run("Success Register (User) : 201 & Success Remove 200", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		// Generate Super admin token to remove created user
		token, err := jwt.CreateToken(cfg.User.SuperAdmin.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		path := fmt.Sprintf("/user/%s", m.Username)
		req, _ = http.NewRequest("DELETE", path, nil)
		req.Header.Set("Authorization", token)
		w = httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("Success Register (Admin) : 201 & Success Remove 200", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		m.Role = "admin"
		payload, _ := json.Marshal(m)
		
		// Generate admin token to create admin
		adminToken, err := jwt.CreateToken(cfg.User.Admins.Username, cfg.User.Admins.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		req.Header.Set("Authorization", adminToken)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		// Generate Super admin token to remove created user
		token, err := jwt.CreateToken(cfg.User.SuperAdmin.Username, cfg.User.SuperAdmin.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		path := fmt.Sprintf("/user/%s", m.Username)
		req, _ = http.NewRequest("DELETE", path, nil)
		req.Header.Set("Authorization", token)
		w = httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})
	t.Run("Failed Register : Username already exists : 400", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		// use super admin credentials (already insert in set up env (migration))
		m.Username = cfg.User.SuperAdmin.Username

		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		// retrive respond body
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		// check message
		message := responseBody["message"]
		assert.Equal(t, "Username already exists", message)
	})

	t.Run("Failed Register : Email already exists : 400", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		// use super admin credentials (already insert in set up env (migration))
		m.Email = cfg.User.SuperAdmin.Email

		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		// retrive respond body
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatalf("Error parsing response body: %v", err)
		}

		// check message
		message := responseBody["message"]
		assert.Equal(t, "Email already exists", message)
	})

	t.Run("Failed Register Bad request field : 400 ", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		// invalid email , short password
		m.Email = "9pZoZ"
		m.Password = "123"

		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Failed Register: no permission (not attach token when needed) : 400", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		m.Role = "admin"

		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Failed Register: invalid token : 401", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		m.Role = "admin"

		token := "Some Missing token"
		// Making a request
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Failed Register: missing token : 401", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		m.Role = "admin"

		// Making a request
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Failed Register: no permission : 409", func(t *testing.T) {
		m := models.RegisterDto{}
		m.MockData()
		m.Role = "admin"

		// Generate user token
		token, err := jwt.CreateToken(cfg.User.Users.Username, cfg.User.Users.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		// Making a request
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
	})

	//* ------ Remove
	t.Run("Failed Remove ; Not attach username ; 404", func(t *testing.T) {
		// Generate user token
		token, err := jwt.CreateToken(cfg.User.Users.Username, cfg.User.Users.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		path := fmt.Sprintf("/user/%s", "")
		req, _ := http.NewRequest("DELETE", path, nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Failed Remove ; Not found username : 404", func(t *testing.T) {
		// Generate user token
		token, err := jwt.CreateToken(cfg.User.Users.Username, cfg.User.Users.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		path := fmt.Sprintf("/user/%s", "unkw0nusername")
		req, _ := http.NewRequest("DELETE", path, nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Failed Remove ; No permission : 403", func(t *testing.T) {
		// Insert some random user
		m := models.RegisterDto{}
		m.MockData()
		payload, _ := json.Marshal(m)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header
		w := httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		// ---- Try to remove with same role `user` ------
		// Same role return 403 forbidden
		// Generate user token
		token, err := jwt.CreateToken(cfg.User.Users.Username, cfg.User.Users.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}

		path := fmt.Sprintf("/user/%s", m.Username)
		req, _ = http.NewRequest("DELETE", path, nil)
		req.Header.Set("Authorization", token)
		w = httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)

		// --- Make sure to delete it with admin-----
		// Generate admin token
		token, err = jwt.CreateToken(cfg.User.Admins.Username, cfg.User.Admins.Role)
		if err != nil {
			t.Fatalf("Error creating token: %v", err)
		}
		path = fmt.Sprintf("/user/%s", m.Username)
		req, _ = http.NewRequest("DELETE", path, nil)
		req.Header.Set("Authorization", token)
		w = httptest.NewRecorder()
		e.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
