package usermodel

import (
	"logger/internal/context"
	"logger/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getModel() *mongo.Collection {
	return models.GetDatabase().Collection("user")
}

func FindByUserNameAndPassword(userName string, password string) (User, error) {
	ctx := context.GetContext()

	filter := bson.D{{"username", userName}, {"password", password}}

	var l User

	err := getModel().FindOne(*ctx, filter).Decode(&l)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return l, err
		}
		panic(err)
	}
	return l, nil
}
