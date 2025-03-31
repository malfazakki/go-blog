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
	postRepo := repositories.NewPostRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo)
	postHandler := handlers.NewPostHandler(postRepo, categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	// Setup specific route groups
	SetupUserRoutes(router, userHandler)
	SetupPostRoutes(router, postHandler)
	SetupCategoryRoutes(router, categoryHandler)

	return router
}
