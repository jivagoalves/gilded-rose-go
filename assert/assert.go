package assert

import (
	"errors"
	"testing"
)

func Equal[V comparable](t *testing.T, expected, actual V) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func Error(t *testing.T, expected error, actual error) {
	t.Helper()
	if !errors.Is(expected, actual) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
