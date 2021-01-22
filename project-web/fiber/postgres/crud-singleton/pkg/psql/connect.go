// Back-End in Go server
// @jeffotoni
package psql

import (
	"database/sql"
	"log"
	"sync"

	"github.com/project.go.standard/project-web/fiber/postgres/crud-singleton/internal/fmts"
	_ "github.com/lib/pq"
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
		if len(DB_CERT_SERVER_CA_PATH) > 0 &&
			len(DB_CERT_CLIENT_PATH) > 0 &&
			len(DB_CERT_KEY_PATH) > 0 {
			DBINFO = fmts.Concat("host=", DB_HOST, " port=",
				DB_PORT, " user=", DB_USER,
				" password=", DB_PASSWORD, " dbname=", DB_NAME,
				" sslmode=", DB_SSL, " sslrootcert=", DB_CERT_SERVER_CA_PATH,
				" sslcert=", DB_CERT_CLIENT_PATH, " sslkey=", DB_CERT_KEY_PATH,
			)
		} else {
			DB_SSL = "require"
			DBINFO = fmts.Concat("host=", DB_HOST, " port=", DB_PORT, " user=", DB_USER,
				" password=", DB_PASSWORD, " dbname=", DB_NAME,
				" sslmode=", DB_SSL)
		}
		if dbLocal, err = sql.Open(DB_SORCE, DBINFO); err != nil {
			log.Println("error open:", err)
			return
		}
	})
	return dbLocal
}
