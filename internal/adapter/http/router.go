package http

import (
	"github.com/aprianfirlanda/go-server/internal/adapter/http/handler"
	"github.com/aprianfirlanda/go-server/internal/adapter/repository"
	"github.com/aprianfirlanda/go-server/internal/database"
	"github.com/aprianfirlanda/go-server/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userRepo := repository.NewUserRepository(database.DB)
	userService := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app.Get("/", handler.RootHandler)
	app.Get("/health", handler.HealthCheckHandler)

	api := app.Group("/api")
	apiV1 := api.Group("/v1")

	apiV1User := apiV1.Group("/users")
	apiV1User.Post("/", userHandler.CreateUser)
	apiV1User.Get("/", userHandler.GetUsers)
}
