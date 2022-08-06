package tasks

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`
	Title     string               `bson:"title,omitempty"`
	Completed bool                 `bson:"completed,omitempty"`
	Tags      []primitive.ObjectID `bson:"tags,omitempty"`
}
