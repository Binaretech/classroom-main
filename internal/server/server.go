package server

import (
	"github.com/Binaretech/classroom-main/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	api := app.Group("/api")
	api.Get("/user", handler.User)
	api.Post("/user", handler.StoreUser)
	api.Put("/user", handler.UpdateUser)
	return app
}
