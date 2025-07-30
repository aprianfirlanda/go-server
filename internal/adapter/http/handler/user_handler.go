package handler

import (
	"net/http"

	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// User represents a sample user object
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUsersHandler handles GET /api/v1/users
func GetUsersHandler(c *fiber.Ctx) error {
	users := []User{
		{ID: 1, Name: "Aprian", Email: "aprian@example.com"},
		{ID: 2, Name: "Firlanda", Email: "firlanda@example.com"},
	}

	logger.Log.WithFields(map[string]interface{}{
		"endpoint": "/api/v1/users",
		"count":    len(users),
	}).Info("Fetched user list")

	return c.Status(http.StatusOK).JSON(users)
}
