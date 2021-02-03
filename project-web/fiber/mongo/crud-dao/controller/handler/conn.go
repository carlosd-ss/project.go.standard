package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Conn struct {
	Client         *mongo.Client
	Db             string
	CollectionName string
}
