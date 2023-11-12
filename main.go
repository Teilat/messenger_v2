package main

import (
	"Messenger v2/messenger_v2/db"
	"log"
)

func main() {
	dbClient := db.NewDatabase(log.New(log.Writer(), "[db]", 0), "")
}
