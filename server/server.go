package server

import (
	"fmt"

	"github.com/Binaretech/classroom-main/errors"
	"github.com/Binaretech/classroom-main/handler"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Listen() error {
	app := echo.New()

	app.HTTPErrorHandler = func(err error, ctx echo.Context) {
		errors.Handler(err, ctx)
	}

	api := app.Group("/api")

	api.GET("/user", handler.User)
	api.POST("/user", handler.StoreUser)
	api.PUT("/user", handler.UpdateUser)

	api.GET("/sections", handler.UserSections)

	return app.Start(fmt.Sprintf(":%s", viper.GetString("port")))
}
