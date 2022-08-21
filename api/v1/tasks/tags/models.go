package tags

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty" validate:"required,min=3,max=64"`
	Slug string             `json:"slug" bson:"slug,omitempty" validate:"required,min=3,max=32"`
}
