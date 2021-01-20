package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Content  string
	Sender   User
	Receiver User
}
