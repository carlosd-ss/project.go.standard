package user

import (
	"database/sql"

	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/zerolog"
	mUser "github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/model/user"
)

//InsertUser ..
func InsertUser(db *sql.DB, user mUser.User) error {
	sqlStatement := `INSERT INTO users(first_name, last_name) VALUES($1,$2)`

	_, err := db.Exec(sqlStatement, user.Name, user.Lasname)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"customer.post.go",
			18,
			"api.crud-dao.com.br",
			"Repo Post user",
			err.Error())
		return err
	}
	return nil
}
