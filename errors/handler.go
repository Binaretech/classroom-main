package errors

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Binaretech/classroom-main/lang"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Handler catch all the errors and returns a propper response based on the error type
func Handler(err error, c echo.Context) {
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		response(c, NewBadRequestWrapped(lang.Trans("invalid data type"), e))
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response(c, WrapError(err, lang.Trans("not found"), http.StatusNotFound))
		return
	}

	if e, ok := err.(ServerError); ok {
		response(c, e)
		return
	}

	response(c, WrapError(err, lang.Trans("internal error"), http.StatusInternalServerError))
	return
}

func response(c echo.Context, err ServerError) error {
	message := echo.Map{
		"error": err.GetMessage(),
	}

	if viper.GetBool("debug") {
		message["debug"] = err.Error()
	}

	return c.JSON(int(err.GetCode()), message)
}
