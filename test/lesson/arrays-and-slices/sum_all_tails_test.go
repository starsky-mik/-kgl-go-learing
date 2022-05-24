package arrays_and_slices

import (
	"fmt"
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/arrays-and-slices"
	"testing"
)

type SumAllTailsCase struct {
	arrays [][]int
	result []int
}

func TestSumAllTails(t *testing.T) {
	testCases := []SumAllTailsCase{
		{[][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}, []int{5, 7, 9}},
		{[][]int{{1, 2, 3}, {2, 3, 4}, {3, 0, 0}}, []int{5, 7, 0}},
		{[][]int{{1, 2, 3}, {2, 3, 4}, {3, -1, 1}}, []int{5, 7, 0}},
		{[][]int{{1, 2, 3}, {2, 3, 4}, {5}}, []int{5, 7, 0}},
		{[][]int{{1, 2, 3}, {2, 3, 4}, {}}, []int{5, 7, 0}},
		{[][]int{{}, {}, {}}, []int{0, 0, 0}},
	}

	for _, testCase := range testCases {
		t.Run("check arrays_and_slices.SumAllTails function", func(t *testing.T) {
			AssertDeepEquals(t, SumAllTails(testCase.arrays...), testCase.result)
		})
	}

	t.Run("check arrays_and_slices.SumAllTails function with empty slices list", func(t *testing.T) {
		actual := SumAllTails()
		var expected []int

		AssertDeepEquals(t, actual, expected)
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
