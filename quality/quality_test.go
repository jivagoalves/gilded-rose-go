package quality

import (
	"gilded-rose/assert"
	"testing"
)

func TestQuality(t *testing.T) {
	t.Run("quality should never be less than zero", func(t *testing.T) {
		q, err := New(-1)

		assertNil(t, q)
		assert.Error(t, ErrNegativeQuality{}, err)
	})

	t.Run("quality should never be less than zero after decreasing", func(t *testing.T) {
		q, _ := New(0)

		AssertEqual(t, Zero, q.Dec())
	})

	t.Run("quality should never be less than zero after adding", func(t *testing.T) {
		q, _ := New(0)

		AssertEqual(t, Zero, q.Inc(-1))
	})
}

func assertNil(t *testing.T, q Q) {
	t.Helper()
	if q != nil {
		t.Errorf("expected %v to be nil", q)
	}
}
