package main

import (
	"log"

	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"

	"github.com/gofiber/fiber"
	route "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler"
	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/fmts"
)

func main() {
	s := route.Conn{
		Client:         mongoconf.Connection(),
		Db:             "Banco1",
		CollectionName: "Colection1",
	}
	//defer s.Client.Cancel()
	app := fiber.New()
	route.AllRoutes(&s, app)
	err := app.Listen(8181)
	if err != nil {
		log.Println(err)
		return
	}
	fmts.Println("Server Started!")
}
