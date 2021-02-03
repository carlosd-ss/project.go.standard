package handler

import (
	"github.com/gofiber/fiber"
)

// Ping returns a pong string through fiber context
func (s *Conn) Ping(c *fiber.Ctx) {
	c.Status(200).SendString("pong")
	return
}
