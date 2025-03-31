package seeders

import (
	"log"

	"github.com/malfazakki/go-blog/models"
	"github.com/malfazakki/go-blog/utils"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	// Check if users already exist
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("Users already seeded")
		return
	}

	// Create sample users
	users := []models.User{
		{
			Name:     "Admin User",
			Email:    "admin@example.com",
			Password: "password123",
		},
		{
			Name:     "Regular User",
			Email:    "user@example.com",
			Password: "password123",
		},
		{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		},
	}

	// Hash password and save users
	for i := range users {
		hashedPassword, err := utils.HashPassword(users[i].Password)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
		}
		users[i].Password = hashedPassword

		if err := db.Create(&users[i]); err != nil {
			log.Printf("Error creating user: %v", err)
		}
	}

	log.Println("Users seeded successfully")
}
