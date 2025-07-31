package handler

import (
	"github.com/aprianfirlanda/go-server/internal/domain/model"
	"github.com/aprianfirlanda/go-server/internal/domain/port"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase port.UserUsecase
}

func NewUserHandler(usecase port.UserUsecase) *UserHandler {
	return &UserHandler{usecase}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if err := h.usecase.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.usecase.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get users"})
	}
	return c.JSON(users)
}
