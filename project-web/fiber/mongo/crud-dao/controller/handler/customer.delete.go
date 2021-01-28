package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type sError struct {
	Msg string `json:"msg"`
}

var serror sError

func (s Conn) Delete(c *fiber.Ctx) {

	c.Set("Content-Type", "application/json")

	if len(c.Params("id")) <= 0 {
		serror.Msg = "Id é obrigatorio "
		c.Status(400).JSON(serror)
		return
	}

	Delete(Id, s.Collection1)

	c.Status(200).SendString("")
	return
}

/// REPO
func Delete(Id string, collectionpost *mongo.Collection) (err error) {

	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		serror.Msg = "Usuário não pode ser deletado!"
		c.Status(400).JSON(serror)
		return
	}

	collection, err := mongoconf.GetMongoDbCollection(db, collectionpost)

	if err != nil {
		serror.Msg = "Id é obrigatorio "
		c.Status(400).JSON(serror)
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

	return
}
