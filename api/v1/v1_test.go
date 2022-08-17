package api_v1

import (
	"net/http"
	"net/http/httptest"
	"pyncz/go-rest/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	tests := []models.TestCase{
		{
			Description:    "responses with 200 OK on Health Check route",
			Req:            httptest.NewRequest("GET", "/ping", nil),
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   "",
			ExpectedError:  nil,
		},
	}

	// Mock fiber app
	app := fiber.New()
	app.Mount("/", App(nil))

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
