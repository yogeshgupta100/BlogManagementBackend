package service

import (
	"BlogManagment/internal/models"
	"BlogManagment/internal/repository"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBlogRepository is a mock implementation of BlogRepository
type MockBlogRepository struct {
	mock.Mock
}

func (m *MockBlogRepository) Create(blog *models.Blog) error {
	args := m.Called(blog)
	return args.Error(0)
}

func (m *MockBlogRepository) GetByID(id string) (*models.Blog, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Blog), args.Error(1)
}

func (m *MockBlogRepository) GetAll() ([]models.Blog, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Blog), args.Error(1)
}

func (m *MockBlogRepository) Update(blog *models.Blog) error {
	args := m.Called(blog)
	return args.Error(0)
}

func (m *MockBlogRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewBlogService(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	assert.NotNil(t, service)
	assert.IsType(t, &blogService{}, service)
}

func TestBlogService_CreateBlog_Success(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	request := &models.BlogCreateRequest{
		Title:       "Test Blog",
		Description: "Test Description",
		Body:        "Test Body",
	}
	
	mockRepo.On("Create", mock.AnythingOfType("*models.Blog")).Return(nil)
	
	response, err := service.CreateBlog(request)
	
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, request.Title, response.Title)
	assert.Equal(t, request.Description, response.Description)
	assert.Equal(t, request.Body, response.Body)
	assert.NotEmpty(t, response.ID)
	assert.NotZero(t, response.CreatedAt)
	assert.NotZero(t, response.UpdatedAt)
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_CreateBlog_ValidationError(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	// Test with nil request
	response, err := service.CreateBlog(nil)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "request cannot be nil", err.Error())
	
	// Test with empty title
	request := &models.BlogCreateRequest{
		Title: "",
		Body:  "Test Body",
	}
	response, err = service.CreateBlog(request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "title is required", err.Error())
	
	// Test with empty body
	request = &models.BlogCreateRequest{
		Title: "Test Title",
		Body:  "",
	}
	response, err = service.CreateBlog(request)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "body is required", err.Error())
}

func TestBlogService_CreateBlog_RepositoryError(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	request := &models.BlogCreateRequest{
		Title: "Test Blog",
		Body:  "Test Body",
	}
	
	mockRepo.On("Create", mock.AnythingOfType("*models.Blog")).Return(errors.New("database error"))
	
	response, err := service.CreateBlog(request)
	
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "database error", err.Error())
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_GetBlogByID_Success(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	blogID := uuid.New().String()
	expectedBlog := &models.Blog{
		ID:          blogID,
		Title:       "Test Blog",
		Description: "Test Description",
		Body:        "Test Body",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	mockRepo.On("GetByID", blogID).Return(expectedBlog, nil)
	
	response, err := service.GetBlogByID(blogID)
	
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, expectedBlog.ID, response.ID)
	assert.Equal(t, expectedBlog.Title, response.Title)
	assert.Equal(t, expectedBlog.Description, response.Description)
	assert.Equal(t, expectedBlog.Body, response.Body)
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_GetBlogByID_EmptyID(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	response, err := service.GetBlogByID("")
	
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "blog ID is required", err.Error())
}

func TestBlogService_GetBlogByID_NotFound(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	blogID := uuid.New().String()
	mockRepo.On("GetByID", blogID).Return(nil, errors.New("blog post not found"))
	
	response, err := service.GetBlogByID(blogID)
	
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "blog post not found", err.Error())
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_GetAllBlogs_Success(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	expectedBlogs := []models.Blog{
		{
			ID:          uuid.New().String(),
			Title:       "Blog 1",
			Description: "Description 1",
			Body:        "Body 1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Title:       "Blog 2",
			Description: "Description 2",
			Body:        "Body 2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	
	mockRepo.On("GetAll").Return(expectedBlogs, nil)
	
	responses, err := service.GetAllBlogs()
	
	assert.NoError(t, err)
	assert.NotNil(t, responses)
	assert.Len(t, responses, 2)
	assert.Equal(t, expectedBlogs[0].Title, responses[0].Title)
	assert.Equal(t, expectedBlogs[1].Title, responses[1].Title)
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_GetAllBlogs_Error(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	mockRepo.On("GetAll").Return(nil, errors.New("database error"))
	
	responses, err := service.GetAllBlogs()
	
	assert.Error(t, err)
	assert.Nil(t, responses)
	assert.Equal(t, "database error", err.Error())
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_UpdateBlog_Success(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	blogID := uuid.New().String()
	existingBlog := &models.Blog{
		ID:          blogID,
		Title:       "Original Title",
		Description: "Original Description",
		Body:        "Original Body",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	newTitle := "Updated Title"
	request := &models.BlogUpdateRequest{
		Title: &newTitle,
	}
	
	mockRepo.On("GetByID", blogID).Return(existingBlog, nil)
	mockRepo.On("Update", mock.AnythingOfType("*models.Blog")).Return(nil)
	
	response, err := service.UpdateBlog(blogID, request)
	
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, newTitle, response.Title)
	assert.Equal(t, existingBlog.Description, response.Description)
	assert.Equal(t, existingBlog.Body, response.Body)
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_UpdateBlog_EmptyID(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	request := &models.BlogUpdateRequest{}
	
	response, err := service.UpdateBlog("", request)
	
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "blog ID is required", err.Error())
}

func TestBlogService_UpdateBlog_NotFound(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	blogID := uuid.New().String()
	request := &models.BlogUpdateRequest{}
	
	mockRepo.On("GetByID", blogID).Return(nil, errors.New("blog post not found"))
	
	response, err := service.UpdateBlog(blogID, request)
	
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "blog post not found", err.Error())
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_UpdateBlog_ValidationError(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	blogID := uuid.New().String()
	existingBlog := &models.Blog{
		ID:          blogID,
		Title:       "Original Title",
		Description: "Original Description",
		Body:        "Original Body",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	emptyTitle := ""
	request := &models.BlogUpdateRequest{
		Title: &emptyTitle,
	}
	
	mockRepo.On("GetByID", blogID).Return(existingBlog, nil)
	
	response, err := service.UpdateBlog(blogID, request)
	
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "title cannot be empty", err.Error())
	
	mockRepo.AssertExpectations(t)
}

func TestBlogService_DeleteBlog_Success(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	blogID := uuid.New().String()
	mockRepo.On("Delete", blogID).Return(nil)
	
	err := service.DeleteBlog(blogID)
	
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlogService_DeleteBlog_EmptyID(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	err := service.DeleteBlog("")
	
	assert.Error(t, err)
	assert.Equal(t, "blog ID is required", err.Error())
}

func TestBlogService_DeleteBlog_NotFound(t *testing.T) {
	mockRepo := &MockBlogRepository{}
	service := NewBlogService(mockRepo)
	
	blogID := uuid.New().String()
	mockRepo.On("Delete", blogID).Return(errors.New("blog post not found"))
	
	err := service.DeleteBlog(blogID)
	
	assert.Error(t, err)
	assert.Equal(t, "blog post not found", err.Error())
	
	mockRepo.AssertExpectations(t)
} 