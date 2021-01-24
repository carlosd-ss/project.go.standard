package mongo

import (
	"context"
	"sync"
	"time"

	"github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/fmts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectOnce sync.Once
var connectOnce2 sync.Once

var (
	session     *mongo.Client
	TimeContext = time.Duration(2 * time.Minute)
	err         error
)
var (
	MgoSrv     = "mongodb+srv"
	MgoUser    = "user.here"
	MgoPass    = "senhahere.."
	MgoCluster = "cluster0-mongo"
	MgoDb      = "test"
	MgoRetry   = "?retryWrites=true&w=majority"
)
var (
	client *mongo.Client
	ctx    context.Context
)

func connectDefault() string {
	return fmts.Concat(MgoSrv, "://", MgoUser, ":", MgoPass, "@", MgoCluster, "/", MgoDb, MgoRetry)
}

func connectClient() (*mongo.Client, error) {
	connectOnce.Do(func() {
		if session == nil {
			session, err = mongo.NewClient(
				options.Client().
					ApplyURI(connectDefault()))
			if err != nil {
				return
			}
		}
	})
	return session, err
}

func Connect() (*mongo.Client, context.Context, error) {
	connectOnce2.Do(func() {
		if client == nil {
			client, err = connectClient()
			if err != nil {
				return
			}
			ctx, _ = context.WithTimeout(context.Background(), TimeContext)
			err := client.Connect(ctx)
			if err != nil {(
				return
			}
		}
	})
	return client, ctx
}
