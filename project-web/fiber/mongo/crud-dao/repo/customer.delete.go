package repo

import (
	"context"
	"errors"
	"log"

	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(Id string, db string, CollectionName string) (err error) {

	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return errors.New("Usuário não pode ser deletado!")
	}

	collection := mongoconf.GetCollection(db, CollectionName)

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		log.Println(err)
		return
	}

	return
}
