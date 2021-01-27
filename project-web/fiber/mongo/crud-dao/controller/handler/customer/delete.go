package hcustomer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber"
	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(c *fiber.Ctx) {
	const (
		db             = "dbcustomers"
		collectionpost = "customerpost"
	)
	collection, err := mongoconf.GetMongoDbCollection(db, collectionpost)

	if err != nil {
		c.Status(400).Send(err)
		return
	}

	objID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		c.Status(400).Send(err)
		return
	}
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(400).Send(err)
		return
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	c.Send(jsonResponse)
	c.Status(200)
}
