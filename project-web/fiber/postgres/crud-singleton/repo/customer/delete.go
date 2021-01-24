package rcustomer

import (
	"errors"

	db "github.com/project.go.standard/project-web/fiber/postgres/crud-dao/internal/psql"
)

func Delete(uuid string) error {
	db := db.Connect()
	//Those ones that are not here on the insert, are default fields
	sqlexec := `DELETE from ad_customer WHERE  imp_uuid = $1`
	res, err := db.Exec(sqlexec, uuid)
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		return errors.New("nenhuma coluna afetada")
	}
	return nil

}
