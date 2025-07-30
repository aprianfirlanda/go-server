package handler

import (
	"net/http"

	"github.com/aprianfirlanda/go-server/internal/domain/port"
	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.FetchUsers()
	if err != nil {
		logger.Log.WithError(err).Error("Failed to fetch users")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	logger.Log.WithField("count", len(users)).Info("Users fetched")
	return c.JSON(users)
}
