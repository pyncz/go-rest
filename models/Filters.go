package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Filters struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}
