package assert

import "testing"

func Equal[V comparable](t *testing.T, expected, actual V) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
