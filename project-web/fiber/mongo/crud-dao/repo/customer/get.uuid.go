package rcustomer

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"

	mcustomer "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/models/customer"
)

func GetUuid(db *mongo.Client, uuid string) (string, error) {
	var customer mcustomer.Customer
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
