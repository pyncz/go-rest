package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task item model
// @Description Item data
type Item struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty" validate:"required,len=24"`                       // ID in the DB collection
	Task        primitive.ObjectID `json:"task" bson:"task,omitempty" validate:"required,len=24"`                    // Related task's ID
	Title       string             `json:"title" bson:"title,omitempty" validate:"required,min=3,max=64"`            // Title of the task item
	Description string             `json:"description" bson:"description,omitempty" validate:"max=256"`              // Description of the task item
	Estimation  uint8              `json:"estimation" bson:"estimation,omitempty" validate:"required,gte=1,lte=100"` // Estimation points of the task item
} // @name Item

// Task item filters
// @Description Task item's read filters
type ItemFilters struct{} // @name ItemFilters

// Task create model
// @Description Task creation model
type ItemCreateForm struct {
	Task        primitive.ObjectID `json:"task" bson:"task,omitempty" validate:"required,len=24"`                    // Related task's ID
	Title       string             `json:"title" bson:"title,omitempty" validate:"required,min=3,max=64"`            // Title of the task item
	Description string             `json:"description" bson:"description,omitempty" validate:"max=256"`              // Description of the task item
	Estimation  uint8              `json:"estimation" bson:"estimation,omitempty" validate:"required,gte=1,lte=100"` // Estimation points of the task item
} // @name ItemCreateForm
