package arrays_and_slices

import (
	"fmt"
	"reflect"
	"testing"
)

func assertEquals(t testing.TB, actual, expected int) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}

func assertDeepEquals(t testing.TB, actual, expected any) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected '%v' but got '%v", expected, actual)
	}
}

func TestSum(t *testing.T) {
	t.Run("test Sum function", func(t *testing.T) {
		assertEquals(t, Sum([]int{1, 2, 3, 4, 5}), 15)
	})
	t.Run("test Sum function with empty slice", func(t *testing.T) {
		assertEquals(t, Sum([]int{}), 0)
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output 15
}

func TestSumAll(t *testing.T) {
	t.Run("test SumAll function", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{2, 3}, []int{3, 4})
		expected := []int{3, 5, 7}

		assertDeepEquals(t, actual, expected)
	})

	t.Run("test SumAll function with one empty slice", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{2, 3}, []int{})
		expected := []int{3, 5, 0}

		assertDeepEquals(t, actual, expected)
	})

	t.Run("test SumAll function with all empty slices", func(t *testing.T) {
		actual := SumAll([]int{}, []int{}, []int{})
		expected := []int{0, 0, 0}

		assertDeepEquals(t, actual, expected)
	})

	t.Run("test SumAll function with empty slices list", func(t *testing.T) {
		actual := SumAll()
		var expected []int

		assertDeepEquals(t, actual, expected)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{2, 3}, []int{3, 4})
	}
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2}, []int{2, 3}, []int{3, 4})
	fmt.Println(sum)
	// Output [3 5 7]
}

func TestSumAllTails(t *testing.T) {
	t.Run("test TestSumAllTails function", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
		expected := []int{5, 7, 9}

		assertDeepEquals(t, actual, expected)
	})

	t.Run("test TestSumAllTails function with one slice without head", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4}, []int{5})
		expected := []int{5, 7, 0}

		assertDeepEquals(t, actual, expected)
	})

	t.Run("test TestSumAllTails function with one empty slice", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4}, []int{})
		expected := []int{5, 7, 0}

		assertDeepEquals(t, actual, expected)
	})

	t.Run("test TestSumAllTails function with all empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{}, []int{})
		expected := []int{0, 0, 0}

		assertDeepEquals(t, actual, expected)
	})

	t.Run("test TestSumAllTails function with empty slices list", func(t *testing.T) {
		actual := SumAllTails()
		var expected []int

		assertDeepEquals(t, actual, expected)
	})
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
	}
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
	fmt.Println(sum)
	// Output [5 7 9]
}
