package tasks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Title     string               `json:"title" bson:"title,omitempty" validate:"required,min=3,max=64"`
	Completed *bool                `json:"completed" bson:"completed,omitempty" validate:"required"`
	Tags      []primitive.ObjectID `json:"tags" bson:"tags,omitempty"`
}
