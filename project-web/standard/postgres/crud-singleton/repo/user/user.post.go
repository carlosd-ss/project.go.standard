package user

import (
	"github.com/jeffotoni/project.go.standard/project-web/standard.libray/crud.user.singleton/internal/zerolog"
	pg "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/psql"
	mUser "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/model/user"
)

//InsertUser ..
func InsertUser(user mUser.User) error {
	sqlStatement := `INSERT INTO users(first_name, last_name) VALUES($1,$2)`
	Db := pg.Connect()
	_, err := Db.Exec(sqlStatement, user.Name, user.Lasname)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"post.go",
			18,
			"api.crud-dao.com.br",
			"Repo Post user",
			err.Error())
		return err
	}
	return nil
}
