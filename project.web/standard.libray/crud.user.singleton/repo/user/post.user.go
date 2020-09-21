package user

import (
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/model/usermodel"
)

func InsertUser(user usermodel.User) (int64, error) {
	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age)
	if err != nil {
		return id, err
	}
	return id, nil
}
