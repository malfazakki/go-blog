package routes

import (
	"github.com/gorilla/mux"
	"github.com/malfazakki/go-blog/handlers"
	"github.com/malfazakki/go-blog/middleware"
)

// SetupPostRoutes configures all post-related routes
func SetupPostRoutes(router *mux.Router, handler *handlers.PostHandler) {
	// Create a subrouter for post endpoints
	postRouter := router.PathPrefix("/posts").Subrouter()

	// Public routes
	postRouter.HandleFunc("", handler.GetPost).Methods("GET")
	postRouter.HandleFunc("/{id:[0-9]+}", handler.GetPost).Methods("GET")

	// Protecting the routes
	protectedPostRouter := router.PathPrefix("/posts").Subrouter()
	protectedPostRouter.Use(middleware.AuthMiddleware)

	// Private routes
	protectedPostRouter.HandleFunc("", handler.CreatePost).Methods("POST")
	protectedPostRouter.HandleFunc("/{id:[0-9]+}", handler.Update).Methods("PUT")
	protectedPostRouter.HandleFunc("/{id:[0-9]+}", handler.DeletePost).Methods("DELETE")
}
