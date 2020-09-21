package user

import (
	pg "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/psql"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/zerolog"
)

func List(uuiduser string) (userjson string) {
	Db := pg.Connect()
	row := Db.QueryRow(`SELECT * FROM adphone WHERE uuiduser=$1`, uuiduser)
	err := row.Scan(&userjson)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"get.user.go",
			13,
			"api.crud.user.singleton.com.br",
			"List user",
			err.Error())
		return "{}"
	}
	if len(userjson) > 0 {
		return userjson
	}
	return "{}"
}
