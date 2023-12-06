package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
}

type database struct {
	log *log.Logger
	db  *mongo.Database
}

func NewDatabase(ctx context.Context, log *log.Logger, url string) Database {
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(url).SetAuth(
			options.Credential{
				Username: "admin",
				Password: "admin",
			}))
	if err != nil {
		panic(err)
	}

	return &database{
		log: log,
		db:  client.Database("messenger"),
	}
}
