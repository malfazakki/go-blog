package seeders

import (
	"log"
	"math/rand"
	"time"

	"github.com/malfazakki/go-blog/models"
	"gorm.io/gorm"
)

func SeedPosts(db *gorm.DB) {
	var count int64
	db.Model(&models.Post{}).Count(&count)
	if count > 0 {
		log.Println("Posts already seeded")
		return
	}

	// Get users
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		log.Printf("Error fetching users: %v", err)
		return
	}
	if len(users) == 0 {
		log.Println("No users found, please seed users first")
		return
	}

	// Get Categories
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Printf("Error fetching categories: %v", err)
		return
	}
	if len(categories) == 0 {
		log.Printf("No categorries found, please seed categories first")
		return
	}

	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// Create sample posts
	posts := []models.Post{
		{
			Title:   "Getting Started with Go",
			Content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
			UserID:  users[rand.Intn(len(users))].ID,
		},
		{
			Title:   "RESTful API Design",
			Content: "REST is an architectural style for designing networked applications. It relies on a stateless, client-server communication protocol, almost always HTTP.",
			UserID:  users[rand.Intn(len(users))].ID,
		},
		{
			Title:   "Database Design Patterns",
			Content: "Database design patterns are reusable solutions to common problems in database design. They help create efficient and maintainable database structures.",
			UserID:  users[rand.Intn(len(users))].ID,
		},
		{
			Title:   "Authentication with JWT",
			Content: "JSON Web Tokens (JWT) provide a way to securely transmit information between parties as a JSON object. This information can be verified and trusted because it is digitally signed.",
			UserID:  users[rand.Intn(len(users))].ID,
		},
		{
			Title:   "Clean Architecture in Go",
			Content: "Clean Architecture is a software design philosophy that separates the elements of a design into ring levels. The main rule of clean architecture is that code dependencies can only come from the outer levels inward.",
			UserID:  users[rand.Intn(len(users))].ID,
		},
	}

	// Save posts and assign random categories
	for i := range posts {
		if err := db.Create(&posts[i]).Error; err != nil {
			log.Printf("Error creating post: %v", err)
			return
		}

		// Assign 1-3 random categories to each post
		numCategories := rand.Intn(3) + 1
		for j := 0; j < numCategories; j++ {
			categoryIndex := rand.Intn(len(categories))
			if err := db.Model(&posts[i]).Association("Categories").Append(&categories[categoryIndex]); err != nil {
				log.Printf("Error assigning category to post: %v", err)
				return
			}
		}
	}

	log.Println("Posts seeded successfully")
}
