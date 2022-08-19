package setup

import (
	"context"
	"pyncz/go-rest/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() (*mongo.Database, func()) {
	dbUri := utils.GetEnv("MONGO_CONNECT_STRING")
	dbName := utils.GetEnv("MONGO_INITDB_DATABASE")

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(dbUri).SetAppName("api"),
	)
	if err != nil {
		panic(err)
	}
	disconnect := func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}

	db := client.Database(dbName)

	return db, disconnect
}
