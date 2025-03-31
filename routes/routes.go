package routes

import (
	"github.com/gorilla/mux"
	"github.com/malfazakki/go-blog/handlers"
	"github.com/malfazakki/go-blog/repositories"
	"gorm.io/gorm"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo)

	// Setup specific route groups
	SetupUserRoutes(router, userHandler)

	return router
}
