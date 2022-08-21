package main

import (
	"net/http"
	"net/http/httptest"
	"pyncz/go-rest/middlewares"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils/test"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestMiddlewares(t *testing.T) {
	tests := []models.HttpTestCase{
		{
			Description:    "returns 404 Not Found on unhandled routes",
			Req:            httptest.NewRequest("GET", "/not-found", nil),
			ExpectedStatus: http.StatusNotFound,
			ExpectedBody:   "Not Found",
		},
	}

	// Mock fiber app
	app := fiber.New()
	app.Use(middlewares.NotFound)

	test.TestRequestList(t, app, &tests)
}
