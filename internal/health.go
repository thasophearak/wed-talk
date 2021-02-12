package internal

import "github.com/gofiber/fiber/v2"

// Health check func
func Health(c *fiber.Ctx) error {
	return c.SendString("alive")
}
