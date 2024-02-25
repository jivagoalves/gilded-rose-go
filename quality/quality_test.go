package quality

import (
	"errors"
	"testing"
)

func TestQuality(t *testing.T) {
	t.Run("quality should never be less than zero", func(t *testing.T) {
		q, err := New(-1)

		assertNil(t, q)
		assertError(t, ErrNegativeQuality{}, err)
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

func assertError(t *testing.T, expected error, actual error) {
	t.Helper()
	if !errors.Is(expected, actual) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func assertNil(t *testing.T, q Q) {
	t.Helper()
	if q != nil {
		t.Errorf("expected %v to be nil", q)
	}
}
