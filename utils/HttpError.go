package utils

type HttpError struct {
	Message string `json:"message"`
} // @name HttpError

type HttpBadRequestError map[string]string // @name HttpBadRequestError

func NewError(message string) *HttpError {
	return &HttpError{message}
}
