package server

import (
	"github.com/Binaretech/classroom-main/errors"
	"github.com/Binaretech/classroom-main/handler"
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errors.Handler,
	})

	api := app.Group("/api")

	api.Get("/user", handler.User)
	api.Post("/user", handler.StoreUser)
	api.Put("/user", handler.UpdateUser)

	// api.Get("/sections", handler.UserSections)
	return app
}
