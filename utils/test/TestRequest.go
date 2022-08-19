package test

import (
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T, app *fiber.App, test *models.TestCase) {
	res, err := app.Test(test.Req, -1)

	// Check response error
	assert.ErrorIs(t, err, test.ExpectedError)
	assert.Equal(t, test.ExpectedStatus, res.StatusCode)

	if err == nil {
		// If no error, check response as well
		assert.Equal(t, test.ExpectedBody, utils.GetResponseBody(res))
	}
}
