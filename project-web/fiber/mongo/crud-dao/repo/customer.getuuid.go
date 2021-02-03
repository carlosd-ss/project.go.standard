package repo

import (
	"context"
	"errors"

	mongoconf "github.com/project.go.standard/project-web/fiber/mongo/crud-dao/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Getid(Id string, db string, CollectionName string) (err error) {

	objID, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return errors.New("Usuário não pode ser encontrado!")
	}

	collection := mongoconf.GetCollection(db, CollectionName)
	collection.FindOne(context.Background(), bson.M{"_id": objID})

	return nil
}
