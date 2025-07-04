package models

import (
	"time"

	"gorm.io/gorm"
)

// Blog represents a blog post in the system
// @Description Blog post entity with all required fields
type Blog struct {
	ID          string         `json:"id" gorm:"primaryKey;type:varchar(36)" example:"550e8400-e29b-41d4-a716-446655440000"`
	Title       string         `json:"title" gorm:"type:varchar(255);not null" example:"My First Blog Post"`
	Description string         `json:"description" gorm:"type:text" example:"This is a brief description of my blog post"`
	Body        string         `json:"body" gorm:"type:text;not null" example:"This is the main content of my blog post..."`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime" example:"2023-01-01T00:00:00Z"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// BlogCreateRequest represents the request structure for creating a blog post
// @Description Request model for creating a new blog post
type BlogCreateRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255" example:"My First Blog Post"`
	Description string `json:"description" validate:"max=1000" example:"This is a brief description of my blog post"`
	Body        string `json:"body" validate:"required,min=1" example:"This is the main content of my blog post..."`
}

// BlogUpdateRequest represents the request structure for updating a blog post
// @Description Request model for updating an existing blog post
type BlogUpdateRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1,max=255" example:"Updated Blog Post Title"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=1000" example:"Updated description"`
	Body        *string `json:"body,omitempty" validate:"omitempty,min=1" example:"Updated blog post content..."`
}

// BlogResponse represents the response structure for blog posts
// @Description Response model for blog post data
type BlogResponse struct {
	ID          string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Title       string    `json:"title" example:"My First Blog Post"`
	Description string    `json:"description" example:"This is a brief description of my blog post"`
	Body        string    `json:"body" example:"This is the main content of my blog post..."`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}
