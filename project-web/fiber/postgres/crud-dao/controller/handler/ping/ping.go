package hping

import (
	"github.com/gofiber/fiber"
)

// Ping returns a pong string through fiber context
func Ping(c *fiber.Ctx) {
	c.Send("Pong")
}
