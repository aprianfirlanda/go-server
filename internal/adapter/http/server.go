package http

import (
	"fmt"
	"os"

	"github.com/aprianfirlanda/go-server/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func NewFiberApp() *fiber.App {
	app := fiber.New()
	SetupRoutes(app) // <- delegate routing to router.go
	return app
}

func StartServer(app *fiber.App, port string) {
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
	}

	logger.Log.Infof("🚀 Server starting on port %s", port)

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		logger.Log.WithError(err).Fatal("❌ Failed to start server")
	}
}
