package rcustomer

import (
	"database/sql"
	"encoding/json"

	mcustomer "github.com/go.standard.project.layout/project-web/fiber/crud-dao/models/customer"
)

func GetUuid(db *sql.DB, uuid string) (string, error) {

	var customer mcustomer.CustomerPost
	sqlexec := `
	SELECT 
		imp_id,
		imp_uuid,
		imp_status,
		imp_nome
	FROM ad_customer WHERE imp_uuid=$1 AND imp_status =$2`
	row := db.QueryRow(sqlexec, uuid, 1)
	//aju_status = 1 ativo, = 2 desativado
	err := row.Scan(
		&customer.ImpId,
		&customer.ImpUuid,
		&customer.ImpStatus,
		&customer.ImpNome,
	)
	if err != nil {
		return "", err
	}
	json, err := json.Marshal(customer)
	return string(json), nil
}
