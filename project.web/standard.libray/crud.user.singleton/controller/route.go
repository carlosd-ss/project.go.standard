package controller

// Go Api server
// @jeffotoni
// @ancogamer

import (
	"net/http"
	"time"

	cf "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/controller/handler"
	mw "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/controller/middleware"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/middleware/stdlib"
	"github.com/ulule/limiter/drivers/store/memory"
	//"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/cors"
)

//Routes
func Routes(cfg cf.Config) *GoServerHttp {

	///////////////////////////////////////
	/////

	// DefaultServeMux
	mux := http.NewServeMux()
	//corsx := cors.Domain()
	// POST handler /ping

	mux.Handle(Endpoint().Ping, mw.Use(http.HandlerFunc(handler.Ping),
		mw.CustomHeaders(),
		mw.Gzip(),
		mw.MaxClients(cf.MaxClients),
		mw.Logger("/ping"),
	))

	///////////////////////////////////////////////////////////
	// Endpoints
	//////////////////////////////////////
	// user
	// user/{uuid}
	/////////////////////////////////////
	mux.Handle("/user", mw.Use(http.HandlerFunc(handler.UserPost),
		mw.CustomHeaders(),
		mw.Logger("/user")))

	mux.Handle("/", mw.Use(http.HandlerFunc(handler.HomeHandler),
		mw.CustomHeaders(),
		mw.Logger("/")))
	//handlerCors := corsx.Handler(mux)
	APIServer := GoServerHttp{
		server: &http.Server{
			Addr: cfg.Host,
			Handler: stdlib.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{Formatted: "",
				Period: 1 * time.Second, Limit: 15})).Handler(mux),
			ReadTimeout:  time.Millisecond * 600,
			WriteTimeout: time.Millisecond * 400,
		},
	}

	// Add to the WaitGroup for the listener goroutine
	APIServer.wg.Add(1)

	Show(cfg)
	APIServer.server.ListenAndServe()
	APIServer.wg.Done()
	return &APIServer
}
