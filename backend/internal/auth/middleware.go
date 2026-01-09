package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// AuthMiddleware creates a middleware for JWT authentication
func AuthMiddleware(jwtService *JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		// Check Bearer scheme
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		tokenString := parts[1]

		// Validate token
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Store user info in context
		c.Locals("user_id", claims.UserID)
		c.Locals("user_email", claims.Email)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}

// GetUserID extracts the user ID from the context
func GetUserID(c *fiber.Ctx) (uuid.UUID, error) {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}
	return userID, nil
}

// GetUserEmail extracts the user email from the context
func GetUserEmail(c *fiber.Ctx) (string, error) {
	email, ok := c.Locals("user_email").(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}
	return email, nil
}

// GetUserRole extracts the user role from the context
func GetUserRole(c *fiber.Ctx) (string, error) {
	role, ok := c.Locals("user_role").(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}
	return role, nil
}

// RequireRole creates a middleware that checks if the user has a specific role
func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, err := GetUserRole(c)
		if err != nil {
			return err
		}

		for _, role := range roles {
			if userRole == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Insufficient permissions",
		})
	}
}
