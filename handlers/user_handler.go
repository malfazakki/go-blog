package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/malfazakki/go-blog/models"
	"github.com/malfazakki/go-blog/repositories"
)

type UserHandler struct {
	userRepo repositories.UserRepository
}

func NewUserHandler(userRepo repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Hashing the Password, but save for later

	err = h.userRepo.Create(&user)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	// Make password empty in the reponse
	user.Password = ""
	SuccessResponse(w, http.StatusCreated, "User created successfully", user)

}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get ID from URL query parameters
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		ErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.userRepo.FindByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	// Make the password empty
	user.Password = ""
	SuccessResponse(w, http.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		ErrorResponse(w, http.StatusBadRequest, "User ID is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	existingUser, err := h.userRepo.FindByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if updatedUser.Name != "" {
		existingUser.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		existingUser.Email = updatedUser.Email
	}
	if updatedUser.Password != "" {
		existingUser.Password = updatedUser.Password
	}

	err = h.userRepo.Update(existingUser)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error updating user")
		return
	}

	// Make password empty in the reponse
	existingUser.Password = ""
	SuccessResponse(w, http.StatusOK, "User updated successfully", existingUser)
}
