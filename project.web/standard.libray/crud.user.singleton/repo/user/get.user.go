package user

import (
	pg "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/psql"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/zerolog"
)

//List ..
func List(id string) (userjson string) {
	Db := pg.Connect()
	row := Db.QueryRow(`SELECT * FROM users WHERE id=$1`, id)
	err := row.Scan(&userjson)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"user.go",
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
