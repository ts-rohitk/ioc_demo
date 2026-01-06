package db

import (
	"context"
	"fmt"
	"log"

	"goat/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*20)
	defer cancel()

	fmt.Println("mongo uri", config.Cfg.Get("mongo_db_url"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.Get("mongo_db_url")))
	if err != nil {
		panic(err)
	}

	log.Println("connection to database is succuessful")
	return client
}
