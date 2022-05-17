package testing_tools

import "testing"

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("unexpected error \"%v\"", err)
	}
}
