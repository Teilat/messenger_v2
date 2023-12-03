package main

import (
	"context"
	"log"
	"messenger_v2/db"
)

func main() {
	ctx, _ := context.WithCancel(context.TODO())
	dbClient := db.NewDatabase(ctx, log.New(log.Writer(), "[db]", 0), "")
	dbClient.Init()
}
