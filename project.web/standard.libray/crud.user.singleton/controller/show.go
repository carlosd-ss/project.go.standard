// Go Api server
// @jeffotoni

package controller

import (
	cfp "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/api.user/config"
	cf "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/pkg/util"
)

func Show(cfg cf.Config) {
	API_URL := cfp.HOST_SERVER
	if cf.TypeEnv() != "PROD" {
		util.Printnew("\033[0;33m  Api Server Run \033[0m \033[0;32mstarted:\033[0m \033[37mHost: ", API_URL, " port: ", cfp.PORT_SERVER, "\033[0m")
		util.Printnew("\033[37m  EndPoints\033[0m")
		util.Printnew("\033[0;33m  [POST]\033[0m.........-> {public}............... -> \033[0;36m", Endpoint().Ping)
		util.Printnew("\033[0;33m  [DELETE]\033[0m.......-> {private key}............ -> \033[0;36m", Endpoint().User, "/{uuid}")
		util.Printnew("\033[0m")
	}
}
