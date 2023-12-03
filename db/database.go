package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Init(ctx context.Context)
	Read()
	Write()
	Update()
}

type database struct {
	log    *log.Logger
	url    string
	client *mongo.Client
}

var Collections = map[string]bool{
	ChatCollection:    false,
	GroupCollection:   false,
	MessageCollection: false,
}

func NewDatabase(ctx context.Context, log *log.Logger, url string) Database {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return &database{
		log:    log,
		url:    url,
		client: client,
	}
}

func (d *database) Init(ctx context.Context) {
	existingCollections, err := d.client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		return
	}
	for _, collection := range existingCollections {
		Collections[collection] = true
	}
	for s, exist := range Collections {
		if !exist {
			fmt.Println(s)
			// create collection
		}
	}
}

func (d *database) Read() {
	//TODO implement me
	panic("implement me")
}

func (d *database) Write() {
	//TODO implement me
	panic("implement me")
}

func (d *database) Update() {
	//TODO implement me
	panic("implement me")
}
