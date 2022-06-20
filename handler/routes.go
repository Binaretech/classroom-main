package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Routes(api *echo.Group) {
	api.GET("/user", h.User)
	api.POST("/user", h.StoreUser)
	api.PUT("/user", h.UpdateUser)

	api.GET("/sections", h.UserSections)
}
