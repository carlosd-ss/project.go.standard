package rcustomer

import (
	"errors"

	mcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/models/customer"
	fmts "github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/pkg/fmts"
	db "github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/pkg/psql"
)

func Update(customer mcustomer.Customer) error {
	db := db.Connect()
	//Those ones that are not here on the insert, are default fields
	sqlexec := `UPDATE ad_customer SET imp_user_update = $1`

	if customer.ImpStatus != 0 && customer.ImpStatus == 1 || customer.ImpStatus == 2 {
		sqlexec = fmts.Concat(sqlexec, ", imp_status = ", customer.ImpStatus, "")
	} else {
		return errors.New("status invalido")
	}
	if customer.ImpNome != "" {
		sqlexec = fmts.Concat(sqlexec, ", imp_nome = '", customer.ImpNome, "'")
	}
	sqlexec = fmts.Concat(sqlexec, ", imp_ip = $2 WHERE imp_uuid = $3")
	res, err := db.Exec(sqlexec,
		customer.ImpUserUpdate,
		customer.ImpIp,
		customer.ImpUuid,
	)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmts.Println(err)
		return err
	}
	if count == 0 {
		fmts.Println("checar imp_uuid q está sendo enviado")
		return errors.New("nenhuma coluna afetada")
	}
	return nil
}
