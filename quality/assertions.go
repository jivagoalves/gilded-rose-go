package quality

import (
	"gilded-rose/assert"
	"testing"
)

func AssertEqual(t *testing.T, expected, actual Q) {
	t.Helper()
	assert.Equal(t, expected.Value(), actual.Value())
}
