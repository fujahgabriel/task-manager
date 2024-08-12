// controllers/user_controller_test.go

package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"task-manager/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserController_Register(t *testing.T) {
	userService := &services.UserService{}
	userController := &UserController{UserService: userService}

	t.Run("ValidRequest", func(t *testing.T) {
		// Create a request body with valid email and password
		requestBody := `{"email": "test@example.com", "password": "password123"}`

		// Create a new HTTP request
		req, _ := http.NewRequest(http.MethodPost, "/api/register", strings.NewReader(requestBody))

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the Register handler function
		userController.Register(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "User registered successfully", response["message"])
		assert.NotNil(t, response["data"])
		assert.NotNil(t, response["data"].(map[string]interface{})["token"])
	})

	t.Run("InvalidRequest", func(t *testing.T) {
		// Create a request body with missing email field
		requestBody := `{"password": ""}`

		// Create a new HTTP request
		req, _ := http.NewRequest(http.MethodPost, "/api/register", strings.NewReader(requestBody))

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the Register handler function
		userController.Register(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusBadRequest, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Email and password are required", response["message"])
		assert.Nil(t, response["data"])
	})
}

func TestUserController_Login(t *testing.T) {
	userService := &services.UserService{}
	userController := &UserController{UserService: userService}

	/* t.Run("ValidRequest", func(t *testing.T) {
		// Create a request body with valid email and password
		requestBody := `{"email": "test@example.com", "password": "password123"}`

		// Create a new HTTP request
		req, _ := http.NewRequest(http.MethodPost, "/api/login", strings.NewReader(requestBody))

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the Login handler function
		userController.Login(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Login successful", response["message"])
		assert.NotNil(t, response["data"])
		assert.NotNil(t, response["data"].(map[string]interface{})["token"])
	})*/

	t.Run("InvalidEmailPassword", func(t *testing.T) {
		// Create a request body with missing password field
		requestBody := `{"email": "test@example.com","password": ""}`

		// Create a new HTTP request
		req, _ := http.NewRequest(http.MethodPost, "/api/login", strings.NewReader(requestBody))

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Call the Login handler function
		userController.Login(rr, req)

		// Check the response status code
		assert.Equal(t, http.StatusUnauthorized, rr.Code)

		// Check the response body
		var response map[string]interface{}
		json.NewDecoder(rr.Body).Decode(&response)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Invalid email or password", response["message"])
		assert.Nil(t, response["data"])
	})
}
