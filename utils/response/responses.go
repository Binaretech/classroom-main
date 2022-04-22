package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResponseBadRequest(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": msg,
	})
}

func ResponseMessage(c echo.Context, status int, content interface{}) error {
	return c.JSON(status, map[string]interface{}{
		"message": content,
	})
}
