package testing_tools

import "testing"

func AssertError(t testing.TB, err, expected error) {
	t.Helper()

	if err == nil {
		t.Errorf("expected error \"%v\" but got \"%v\"", expected, err)
	}

	if err != expected {
		t.Errorf("expected error \"%v\" but got \"%v\"", expected, err)
	}
}
