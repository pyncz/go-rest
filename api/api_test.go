package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	description    string
	req            *http.Request
	expectedStatus int
	expectedBody   string
	expectedError  error
}

func TestRoutes(t *testing.T) {
	tests := []TestCase{
		{
			description:    "returns 404 Not Found on unhandled routes",
			req:            httptest.NewRequest("GET", "/not-found", nil),
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Not found",
			expectedError:  nil,
		},
	}

	// Mock fiber app
	app := fiber.New()
	App(app, nil)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			res, err := app.Test(test.req, -1)

			// Check response error
			assert.ErrorIsf(t, err, test.expectedError, test.description)
			assert.Equalf(t, test.expectedStatus, res.StatusCode, test.description)

			// if err == nil {
			// 	// If no error, check response as well
			// 	assert.HTTPBody()
			// }
		})
	}
}

func TestHealthCheck(t *testing.T) {
	tests := []TestCase{
		{
			description:    "responses with 200 OK on Health Check route",
			req:            httptest.NewRequest("GET", "/ping", nil),
			expectedStatus: http.StatusOK,
			expectedBody:   "",
			expectedError:  nil,
		},
	}

	// Mock fiber app
	app := fiber.New()
	App(app, nil)

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			res, err := app.Test(test.req, -1)

			// Check response error
			assert.ErrorIsf(t, err, test.expectedError, test.description)
			assert.Equalf(t, test.expectedStatus, res.StatusCode, test.description)

			// if err == nil {
			// 	// If no error, check response as well
			// 	assert.HTTPBody()
			// }
		})
	}
}
