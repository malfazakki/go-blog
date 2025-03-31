package routes

import (
	"github.com/gorilla/mux"
	"github.com/malfazakki/go-blog/handlers"
	"github.com/malfazakki/go-blog/middleware"
)

func SetupUserRoutes(router *mux.Router, handler *handlers.UserHandler) {
	// Create a subrouter for user endpoints
	userRouter := router.PathPrefix("/users").Subrouter()

	// Public routes
	userRouter.HandleFunc("", handler.CreateUser).Methods("POST")

	// Use Auth Middleware for protecting routes
	userRouter.Use(middleware.AuthMiddleware)

	// Private routes
	userRouter.HandleFunc("", handler.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", handler.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
}
