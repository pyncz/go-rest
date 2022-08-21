package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Task        primitive.ObjectID `json:"task" bson:"task,omitempty" validate:"required"`
	Title       string             `json:"title" bson:"title,omitempty" validate:"required,min=3,max=64"`
	Description string             `json:"description" bson:"description,omitempty" validate:"max=256"`
	Estimation  uint8              `json:"estimation" bson:"estimation,omitempty" validate:"required,gte=1,lte=100"`
}

type ItemFilters struct{}
