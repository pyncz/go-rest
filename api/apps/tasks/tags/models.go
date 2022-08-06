package tags

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Slug string             `bson:"slug,omitempty"`
	Name string             `bson:"slug,omitempty"`
}
