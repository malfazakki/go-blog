package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/malfazakki/go-blog/models"
	"github.com/malfazakki/go-blog/repositories"
)

type CategoryHandler struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryHandler(categoryRepo repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{categoryRepo}
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.categoryRepo.Create(&category)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error creating category")
		return
	}

	SuccessResponse(w, http.StatusOK, "Category created successfully", category)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract ID from URL query parameters
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		// If no ID is provided, return all categories
		categories, err := h.categoryRepo.FindAll()
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, "Error retrieving categories")
			return
		}

		SuccessResponse(w, http.StatusOK, "Categories retrieved successfully", categories)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.categoryRepo.FindByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Category not found")
		return
	}

	SuccessResponse(w, http.StatusOK, "Category retrieved successfully", category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract category ID from URL query parameters
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		ErrorResponse(w, http.StatusBadRequest, "Category ID is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	existingCategory, err := h.categoryRepo.FindByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Category not found")
		return
	}

	var updatedCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update only if value provided
	if updatedCategory.Name != "" {
		existingCategory.Name = updatedCategory.Name
	}

	err = h.categoryRepo.Update(existingCategory)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error updating category")
		return
	}

	SuccessResponse(w, http.StatusOK, "Category updated successfully", existingCategory)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract category ID from URL query parameters
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		ErrorResponse(w, http.StatusBadRequest, "Category ID is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	err = h.categoryRepo.Delete(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error deleting category")
		return
	}

	SuccessResponse(w, http.StatusOK, "Category deleted successfully", nil)
}
