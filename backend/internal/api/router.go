package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jay/dadmail/internal/config"
)

// SetupRoutes configures all API routes
func SetupRoutes(app *fiber.App, cfg *config.Config) {
	// API v1 group
	v1 := app.Group("/api/v1")

	// Auth routes
	auth := v1.Group("/auth")
	setupAuthRoutes(auth, cfg)

	// User routes (protected)
	users := v1.Group("/users")
	// TODO: Add auth middleware
	setupUserRoutes(users, cfg)

	// Email routes (protected)
	emails := v1.Group("/emails")
	// TODO: Add auth middleware
	setupEmailRoutes(emails, cfg)

	// Caregiver routes (protected)
	caregivers := v1.Group("/caregivers")
	// TODO: Add auth middleware
	setupCaregiverRoutes(caregivers, cfg)
}

// setupAuthRoutes configures authentication endpoints
func setupAuthRoutes(router fiber.Router, cfg *config.Config) {
	router.Post("/register", func(c *fiber.Ctx) error {
		// TODO: Implement registration
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Registration endpoint - coming soon",
		})
	})

	router.Post("/login", func(c *fiber.Ctx) error {
		// TODO: Implement login
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Login endpoint - coming soon",
		})
	})

	router.Post("/refresh", func(c *fiber.Ctx) error {
		// TODO: Implement token refresh
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Token refresh endpoint - coming soon",
		})
	})

	router.Post("/logout", func(c *fiber.Ctx) error {
		// TODO: Implement logout
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Logout endpoint - coming soon",
		})
	})
}

// setupUserRoutes configures user management endpoints
func setupUserRoutes(router fiber.Router, cfg *config.Config) {
	router.Get("/me", func(c *fiber.Ctx) error {
		// TODO: Get current user
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Get current user - coming soon",
		})
	})

	router.Patch("/me", func(c *fiber.Ctx) error {
		// TODO: Update current user
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Update user - coming soon",
		})
	})
}

// setupEmailRoutes configures email-related endpoints
func setupEmailRoutes(router fiber.Router, cfg *config.Config) {
	router.Get("/", func(c *fiber.Ctx) error {
		// TODO: List emails
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "List emails - coming soon",
		})
	})

	router.Get("/:id", func(c *fiber.Ctx) error {
		// TODO: Get email by ID
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Get email - coming soon",
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		// TODO: Send email
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Send email - coming soon",
		})
	})

	router.Get("/categories/:category", func(c *fiber.Ctx) error {
		// TODO: Get emails by category
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Get emails by category - coming soon",
		})
	})
}

// setupCaregiverRoutes configures caregiver-specific endpoints
func setupCaregiverRoutes(router fiber.Router, cfg *config.Config) {
	router.Get("/dashboard", func(c *fiber.Ctx) error {
		// TODO: Caregiver dashboard
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Caregiver dashboard - coming soon",
		})
	})

	router.Post("/rules", func(c *fiber.Ctx) error {
		// TODO: Create rule
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Create rule - coming soon",
		})
	})

	router.Get("/activity", func(c *fiber.Ctx) error {
		// TODO: View activity log
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Activity log - coming soon",
		})
	})
}
