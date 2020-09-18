// Go Api server
// @jeffotoni

package controller

import (
	cf "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/config"
	cfp "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/fmts"
)

func Show(cfg cf.Config) {
	API_URL := cfp.HOST_SERVER
	if cf.TypeEnv() != "PROD" {
		fmts.Println("\033[0;33m  Api Server Run \033[0m \033[0;32mstarted:\033[0m \033[37mHost: ", API_URL, " port: ", cfp.PORT_SERVER, "\033[0m")
		fmts.Println("\033[37m  EndPoints\033[0m")
		fmts.Println("\033[0;33m  [POST]\033[0m.........-> {public}............... -> \033[0;36m", Endpoint().Ping)
		fmts.Println("\033[0;33m  [DELETE]\033[0m.......-> {private key}............ -> \033[0;36m", Endpoint().User, "/{uuid}")
		fmts.Println("\033[0m")
	}
}
