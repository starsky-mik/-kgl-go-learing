package integers

import (
	"fmt"
	. "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgl-go-learing/lesson/integers"
	"testing"
)

type AddTestCase struct {
	a, b   int
	result int
}

func TestAdd(t *testing.T) {
	testCases := []AddTestCase{
		{1, 5, 6},
		{-1, 5, 4},
		{1, -1, 0},
		{0, 5, 5},
		{1, 0, 1},
		{0, 0, 0},
	}

	for _, testCase := range testCases {
		t.Run("check integers.Add function", func(t *testing.T) {
			AssertEquals(t, Add(testCase.a, testCase.b), testCase.result)
		})
	}
}

func ExampleAdd() {
	sum := Add(3, 6)
	fmt.Println(sum)
	// Output 9
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(5, 6)
	}
}
