package main

import (
	"context"

	"github.com/Binaretech/chattimus/config"
	"github.com/Binaretech/chattimus/pkg/database"
	"github.com/Binaretech/chattimus/pkg/database/entities"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config.InitConfig()
	database := database.Database()
	createUsers(database.Collection("users"))
}

func createUsers(collection *mongo.Collection) {
	users := make([]interface{}, 50)

	for i := 0; i < 50; i++ {
		user := entities.User{}
		gofakeit.Struct(&user)
		users[i] = user
	}

	if _, err := collection.InsertMany(context.Background(), users); err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Created %d users\n", len(users))
}
