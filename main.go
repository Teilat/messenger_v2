package main

import (
	"context"
	"fmt"
	"log"
	"messenger_v2/db"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	dbLog := log.New(log.Writer(), "[db]", 0)
	db.NewDatabase(ctx, dbLog, "mongodb://localhost:27017")
	for true {
		fmt.Scanln()
		cancel()
	}
}
