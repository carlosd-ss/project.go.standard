package user

import (
	"database/sql"

	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/zerolog"
)

//List ..
func List(db *sql.DB, rid string) (name, lastname, id string) {

	row := db.QueryRow(`SELECT * FROM users WHERE id=$1`, rid)
	err := row.Scan(&name, &lastname, &id)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"user.go",
			13,
			"api.crud-dao.com.br",
			"Repo List user",
			err.Error())
		return "", "", ""
	}
	if len(name) > 0 {
		return name, lastname, id
	}
	return "", "", ""
}
