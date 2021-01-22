package main

import (
	"github.com/project.go.standard/project-web/fiber/crud.mongoa/controller/route"
	"github.com/project.go.standard/project-web/fiber/crud.mongoa/internal/fmts"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	route.AllRoutes(app)

	app.Listen(8181)

	fmts.Println("Server Started!")
}
