package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func assertEquals(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected '%q' but got '%q'", expected, actual)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("empty char repeating", func(t *testing.T) {
		assertEquals(t, Repeat("", 5), "")
		assertEquals(t, Repeat("", 5), strings.Repeat("", 5))
	})

	t.Run("char repeating", func(t *testing.T) {
		assertEquals(t, Repeat("m", 5), "mmmmm")
		assertEquals(t, Repeat("m", 5), strings.Repeat("m", 5))
	})

	t.Run("multi-char string repeating", func(t *testing.T) {
		assertEquals(t, Repeat("xyz", 3), "xyzxyzxyz")
		assertEquals(t, Repeat("xyz", 5), strings.Repeat("xyz", 5))
	})

	t.Run("char repeating 0 times", func(t *testing.T) {
		assertEquals(t, Repeat("m", 0), "")
		assertEquals(t, Repeat("m", 0), strings.Repeat("m", 0))
	})

	t.Run("char repeating negative times", func(t *testing.T) {
		assertEquals(t, Repeat("m", -5), "")
	})
}

func ExampleRepeat() {
	repeated := Repeat("w", 3)
	fmt.Println(repeated)
	// Output: www
}
