package controller

import (
	"BlogManagment/internal/models"
	"BlogManagment/internal/service"
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBlogService is a mock implementation of BlogService
type MockBlogService struct {
	mock.Mock
}

func (m *MockBlogService) CreateBlog(request *models.BlogCreateRequest) (*models.BlogResponse, error) {
	args := m.Called(request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.BlogResponse), args.Error(1)
}

func (m *MockBlogService) GetBlogByID(id string) (*models.BlogResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.BlogResponse), args.Error(1)
}

func (m *MockBlogService) GetAllBlogs() ([]models.BlogResponse, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.BlogResponse), args.Error(1)
}

func (m *MockBlogService) UpdateBlog(id string, request *models.BlogUpdateRequest) (*models.BlogResponse, error) {
	args := m.Called(id, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.BlogResponse), args.Error(1)
}

func (m *MockBlogService) DeleteBlog(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// setupTestApp creates a test Fiber app with the blog controller
func setupTestApp() (*fiber.App, *MockBlogService) {
	app := fiber.New()
	mockService := &MockBlogService{}
	controller := NewBlogController(mockService)
	
	// Setup routes for testing
	app.Post("/api/blog-post", controller.CreateBlog)
	app.Get("/api/blog-post", controller.GetAllBlogs)
	app.Get("/api/blog-post/:id", controller.GetBlogByID)
	app.Patch("/api/blog-post/:id", controller.UpdateBlog)
	app.Delete("/api/blog-post/:id", controller.DeleteBlog)
	
	return app, mockService
}

func TestNewBlogController(t *testing.T) {
	mockService := &MockBlogService{}
	controller := NewBlogController(mockService)
	
	assert.NotNil(t, controller)
	assert.IsType(t, &BlogController{}, controller)
}

func TestBlogController_CreateBlog_Success(t *testing.T) {
	app, mockService := setupTestApp()
	
	requestBody := models.BlogCreateRequest{
		Title:       "Test Blog",
		Description: "Test Description",
		Body:        "Test Body",
	}
	
	expectedResponse := &models.BlogResponse{
		ID:          uuid.New().String(),
		Title:       requestBody.Title,
		Description: requestBody.Description,
		Body:        requestBody.Body,
	}
	
	mockService.On("CreateBlog", &requestBody).Return(expectedResponse, nil)
	
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/blog-post", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog post created successfully", result["message"])
	assert.NotNil(t, result["data"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_CreateBlog_InvalidBody(t *testing.T) {
	app, _ := setupTestApp()
	
	req := httptest.NewRequest("POST", "/api/blog-post", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Invalid request body", result["error"])
}

func TestBlogController_CreateBlog_ServiceError(t *testing.T) {
	app, mockService := setupTestApp()
	
	requestBody := models.BlogCreateRequest{
		Title: "Test Blog",
		Body:  "Test Body",
	}
	
	mockService.On("CreateBlog", &requestBody).Return(nil, errors.New("validation error"))
	
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/blog-post", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Failed to create blog post", result["error"])
	assert.Equal(t, "validation error", result["message"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_GetBlogByID_Success(t *testing.T) {
	app, mockService := setupTestApp()
	
	blogID := uuid.New().String()
	expectedResponse := &models.BlogResponse{
		ID:          blogID,
		Title:       "Test Blog",
		Description: "Test Description",
		Body:        "Test Body",
	}
	
	mockService.On("GetBlogByID", blogID).Return(expectedResponse, nil)
	
	req := httptest.NewRequest("GET", "/api/blog-post/"+blogID, nil)
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog post retrieved successfully", result["message"])
	assert.NotNil(t, result["data"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_GetBlogByID_EmptyID(t *testing.T) {
	app, _ := setupTestApp()
	
	req := httptest.NewRequest("GET", "/api/blog-post/", nil)
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestBlogController_GetBlogByID_NotFound(t *testing.T) {
	app, mockService := setupTestApp()
	
	blogID := uuid.New().String()
	mockService.On("GetBlogByID", blogID).Return(nil, errors.New("blog post not found"))
	
	req := httptest.NewRequest("GET", "/api/blog-post/"+blogID, nil)
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog post not found", result["error"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_GetAllBlogs_Success(t *testing.T) {
	app, mockService := setupTestApp()
	
	expectedResponses := []models.BlogResponse{
		{
			ID:          uuid.New().String(),
			Title:       "Blog 1",
			Description: "Description 1",
			Body:        "Body 1",
		},
		{
			ID:          uuid.New().String(),
			Title:       "Blog 2",
			Description: "Description 2",
			Body:        "Body 2",
		},
	}
	
	mockService.On("GetAllBlogs").Return(expectedResponses, nil)
	
	req := httptest.NewRequest("GET", "/api/blog-post", nil)
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog posts retrieved successfully", result["message"])
	assert.Equal(t, float64(2), result["count"])
	assert.NotNil(t, result["data"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_GetAllBlogs_ServiceError(t *testing.T) {
	app, mockService := setupTestApp()
	
	mockService.On("GetAllBlogs").Return(nil, errors.New("database error"))
	
	req := httptest.NewRequest("GET", "/api/blog-post", nil)
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Failed to retrieve blog posts", result["error"])
	assert.Equal(t, "database error", result["message"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_UpdateBlog_Success(t *testing.T) {
	app, mockService := setupTestApp()
	
	blogID := uuid.New().String()
	requestBody := models.BlogUpdateRequest{
		Title: stringPtr("Updated Title"),
	}
	
	expectedResponse := &models.BlogResponse{
		ID:          blogID,
		Title:       "Updated Title",
		Description: "Original Description",
		Body:        "Original Body",
	}
	
	mockService.On("UpdateBlog", blogID, &requestBody).Return(expectedResponse, nil)
	
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("PATCH", "/api/blog-post/"+blogID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog post updated successfully", result["message"])
	assert.NotNil(t, result["data"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_UpdateBlog_InvalidBody(t *testing.T) {
	app, _ := setupTestApp()
	
	blogID := uuid.New().String()
	req := httptest.NewRequest("PATCH", "/api/blog-post/"+blogID, bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Invalid request body", result["error"])
}

func TestBlogController_UpdateBlog_NotFound(t *testing.T) {
	app, mockService := setupTestApp()
	
	blogID := uuid.New().String()
	requestBody := models.BlogUpdateRequest{
		Title: stringPtr("Updated Title"),
	}
	
	mockService.On("UpdateBlog", blogID, &requestBody).Return(nil, errors.New("blog post not found"))
	
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("PATCH", "/api/blog-post/"+blogID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog post not found", result["error"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_DeleteBlog_Success(t *testing.T) {
	app, mockService := setupTestApp()
	
	blogID := uuid.New().String()
	mockService.On("DeleteBlog", blogID).Return(nil)
	
	req := httptest.NewRequest("DELETE", "/api/blog-post/"+blogID, nil)
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog post deleted successfully", result["message"])
	
	mockService.AssertExpectations(t)
}

func TestBlogController_DeleteBlog_NotFound(t *testing.T) {
	app, mockService := setupTestApp()
	
	blogID := uuid.New().String()
	mockService.On("DeleteBlog", blogID).Return(errors.New("blog post not found"))
	
	req := httptest.NewRequest("DELETE", "/api/blog-post/"+blogID, nil)
	
	resp, _ := app.Test(req)
	
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	assert.Equal(t, "Blog post not found", result["error"])
	
	mockService.AssertExpectations(t)
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
} 