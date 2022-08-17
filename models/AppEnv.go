package models

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppEnv struct {
	DB  *mongo.Database
	Log *log.Logger
}
