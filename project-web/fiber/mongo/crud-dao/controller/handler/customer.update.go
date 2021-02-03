package handler

import (
	"context"
	"encoding/json"
	"log"

	mcustomer "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/models/customer"
	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Conn) Update(c *fiber.Ctx) {
	const (
		db             = "dbcustomers"
		collectionpost = "customerpost"
	)
	collection, err := mongoconf.GetMongoDbCollection(db, collectionpost)
	if err != nil {
		c.Status(400).Send(err)
		return
	}
	var customer mcustomer.CustomerPost
	json.Unmarshal([]byte(c.Body()), &customer)

	update := bson.M{
		"$set": customer,
	}

	objID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		c.Status(400).Send(err)
		return

	}
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(400).Send(err)
		return
	}

	response, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	c.Send(response)
	c.Status(200)
}
