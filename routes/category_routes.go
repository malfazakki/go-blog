package routes

import (
	"github.com/gorilla/mux"
	"github.com/malfazakki/go-blog/handlers"
	"github.com/malfazakki/go-blog/middleware"
)

// SetupCategoryRoutes configures all category-related routes
func SetupCategoryRoutes(router *mux.Router, handler *handlers.CategoryHandler) {
	// Create a subrouter for category endpoints
	categoryRouter := router.PathPrefix("/categories").Subrouter()

	// Public routes
	categoryRouter.HandleFunc("", handler.GetCategory).Methods("GET")
	categoryRouter.HandleFunc("/{id:[0-9]+}", handler.GetCategory).Methods("GET")

	// Protecting the routes
	categoryRouter.Use(middleware.AuthMiddleware)

	// Private routes
	categoryRouter.HandleFunc("", handler.CreateCategory).Methods("POST")
	categoryRouter.HandleFunc("/{id:[0-9]+}", handler.UpdateCategory).Methods("PUT")
	categoryRouter.HandleFunc("/{id:[0-9]+}", handler.DeleteCategory).Methods("DELETE")
}
