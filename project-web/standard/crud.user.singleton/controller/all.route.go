package controller

// Go Api server
// @jeffotoni
// @ancogamer

import (
	"net/http"
	"time"

	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/middleware/stdlib"
	"github.com/ulule/limiter/drivers/store/memory"

	cf "github.com/jeffotoni/project.go.standard/project-web/standard/crud.user.singleton/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/controller/handler"
	mw "github.com/jeffotoni/project.go.standard/project-web/standard/crud.user.singleton/controller/middleware"
	//"github.com/jeffotoni/project.go.standard/project-web/standard/crud.user.singleton/internal/cors"
)

//Routes
func Routes(cfg cf.Config) *GoServerHttp {

	///////////////////////////////////////
	/////

	// DefaultServeMux
	mux := http.NewServeMux()
	//corsx := cors.Domain()
	// POST handler /ping

	mux.Handle(Endpoint().Ping, mw.Use(stdlib.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{Formatted: "",
		Period: 1 * time.Second, Limit: 1})).Handler(http.HandlerFunc(handler.Ping)),
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
	mux.Handle("/user", mw.Use(stdlib.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{Formatted: "",
		Period: 1 * time.Second, Limit: 15})).Handler(http.HandlerFunc(handler.UserPost)),
		mw.CustomHeaders(),
		mw.Logger("/user")))
	// vale lembrar que alguns endpoints utilizam regex e tem sua criação dinamica, ou isto significa que este rate limit de 15 requests/s será atribuido a todos,
	// que nascem desta maneira
	mux.Handle("/", mw.Use(stdlib.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{Formatted: "",
		Period: 1 * time.Second, Limit: 15})).Handler(http.HandlerFunc(handler.HomeHandler)),
		mw.CustomHeaders(),
		mw.Logger("/")))

	APIServer := GoServerHttp{
		server: &http.Server{
			Addr:         cfg.Host,
			Handler:      mux,
			ReadTimeout:  time.Millisecond * 600,
			WriteTimeout: time.Millisecond * 400,
		},
	}
	//handlerCors := corsx.Handler(mux)
	// Isto ira atribuir o rate limit para todos os endpoints
	/*
		APIServer := GoServerHttp{
			server: &http.Server{
				Addr: cfg.Host,
				Handler: stdlib.NewMiddleware(limiter.New(memory.NewStore(), limiter.Rate{Formatted: "",
					Period: 1 * time.Second, Limit: 15})).Handler(mux),
				ReadTimeout:  time.Millisecond * 600,
				WriteTimeout: time.Millisecond * 400,
			},
		}
	*/
	// Add to the WaitGroup for the listener goroutine
	APIServer.wg.Add(1)

	Show(cfg)
	APIServer.server.ListenAndServe()
	APIServer.wg.Done()
	return &APIServer
}
