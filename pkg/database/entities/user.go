package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `faker:"{username}"`
	IsOnline bool               `faker:"{boolean}"`
}
