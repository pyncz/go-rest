package items

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Task        primitive.ObjectID `json:"task" bson:"task,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Estimation  int32              `json:"estimation" bson:"estimation,omitempty"`
}
