package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jay/dadmail/internal/auth"
	"github.com/jay/dadmail/internal/config"
	"github.com/jay/dadmail/internal/repository"
	"github.com/jmoiron/sqlx"
)

// SetupRoutes configures all API routes
func SetupRoutes(app *fiber.App, cfg *config.Config, db *sqlx.DB) {
	// Initialize services
	jwtService := auth.NewJWTService(cfg.JWT.Secret, cfg.JWT.AccessTokenTTL, cfg.JWT.RefreshTokenTTL)
	authHandler := NewAuthHandler(db, cfg)
	userRepo := repository.NewUserRepository(db)

	// API v1 group
	v1 := app.Group("/api/v1")

	// Auth routes (public)
	authGroup := v1.Group("/auth")
	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/refresh", authHandler.Refresh)
	authGroup.Post("/logout", authHandler.Logout)

	// Protected routes
	protected := v1.Group("", auth.AuthMiddleware(jwtService))

	// User routes (protected)
	users := protected.Group("/users")
	users.Get("/me", func(c *fiber.Ctx) error {
		userID, err := auth.GetUserID(c)
		if err != nil {
			return err
		}

		user, err := userRepo.GetByID(userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(fiber.Map{
			"id":           user.ID,
			"email":        user.Email,
			"full_name":    user.FullName,
			"role":         user.Role,
			"created_at":   user.CreatedAt,
			"last_login_at": user.LastLoginAt,
		})
	})

	users.Patch("/me", func(c *fiber.Ctx) error {
		userID, err := auth.GetUserID(c)
		if err != nil {
			return err
		}

		var req struct {
			FullName string `json:"full_name"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if req.FullName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Full name is required",
			})
		}

		if err := userRepo.Update(userID, req.FullName); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update user",
			})
		}

		return c.JSON(fiber.Map{
			"message": "User updated successfully",
		})
	})

	// Email routes (protected)
	emails := protected.Group("/emails")
	emails.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "List emails - coming soon",
		})
	})

	emails.Get("/:id", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Get email - coming soon",
		})
	})

	emails.Post("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Send email - coming soon",
		})
	})

	emails.Get("/categories/:category", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Get emails by category - coming soon",
		})
	})

	// Caregiver routes (protected, caregiver role required)
	caregivers := protected.Group("/caregivers")
	caregivers.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Caregiver dashboard - coming soon",
		})
	})

	caregivers.Post("/rules", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Create rule - coming soon",
		})
	})

	caregivers.Get("/activity", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "Activity log - coming soon",
		})
	})
}
