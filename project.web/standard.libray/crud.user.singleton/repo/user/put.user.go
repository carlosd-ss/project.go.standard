package user

import (
	mUser "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/model/user"
	pg "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/psql"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/zerolog"
)

//UpdateUser ..
func UpdateUser(id string, user mUser.User) error {
	//Atualiza todos os valores
	sqlStatement := `UPDATE users SET first_name=$1, last_name=$2 WHERE id = $3`
	Db := pg.Connect()
	_, err := Db.Exec(sqlStatement, user.Name, user.Lasname, id)
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
