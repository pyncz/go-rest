package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Filters struct {
	_id        primitive.ObjectID
	created_at primitive.DateTime
	updated_at primitive.DateTime
}
