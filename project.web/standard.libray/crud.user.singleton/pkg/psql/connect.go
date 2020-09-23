package psql

import (
	"database/sql"
	"log"
	"sync"

	cf "github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/config"
	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/fmts"
	_ "github.com/lib/pq"
)

var (
	once    sync.Once
	err     error
	dbLocal *sql.DB
)

//Connect ..
func Connect() *sql.DB {
	once.Do(func() {
		if dbLocal != nil {
			return
		}
		print(cf.DBNAME)
		var DBINFO string
		DBINFO = fmts.Concat(" host=", cf.DBHOST, " port=",
			cf.DBPORT, " user=", cf.DBUSER,
			" password=", cf.DBPASSWORD, " dbname=", cf.DBNAME,
			" sslmode=", cf.DBSSL)
		if dbLocal, err = sql.Open(cf.DBSORCE, DBINFO); err != nil {
			log.Println("error open:", err)
			return
		}
	})
	return dbLocal
}
