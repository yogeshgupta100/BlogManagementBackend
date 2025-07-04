package service

import (
	"BlogManagment/internal/models"
	"BlogManagment/internal/repository"
	"errors"
	"time"

	"github.com/google/uuid"
)

// BlogService defines the interface for blog business operations
type BlogService interface {
	CreateBlog(request *models.BlogCreateRequest) (*models.BlogResponse, error)
	GetBlogByID(id string) (*models.BlogResponse, error)
	GetAllBlogs() ([]models.BlogResponse, error)
	UpdateBlog(id string, request *models.BlogUpdateRequest) (*models.BlogResponse, error)
	DeleteBlog(id string) error
}

// blogService implements BlogService interface
type blogService struct {
	blogRepo repository.BlogRepository
}

// NewBlogService creates a new blog service instance
func NewBlogService(blogRepo repository.BlogRepository) BlogService {
	return &blogService{blogRepo: blogRepo}
}

// CreateBlog creates a new blog post
func (s *blogService) CreateBlog(request *models.BlogCreateRequest) (*models.BlogResponse, error) {
	// Validate request
	if err := s.validateCreateRequest(request); err != nil {
		return nil, err
	}

	// Create blog model
	blog := &models.Blog{
		ID:          uuid.New().String(),
		Title:       request.Title,
		Description: request.Description,
		Body:        request.Body,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save to database
	if err := s.blogRepo.Create(blog); err != nil {
		return nil, err
	}

	// Return response
	return s.blogToResponse(blog), nil
}

// GetBlogByID retrieves a blog post by ID
func (s *blogService) GetBlogByID(id string) (*models.BlogResponse, error) {
	if id == "" {
		return nil, errors.New("blog ID is required")
	}

	blog, err := s.blogRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.blogToResponse(blog), nil
}

// GetAllBlogs retrieves all blog posts
func (s *blogService) GetAllBlogs() ([]models.BlogResponse, error) {
	blogs, err := s.blogRepo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.BlogResponse, len(blogs))
	for i, blog := range blogs {
		responses[i] = *s.blogToResponse(&blog)
	}

	return responses, nil
}

// UpdateBlog updates an existing blog post
func (s *blogService) UpdateBlog(id string, request *models.BlogUpdateRequest) (*models.BlogResponse, error) {
	if id == "" {
		return nil, errors.New("blog ID is required")
	}

	// Get existing blog
	existingBlog, err := s.blogRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if request.Title != nil {
		if *request.Title == "" {
			return nil, errors.New("title cannot be empty")
		}
		existingBlog.Title = *request.Title
	}

	if request.Description != nil {
		existingBlog.Description = *request.Description
	}

	if request.Body != nil {
		if *request.Body == "" {
			return nil, errors.New("body cannot be empty")
		}
		existingBlog.Body = *request.Body
	}

	existingBlog.UpdatedAt = time.Now()

	// Save to database
	if err := s.blogRepo.Update(existingBlog); err != nil {
		return nil, err
	}

	return s.blogToResponse(existingBlog), nil
}

// DeleteBlog deletes a blog post
func (s *blogService) DeleteBlog(id string) error {
	if id == "" {
		return errors.New("blog ID is required")
	}

	return s.blogRepo.Delete(id)
}

// validateCreateRequest validates the create request
func (s *blogService) validateCreateRequest(request *models.BlogCreateRequest) error {
	if request == nil {
		return errors.New("request cannot be nil")
	}

	if request.Title == "" {
		return errors.New("title is required")
	}

	if request.Body == "" {
		return errors.New("body is required")
	}

	return nil
}

// blogToResponse converts a Blog model to BlogResponse
func (s *blogService) blogToResponse(blog *models.Blog) *models.BlogResponse {
	return &models.BlogResponse{
		ID:          blog.ID,
		Title:       blog.Title,
		Description: blog.Description,
		Body:        blog.Body,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}
} 