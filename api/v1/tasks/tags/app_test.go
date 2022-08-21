package tags

import (
	"pyncz/go-rest/models"
	"pyncz/go-rest/utils"
	"pyncz/go-rest/utils/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	tests := []models.ValidationTestCase[Tag]{
		{
			Description: "should pass on completed object",
			Input: &Tag{
				Name: "Test tag",
				Slug: "Test tag's slug",
			},
			IfExpectError: false,
		},

		// Name checks
		{
			Description: "should fail on the object without 'Name'",
			Input: &Tag{
				Slug: "Test tag's slug",
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Name' len < 3",
			Input: &Tag{
				Name: mock.MockString(2),
				Slug: "Test tag's slug",
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Name' len > 64",
			Input: &Tag{
				Name: mock.MockString(65),
				Slug: "Test tag's slug",
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with empty 'Name'",
			Input: &Tag{
				Name: "",
				Slug: "Test tag's slug",
			},
			IfExpectError: true,
		},

		// Slug checks
		{
			Description: "should fail on the object without 'Slug'",
			Input: &Tag{
				Name: "Test tag",
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Slug' len < 3",
			Input: &Tag{
				Name: "Test tag",
				Slug: mock.MockString(2),
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with 'Slug' len > 32",
			Input: &Tag{
				Name: "Test tag",
				Slug: mock.MockString(33),
			},
			IfExpectError: true,
		},
		{
			Description: "should fail on the object with empty 'Slug'",
			Input: &Tag{
				Name: "Test tag",
				Slug: "",
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
