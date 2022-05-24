package concurrency

import (
	"github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/concurrency"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://error" {
		return false
	}

	return true
}

func mockSlowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://no-error",
		"http://no-error-too",
		"http://error",
	}

	expected := map[string]bool{
		"http://no-error":     true,
		"http://no-error-too": true,
		"http://error":        false,
	}

	got := CheckWebsites(mockWebsiteChecker, urls)
	testing_tools.AssertDeepEquals(t, got, expected)
}

func BenchmarkCheckWebsites(b *testing.B) {
	var urls = make([]string, 100)

	for i := 0; i < 100; i++ {
		urls[i] = "url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CheckWebsites(mockSlowWebsiteChecker, urls)
	}
}
