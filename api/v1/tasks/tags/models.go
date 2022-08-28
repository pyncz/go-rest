package tags

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tag model
// @Description Tag data
type Tag struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty" validate:"required,len=24"`          // ID in the DB collection
	Name string             `json:"name" bson:"name,omitempty" validate:"required,min=3,max=64"` // Name of the tag
	Slug string             `json:"slug" bson:"slug,omitempty" validate:"required,min=3,max=32"` // Slug of the tag
} // @name Tag

// Tag filters model
// @Description Tag read filters
type TagFilters struct{} // @name TagFilters

// Tag create model
type TagCreateForm struct {
	Name string `json:"name" bson:"name,omitempty" validate:"required,min=3,max=64"` // Name of the tag
	Slug string `json:"slug" bson:"slug,omitempty" validate:"required,min=3,max=32"` // Slug of the tag
} // @name TagCreateForm
