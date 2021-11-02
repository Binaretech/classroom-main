package server

import (
	"errors"
	"github.com/Binaretech/classroom-main/internal/lang"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	msg := lang.Trans("internal error")

	if errors.Is(err, gorm.ErrRecordNotFound) {
		msg = lang.Trans("not found")
		code = fiber.StatusNotFound
	}

	response := map[string]string{
		"message": msg,
	}

	if viper.GetBool("DEBUG") {
		response["error"] = err.Error()
	}

	return c.Status(code).JSON(response)
}
