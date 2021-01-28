package mongoconf

import (
	"context"
	"time"

	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/fmts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoRetry string = ""
var mgoSchema string = ""
var mgoUri string = ""

func connMongo(MgoDb, mgoUser, mgoPass string) (connectionString string) {
	if len(mgoUser) > 0 && len(mgoPass) > 0 {
		connectionString = fmts.Concat(mgoSchema, "://", mgoUser, ":", mgoPass, "@", mgoUri, "/", MgoDb, "?", mgoRetry)
	} else {
		connectionString = fmts.Concat(mgoSchema, "://", mgoUri, "/", MgoDb, "?", mgoRetry)
	}
	return
}

func connClient(MgoDb, mgoUser, mgoPass string) *mongo.Client {

	session, err = mongo.NewClient(
		options.Client().
			ApplyURI(connMongo(MgoDb, mgoUser, mgoPass)))
	if err != nil {
		// zerolog here..
	}

	return session
}

func ConnectMongo(MgoDb, mgoUser, mgoPass string) *mongo.Client {
	client = connClient(MgoDb, mgoUser, mgoPass)
	ctx, _ = context.WithTimeout(context.Background(), ConnectionTimeout)
	err := client.Connect(ctx)
	if err != nil {
		// zero log
		return
	}
	return client
}

func GetCollection(client *mongo.Client, MgoDb, CollectionName string) *mongo.Collection {
	CollectionOnce.Do(func() {
		if collection == nil {
			collection = client.Database(MgoDb).Collection(CollectionName)
		}
	})
	return collection
}

func PingTimeout() time.Duration {
	return time.Duration(time.Second * 5)
}

func ReadTimeout() time.Duration {
	return time.Duration(time.Second * 3)
}

func WriteTimeout() time.Duration {
	return time.Duration(time.Second * 3)
}

// //GetMongoDbConnection get connection of mongodb
// func GetMongoDbConnection() (*mongo.Client, error) {

// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.Background(), readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return client, nil
// }

// func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
// 	client, err := GetMongoDbConnection()

// 	if err != nil {
// 		return nil, err
// 	}

// 	collection := client.Database(DbName).Collection(CollectionName)

// 	return collection, nil
// }
