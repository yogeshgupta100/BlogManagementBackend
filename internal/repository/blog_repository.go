package repository

import (
	"BlogManagment/internal/models"
	"errors"

	"gorm.io/gorm"
)

// BlogRepository defines the interface for blog data operations
type BlogRepository interface {
	Create(blog *models.Blog) error
	GetByID(id string) (*models.Blog, error)
	GetAll() ([]models.Blog, error)
	Update(blog *models.Blog) error
	Delete(id string) error
}

// blogRepository implements BlogRepository interface
type blogRepository struct {
	db *gorm.DB
}

// NewBlogRepository creates a new blog repository instance
func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}

// Create adds a new blog post to the database
func (r *blogRepository) Create(blog *models.Blog) error {
	result := r.db.Create(blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByID retrieves a blog post by its ID
func (r *blogRepository) GetByID(id string) (*models.Blog, error) {
	var blog models.Blog
	result := r.db.Where("id = ?", id).First(&blog)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("blog post not found")
		}
		return nil, result.Error
	}
	return &blog, nil
}

// GetAll retrieves all blog posts from the database
func (r *blogRepository) GetAll() ([]models.Blog, error) {
	var blogs []models.Blog
	result := r.db.Order("created_at DESC").Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

// Update modifies an existing blog post
func (r *blogRepository) Update(blog *models.Blog) error {
	result := r.db.Save(blog)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("blog post not found")
	}
	return nil
}

// Delete removes a blog post from the database
func (r *blogRepository) Delete(id string) error {
	result := r.db.Where("id = ?", id).Delete(&models.Blog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("blog post not found")
	}
	return nil
} 