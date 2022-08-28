package tasks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task model
// @Description Task data
type Task struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty" validate:"required,len=24"`            // ID in the DB collection
	Title     string               `json:"title" bson:"title,omitempty" validate:"required,min=3,max=64"` // Title of the task
	Completed bool                 `json:"completed" bson:"completed,omitempty"`                          // If the task is completed
	Tags      []primitive.ObjectID `json:"tags" bson:"tags,omitempty" validate:"dive,len=24"`             // Task's related tags' IDs
} // @name Task

// Task filters model
// @Description Task read filters
type TaskFilters struct{} // @name TaskFilters

// Task create model
// @Description Task creation model
type TaskCreateForm struct {
	Title     string               `json:"title" bson:"title,omitempty" validate:"required,min=3,max=64"` // Title of the task
	Completed bool                 `json:"completed" bson:"completed,omitempty"`                          // If the task is completed
	Tags      []primitive.ObjectID `json:"tags" bson:"tags,omitempty"`                                    // Task's related tags' IDs
} // @name TaskCreateForm
