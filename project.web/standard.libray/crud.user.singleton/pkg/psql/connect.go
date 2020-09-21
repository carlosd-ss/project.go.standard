package psql

import (
	"database/sql"
	"log"
	"sync"

	cf "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/fmts"
)

var (
	once    sync.Once
	err     error
	dbLocal *sql.DB
)

func Connect() *sql.DB {
	once.Do(func() {
		if dbLocal != nil {
			return
		}
		var DBINFO string
		//if len(DB_CERT_SERVER_CA_PATH) > 0 &&
		//len(DB_CERT_CLIENT_PATH) > 0 &&
		//len(DB_CERT_KEY_PATH) > 0 {
		DBINFO = fmts.Concat(" host=", cf.DB_HOST, " port=",
			cf.DB_PORT, " user=", cf.DB_USER,
			" password=", cf.DB_PASSWORD, " dbname=", cf.DB_NAME,
			" sslmode=", cf.DB_SSL)
		//} else {
		//DB_SSL = "require"
		//DBINFO = fmts.Concat("host=", DB_HOST, " port=", DB_PORT, " user=", DB_USER,
		//" password=", DB_PASSWORD, " dbname=", DB_NAME,
		//" sslmode=", DB_SSL)
		//}

		if dbLocal, err = sql.Open(cf.DB_SORCE, DBINFO); err != nil {
			log.Println("error open:", err)
			return
		}
	})
	return dbLocal
}
