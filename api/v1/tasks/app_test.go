package tasks

import (
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"
	"pyncz/go-rest/utils/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	tests := []models.ValidationTestCase[Task]{
		{
			Description: "should pass on completed object with 'true' booleans",
			Input: &Task{
				Title:     "Test task",
				Completed: true,
			},
			IfExpectError: false,
		},
		{
			Description: "should pass on completed object with 'false' booleans",
			Input: &Task{
				Title:     "Test task",
				Completed: false,
			},
			IfExpectError: false,
		},

		// Completed checks
		{
			Description: "should pass on the object without 'Completed'",
			Input: &Task{
				Title: "Test task",
			},
			IfExpectError: false,
		},

		// Title checks
		{
			Description: "should fail on the object without 'Title'",
			Input: &Task{
				Completed: true,
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Title' len < 3",
			Input: &Task{
				Title:     mock.MockString(2),
				Completed: true,
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Title' len > 64",
			Input: &Task{
				Title:     mock.MockString(65),
				Completed: true,
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with empty 'Title'",
			Input: &Task{
				Title:     "",
				Completed: true,
			},
			IfExpectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			_, err := utils.Validate(test.Input)

			assert.Equal(t, test.IfExpectError, err != nil)
		})
	}
}
