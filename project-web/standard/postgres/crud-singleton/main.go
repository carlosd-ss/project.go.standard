// Go Api server
// @jeffotoni

package main

import (
	"os"
	"os/signal"

	cf "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-singleton/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/controller"
	cfp "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-singleton/controller"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/internal/fmts"
)

var confserv = cfp.Endpoint()

func main() {
	serverCfg := cf.Config{
		Host:           cf.HOSTCONFIG,
		ReadTimeout:    cf.READTIMEOUT,
		WriteTimeout:   cf.WRITETIMEOUT,
		MaxHeaderBytes: cf.HOSTMAXBYTE,
	}

	//star server
	hServer := controller.Routes(serverCfg)

	// stop server
	defer hServer.StopServer()

	//channel with notify
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmts.Print("\nmain : Shutting down goapiserver\n")
}
