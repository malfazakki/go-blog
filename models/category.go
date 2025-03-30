package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Posts []Post `gorm:"many2many:post_categories;" json:"posts"`
}
