package iteration

import (
	"fmt"
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/iteration"
	"strings"
	"testing"
)

type RepeatTestCase struct {
	word   string
	count  int
	result string
}

func TestRepeat(t *testing.T) {
	testCases := []RepeatTestCase{
		{"", 3, ""},
		{"x", 3, "xxx"},
		{"xyz", 3, "xyzxyzxyz"},
		{"x", 0, ""},
	}

	for _, testCase := range testCases {
		t.Run("check iteration.Repeat function and compare result with strings.Repeat", func(t *testing.T) {
			AssertEquals(t, Repeat(testCase.word, testCase.count), testCase.result)
			AssertEquals(t, Repeat(testCase.word, testCase.count), strings.Repeat(testCase.word, testCase.count))
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("w", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("w", 3)
	fmt.Println(repeated)
	// Output: www
}
