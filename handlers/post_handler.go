package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/malfazakki/go-blog/models"
	"github.com/malfazakki/go-blog/repositories"
)

type PostHandler struct {
	postRepo     repositories.PostRepository
	categoryRepo repositories.CategoryRepository
}

func NewPostHandler(postRepo repositories.PostRepository, categoryRepo repositories.CategoryRepository) *PostHandler {
	return &PostHandler{postRepo, categoryRepo}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.postRepo.Create(&post)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error creating post")
		return
	}

	SuccessResponse(w, http.StatusCreated, "Post created successfully", post)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract post ID from URL query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		// If no ID is provided, return all posts
		posts, err := h.postRepo.FindAll()
		if err != nil {
			ErrorResponse(w, http.StatusInternalServerError, "Error retrieving posts")
			return
		}
		SuccessResponse(w, http.StatusOK, "Posts retrieved successfully", posts)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	post, err := h.postRepo.FindByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	SuccessResponse(w, http.StatusOK, "Post retrieved successfully", post)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract post ID from URL Query Parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		ErrorResponse(w, http.StatusBadRequest, "Post ID is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	existingPost, err := h.postRepo.FindByID(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusNotFound, "Post not found")
		return
	}

	var updatedPost models.Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update only the fields that are provided
	if updatedPost.Title != "" {
		existingPost.Title = updatedPost.Title
	}
	if updatedPost.Content != "" {
		existingPost.Content = updatedPost.Content
	}
	// Handle category updates if provided
	if len(updatedPost.Categories) > len(existingPost.Categories) {
		existingPost.Categories = updatedPost.Categories
	}

	err = h.postRepo.Update(existingPost)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error updating post")
		return
	}

	SuccessResponse(w, http.StatusOK, "Post updated successsfully", existingPost)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract ID from URL query parameters
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		ErrorResponse(w, http.StatusBadRequest, "Post ID is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	err = h.postRepo.Delete(uint(id))
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "Error deleting post")
		return
	}

	SuccessResponse(w, http.StatusOK, "Post deleted successfully", nil)
}
