package utils

import (
	"context"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDb() func() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	dbUri := GetEnv("MONGODB_URI")
	dbName := GetEnv("DB_NAME")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}
	Disconnect := func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}

	DB = client.Database(dbName)

	return Disconnect
}
