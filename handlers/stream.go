package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func StartStream(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

func StopStream(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}

func GetStream(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
	})
}
