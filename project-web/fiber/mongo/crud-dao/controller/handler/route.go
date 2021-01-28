package handler

import (
	"github.com/gofiber/cors"
	mw "github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"

	//handlerPing "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler/ping"
	conn "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler"
	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"
	//handlerPing "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler/ping"
	//handlerPing "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/controller/handler/ping"
)

//
func AllRoutes(app *fiber.App) {

	s := conn.Conn{
		Client1:      mongoconf.Connect("banco1"),
		Client2:      mongoconf.Connect("banco2"),
		Client3:      mongoconf.Connect("banco3"),
		Collection1:  mongoconf.GetCollection(Client1, "banco1"),
		Collection2:  mongoconf.GetCollection(Client2, "banco2"),
		Collection3:  mongoconf.GetCollection(Client3, "banco3"),
		PingTimeout:  mongoconf.PingTimeout(),
		ReadTimeout:  mongoconf.ReadTimeout(),
		WriteTimeout: mongoconf.WriteTimeout(),
	}
	////
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
	//app.Get("/v1/customer/:uuid", handlerImp.GetUuid)
	app.Delete("/v1/customer/:uuid", s.Delete)
	app.Put("/v1/customer", s.Update)
}
