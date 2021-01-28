package handler

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Conn struct {
	Client       *mongo.Client
	PingTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Collection   *mongo.Collection
	MgoDb        string
}
