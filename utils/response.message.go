package utils

import "github.com/gofiber/fiber/v2"

func HandleError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).SendString(err.Error())
}

func HandleNotContent(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNoContent).SendString(message)
}
