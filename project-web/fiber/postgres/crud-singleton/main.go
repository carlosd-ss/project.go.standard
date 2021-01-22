package main

import (
	"github.com/go.standard.project.layout/project.web/fiber/postgres/crud-singleton/controller/route"
	"github.com/go.standard.project.layout/project.web/fiber/postgres/crud-singleton/internal/fmts"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	route.AllRoutes(app)

	app.Listen(8181)

	fmts.Println("Server Started!")
}
