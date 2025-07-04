package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ErrorHandler is a middleware that handles panics and errors
func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// Log the error
		log.Printf("Error: %v", err)

		// Default error response
		code := fiber.StatusInternalServerError
		message := "Internal Server Error"

		// Check if it's a fiber error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			message = e.Message
		}

		// Return JSON error response
		return c.Status(code).JSON(fiber.Map{
			"error":   "Request failed",
			"message": message,
		})
	}
}

// Logger is a middleware that logs HTTP requests
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Log request
		log.Printf("Request: %s %s", c.Method(), c.Path())

		// Continue to next middleware/handler
		err := c.Next()

		// Log response
		log.Printf("Response: %s %s - %d", c.Method(), c.Path(), c.Response().StatusCode())

		return err
	}
} 