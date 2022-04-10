package errors

import (
	"encoding/json"
	"errors"

	"github.com/Binaretech/classroom-main/lang"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Handler catch all the errors and returns a propper response based on the error type
func Handler(c *fiber.Ctx, err error) error {

	if e, ok := err.(*json.UnmarshalTypeError); ok {
		return response(c, NewBadRequestWrapped(lang.Trans("invalid data type"), e))
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response(c, WrapError(err, lang.Trans("not found"), fiber.StatusNotFound))
	}

	if e, ok := err.(ServerError); ok {
		return response(c, e)
	}

	return response(c, WrapError(err, lang.Trans("internal error"), fiber.StatusInternalServerError))
}

func response(c *fiber.Ctx, err ServerError) error {
	message := fiber.Map{
		"error": err.GetMessage(),
	}

	if viper.GetBool("debug") {
		message["debug"] = err.Error()
	}

	return c.Status(int(err.GetCode())).JSON(message)
}
