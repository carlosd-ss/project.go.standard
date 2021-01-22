package main

import (
	"log"

	hcustomer "github.com/go.standard.project.layout/project-web/fiber/crud-dao/controller/handler/customer"
	"github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/pkg/psql"

	"github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/controller/route"
	"github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/pkg/fmts"

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
