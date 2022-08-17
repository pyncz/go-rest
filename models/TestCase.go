package models

import "net/http"

type TestCase struct {
	Description    string
	Req            *http.Request
	ExpectedStatus int
	ExpectedBody   string
	ExpectedError  error
}
