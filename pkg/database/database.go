package database

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_URL"),
	)
}

func Database() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString()))

	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		logrus.Fatal(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logrus.Fatal(err)
	}

	database := client.Database("chattimus")

	ConfigureDatabase(database)

	return database
}

func ConfigureDatabase(database *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	index := mongo.IndexModel{
		Keys: bson.M{
			"username": 1,
		},

		Options: options.Index().SetUnique(true),
	}

	if _, err := database.Collection("users").Indexes().CreateOne(ctx, index); err != nil {
		logrus.Fatal(err)
	}

	return nil
}
