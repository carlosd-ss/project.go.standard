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

func GetAll(c *fiber.Ctx) {
	const (
		db             = "dbcustomers"
		collectionpost = "customerpost"
	)
	collection, err := mongoconf.GetMongoDbCollection(db, collectionpost)
	if err != nil {
		c.Status(400).Send(err)
		return
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println(err)
			return
		}
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.Status(400).Send(err)
		return
	}
	defer cur.Close(context.Background())

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(400)
		return
	}

	json, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
		return
	}
	c.Send(json)
	c.Status(200)
}
