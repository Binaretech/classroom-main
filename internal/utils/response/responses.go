package response

import "github.com/gofiber/fiber/v2"

func ResponseBadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusBadRequest).JSON(map[string]string{
		"message": msg,
	})
}

func ResponseMessage(c *fiber.Ctx, status int, content interface{}) error {
	return c.Status(status).JSON(map[string]interface{}{
		"message": content,
	})
}
