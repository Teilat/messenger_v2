package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Read()
	Write()
	Update()
}

type database struct {
	log    *log.Logger
	url    string
	client interface{}
}

func NewDatabase(log *log.Logger, url string) Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &database{
		log:    log,
		url:    url,
		client: client,
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
