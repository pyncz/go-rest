package models

type PaginationQuery struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

type ListResults[TData any] struct {
	Count   int64   `json:"count" bson:"count,omitempty"`
	Results []TData `json:"results" bson:"results,omitempty"`
}

type PaginatedListResults[TData any] struct {
	Count   int64   `json:"count" bson:"count,omitempty"`
	Results []TData `json:"results" bson:"results,omitempty"`
	Cursor  int64   `json:"cursor" bson:"cursor,omitempty"`
	Offset  int64   `json:"offset" bson:"offset,omitempty"`
	Limit   int64   `json:"limit" bson:"limit,omitempty"`
}

var DEFAULT_LIMIT int64 = 12
