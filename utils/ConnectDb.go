package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() (*mongo.Database, func()) {
	dbUri := GetEnv("MONGO_CONNECT_STRING")
	dbName := GetEnv("MONGO_INITDB_DATABASE")

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(dbUri).SetAppName("api"),
	)
	if err != nil {
		panic(err)
	}
	Disconnect := func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}

	DB := client.Database(dbName)

	return DB, Disconnect
}
