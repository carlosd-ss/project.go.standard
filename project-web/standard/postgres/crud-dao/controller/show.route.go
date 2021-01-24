// Go Api server
// @jeffotoni

package controller

import (
	cf "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/config"
	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/fmts"
)

//Show log
func Show(cfg cf.Config) {
	APIURL := cf.HOSTSERVER
	fmts.Println("\033[0;33m  Api Server Run \033[0m \033[0;32mstarted:\033[0m \033[37mHost: ", APIURL, " port: ", cf.PORTSERVER, "\033[0m")
	fmts.Println("\033[37m  EndPoints\033[0m")
	fmts.Println("\033[0;33m  [POST]\033[0m.........-> {public}............... -> \033[0;36m", Endpoint().Ping)
	fmts.Println("\033[0;33m  [POST]\033[0m.........-> {public}............... -> \033[0;36m", Endpoint().User)
	fmts.Println("\033[0;33m  [GET]\033[0m.........-> {private key}............... -> \033[0;36m", Endpoint().User, "/{uuid}")
	fmts.Println("\033[0;33m  [PUT]\033[0m.........-> {private key}............... -> \033[0;36m", Endpoint().User, "/{uuid}")
	fmts.Println("\033[0;33m  [DELETE]\033[0m.......-> {private key}............ -> \033[0;36m", Endpoint().User, "/{uuid}")
	fmts.Println("\033[0m")

}
