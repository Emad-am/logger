package models

import (
	"logger/config"
	"logger/internal/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetDatabase() *mongo.Database {
	db := db.GetMongoDBClient()
	cnfg := config.GetConfig()
	return db.Client.Database(cnfg.Db.Name)
}
