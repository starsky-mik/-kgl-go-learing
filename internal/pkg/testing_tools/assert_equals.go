package testing_tools

import "testing"

func AssertEquals(t testing.TB, actual, expected any) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}
