package seeders

import (
	"log"

	"github.com/malfazakki/go-blog/models"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	// Check if categories already exist
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count > 0 {
		log.Println("Categories already seeded")
		return
	}

	// Create sample categories
	categories := []models.Category{
		{
			Name: "Technology",
		},
		{
			Name: "Programming",
		},
		{
			Name: "Web Development",
		},
		{
			Name: "Mobile Development",
		},
		{
			Name: "DevOps",
		},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			log.Printf("Error creating category: %v", err)
			return
		}
	}

	log.Println("Categories seeded successfully")
}
