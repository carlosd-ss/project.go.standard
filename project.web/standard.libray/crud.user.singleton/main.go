// Go Api server
// @jeffotoni

package main

import (
	"os"
	"os/signal"

	cfp "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user/api.user/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user/api.user/controller"
	cf "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user/pkg/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user/pkg/pkg/util"
)

//var confserv = cfp.Endpoint()
func main() {
	serverCfg := cf.Config{
		Host:           cfp.HOST_CONFIG,
		ReadTimeout:    cf.READTIMEOUT,
		WriteTimeout:   cf.WRITETIMEOUT,
		MaxHeaderBytes: cf.HOST_MAXBYTE,
	}

	// star server
	hServer := controller.StartServer(serverCfg)

	// stop server
	defer hServer.StopServer()

	// channel with notify
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	util.Print("\nmain : Shutting down goapiserver\n")
}
