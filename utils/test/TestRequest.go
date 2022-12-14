package test

import (
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T, app *fiber.App, test *models.HttpTestCase) {
	res, err := app.Test(test.Req, -1)

	// Check response error
	assert.Equal(t, test.IfExpectError, err != nil)
	if test.IfExpectError && test.ExpectedError != nil {
		assert.ErrorIs(t, test.ExpectedError, err)
	}

	// Check response status
	assert.Equal(t, test.ExpectedStatus, res.StatusCode)

	if err == nil {
		// If no error, check response as well
		assert.Equal(t, test.ExpectedBody, utils.GetResponseBody(res))
	}
}
