package models

type PaginationQuery struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

type PaginatedResponse[TData any] struct {
	Count   int64   `json:"count" bson:"count,omitempty"`
	Cursor  int64   `json:"cursor" bson:"cursor,omitempty"`
	Offset  int64   `json:"offset" bson:"offset,omitempty"`
	Limit   int64   `json:"limit" bson:"limit,omitempty"`
	Results []TData `json:"results" bson:"results,omitempty"`
}

var DEFAULT_LIMIT int64 = 12
