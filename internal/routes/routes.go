package routes

import (
	"BlogManagment/internal/controller"
	"BlogManagment/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, blogController *controller.BlogController) {
	// Global middleware
	app.Use(middleware.Logger())

	// API routes group
	api := app.Group("/api")

	// Blog routes
	blogRoutes := api.Group("/blog-post")
	blogRoutes.Post("/", blogController.CreateBlog)      // POST /api/blog-post
	blogRoutes.Get("/", blogController.GetAllBlogs)      // GET /api/blog-post
	blogRoutes.Get("/:id", blogController.GetBlogByID)   // GET /api/blog-post/:id
	blogRoutes.Patch("/:id", blogController.UpdateBlog)  // PATCH /api/blog-post/:id
	blogRoutes.Delete("/:id", blogController.DeleteBlog) // DELETE /api/blog-post/:id

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "Blog Management API is running",
		})
	})

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Not Found",
			"message": "The requested endpoint does not exist",
		})
	})
}
