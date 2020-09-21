// Go Api server
// @jeffotoni

package cf

import (
	"math/rand"
	"os"
	"time"

	"github.com/jeffotoni/project.go.standard/project.web/standard.libray/crud.user.singleton/pkg/fmts"
	_ "github.com/lib/pq"
)

var (
	PORT_SERVER = "8181"
	HOST_SERVER = "0.0.0.0"
	HOST_CONFIG = fmts.Concat(HOST_SERVER, ":", PORT_SERVER)
)

const LayoutDateLog = "2006-01-02 15:04:05"
const LayoutDate = "2006-01-02"
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
	NAME_APP     = "crud.user.singleton"
	PROJECT_SITE = ""
)

const (

	// a sleep for goroutine
	// withstand a very large
	// load of requests
	MaxClients = 200000
)

// System
// go build -ldflags '-X main.DEBUG=YES'
// go build -ldflags '-X main.DEBUG=YES -x main.ENV=PROD'
var (
	DEBUG = "NO"
)

var (
	HASHTOKEN           = "1234567890"
	AUTHORIZATION_REGEX = "Basic %s:%s"

	HOST_SCHEME_HTTP  = "http://"
	HOST_SCHEME_HTTPS = "https://"

	// 10,000 requests per second
	REQUEST_SEC  = 50000.00
	HOST_MAXBYTE = 1 << 26 // 64MB
	PORT_METRICS = "6010"

	HOST_SERVER_METRICS   = ""
	API_HOST_CRYPT        = ""
	API_REMOTE_SERVER_ONE = ""

	// Authorization Basic
	X_KEY_1 = "xxxxxx" // base64->
	X_KEY_2 = "xxxxxx" // base64->

	AUTHORIZATION = fmts.Concat(AUTHORIZATION_REGEX, X_KEY_1, X_KEY_2)
	READTIMEOUT   = time.Duration(rand.Intn(1000-700)+700) * time.Millisecond
	WRITETIMEOUT  = time.Duration(rand.Intn(1400-900)+900) * time.Millisecond

	WD_LEVEL = 0 // Path: Getwd - dir absolute
)

/////// DATA BASE
var (
	DB_NAME = "user"
	// DB_HOST     = ""
	DB_HOST     = "localhost" // container
	DB_USER     = "postgres"
	DB_PASSWORD = "20222022"
	DB_PORT     = "5432"
	DB_SSL      = "disable"
	DB_SORCE    = "postgres"

	/////////server
	// HOST_CONFIG = HOST_SERVER + ":" + PORT_SERVER
	/////////////

	///api
	API_ENV  = ""
	API_HOST = ""
	ENV_AMBI = os.Getenv("ENV_AMBI")
)

// Config provides
// basic configuration
type Config struct {
	Host           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func init() {
	if len(ENV_AMBI) <= 0 {
		// ACESSO REMOTO
		DB_NAME = "user"
		DB_HOST = "localhost"
	} else if ENV_AMBI == "LOCAL" {
		DB_NAME = "user"
		DB_HOST = "localhost"
	} else if ENV_AMBI == "EC2_DB" {
		DB_NAME = ""
		DB_HOST = ""
	} else if ENV_AMBI == "PROD" {
		DB_NAME = ""
		DB_HOST = ""
	}

	//println("Connect ambiente: ", ENV_AMBI)
	//println("Database ", DB_NAME)
}

func TypeEnv() string {
	return ENV_AMBI
}
