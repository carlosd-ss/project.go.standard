package user

import (
	mUser "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/model/user"
	pg "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/psql"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/zerolog"
)

//UpdateUser ..
func UpdateUser(id string, user mUser.User) error {
	//Atualiza todos os valores
	sqlStatement := `UPDATE INTO users (first_name, last_name) VALUES ($1, $2)`
	Db := pg.Connect()
	_, err := Db.Exec(sqlStatement, id, user.Name, user.Lasname)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"put.go",
			19,
			"api.crud.user.singleton.com.br",
			"Repo Put user",
			err.Error())
		return err
	}
	return nil
}
