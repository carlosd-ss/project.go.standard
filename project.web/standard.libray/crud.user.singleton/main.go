// Go Api server
// @jeffotoni

package main

import (
	cf "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/controller"
	cfp "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/controller"
)

var confserv = cfp.Endpoint()

func main() {
	serverCfg := cf.Config{
		Host:           cf.HOST_CONFIG,
		ReadTimeout:    cf.READTIMEOUT,
		WriteTimeout:   cf.WRITETIMEOUT,
		MaxHeaderBytes: cf.HOST_MAXBYTE,
	}

	//star server
	hServer := controller.StartServer(serverCfg)

	// stop server
	defer hServer.StopServer()

	// channel with notify
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, os.Interrupt)
	//<-sigChan

	//util.Print("\nmain : Shutting down goapiserver\n")
}
