package server

import (
	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/gofiber/fiber/v2"
)

// NewFiberApp sets up and returns a configured Fiber app.
func NewFiberApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Fiber!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}

// StartServer starts the Fiber app on a given port.
func StartServer(app *fiber.App, port string) {
	if port == "" {
		port = "8080"
	}
	logger.Log.Infof("🚀 Server starting on port %s", port)

	if err := app.Listen(":" + port); err != nil {
		logger.Log.WithError(err).Fatal("❌ Failed to start server")
	}
}
