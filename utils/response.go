package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": message,
	})
}
