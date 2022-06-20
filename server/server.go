package server

import (
	"fmt"

	"github.com/Binaretech/classroom-main/errors"
	"github.com/Binaretech/classroom-main/handler"
	"github.com/Binaretech/classroom-main/validation"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Listen(db *gorm.DB) error {
	app := echo.New()

	app.HTTPErrorHandler = func(err error, ctx echo.Context) {
		errors.Handler(err, ctx)
	}

	app.Validator = validation.SetUpValidator(db)

	api := app.Group("/api")

	handler := handler.New(db)
	handler.Routes(api)

	return app.Start(fmt.Sprintf(":%s", viper.GetString("port")))
}
