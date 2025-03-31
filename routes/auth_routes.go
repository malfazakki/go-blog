package routes

import (
	"github.com/gorilla/mux"
	"github.com/malfazakki/go-blog/handlers"
)

// SetupAuthRoutes configures all authentication-related routes
func SetupAuthRoutes(router *mux.Router, handler *handlers.AuthHandler) {
	// Create a subrouter for auth endpoints
	authRouter := router.PathPrefix("/api/auth").Subrouter()

	// Register routes
	authRouter.HandleFunc("/login", handler.Login).Methods("POST")
}
