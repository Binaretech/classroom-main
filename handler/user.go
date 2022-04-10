package handler

import (
	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/errors"
	"github.com/Binaretech/classroom-main/lang"
	"github.com/Binaretech/classroom-main/validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// User handler look up the user's data in the database and return it or returns 404 if not found
func User(c *fiber.Ctx) error {
	userID := c.Get("X-User")

	user := model.User{}
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		return err
	}

	if user.ID == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// StoreUser create a new user in the database with the request data
func StoreUser(c *fiber.Ctx) error {
	req := model.StoreUserRequest{}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	req.ID = c.Get("X-User")

	if err := validation.Struct(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

// UpdateUser update the user's data in the database with the request data
func UpdateUser(c *fiber.Ctx) error {
	req := model.UpdateUserRequest{}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	req.ID = c.Get("X-User")

	if err := validation.Struct(req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	if err := db.Model(&model.User{}).Where("id = ?", req.ID).Updates(req.Data()).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": lang.Trans("user updated"),
	})
}
