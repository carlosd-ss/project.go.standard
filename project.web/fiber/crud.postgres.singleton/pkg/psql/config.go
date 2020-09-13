package psql

import (
	"os"
)

var (
	DB_NAME     = os.Getenv("DB_NAME")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_PORT     = "5432"
	DB_SSL      = "verify-ca"
	DB_SORCE    = "postgres"
	//DB_CERT_PATH = ""
	DB_CERT_SERVER_CA_PATH = "rds-certs/server-ca.pem"
	DB_CERT_CLIENT_PATH    = "rds-certs/client-cert.pem"
	DB_CERT_KEY_PATH       = "rds-certs/client-key.pem"
	//DB_CERT_CLIENT_PATH = "./rds-certs/client-cert.pem"

	/////////server
	// HOST_CONFIG = HOST_SERVER + ":" + PORT_SERVER
	/////////////
	///api
	API_ENV  = ""
	API_HOST = ""
	ENV_AMBI = os.Getenv("ENV_AMBI")
)
