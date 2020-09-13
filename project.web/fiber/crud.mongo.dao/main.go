package main

import (
	"github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/controller/route"
	"github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/pkg/fmts"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	route.AllRoutes(app)

	app.Listen(8181)

	fmts.Println("Server Started!")
}
