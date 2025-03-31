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
	postRouter.Use(middleware.AuthMiddleware)

	// Private routes
	postRouter.HandleFunc("", handler.CreatePost).Methods("POST")
	postRouter.HandleFunc("/{id:[0-9]+}", handler.Update).Methods("PUT")
	postRouter.HandleFunc("/{id:[0-9]+}", handler.DeletePost).Methods("DELETE")
}
