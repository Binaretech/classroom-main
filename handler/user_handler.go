package handler

import (
	"net/http"

	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/errors"
	"github.com/Binaretech/classroom-main/handler/request"
	"github.com/Binaretech/classroom-main/handler/service"
	"github.com/Binaretech/classroom-main/lang"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// User handler look up the user's data in the database and return it or returns 404 if not found
func (h *Handler) User(c echo.Context) error {
	userID := c.Request().Header.Get("X-User")

	user := model.User{}
	if err := h.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

// StoreUser create a new user in the database with the request data
func (h *Handler) StoreUser(c echo.Context) error {
	req := request.StoreUserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	req.ID = c.Request().Header.Get("X-User")

	if err := c.Validate(req); err != nil {
		return err
	}

	user := req.User()

	if err := h.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		file, _ := c.FormFile("image")

		if err := service.UpdateProfileImage(user, tx, file, user.ID); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return errors.NewInternalError(err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"user": user,
	})
}

// UpdateUser update the user's data in the database with the request data
func (h *Handler) UpdateUser(c echo.Context) error {
	req := request.UpdateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	req.ID = c.Request().Header.Get("X-User")

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := h.DB.Model(&model.User{}).Where("id = ?", req.ID).Updates(req.Data()).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": lang.Trans("user updated"),
	})
}
