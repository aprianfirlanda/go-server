package http

import (
	"github.com/aprianfirlanda/go-server/internal/adapter/http/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.RootHandler)
	app.Get("/health", handler.HealthCheckHandler)

	// You can add more route groups, e.g., API v1
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/users", handler.GetUsersHandler) // Example
}
