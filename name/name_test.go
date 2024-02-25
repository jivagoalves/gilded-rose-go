package name

import (
	"gilded-rose/assert"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("name should not be blank", func(t *testing.T) {
		n, err := New("")
		assertNil(t, n)
		assert.Error(t, ErrBlankName{}, err)

		n, err = New(" ")
		assertNil(t, n)
		assert.Error(t, ErrBlankName{}, err)
	})
}

func assertNil(t *testing.T, n N) {
	t.Helper()
	if n != nil {
		t.Errorf("expected %v to be nil", n)
	}
}
