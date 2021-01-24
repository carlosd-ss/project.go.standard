package mongo

import (
	"context"
	"sync"
	"time"

	// "internal/pkg/fmts"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Schema            = "MONGODB_SCHEMA"
	Uri               = "MONGODB_URI"
	Username          = "MONGODB_USERNAME"
	Password          = "MONGODB_PASSWORD"
	databaseName      = "MONGODB_DATABASE"
	collectionName    = "MONGODB_COLLECTION"
	pingTimeout       = "MONGODB_PING_TIMEOUT"
	readTimeout       = "MONGODB_READ_TIMEOUT"
	writeTimeout      = "MONGODB_WRITE_TIMEOUT"
	Options           = "MONGODB_OPTIONS"
	connectionTimeout = "MONGODB_CONNECTION_TIMEOUT_IN_SECONDS"
)

var (
	client               *mongo.Client
	ctx                  context.Context
	connectOnce          sync.Once
	connectOnce2         sync.Once
	CollectionOnce       sync.Once
	session              *mongo.Client
	collection           *mongo.Collection
	ConnectionTimeout    = time.Duration(connectionTimeout, 2*time.Minute)
	pingTimeoutDuration  = time.Duration(pingTimeout, 30*time.Second)
	readTimeoutDuration  = time.Duration(readTimeout, 30*time.Second)
	writeTimeoutDuration = time.Duration(writeTimeout, 30*time.Second)
	err                  error
)

var (
	mgoSchema      = "mongodb"
	mgoUri         = "localhost:27017"
	mgoUser        = "root"
	mgoPass        = "senhahere"
	MgoDb          = "databaseHere"
	mgoRetry       = "authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	CollectionName = "Clientes"
)

func connectDefault() (connectionString string) {
	if len(mgoUser) > 0 && len(mgoPass) > 0 {
		connectionString = fmts.Concat(mgoSchema, "://", mgoUser, ":", mgoPass, "@", mgoUri, "/", MgoDb, "?", mgoRetry)
	} else {
		connectionString = fmts.Concat(mgoSchema, "://", mgoUri, "/", MgoDb, "?", mgoRetry)
	}
	return
}

func connectClient() *mongo.Client {
	connectOnce.Do(func() {
		if session == nil {
			session, err = mongo.NewClient(
				options.Client().
					ApplyURI(connectDefault()))
			if err != nil {
				logone.Error(
					"10000",
					"enginer",
					"master",
					"global.env",
					"0001",
					"localhost",
					"Clientes",
					"100001", "", "",
					"admin", "127.0.0.1",
					"mongo.NewClient",
					`options.Client()`,
					"msg,err", err.Error())
			}
		}
	})
	return session
}

func Connect() *mongo.Client {
	connectOnce2.Do(func() {
		if client == nil {
			client = connectClient()
			ctx, _ = context.WithTimeout(context.Background(), ConnectionTimeout)
			err := client.Connect(ctx)
			if err != nil {
				logone.Error(
					"10000",
					"enginer",
					"master",
					"global.env",
					"0001",
					"localhost",
					"Clientes",
					"100001", "", "",
					"admin", "127.0.0.1",
					"client.Connect(ctx)",
					`ctx`,
					"msg,err", err.Error())
				return
			}

		}
	})
	return client
}

func GetCollectionClientes() *mongo.Collection {
	//CollectionOnce.Do(func() {
	//	if collection == nil {

	//	}
	//})
	collection = client.Database(MgoDb).Collection(CollectionName)
	return collection
}
