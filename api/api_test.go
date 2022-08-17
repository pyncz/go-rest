package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/v3/assert"
)

type TestCase struct {
	name           string
	in             *http.Request
	out            *httptest.ResponseRecorder
	expectedStatus int
	expectedBody   string
}

func TestPing(t *testing.T) {
	tests := []TestCase{
		{
			name:           "Health Check",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			// TODO: Mock DB?
			handler := Ping(nil)

			// Mock gin context
			ctx, _ := gin.CreateTestContext(test.out)
			ctx.Request = test.in
			ctx.Params = []gin.Param{
				{Key: "limit", Value: "12"},
			}

			handler(ctx)

			assert.Equal(t, test.expectedStatus, test.out.Code)
			assert.Equal(t, test.expectedBody, test.out.Body.String())
		})
	}
}
