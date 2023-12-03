package db

import "github.com/google/uuid"

const ChatCollection = "Chats"

type Chat struct {
	Users [2]uuid.UUID `bson:"users"`
}
