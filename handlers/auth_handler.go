package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/malfazakki/go-blog/models"
	"github.com/malfazakki/go-blog/repositories"
	"github.com/malfazakki/go-blog/utils"
)

type AuthHandler struct {
	userRepo repositories.UserRepository
}

func NewAuthHandler(userRepo repositories.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo}
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the login response body
type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Login handles user authentication
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Find user by email
	user, err := h.userRepo.FindByEmail(loginReq.Email)
	if err != nil {
		ErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Check password
	if !utils.CheckPasswordHash(loginReq.Password, user.Password) {
		ErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	// Return Empty password
	user.Password = ""

	// Return token and user
	SuccessResponse(w, http.StatusOK, "Login successful",
		LoginResponse{
			Token: token,
			User:  *user,
		})
}
