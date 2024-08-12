package controllers

import (
	"encoding/json"
	"net/http"
	"task-manager/services"
	"task-manager/utils"
)

// UserController handles user-related HTTP requests.
type UserController struct {
	UserService *services.UserService
}

// Register handles user registration requests.
// It expects a JSON payload with "email" and "password" fields.
// On success, it returns a JWT token in the response.
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid request", nil)
		return
	}

	// Check if email or password is empty
	if input.Email == "" || input.Password == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Email and password are required", nil)
		return
	}

	user, err := uc.UserService.Register(input.Email, input.Password)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "error", "Could not generate token", nil)
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "success", "User registered successfully", map[string]string{"token": token})
}

// Login handles user login requests.
// It expects a JSON payload with "email" and "password" fields.
// On success, it returns a JWT token in the response.
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "error", "Invalid input", nil)
		return
	}

	user, err := uc.UserService.Authenticate(input.Email, input.Password)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusUnauthorized, "error", "Invalid email or password", nil)
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, "error", "Could not generate token", nil)
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, "success", "Login successful", map[string]string{"token": token})
}
