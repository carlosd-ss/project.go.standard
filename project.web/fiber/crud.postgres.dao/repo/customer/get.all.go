package rcustomer

import (
	"database/sql"
	"encoding/json"

	mcustomer "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/models/customer"
	fmts "github.com/go.standard.project.layout/project.web/fiber/crud.postgres.dao/pkg/fmts"
)

func GetAll(db *sql.DB, offset string, limit string) ([]string, error) {

	var customer mcustomer.CustomerPost
	sqlexec := `
	SELECT 
		imp_uuid,
		imp_nome,
		imp_status
	FROM ad_customer WHERE imp_status =$1 order by imp_id  offset $2 limit $3`
	rows, err := db.Query(sqlexec, 1, offset, limit)
	if err != nil {
		fmts.Println(err)
		return nil, err
	}
	var results []string
	//aju_status = 1 ativo, = 2 desativado
	for rows.Next() {
		err := rows.Scan(
			&customer.ImpUuid,
			&customer.ImpNome,
			&customer.ImpStatus,
		)
		if err != nil {
			return nil, err
		}
		json, err := json.Marshal(customer)
		results = append(results, string(json))
	}
	if err := rows.Err(); err != nil {
		fmts.Println(err)
		return nil, err
	}
	return results, nil

}
