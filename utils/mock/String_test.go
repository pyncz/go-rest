package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	t.Run("should return empty string on length=0", func(t *testing.T) {
		assert.Equal(t, "", MockString(0))
	})

	t.Run("should return N-len-string on length=N for odd numbers", func(t *testing.T) {
		var length uint = 5
		mocked := MockString(length)
		assert.Equal(t, length, uint(len(mocked)))
	})

	t.Run("should return N-len-string on length=N for even numbers", func(t *testing.T) {
		var length uint = 8
		mocked := MockString(length)
		assert.Equal(t, length, uint(len(mocked)))
	})
}
