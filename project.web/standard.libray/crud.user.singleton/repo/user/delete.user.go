package user

import (
	pg "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/psql"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/zerolog"
)

func Delete(userUuid string) bool {
	Db := pg.Connect()
	dell := `DELETE FROM adphone WHERE uuiduser=$1`
	_, err := Db.Exec(dell, userUuid)
	if err != nil {
		zerolog.Error(
			"1.0.0",
			"delete.user.go",
			17,
			"localhost",
			"Delete",
			err.Error())
		return false
	}
	return true
}
