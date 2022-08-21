package items

import (
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"
	"pyncz/go-rest/utils/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidation(t *testing.T) {
	tests := []models.ValidationTestCase[Item]{
		{
			Description: "should pass on completed object",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       "Test item",
				Description: "Test item's description",
				Estimation:  1,
			},
			IfExpectError: false,
		},

		// Task checks
		{
			Description: "should fail on the object without 'Task' ID",
			Input: &Item{
				Title:       "Test item",
				Description: "Test item's description",
				Estimation:  1,
			},
			IfExpectError: true,
		},

		// Title checks
		{
			Description: "should fail on the object without 'Title'",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Description: "Test item's description",
				Estimation:  1,
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Title' len < 3",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       mock.MockString(2),
				Description: "Test item's description",
				Estimation:  1,
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Title' len > 64",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       mock.MockString(65),
				Description: "Test item's description",
				Estimation:  1,
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with empty 'Title'",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       "",
				Description: "Test item's description",
				Estimation:  1,
			},
			IfExpectError: true,
		},

		// Description checks
		{
			Description: "should pass on the object without 'Description'",
			Input: &Item{
				Task:       primitive.NewObjectID(),
				Title:      "Test item",
				Estimation: 1,
			},
			IfExpectError: false,
		},
		{
			Description: "should fail on the object with 'Description' len > 256",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       "Test item",
				Description: mock.MockString(257),
				Estimation:  1,
			},
			IfExpectError: true,
		},
		{
			Description: "should pass on the object with empty 'Description'",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       "Test item",
				Description: "",
				Estimation:  1,
			},
			IfExpectError: false,
		},

		// Estimation checks
		{
			Description: "should fail on the object without 'Estimation'",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       "Test item",
				Description: "Test item's description",
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with empty 'Estimation'",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       "Test item",
				Description: "Test item's description",
				Estimation:  0,
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Estimation' > 100",
			Input: &Item{
				Task:        primitive.NewObjectID(),
				Title:       "Test item",
				Description: "Test item's description",
				Estimation:  101,
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
