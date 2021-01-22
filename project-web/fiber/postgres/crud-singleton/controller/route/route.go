package route

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	mw "github.com/gofiber/fiber/middleware"

	handlerImp "github.com/go.standard.project.layout/project-web/fiber/postgres/crud-singleton/controller/handler/customer"
	handlerPing "github.com/go.standard.project.layout/project-web/fiber/postgres/crud-singleton/controller/handler/ping"
)

//
func AllRoutes(app *fiber.App) {
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
	app.Post("/v1/customer", handlerImp.Post)
	app.Get("/v1/customer/:offset/:limit", handlerImp.GetAll)
	app.Get("/v1/customer/:uuid", handlerImp.GetUuid)
	app.Delete("/v1/customer/:uuid", handlerImp.Delete)
	app.Put("/v1/customer", handlerImp.Update)
}
