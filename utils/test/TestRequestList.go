package test

import (
	"pyncz/go-rest/models"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRequestList(t *testing.T, app *fiber.App, tests *[]models.HttpTestCase) {
	for _, test := range *tests {
		t.Run(test.Description, func(t *testing.T) {
			TestRequest(t, app, &test)
		})
	}
}
