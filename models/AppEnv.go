package models

import "go.mongodb.org/mongo-driver/mongo"

type AppEnv struct {
	DB *mongo.Database
}
