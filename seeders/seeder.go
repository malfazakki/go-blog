package seeders

import "gorm.io/gorm"

func RunSeeders(db *gorm.DB) {
	//  Send users first
	SeedUsers(db)
	// Seed Categories
	SeedCategories(db)
	// Seed posts (depends on users and categories)
	SeedPosts(db)
}
