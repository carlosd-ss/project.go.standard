package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber"
	mcustomer "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/models/customer"
	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"
)

func Post(c *fiber.Ctx) {

	const (
		db             = "dbcustomers"
		collectionpost = "customerpost"
	)

	collection, err := mongoconf.GetMongoDbCollection(db, collectionpost)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var customer mcustomer.CustomerPost
	json.Unmarshal([]byte(c.Body()), &customer)

	res, err := collection.InsertOne(context.Background(), customer)
	if err != nil {
		c.Status(500).Send(err)
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
