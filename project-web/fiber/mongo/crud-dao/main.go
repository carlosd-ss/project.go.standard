package main

import (
	"log"

	hcustomer "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler/customer"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber"
	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/route"
	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/fmts"
)

func main() {
	s := hcustomer.Server{
		Db: mongo.Connect(),
	}
	defer s.Db.Close()
	app := fiber.New()
	route.AllRoutes(&s, app)
	err := app.Listen(8181)
	if err != nil {
		log.Println(err)
		return
	}
	fmts.Println("Server Started!")
}
