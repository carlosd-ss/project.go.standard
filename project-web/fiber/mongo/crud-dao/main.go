package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/route"
	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/fmts"
)

func main() {
	app := fiber.New()
	route.AllRoutes(app)
	err := app.Listen(8181)
	if err != nil {
		log.Println(err)
		return
	}
	fmts.Println("Server Started!")
}
