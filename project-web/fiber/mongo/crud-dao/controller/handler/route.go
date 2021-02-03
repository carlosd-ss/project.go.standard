package handler

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	mw "github.com/gofiber/fiber/middleware"
	//handlerPing "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler/ping"
	//handlerPing "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler/ping"
	//handlerPing "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler/ping"
)

//
func AllRoutes(s *Conn, app *fiber.App) {

	app.Use(cors.New())
	app.Use(mw.Compress(mw.CompressLevelBestSpeed))
	app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))

	///
	app.Get("/ping", s.Ping)

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
