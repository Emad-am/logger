package logmodel

import (
	"logger/internal/context"
	"logger/src/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func getModel() *mongo.Collection {
	return models.GetDatabase().Collection("log")
}

func (l Log) Save() (*mongo.InsertOneResult, error) {
	ctx := context.GetContext()
	return getModel().InsertOne(*ctx, l)
}
