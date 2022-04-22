package handler

import (
	"net/http"

	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/errors"
	"github.com/Binaretech/classroom-main/lang"
	"github.com/Binaretech/classroom-main/validation"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// User handler look up the user's data in the database and return it or returns 404 if not found
func User(c echo.Context) error {
	userID := c.Request().Header.Get("X-User")

	user := model.User{}
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

// StoreUser create a new user in the database with the request data
func StoreUser(c echo.Context) error {
	req := model.StoreUserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	req.ID = c.Request().Header.Get("X-User")

	if err := validation.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	user := req.User()

	if err := db.Query().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		file, _ := c.FormFile("image")

		if err := user.UpdateProfileImage(tx, file, user.ID); err != nil {
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
func UpdateUser(c echo.Context) error {
	req := model.UpdateUserRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	req.ID = c.Request().Header.Get("X-User")

	if err := validation.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := db.Model(&model.User{}).Where("id = ?", req.ID).Updates(req.Data()).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": lang.Trans("user updated"),
	})
}
