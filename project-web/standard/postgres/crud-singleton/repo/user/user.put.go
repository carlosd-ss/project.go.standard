package user

import (
	"github.com/jeffotoni/project.go.standard/project-web/standard.libray/crud.user.singleton/internal/zerolog"
	pg "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/psql"
	mUser "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/model/user"
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
			"api.crud-dao.com.br",
			"Repo Put user",
			err.Error())
		return err
	}
	return nil
}
