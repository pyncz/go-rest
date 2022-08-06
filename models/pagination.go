package models

type PaginatedResponse[TData any] struct {
	Count   int     `json:"count" bson:"count,omitempty"`
	Cursor  int     `json:"cursor" bson:"cursor,omitempty"`
	Results []TData `json:"results" bson:"results,omitempty"`
}

var DEFAULT_LIMIT int64 = 0
