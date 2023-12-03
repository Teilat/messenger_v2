package db

import "github.com/google/uuid"

const GroupCollection = "Groups"

type Group struct {
	Name  string      `bson:"name"`
	Users []uuid.UUID `bson:"users"`
}
