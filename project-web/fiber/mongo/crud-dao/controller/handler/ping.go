package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Ping returns a pong string through fiber context
func (c Conn) Ping(c *fiber.Ctx) {
	c.Status(200).SendString("pong")
	return
}
