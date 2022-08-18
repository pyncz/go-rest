package tasks

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Title     string               `json:"title" bson:"title,omitempty"`
	Completed bool                 `json:"completed" bson:"completed,omitempty"`
	Tags      []primitive.ObjectID `json:"tags" bson:"tags,omitempty"`
}