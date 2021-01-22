package main

import (
	"log"

	hcustomer "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/controller/handler/customer"
	"github.com/project.go.standard/project-web/fiber/postgres/crud-dao/pkg/psql"

	"github.com/project.go.standard/project-web/fiber/postgres/crud-dao/controller/route"
	"github.com/project.go.standard/project-web/fiber/postgres/crud-dao/pkg/fmts"

	"github.com/gofiber/fiber"
)

func main() {
	s := hcustomer.Server{
		Db: psql.Connect(),
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
