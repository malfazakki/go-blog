package repositories

import (
	"github.com/malfazakki/go-blog/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	FindByID(id uint) (*models.Post, error)
	FindAll() ([]models.Post, error)
	FindByUser(userID uint) ([]models.Post, error)
	Update(post *models.Post) error
	Delete(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

// NewPostRepository creates a new instance of PostRepository
func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.Preload("User").Preload("Categories").First(&post, id).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("User").Preload("Categories").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepository) FindByUser(userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Where("user_id = ?", userID).Preload("Categories").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}
