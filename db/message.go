package db

import "github.com/google/uuid"

const MessageCollection = "Messages"

type Message struct {
	Id        uuid.UUID `bson:"id"`
	ChatId    uuid.UUID `bson:"chatId"`
	UserId    uuid.UUID `bson:"userId"`
	Text      string    `bson:"text"`
	Timestamp int64     `bson:"timestamp"`
}
