package integers

import (
	"fmt"
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	"github.com/starsky-mik/kgl-go-learing/lesson/integers"
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
			testingTools.AssertEquals(t, integers.Add(testCase.a, testCase.b), testCase.result)
		})
	}
}

func ExampleAdd() {
	sum := integers.Add(3, 6)
	fmt.Println(sum)
	// Output 9
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integers.Add(5, 6)
	}
}
