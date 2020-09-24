package cf

// Go Api server
// @jeffotoni
// @ancogamer

import (
	"math/rand"
	"os"
	"time"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/internal/fmts"
)

var (
	//PORTSERVER ..
	PORTSERVER = "8181"
	//HOSTSERVER ..
	HOSTSERVER = "0.0.0.0"
	//HOSTCONFIG ..
	HOSTCONFIG = fmts.Concat(HOSTSERVER, ":", PORTSERVER)
)

//LayoutDateLog ..
const LayoutDateLog = "2006-01-02 15:04:05"

//LayoutDate ..
const LayoutDate = "2006-01-02"

//LayoutHour ..
const LayoutHour = "15:04:05"

type ByteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	//ZB
	//YB
)

const (
	// a sleep for goroutine
	// withstand a very large
	// load of requests

	//MaxClients ...
	MaxClients = 200000
)

// System
// go build -ldflags '-X main.DEBUG=YES'
// go build -ldflags '-X main.DEBUG=YES -x main.ENV=PROD'
var (
	DEBUG = "NO"
)

var (
	//HASHTOKEN ..
	HASHTOKEN = "1234567890"
	//AUTHORIZATIONRegex ..
	AUTHORIZATIONRegex = "Basic %s:%s"
	//HostSchemeHTTP ..
	HostSchemeHTTP = "http://"
	//HostSchemeHTTPS ..
	HostSchemeHTTPS = "https://"

	//REQUESTSEC ..
	REQUESTSEC = 50000.00
	//HOSTMAXBYTE ..
	HOSTMAXBYTE = 1 << 26 // 64MB
	//PORTMETRICS ..
	PORTMETRICS = "6010"
	//HOSTSERVERMetrics ..
	HOSTSERVERMetrics = ""
	//APIHostCrypt ..
	APIHostCrypt = ""
	//APIRemoteServerONE ..
	APIRemoteServerONE = ""

	//Authorization Basic

	//Xkey1 ..
	Xkey1 = "xxxxxx" // base64->
	//Xkey2 ..
	Xkey2 = "xxxxxx" // base64->
	//AUTHORIZATION ..
	AUTHORIZATION = fmts.Concat(AUTHORIZATIONRegex, Xkey1, Xkey2)
	//READTIMEOUT ..
	READTIMEOUT = time.Duration(rand.Intn(1000-700)+700) * time.Millisecond
	//WRITETIMEOUT ..
	WRITETIMEOUT = time.Duration(rand.Intn(1400-900)+900) * time.Millisecond
	//WDlevel ..
	WDlevel = 0 // Path: Getwd - dir absolute
)

/////// DATA BASE
var (
	DBNAME     = os.Getenv("DB_NAME")
	DBHOST     = os.Getenv("DB_HOST") // container
	DBUSER     = os.Getenv("DB_USER")
	DBPASSWORD = os.Getenv("DB_PASSWORD")
	DBPORT     = os.Getenv("DB_PORT")
	DBSSL      = os.Getenv("DB_SSL")
	DBSORCE    = os.Getenv("DB_SORCE")

	/////////server
	// HOST_CONFIG = HOST_SERVER + ":" + PORT_SERVER
	/////////////

	/////api
	//APIENV  = os.Getenv("APIENV")
	//APIHOST = ""
	//ENVAMBI = os.Getenv("ENV_AMBI")
)

// Config provides
// basic configuration
type Config struct {
	Host           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

/*func init() {
	if len(ENVAMBI) <= 0 {
		// ACESSO REMOTO
		DBNAME = "user"
		DBHOST = "localhost"
	} else if ENV_AMBI == "LOCAL" {
		DBNAME = "user"
		DBHOST = "localhost"
	} else if ENV_AMBI == "EC2_DB" {
		DBNAME = ""
		DBHOST = ""
	} else if ENV_AMBI == "PROD" {
		DBNAME = ""
		DBHOST = ""
	}

	//println("Connect ambiente: ", ENV_AMBI)
	//println("Database ", DB_NAME)
//}

//func TypeEnv() string {
//	return ENV_AMBI
//}
*/
