package main

import (
	"net/http"
	"net/http/httptest"
	"pyncz/go-rest/middlewares"
	"pyncz/go-rest/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewares(t *testing.T) {
	tests := []models.TestCase{
		{
			Description:    "returns 404 Not Found on unhandled routes",
			Req:            httptest.NewRequest("GET", "/not-found", nil),
			ExpectedStatus: http.StatusNotFound,
			ExpectedBody:   "Not found",
			ExpectedError:  nil,
		},
	}

	// Mock fiber app
	app := fiber.New()
	app.Use(middlewares.NotFound)

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			res, err := app.Test(test.Req, -1)

			// Check response error
			assert.ErrorIsf(t, err, test.ExpectedError, test.Description)
			assert.Equalf(t, test.ExpectedStatus, res.StatusCode, test.Description)

			// if err == nil {
			// 	// If no error, check response as well
			// 	assert.HTTPBody()
			// }
		})
	}
}
