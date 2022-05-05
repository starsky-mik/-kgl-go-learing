package main

import (
	"fmt"
	"testing"
)

func assertEquals(t testing.TB, actual, expected int) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", actual, expected)
	}
}

func TestAdd(t *testing.T) {
	t.Run("Check sum", func(t *testing.T) {
		assertEquals(t, Add(-1, 5), 4)
	})

	t.Run("Check sum", func(t *testing.T) {
		assertEquals(t, Add(0, 4), 4)
	})

	t.Run("Check sum", func(t *testing.T) {
		assertEquals(t, Add(2, 2), 4)
	})

	t.Run("Check sum", func(t *testing.T) {
		assertEquals(t, Add(5, -1), 4)
	})
}

func ExampleAdd() {
	sum := Add(3, 6)
	fmt.Println(sum)
	// Output 9
}
