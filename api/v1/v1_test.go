package api_v1

import (
	"net/http"
	"net/http/httptest"
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils/test"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHealthCheck(t *testing.T) {
	tests := []models.HttpTestCase{
		{
			Description:    "responses with 200 OK on Health Check route",
			Req:            httptest.NewRequest("GET", "/ping", nil),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   "OK",
		},
	}

	// Mock fiber app
	app := fiber.New()
	app.Mount("/", App(nil))

	test.TestRequestList(t, app, &tests)
}
