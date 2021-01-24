package user

import (
	"database/sql"

	"github.com/jeffotoni/project.go.standard/project-web/standard/postgres/crud-dao/internal/zerolog"
)

//Delete ..
func Delete(db *sql.DB, id string) error {
	dell := `DELETE FROM users WHERE id=$1`
	_, err := db.Exec(dell, id)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"delete.user.go",
			17,
			"localhost",
			"Repo Delete",
			err.Error())
		return err
	}
	return nil
}
