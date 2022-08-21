package models

import "net/http"

type TestCase struct {
	Description   string
	IfExpectError bool
	ExpectedError error
}

type ValidationTestCase[T any] struct {
	Description   string
	Input         *T
	IfExpectError bool
	ExpectedError error
}

type HttpTestCase struct {
	Description    string
	IfExpectError  bool
	ExpectedError  error
	Req            *http.Request
	ExpectedStatus int
	ExpectedBody   string
}
