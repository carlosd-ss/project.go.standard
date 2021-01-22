package route

import (
	hcustomer "github.com/go.standard.project.layout/project-web/fiber/crud-dao/controller/handler/customer"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	mw "github.com/gofiber/fiber/middleware"

	handlerPing "github.com/go.standard.project.layout/project-web/fiber/crud-dao/controller/handler/ping"
)

//
func AllRoutes(s *hcustomer.Server, app *fiber.App) {
	////
	app.Use(cors.New())
	app.Use(mw.Compress(mw.CompressLevelBestSpeed))
	app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

	///
	app.Get("/ping", handlerPing.Ping)

	// application/json
	// application/x-www-form-urlencoded
	// multipart/form-data

	//customer (customer) 	// application/json
	app.Post("/v1/customer", s.Post)
	app.Get("/v1/customer/:offset/:limit", s.GetAll)
	app.Get("/v1/customer/:uuid", s.GetUuid)
	app.Delete("/v1/customer/:uuid", s.Delete)
	app.Put("/v1/customer", s.Update)
}
