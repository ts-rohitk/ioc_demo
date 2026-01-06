package queue

import (
	"goat/config"
	"goat/db"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MongoClient *mongo.Client
)

func init() {
	config.Load()
	MongoClient = db.Connect()
}
