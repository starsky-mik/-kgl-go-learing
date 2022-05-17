package testing_tools

import (
	"reflect"
	"testing"
)

func AssertDeepEquals(t testing.TB, actual, expected any) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected '%v' but got '%v", expected, actual)
	}
}
