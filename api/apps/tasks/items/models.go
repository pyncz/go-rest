package items

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Task        primitive.ObjectID `bson:"task,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Estimation  int32              `bson:"estimation,omitempty"`
}
