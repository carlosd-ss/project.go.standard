// Go Api server
// @jeffotoni

package main

import (
	"os"
	"os/signal"

	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/controller/handler"
	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/psql"

	cf "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/config"
	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/controller"
	cfp "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/controller"
	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/fmts"
)

var confserv = cfp.Endpoint()

func main() {
	serverCfg := cf.Config{
		Host:           cf.HOSTCONFIG,
		ReadTimeout:    cf.READTIMEOUT,
		WriteTimeout:   cf.WRITETIMEOUT,
		MaxHeaderBytes: cf.HOSTMAXBYTE,
	}
	s := handler.Server{
		Db: psql.Connect(),
	}
	defer s.Db.Close()
	//star server
	hServer := controller.Routes(&s, serverCfg)

	// stop server
	defer hServer.StopServer()

	//channel with notify
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmts.Print("\nmain : Shutting down goapiserver\n")
}
