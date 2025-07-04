package controller

import (
	"BlogManagment/internal/models"
	"BlogManagment/internal/service"
	_ "strconv"

	"github.com/gofiber/fiber/v2"
)

// BlogController handles HTTP requests for blog operations
type BlogController struct {
	blogService service.BlogService
}

// NewBlogController creates a new blog controller instance
func NewBlogController(blogService service.BlogService) *BlogController {
	return &BlogController{blogService: blogService}
}

// CreateBlog handles POST /api/blog-post
// @Summary Create a new blog post
// @Description Create a new blog post with title, description, and body
// @Tags blog
// @Accept json
// @Produce json
// @Param blog body models.BlogCreateRequest true "Blog post data"
// @Success 201 {object} map[string]interface{} "Blog post created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog-post [post]
func (c *BlogController) CreateBlog(ctx *fiber.Ctx) error {
	var request models.BlogCreateRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	blog, err := c.blogService.CreateBlog(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to create blog post",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Blog post created successfully",
		"data":    blog,
	})
}

// GetBlogByID handles GET /api/blog-post/:id
// @Summary Get a blog post by ID
// @Description Retrieve a specific blog post by its unique identifier
// @Tags blog
// @Accept json
// @Produce json
// @Param id path string true "Blog post ID"
// @Success 200 {object} map[string]interface{} "Blog post retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid ID"
// @Failure 404 {object} map[string]interface{} "Blog post not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog-post/{id} [get]
func (c *BlogController) GetBlogByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Blog ID is required",
			"message": "Please provide a valid blog ID",
		})
	}

	blog, err := c.blogService.GetBlogByID(id)
	if err != nil {
		if err.Error() == "blog post not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "Blog post not found",
				"message": "The requested blog post does not exist",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve blog post",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Blog post retrieved successfully",
		"data":    blog,
	})
}

// GetAllBlogs handles GET /api/blog-post
// @Summary Get all blog posts
// @Description Retrieve all blog posts from the database
// @Tags blog
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Blog posts retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog-post [get]
func (c *BlogController) GetAllBlogs(ctx *fiber.Ctx) error {
	blogs, err := c.blogService.GetAllBlogs()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve blog posts",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Blog posts retrieved successfully",
		"data":    blogs,
		"count":   len(blogs),
	})
}

// UpdateBlog handles PATCH /api/blog-post/:id
// @Summary Update a blog post
// @Description Update an existing blog post by ID with partial data
// @Tags blog
// @Accept json
// @Produce json
// @Param id path string true "Blog post ID"
// @Param blog body models.BlogUpdateRequest true "Blog post update data"
// @Success 200 {object} map[string]interface{} "Blog post updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - validation error"
// @Failure 404 {object} map[string]interface{} "Blog post not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog-post/{id} [patch]
func (c *BlogController) UpdateBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Blog ID is required",
			"message": "Please provide a valid blog ID",
		})
	}

	var request models.BlogUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"message": err.Error(),
		})
	}

	blog, err := c.blogService.UpdateBlog(id, &request)
	if err != nil {
		if err.Error() == "blog post not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "Blog post not found",
				"message": "The requested blog post does not exist",
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Failed to update blog post",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Blog post updated successfully",
		"data":    blog,
	})
}

// DeleteBlog handles DELETE /api/blog-post/:id
// @Summary Delete a blog post
// @Description Delete a blog post by its unique identifier
// @Tags blog
// @Accept json
// @Produce json
// @Param id path string true "Blog post ID"
// @Success 200 {object} map[string]interface{} "Blog post deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid ID"
// @Failure 404 {object} map[string]interface{} "Blog post not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /blog-post/{id} [delete]
func (c *BlogController) DeleteBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Blog ID is required",
			"message": "Please provide a valid blog ID",
		})
	}

	err := c.blogService.DeleteBlog(id)
	if err != nil {
		if err.Error() == "blog post not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "Blog post not found",
				"message": "The requested blog post does not exist",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete blog post",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Blog post deleted successfully",
	})
}
