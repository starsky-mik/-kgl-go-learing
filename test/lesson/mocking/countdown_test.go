package mocking

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/mocking"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	var output = SpyCountdownOutputProxy{Interval: 1 * time.Second}

	Countdown(&output)

	AssertDeepEquals(t, output.Log(), []string{
		"Wait 1s",
		"Write line: \"3\"",
		"Wait 1s",
		"Write line: \"2\"",
		"Wait 1s",
		"Write line: \"1\"",
		"Wait 1s",
		"Write line: \"Go!\"",
	})

	expOut := `3
2
1
Go!
`

	AssertEquals(t, output.Out(), expOut)
}
