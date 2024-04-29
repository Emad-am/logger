package db

import (
	"fmt"
	"log"

	"logger/config"
	context "logger/internal/context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Client mongo.Client
}

var client *mongo.Client

func init() {

	cnfg := config.GetConfig()

	host := cnfg.Db.Host
	port := cnfg.Db.Port
	userName := cnfg.Db.UserName
	pass := cnfg.Db.Password

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s", userName, pass, host, port)

	ctx := context.GetContext()

	var err error
	client, err = mongo.Connect(*ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	err = client.Ping(*ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")
}

func GetMongoDBClient() *MongoDBClient {
	return &MongoDBClient{
		Client: *client,
	}
}

func Disconnect() {
	ctx := context.GetContext()
	if err := client.Disconnect(*ctx); err != nil {
		panic(err)
	}
}
