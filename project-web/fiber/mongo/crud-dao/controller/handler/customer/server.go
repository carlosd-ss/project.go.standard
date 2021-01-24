package hcustomer

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber"
)

type Server struct {
	App *fiber.App
	Db  *mongo.Client
}
